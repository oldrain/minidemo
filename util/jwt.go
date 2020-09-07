package util

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	Token string
}

type JwtClaims struct {
	jwt.StandardClaims

	PartnerId string
	AppId string
	CustomerCode string
}

func JwtClaimsCreate(partnerId, appId, customerCode string, expireAt int64) *JwtClaims {
	return &JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
		PartnerId:      partnerId,
		AppId:          appId,
		CustomerCode:   customerCode,
	}
}

func JwtNewToken(claims *JwtClaims, key string) (*JwtToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(key)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &JwtToken{Token:tokenStr}, nil
}

func JwtParse(jwtToken *JwtToken, claims *JwtClaims, key string) error {
	token, err := jwt.ParseWithClaims(jwtToken.Token, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(key), nil
	})
	if token == nil {
		return err
	}
	if _, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return nil
	}
	return err
}
