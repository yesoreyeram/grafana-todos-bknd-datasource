package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type jsonPlaceholderDatasource struct {
	logger log.Logger
}

func (td *jsonPlaceholderDatasource) Query(jsonEntity string, instance *instanceSettings, refID string) (frame data.Frame, err error) {
	frame.Name, frame.RefID = refID, refID
	TodoURL := fmt.Sprintf("%s/%s", "https://jsonplaceholder.typicode.com", jsonEntity)
	res, err := instance.httpClient.Get(TodoURL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	var results []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return
	}
	keys := make([]string, 0)
	for k := range results[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		switch results[0][key].(type) {
		case string:
			var items []string
			for _, result := range results {
				items = append(items, fmt.Sprintf("%v", result[key]))
			}
			frame.Fields = append(frame.Fields, data.NewField(key, nil, items))
		case int, int16, int32, int64, float32, float64:
			var items []float64
			for _, result := range results {
				items = append(items, result[key].(float64))
			}
			frame.Fields = append(frame.Fields, data.NewField(key, nil, items))
		case bool:
			var items []bool
			for _, result := range results {
				items = append(items, result[key].(bool))
			}
			frame.Fields = append(frame.Fields, data.NewField(key, nil, items))
		default:
			var items []string
			for _, result := range results {
				j, err := json.Marshal(result[key])
				if err == nil {
					items = append(items, fmt.Sprintf("%s", j))
				}
			}
			frame.Fields = append(frame.Fields, data.NewField(key, nil, items))
		}
	}
	return frame, nil
}
