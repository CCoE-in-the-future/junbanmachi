package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	NumberPeople int       `json:"numberPeople"`
	WaitStatus   bool      `json:"waitStatus"`
	ArrivalTime  time.Time `json:"arrivalTime"`
}

var tableName = "junbanmachi-table"

var users = []User{
    {ID: "578e084d-a147-4797-8fca-6a9b383320e6", Name: "太郎", NumberPeople: 2, WaitStatus: false, ArrivalTime: time.Now().Add(-60 * time.Minute)},
    {ID: "d9c62f2d-73cf-4e3e-960f-392ba6fd59fa", Name: "花子", NumberPeople: 2, WaitStatus: false, ArrivalTime: time.Now().Add(-50 * time.Minute)},
    {ID: "850b0afa-5066-47b7-a460-a5685090580b", Name: "次郎", NumberPeople: 4, WaitStatus: false, ArrivalTime: time.Now().Add(-40 * time.Minute)},
    {ID: "5daf7be7-8ac1-443e-b522-e441e737ac6d", Name: "三郎", NumberPeople: 1, WaitStatus: true, ArrivalTime: time.Now().Add(-30 * time.Minute)},
    {ID: "2717da2c-7ba6-4dfe-8067-50cbac9c7792", Name: "四郎", NumberPeople: 2, WaitStatus: true, ArrivalTime: time.Now().Add(-20 * time.Minute)},
    {ID: "69e05da4-d730-4983-a71a-d95ae60bcd31", Name: "五郎", NumberPeople: 3, WaitStatus: true, ArrivalTime: time.Now().Add(-10 * time.Minute)},
    {ID: "77e6f4b2-39aa-47ca-b669-6c1ba68c0869", Name: "六子", NumberPeople: 2, WaitStatus: true, ArrivalTime: time.Now().Add(-0 * time.Minute)},
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))

	svc := dynamodb.New(sess)

	for _, user := range users {
		av, err := dynamodbattribute.MarshalMap(user)
		if err != nil {
			log.Fatalf("Marshalling failed: %v", err)
		}

		input := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      av,
		}

		_, err = svc.PutItem(input)
		if err != nil {
			log.Fatalf("Failed to put item: %v", err)
		}

		fmt.Printf("Inserted user: %s\n", user.Name)
	}

	fmt.Println("All users inserted successfully!")

	// Echo instance
	e := echo.New()
	e.Use(middleware.CORS())

	// Get all users from DynamoDB
	e.GET("/api/users", func(c echo.Context) error {
		input := &dynamodb.ScanInput{
			TableName: aws.String(tableName),
		}

		result, err := svc.Scan(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch users from DynamoDB",
			})
		}

		var users []User
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to unmarshal users",
			})
		}

		return c.JSON(http.StatusOK, users)
	})

	// Create a new user and insert into DynamoDB
	e.POST("/api/users", func(c echo.Context) error {
		var newUser User
		if err := c.Bind(&newUser); err != nil {
			return err
		}
		newUser.ID = uuid.New().String()
		newUser.WaitStatus = true
		newUser.ArrivalTime = time.Now()

		av, err := dynamodbattribute.MarshalMap(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to marshal new user",
			})
		}

		input := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      av,
		}

		_, err = svc.PutItem(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to insert user into DynamoDB",
			})
		}

		return c.JSON(http.StatusCreated, newUser)
	})

	// Delete a user from DynamoDB
	e.DELETE("/api/users", func(c echo.Context) error {
		var request struct {
			ID string `json:"id"`
		}
		if err := c.Bind(&request); err != nil {
			return err
		}

		input := &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(request.ID),
				},
			},
		}

		_, err := svc.DeleteItem(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to delete user from DynamoDB",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "User deleted successfully",
		})
	})

	// Update a user's WaitStatus in DynamoDB
	e.PUT("/api/users", func(c echo.Context) error {
		var request struct {
			ID string `json:"id"`
		}
		if err := c.Bind(&request); err != nil {
			return err
		}

		input := &dynamodb.UpdateItemInput{
			TableName: aws.String(tableName),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(request.ID),
				},
			},
			UpdateExpression: aws.String("SET WaitStatus = :ws"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":ws": {
					BOOL: aws.Bool(false),
				},
			},
			ReturnValues: aws.String("UPDATED_NEW"),
		}

		_, err := svc.UpdateItem(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update user in DynamoDB",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "User updated successfully",
		})
	})

	// Get estimated wait time
	e.GET("/api/wait-time", func(c echo.Context) error {
		input := &dynamodb.ScanInput{
			TableName:        aws.String(tableName),
			FilterExpression: aws.String("WaitStatus = :ws"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":ws": {
					BOOL: aws.Bool(true),
				},
			},
		}

		result, err := svc.Scan(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to scan DynamoDB for wait status",
			})
		}

		var waitUsers []User
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &waitUsers)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to unmarshal wait users",
			})
		}

		estimatedWaitTime := 0
		for _, user := range waitUsers {
			estimatedWaitTime += user.NumberPeople * 15
		}

		return c.JSON(http.StatusOK, map[string]int{
			"waitTime": estimatedWaitTime,
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
