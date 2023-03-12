package http

import "encoding/json"

func getBody[T any](body []byte) (*T, error) {
	var decodedData T
	err := json.Unmarshal(body, &decodedData)
	if err != nil {
		return nil, err
	}
	return &decodedData, nil
}
