package token

import (
	"errors"
	"strings"
	"time"

	"wtg_desktop/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey       = []byte(config.GetJwtSecret())
	refreshSecret   = []byte(config.GetJwtRefreshSecret())
	accessTokenTTL  = time.Minute * 15
	refreshTokenTTL = time.Hour * 24 * 7
)

type Claims struct {
	UserID int      `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

func GenerateTokens(userID int, roles []string) (accessToken string, refreshToken string, err error) {
	now := time.Now()

	accessClaims := &Claims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	refreshClaims := &Claims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(refreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(refreshSecret)
	return
}

func ParseToken(tokenStr string, isRefresh bool) (*Claims, error) {
	key := secretKey
	if isRefresh {
		key = refreshSecret
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ExtractTokenFromHeader(header string) (string, error) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}
	return parts[1], nil
}
