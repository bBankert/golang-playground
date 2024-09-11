package lib

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email, userId string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    time.Now().Add(time.Minute * 5).Unix(),
		})

	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, correctSigningType := token.Method.(*jwt.SigningMethodHMAC)

		if !correctSigningType {
			return nil, errors.New("invalid token signing method")
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
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
