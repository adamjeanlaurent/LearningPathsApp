package security

import (
	"errors"
	"time"

	configuration "github.com/adamjeanlaurent/LearningPathsApp/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(configuration.GetJwtSecretKey())

const jwtExpiryTime time.Duration = time.Hour * 24

type JwtClaims struct {
	StableId string `json:"userId"`
	jwt.RegisteredClaims
}

func CreateNewJwt(stableId string) (string, error) {
	var expirationTime time.Time = time.Now().Add(jwtExpiryTime)

	claims := &JwtClaims{
		StableId: stableId,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJwt(tokenString string) (string, error) {
	claims := &JwtClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token is invalid")
	}

	return claims.StableId, nil
}
