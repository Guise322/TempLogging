package main

import (
	"fmt"
	"temp-logging/internal/metrics"
	"temp-logging/internal/rest"
	"temp-logging/internal/service"
	"temp-logging/internal/temp"
	"time"

	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	slog.Info("Start gathering the CPU temperature data")

	tempGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "CPU_temperature",
		Help: "The value of the current CPU temperature.",
	})

	tempDataGetter := temp.TemperatureDataGetter{}
	tempDataSetter := metrics.NewGaugeDataWriter(tempGauge)
	tempDataService := service.NewTempDataService()
	dataWritingPeriod := 5 * time.Second

	go func() {
		err := tempDataService.ServeTempData(tempDataGetter, tempDataSetter, dataWritingPeriod)
		errMsg := fmt.Sprintf("Get an error at serving temperature data: %v", err)
		slog.Error(errMsg)
	}()

	err := rest.ServePromMetrics(tempGauge)
	if err != nil {
		errMsg := fmt.Sprintf("Get an error at serving Prometheus: %v", err)
		slog.Error(errMsg)
	}
}
