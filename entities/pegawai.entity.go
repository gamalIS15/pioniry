package entities

import "time"

type Pegawai struct {
	Id               int       `json:"id" gorm:"index"`
	NipLama          string    `json:"nip_lama" gorm:"unique"`
	Nip              string    `json:"nip" gorm:"unique"`
	Nama             string    `json:"nama"`
	NamaGelar        string    `json:"nama_gelar"`
	Jabatan          string    `json:"jabatan"`
	JenjangJabatan   string    `json:"jenjang_jabatan"`
	UnitKerja        string    `json:"unit_kerja"`
	SatuanOrganisasi string    `json:"satuan_organisasi"`
	UpdateAt         time.Time `json:"update_at"`
}

func (e *Pegawai) TableName() string {
	return "pegawai"
}
