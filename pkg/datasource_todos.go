package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type todoItem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type todoDatasource struct {
	logger log.Logger
}

func filterTodosByState(tds []todoItem, hideFinishedTodos bool) (ret []todoItem) {
	for _, td := range tds {
		if hideFinishedTodos == false || td.Completed == false {
			ret = append(ret, td)
		}
	}
	return ret
}

func (td *todoDatasource) Query(numberOfTodos int, hideFinishedTodos bool, instance *instanceSettings, refID string) (frame data.Frame, err error) {
	frame.Name, frame.RefID = refID, refID
	TodoURL := fmt.Sprintf("%s/%s", "https://jsonplaceholder.typicode.com", "todos")
	res, err := instance.httpClient.Get(TodoURL)
	if err != nil {
		td.logger.Warn("Error getting data from jsonplaceholder")
		return
	}
	defer res.Body.Close()
	var todos []todoItem
	err = json.NewDecoder(res.Body).Decode(&todos)
	if err != nil {
		td.logger.Warn("Error parsing data from jsonplaceholder")
		return
	}
	var todoIDs []int64
	var todoTitles []string
	var todoStatuses []string
	filteredTodos := filterTodosByState(todos, hideFinishedTodos)
	if numberOfTodos == 0 {
		numberOfTodos = 200
	}
	for i := 0; i < int(numberOfTodos) && i < len(filteredTodos); i++ {
		todoIDs = append(todoIDs, filteredTodos[i].ID)
		todoTitles = append(todoTitles, filteredTodos[i].Title)
		todoStatuses = append(todoStatuses, strconv.FormatBool(filteredTodos[i].Completed))
	}
	frame.Fields = append(frame.Fields, data.NewField("ID", nil, todoIDs))
	frame.Fields = append(frame.Fields, data.NewField("Title", nil, todoTitles))
	frame.Fields = append(frame.Fields, data.NewField("Status", nil, todoStatuses))
	return frame, nil
}
