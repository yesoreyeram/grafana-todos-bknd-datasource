package jsonplaceholder

import (
	"encoding/json"
	"net/http"
)

const jsonPlaceHolderURL = "https://jsonplaceholder.typicode.com"

func get(url string, obj interface{}) (err error) {
	res, err := http.Get(jsonPlaceHolderURL + "/" + url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&obj)
	if err != nil {
		return err
	}
	return nil
}
