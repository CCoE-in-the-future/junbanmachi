package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
	NumberPeople int      `json:"numberPeople"` 
	WaitStatus  bool      `json:"waitStatus"`
    ArrivalTime time.Time `json:"arrivalTime"`
}

var users = []User{
    {ID: 1, Name: "太郎", NumberPeople: 2, WaitStatus: true, ArrivalTime: time.Now()},
    {ID: 2, Name: "花子", NumberPeople: 2,WaitStatus: true, ArrivalTime: time.Now()},
    {ID: 3, Name: "次郎", NumberPeople: 4,WaitStatus: true, ArrivalTime: time.Now()},
    {ID: 4, Name: "三郎", NumberPeople: 1,WaitStatus: false, ArrivalTime: time.Now()},
    {ID: 5, Name: "四郎", NumberPeople: 2,WaitStatus: false, ArrivalTime: time.Now()},
    {ID: 6, Name: "五郎", NumberPeople: 3,WaitStatus: false, ArrivalTime: time.Now()},
    {ID: 7, Name: "六子", NumberPeople: 2,WaitStatus: false, ArrivalTime: time.Now()},
}

func main() {
    e := echo.New()
		e.Use(middleware.CORS())


    e.GET("/api/users", func(c echo.Context) error {
        return c.JSON(http.StatusOK, users)
    })

    e.POST("/api/users", func(c echo.Context) error {
        var newUser User
        if err := c.Bind(&newUser); err != nil {
            return err
        }
        newUser.ID = len(users) + 1
        newUser.ArrivalTime = time.Now()
        users = append(users, newUser)
        return c.JSON(http.StatusCreated, newUser)
    })

    e.DELETE("/api/users", func(c echo.Context) error {
        var request struct {
            ID int `json:"id"`
        }
        if err := c.Bind(&request); err != nil {
            return err
        }

        filteredUsers := []User{}
        for _, user := range users {
            if user.ID != request.ID {
                filteredUsers = append(filteredUsers, user)
            }
        }
        users = filteredUsers
        return c.JSON(http.StatusOK, users)
    })

	e.GET("/api/wait-time", func(c echo.Context) error {
		userLength := len(users)
		estimatedWaitTime := userLength * 20 
		return c.JSON(http.StatusOK, map[string]int{
			"waitTime": estimatedWaitTime,
		})
	})

    e.Logger.Fatal(e.Start(":8080"))
}
