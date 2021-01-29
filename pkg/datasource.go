package main

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
)

type queryModel struct {
	Constant  float64 `json:"constant"`
	QueryText string  `json:"queryText"`
}

type todosBkdnDatasource struct {
	im instancemgmt.InstanceManager
}

func (td *todosBkdnDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	response := backend.NewQueryDataResponse()
	for _, q := range req.Queries {
		res := td.query(ctx, q)
		response.Responses[q.RefID] = res
	}
	return response, nil
}

func (td *todosBkdnDatasource) query(ctx context.Context, query backend.DataQuery) backend.DataResponse {
	var qm queryModel
	response := backend.DataResponse{}
	response.Error = json.Unmarshal(query.JSON, &qm)
	if response.Error != nil {
		return response
	}
	dataFrame, err := getDummyData(int(qm.Constant), qm.QueryText)
	if err != nil {
		response.Error = errors.New("Error parsing dataframes")
		return response
	}
	response.Frames = append(response.Frames, &dataFrame)
	return response
}

func (td *todosBkdnDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Health status not configured",
	}, nil
}
