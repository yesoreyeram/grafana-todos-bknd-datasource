package plugin

import (
	"context"
	"encoding/json"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type queryModel struct {
	EntityType            string  `json:"entityType"`
	Constant              float64 `json:"constant"`
	QueryText             string  `json:"queryText"`
	NumberOfTodos         int     `json:"numberOfTodos"`
	HideFinishedTodos     bool    `json:"hideFinishedTodos"`
	JSONPlaceholderEntity string  `json:"jsonPlaceholderEntity"`
	JSONURL               string  `json:"jsonURL"`
}

// dataSource structure
type dataSource struct {
	InstanceManager           instancemgmt.InstanceManager
	Logger                    log.Logger
	JSONDatasource            jsonDatasource
	JSONplaceholderDatasource jsonPlaceholderDatasource
	DummyDatasource           dummyDatasource
	TodoDatasource            todoDatasource
}

// DataSourceConfig return structure of DataSource Configuration
type DataSourceConfig struct {
	Path           string `json:"path"`
	DefaultJSONURL string `json:"defaultJSONURL"`
}

// CheckHealth returns healthstatus of the datasource
func (td *dataSource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Health status not configured",
	}, nil
}

// QueryData return results Grafana format
func (td *dataSource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	response := backend.NewQueryDataResponse()
	instance, err := td.getInstance(req.PluginContext)
	if err != nil {
		return response, err
	}
	config, err := td.getInstanceConfig(req)
	if err != nil {
		return response, err
	}
	for _, q := range req.Queries {
		res := td.query(ctx, q, instance, *config)
		response.Responses[q.RefID] = res
	}
	return response, nil
}

func (td *dataSource) query(ctx context.Context, query backend.DataQuery, instance *instanceSettings, config DataSourceConfig) backend.DataResponse {
	var qm queryModel
	response := backend.DataResponse{}
	response.Error = json.Unmarshal(query.JSON, &qm)
	if response.Error != nil {
		return response
	}
	switch qm.EntityType {
	case "dummy":
		dataFrameDummy, err := td.DummyDatasource.Query(int(qm.Constant), qm.QueryText, query.RefID)
		if err != nil {
			response.Error = err
			return response
		}
		response.Frames = append(response.Frames, &dataFrameDummy)
	case "todos":
		dataFrameTodos, err := td.TodoDatasource.Query(qm.NumberOfTodos, qm.HideFinishedTodos, instance, query.RefID)
		if err != nil {
			response.Error = err
			return response
		}
		response.Frames = append(response.Frames, &dataFrameTodos)
	case "jsonplaceholder":
		dataFrameJSONPlaceholders, err := td.JSONplaceholderDatasource.Query(qm.JSONPlaceholderEntity, instance, query.RefID)
		if err != nil {
			response.Error = err
			return response
		}
		response.Frames = append(response.Frames, &dataFrameJSONPlaceholders)
	case "json":
		dataFrameJSON, err := td.JSONDatasource.Query(qm.JSONURL, instance, query.RefID, config)
		if err != nil {
			response.Error = err
			return response
		}
		response.Frames = append(response.Frames, &dataFrameJSON)
	}
	return response
}

func (td *dataSource) getInstance(ctx backend.PluginContext) (*instanceSettings, error) {
	instance, err := td.InstanceManager.Get(ctx)
	if err != nil {
		return nil, err
	}
	return instance.(*instanceSettings), nil
}

func (td *dataSource) getInstanceConfig(req *backend.QueryDataRequest) (config *DataSourceConfig, err error) {
	err = json.Unmarshal(req.PluginContext.DataSourceInstanceSettings.JSONData, &config)
	if err != nil {
		return nil, err
	}
	return config, err
}
