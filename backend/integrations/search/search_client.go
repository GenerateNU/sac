package search

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GenerateNU/sac/backend/entities/models"
	"gorm.io/gorm"

	"github.com/go-resty/resty/v2"
)

type SearchClientInterface interface {
	Seed(db *gorm.DB) error
}

type OpenSearchClient struct {
	Url string
}

func NewOpenSearchClient(Url string) SearchClientInterface {
	return &OpenSearchClient{
		Url: Url,
	}
}

func jsonString(p map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(p)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (c *OpenSearchClient) fullUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.Url, endpoint)
}

func (c *OpenSearchClient) Seed(db *gorm.DB) error {
	var clubs []models.Club

	if err := db.Find(&clubs).Error; err != nil {
		return err
	}

	var requestBodyBuilder strings.Builder

	for _, club := range clubs {
		indexData, err := jsonString(map[string]interface{}{
			"create": map[string]interface{}{
				"_index": "clubs",
				"_id":    club.ID.String(),
			},
		})

		if err != nil {
			return err
		}

		bodyData, err := jsonString(map[string]interface{}{
			"Name":                 club.Name,
			"Preview":              club.Preview,
			"Description":          club.Description,
			"NumMembers":           club.NumMembers,
			"IsRecruiting":         club.IsRecruiting,
			"RecruitmentCycle":     club.RecruitmentCycle,
			"RecruitmentType":      club.RecruitmentType,
			"WeeklyTimeCommitment": club.WeeklyTimeCommitment,
			"OneWordToDescribeUs":  club.OneWordToDescribeUs,
		})

		if err != nil {
			return err
		}

		requestBodyBuilder.WriteString(fmt.Sprintf("%s\n%s\n", indexData, bodyData))
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBodyBuilder.String()).
		Post(c.fullUrl("_bulk"))

	if err != nil {
		return err
	}

	print(resp.String())
	return nil
}
