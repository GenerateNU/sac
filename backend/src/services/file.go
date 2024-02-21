package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileServiceInterface interface {
	CreateFile(fileBody models.FileBody, file models.File, data *multipart.FileHeader, reader io.Reader) (*models.File, *errors.Error)
	GetFile(id string) (*models.File, *errors.Error)
	GetFileInfo(id string, days string) (*models.File, *errors.Error)
	DeleteFile(id string, s3Only bool) error
}

type FileService struct {
	DB       *gorm.DB
	Settings config.AWSSettings
	Validate *validator.Validate
}

func NewFileService(db *gorm.DB, settings config.AWSSettings, validate *validator.Validate) *FileService {
	return &FileService{DB: db, Settings: settings, Validate: validate}
}

func createAWSSession(settings config.AWSSettings) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials(settings.ID, settings.SECRET, ""),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}

// Get File
func (f *FileService) GetFile(id string) (*models.File, *errors.Error) {
	var file models.File

	idAsUUID, errUUID := utilities.ValidateID(id)

	if errUUID != nil {
		return nil, &errors.FailedToValidateID
	}

	if err := f.DB.First(&file, idAsUUID).Error; err != nil {
		return &models.File{}, &errors.FailedToGetFile
	}

	sess, err := createAWSSession(f.Settings)
	if err != nil {
		return nil, &errors.FailedToCreateAWSSession
	}

	downloader := s3manager.NewDownloader(sess)

	downloadedFile, err := os.Create(file.FileName)

	_, err = downloader.Download(downloadedFile,
		&s3.GetObjectInput{
			Bucket: aws.String("generate-sac-storage"),
			Key:    aws.String(file.FileName),
		})

	fileData, fileDataErr := io.ReadAll(downloadedFile)
	if fileDataErr != nil {
		return nil, &errors.FailedToDownloadFile
	}
	file.FileData = fileData

	return &file, nil
}

func (f *FileService) GetFileInfo(id string, days string) (*models.File, *errors.Error) {
	var file models.File

	idAsUUID, errUUID := utilities.ValidateID(id)

	if errUUID != nil {
		return nil, &errors.FailedToValidateID
	}

	if err := f.DB.First(&file, idAsUUID).Error; err != nil {
		return nil, &errors.FailedToGetFile
	}

	sess, err := createAWSSession(f.Settings)
	if err != nil {
		return nil, &errors.FailedToCreateAWSSession
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("generate-sac-storage"),
		Key:    aws.String(file.ObjectKey),
	})

	daysInt, err := strconv.Atoi(days)
	if err != nil {
		return nil, &errors.FailedToParseDaysToInt
	}
	expiration := time.Duration(24*time.Hour) * time.Duration(daysInt)

	url, err := req.Presign(expiration)
	if err != nil {
		return nil, &errors.FailedToGetSignedURL
	}

	file.S3Url = url
	return &file, nil
}

// Create File
func (f *FileService) CreateFile(fileBody models.FileBody, file models.File, data *multipart.FileHeader, reader io.Reader) (*models.File, *errors.Error) {
	var testFile models.File
	var searchFiles []models.File
	file.FileName = data.Filename

	// if file type and file id are populated, if both are then search to see if id exists in corresponding table. Throw error
	// if entry does not exist
	ownerType := fileBody.OwnerType
	ownerID := fileBody.OwnerID
	print(ownerID.String())

	possibleTypes := []string{"users", "clubs", "events", "point_of_contacts"}

	if ownerID != uuid.Nil {
		if !isInArray(possibleTypes, ownerType) {
			return nil, &errors.InvalidAssociationType
		}
		result := f.DB.Table(ownerType).Where("id = ?", ownerID).RowsAffected
		if result == 0 {
			return nil, &errors.FailedToFindAssociationID
		}
	}

	// check if filename is already taken, and add (filenumber) to name if it is
	objectKey := file.FileName
	dotIndex := strings.LastIndex(objectKey, ".")
	file_substring := objectKey[:dotIndex]
	file_extension := objectKey[dotIndex:]
	searchKey := file_substring + "%" + file_extension

	file.ObjectKey = objectKey

	if err := f.DB.Where("object_key = ?", objectKey).Find(&testFile).Error; err == nil {
		f.DB.Where("object_key LIKE ?", searchKey).Find(&searchFiles)
		i := len(searchFiles)

		file_num := fmt.Sprintf(" (%v)", i)
		file.ObjectKey = file_substring + file_num + file_extension
		file.FileName = file_substring[strings.Index(file_substring, "-")+1:] + file_num + file_extension
	}

	// Check if the file size is greater than 5 MB
	if data.Size > 5000000 {
		return nil, &errors.InvalidFileSize
	}

	file.FileSize = data.Size

	// Upload the file to the S3 bucket
	sess, err := createAWSSession(f.Settings)
	if err != nil {
		return nil, &errors.FailedToCreateAWSSession
	}

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(f.Settings.BUCKET_NAME),
		Key:    aws.String(file.ObjectKey),
		Body:   reader,
	})
	if err != nil {
		return nil, &errors.FailedToUploadToS3
	}

	// Create the file in the database
	if err := f.DB.Create(&file).Error; err != nil {
		f.DeleteFile(fmt.Sprint(file.ID), true) // delete file from s3 if it cant be made in database
		return nil, &errors.FailedToCreateFileInDB
	}
	return &file, nil
}

// check if a string is in the array
func isInArray(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// Delete File
func (f *FileService) DeleteFile(id string, s3Only bool) error {
	var file models.File

	idAsUUID, errUUID := utilities.ValidateID(id)
	if errUUID != nil {
		return &errors.FailedToValidateID
	}

	if err := f.DB.First(&file, idAsUUID).Error; err != nil {
		return err
	}

	// create session and service client, then delete file
	sess, err := createAWSSession(f.Settings)
	if err != nil {
		return err
	}

	svc := s3.New(sess)
	objectKey := file.FileName

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(f.Settings.BUCKET_NAME),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}

	// Required to delete the file from the database permanently
	if !s3Only {
		if err := f.DB.Unscoped().Delete(&file).Error; err != nil {
			return err
		}
	}
	return nil
}

func ValidateData(c *fiber.Ctx, model interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(model); err != nil {
		return err
	}

	return nil
}
