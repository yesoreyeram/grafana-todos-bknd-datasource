package plugin

import (
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type instanceSettings struct {
	httpClient *http.Client
}

func (s *instanceSettings) Dispose() {
}

func newDataSourceInstance(setting backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	return &instanceSettings{
		httpClient: &http.Client{},
	}, nil
}

// GetDatasourceServeOpts returns server options required to run as a grafana plugin
func GetDatasourceServeOpts() datasource.ServeOpts {
	loggerInstance := log.New()
	dummyds := &dummyDatasource{
		Logger: loggerInstance,
	}
	todods := &todoDatasource{
		Logger: loggerInstance,
	}
	jsonplaceholderds := &jsonPlaceholderDatasource{
		Logger: loggerInstance,
	}
	jsonds := &jsonDatasource{
		Logger: loggerInstance,
	}
	handler := &dataSource{
		InstanceManager:           datasource.NewInstanceManager(newDataSourceInstance),
		DummyDatasource:           *dummyds,
		TodoDatasource:            *todods,
		JSONplaceholderDatasource: *jsonplaceholderds,
		JSONDatasource:            *jsonds,
	}
	return datasource.ServeOpts{
		CheckHealthHandler: handler,
		QueryDataHandler:   handler,
	}
}
