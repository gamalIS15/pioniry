package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"pioniry/helpers"
	"pioniry/models/domain"
	"strings"
)

type PegawaiController struct {
	model domain.PegawaiModel
}

func (p PegawaiController) GetPegawai(c echo.Context) error {
	//var pegawai entities.Pegawai
	type pegawaiReq struct {
		Nip string `query:"nip"`
	}
	var pegawai pegawaiReq
	err := c.Bind(&pegawai)

	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}

	sUser, errN := p.model.GetPegawaiByNip(pegawai.Nip)
	if errN != nil {
		return helpers.HandleErrorResponse(c, "Cannot get data pegawai", http.StatusBadRequest)
	}

	return helpers.HandleSuccessResponse(c, "User By Nip", sUser, http.StatusOK)
}

func (p PegawaiController) SearchPegawai(c echo.Context) error {
	//var pegawai entities.Pegawai
	//q := c.QueryParam("search")

	type req struct {
		Search string `query:"search"`
		Limit  int    `query:"limit"`
		Offset int    `query:"offset"`
		Order  string `query:"order"`
	}
	var r req
	err := c.Bind(&r)

	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusNotFound)
	}

	//Replace _ in Order
	r.Order = strings.Replace(r.Order, "_", " ", 1)

	sUser, errN := p.model.GetPegawaiByParam(r.Search, r.Limit, r.Offset, r.Order)
	fmt.Println(r)
	if errN != nil {
		return helpers.HandleErrorResponse(c, "Cannot get data pegawai", http.StatusBadRequest)
	}

	return helpers.HandleSuccessResponse(c, "Search User", sUser, http.StatusOK)
}
