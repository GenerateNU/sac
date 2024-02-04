package errors

const (
	FailedToValidateFileId   = "failed to validate file id"
	InvalidFileSize          = "file size is greater than 5 MB"
	FailedToCreateAWSSession = "failed to create AWS session"
	FailedToUploadToS3       = "failed to upload to S3 Bucket"
	FailedToCreateFileInDB   = "failed to create file in database"
	FailedToGetFile          = "failed to get file"
	FailedToProcessRequest   = "failed to process the request"
	FailedToValidatedData    = "failed to validate data"
	FailedToOpenFile         = "failed to open file"
	InvalidImageFormat       = "invalid image format"
	FailedToDeleteFile       = "failed to delete file from database"
)
