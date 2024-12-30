package main

import (
	"fmt"
	"log/slog"
	"metrics-exporter/internal/exporter"
	"metrics-exporter/internal/model"
	clickhouserepo "metrics-exporter/internal/repository/clickhouse-repo"
	clickhouseconnect "metrics-exporter/pkg/clickhouse-connect"
)

const (
	address      = "clickhouse:9000"
	user         = "default"
	password     = ""
	database     = "default"
	metricsTable = "metrics"
)

func main() {
	conn, err := clickhouseconnect.New(
		address,
		user,
		password,
		database,
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect to ClickHouse: %w", err))
	}

	slog.Info("Connected to ClickHouse")

	repo := clickhouserepo.NewMetricsRepository(
		*conn,
		metricsTable,
	)

	exporter := exporter.New(
		repo,
		make(chan []model.Metric),
	)
	
	exporter.Start()
}
