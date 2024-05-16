package utilities

import (
	"strings"

	"github.com/mitchellh/mapstructure"
)

// MapRequestToModel maps request data to a target model using mapstructure
func MapRequestToModel[T any, U any](responseData T, targetModel *U) (*U, error) {
	config := &mapstructure.DecoderConfig{
		Result:  targetModel,
		TagName: "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(responseData)
	if err != nil {
		return nil, err
	}

	return targetModel, nil
}

func NormalizeEmail(email string) string {
	return strings.ToLower(email)
}
