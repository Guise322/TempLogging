package metrics

import "github.com/prometheus/client_golang/prometheus"

type DataWriter interface {
	WriteData(data float64)
}

type GaugeDataWriter struct {
	gauge prometheus.Gauge
}

func NewGaugeDataWriter(gauge prometheus.Gauge) GaugeDataWriter {
	return GaugeDataWriter{gauge: gauge}
}

func (g GaugeDataWriter) WriteData(data float64) {
	g.gauge.Set(data)
}
