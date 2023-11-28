package domain

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"pioniry/db"
	"pioniry/entities"
)

type UserModel struct {
	db.DatabaseConn
}

func (u *UserModel) Create(users *entities.User) (*entities.User, error) {
	check := u.GetDB().Limit(1).Find(&users, "nip = ?", users.Nip)
	if check.RowsAffected < 1 {
		err := u.GetDB().Create(&users).Error
		if err != nil {
			return nil, fmt.Errorf("Failed create data users")
		}
	} else {
		return nil, fmt.Errorf("Duplicate user")
	}

	return users, nil
}

func (u *UserModel) Update(id int, data *entities.User) (*entities.User, error) {
	var user entities.User
	err := u.GetDB().Select("id").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Id not found")
	}

	err = u.GetDB().Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to update")
	}

	return nil, nil
}

func (u *UserModel) Delete(id int) error {
	var user entities.User
	err := u.GetDB().Select("id").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return fmt.Errorf("Id not found")
	}
	err = u.GetDB().Delete(&user, id).Error
	if err != nil {
		return fmt.Errorf("Failed to delete")
	}
	return nil

}

func (u *UserModel) GetUsers() (*[]entities.User, error) {
	var users []entities.User
	err := u.GetDB().Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to show all user")
	}
	return &users, nil
}

func (u *UserModel) GetUserById(id int) (*entities.User, error) {
	var user entities.User
	err := u.GetDB().Where("id = ?", id).Find(&user).Error

	if err != nil {
		return nil, fmt.Errorf("Failed to show user")
	}
	return &user, nil
}

func (u *UserModel) GetUserByNip(nip string) (*entities.User, error) {
	var user entities.User
	err := u.GetDB().Limit(1).Find(&user, "nip = ?", nip).Error

	if err != nil {
		return nil, fmt.Errorf("Failed to show user")
	}
	return &user, nil
}

func (u *UserModel) CheckPassword(password string, reqPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(reqPassword))
	if err != nil {
		return fmt.Errorf("password not matched")
	}
	return nil
}
