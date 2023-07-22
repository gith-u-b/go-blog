package util

import (
	"github.com/dgrijalva/jwt-go"
	"go-blog-step-by-step/pkg/setting"
	"golang.org/x/xerrors"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string)(string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt : expireTime.Unix(),
			Issuer : "go-blog-step-by-step",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	if err != nil {
		return token, xerrors.Errorf("Generate Token Failed: %w", err)
	}

	return token, nil
}


func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	if err != nil {
		return nil, xerrors.Errorf("Parse Token Failed: %w", err)
	}

	return nil, err
}

