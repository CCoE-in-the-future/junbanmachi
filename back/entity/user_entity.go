package entity

import "time"

type User struct {
	ID           string    
	Name         string    
	NumberPeople int       
	WaitStatus   bool     
	ArrivalTime  time.Time
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