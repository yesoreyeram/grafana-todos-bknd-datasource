package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/yesoreyeram/grafana-todos-bknd-datasource/pkg/plugin"
)

func main() {

	backend.SetupPluginEnvironment("yesoreyeram-todosbknd-datasource")

	ds := plugin.NewDataSource()

	err := datasource.Serve(datasource.ServeOpts{
		CheckHealthHandler:  ds,
		QueryDataHandler:    ds,
		CallResourceHandler: ds.ResourceHandler,
	})

	if err != nil {
		ds.Logger.Error(err.Error())
		os.Exit(1)
	}

}
