package csvprovider

import (
	"fmt"
	"io/fs"
	"log/slog"
	"metrics-exporter/internal/model"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
)

type MetricsWriter interface {
	Write(metrics map[string][]model.Metric)
	FilterFiles(files []string) []string
}

type provider struct {
	writer MetricsWriter
}

func New(writer MetricsWriter) *provider {
	return &provider{
		writer: writer,
	}
}

func (p *provider) ProvideFromCsv(directory string) error {
	var csvFiles []string

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk directory: %w", err)
		}

		// check if it's a .csv file
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".csv") {
			csvFiles = append(csvFiles, path)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	if len(csvFiles) == 0 {
		slog.Info("no .csv files found")
		return nil
	}

	csvFiles = p.writer.FilterFiles(csvFiles)
	if len(csvFiles) == 0 {
		slog.Info("no new .csv files found")
		return nil
	}

	slog.Info("found new .csv files", "files len", len(csvFiles))
	finalMetrics := make(map[string][]model.Metric)

	for _, f := range csvFiles {
		metricsFile, err := os.Open(f)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer metricsFile.Close()

		var metrics []model.Metric

		if err := gocsv.UnmarshalFile(metricsFile, &metrics); err != nil {
			slog.Warn("failed to unmarshal file", "file", f, "error", err)
			continue
		}

		finalMetrics[f] = metrics
	}

	p.writer.Write(finalMetrics)

	return nil
}
