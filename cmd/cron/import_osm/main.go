package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/qedus/osmpbf"
	"github.com/samber/lo"

	"github.com/yasonofriychuk/real-estate-insight/internal/config"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/postgres"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/pbf_scanner"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/utils"
)

const (
	batchSize       = 15000
	goroutinesCount = 2
	filePath        = "osmfiles/RU.osm.pbf"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(slog.LevelDebug, "dev", os.Stdout)
	cfg := config.MustNewConfigWithEnv()

	// Инициализация компонентов
	file, err := os.Open(filePath)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to open file")
		os.Exit(1)
	}
	defer file.Close()

	pg, err := postgres.New(cfg.PgUrl())
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
		os.Exit(1)
	}
	defer pg.Close()

	sc, err := pbf_scanner.New(file, goroutinesCount)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to initialize pbf scanner")
		os.Exit(1)
	}

	// Накопление узлов для батча
	var totalImport uint64
	nodesBatch := make([]osmpbf.Node, 0, batchSize)
	for {
		node, err := sc.Next()
		if err != nil {
			break
		}
		if len(utils.ObjectTypeByTags(node.Tags)) == 0 {
			continue
		}
		if strings.TrimSpace(node.Tags["name"]) == "" {
			continue
		}

		nodesBatch = append(nodesBatch, node)
		if len(nodesBatch) >= batchSize {
			if err := insertBatch(ctx, pg.Pool, nodesBatch); err != nil {
				log.WithContext(ctx).WithError(err).Warning("failed to insert nodes")
				continue
			}
			totalImport += uint64(len(nodesBatch))
			nodesBatch = nodesBatch[:0]

			log.WithContext(ctx).WithFields(map[string]any{
				"counts": totalImport,
			}).Info("imported nodes")
		}
	}

	if len(nodesBatch) > 0 {
		if err := insertBatch(ctx, pg.Pool, nodesBatch); err != nil {
			fmt.Printf("Failed to insert final batch: %v\n", err)
		}
	}

	log.WithContext(ctx).Info("OSM nodes have been successfully scanned and inserted into the database.")
}

func argsByNode(node osmpbf.Node) []any {
	return []any{
		node.ID,
		node.Tags["name"],
		strings.Join(lo.Map(utils.ObjectTypeByTags(node.Tags), func(t osm.ObjType, _ int) string {
			return string(t) + " => yes"
		}), ","),
		node.Lon,
		node.Lat,
	}
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
