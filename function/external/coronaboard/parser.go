package coronaboard

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

func ParseBoard(ctx context.Context) (*CoronaBoardData, error) {
	client := http.Client{}
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://coronaboard.kr", nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	jsonRegex := regexp.MustCompile(`jsonData = (.*?);</script`)
	jsonBody := jsonRegex.FindSubmatch(body)[1]
	if err != nil {
		return nil, err
	}

	var data CoronaBoardData
	json.Unmarshal(jsonBody, &data)
	return &data, nil
}
