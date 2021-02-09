package plugin

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlePing(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "pong\n")
}

func (td *TodosDataSource) handleRoutes(mux *mux.Router) {
	mux.HandleFunc("/ping", handlePing)
	mux.Handle("/metrics", promMetricsForInstance())
}
