package utilities

import "encoding/json"

func ToJsonString(p interface{}) (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
