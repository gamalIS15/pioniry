package entities

import (
	"fmt"
	"gorm.io/gorm"
	"path"
	"time"
)

type DokumenPelatihan struct {
	gorm.Model
	IdRelasi    string `json:"id_relasi" gorm:"index"`
	Jenis       string `json:"jenis"`
	NamaDokumen string `json:"nama_dokumen"`
	File        string `json:"file"`
	//Configs     datatypes.JSON `json:"configs"`
}

func (d *DokumenPelatihan) FormatFileName() error {
	ext := path.Ext(d.File)
	d.File = fmt.Sprint("Dokumen_", d.IdRelasi, "_", d.Jenis, "_", time.Now().Format("2006-02-01"), ext)
	return nil
}

func (d *DokumenPelatihan) TableName() string {
	return "dokumen_pelatihan"
}
