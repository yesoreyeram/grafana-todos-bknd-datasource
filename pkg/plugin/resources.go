package plugin

import (
	"fmt"
	"net/http"
)

const (
	resourcesURLPing    = "/ping"
	resourcesURLMetrics = "/metrics"
)

func handlePing(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "pong\n")
}

func (td *TodosDataSource) handleRoutes(mux *http.ServeMux) {
	mux.HandleFunc(resourcesURLPing, handlePing)
	mux.Handle(resourcesURLMetrics, promMetricsForInstance())
}
