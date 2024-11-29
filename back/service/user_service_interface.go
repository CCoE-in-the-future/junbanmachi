package service

import "back/dto"

type UserServiceInterface interface {
	GetAllUsers() ([]dto.UserDTO, error)
	CreateUser(user dto.UserDTO) (dto.UserDTO, error)
	DeleteUser(id string) error
	UpdateUserWaitStatus(id string) error
	GetEstimatedWaitTime() (int, error)
}
