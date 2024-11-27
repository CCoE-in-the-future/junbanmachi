package service

import "back/entity"

type UserRepositoryInterface interface {
	GetAllUsers() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	DeleteUser(id string) error
	UpdateUserWaitStatus(id string, status bool) error
	GetWaitingUsers() ([]entity.User, error)
}
