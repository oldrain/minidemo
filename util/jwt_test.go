package util

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestJwtTokenExpireAfter(t *testing.T) {
	now := time.Now()
	key := "xxx"

	claimsAfter := &JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Unix() + 1000,
		},
		PartnerId:    "xx",
		AppId:        "xx",
		CustomerCode: "22",
	}
	
	tokenAfter, err := JwtNewToken(claimsAfter, key)
	if err != nil {
		t.Error(err)
	}

	claimsParseAfter := new(JwtClaims)
	err = JwtParse(tokenAfter, claimsParseAfter, key)
	if err != nil {
		t.Error(err)
	}
	if claimsParseAfter.ExpiresAt <= now.Unix() {
		t.Error("should not expired")
	}

	if claimsParseAfter.VerifyExpiresAt(now.Unix() + 1000, false) == false {
		t.Error("should expired")
	}
	if claimsParseAfter.VerifyExpiresAt(now.Unix() + 1001, true) {
		t.Error("should expired")
	}
	if claimsParseAfter.VerifyExpiresAt(now.Unix() + 10001, true) {
		t.Error("should expired")
	}
	if claimsParseAfter.VerifyExpiresAt(now.Unix() + 999, false) == false {
		t.Error("should expired")
	}

}

func TestJwtTokenExpireBefore(t *testing.T) {
	now := time.Now()
	key := "xxx"

	claimsBefore := &JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Unix() - 1000,
		},
		PartnerId:      "xx",
		AppId:          "xx",
		CustomerCode:   "22",
	}

	tokenBefore, err := JwtNewToken(claimsBefore, key)
	if err != nil {
		t.Error(err)
	}

	claimsParseBefore := new(JwtClaims)
	err = JwtParse(tokenBefore, claimsParseBefore, key)
	if err == nil {
		t.Error("should expired")
	}
	if claimsParseBefore.ExpiresAt > now.Unix() {
		t.Error("should expired")
	}
}
