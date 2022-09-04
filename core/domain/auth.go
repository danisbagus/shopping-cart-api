package domain

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AccessTokenClaims struct {
	UserID uint64 `json:"user_id"`
	RoleID uint64 `json:"role_id"`
	jwt.StandardClaims
}

type AuthToken struct {
	token *jwt.Token
}

const ACCESS_TOKEN_DURATION = time.Minute
const REFRESH_TOKEN_DURATION = time.Hour * 24 * 30
const HMAC_SAMPLE_SECRET = "shopping-cart-secret"

func (a AuthToken) NewAccessToken() (string, error) {
	signedString, err := a.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		return "", fmt.Errorf("cannot generate access token")
	}
	return signedString, nil
}

func GenerateAccessToken(userID, roleID uint64) (string, error) {

	claims := AccessTokenClaims{
		UserID: userID,
		RoleID: roleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}

	authToken := newAuthToken(claims)

	accessToken, err := authToken.NewAccessToken()
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func newAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}
