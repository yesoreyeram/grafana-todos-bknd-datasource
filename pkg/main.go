package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func main() {

	backend.SetupPluginEnvironment("yesoreyeram-todosbknd-datasource")

	logger := log.New()
	dummyds := &dummyDatasource{
		logger: logger,
	}
	todods := &todoDatasource{
		logger: logger,
	}
	jsonplaceholderds := &jsonPlaceholderDatasource{
		logger: logger,
	}
	jsonds := &jsonDatasource{
		logger: logger,
	}
	ds := &dataSource{
		im:                        datasource.NewInstanceManager(newDataSourceInstance),
		logger:                    logger,
		jsonplaceholderDatasource: *jsonplaceholderds,
		jsonDatasource:            *jsonds,
		dummyDatasource:           *dummyds,
		todoDatasource:            *todods,
	}
	err := datasource.Serve(datasource.ServeOpts{
		QueryDataHandler:   ds,
		CheckHealthHandler: ds,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
