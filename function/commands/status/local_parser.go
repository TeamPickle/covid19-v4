package status

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseBoard() (*CoronaBoardData, error) {
	resp, err := http.Get("https://coronaboard.kr")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data CoronaBoardData
	json.Unmarshal(body, &data)
	return &data, nil
}
