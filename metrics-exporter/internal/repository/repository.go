package repository

import "metrics-exporter/internal/model"

type MetricsRepository interface {
	SaveMetrics(metrics []model.Metric) error
}
