package base

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/constants"
	authEntities "github.com/GenerateNU/sac/backend/entities/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	usersEntities "github.com/GenerateNU/sac/backend/entities/users/base"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthServicer interface {
	GetRole(ctx context.Context, id string) (*models.UserRole, error)
	Register(ctx context.Context, userBody usersEntities.CreateUserRequestBody) (*models.User, *auth.Token, error)
	Login(ctx context.Context, userBody authEntities.LoginResponseBody) (*models.User, *auth.Token, error)
	Refresh(ctx context.Context, refreshToken string) (*auth.Token, error)
	Logout(ctx context.Context, c *fiber.Ctx) error

	SendCode(ctx context.Context, email string) error
	VerifyEmail(ctx context.Context, emailBody VerifyEmailRequestBody) error

	ForgotPassword(ctx context.Context, email string) error
	VerifyPasswordResetToken(ctx context.Context, passwordBody VerifyPasswordResetTokenRequestBody) error
}

type AuthService struct {
	types.ServiceParams
}

func NewAuthService(serviceParams types.ServiceParams) AuthServicer {
	return &AuthService{serviceParams}
}

func (a *AuthService) Login(ctx context.Context, loginBody authEntities.LoginResponseBody) (*models.User, *auth.Token, error) {
	if err := utilities.Validate(a.Validate, loginBody, *utilities.NewMaybeError("password", auth.ValidatePassword(loginBody.Password))); err != nil {
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

func (a *AuthService) Register(ctx context.Context, userBody usersEntities.CreateUserRequestBody) (*models.User, *auth.Token, error) {
	if err := utilities.Validate(a.Validate, userBody, *utilities.NewMaybeError("password", auth.ValidatePassword(userBody.Password))); err != nil {
		return nil, nil, err
	}

	user, err := utilities.MapJsonTags(userBody, &models.User{})
	if err != nil {
		return nil, nil, err
	}

	user.Role = models.Student

	passwordHash, err := auth.ComputeHash(userBody.Password)
	if err != nil {
		return nil, nil, err
	}

	user.Email = utilities.NormalizeEmail(userBody.Email)
	user.PasswordHash = *passwordHash

	sendWelcomeCtx, sendWelcomeCancel := context.WithTimeout(ctx, constants.THIRD_PARTY_TIMEOUT)
	defer sendWelcomeCancel()

	if err := a.Integrations.Email.SendWelcomeEmail(sendWelcomeCtx, fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email); err != nil {
		return nil, nil, fmt.Errorf("failed to send welcome email: %w", err)
	}

	createUserCtx, createUserCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer createUserCancel()

	user, err = CreateUser(createUserCtx, a.DB, user)
	if err != nil {
		return nil, nil, err
	}

	sendCodeCtx, sendCodeCancel := context.WithTimeout(ctx, constants.THIRD_PARTY_TIMEOUT)
	defer sendCodeCancel()

	if err := a.SendCode(sendCodeCtx, user.Email); err != nil {
		return nil, nil, err
	}

	loginCtx, loginCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer loginCancel()

	_, tokens, err := a.Login(loginCtx, authEntities.LoginResponseBody{Email: user.Email, Password: userBody.Password})
	if err != nil {
		return nil, nil, err
	}

	return user, tokens, nil
}

func (a *AuthService) Refresh(ctx context.Context, refreshToken string) (*auth.Token, error) {
	claims, err := a.JWT.ExtractClaims(refreshToken, auth.RefreshToken)
	if err != nil {
		return nil, err
	}

	getRoleCtx, getRoleCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer getRoleCancel()

	role, err := a.GetRole(getRoleCtx, claims["iss"].(string))
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

func (a *AuthService) GetRole(ctx context.Context, id string) (*models.UserRole, error) {
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

func (a *AuthService) ForgotPassword(ctx context.Context, email string) error {
	user, err := users.GetUserByEmail(a.DB, email)
	if err != nil {
		return nil
	}

	getActiveTokenCtx, getActiveTokenCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer getActiveTokenCancel()

	activeToken, err := GetActiveTokenByUserID(getActiveTokenCtx, a.DB, user.ID, models.PasswordResetType)
	if err != nil {
		if !utilities.IsNotFound(err) {
			return err
		}
	}

	if activeToken != nil {
		sendEmailCtx, sendEmailCancel := context.WithTimeout(ctx, constants.THIRD_PARTY_TIMEOUT)
		defer sendEmailCancel()

		err := a.Integrations.Email.SendPasswordResetEmail(sendEmailCtx, user.FirstName, user.Email, activeToken.Token)
		if err != nil {
			return err
		}

		return nil
	}

	token, err := auth.GenerateURLSafeToken(64)
	if err != nil {
		return err
	}

	saveTokenCtx, saveTokenCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer saveTokenCancel()

	if err := SaveToken(saveTokenCtx, a.DB, user.ID, *token, models.PasswordResetType, time.Now().Add(time.Hour*24).UTC()); err != nil {
		return err
	}

	sendPasswordResetCtx, sendPasswordResetCancel := context.WithTimeout(ctx, constants.THIRD_PARTY_TIMEOUT)
	defer sendPasswordResetCancel()

	if err := a.Integrations.Email.SendPasswordResetEmail(sendPasswordResetCtx, user.FirstName, user.Email, *token); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) VerifyPasswordResetToken(ctx context.Context, passwordBody VerifyPasswordResetTokenRequestBody) error {
	if err := utilities.Validate(a.Validate, passwordBody); err != nil {
		return err
	}

	getTokenCtx, getTokenCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer getTokenCancel()

	token, err := GetToken(getTokenCtx, a.DB, passwordBody.Token, models.PasswordResetType)
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

	passwordResetCtx, passwordResetCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer passwordResetCancel()

	return PasswordReset(passwordResetCtx, a.DB, token.UserID, *hash, passwordBody.Token)
}

func (a *AuthService) SendCode(ctx context.Context, email string) error {
	user, err := users.GetUserByEmail(a.DB, email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return fmt.Errorf("user is already verified")
	}

	activeOTPCtx, activeOTPCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer activeOTPCancel()

	activeOTP, err := GetActiveTokenByUserID(activeOTPCtx, a.DB, user.ID, models.EmailVerificationType)
	if err != nil {
		if !utilities.IsNotFound(err) {
			return err
		}
	}

	if activeOTP != nil {
		sendEmailCtx, sendEmailCancel := context.WithTimeout(ctx, constants.THIRD_PARTY_TIMEOUT)
		defer sendEmailCancel()

		if err := a.Integrations.Email.SendEmailVerification(sendEmailCtx, user.Email, activeOTP.Token); err != nil {
			return err
		}

		return nil
	}

	otp, err := auth.GenerateOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	saveTokenCtx, saveTokenCancel := context.WithTimeout(ctx, constants.QUERY_TIMEOUT)
	defer saveTokenCancel()

	if err := SaveToken(saveTokenCtx, a.DB, user.ID, *otp, models.EmailVerificationType, time.Now().Add(time.Minute*5).UTC()); err != nil {
		return err
	}

	sendEmailCtx, sendEmailCancel := context.WithTimeout(ctx, constants.THIRD_PARTY_TIMEOUT)
	defer sendEmailCancel()

	if err := a.Integrations.Email.SendEmailVerification(sendEmailCtx, user.Email, *otp); err != nil {
		return err
	}

	return nil
}

func verifyEmailHelper(ctx context.Context, user *models.User, token string, db *gorm.DB) error {
	if user.IsVerified {
		return fmt.Errorf("user is already verified")
	}

	otp, otpErr := GetToken(ctx, db, token, models.EmailVerificationType)
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

func (a *AuthService) VerifyEmail(ctx context.Context, emailBody VerifyEmailRequestBody) error {
	if err := utilities.Validate(a.Validate, emailBody); err != nil {
		return err
	}

	user, err := users.GetUserByEmail(a.DB, emailBody.Email)
	if err != nil {
		return err
	}

	if err := verifyEmailHelper(ctx, user, emailBody.Token, a.DB); err != nil {
		return err
	}

	return VerifyEmail(ctx, a.DB, user.ID, emailBody.Token)
}

func (a *AuthService) Logout(ctx context.Context, c *fiber.Ctx) error {
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
