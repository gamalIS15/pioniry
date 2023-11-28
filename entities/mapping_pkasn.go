package entities

type MappingPkasn struct {
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	KodePkasn1 string `json:"kode_pkasn1" gorm:"index; type:char(20)"`
	KodePkasn2 string `json:"kode_pkasn2" gorm:"index;type:char(20)"`
	KodePkasn3 string `json:"kode_pkasn3" gorm:"type:char(20)"`
	KodePkasn4 string `json:"kode_pkasn4" gorm:"type:char(20)"`
	KodePkasn5 string `json:"kode_pkasn5" gorm:"type:char(20)"`
	IsActive   uint   `json:"is_active"`
}

func (e *MappingPkasn) TableName() string {
	return "mapping_pkasn"
}
