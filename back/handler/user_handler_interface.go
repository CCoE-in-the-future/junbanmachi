package handler

import (
	"github.com/labstack/echo/v4"
)

type UserHandlerInterface interface {
	GetAllUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	UpdateUserWaitStatus(c echo.Context) error
	GetEstimatedWaitTime(c echo.Context) error
}
