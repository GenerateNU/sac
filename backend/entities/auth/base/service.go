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
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/errs"
	"github.com/GenerateNU/sac/backend/types"

	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthService interface {
	GetRole(id string) (*models.UserRole, error)
	Register(ctx context.Context, userBody usersEntities.CreateUserRequestBody) (*models.User, *auth.Token, error)
	Login(ctx context.Context, userBody authEntities.LoginResponseBody) (*models.User, *auth.Token, error)
	Logout(ctx context.Context, userID uuid.UUID) error

	SendCode(email string) error
	VerifyEmail(emailBody VerifyEmailRequestBody) error

	ForgotPassword(email string) error
	VerifyPasswordResetToken(passwordBody VerifyPasswordResetTokenRequestBody) error
}

type AuthServiceHandler struct {
	types.ServiceParams
}

func NewAuthService(serviceParams types.ServiceParams) AuthService {
	return &AuthServiceHandler{serviceParams}
}

func (a *AuthServiceHandler) Login(ctx context.Context, loginBody authEntities.LoginResponseBody) (*models.User, *auth.Token, error) {
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

	storeRefreshTokenCtx, storeRefreshTokenCancel := context.WithTimeoutCause(ctx, constants.REDIS_TIMEOUT, errs.ErrRedisTimeout)
	defer storeRefreshTokenCancel()
	if err := a.Stores.Refresher.Set(storeRefreshTokenCtx, string(tokens.RefreshToken), user.ID.String(), constants.REFRESH_TOKEN_EXPIRY); err != nil {
		return nil, nil, err
	}

	return user, tokens, nil
}

func (a *AuthServiceHandler) Register(ctx context.Context, userBody usersEntities.CreateUserRequestBody) (*models.User, *auth.Token, error) {
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

	if err := a.Integrations.Email.SendWelcomeEmail(fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email); err != nil {
		return nil, nil, fmt.Errorf("failed to send welcome email: %w", err)
	}

	user, err = CreateUser(a.DB, user)
	if err != nil {
		return nil, nil, err
	}

	if err := a.SendCode(user.Email); err != nil {
		return nil, nil, err
	}

	loginCtx, loginCancel := context.WithTimeoutCause(ctx, constants.REDIS_TIMEOUT, errs.ErrRedisTimeout)
	defer loginCancel()
	_, tokens, err := a.Login(loginCtx, authEntities.LoginResponseBody{Email: user.Email, Password: userBody.Password})
	if err != nil {
		return nil, nil, err
	}

	storeRefreshTokenCtx, storeRefreshTokenCancel := context.WithTimeoutCause(ctx, constants.REDIS_TIMEOUT, errs.ErrRedisTimeout)
	defer storeRefreshTokenCancel()
	if err := a.Stores.Refresher.Set(storeRefreshTokenCtx, string(tokens.RefreshToken), user.ID.String(), constants.REFRESH_TOKEN_EXPIRY); err != nil {
		return nil, nil, err
	}

	return user, tokens, nil
}

func (a *AuthServiceHandler) GetRole(id string) (*models.UserRole, error) {
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

func (a *AuthServiceHandler) ForgotPassword(email string) error {
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

	if err := a.Integrations.Email.SendPasswordResetEmail(user.FirstName, user.Email, *token); err != nil {
		return err
	}

	return nil
}

func (a *AuthServiceHandler) VerifyPasswordResetToken(passwordBody VerifyPasswordResetTokenRequestBody) error {
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

func (a *AuthServiceHandler) SendCode(email string) error {
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

	otp, err := auth.GenerateOTP(constants.OTP_LENGTH)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	if err := SaveToken(a.DB, user.ID, *otp, models.EmailVerificationType, time.Now().Add(constants.OTP_EXPIRY).UTC()); err != nil {
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

func (a *AuthServiceHandler) VerifyEmail(emailBody VerifyEmailRequestBody) error {
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

func (a *AuthServiceHandler) Logout(ctx context.Context, userID uuid.UUID) error {
	deleteRefreshTokenCtx, deleteRefreshTokenCancel := context.WithTimeoutCause(ctx, constants.REDIS_TIMEOUT, errs.ErrRedisTimeout)
	defer deleteRefreshTokenCancel()
	if err := a.Stores.Refresher.Delete(deleteRefreshTokenCtx, userID.String()); err != nil {
		return err
	}

	return nil
}
