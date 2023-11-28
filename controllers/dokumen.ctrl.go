package controllers

import (
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pioniry/entities"
	"pioniry/helpers"
	"pioniry/models/domain"
	"strconv"
	"strings"
	"time"
)

type DokumenController struct {
	model domain.DokumenPelatihan
}

func (d DokumenController) Create(c echo.Context) error {
	var allFile []string
	idRelasi := c.FormValue("id_relasi")
	jenis := c.FormValue("jenis")
	namaDokumen := c.FormValue("nama_dokumen")

	form, err := c.MultipartForm()
	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusBadRequest)
	}

	files := form.File["files"]

	for index, file := range files {
		//Source
		src, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer src.Close()

		fileExt := filepath.Ext(file.Filename)
		currentTime := time.Now()
		filename := "Dokumen_" + namaDokumen + "_" + currentTime.Format(time.DateOnly) + "_" + strconv.Itoa(index) + fileExt

		// Destination
		dst, err := os.Create("static/upload/" + filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		allFile = append(allFile, filename)

	}

	fileResp := strings.Join(allFile, " ")
	body := entities.DokumenPelatihan{
		IdRelasi:    idRelasi,
		Jenis:       jenis,
		NamaDokumen: namaDokumen,
		File:        fileResp,
	}
	erd, errd := d.model.Create(&body)
	if errd != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusBadRequest)
	}

	return helpers.HandleSuccessResponse(c, "Document Uploaded", erd, http.StatusCreated)
}

func (d DokumenController) Delete(c echo.Context) error {
	fileloc := "static/upload/"
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	data, err := d.model.GetDokumen(id)

	if err != nil || data.ID == 0 {
		return helpers.HandleErrorResponse(c, "Data not found", http.StatusNotFound)
	}

	//Remove File
	if _, erra := os.Stat(fileloc + data.File); erra == nil {
		err = os.Remove(fileloc + data.File)
		if err != nil {
			return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusBadRequest)
		}
	} else {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusBadRequest)
	}

	//Remove Data in Database
	err = d.model.Delete(id)
	if err != nil {
		return helpers.HandleErrorResponse(c, "Something Wrong", http.StatusBadRequest)
	}

	return helpers.HandleSuccessResponse(c, "Document Deleted", data, http.StatusOK)
}
