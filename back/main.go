package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"back/handler"
	"back/repository"
	"back/service"
)

func main() {
	// AWS Session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	db := dynamodb.New(sess)
	
	var userRepo service.UserRepositoryInterface = repository.NewUserRepository(db, "junbanmachi-table")
	var userService service.UserServiceInterface = service.NewUserService(userRepo) 
	var userHandler handler.UserHandlerInterface = handler.NewUserHandler(userService) 

	// Echo instance
	e := echo.New()
	e.Use(middleware.CORS())

	// Routes
	e.GET("/api/users", userHandler.GetAllUsers)
	e.POST("/api/users", userHandler.CreateUser)
	e.DELETE("/api/users", userHandler.DeleteUser)
	e.PUT("/api/users", userHandler.UpdateUserWaitStatus)
	e.GET("/api/wait-time", userHandler.GetEstimatedWaitTime)

	// Start server
	log.Fatal(e.Start(":8080"))
}
