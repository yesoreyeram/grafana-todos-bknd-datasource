package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type jsonDatasource struct {
	logger log.Logger
}

func (td *jsonDatasource) Query(jsonURL string, instance *instanceSettings, refID string, config dataSourceConfig) (frame data.Frame, err error) {
	frame.Name, frame.RefID = refID, refID
	TodoURL := fmt.Sprintf("%s", jsonURL)
	if TodoURL == "" {
		TodoURL = config.DefaultJSONURL
	}
	res, err := instance.httpClient.Get(TodoURL)
	if err != nil {
		td.logger.Error("Error retreiving data from URL")
		return
	}
	defer res.Body.Close()
	var results []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		td.logger.Error("Error parsing data received")
		return frame, err
	}
	keys := make([]string, 0)
	for k := range results[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		switch results[0][key].(type) {
		case string:
			items := make([]string, len(results))
			for i, result := range results {
				if result[key] != nil {
					items[i] = fmt.Sprintf("%v", result[key])
				}
			}
			frame.Fields = append(frame.Fields, data.NewField(key, nil, items))
		case int, int16, int32, int64, float32, float64:
			items := make([]float64, len(results))
			for i, result := range results {
				if result[key] != nil {
					items[i] = result[key].(float64)
				}
			}
			frame.Fields = append(frame.Fields, data.NewField(key, nil, items))
		case bool:
			items := make([]bool, len(results))
			for i, result := range results {
				if result[key] != nil {
					items[i] = result[key].(bool)
				}
			}
			frame.Fields = append(frame.Fields, data.NewField(key, nil, items))
		default:
			items := make([]string, len(results))
			for i, result := range results {
				if result[key] != nil {
					j, err := json.Marshal(result[key])
					if err == nil {
						items[i] = fmt.Sprintf("%s", j)
					}
				}
			}
			frame.Fields = append(frame.Fields, data.NewField(key, nil, items))
		}
	}
	return frame, nil
}
