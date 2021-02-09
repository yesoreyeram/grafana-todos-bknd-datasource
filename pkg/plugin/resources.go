package plugin

import (
	"fmt"
	"net/http"
)

func handlePing(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "pong\n")
}

func (td *TodosDataSource) handleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", handlePing)
	mux.Handle("/metrics", promMetricsForInstance())
}
