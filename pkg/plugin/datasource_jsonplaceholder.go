package plugin

import (
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/yesoreyeram/grafana-todos-bknd-datasource/pkg/parser"
)

type jsonPlaceholderDatasource struct{}

func (td *jsonPlaceholderDatasource) Query(jsonEntity string, instance *instanceSettings, refID string) (frame data.Frame, err error) {
	TodoURL := fmt.Sprintf("%s/%s", "https://jsonplaceholder.typicode.com", jsonEntity)
	res, err := instance.httpClient.Get(TodoURL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	frame, err = parser.GetDataframeFromJSONReponse(res.Body, refID)
	return frame, nil
}
