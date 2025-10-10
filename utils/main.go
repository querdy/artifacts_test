package utils

import (
	"encoding/json"
	"net/http"
)

func Stringify(s interface{}) string {
	formattedStruct, _ := json.MarshalIndent(s, "", "    ")
	return string(formattedStruct)
}

func DecodeJSONBody[T any](resp *http.Response) (*T, error) {
	var v T
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}
