package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/yesoreyeram/grafana-todos-bknd-datasource/pkg/plugin"
)

func main() {

	const pluginID = "yesoreyeram-todosbknd-datasource"

	backend.SetupPluginEnvironment(pluginID)

	err := datasource.Serve(plugin.GetDatasourceServeOpts())

	if err != nil {
		backend.Logger.Error(err.Error())
		os.Exit(1)
	}

}
