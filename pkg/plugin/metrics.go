package plugin

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/prometheus/client_golang/prometheus"
)

const appPrefix = "todo"

var (
	promRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_plugin_requests_total", appPrefix),
		Help: "The total number of  requests for the plugin",
	})
	promMetricsRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_plugin_metrics_requests_total", appPrefix),
		Help: "The total number of metric requests for the plugin",
	})
)

func init() {
	prometheus.MustRegister(promRequestsTotal)
	prometheus.MustRegister(promMetricsRequestsTotal)
}

type metricRegistryInstance struct {
	Registry *prometheus.Registry
	Metrics  struct {
		TotalMetricsRequest prometheus.Counter
		Constant            *prometheus.GaugeVec
		Random              *prometheus.GaugeVec
	}
}

func newMetricRegistryInstance() *metricRegistryInstance {
	instance := &metricRegistryInstance{}
	instance.Registry = prometheus.NewRegistry()
	instance.Metrics.Constant = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_constant", appPrefix),
		Help: "Constant coming from Grafana datasource config",
	}, []string{"instanceName", "instanceId"})
	instance.Registry.MustRegister(instance.Metrics.Constant)
	instance.Metrics.Random = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_random", appPrefix),
		Help: "Random",
	}, []string{"instanceName", "instanceId"})
	instance.Registry.MustRegister(instance.Metrics.Random)
	instance.Metrics.TotalMetricsRequest = prometheus.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_metrics_requests_total", appPrefix),
		Help: "The total number of requests",
	})
	instance.Registry.MustRegister(instance.Metrics.TotalMetricsRequest)
	return instance
}

var registryInstances map[int64]metricRegistryInstance = make(map[int64]metricRegistryInstance)

func getInstanceRegistry(id int64) metricRegistryInstance {
	if _, ok := registryInstances[id]; ok {
		return registryInstances[id]
	}
	registryInstances[id] = *newMetricRegistryInstance()
	return registryInstances[id]
}

func (registry *metricRegistryInstance) collectMetrics(pluginContext backend.PluginContext, config *instanceConfig) {
	// Global metric
	promMetricsRequestsTotal.Inc()
	// Metric 0 - Total Metrics Requests collection
	registry.Metrics.TotalMetricsRequest.Inc()
	// Metric 1 - Constant value involve datasource instance config
	value := float64(config.PromValue)
	registry.Metrics.Constant.With(prometheus.Labels{
		"instanceName": pluginContext.DataSourceInstanceSettings.Name,
		"instanceId":   fmt.Sprint(pluginContext.DataSourceInstanceSettings.ID),
	}).Set(value)
	// Metric 2 - Another Random logic involve datasource instance config
	randomSeed := rand.NewSource(time.Now().UnixNano())
	value1 := float64(rand.New(randomSeed).Intn(100)) * 0.01
	registry.Metrics.Random.With(prometheus.Labels{
		"instanceName": pluginContext.DataSourceInstanceSettings.Name,
		"instanceId":   fmt.Sprint(pluginContext.DataSourceInstanceSettings.ID),
	}).Set(value1)
}
