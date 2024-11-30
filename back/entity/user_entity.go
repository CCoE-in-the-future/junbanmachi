package entity

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	NumberPeople int       `json:"numberPeople"`
	WaitStatus   bool      `json:"waitStatus"`
	ArrivalTime  time.Time `json:"arrivalTime"`
}


func NewUser(id string, name string, numberPeople int, waitStatus bool, arrivalTime time.Time) User {
	return User{
		ID           :id,
		Name         :name,    
		NumberPeople :numberPeople,       
		WaitStatus   :waitStatus,
		ArrivalTime  :arrivalTime,
	}
}