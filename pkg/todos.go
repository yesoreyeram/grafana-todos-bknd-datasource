package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type todoItem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type dummyServer struct {
	logger log.Logger
}

func (td *dummyServer) getDummyData(constant int, queryText string) (data.Frame, error) {
	var timeslices []time.Time
	var valueslices []int64
	var stringslices []string
	for i := 0; i < int(constant); i++ {
		timeslices = append(timeslices, time.Now())
		stringslices = append(stringslices, "hello")
		valueslices = append(valueslices, int64(i+1))
	}
	frame := data.NewFrame("Dummy Data")
	frame.Fields = append(frame.Fields, data.NewField("Time", nil, timeslices))
	frame.Fields = append(frame.Fields, data.NewField("Strings", nil, stringslices))
	frame.Fields = append(frame.Fields, data.NewField(queryText, nil, valueslices))
	return *frame, nil
}

func (td *dummyServer) getTodos() (data.Frame, error) {
	TodoURL := fmt.Sprintf("%s/%s", "https://jsonplaceholder.typicode.com", "todos")
	res, err := http.Get(TodoURL)
	if err != nil {
		td.logger.Warn("Error getting data from jsonplaceholder")
	}
	defer res.Body.Close()
	var todos []todoItem
	err = json.NewDecoder(res.Body).Decode(&todos)
	if err != nil {
		td.logger.Warn("Error parsing data from jsonplaceholder")
	}
	var todoIDs []int64
	var todoTitles []string
	var todoStatuses []string
	for _, todoitem := range todos {
		todoIDs = append(todoIDs, todoitem.ID)
		todoTitles = append(todoTitles, todoitem.Title)
		todoStatuses = append(todoStatuses, strconv.FormatBool(todoitem.Completed))
	}
	frame := data.NewFrame("Todos")
	frame.Fields = append(frame.Fields, data.NewField("ID", nil, todoIDs))
	frame.Fields = append(frame.Fields, data.NewField("Title", nil, todoTitles))
	frame.Fields = append(frame.Fields, data.NewField("Status", nil, todoStatuses))
	return *frame, nil
}
