package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"luthfi/pemilu/domain"
	"luthfi/pemilu/internal/components"
	"net/http"
	"strconv"
	"time"
)

type paslonHandler struct {
	domain.PaslonUseCase
}

func NewPaslonHandler(e *echo.Echo, pu domain.PaslonUseCase) {
	paslonHandler := &paslonHandler{pu}
	g := e.Group("api/v1")
	g.GET("/paslon", components.TokenMiddleware(paslonHandler.FetchPaslon))
	g.GET("/paslon/:id", components.TokenMiddleware(paslonHandler.GetByIDPaslon))
	g.POST("/paslon/post", components.TokenMiddleware(paslonHandler.CreatePaslon))
	g.PUT("/paslon/update/:id", components.TokenMiddleware(paslonHandler.UpdatePaslon))
	g.DELETE("/paslon/delete/:id", components.TokenMiddleware(paslonHandler.DeletePaslon))
}

func (h *paslonHandler) FetchPaslon(c echo.Context) error {
	paslons, err := h.PaslonUseCase.Fetch()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	
	return c.JSON(200, map[string]any{
		"code":   200,
		"status": "success",
		"data":   paslons,
	})
}

func (h *paslonHandler) GetByIDPaslon(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	
	paslon, err := h.PaslonUseCase.GetByID(int64(numId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   paslon,
	})
}

func (h *paslonHandler) CreatePaslon(c echo.Context) error {
	var newPaslon domain.Paslon
	cld := components.CloudinaryUploadConfig()
	name := c.FormValue("name")
	serialNumber := c.FormValue("serial_number")
	visionMission := c.FormValue("vision_mission")
	imageFileHeader, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Image is required",
		})
	}
	
	imageFile, err := imageFileHeader.Open()
	
	fileName := fmt.Sprintf("%s-%s", time.Now().Format("20060102150405"), imageFileHeader.Filename)
	
	err = components.UpFileToLocal(imageFileHeader)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	cloudinaryReq := components.CloudinaryUploadReq{
		CTX:       context.Background(),
		ImageFile: imageFile,
		FileName:  fileName,
	}
	
	scrUrl, err := components.CloudinaryUploadImage(cld, cloudinaryReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	newPaslon = domain.Paslon{
		Name:          name,
		SerialNumber:  serialNumber,
		VisionMission: visionMission,
		Image:         scrUrl,
	}
	
	paslon, err := h.PaslonUseCase.Store(newPaslon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   paslon,
	})
}

func (h *paslonHandler) UpdatePaslon(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	
	var updatePaslon domain.Paslon
	_, err = h.PaslonUseCase.Update(int64(numId), updatePaslon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	cld := components.CloudinaryUploadConfig()
	name := c.FormValue("name")
	serialNumber := c.FormValue("serial_number")
	visionMission := c.FormValue("vision_mission")
	imageFileHeader, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Image is required",
		})
	}
	
	imageFile, err := imageFileHeader.Open()
	
	fileName := fmt.Sprintf("%s-%s", time.Now().Format("20060102150405"), imageFileHeader.Filename)
	
	err = components.UpFileToLocal(imageFileHeader)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	cloudinaryReq := components.CloudinaryUploadReq{
		CTX:       context.Background(),
		ImageFile: imageFile,
		FileName:  fileName,
	}
	
	scrUrl, err := components.CloudinaryUploadImage(cld, cloudinaryReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	updatePaslon = domain.Paslon{
		Name:          name,
		SerialNumber:  serialNumber,
		VisionMission: visionMission,
		Image:         scrUrl,
	}
	
	paslon, _ := h.PaslonUseCase.Update(int64(numId), updatePaslon)
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   paslon,
	})
}

func (h *paslonHandler) DeletePaslon(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	
	err = h.PaslonUseCase.Delete(int64(numId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
	})
}
