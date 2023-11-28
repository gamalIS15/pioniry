package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pioniry/entities"
	"pioniry/helpers"
	"pioniry/models/domain"
	"strconv"
)

type PelatihanUnit struct {
	pelatihanModel domain.PelatihanModel
}

func (p PelatihanUnit) Create(c echo.Context) error {
	var pelatihan entities.Pelatihan
	err := c.Bind(&pelatihan)
	if err != nil {
		helpers.ReportError(err)
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}

	mPel, errN := p.pelatihanModel.Create(&pelatihan)
	if errN != nil {
		helpers.ReportError(errN)
		return helpers.HandleErrorResponse(c, "Save error", http.StatusBadRequest)
	}

	//Map Response to struct
	response := helpers.HandleResponsePelatihan(mPel)

	return helpers.HandleSuccessResponse(c, "Create Pelatihan", response, http.StatusCreated)
}

func (p PelatihanUnit) Update(c echo.Context) error {
	var pelatihan entities.Pelatihan
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := c.Bind(&pelatihan)
	if err != nil {
		//helpers.ReportError(err)
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}

	mPel, erPel := p.pelatihanModel.Update(id, &pelatihan)
	if erPel != nil {
		//helpers.ReportError(erPel)
		return helpers.HandleErrorResponse(c, "Cannot update to Database", http.StatusBadRequest)
	}

	return helpers.HandleSuccessResponse(c, "Update Pelatihan", mPel, http.StatusAccepted)

}
