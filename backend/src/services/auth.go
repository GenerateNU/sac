package services

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
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
	ForgotPassword(userBody models.PasswordResetRequestBody) *errors.Error
	VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) *errors.Error
}

type AuthService struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Email    *auth.EmailService
}

func NewAuthService(db *gorm.DB, validate *validator.Validate, email *auth.EmailService) *AuthService {
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

	passwordHash, err := transactions.GetUserPasswordHash(a.DB, *idAsUint)
	if err != nil {
		return err
	}

	correct, passwordErr := auth.CompareHash(passwordBody.OldPassword, passwordHash)
	if passwordErr != nil || !correct {
		return &errors.FailedToValidateUser
	}

	hash, hashErr := auth.ComputeHash(passwordBody.NewPassword)
	if hashErr != nil {
		return &errors.FailedToValidateUser
	}

	updateErr := transactions.UpdatePassword(a.DB, *idAsUint, *hash)
	if updateErr != nil {
		return updateErr
	}

	return nil
}

func (a *AuthService) ForgotPassword(userBody models.PasswordResetRequestBody) *errors.Error {
	if err := a.Validate.Struct(userBody); err != nil {
		return &errors.FailedToValidateUser
	}

	user, err := transactions.GetUserByEmail(a.DB, userBody.Email)
	if err != nil {
		return nil // Do not return error if user does not exist
	}

	// check if user has a password reset token, if not, generate one, if yes, use the existing one
	activeToken, tokenErr := transactions.GetActivePasswordResetTokenByUserID(a.DB, user.ID)
	if tokenErr != nil {
		if tokenErr != &errors.PasswordResetTokenNotFound {
			return tokenErr
		}
	}

	if activeToken != nil {
		return nil
	}

	token, generateErr := auth.GeneratePasswordResetToken()
	if generateErr != nil {
		return &errors.FailedToGenerateToken
	}

	// save token to db
	saveErr := transactions.SavePasswordResetToken(a.DB, user.ID, token)
	if saveErr != nil {
		return saveErr
	}

	// PLEASE NOTE: don't overuse this email service in testing (we only have 1000 free emails per month)
	sendErr := a.Email.SendPasswordResetEmail(user.FirstName, user.Email, token)
	if sendErr != nil {
		deleteErr := transactions.DeletePasswordResetToken(a.DB, token)
		if deleteErr != nil {
			return deleteErr
		}
		return &errors.FailedToSendEmail
	}

	return nil
}

func (a *AuthService) VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) *errors.Error {
	if err := a.Validate.Struct(passwordBody); err != nil {
		return &errors.FailedToValidateUser
	}

	token, tokenErr := transactions.GetPasswordResetToken(a.DB, passwordBody.Token)
	if tokenErr != nil {
		return tokenErr
	}

	if token.ExpiresAt.Before(time.Now().UTC()) {
		return &errors.TokenExpired
	}

	hash, hashErr := auth.ComputeHash(passwordBody.NewPassword)
	if hashErr != nil {
		return &errors.FailedToValidateUser
	}

	updateErr := transactions.UpdatePassword(a.DB, token.UserID, *hash)
	if updateErr != nil {
		return updateErr
	}

	deleteErr := transactions.DeletePasswordResetToken(a.DB, passwordBody.Token)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
