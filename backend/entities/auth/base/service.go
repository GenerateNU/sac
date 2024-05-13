package base

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	GetRole(id string) (*models.UserRole, error)
	Login(userBody models.LoginUserResponseBody) (*models.User, *auth.Token, error)
	Refresh(refreshToken string) (*auth.Token, error)
	Logout(c *fiber.Ctx) error

	SendCode(email string) error
	VerifyEmail(emailBody models.VerifyEmailRequestBody) error

	ForgotPassword(email string) error
	VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) error
}

type AuthService struct {
	types.ServiceParams
}

func NewAuthService(serviceParams types.ServiceParams) AuthServiceInterface {
	return &AuthService{serviceParams}
}

func (a *AuthService) Login(loginBody models.LoginUserResponseBody) (*models.User, *auth.Token, error) {
	if err := utilities.Validate(a.Validate, loginBody); err != nil {
		return nil, nil, err
	}

	if err := auth.ValidatePassword(loginBody.Password); err != nil {
		return nil, nil, err
	}

	user, err := users.GetUserByEmail(a.DB, loginBody.Email)
	if err != nil {
		return nil, nil, err
	}

	correct, err := auth.CompareHash(loginBody.Password, user.PasswordHash)
	if err != nil || !correct {
		return nil, nil, err
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
		return nil, nil, err
	}

	return user, tokens, nil
}

func (a *AuthService) Refresh(refreshToken string) (*auth.Token, error) {
	claims, err := a.JWT.ExtractClaims(refreshToken, auth.RefreshToken)
	if err != nil {
		return nil, err
	}

	role, err := a.GetRole(claims["iss"].(string))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return tokens, nil
}

func (a *AuthService) GetRole(id string) (*models.UserRole, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	user, err := users.GetUser(a.DB, *idAsUUID)
	if err != nil {
		return nil, err
	}

	return &user.Role, nil
}

func (a *AuthService) ForgotPassword(email string) error {
	user, err := users.GetUserByEmail(a.DB, email)
	if err != nil {
		return nil
	}

	activeToken, err := GetActiveTokenByUserID(a.DB, user.ID, models.PasswordResetType)
	if err != nil {
		if !utilities.IsNotFound(err) {
			return err
		}
	}

	if activeToken != nil {
		err := a.Integrations.Email.SendPasswordResetEmail(user.FirstName, user.Email, activeToken.Token)
		if err != nil {
			return err
		}

		return nil
	}

	token, err := auth.GenerateURLSafeToken(64)
	if err != nil {
		return err
	}

	if err := SaveToken(a.DB, user.ID, *token, models.PasswordResetType, time.Now().Add(time.Hour*24).UTC()); err != nil {
		return err
	}

	// TODO fmt.Sprintf token to URl
	if err := a.Integrations.Email.SendPasswordResetEmail(user.FirstName, user.Email, *token); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) VerifyPasswordResetToken(passwordBody models.VerifyPasswordResetTokenRequestBody) error {
	if err := utilities.Validate(a.Validate, passwordBody); err != nil {
		return err
	}

	token, err := GetToken(a.DB, passwordBody.Token, models.PasswordResetType)
	if err != nil {
		return err
	}

	if token.ExpiresAt.Before(time.Now().UTC()) {
		return utilities.Unauthorized()
	}

	hash, err := auth.ComputeHash(passwordBody.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := users.UpdatePassword(tx, token.UserID, *hash); err != nil {
		tx.Rollback()
		return err
	}

	if err := DeleteToken(tx, passwordBody.Token, models.PasswordResetType); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (a *AuthService) SendCode(email string) error {
	user, err := users.GetUserByEmail(a.DB, email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return fmt.Errorf("user is already verified")
	}

	activeOTP, err := GetActiveTokenByUserID(a.DB, user.ID, models.EmailVerificationType)
	if err != nil {
		if !utilities.IsNotFound(err) {
			return err
		}
	}

	if activeOTP != nil {
		if err := a.Integrations.Email.SendEmailVerification(user.Email, activeOTP.Token); err != nil {
			return err
		}

		return nil
	}

	otp, err := auth.GenerateOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	if err := SaveToken(a.DB, user.ID, *otp, models.EmailVerificationType, time.Now().Add(time.Minute*5).UTC()); err != nil {
		return err
	}

	if err := a.Integrations.Email.SendEmailVerification(user.Email, *otp); err != nil {
		return err
	}

	return nil
}

func verifyEmailHelper(user *models.User, token string, db *gorm.DB) error {
	if user.IsVerified {
		return fmt.Errorf("user is already verified")
	}

	otp, otpErr := GetToken(db, token, models.EmailVerificationType)
	if otpErr != nil {
		return otpErr
	}

	if otp.Token != token {
		slog.Error("invalid otp", "otp", otp.Token, "token", token)
		return utilities.BadRequest(errors.New("invalid otp"))
	}

	if otp.ExpiresAt.Before(time.Now().UTC()) {
		return utilities.BadRequest(errors.New("otp expired"))
	}

	return nil
}

func (a *AuthService) VerifyEmail(emailBody models.VerifyEmailRequestBody) error {
	if err := utilities.Validate(a.Validate, emailBody); err != nil {
		return err
	}

	user, err := users.GetUserByEmail(a.DB, emailBody.Email)
	if err != nil {
		return err
	}

	if err := verifyEmailHelper(user, emailBody.Token, a.DB); err != nil {
		return err
	}

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := users.UpdateEmailVerification(tx, user.ID); err != nil {
		tx.Rollback()
		return err
	}

	if err := DeleteToken(tx, emailBody.Token, models.EmailVerificationType); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (a *AuthService) Logout(c *fiber.Ctx) error {
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
