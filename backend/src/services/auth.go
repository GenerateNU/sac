package services

import (
	"fmt"
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthServiceInterface interface {
	GetRole(id string) (*models.UserRole, *errors.Error)
	SetResponseTokens(c *fiber.Ctx, tokens *auth.Token) *errors.Error
	Login(userBody models.LoginUserResponseBody) (*models.User, *auth.Token, *errors.Error)
	Refresh(refreshToken string) (*auth.Token, *errors.Error)

	SendCode(email string) *errors.Error
	VerifyEmail(emailBody models.VerifyEmailRequestBody) *errors.Error

	ForgotPassword(email string) *errors.Error
	VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) *errors.Error
}

type AuthService struct {
	types.ServiceParams
}

func NewAuthService(serviceParams types.ServiceParams) *AuthService {
	return &AuthService{serviceParams}
}

// TODO: organize this
func (a *AuthService) SetResponseTokens(c *fiber.Ctx, tokens *auth.Token) *errors.Error {
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "refresh_token",
	// 	Value:    "",
	// 	Expires:  time.Now().Add(-time.Hour),
	// 	HTTPOnly: true,
	// })
	// c.Set("Authorization", "")

	// Set the tokens in the response
	// should also blacklist the old refresh and access tokens

	c.Set("Authorization", fmt.Sprintf("Bearer %s", tokens.AccessToken))
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    string(tokens.RefreshToken),
		Expires:  time.Now().Add(time.Hour * time.Duration(a.AuthSettings.RefreshTokenExpiry)),
		HTTPOnly: true,
	})

	return nil
}

func (a *AuthService) Login(loginBody models.LoginUserResponseBody) (*models.User, *auth.Token, *errors.Error) {
	if err := a.Validate.Struct(loginBody); err != nil {
		return nil, nil, &errors.FailedToValidateUser
	}

	user, getUserByEmailErr := transactions.GetUserByEmail(a.DB, loginBody.Email)
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

	user, err := transactions.GetUser(a.DB, *idAsUUID)
	if err != nil {
		return nil, err
	}

	return &user.Role, nil
}

func (a *AuthService) ForgotPassword(email string) *errors.Error {
	user, err := transactions.GetUserByEmail(a.DB, email)
	if err != nil {
		return nil
	}

	activeToken, tokenErr := transactions.GetActiveTokenByUserID(a.DB, user.ID, models.PasswordResetType)
	if tokenErr != nil {
		if tokenErr != &errors.TokenNotFound {
			return tokenErr
		}
	}

	if activeToken != nil {
		sendErr := a.Resend.SendPasswordResetEmail(user.FirstName, user.Email, activeToken.Token)
		if sendErr != nil {
			return &errors.FailedToSendEmail
		}

		return nil
	}

	token, generateErr := auth.GenerateURLSafeToken(64)
	if generateErr != nil {
		return &errors.FailedToGenerateToken
	}

	saveErr := transactions.SaveToken(a.DB, user.ID, *token, models.PasswordResetType, time.Now().Add(time.Hour*24).UTC())
	if saveErr != nil {
		return saveErr
	}

	sendErr := a.Resend.SendPasswordResetEmail(user.FirstName, user.Email, *token)
	if sendErr != nil {
		return sendErr
	}

	return nil
}

func (a *AuthService) VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) *errors.Error {
	if err := a.Validate.Struct(passwordBody); err != nil {
		return &errors.FailedToValidateUser
	}

	token, tokenErr := transactions.GetToken(a.DB, passwordBody.Token, models.PasswordResetType)
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

	updateErr := transactions.UpdatePassword(tx, token.UserID, *hash)
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	deleteErr := transactions.DeleteToken(tx, passwordBody.Token, models.PasswordResetType)
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
	user, err := transactions.GetUserByEmail(a.DB, email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return &errors.EmailAlreadyVerified
	}

	activeOTP, tokenErr := transactions.GetActiveTokenByUserID(a.DB, user.ID, models.EmailVerificationType)
	if tokenErr != nil {
		if tokenErr != &errors.TokenNotFound {
			return tokenErr
		}
	}

	if activeOTP != nil {
		sendErr := a.Resend.SendEmailVerification(user.Email, activeOTP.Token)
		if sendErr != nil {
			return &errors.FailedToSendEmail
		}

		return nil
	}

	otp, otpErr := auth.GenerateOTP(6)
	if otpErr != nil {
		return &errors.FailedToGenerateOTP
	}

	saveErr := transactions.SaveToken(a.DB, user.ID, *otp, models.EmailVerificationType, time.Now().Add(time.Minute*5).UTC())
	if saveErr != nil {
		return saveErr
	}

	sendErr := a.Resend.SendEmailVerification(user.Email, *otp)
	if sendErr != nil {
		return &errors.FailedToSendEmail
	}

	return nil
}

/* trunk-ignore(golangci-lint/cyclop) */
func (a *AuthService) VerifyEmail(emailBody models.VerifyEmailRequestBody) *errors.Error {
	if err := a.Validate.Struct(emailBody); err != nil {
		return &errors.FailedToValidateUser
	}

	user, err := transactions.GetUserByEmail(a.DB, emailBody.Email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return &errors.EmailAlreadyVerified
	}

	otp, otpErr := transactions.GetToken(a.DB, emailBody.Token, models.EmailVerificationType)
	if otpErr != nil {
		return otpErr
	}

	if otp.Token != emailBody.Token {
		return &errors.InvalidOTP
	}

	if otp.ExpiresAt.Before(time.Now().UTC()) {
		return &errors.OTPExpired
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	updateErr := transactions.UpdateEmailVerification(tx, user.ID)
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	deleteErr := transactions.DeleteToken(tx, emailBody.Token, models.EmailVerificationType)
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
