package domain

import (
	"fmt"
	"pioniry/db"
	"pioniry/entities"
)

type PegawaiModel struct {
	db.DatabaseConn
}

func (p *PegawaiModel) GetPegawaiByNip(nip string) (*entities.Pegawai, error) {
	var pegawai entities.Pegawai

	err := p.GetDB().Where("nip = ?", nip).Find(&pegawai).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to get pegawai", err)
	}

	return &pegawai, nil
}
func (p *PegawaiModel) GetPegawaiByParam(param string, limit int, offset int, order string) (interface{}, error) {
	type resp struct {
		Id   int    `json:"id"`
		Nip  string `json:"nip"`
		Nama string `json:"nama"`
	}

	var pegawai entities.Pegawai
	var selectData []resp
	var err error
	if limit != 0 && offset != 0 {
		err = p.GetDB().Model(&pegawai).Where("nip = ?", param).Or("nip_lama =?", param).Or("nama like ?", "%"+param+"%").Or("nama_gelar like ?", "%"+param+"%").Order(order).Limit(limit).Offset(offset).Find(&selectData).Error
	} else if limit != 0 {
		err = p.GetDB().Model(&pegawai).Where("nip = ?", param).Or("nip_lama =?", param).Or("nama like ?", "%"+param+"%").Or("nama_gelar like ?", "%"+param+"%").Order(order).Limit(limit).Find(&selectData).Error
	} else if offset != 0 {
		err = p.GetDB().Model(&pegawai).Where("nip = ?", param).Or("nip_lama =?", param).Or("nama like ?", "%"+param+"%").Or("nama_gelar like ?", "%"+param+"%").Order(order).Offset(offset).Find(&selectData).Error
	} else {
		err = p.GetDB().Model(&pegawai).Where("nip = ?", param).Or("nip_lama =?", param).Or("nama like ?", "%"+param+"%").Or("nama_gelar like ?", "%"+param+"%").Order(order).Find(&selectData).Error
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to get pegawai", err)
	}

	return &selectData, nil
}

func (p *PegawaiModel) GetCountPegawai() (int64, error) {
	var count int64
	var pegawai entities.Pegawai
	err := p.GetDB().Model(&pegawai).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("Failed to get count ", err)
	}
	return count, nil
}
