package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
	NumberPeople int      `json:"numberPeople"` 
	WaitStatus  bool      `json:"waitStatus"`
    ArrivalTime time.Time `json:"arrivalTime"`
}

var users = []User{
    {ID: "578e084d-a147-4797-8fca-6a9b383320e6", Name: "太郎", NumberPeople: 2, WaitStatus: false, ArrivalTime: time.Now()},
    {ID: "d9c62f2d-73cf-4e3e-960f-392ba6fd59fa", Name: "花子", NumberPeople: 2, WaitStatus: false, ArrivalTime: time.Now()},
    {ID: "850b0afa-5066-47b7-a460-a5685090580b", Name: "次郎", NumberPeople: 4, WaitStatus: false, ArrivalTime: time.Now()},
    {ID: "5daf7be7-8ac1-443e-b522-e441e737ac6d", Name: "三郎", NumberPeople: 1, WaitStatus: true, ArrivalTime: time.Now()},
    {ID: "2717da2c-7ba6-4dfe-8067-50cbac9c7792", Name: "四郎", NumberPeople: 2, WaitStatus: true, ArrivalTime: time.Now()},
    {ID: "69e05da4-d730-4983-a71a-d95ae60bcd31", Name: "五郎", NumberPeople: 3, WaitStatus: true, ArrivalTime: time.Now()},
    {ID: "77e6f4b2-39aa-47ca-b669-6c1ba68c0869", Name: "六子", NumberPeople: 2, WaitStatus: true, ArrivalTime: time.Now()},
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
        newUser.ID = uuid.New().String()
        newUser.WaitStatus = true
        newUser.ArrivalTime = time.Now()
        users = append(users, newUser)
        return c.JSON(http.StatusCreated, newUser)
    })

    e.DELETE("/api/users", func(c echo.Context) error {
        var request struct {
            ID string `json:"id"`
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

    e.PUT("/api/users", func(c echo.Context) error {
        var request struct {
            ID string `json:"id"`
        }
        if err := c.Bind(&request); err != nil {
            return err
        }

        updatedUsers := users
        for idx, user := range users {
            if user.ID == request.ID {
            updatedUsers[idx].WaitStatus = false
        }          
        }
        users = updatedUsers
        return c.JSON(http.StatusOK, users)
    })
    

	e.GET("/api/wait-time", func(c echo.Context) error {
		
        waitUsers := []User{}
        for _, user := range users {
            if user.WaitStatus {
                waitUsers = append(waitUsers, user)        
            }          
        }
        estimatedWaitTime :=0
        for _, user := range waitUsers {
            estimatedWaitTime += user.NumberPeople * 15
        }

		return c.JSON(http.StatusOK, map[string]int{
			"waitTime": estimatedWaitTime,
		})
	})

    e.Logger.Fatal(e.Start(":8080"))
}
