package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/qedus/osmpbf"

	"github.com/yasonofriychuk/real-estate-insight/interanal/osm/pbf_scanner"
)

const batchSize = 10000

func main() {
	ctx := context.Background()

	// Открытие OSM PBF файла
	file, err := os.Open("osmfiles/RU.osm.pbf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Подключение к базе данных
	conn, err := pgx.Connect(ctx, "postgres://osmuser:osmpassword@localhost:5432/osm")
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	defer conn.Close(ctx)

	// Инициализация сканера
	sc, err := pbf_scanner.New(file, runtime.GOMAXPROCS(-1))
	if err != nil {
		panic(err)
	}

	// Накопление узлов для батча
	var total uint64
	nodesBatch := make([]osmpbf.Node, 0, batchSize)
	for {
		node, err := sc.Next()
		if err != nil {
			break
		}

		nodesBatch = append(nodesBatch, node)
		if len(nodesBatch) >= batchSize {
			if err := insertBatch(ctx, conn, nodesBatch); err != nil {
				log.Fatalf("Failed to insert batch: %v\n", err)
			}
			total += uint64(len(nodesBatch))
			nodesBatch = nodesBatch[:0]
		}
	}

	if len(nodesBatch) > 0 {
		if err := insertBatch(ctx, conn, nodesBatch); err != nil {
			fmt.Printf("Failed to insert final batch: %v\n", err)
		}
	}

	fmt.Println("OSM nodes have been successfully scanned and inserted into the database.")
}
func hQuote(str string) string {
	str = strings.Replace(str, "\\", "\\\\", -1)
	return `"` + strings.Replace(str, "\"", "\\\"", -1) + `"`
}

func argsByNode(node osmpbf.Node) []any {
	tagParts := make([]string, 0, len(node.Tags))
	for k, v := range node.Tags {
		if k == "" || v == "" {
			continue
		}
		tagParts = append(tagParts, hQuote(k)+" => "+hQuote(v))
	}
	tags := strings.Join(tagParts, ",")

	return []any{node.ID, node.Tags["name"], tags, node.Lon, node.Lat}
}

// insertBatch выполняет вставку данных в таблицу батчем с использованием squirrel и поддержкой upsert
func insertBatch(ctx context.Context, conn *pgx.Conn, batch []osmpbf.Node) error {
	query := `
		INSERT INTO osm_node (osm_id, name, tags, way)
		VALUES ($1, $2, $3::hstore, ST_SetSRID(ST_MakePoint($4, $5), 4326))
		ON CONFLICT (osm_id) DO UPDATE SET
			name = EXCLUDED.name,
			tags = EXCLUDED.tags,
			way = EXCLUDED.way;
	`

	batchQueue := &pgx.Batch{}
	for _, node := range batch {
		batchQueue.Queue(query, argsByNode(node)...)
	}

	br := conn.SendBatch(ctx, batchQueue)
	defer br.Close()

	for i := 0; i < len(batch); i++ {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("failed to execute batch: %w", err)
		}
	}

	// Закрываем BatchReader после успешного выполнения
	if err := br.Close(); err != nil {
		return fmt.Errorf("failed to close batch reader: %w", err)
	}

	return nil
}
