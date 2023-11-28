package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pioniry/entities"
	"pioniry/helpers"
	"pioniry/models/domain"
	"pioniry/models/web"
	"time"
)

type SignInController struct {
	user domain.UserModel
}

func (s SignInController) SignIn(c echo.Context) error {
	var dataLogin = web.UserLogin{}

	storedUser := new(entities.User)
	err := c.Bind(&dataLogin)
	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}

	storedUser, err = s.user.GetUserByNip(dataLogin.Nip)

	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}

	//Update last login
	storedUser.LastLogin = time.Now()
	_, err = s.user.Update(int(storedUser.ID), storedUser)
	//fmt.Println(checkUser.Nip)

	err = s.user.CheckPassword(storedUser.Password, dataLogin.Password)
	if err != nil {
		return helpers.HandleErrorResponse(c, "Password incorrect", http.StatusBadRequest)
	}

	err = helpers.GenerateTokenAndSetCookie(storedUser, c)
	if err != nil {
		return helpers.HandleErrorResponse(c, "Token is incorrect", http.StatusUnauthorized)
	}
	return helpers.HandleSuccessResponse(c, "Success", nil, http.StatusAccepted)

}

func (s SignInController) SignOut(c echo.Context) error {
	cookieUser := new(http.Cookie)
	cookieUser.Name = "user"
	cookieUser.Value = "-"
	cookieUser.MaxAge = -1

	c.SetCookie(cookieUser)
	cookieToken := new(http.Cookie)
	cookieToken.Name = "access-token"
	cookieToken.Value = "-"
	cookieToken.MaxAge = -1

	c.SetCookie(cookieToken)

	cookieRToken := new(http.Cookie)
	cookieRToken.Name = "refresh-token"
	cookieRToken.Value = "-"
	cookieRToken.MaxAge = -1

	c.SetCookie(cookieRToken)
	return c.JSON(http.StatusMovedPermanently, "Success Logout")
}
