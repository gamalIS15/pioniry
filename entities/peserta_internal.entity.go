package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PesertaInternal struct {
	gorm.Model
	IdPelatihan uuid.UUID `json:"-"`
	IdPegawai   int       `json:"-"`
	Kepesertaan string    `json:"kepesertaan"`
	Status      uint8     `json:"status" gorm:"type:tinyint(2)"`
	Pelatihan   Pelatihan `json:"pelatihan,omitempty" gorm:"foreignKey:IdPelatihan;references:KodePelatihan"`
	Pegawai     Pegawai   `json:"pegawai,omitempty" gorm:"foreignKey:IdPegawai;references:Id"`
}

func (e *PesertaInternal) TableName() string {
	return "peserta_internal"
}
