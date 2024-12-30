package csvprovider

import "metrics-exporter/internal/model"

type provider struct {
	metricsCh chan<- []model.Metric
}

func New(metricsCh chan<- []model.Metric) *provider {
	return &provider{
		metricsCh: metricsCh,
	}
}




