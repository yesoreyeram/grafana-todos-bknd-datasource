package jsonplaceholder

import "fmt"

// ToDoItem structure
type ToDoItem struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// GetToDoItems returns list of ToDoItem
func GetToDoItems() (todos []ToDoItem, err error) {
	err = get("todos", &todos)
	if err != nil {
		return nil, err
	}
	return todos, err
}

// GetToDoItem returns ToDoItem by id
func GetToDoItem(id int) (todo ToDoItem, err error) {
	err = get(fmt.Sprintf("todos/%d", id), &todo)
	if err != nil {
		return todo, err
	}
	return todo, err
}
