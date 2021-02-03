package plugin

import (
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/yesoreyeram/grafana-todos-bknd-datasource/pkg/parser"
)

type jsonDatasource struct{}

func (td *jsonDatasource) Query(jsonURL string, instance *dsInstance, refID string, config instanceConfig) (frame data.Frame, err error) {
	JSONURL := fmt.Sprintf("%s", jsonURL)
	if JSONURL == "" {
		JSONURL = config.DefaultJSONURL
	}
	res, err := instance.httpClient.Get(JSONURL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	frame, err = parser.GetDataframeFromJSONReponse(res.Body, refID)
	return frame, err
}
