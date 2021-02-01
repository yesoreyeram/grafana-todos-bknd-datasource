package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func GetDataframeFromJSONReponse(body io.ReadCloser, refID string) (frame data.Frame, err error) {
	frame.Name, frame.RefID = refID, refID
	var results []map[string]interface{}
	err = json.NewDecoder(body).Decode(&results)
	if err != nil {
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
