package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

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

type dataSource struct {
	im                        instancemgmt.InstanceManager
	logger                    log.Logger
	jsonDatasource            jsonDatasource
	jsonplaceholderDatasource jsonPlaceholderDatasource
	dummyDatasource           dummyDatasource
	todoDatasource            todoDatasource
}

func (td *dataSource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	td.logger.Warn(fmt.Sprintf("%v Query(s) Received.", len(req.Queries)))
	response := backend.NewQueryDataResponse()
	for _, q := range req.Queries {
		res := td.query(ctx, q)
		response.Responses[q.RefID] = res
	}
	return response, nil
}

func (td *dataSource) query(ctx context.Context, query backend.DataQuery) backend.DataResponse {
	var qm queryModel
	response := backend.DataResponse{}
	response.Error = json.Unmarshal(query.JSON, &qm)
	if response.Error != nil {
		return response
	}
	switch qm.EntityType {
	case "dummy":
		dataFrameDummy, err := td.dummyDatasource.Query(int(qm.Constant), qm.QueryText)
		if err != nil {
			response.Error = errors.New("Error parsing dataframes")
			return response
		}
		response.Frames = append(response.Frames, &dataFrameDummy)
	case "todos":
		dataFrameTodos, err := td.todoDatasource.Query(qm.NumberOfTodos, qm.HideFinishedTodos)
		if err != nil {
			response.Error = errors.New("Error parsing dataframes")
			return response
		}
		response.Frames = append(response.Frames, &dataFrameTodos)
	case "jsonplaceholder":
		dataFrameJSONPlaceholders, err := td.jsonplaceholderDatasource.Query(qm.JSONPlaceholderEntity)
		if err != nil {
			response.Error = errors.New("Error parsing dataframes")
			return response
		}
		response.Frames = append(response.Frames, &dataFrameJSONPlaceholders)
	case "json":
		dataFrameJSON, err := td.jsonDatasource.Query(qm.JSONURL)
		if err != nil {
			response.Error = errors.New("Error parsing dataframes")
			return response
		}
		response.Frames = append(response.Frames, &dataFrameJSON)
	}
	return response
}

func (td *dataSource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Health status not configured",
	}, nil
}
