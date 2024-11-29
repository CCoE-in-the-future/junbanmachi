package dto

import "time"

type UserDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	NumberPeople int  `json:"numberPeople"`     
	WaitStatus bool   `json:"waitStatus"`
	ArrivalTime  time.Time `json:"arrivalTime"`
}


func NewUserDTO(id string, name string, numberPeople int, waitStatus bool, arrivalTime time.Time) UserDTO {
	return UserDTO{
		ID           :id,
		Name         :name,    
		NumberPeople :numberPeople,       
		WaitStatus   :waitStatus,
		ArrivalTime  :arrivalTime,
	}
}
