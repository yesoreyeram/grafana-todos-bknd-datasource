package main

import (
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type dummyDatasource struct {
	logger log.Logger
}

func (ds *dummyDatasource) Query(constant int, queryText string) (data.Frame, error) {
	var timeslices []time.Time
	var valueslices []int64
	var stringslices []string
	currentTime := time.Now()
	for i := 0; i < int(constant); i++ {
		timeslices = append(timeslices, currentTime.Add(time.Minute))
		stringslices = append(stringslices, "hello")
		valueslices = append(valueslices, int64(i+1))
	}
	frame := data.NewFrame("Dummy Data")
	frame.Fields = append(frame.Fields, data.NewField("Time", nil, timeslices))
	frame.Fields = append(frame.Fields, data.NewField("Strings", nil, stringslices))
	frame.Fields = append(frame.Fields, data.NewField(queryText, nil, valueslices))
	return *frame, nil
}
