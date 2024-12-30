package repository

import (
	"wkm/entity"
	"wkm/request"
	"wkm/response"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uint32) entity.User
	FindByUsername(username string) entity.User
	ResetPassword(data request.ResetPassword)
	All() []response.ClientUser
}

type userRepository struct {
	connUser *gorm.DB
}

func NewUserRepository(connUser *gorm.DB) UserRepository {
	return &userRepository{
		connUser: connUser,
	}
}

func (lR *userRepository) FindById(id uint32) entity.User {
	user := entity.User{ID: id}
	lR.connUser.Preload("Role").Find(&user)

	var permissions []entity.Permission
	lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	for _, v := range permissions {
		user.Permissions = append(user.Permissions, v.Name)
	}

	return user
}

func (lR *userRepository) ResetPassword(data request.ResetPassword) {
	user := entity.User{ID: data.IdUser}
	lR.connUser.Find(&user)
	user.Password = data.Password
	lR.connUser.Save(&user)
}

func (lR *userRepository) FindByUsername(username string) entity.User {
	var user entity.User
	lR.connUser.Where("username", username).First(&user)

	var permissions []entity.Permission
	lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	for _, v := range permissions {
		user.Permissions = append(user.Permissions, v.Name)
	}

	return user
}

func (lR *userRepository) All() []response.ClientUser {
	var data []response.ClientUser
	var user []entity.User
	lR.connUser.Where("role_id", 1).Find(&user)

	for _, v := range user {
		data = append(data, response.ClientUser{
            Name : v.Name,
        })
	}

	return data
}
