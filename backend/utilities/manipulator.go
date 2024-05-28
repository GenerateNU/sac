package utilities

import (
	"fmt"
	"strings"

	"github.com/go-viper/mapstructure/v2"
)

// MapJsonTags maps request data to a target model using mapstructure
// NOTE: this function does not correctly map time.Time fields from from to time.Time fields in intoType
// see the following GitHub issues:
// https://github.com/mitchellh/mapstructure/issues/334
// https://github.com/go-viper/mapstructure/issues/20
func MapJsonTags[From any, Into any](from From, intoType *Into) (*Into, error) {
	config := &mapstructure.DecoderConfig{
		Result:  intoType,
		TagName: "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create decoder: %w", err)
	}

	err = decoder.Decode(from)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	return intoType, nil
}

func NormalizeEmail(email string) string {
	return strings.ToLower(email)
}

func SplitBearerToken(accessToken string) (*string, error) {
	token := strings.Split(accessToken, "Bearer ")
	if len(token) != 2 {
		return nil, fmt.Errorf("invalid token")
	}

	return &token[1], nil
}
