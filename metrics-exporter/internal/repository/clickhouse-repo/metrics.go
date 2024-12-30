package clickhouserepo

import (
	"context"
	"fmt"
	"log/slog"
	"metrics-exporter/internal/model"
	"metrics-exporter/internal/repository"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type metricsRepository struct {
	conn            driver.Conn
	metricsTable    string
	savedFilesTable string
}

func NewMetricsRepository(conn driver.Conn, metricsTable string, savedFilesTable string) repository.MetricsRepository {
	return &metricsRepository{
		conn:            conn,
		metricsTable:    metricsTable,
		savedFilesTable: savedFilesTable,
	}
}

func (m *metricsRepository) SaveMetricsAndFiles(ctx context.Context, metrics []model.Metric, files []string) error {
	if err := m.saveMetrics(ctx, metrics); err != nil {
		return err
	}

	if err := m.saveFiles(ctx, files); err != nil {
		return err
	}

	return nil
}

func (m *metricsRepository) saveFiles(ctx context.Context, files []string) error {
	t := time.Now()
	query := fmt.Sprintf("INSERT INTO %s", m.savedFilesTable)
	batch, err := m.conn.PrepareBatch(ctx, query)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := batch.Append(
			file,
			time.Now(),
		)
		if err != nil {
			return err
		}
	}

	slog.Info("saved files",
		"count", len(files),
		"table", m.savedFilesTable,
		"took", time.Since(t),
	)

	return batch.Send()
}

func (m *metricsRepository) saveMetrics(ctx context.Context, metrics []model.Metric) error {
	t := time.Now()
	slog.Info("saving metrics", "count", len(metrics))
	query := fmt.Sprintf("INSERT INTO %s", m.metricsTable)
	batch, err := m.conn.PrepareBatch(ctx, query)
	if err != nil {
		return err
	}

	for _, metric := range metrics {
		t, err := time.Parse("2006-01-02 15:04:05.000", fmt.Sprintf("%v %v", metric.Date, metric.Time))
		if err != nil {
			slog.Error("failed to parse time", "time", metric.Time, "error", err)
			return err
		}

		err = batch.Append(
			t,
			t,
			metric.Latitude,
			metric.Longitude,
			metric.GPSSpeed,
			metric.GPSAlt,
			metric.GPSHeading,
			metric.GPSDistance,
			metric.Speed,
			metric.Voltage,
			metric.PhaseCurrent,
			metric.Current,
			metric.Power,
			metric.Torque,
			metric.PWM,
			metric.BatteryLevel,
			metric.Distance,
			metric.TotalDistance,
			metric.SystemTemp,
			metric.Temp2,
			metric.Tilt,
			metric.Roll,
			metric.Mode,
			metric.Alert,
		)
		if err != nil {
			return err
		}
	}

	slog.Info("saved metrics",
		"count", len(metrics),
		"table", m.metricsTable,
		"took", time.Since(t),
	)
	return batch.Send()
}

func (m *metricsRepository) GetExportedFiles(ctx context.Context) ([]string, error) {
	files := []string{}
	query := fmt.Sprintf("SELECT name FROM %s", m.savedFilesTable)
	rows, err := m.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var file string
		if err := rows.Scan(&file); err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}
