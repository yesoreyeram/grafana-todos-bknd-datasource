package plugin

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

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
	promRandom = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_randmom", appPrefix),
		Help: "Random",
	})
)

func init() {
	promMetricsRegistry.MustRegister(promRequestsTotal)
	promMetricsRegistry.MustRegister(promQueriesTotal)
	promMetricsRegistry.MustRegister(promRandom)
	s1 := rand.NewSource(time.Now().UnixNano())
	go func() {
		for {
			promRandom.Set(float64(rand.New(s1).Intn(100)) * 0.01)
			time.Sleep(time.Second * 2)
		}
	}()
}

func handlePing(rw http.ResponseWriter, req *http.Request) {
	requestContenct := req.Context()
	pluginContext := httpadapter.PluginConfigFromContext(requestContenct)
	if pluginContext.DataSourceInstanceSettings != nil {
		backend.Logger.Warn("Received instance resource call", "url", req.URL.String(), "method", req.Method, "pluginid", pluginContext.PluginID, "instanceID", pluginContext.DataSourceInstanceSettings.ID, "instanceName", pluginContext.DataSourceInstanceSettings.Name)
		config := &instanceConfig{}
		json.Unmarshal(pluginContext.DataSourceInstanceSettings.JSONData, &config)
		fmt.Fprintf(rw, "pong "+fmt.Sprint(config.PromValue)+"\n")
	} else {
		backend.Logger.Warn("Received plugin resource call", "url", req.URL.String(), "method", req.Method, "pluginid", pluginContext.PluginID)
		fmt.Fprintf(rw, "pong\n")
	}
}

func (td *TodosDataSource) handleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", handlePing)
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/todos/metrics", promhttp.HandlerFor(promMetricsRegistry, promhttp.HandlerOpts{}))
}
