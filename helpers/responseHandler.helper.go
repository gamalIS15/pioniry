package helpers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pioniry/entities"
	"pioniry/models/web"
)

func HandleSuccessResponse(c echo.Context, message string, data interface{}, status int) error {
	if status == 0 {
		status = http.StatusOK
	}
	res := web.Response{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
	return c.JSON(status, res)
}

func HandleErrorResponse(c echo.Context, message string, status int) error {
	res := web.Response{
		Status:  "Failed",
		Message: message,
	}
	return c.JSON(status, res)
}

func HandleResponsePelatihan(data *entities.Pelatihan) interface{} {
	res := web.ResponsePelatihanInternal{
		IdPelatihan:   data.KodePelatihan,
		NamaPelatihan: data.NamaPelatihan,
		WaktuMulai:    data.WaktuMulai,
		WaktuSelesai:  data.WaktuSelesai,
	}
	return res
}
