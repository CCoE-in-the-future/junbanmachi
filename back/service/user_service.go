package service

import (
	"back/dto"
	"back/entity"

	"time"

	"github.com/google/uuid"
)


type UserService struct {
	userRepo UserRepositoryInterface
}

func NewUserService(ur UserRepositoryInterface) *UserService {
	return &UserService{userRepo: ur}
}

func GenerateUUID() string {
	return uuid.New().String()
}


func (s *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	users, _ := s.userRepo.GetAllUsers()

	userDTOs := make([]dto.UserDTO, len(users))

	for i, user := range users {
		userDTOs[i] = s.convertToUserDTO(user) 
	}

	return userDTOs, nil
}

func (s *UserService) CreateUser(userDTO dto.UserDTO) (dto.UserDTO, error) {
	userDTO.ID = GenerateUUID()
	userDTO.WaitStatus = true
	userDTO.ArrivalTime = time.Now()

	user := s.convertToUser(userDTO)
	newUser, _ := s.userRepo.CreateUser(user)
	newUserDTO := s.convertToUserDTO(newUser)

	return newUserDTO, nil
}


func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}

func (s *UserService) UpdateUserWaitStatus(id string) error {
	return s.userRepo.UpdateUserWaitStatus(id, false)
}

func (s *UserService) GetEstimatedWaitTime() (int, error) {
	users, _ := s.userRepo.GetWaitingUsers()
	waitTime := 0
	for _, user := range users {
		waitTime += user.NumberPeople * 15
	}
	return waitTime, nil
}

// DTO to entity
func (s *UserService) convertToUser(user dto.UserDTO) entity.User {
	return entity.NewUser(user.ID, user.Name, user.NumberPeople, user.WaitStatus, user.ArrivalTime)
}

// entity to DTO
func (s *UserService) convertToUserDTO(user entity.User) dto.UserDTO {
	return dto.NewUserDTO(user.ID, user.Name, user.NumberPeople, user.WaitStatus, user.ArrivalTime)
}