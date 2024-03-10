package rest

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServePromMetrics(gauge prometheus.Gauge) error {
	reg := prometheus.NewRegistry()
	reg.MustRegister(gauge)

	http.Handle(
		"/metrics", promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{}),
	)

	return http.ListenAndServe(":8080", nil)
}
