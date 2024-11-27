package repository

import (
	"back/entity"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UserRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func NewUserRepository(db *dynamodb.DynamoDB, tableName string) *UserRepository {
	return &UserRepository{db: db, tableName: tableName}
}

func (r *UserRepository) GetAllUsers() ([]entity.User, error) {
	input := &dynamodb.ScanInput{
		TableName: &r.tableName,
	}
	result, err := r.db.Scan(input)
	if err != nil {
		return nil, err
	}
	var users []entity.User
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	return users, err
}

func (r *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return entity.User{}, err
	}
	input := &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	}
	_, err = r.db.PutItem(input)
	return user, err
}

func (r *UserRepository) DeleteUser(id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &r.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: &id},
		},
	}
	_, err := r.db.DeleteItem(input)
	return err
}

func (r *UserRepository) UpdateUserWaitStatus(id string, status bool) error {
	input := &dynamodb.UpdateItemInput{
		TableName: &r.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: &id},
		},
		UpdateExpression: aws.String("SET waitStatus = :ws"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ws": {BOOL: &status},
		},
	}
	_, err := r.db.UpdateItem(input)
	return err
}

func (r *UserRepository) GetWaitingUsers() ([]entity.User, error) {
	input := &dynamodb.ScanInput{
		TableName:        &r.tableName,
		FilterExpression: aws.String("waitStatus = :ws"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ws": {BOOL: aws.Bool(true)},
		},
	}

	result, err := r.db.Scan(input)
	if err != nil {
		return nil, err
	}

	var waitingUsers []entity.User
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &waitingUsers)
	if err != nil {
		return nil, err
	}

	return waitingUsers, nil
}
