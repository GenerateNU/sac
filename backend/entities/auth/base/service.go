package base

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
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

type AuthServiceInterface interface {
	GetRole(id string) (*models.UserRole, error)
	Register(userBody usersEntities.CreateUserRequestBody) (*models.User, *auth.Token, error)
	Login(userBody authEntities.LoginResponseBody) (*models.User, *auth.Token, error)
	Refresh(refreshToken string) (*auth.Token, error)
	Logout(c *fiber.Ctx) error

	SendCode(email string) error
	VerifyEmail(emailBody VerifyEmailRequestBody) error

	ForgotPassword(email string) error
	VerifyPasswordResetToken(passwordBody VerifyPasswordResetTokenRequestBody) error
}

type AuthService struct {
	types.ServiceParams
}

func NewAuthService(serviceParams types.ServiceParams) AuthServiceInterface {
	return &AuthService{serviceParams}
}

func (a *AuthService) Login(loginBody authEntities.LoginResponseBody) (*models.User, *auth.Token, error) {
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

	if err := a.Stores.ActiveToken.StoreRefreshToken(string(tokens.RefreshToken), user.ID.String(), constants.REFRESH_TOKEN_EXPIRY); err != nil {
		return nil, nil, err
	}

	return user, tokens, nil
}

func (a *AuthService) Register(userBody usersEntities.CreateUserRequestBody) (*models.User, *auth.Token, error) {
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

	_, tokens, err := a.Login(authEntities.LoginResponseBody{Email: user.Email, Password: userBody.Password})
	if err != nil {
		return nil, nil, err
	}

	return user, tokens, nil
}

func (a *AuthService) Refresh(refreshToken string) (*auth.Token, error) {
	isActive, err := a.Stores.ActiveToken.IsActive(refreshToken)
	if err != nil {
		return nil, err
	}

	if !isActive {
		return nil, utilities.Unauthorized()
	}

	isblacklisted, err := a.Stores.Blacklist.IsTokenBlacklisted(refreshToken)
	if err != nil {
		return nil, err
	}

	if isblacklisted {
		return nil, fmt.Errorf("refresh token is blacklisted")
	}

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

	if err := a.Stores.Blacklist.BlacklistToken(refreshToken, time.Until(time.Unix(int64(claims["exp"].(float64)), 0))); err != nil {
		return nil, err
	}

	if err := a.Stores.ActiveToken.DeleteRefreshToken(refreshToken); err != nil {
		return nil, err
	}

	if err := a.Stores.ActiveToken.StoreRefreshToken(string(tokens.RefreshToken), claims["iss"].(string), constants.REFRESH_TOKEN_EXPIRY); err != nil {
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

	if err := a.Integrations.Email.SendPasswordResetEmail(user.FirstName, user.Email, *token); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) VerifyPasswordResetToken(passwordBody VerifyPasswordResetTokenRequestBody) error {
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

func (a *AuthService) VerifyEmail(emailBody VerifyEmailRequestBody) error {
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
	accessToken := c.Get("Authorization")
	if accessToken == "" {
		return utilities.Unauthorized()
	}

	token, err := utilities.SplitBearerToken(accessToken)
	if err != nil {
		return utilities.Unauthorized()
	}

	blacklist(a, token)

	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return utilities.Unauthorized()
	}

	isActive, err := a.Stores.ActiveToken.IsActive(refreshToken)
	if err != nil {
		return err
	}

	if !isActive {
		return utilities.Unauthorized()
	}

	blacklist(a, &refreshToken)

	if err := a.Stores.ActiveToken.DeleteRefreshToken(refreshToken); err != nil {
		slog.Error("Failed to delete refresh token", "error", err)
	}

	auth.ExpireResponseTokens(c)

	return utilities.FiberMessage(c, http.StatusOK, "success")
}

func blacklist(a *AuthService, token *string) {
	claims, err := a.JWT.ExtractClaims(*token, auth.AccessToken)
	if err != nil {
		slog.Warn("Invalid token", "error", err)
	} else {
		exp, ok := claims["exp"].(float64)
		if ok {
			_ = a.Stores.Blacklist.BlacklistToken(*token, time.Until(time.Unix(int64(exp), 0)))
		}
	}
}
