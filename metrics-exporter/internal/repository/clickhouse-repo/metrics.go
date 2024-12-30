package clickhouserepo

import (
	"metrics-exporter/internal/model"
	"metrics-exporter/internal/repository"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type metricsRepository struct {
	conn         driver.Conn
	metricsTable string
}

func NewMetricsRepository(conn driver.Conn, metricsTable string) repository.MetricsRepository {
	return &metricsRepository{
		conn:         conn,
		metricsTable: metricsTable,
	}
}

func (m *metricsRepository) SaveMetrics(metrics []model.Metric) error {
	return nil
}
