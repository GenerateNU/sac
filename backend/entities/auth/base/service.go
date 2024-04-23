package base

import (
	"time"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	GetRole(id string) (*models.UserRole, *errors.Error)
	Login(userBody models.LoginUserResponseBody) (*models.User, *auth.Token, *errors.Error)
	Refresh(refreshToken string) (*auth.Token, *errors.Error)
	Logout(c *fiber.Ctx) *errors.Error

	SendCode(email string) *errors.Error
	VerifyEmail(emailBody models.VerifyEmailRequestBody) *errors.Error

	ForgotPassword(email string) *errors.Error
	VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) *errors.Error
}

type AuthService struct {
	types.ServiceParams
}

func NewAuthService(serviceParams types.ServiceParams) AuthServiceInterface {
	return &AuthService{serviceParams}
}

func (a *AuthService) Login(loginBody models.LoginUserResponseBody) (*models.User, *auth.Token, *errors.Error) {
	if err := a.Validate.Struct(loginBody); err != nil {
		return nil, nil, &errors.FailedToValidateUser
	}

	if pwordValErr := auth.ValidatePassword(loginBody.Password); pwordValErr != nil {
		return nil, nil, pwordValErr
	}

	user, getUserByEmailErr := users.GetUserByEmail(a.DB, loginBody.Email)
	if getUserByEmailErr != nil {
		return nil, nil, getUserByEmailErr
	}

	correct, passwordErr := auth.CompareHash(loginBody.Password, user.PasswordHash)
	if passwordErr != nil || !correct {
		return nil, nil, &errors.FailedToValidateUser
	}

	tokens, err := a.JWT.GenerateTokenPair(auth.Claims{
		StandardClaims: &jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer:   user.ID.String(),
		},
		CustomClaims: &jwt.MapClaims{
			"role": user.Role,
		},
	}, auth.Claims{
		StandardClaims: &jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer:   user.ID.String(),
		},
	})
	if err != nil {
		return nil, nil, &errors.FailedToGenerateToken
	}

	return user, tokens, nil
}

func (a *AuthService) Refresh(refreshToken string) (*auth.Token, *errors.Error) {
	claims, err := a.JWT.ExtractClaims(refreshToken, auth.RefreshToken)
	if err != nil {
		return nil, &errors.FailedToExtractClaims
	}

	role, roleErr := a.GetRole(claims["iss"].(string))
	if roleErr != nil {
		return nil, roleErr
	}

	tokens, err := a.JWT.GenerateTokenPair(auth.Claims{
		StandardClaims: &jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer:   claims["iss"].(string),
		},
		CustomClaims: &jwt.MapClaims{
			"role": role,
		},
	}, auth.Claims{
		StandardClaims: &jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer:   claims["iss"].(string),
		},
	})
	if err != nil {
		return nil, &errors.FailedToGenerateToken
	}

	return tokens, nil
}

func (a *AuthService) GetRole(id string) (*models.UserRole, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	user, err := users.GetUser(a.DB, *idAsUUID)
	if err != nil {
		return nil, err
	}

	return &user.Role, nil
}

func (a *AuthService) ForgotPassword(email string) *errors.Error {
	user, err := users.GetUserByEmail(a.DB, email)
	if err != nil {
		return nil
	}

	activeToken, tokenErr := GetActiveTokenByUserID(a.DB, user.ID, models.PasswordResetType)
	if tokenErr != nil {
		if tokenErr != &errors.TokenNotFound {
			return tokenErr
		}
	}

	if activeToken != nil {
		sendErr := a.Deps.Email.SendPasswordResetEmail(user.FirstName, user.Email, activeToken.Token)
		if sendErr != nil {
			return &errors.FailedToSendEmail
		}

		return nil
	}

	token, generateErr := auth.GenerateURLSafeToken(64)
	if generateErr != nil {
		return &errors.FailedToGenerateToken
	}

	saveErr := SaveToken(a.DB, user.ID, *token, models.PasswordResetType, time.Now().Add(time.Hour*24).UTC())
	if saveErr != nil {
		return saveErr
	}

	sendErr := a.Deps.Email.SendPasswordResetEmail(user.FirstName, user.Email, *token)
	if sendErr != nil {
		return sendErr
	}

	return nil
}

func (a *AuthService) VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) *errors.Error {
	if err := a.Validate.Struct(passwordBody); err != nil {
		return &errors.FailedToValidateUser
	}

	token, tokenErr := GetToken(a.DB, passwordBody.Token, models.PasswordResetType)
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

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	updateErr := users.UpdatePassword(tx, token.UserID, *hash)
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	deleteErr := DeleteToken(tx, passwordBody.Token, models.PasswordResetType)
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

func (a *AuthService) SendCode(email string) *errors.Error {
	user, err := users.GetUserByEmail(a.DB, email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return &errors.EmailAlreadyVerified
	}

	activeOTP, tokenErr := GetActiveTokenByUserID(a.DB, user.ID, models.EmailVerificationType)
	if tokenErr != nil {
		if tokenErr != &errors.TokenNotFound {
			return tokenErr
		}
	}

	if activeOTP != nil {
		sendErr := a.Deps.Email.SendEmailVerification(user.Email, activeOTP.Token)
		if sendErr != nil {
			return &errors.FailedToSendEmail
		}

		return nil
	}

	otp, otpErr := auth.GenerateOTP(6)
	if otpErr != nil {
		return &errors.FailedToGenerateOTP
	}

	saveErr := SaveToken(a.DB, user.ID, *otp, models.EmailVerificationType, time.Now().Add(time.Minute*5).UTC())
	if saveErr != nil {
		return saveErr
	}

	sendErr := a.Deps.Email.SendEmailVerification(user.Email, *otp)
	if sendErr != nil {
		return &errors.FailedToSendEmail
	}

	return nil
}

func verifyEmailHelper(user *models.User, token string, db *gorm.DB) *errors.Error {
	if user.IsVerified {
		return &errors.EmailAlreadyVerified
	}

	otp, otpErr := GetToken(db, token, models.EmailVerificationType)
	if otpErr != nil {
		return otpErr
	}

	if otp.Token != token {
		return &errors.InvalidOTP
	}

	if otp.ExpiresAt.Before(time.Now().UTC()) {
		return &errors.OTPExpired
	}

	return nil
}

func (a *AuthService) VerifyEmail(emailBody models.VerifyEmailRequestBody) *errors.Error {
	if err := a.Validate.Struct(emailBody); err != nil {
		return &errors.FailedToValidateUser
	}

	user, err := users.GetUserByEmail(a.DB, emailBody.Email)
	if err != nil {
		return err
	}

	verifyErr := verifyEmailHelper(user, emailBody.Token, a.DB)
	if verifyErr != nil {
		return verifyErr
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	updateErr := users.UpdateEmailVerification(tx, user.ID)
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	deleteErr := DeleteToken(tx, emailBody.Token, models.EmailVerificationType)
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

func (a *AuthService) Logout(c *fiber.Ctx) *errors.Error {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	c.Set("Authorization ", "")

	// TODO: blacklist the refresh token

	return nil
}
