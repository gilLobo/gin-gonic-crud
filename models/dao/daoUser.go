package dao

import (
	"gin-gonic-crud/config"
	"gin-gonic-crud/models/model"
)

// CreateUser ...
func CreateUser(user *model.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetAllUser ...
func GetAllUser(user *[]model.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// GetByIDUser ...
func GetByIDUser(user *model.User, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ...
func UpdateUser(user *model.User, id string) (err error) {
	config.DB.Save(user)
	return nil
}

// DeleteUser ...
func DeleteUser(user *model.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
