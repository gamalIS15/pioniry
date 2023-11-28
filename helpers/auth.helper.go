package helpers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"pioniry/entities"
	"time"
)

func GetJWTSecret() string {
	return viper.GetString("JWTSECRETKEY")
	//return jwtSecretKey
}

func GetRefreshJWTSecret() string {
	return viper.GetString("JWTREFRESHSECRETKEY")
	//return jwtRefreshSecretKey
}

type Claims struct {
	Nip   string `json:"nip"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenAndSetCookie(user *entities.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}
	setTokenCookie(viper.GetString("ACCESSCOOKIENAME"), accessToken, exp, c)
	setUserCookie(user, exp, c)

	refreshToken, exp, err := generateRefreshToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(viper.GetString("REFRESHCOOKIENAME"), refreshToken, exp, c)
	return nil
}

func generateAccessToken(user *entities.User) (string, time.Time, error) {
	//Expiration token ( 1 hour)
	expirationTime := time.Now().Add(1 * time.Hour)
	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(user *entities.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
		Nip: user.Nip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil

}

func generateRefreshToken(user *entities.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	return generateToken(user, expirationTime, []byte(GetRefreshJWTSecret()))
}

func setTokenCookie(name, token string, expirationTime time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expirationTime
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func setUserCookie(user *entities.User, expirationTime time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Nip
	cookie.Expires = expirationTime
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func JWTErrorChecker(err error, c echo.Context) error {
	return HandleErrorResponse(c, "Invalid auth", http.StatusUnauthorized)
}
