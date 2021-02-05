package plugin

import (
	"fmt"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const appPrefix = "todo"

var (
	promMetricsRegistry = prometheus.NewRegistry()
	promRequestsTotal   = prometheus.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_requests_total", appPrefix),
		Help: "The total number of requests",
	})
	promQueriesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_queries_total", appPrefix),
			Help: "The total number of queries",
		},
		[]string{"entityType"},
	)
)

func handlePing(rw http.ResponseWriter, req *http.Request) {
	backend.Logger.Warn("Received resource call", "url", req.URL.String(), "method", req.Method)
	fmt.Fprintf(rw, "pong\n")
}

func handleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", handlePing)

	mux.Handle("/internal/metrics", promhttp.Handler())

	promMetricsRegistry.MustRegister(promRequestsTotal)
	promMetricsRegistry.MustRegister(promQueriesTotal)
	mux.Handle("/metrics", promhttp.HandlerFor(promMetricsRegistry, promhttp.HandlerOpts{}))
}
