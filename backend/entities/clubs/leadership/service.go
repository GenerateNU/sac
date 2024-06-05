package leadership

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/entities/files/base"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubLeaderServiceInterface interface {
	GetClubLeadership(clubID string) ([]models.Leader, error)
	GetClubLeader(clubID, leaderID string) (*models.Leader, error)
	CreateClubLeader(clubID string, leaderBody CreateLeaderBody, fileHeader *multipart.FileHeader) (*models.Leader, error)
	UpdateClubLeaderPhoto(clubID, leaderID string, fileHeader *multipart.FileHeader) (*models.Leader, error)
	UpdateClubLeader(clubID, leaderID string, leaderBody UpdateLeaderBody) (*models.Leader, error)
	DeleteClubLeader(clubID, leaderID string) error
}

type ClubLeaderService struct {
	types.ServiceParams
}

func NewClubLeaderService(serviceParams types.ServiceParams) ClubLeaderServiceInterface {
	return &ClubLeaderService{serviceParams}
}

func (cl *ClubLeaderService) GetClubLeadership(clubID string) ([]models.Leader, error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubLeadership(cl.DB, *clubIdAsUUID)
}

func (cl *ClubLeaderService) GetClubLeader(clubID, leaderID string) (*models.Leader, error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	leaderIDAsUUID, err := utilities.ValidateID(leaderID)
	if err != nil {
		return nil, err
	}

	return GetClubLeader(cl.DB, *clubIdAsUUID, *leaderIDAsUUID)
}

// MOVE THIS TO TRANSACTIOPNS
func (cl *ClubLeaderService) CreateClubLeader(clubID string, leaderBody CreateLeaderBody, fileHeader *multipart.FileHeader) (*models.Leader, error) {
	if err := utilities.Validate(cl.Validate, leaderBody); err != nil {
		return nil, err
	}

	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	leader, err := GetClubLeaderByClubIDAndEmail(cl.DB, *clubIdAsUUID, leaderBody.Email)
	if err == nil {
		return leader, nil
	}

	return CreateClubLeader(cl.DB,
		models.Leader{
			Name:     leaderBody.Name,
			Email:    leaderBody.Email,
			Position: leaderBody.Position,
			ClubID:   *clubIdAsUUID,
		},
		func() (*models.FileInfo, error) {
			return cl.Integrations.File.UploadFile("leadership", fileHeader, []file.FileType{file.IMAGE})
		},
		cl.Integrations.File.DeleteFile,
	)
}

func (cl *ClubLeaderService) UpdateClubLeaderPhoto(clubID, leaderID string, fileHeader *multipart.FileHeader) (*models.Leader, error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	leaderIDAsUUID, err := utilities.ValidateID(leaderID)
	if err != nil {
		return nil, err
	}

	leader, err := GetClubLeader(cl.DB, *clubIdAsUUID, *leaderIDAsUUID)
	if err != nil {
		return nil, err
	}

	if err := cl.Integrations.File.DeleteFile(leader.PhotoFile.FileURL); err != nil {
		return nil, err
	}

	fileInfo, err := cl.Integrations.File.UploadFile("point_of_contacts", fileHeader, []file.FileType{file.IMAGE})
	if err != nil {
		return nil, err
	}

	file, err := base.UpdateFile(cl.DB, leader.PhotoFile.ID, *fileInfo)
	if err != nil {
		return nil, err
	}

	leader.PhotoFile = *file

	return leader, nil
}

func (cl *ClubLeaderService) UpdateClubLeader(clubID, leaderID string, leaderBody UpdateLeaderBody) (*models.Leader, error) {
	if err := utilities.Validate(cl.Validate, leaderBody); err != nil {
		return nil, err
	}

	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	leaderIDAsUUID, err := utilities.ValidateID(leaderID)
	if err != nil {
		return nil, err
	}

	leader, err := utilities.MapJsonTags(leaderBody, &models.Leader{})
	if err != nil {
		return nil, err
	}

	return UpdateClubLeader(cl.DB, *clubIdAsUUID, *leaderIDAsUUID, *leader)
}

func (cl *ClubLeaderService) DeleteClubLeader(clubID, leaderID string) error {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	leaderIDAsUUID, err := utilities.ValidateID(leaderID)
	if err != nil {
		return err
	}

	return DeleteClubLeader(cl.DB, *clubIdAsUUID, *leaderIDAsUUID, cl.Integrations.File.DeleteFile)
}
