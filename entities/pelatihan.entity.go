package entities

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Pelatihan struct {
	gorm.Model
	KodePelatihan           uuid.UUID      `json:"kode_pelatihan" gorm:"type:char(36);index"`
	KodePKASN               int            `json:"kode_pkasn" gorm:"index"`
	JenisPelatihan          string         `json:"jenis_pelatihan"`
	NamaPelatihan           string         `json:"nama_pelatihan"`
	WaktuMulai              time.Time      `json:"waktu_mulai"`
	WaktuSelesai            time.Time      `json:"waktu_selesai"`
	JenisPengajuanPelatihan string         `json:"jenis_pengajuan_pelatihan" gorm:"index"`
	MetodePelaksanaan       string         `json:"metode_pelaksanaan"`
	Tahun                   string         `json:"tahun"`
	Durasi                  int32          `json:"durasi"`
	Status                  uint8          `json:"status"`
	Penyelenggara           int            `json:"penyelenggara"`
	CreatedBy               int            `json:"created_by" gorm:"index"`
	Configs                 datatypes.JSON `json:"configs"`
}

func (p *Pelatihan) BeforeCreate(tx *gorm.DB) error {
	if p.KodePelatihan == uuid.Nil {
		p.KodePelatihan = uuid.New()
	}
	return nil
}

func (p *Pelatihan) TableName() string {
	return "pelatihan"
}
