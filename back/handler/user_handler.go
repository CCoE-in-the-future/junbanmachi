package handler

import (
	"net/http"

	"back/dto"
	"back/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(us service.UserServiceInterface) *UserHandler {
	return &UserHandler{userService: us}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch users",
			"detail": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user dto.UserDTO
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
			"detail": err.Error(),
		})
	}
	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user",
			"detail": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var user dto.UserDTO
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
			"detail": err.Error(),
		})
	}
	if err := h.userService.DeleteUser(user.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete user",
			"detail": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}

func (h *UserHandler) UpdateUserWaitStatus(c echo.Context) error {
	var user dto.UserDTO
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
			"detail": err.Error(),
		})
	}
	if err := h.userService.UpdateUserWaitStatus(user.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update user",
			"detail": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})
}

func (h *UserHandler) GetEstimatedWaitTime(c echo.Context) error {
	waitTime, err := h.userService.GetEstimatedWaitTime()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to calculate wait time",
			"detail": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]int{
		"waitTime": waitTime,
	})
}
