package http

import (
	"github.com/labstack/echo/v4"
	"luthfi/pemilu/domain"
	"luthfi/pemilu/internal/components"
	"net/http"
)

type userHandler struct {
	domain.UserUseCase
}

func NewUserHandler(e *echo.Echo, uu domain.UserUseCase) {
	handler := &userHandler{uu}
	g := e.Group("/api/v1")
	g.POST("/users/register", handler.RegisterUser)
	g.POST("/users/login", handler.LoginUser)
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	var newUser domain.User
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	user, err := h.UserUseCase.RegisterUser(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   user,
	})
}

func (h *userHandler) LoginUser(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	userData, err := h.UserUseCase.LoginUser(user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	token := components.GenerateToken(userData)
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"token":  token,
	})
	
}
