package main

import (
	"os"

	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
	"github.com/yesoreyeram/grafana-todos-bknd-datasource/pkg/plugin"
)

func main() {

	backend.SetupPluginEnvironment("yesoreyeram-todosbknd-datasource")

	mux := http.NewServeMux()
	httpResourceHandler := httpadapter.New(mux)
	ds := plugin.NewDataSource(mux)

	err := datasource.Serve(datasource.ServeOpts{
		CheckHealthHandler:  ds,
		QueryDataHandler:    ds,
		CallResourceHandler: httpResourceHandler,
	})

	if err != nil {
		ds.Logger.Error(err.Error())
		os.Exit(1)
	}

}
