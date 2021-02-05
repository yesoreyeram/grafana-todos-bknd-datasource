package plugin

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
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

// TodosDataSource structure
type TodosDataSource struct {
	InstanceManager instancemgmt.InstanceManager
	Logger          log.Logger
}

// NewDataSource return instance of new DataSource
func NewDataSource(mux *http.ServeMux) (ds *TodosDataSource) {
	loggerInstance := log.New()
	ds = &TodosDataSource{
		Logger:          loggerInstance,
		InstanceManager: datasource.NewInstanceManager(newDataSourceInstance),
	}
	handleRoutes(mux)
	return ds
}

// CheckHealth returns healthstatus of the datasource
func (td *TodosDataSource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Health status not configured",
	}, nil
}

// QueryData return results Grafana format
func (td *TodosDataSource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	response := backend.NewQueryDataResponse()
	instance, err := getInstance(td.InstanceManager, req.PluginContext)
	if err != nil {
		return response, err
	}
	config, err := instance.getInstanceConfig(req)
	if err != nil {
		return response, err
	}
	for _, q := range req.Queries {
		res := td.query(ctx, q, instance, *config)
		response.Responses[q.RefID] = res
	}
	return response, nil
}

func (td *TodosDataSource) query(ctx context.Context, query backend.DataQuery, instance *dsInstance, config instanceConfig) backend.DataResponse {
	var qm queryModel
	response := backend.DataResponse{}
	response.Error = json.Unmarshal(query.JSON, &qm)
	if response.Error != nil {
		return response
	}
	switch qm.EntityType {
	case "dummy":
		dummyDatasource := &dummyDatasource{}
		dataFrameDummy, err := dummyDatasource.Query(int(qm.Constant), qm.QueryText, query.RefID)
		if err != nil {
			response.Error = err
			return response
		}
		response.Frames = append(response.Frames, &dataFrameDummy)
	case "todos":
		todoDatasource := &todoDatasource{}
		dataFrameTodos, err := todoDatasource.Query(qm.NumberOfTodos, qm.HideFinishedTodos, instance, query.RefID)
		if err != nil {
			response.Error = err
			return response
		}
		response.Frames = append(response.Frames, &dataFrameTodos)
	case "jsonplaceholder":
		jsonPlaceholderDatasource := jsonPlaceholderDatasource{}
		dataFrameJSONPlaceholders, err := jsonPlaceholderDatasource.Query(qm.JSONPlaceholderEntity, instance, query.RefID)
		if err != nil {
			response.Error = err
			return response
		}
		response.Frames = append(response.Frames, &dataFrameJSONPlaceholders)
	case "json":
		jsonDatasource := jsonDatasource{}
		dataFrameJSON, err := jsonDatasource.Query(qm.JSONURL, instance, query.RefID, config)
		if err != nil {
			response.Error = err
			return response
		}
		response.Frames = append(response.Frames, &dataFrameJSON)
	}
	return response
}
