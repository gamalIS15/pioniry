package db

import (
	"gorm.io/gorm"
	"log"
	"pioniry/entities"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&entities.User{},
		&entities.Pegawai{},
		&entities.Pelatihan{},
		&entities.DokumenPelatihan{},
		&entities.MasterPKASN{},
		&entities.MappingPkasn{},
		&entities.PesertaInternal{},
	)
	if err != nil {
		log.Fatal("Migration Error", err)
	}
}
