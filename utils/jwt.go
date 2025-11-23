package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "asew234rtpoy90_322!3x"

func TokenGeneration(email string, id int) (string, error) {
	claims := &jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(59 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedInString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedInString, nil
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method for token creation")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return -1, errors.New("could not parse token")
	} 
	
	if !parsedToken.Valid{
		return -1, errors.New("token validation failed")
	}
	
	claims,ok:= parsedToken.Claims.(jwt.MapClaims) 
	if !ok{
		return -1, errors.New("invalid tokens claim")
	}
	
	id,ok:= claims["id"].(float64)
	if !ok{
		return -1, errors.New("id not found in the claims")
	}

	return int64(id), nil  
}
