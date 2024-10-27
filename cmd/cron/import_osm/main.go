package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/qedus/osmpbf"

	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/postgres"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/pbf_scanner"
)

const (
	batchSize = 10000
	filePath  = "osmfiles/RU.osm.pbf"
	pgConnStr = "postgres://osmuser:osmpassword@localhost:5432/osm"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(slog.LevelDebug, "dev", os.Stdout)

	// Инициализация компонентов
	file, err := os.Open(filePath)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to open file")
		os.Exit(1)
	}
	defer file.Close()

	pg, err := postgres.New(pgConnStr)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
		os.Exit(1)
	}
	defer pg.Close()

	sc, err := pbf_scanner.New(file, runtime.GOMAXPROCS(-1))
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to initialize pbf scanner")
		os.Exit(1)
	}

	// Накопление узлов для батча
	nodesBatch := make([]osmpbf.Node, 0, batchSize)
	for {
		node, err := sc.Next()
		if err != nil {
			break
		}

		nodesBatch = append(nodesBatch, node)
		if len(nodesBatch) >= batchSize {
			if err := insertBatch(ctx, pg.Pool, nodesBatch); err != nil {
				log.WithContext(ctx).WithError(err).Warning("failed to insert nodes")
				continue
			}
			nodesBatch = nodesBatch[:0]
		}
	}

	if len(nodesBatch) > 0 {
		if err := insertBatch(ctx, pg.Pool, nodesBatch); err != nil {
			fmt.Printf("Failed to insert final batch: %v\n", err)
		}
	}

	log.WithContext(ctx).Info("OSM nodes have been successfully scanned and inserted into the database.")
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
func insertBatch(ctx context.Context, pg postgres.PgxPool, batch []osmpbf.Node) error {
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

	br := pg.SendBatch(ctx, batchQueue)
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
