package plugin

import (
	"encoding/json"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
)

type instanceConfig struct {
	PromValue      int    `json:"promValue"`
	Path           string `json:"path"`
	DefaultJSONURL string `json:"defaultJSONURL"`
}

type dsInstance struct {
	httpClient *http.Client
}

func getInstanceConfig(req *backend.QueryDataRequest) (config *instanceConfig, err error) {
	err = json.Unmarshal(req.PluginContext.DataSourceInstanceSettings.JSONData, &config)
	if err != nil {
		return nil, err
	}
	return config, err
}

func (ins *dsInstance) Dispose() {
}

func getInstance(ins instancemgmt.InstanceManager, ctx backend.PluginContext) (*dsInstance, error) {
	instance, err := ins.Get(ctx)
	if err != nil {
		return nil, err
	}
	return instance.(*dsInstance), nil
}

func newDataSourceInstance(setting backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	return &dsInstance{
		httpClient: &http.Client{},
	}, nil
}
