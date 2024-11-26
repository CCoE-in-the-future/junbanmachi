package entity

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	NumberPeople int       `json:"numberPeople"`
	WaitStatus   bool      `json:"waitStatus"`
	ArrivalTime  time.Time `json:"arrivalTime"`
}
