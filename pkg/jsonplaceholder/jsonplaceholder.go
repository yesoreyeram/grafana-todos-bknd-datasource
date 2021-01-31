package jsonplaceholder

import (
	"encoding/json"
	"net/http"
)

const jsonPlaceHolderURL = "https://jsonplaceholder.typicode.com"

// ToDoItem structure
type ToDoItem struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// GetToDoItems returns list of ToDoItem
func GetToDoItems() (todos []ToDoItem, err error) {
	res, err := http.Get(jsonPlaceHolderURL + "/todos")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&todos)
	if err != nil {
		return nil, err
	}
	return
}
