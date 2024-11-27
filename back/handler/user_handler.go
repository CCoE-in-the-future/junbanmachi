package handler

import (
	"net/http"

	"back/entity"
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
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}
	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user",
		})
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var req struct {
		ID string `json:"id"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}
	if err := h.userService.DeleteUser(req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete user",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}

func (h *UserHandler) UpdateUserWaitStatus(c echo.Context) error {
	var req struct {
		ID string `json:"id"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}
	if err := h.userService.UpdateUserWaitStatus(req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update user",
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
		})
	}
	return c.JSON(http.StatusOK, map[string]int{
		"waitTime": waitTime,
	})
}
