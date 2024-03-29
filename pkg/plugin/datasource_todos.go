package plugin

import (
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/yesoreyeram/grafana-todos-bknd-datasource/pkg/jsonplaceholder"
)

type todoItem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type todoDatasource struct{}

func (td *todoDatasource) Query(numberOfTodos int, hideFinishedTodos bool, instance *dsInstance, refID string) (frame data.Frame, err error) {
	frame.Name, frame.RefID = refID, refID
	var todos []jsonplaceholder.ToDoItem
	todos, err = jsonplaceholder.GetToDoItems()
	if err != nil {
		return
	}
	filteredTodos := filterTodosByState(todos, hideFinishedTodos)
	if numberOfTodos == 0 {
		numberOfTodos = 200
	}
	var todoUserIDs []int64
	var todoIDs []int64
	var todoTitles []string
	var todoStatuses []bool
	for i := 0; i < int(numberOfTodos) && i < len(filteredTodos); i++ {
		todoIDs = append(todoIDs, int64(filteredTodos[i].ID))
		todoUserIDs = append(todoUserIDs, int64(filteredTodos[i].UserID))
		todoTitles = append(todoTitles, filteredTodos[i].Title)
		todoStatuses = append(todoStatuses, filteredTodos[i].Completed)
	}
	frame.Fields = append(frame.Fields, data.NewField("ID", nil, todoIDs))
	frame.Fields = append(frame.Fields, data.NewField("User ID", nil, todoUserIDs))
	frame.Fields = append(frame.Fields, data.NewField("Title", nil, todoTitles))
	frame.Fields = append(frame.Fields, data.NewField("Status", nil, todoStatuses))
	return frame, nil
}

func filterTodosByState(tds []jsonplaceholder.ToDoItem, hideFinishedTodos bool) (ret []jsonplaceholder.ToDoItem) {
	for _, td := range tds {
		if hideFinishedTodos == false || td.Completed == false {
			ret = append(ret, td)
		}
	}
	return ret
}
