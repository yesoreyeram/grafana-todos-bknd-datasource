package plugin

import (
	"fmt"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func handlePing(rw http.ResponseWriter, req *http.Request) {
	backend.Logger.Warn("Received resource call", "url", req.URL.String(), "method", req.Method)
	fmt.Fprintf(rw, "pong\n")
}

func handleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", handlePing)
}
