package plugin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
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
	ctx := req.Context()
	ctxt := httpadapter.PluginConfigFromContext(ctx)
	if ctxt.DataSourceInstanceSettings != nil {
		backend.Logger.Warn("Received instance resource call", "url", req.URL.String(), "method", req.Method, "pluginid", ctxt.PluginID)
		backend.Logger.Warn(string(ctxt.DataSourceInstanceSettings.JSONData))
		config := &instanceConfig{}
		json.Unmarshal(ctxt.DataSourceInstanceSettings.JSONData, &config)
		fmt.Fprintf(rw, "pong "+fmt.Sprint(config.PromValue)+"\n")
	} else {
		backend.Logger.Warn("Received plugin resource call", "url", req.URL.String(), "method", req.Method, "pluginid", ctxt.PluginID)
		fmt.Fprintf(rw, "pong\n")
	}
}

func (td *TodosDataSource) handleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", handlePing)

	mux.Handle("/metrics", promhttp.Handler())

	promMetricsRegistry.MustRegister(promRequestsTotal)
	promMetricsRegistry.MustRegister(promQueriesTotal)
	mux.Handle("/todos/metrics", promhttp.HandlerFor(promMetricsRegistry, promhttp.HandlerOpts{}))
}
