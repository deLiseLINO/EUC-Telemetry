package repository

import (
	"context"
	"metrics-exporter/internal/model"
)

type MetricsRepository interface {
	SaveMetricsAndFiles(ctx context.Context, metrics []model.Metric, files []string) error
	GetExportedFiles(ctx context.Context) ([]string, error)
}
