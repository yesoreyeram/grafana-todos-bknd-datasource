package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (td *todoDatasource) Query(numberOfTodos int, hideFinishedTodos bool) (data.Frame, error) {
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
	filteredTodos := filterTodosByState(todos, hideFinishedTodos)
	for i := 0; i < int(numberOfTodos) && i < len(filteredTodos); i++ {
		todoIDs = append(todoIDs, filteredTodos[i].ID)
		todoTitles = append(todoTitles, filteredTodos[i].Title)
		todoStatuses = append(todoStatuses, strconv.FormatBool(filteredTodos[i].Completed))
	}
	frame := data.NewFrame("Todos")
	frame.Fields = append(frame.Fields, data.NewField("ID", nil, todoIDs))
	frame.Fields = append(frame.Fields, data.NewField("Title", nil, todoTitles))
	frame.Fields = append(frame.Fields, data.NewField("Status", nil, todoStatuses))
	return *frame, nil
}
