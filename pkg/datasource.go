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
	Constant  float64 `json:"constant"`
	QueryText string  `json:"queryText"`
}

type dataSource struct {
	im              instancemgmt.InstanceManager
	logger          log.Logger
	dummyDatasource dummyDatasource
	todoDatasource  todoDatasource
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
	dataFrameDummy, err := td.dummyDatasource.Query(int(qm.Constant), qm.QueryText)
	if err != nil {
		response.Error = errors.New("Error parsing dataframes")
		return response
	}
	response.Frames = append(response.Frames, &dataFrameDummy)
	dataFrameTodos, err := td.todoDatasource.Query()
	if err != nil {
		response.Error = errors.New("Error parsing dataframes")
		return response
	}
	response.Frames = append(response.Frames, &dataFrameTodos)
	return response
}

func (td *dataSource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Health status not configured",
	}, nil
}
