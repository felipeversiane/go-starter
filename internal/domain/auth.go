package domain

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/felipeversiane/go-starter/internal/infra/config/response"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY         = "JWT_SECRET_KEY"
	JWT_REFRESH_SECRET_KEY = "JWT_SECRET_REFRESH_KEY"
)

func NewUserLogin(
	email, password string,
) (UserInterface, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return &user{
		email:    email,
		password: hashedPassword,
	}, nil
}

func (ud *user) GenerateToken() (string, string, *response.ErrorResponse) {
	access, err := ud.GenerateAcessToken()
	if err != nil {
		return "", "", err
	}
	refresh, err := ud.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func (ud *user) GenerateAcessToken() (string, *response.ErrorResponse) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":         ud.id,
		"email":      ud.email,
		"first_name": ud.firstName,
		"last_name":  ud.lastName,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", response.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}

func (ud *user) GenerateRefreshToken() (string, *response.ErrorResponse) {
	secret := os.Getenv(JWT_REFRESH_SECRET_KEY)

	refreshTokenClaims := jwt.MapClaims{
		"id":  ud.id,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", response.NewInternalServerError(
			fmt.Sprintf("error trying to generate refresh token, err=%s", err.Error()))
	}

	return refreshTokenString, nil
}

func VerifyAcessToken(tokenValue string) (UserInterface, *response.ErrorResponse) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, response.NewBadRequestError("invalid token")
	})
	if err != nil {
		return nil, response.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, response.NewUnauthorizedRequestError("invalid token")
	}

	return &user{
		id:        claims["id"].(string),
		email:     claims["email"].(string),
		firstName: claims["first_name"].(string),
		lastName:  claims["last_name"].(string),
	}, nil
}

func VerifyRefreshToken(tokenValue string) (UserInterface, *response.ErrorResponse) {
	secret := os.Getenv(JWT_REFRESH_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, response.NewBadRequestError("invalid token")
	})
	if err != nil {
		return nil, response.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, response.NewUnauthorizedRequestError("invalid token")
	}

	return &user{
		id: claims["id"].(string),
	}, nil
}

func RemoveBearerPrefix(tokenValue string) string {
	if strings.HasPrefix(tokenValue, "Bearer ") {
		return tokenValue[len("Bearer "):]
	}
	return tokenValue
}
