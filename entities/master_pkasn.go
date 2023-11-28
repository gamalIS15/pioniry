package entities

import "time"

type MasterPKASN struct {
	Id             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	KodePkasn      string    `json:"kode_pkasn" gorm:"index; type:char(20)"`
	NamaPkasn      string    `json:"nama_pkasn" gorm:"type:varchar(255)"`
	KonversiSimsdm string    `json:"konversi_simsdm" gorm:"type:char(20)"`
	Level          int       `json:"level" gorm:"index"`
	CreatedAt      time.Time `json:"created_at"`
}

func (m *MasterPKASN) TableName() string {
	return "master_pkasn"
}
