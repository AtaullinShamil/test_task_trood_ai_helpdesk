package nlp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type AnalyzeRequest struct {
	Text string `json:"text"`
}

type AnalyzeResponse struct {
	Tokens   []string    `json:"tokens"`
	Entities [][2]string `json:"entities"`
	PosTags  [][2]string `json:"pos_tags"`
}

func AnalyzeText(text string) (*AnalyzeResponse, error) {
	url := "http://localhost:5000/analyze"
	reqBody := AnalyzeRequest{Text: text}
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

	var response AnalyzeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
