package utilities

import go_json "github.com/goccy/go-json"

func ToJsonString(p interface{}) (string, error) {
	jsonData, err := go_json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
