package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type queryModel struct {
	Format string `json:"format"`
}

type todosBkdnDatasource struct {
	im instancemgmt.InstanceManager
}

func (td *todosBkdnDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	log.DefaultLogger.Info("QueryData", "request", req)
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

	if qm.Format == "" {
		log.DefaultLogger.Warn("format is empty. defaulting to time series")
	}

	frame := data.NewFrame("response")

	frame.Fields = append(frame.Fields,
		data.NewField("time", nil, []time.Time{query.TimeRange.From, query.TimeRange.To}),
	)

	frame.Fields = append(frame.Fields,
		data.NewField("values", nil, []int64{10, 20}),
	)

	response.Frames = append(response.Frames, frame)

	return response
}

func (td *todosBkdnDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Health status not configured",
	}, nil
}
