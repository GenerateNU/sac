package search

import (
	"fmt"
	"strings"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

func fullUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", constants.SEARCH_URL, endpoint)
}

func Seed(db *gorm.DB) error {
	var clubs []models.Club

	if err := db.Find(&clubs).Error; err != nil {
		return err
	}

	var requestBodyBuilder strings.Builder

	for _, club := range clubs {
		indexData, err := utilities.ToJsonString(map[string]interface{}{
			"create": map[string]interface{}{
				"_index": "clubs",
				"_id":    club.ID.String(),
			},
		})

		if err != nil {
			return err
		}

		bodyData, err := utilities.ToJsonString(map[string]interface{}{
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
		Post(fullUrl("_bulk"))

	if err != nil {
		return err
	}

	print(resp.String())
	return nil
}
