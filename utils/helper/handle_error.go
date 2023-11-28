package helper

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func BadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, map[string]any{
		"Code":   400,
		"Status": http.StatusText(http.StatusBadRequest),
		"Error":  err.Error(),
	})
}

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, map[string]any{
		"Code":   500,
		"Status": http.StatusText(http.StatusInternalServerError),
		"Error":  err.Error(),
	})
}

func NotFound(c echo.Context, err error) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"Code":   404,
		"Status": http.StatusText(http.StatusNotFound),
		"Error":  err.Error(),
	})
}

func ErrParams(c echo.Context, err error) {
	c.JSON(http.StatusBadRequest, nil)
}
