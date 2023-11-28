package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"pioniry/entities"
	"pioniry/helpers"
	"time"
)

func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user") == nil {
			return next(c)
		}
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(*helpers.Claims)
		//Create new token 15 before expire
		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 15*time.Minute {
			re, err := c.Cookie(viper.GetString("REFRESHCOOKIENAME"))
			if err == nil && re != nil {
				tkn, err := jwt.ParseWithClaims(re.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(helpers.GetRefreshJWTSecret()), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}

				if tkn != nil && tkn.Valid {
					_ = helpers.GenerateTokenAndSetCookie(&entities.User{
						Nip: claims.Nip,
					}, c)
				}

			}
		}

		return next(c)
	}
}
