package domain

import (
	"fmt"
	"pioniry/db"
	"pioniry/entities"
)

type PelatihanModel struct {
	db.DatabaseConn
}

func (p *PelatihanModel) Create(pelatihan *entities.Pelatihan) (*entities.Pelatihan, error) {
	err := p.GetDB().Create(&pelatihan).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to create pelatihan")
	}
	return pelatihan, nil
}

func (p *PelatihanModel) Update(id int, data *entities.Pelatihan) (*entities.Pelatihan, error) {
	//var pelatihan entities.Pelatihan
	err := p.GetDB().Select("id").Where("id = ?", id).Find(data).Error
	if err != nil {
		return nil, fmt.Errorf("Id not found")
	}

	err = p.GetDB().Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to update")
	}
	return data, nil
}

func (p *PelatihanModel) Delete(id int) error {
	var pelatihan entities.Pelatihan
	err := p.GetDB().Select("id").Where("id = ?", id).Find(&pelatihan).Error
	if err != nil {
		return fmt.Errorf("Id not found")
	}
	err = p.GetDB().Delete(&pelatihan, id).Error
	if err != nil {
		return fmt.Errorf("Failed to delete")
	}
	return nil
}

func (p *PelatihanModel) GetAllPelatihan() (*[]entities.Pelatihan, error) {
	var pelatihan []entities.Pelatihan
	err := p.GetDB().Find(&pelatihan).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to show all user")
	}
	return &pelatihan, nil
}

func (p *PelatihanModel) GetPelatihanById(id int) (*entities.Pelatihan, error) {
	var pelatihan entities.Pelatihan
	err := p.GetDB().Limit(1).Find(&pelatihan, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to show user")
	}
	return &pelatihan, nil
}
