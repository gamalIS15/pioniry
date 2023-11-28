package domain

import (
	"fmt"
	"pioniry/db"
	"pioniry/entities"
)

type DokumenPelatihan struct {
	db.DatabaseConn
}

func (d *DokumenPelatihan) Create(dokumen *entities.DokumenPelatihan) (*entities.DokumenPelatihan, error) {
	check := d.GetDB().Limit(1).Find(&dokumen, "id", dokumen.ID)
	if check.RowsAffected < 1 {
		err := d.GetDB().Create(&dokumen).Error
		if err != nil {
			return nil, fmt.Errorf("Failed create data dokumen")
		}
	} else {
		return nil, fmt.Errorf("Duplicate dokumen")
	}
	return dokumen, nil
}

func (d *DokumenPelatihan) Update(id int, data *entities.DokumenPelatihan) (*entities.DokumenPelatihan, error) {
	err := d.GetDB().Select("id").Where("id =?", id).Find(data).Error
	if err != nil {
		return nil, fmt.Errorf("Id not found")
	}

	err = d.GetDB().Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to update")
	}

	return nil, nil
}

func (d *DokumenPelatihan) Delete(id int) error {
	var dokumen entities.DokumenPelatihan
	err := d.GetDB().Select("id").Where("id = ?", id).Find(&dokumen).Error
	if err != nil {
		return fmt.Errorf("Id not found")
	}
	err = d.GetDB().Unscoped().Delete(&dokumen, id).Error
	if err != nil {
		return fmt.Errorf("Failed to delete")
	}
	return nil
}

func (d *DokumenPelatihan) GetDokumen(id int) (*entities.DokumenPelatihan, error) {
	var dokumen entities.DokumenPelatihan
	err := d.GetDB().Where("id = ?", id).Find(&dokumen).Error

	if err != nil {
		return nil, fmt.Errorf("Failed to get pegawai", err)
	}

	return &dokumen, nil
}
