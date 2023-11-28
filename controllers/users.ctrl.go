package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pioniry/entities"
	"pioniry/helpers"
	"pioniry/models/domain"
	"strconv"
)

type UserController struct {
	model domain.UserModel
}

type resUser struct {
	Email     string `json:"email"`
	Nip       string `json:"nip"`
	Image     string `json:"image"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	IsActive  int8   `json:"is_active"`
}

func (u UserController) Create(c echo.Context) error {
	var user entities.User
	var resp resUser

	err := c.Bind(&user)
	if err != nil {
		//return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusConflict)
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}

	err = c.Validate(&user)
	if err != nil {
		//return helpers.HandleErrorResponse(c, "Fill required field", http.StatusHTTPVersionNotSupported)
		return helpers.HandleErrorResponse(c, "Fill required field", http.StatusNotFound)
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		return helpers.HandleErrorResponse(c, "Hashing error", http.StatusNotFound)
	}
	mUser, erUser := u.model.Create(&user)

	if erUser != nil {
		return helpers.HandleErrorResponse(c, "Cannot save to Database", http.StatusBadRequest)
	} else {
		resp = resUser{
			Email:     mUser.Email,
			Nip:       mUser.Nip,
			Image:     mUser.Image,
			Firstname: mUser.Firstname,
			Lastname:  mUser.Lastname,
		}
	}
	return helpers.HandleSuccessResponse(c, "Create User", resp, http.StatusCreated)
}

func (u UserController) Update(c echo.Context) error {
	var user entities.User
	var resp resUser

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = c.Bind(&user)
	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}

	mUser, erUser := u.model.Update(id, &user)
	if erUser != nil {
		return helpers.HandleErrorResponse(c, "Cannot update to Database", http.StatusBadRequest)
	} else {
		resp = resUser{
			Email:     mUser.Email,
			Nip:       mUser.Nip,
			Image:     mUser.Image,
			Firstname: mUser.Firstname,
			Lastname:  mUser.Lastname,
		}
	}
	return helpers.HandleSuccessResponse(c, "Update User", resp, http.StatusAccepted)
}

func (u UserController) GetAllUser(c echo.Context) error {
	//var resp resUser
	users, err := u.model.GetUsers()

	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}
	return helpers.HandleSuccessResponse(c, "Get Users", users, 0)
}

func (u UserController) GetUser(c echo.Context) error {
	type usrGet struct {
		ID int `param:"id"`
	}
	var usr usrGet
	var resp resUser

	err := c.Bind(&usr)
	if err != nil {
		return err
	}

	user, errN := u.model.GetUserById(usr.ID)
	if errN != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	} else {
		resp = resUser{
			Email:     user.Email,
			Nip:       user.Nip,
			Image:     user.Image,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
	}

	return helpers.HandleSuccessResponse(c, "Get User by Id", resp, 0)

}

func (u UserController) Delete(c echo.Context) error {
	type usrDel struct {
		ID int `param:"id"`
	}

	var user usrDel
	err := c.Bind(&user)
	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusBadRequest)
	}

	err = u.model.Delete(user.ID)
	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}
	return helpers.HandleSuccessResponse(c, "Data Deleted", nil, http.StatusNoContent)
}
