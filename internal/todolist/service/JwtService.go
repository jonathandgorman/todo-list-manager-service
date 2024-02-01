package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtService struct{}

func (s *JwtService) GetToken(username string) (string, error) {
	signingKey := []byte("signing-key")
	expirationTime := time.Now().Add(2 * time.Minute).Unix()
	issuedAt := time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     "user",
		"iat":      issuedAt,
		"aud":      "todo-list-manager",
		"exp":      expirationTime,
	})

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JwtService) VerifyToken(token string) (bool, error) {
	signingKey := []byte("signing-key")
	parsedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	now := time.Now().Unix()
	username := parsedToken.Claims.(jwt.MapClaims)["username"].(string)
	expires := parsedToken.Claims.(jwt.MapClaims)["exp"].(float64)
	role := parsedToken.Claims.(jwt.MapClaims)["role"].(string)
	aud := parsedToken.Claims.(jwt.MapClaims)["aud"].(string)

	if (now < int64(expires)) && username == "jonathan.gorman" && role == "user" && aud == "todo-list-manager" {
		return true, nil
	}
	return false, nil
}
