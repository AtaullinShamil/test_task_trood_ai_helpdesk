package nlp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type IntentRequest struct {
	Text string `json:"text"`
}

type IntentResponse struct {
	Intent string `json:"intent"`
}

func GetIntent(text string) (*IntentResponse, error) {
	url := "http://localhost:5000/intent"
	reqBody := IntentRequest{Text: text}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response IntentResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
