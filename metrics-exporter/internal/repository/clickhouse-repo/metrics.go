package clickhouserepo

import (
	"context"
	"fmt"
	"metrics-exporter/internal/model"
	"metrics-exporter/internal/repository"

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
	query := fmt.Sprintf("INSERT INTO %s", m.metricsTable)
	batch, err := m.conn.PrepareBatch(ctx, query)
	if err != nil {
		return err
	}

	for _, metric := range metrics {
		err := batch.Append(
			metric.Date,
			metric.Time,
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
	return batch.Send()
}

func (m *metricsRepository) GetExportedFiles(ctx context.Context) ([]string, error) {
	return nil, nil
}
