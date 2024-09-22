package lib

import (
	"errors"
	"strconv"
	"time"

	"example.com/config"
	"github.com/golang-jwt/jwt/v5"
)

type JwtAuthorizer struct{}

func (j *JwtAuthorizer) GenerateToken(email, userId string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    time.Now().Add(time.Minute * 5).Unix(),
		})

	appConfig := config.AppConfiguration()

	secretKey, err := appConfig.JwtSecretKey()

	if err != nil {
		return "", err
	}

	return token.SignedString([]byte(secretKey))
}

func (j *JwtAuthorizer) ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, correctSigningType := token.Method.(*jwt.SigningMethodHMAC)

		if !correctSigningType {
			return nil, errors.New("invalid token signing method")
		}

		appConfig := config.AppConfiguration()

		secretKey, err := appConfig.JwtSecretKey()

		if err != nil {
			return "", err
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, validClaimsType := parsedToken.Claims.(jwt.MapClaims)

	if !validClaimsType {
		return 0, errors.New("malformed token")
	}

	userId := claims["userId"].(string)

	parsedUserId, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		return 0, errors.New("invalid user id format")
	}

	return parsedUserId, nil
}

func NewJwtAuthorizer() *JwtAuthorizer {
	return &JwtAuthorizer{}
}
