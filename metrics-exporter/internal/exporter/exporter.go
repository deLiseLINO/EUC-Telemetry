package exporter

import (
	"context"
	"log/slog"
	"metrics-exporter/internal/model"
	"metrics-exporter/internal/repository"
	"sync"
	"time"
)

// key of the map is the .csv file name
type exporter struct {
	repo        repository.MetricsRepository
	metricsCh   chan map[string][]model.Metric
	wg          sync.WaitGroup
	exportedMap map[string]struct{}
	mu          sync.Mutex
	closed      bool
}

func New(repo repository.MetricsRepository) *exporter {
	return &exporter{
		repo:        repo,
		metricsCh:   make(chan map[string][]model.Metric),
		wg:          sync.WaitGroup{},
		exportedMap: make(map[string]struct{}),
	}
}

func (e *exporter) Run(ctx context.Context) error {
	// get exported files first
	files, err := e.repo.GetExportedFiles(ctx)
	if err != nil {
		return err
	}
	for _, f := range files {
		e.exportedMap[f] = struct{}{}
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		for {
			select {
			case <-ctx.Done():
				e.stop()
				return
			case metricsMap := <-e.metricsCh:
				var (
					metrics []model.Metric
					files   []string
				)
				e.mu.Lock()
				for k, v := range metricsMap {
					if _, found := e.exportedMap[k]; !found {
						e.exportedMap[k] = struct{}{}
						metrics = append(metrics, v...)
						files = append(files, k)
					}
				}
				e.mu.Unlock()
				err := e.repo.SaveMetricsAndFiles(ctxTimeout, metrics, files)
				if err != nil {
					slog.Error("failed to save metrics", slog.Any("error", err))
				}
			}
		}
	}()
	return nil
}

func (e *exporter) Write(metricsMap map[string][]model.Metric) {
	if e.closed {
		return
	}
	if len(metricsMap) > 0 {
		e.metricsCh <- metricsMap
	}
}

func (e *exporter) FilterFiles(files []string) []string {
	e.mu.Lock()
	defer e.mu.Unlock()
	var filteredFiles []string
	for _, f := range files {
		if _, found := e.exportedMap[f]; !found {
			filteredFiles = append(filteredFiles, f)
		}
	}
	return filteredFiles
}

func (e *exporter) stop() {
	e.closed = true
	close(e.metricsCh)
	slog.Info("waiting for exporter to finish the tasks")
	e.wg.Wait()
	slog.Info("exporter stopped")
}
