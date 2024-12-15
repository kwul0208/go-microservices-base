package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	models "github.com/kwul0208/user/pkg/model"
)

type JWTWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Id    int64
	Email string
}

// GenerateToken generates a JWT token for the given user and role
func (w *JWTWrapper) GenerateToken(user models.User, role string) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(w.ExpirationHours))),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates a signed JWT token and extracts claims
func (w *JWTWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return nil, errors.New("failed to parse JWT token: " + err.Error())
	}

	// Ensure the token is valid and extract claims
	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid JWT token or claims")
	}

	// Check if the token has expired
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("JWT token has expired")
	}

	return claims, nil
}
