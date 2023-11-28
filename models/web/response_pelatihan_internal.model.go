package web

import (
	"github.com/google/uuid"
	"time"
)

type ResponsePelatihanInternal struct {
	IdPelatihan   uuid.UUID `json:"id_pelatihan"`
	NamaPelatihan string    `json:"nama_pelatihan"`
	WaktuMulai    time.Time `json:"waktu_mulai"`
	WaktuSelesai  time.Time `json:"waktu_selesai"`
}
