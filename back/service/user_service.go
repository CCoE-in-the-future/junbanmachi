package service

import (
	"time"

	"back/entity"
	"back/repository"
)


type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepo: ur}
}

func (s *UserService) GetAllUsers() ([]entity.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserService) CreateUser(user entity.User) (entity.User, error) {
	user.ID = repository.GenerateUUID()
	user.WaitStatus = true
	user.ArrivalTime = time.Now()
	return s.userRepo.CreateUser(user)
}


func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}

func (s *UserService) UpdateUserWaitStatus(id string) error {
	return s.userRepo.UpdateUserWaitStatus(id, false)
}

func (s *UserService) GetEstimatedWaitTime() (int, error) {
	users, err := s.userRepo.GetWaitingUsers()
	if err != nil {
		return 0, err
	}
	waitTime := 0
	for _, user := range users {
		waitTime += user.NumberPeople * 15
	}
	return waitTime, nil
}
