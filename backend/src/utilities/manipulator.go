package utilities

import (
	"github.com/mitchellh/mapstructure"
)

// MapResponseToModel maps response data to a target model using mapstructure
func MapResponseToModel[T any, U any](responseData T, targetModel *U) (*U, error) {
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