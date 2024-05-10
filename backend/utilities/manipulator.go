package utilities

import (
	"fmt"

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
		return nil, fmt.Errorf("failed to create decoder: %w", err)
	}

	err = decoder.Decode(responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	return targetModel, nil
}
