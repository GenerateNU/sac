package services

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/email"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	GetRole(id string) (*models.UserRole, *errors.Error)
	Me(id string) (*models.User, *errors.Error)
	Login(userBody models.LoginUserResponseBody) (*models.User, *errors.Error)
	UpdatePassword(id string, passwordBody models.UpdatePasswordRequestBody) *errors.Error
	SendCode(userID string) *errors.Error
	VerifyEmail(emailBody models.VerifyEmailRequestBody) *errors.Error
	ForgotPassword(userBody models.PasswordResetRequestBody) *errors.Error
	VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) *errors.Error
}

type AuthService struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Email    *email.EmailService
}

func NewAuthService(db *gorm.DB, validate *validator.Validate, email *email.EmailService) *AuthService {
	return &AuthService{
		DB:       db,
		Validate: validate,
		Email:    email,
	}
}

func (a *AuthService) Me(id string) (*models.User, *errors.Error) {
	idAsUint, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	user, err := transactions.GetUser(a.DB, *idAsUint)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *AuthService) Login(userBody models.LoginUserResponseBody) (*models.User, *errors.Error) {
	if err := a.Validate.Struct(userBody); err != nil {
		return nil, &errors.FailedToValidateUser
	}

	user, err := transactions.GetUserByEmail(a.DB, userBody.Email)
	if err != nil {
		return nil, err
	}

	correct, passwordErr := auth.CompareHash(userBody.Password, user.PasswordHash)
	if passwordErr != nil || !correct {
		return nil, &errors.FailedToValidateUser
	}

	return user, nil
}

func (a *AuthService) GetRole(id string) (*models.UserRole, *errors.Error) {
	idAsUint, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	user, err := transactions.GetUser(a.DB, *idAsUint)
	if err != nil {
		return nil, err
	}

	role := user.Role

	return &role, nil
}

func (a *AuthService) UpdatePassword(id string, passwordBody models.UpdatePasswordRequestBody) *errors.Error {
	idAsUint, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return idErr
	}

	if err := a.Validate.Struct(passwordBody); err != nil {
		return &errors.FailedToValidateUser
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	passwordHash, err := transactions.GetUserPasswordHash(tx, *idAsUint)
	if err != nil {
		tx.Rollback()
		return err
	}

	correct, passwordErr := auth.CompareHash(passwordBody.OldPassword, passwordHash)
	if passwordErr != nil || !correct {
		tx.Rollback()
		return &errors.FailedToValidateUser
	}

	hash, hashErr := auth.ComputeHash(passwordBody.NewPassword)
	if hashErr != nil {
		tx.Rollback()
		return &errors.FailedToValidateUser
	}

	updateErr := transactions.UpdatePassword(tx, *idAsUint, *hash)
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &errors.FailedToUpdatePassword
	}

	return nil
}

func (a *AuthService) ForgotPassword(userBody models.PasswordResetRequestBody) *errors.Error {
	if err := a.Validate.Struct(userBody); err != nil {
		return &errors.FailedToValidateUser
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := transactions.GetUserByEmail(tx, userBody.Email)
	if err != nil {
		return nil // Do not return error if user does not exist
	}

	// Check for existing or generate new password reset token
	activeToken, tokenErr := transactions.GetActivePasswordResetTokenByUserID(a.DB, user.ID)
	if tokenErr != nil {
		if tokenErr != &errors.PasswordResetTokenNotFound {
			return tokenErr
		}
	}

	if activeToken != nil {
		sendErr := a.Email.SendPasswordResetEmail(user.FirstName, user.Email, activeToken.Token)
		if sendErr != nil {
			tx.Rollback()
			return &errors.FailedToSendEmail
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return &errors.FailedToGenerateToken
		}

		return nil
	}

	// Generate token if none exists
	token, generateErr := auth.GenerateURLSafeToken(64)
	if generateErr != nil {
		tx.Rollback()
		return &errors.FailedToGenerateToken
	}

	// Save token to database
	saveErr := transactions.SavePasswordResetToken(tx, user.ID, *token)
	if saveErr != nil {
		tx.Rollback()
		return saveErr
	}

	// Send email
	sendErr := a.Email.SendPasswordResetEmail(user.FirstName, user.Email, *token)
	if sendErr != nil {
		tx.Rollback()
		return sendErr
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &errors.FailedToGenerateToken
	}

	return nil
}

func (a *AuthService) VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) *errors.Error {
	if err := a.Validate.Struct(passwordBody); err != nil {
		return &errors.FailedToValidateUser
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	token, tokenErr := transactions.GetPasswordResetToken(tx, passwordBody.Token)
	if tokenErr != nil {
		tx.Rollback()
		return tokenErr
	}

	if token.ExpiresAt.Before(time.Now().UTC()) {
		tx.Rollback()
		return &errors.TokenExpired
	}

	hash, hashErr := auth.ComputeHash(passwordBody.NewPassword)
	if hashErr != nil {
		tx.Rollback()
		return &errors.FailedToValidateUser
	}

	updateErr := transactions.UpdatePassword(tx, token.UserID, *hash)
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	deleteErr := transactions.DeletePasswordResetToken(tx, passwordBody.Token)
	if deleteErr != nil {
		tx.Rollback()
		return deleteErr
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &errors.FailedToUpdatePassword
	}

	return nil
}

func (a *AuthService) SendCode(userID string) *errors.Error {
	idAsUint, idErr := utilities.ValidateID(userID)
	if idErr != nil {
		return idErr
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := transactions.GetUser(tx, *idAsUint)
	if err != nil {
		tx.Rollback()
		return err
	}

	otp, otpErr := auth.GenerateOTP(6)
	if otpErr != nil {
		tx.Rollback()
		return &errors.FailedToGenerateOTP
	}

	saveErr := transactions.SaveOTP(tx, user.ID, *otp)
	if saveErr != nil {
		tx.Rollback()
		return saveErr
	}

	sendErr := a.Email.SendEmailVerification(user.Email, *otp)
	if sendErr != nil {
		tx.Rollback()
		return &errors.FailedToSendEmail
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &errors.FailedToSendCode
	}

	return nil
}

func (a *AuthService) VerifyEmail(emailBody models.VerifyEmailRequestBody) *errors.Error {
	if err := a.Validate.Struct(emailBody); err != nil {
		return &errors.FailedToValidateUser
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := transactions.GetUserByEmail(tx, emailBody.Email)
	if err != nil {
		tx.Rollback()
		return err
	}

	if user.IsVerified {
		tx.Rollback()
		return &errors.EmailAlreadyVerified
	}

	otp, otpErr := transactions.GetOTP(tx, user.ID)
	if otpErr != nil {
		tx.Rollback()
		return otpErr
	}

	if otp.Token != emailBody.Token {
		tx.Rollback()
		return &errors.InvalidOTP
	}

	if otp.ExpiresAt.Before(time.Now().UTC()) {
		tx.Rollback()
		return &errors.OTPExpired
	}

	updateErr := transactions.UpdateEmailVerification(tx, user.ID)
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	deleteErr := transactions.DeleteOTP(tx, user.ID)
	if deleteErr != nil {
		tx.Rollback()
		return deleteErr
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateEmailVerification
	}

	return nil
}
