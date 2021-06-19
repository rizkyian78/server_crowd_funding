package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenEncoded string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID string) (string, error) {
	godotenv.Load(".env")
	SECRET_KEY := []byte(os.Getenv("secret_key"))
	payload := jwt.MapClaims{}
	payload["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return fmt.Sprintf("Bearer %s", signedToken), nil
}

func (s *jwtService) ValidateToken(tokenEncoded string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenEncoded, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(os.Getenv("secret_key")), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
