package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func main() {
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
	ds := &dataSource{
		im:                        datasource.NewInstanceManager(newDataSourceInstance),
		logger:                    logger,
		jsonplaceholderDatasource: *jsonplaceholderds,
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
