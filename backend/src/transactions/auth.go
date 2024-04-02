package transactions

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// func SaveOTP(db *gorm.DB, userID uuid.UUID, otp string) *errors.Error {
// 	otpModel := models.Verification{
// 		UserID:    userID,
// 		Token:     otp,
// 		ExpiresAt: time.Now().Add(time.Minute * 30).UTC(),
// 		Type:      models.EmailVerificationType,
// 	}

// 	if err := db.Create(&otpModel).Error; err != nil {
// 		return &errors.FailedToSaveOTP
// 	}

// 	return nil
// }

// // @Deprecated
// func GetOTP(db *gorm.DB, userID uuid.UUID) (*models.Verification, *errors.Error) {
// 	var otp models.Verification
// 	if err := db.Where("user_id = ? AND type = ?", userID, models.EmailVerificationType).First(&otp).Error; err != nil {
// 		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, &errors.OTPNotFound
// 		}

// 		return nil, &errors.FailedToGetOTP
// 	}

// 	return &otp, nil
// }

// func DeleteOTP(db *gorm.DB, userID uuid.UUID) *errors.Error {
// 	if err := db.Where("user_id = ? AND type = ?", userID, models.EmailVerificationType).Delete(&models.Verification{}).Error; err != nil {
// 		return &errors.FailedToDeleteOTP
// 	}

// 	return nil
// }

// func SavePasswordResetToken(db *gorm.DB, userID uuid.UUID, token string) *errors.Error {
// 	passwordReset := models.Verification{
// 		UserID:    userID,
// 		Token:     token,
// 		ExpiresAt: time.Now().Add(time.Hour * 24).UTC(),
// 		Type:      models.PasswordResetType,
// 	}

// 	if err := db.Create(&passwordReset).Error; err != nil {
// 		return &errors.FailedToCreatePasswordReset
// 	}

// 	return nil
// }

// func DeletePasswordResetToken(db *gorm.DB, token string) *errors.Error {
// 	if err := db.Where("token = ? AND type = ?", token, models.PasswordResetType).Delete(&models.Verification{}).Error; err != nil {
// 		return &errors.FailedToDeletePasswordReset
// 	}

// 	return nil
// }

// // @Deprecated
// func GetPasswordResetToken(db *gorm.DB, token string) (*models.Verification, *errors.Error) {
// 	passwordReset := models.Verification{}
// 	if err := db.Where("token = ? AND type = ?", token, models.PasswordResetType).First(&passwordReset).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, &errors.PasswordResetTokenNotFound
// 		}
// 		return nil, &errors.FailedToGetPasswordResetToken
// 	}

// 	return &passwordReset, nil
// }

func GetToken(db *gorm.DB, token string, tokenType models.VerificationType) (*models.Verification, *errors.Error) {
	tokenModel := models.Verification{}
	if err := db.Where("token = ? AND type = ?", token, tokenType).First(&tokenModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &errors.TokenNotFound
		}
		return nil, &errors.FailedToGetToken
	}

	return &tokenModel, nil
}

func GetActiveTokenByUserID(db *gorm.DB, userID uuid.UUID, tokenType models.VerificationType) (*models.Verification, *errors.Error) {
	token := models.Verification{}
	if err := db.Where("user_id = ? AND expires_at > ? AND type = ?", userID, time.Now().UTC(), tokenType).First(&token).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &errors.TokenNotFound
		}
		return nil, &errors.FailedToGetToken
	}

	return &token, nil
}

func DeleteToken(db *gorm.DB, token string, tokenType models.VerificationType) *errors.Error {
	if err := db.Where("token = ? AND type = ?", token, tokenType).Delete(&models.Verification{}).Error; err != nil {
		return &errors.FailedToDeleteToken
	}

	return nil
}

func SaveToken(db *gorm.DB, userID uuid.UUID, token string, tokenType models.VerificationType, expiry time.Time) *errors.Error {
	tokenModel := models.Verification{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiry,
		Type:      tokenType,
	}

	if err := db.Create(&tokenModel).Error; err != nil {
		return &errors.FailedToSaveToken
	}

	return nil
}
