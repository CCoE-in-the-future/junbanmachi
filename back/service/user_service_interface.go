package service

import "back/entity"

type UserServiceInterface interface {
	GetAllUsers() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	DeleteUser(id string) error
	UpdateUserWaitStatus(id string) error
	GetEstimatedWaitTime() (int, error)
}
