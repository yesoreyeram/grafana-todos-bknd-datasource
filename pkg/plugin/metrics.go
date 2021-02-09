package plugin

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const appPrefix = "todo"

var promRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: fmt.Sprintf("%s_requests_total", appPrefix),
	Help: "The total number of requests",
})

func init() {
	prometheus.MustRegister(promRequestsTotal)
}

func promMetricsForInstance() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestContenct := req.Context()
		pluginContext := httpadapter.PluginConfigFromContext(requestContenct)
		if pluginContext.DataSourceInstanceSettings != nil {
			config := &instanceConfig{}
			json.Unmarshal(pluginContext.DataSourceInstanceSettings.JSONData, &config)
			newMetricsRegistry := prometheus.NewRegistry()
			// Metrics Starts here
			// Metric 1 - Random logic involve datasource instance config
			myConstantMetrics := prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: fmt.Sprintf("%s_constant", appPrefix),
				Help: "Constant coming from Grafana datasource config",
			}, []string{"instanceName", "instanceId"})
			newMetricsRegistry.MustRegister(myConstantMetrics)
			myConstantMetrics.With(prometheus.Labels{
				"instanceName": pluginContext.DataSourceInstanceSettings.Name,
				"instanceId":   fmt.Sprint(pluginContext.DataSourceInstanceSettings.ID),
			}).Set(float64(config.PromValue))
			// Metric 2 - Another Random logic involve datasource instance config
			myRandomtMetrics := prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: fmt.Sprintf("%s_random", appPrefix),
				Help: "Random",
			}, []string{"instanceName", "instanceId"})
			newMetricsRegistry.MustRegister(myRandomtMetrics)
			randomSeed := rand.NewSource(time.Now().UnixNano())
			myRandomtMetrics.With(prometheus.Labels{
				"instanceName": pluginContext.DataSourceInstanceSettings.Name,
				"instanceId":   fmt.Sprint(pluginContext.DataSourceInstanceSettings.ID),
			}).Set(float64(rand.New(randomSeed).Intn(100)) * 0.01)
			// Mertics Ends here
			han := promhttp.HandlerFor(newMetricsRegistry, promhttp.HandlerOpts{})
			han.ServeHTTP(rw, req)
		} else {
			han := promhttp.Handler()
			han.ServeHTTP(rw, req)
		}
	})
}
