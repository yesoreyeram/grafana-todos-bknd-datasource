package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/data"
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

	var timeslices []time.Time
	var valueslices []int64
	for i := 0; i < int(qm.Constant); i++ {
		timeslices = append(timeslices, query.TimeRange.From)
		valueslices = append(valueslices, int64(i))
	}
	frame := data.NewFrame("response")
	frame.Fields = append(frame.Fields, data.NewField("time", nil, timeslices))
	frame.Fields = append(frame.Fields, data.NewField(qm.QueryText, nil, valueslices))
	response.Frames = append(response.Frames, frame)
	return response
}

func (td *todosBkdnDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Health status not configured",
	}, nil
}
