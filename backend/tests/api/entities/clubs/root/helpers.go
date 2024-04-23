package root

import (
	"encoding/json"
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SampleClubFactory(userID *uuid.UUID) *map[string]interface{} {
	return &map[string]interface{}{
		"user_id":                 userID,
		"name":                    "Generate",
		"preview":                 "Generate is Northeastern's premier student-led product development studio.",
		"description":             "https://s3.amazonaws.com/<bucket>/<object>",
		"is_recruiting":           true,
		"recruitment_cycle":       "always",
		"recruitment_type":        "application",
		"weekly_time_commitment":  10,
		"one_word_to_describe_us": "muneer",
		"application_link":        "https://generatenu.com/apply",
		"logo":                    "https://s3.amazonaws.com/<bucket>/<object>",
	}
}

func AssertClubBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respClub models.Club

	err := json.NewDecoder(resp.Body).Decode(&respClub)

	eaa.Assert.NilError(err)

	var dbClubs []models.Club

	err = eaa.App.Conn.Order("created_at desc").Find(&dbClubs).Error

	eaa.Assert.NilError(err)

	dbClub := dbClubs[0]

	eaa.Assert.Equal(dbClub.ID, respClub.ID)
	eaa.Assert.Equal(dbClub.Name, respClub.Name)
	eaa.Assert.Equal(dbClub.Preview, respClub.Preview)
	eaa.Assert.Equal(dbClub.Description, respClub.Description)
	eaa.Assert.Equal(dbClub.NumMembers, respClub.NumMembers)
	eaa.Assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
	eaa.Assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
	eaa.Assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
	eaa.Assert.Equal(dbClub.WeeklyTimeCommitment, respClub.WeeklyTimeCommitment)
	eaa.Assert.Equal(dbClub.OneWordToDescribeUs, respClub.OneWordToDescribeUs)
	eaa.Assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
	eaa.Assert.Equal(dbClub.Logo, respClub.Logo)

	var dbAdmins []models.Membership

	err = eaa.App.Conn.Where("club_id = ? AND membership_type = ?", dbClub.ID, models.MembershipTypeAdmin).Find(&dbAdmins).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(1, len(dbAdmins))

	eaa.Assert.Equal(*(*body)["user_id"].(*uuid.UUID), dbAdmins[0].UserID)
	eaa.Assert.Equal((*body)["name"].(string), dbClub.Name)
	eaa.Assert.Equal((*body)["preview"].(string), dbClub.Preview)
	eaa.Assert.Equal((*body)["description"].(string), dbClub.Description)
	eaa.Assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
	eaa.Assert.Equal(models.RecruitmentCycle((*body)["recruitment_cycle"].(string)), dbClub.RecruitmentCycle)
	eaa.Assert.Equal(models.RecruitmentType((*body)["recruitment_type"].(string)), dbClub.RecruitmentType)
	eaa.Assert.Equal((*body)["weekly_time_commitment"].(int), dbClub.WeeklyTimeCommitment)
	eaa.Assert.Equal((*body)["one_word_to_describe_us"].(string), dbClub.OneWordToDescribeUs)
	eaa.Assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
	eaa.Assert.Equal((*body)["logo"].(string), dbClub.Logo)

	return dbClub.ID
}

func AssertClubWithBodyRespDBMostRecent(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respClub models.Club

	err := json.NewDecoder(resp.Body).Decode(&respClub)

	eaa.Assert.NilError(err)

	var dbClub models.Club

	err = eaa.App.Conn.Order("created_at desc").First(&dbClub).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(dbClub.ID, respClub.ID)
	eaa.Assert.Equal(dbClub.Name, respClub.Name)
	eaa.Assert.Equal(dbClub.Preview, respClub.Preview)
	eaa.Assert.Equal(dbClub.Description, respClub.Description)
	eaa.Assert.Equal(dbClub.NumMembers, respClub.NumMembers)
	eaa.Assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
	eaa.Assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
	eaa.Assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
	eaa.Assert.Equal(dbClub.WeeklyTimeCommitment, respClub.WeeklyTimeCommitment)
	eaa.Assert.Equal(dbClub.OneWordToDescribeUs, respClub.OneWordToDescribeUs)
	eaa.Assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
	eaa.Assert.Equal(dbClub.Logo, respClub.Logo)

	var dbAdmins []models.User

	err = eaa.App.Conn.Model(&dbClub).Association("Admins").Find(&dbAdmins)

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(1, len(dbAdmins))

	dbAdmin := dbAdmins[0]

	eaa.Assert.Equal((*body)["user_id"].(uuid.UUID), dbAdmin.ID)
	eaa.Assert.Equal((*body)["name"].(string), dbClub.Name)
	eaa.Assert.Equal((*body)["preview"].(string), dbClub.Preview)
	eaa.Assert.Equal((*body)["description"].(string), dbClub.Description)
	eaa.Assert.Equal((*body)["num_members"].(int), dbClub.NumMembers)
	eaa.Assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
	eaa.Assert.Equal((*body)["recruitment_cycle"].(string), dbClub.RecruitmentCycle)
	eaa.Assert.Equal((*body)["recruitment_type"].(string), dbClub.RecruitmentType)
	eaa.Assert.Equal((*body)["weekly_time_commitment"].(int), dbClub.WeeklyTimeCommitment)
	eaa.Assert.Equal((*body)["one_word_to_describe_us"].(string), dbClub.OneWordToDescribeUs)
	eaa.Assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
	eaa.Assert.Equal((*body)["logo"].(string), dbClub.Logo)

	return dbClub.ID
}

func AssertSampleClubBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, userID uuid.UUID) uuid.UUID {
	sampleClub := SampleClubFactory(&userID)
	return AssertClubBodyRespDB(eaa, resp, sampleClub)
}

func CreateSampleClub(existingAppAssert h.ExistingAppAssert) (eaa h.ExistingAppAssert, studentUUID uuid.UUID, clubUUID uuid.UUID) {
	var sampleClubUUID uuid.UUID

	newAppAssert := existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/clubs/",
			Body:               SampleClubFactory(nil),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer("user_id"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				sampleClubUUID = AssertSampleClubBodyRespDB(eaa, resp, eaa.App.TestUser.UUID)
			},
		},
	)

	return newAppAssert, newAppAssert.App.TestUser.UUID, sampleClubUUID
}
