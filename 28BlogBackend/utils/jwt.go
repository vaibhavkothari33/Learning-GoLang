package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	UserID   primitive.ObjectID `json:"user_id"`
	Email    string             `json : "email"`
	Username string             `json:"username`
	jwt.RegisteredClaims
}

func GenerateToken(userID primitive.ObjectID, email, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil

}

func ValidateToken(tokenString string) (*Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil

	})
	if err != nil {
		log.Fatal(err)
	}
	if !token.Valid {
		return nil, errors.New("invalid")
	}

	return claims, nil

}
