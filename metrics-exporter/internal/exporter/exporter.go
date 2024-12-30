package exporter

import (
	"log/slog"
	"metrics-exporter/internal/model"
	"metrics-exporter/internal/repository"
)

type exporter struct {
	repo      repository.MetricsRepository
	metricsCh <-chan []model.Metric
}

func New(repo repository.MetricsRepository, metricsCh <-chan []model.Metric) *exporter {
	return &exporter{
		repo:      repo,
		metricsCh: metricsCh,
	}
}

func (e *exporter) Start() {
	for metrics := range e.metricsCh {
		err := e.repo.SaveMetrics(metrics)
		if err != nil {
			slog.Error("failed to save metrics", "error", err)
		}
	}
	slog.Info("exporter stopped")
}
