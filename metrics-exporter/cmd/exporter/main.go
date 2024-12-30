package main

import (
	"context"
	"fmt"
	"log/slog"
	csvprovider "metrics-exporter/internal/csv-provider"
	"metrics-exporter/internal/exporter"
	clickhouserepo "metrics-exporter/internal/repository/clickhouse-repo"
	clickhouseconnect "metrics-exporter/pkg/clickhouse-connect"
	"sync"
)

const (
	address         = "localhost:19000"
	user            = "default"
	password        = ""
	database        = "default"
	metricsTable    = "metrics"
	savedFilesTable = "files"

	csvFileDir = "wheel-logs"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := clickhouseconnect.New(
		address,
		user,
		password,
		database,
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect to clickhouse: %w", err))
	}
	slog.Info("connected to clickhouse")

	repo := clickhouserepo.NewMetricsRepository(
		*conn,
		metricsTable,
		savedFilesTable,
	)

	wg := sync.WaitGroup{}
	exporter := exporter.New(repo, &wg)
	if err := exporter.Run(ctx); err != nil {
		panic(fmt.Errorf("failed to run exporter: %w", err))
	}

	provider := csvprovider.New(exporter)
	if err = provider.ProvideFromCsv(csvFileDir); err != nil {
		panic(fmt.Errorf("failed to provide metrics: %w", err))
	}

	exporter.Stop()
}
