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

type partaiHandler struct {
	domain.PartaiUseCase
}

func NewPartaiHandler(e *echo.Echo, pu domain.PartaiUseCase) {
	handler := &partaiHandler{pu}
	
	g := e.Group("/api/v1")
	g.GET("/partai", components.TokenMiddleware(handler.FetchPartai))
	g.GET("/partai/:id", components.TokenMiddleware(handler.GetByIDPartai))
	g.POST("/partai/post", components.TokenMiddleware(handler.CreatePartai))
	g.PUT("/partai/update/:id", components.TokenMiddleware(handler.UpdatePartai))
	g.DELETE("/partai/delete/:id", components.TokenMiddleware(handler.DeletePartai))
}

func (h *partaiHandler) FetchPartai(c echo.Context) error {
	partais, err := h.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   partais,
	})
}

func (h *partaiHandler) GetByIDPartai(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	
	partai, err := h.PartaiUseCase.GetByID(int64(numId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   partai,
	})
}

func (h *partaiHandler) CreatePartai(c echo.Context) error {
	var newPartai domain.Partai
	cld := components.CloudinaryUploadConfig()
	name := c.FormValue("name")
	chairman := c.FormValue("chairman")
	visionMission := c.FormValue("vision_mission")
	address := c.FormValue("address")
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	cloudinaryReq := components.CloudinaryUploadReq{
		CTX:       context.Background(),
		ImageFile: imageFile,
		FileName:  fileName,
	}
	
	scrURL, err := components.CloudinaryUploadImage(cld, cloudinaryReq)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	newPartai = domain.Partai{
		Name:          name,
		Chairman:      chairman,
		VisionMission: visionMission,
		Address:       address,
		Image:         scrURL,
	}
	
	partai, err := h.PartaiUseCase.Store(newPartai)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   partai,
	})
}

func (h *partaiHandler) UpdatePartai(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	
	updatePartai := domain.Partai{}
	
	_, err = h.PartaiUseCase.Update(int64(numId), updatePartai)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	
	cld := components.CloudinaryUploadConfig()
	name := c.FormValue("name")
	chairman := c.FormValue("chairman")
	visionMission := c.FormValue("vision_mission")
	address := c.FormValue("address")
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	cloudinaryReq := components.CloudinaryUploadReq{
		CTX:       context.Background(),
		ImageFile: imageFile,
		FileName:  fileName,
	}
	
	scrURL, err := components.CloudinaryUploadImage(cld, cloudinaryReq)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	updatePartai = domain.Partai{
		Name:          name,
		Chairman:      chairman,
		VisionMission: visionMission,
		Address:       address,
		Image:         scrURL,
	}
	
	partai, _ := h.PartaiUseCase.Update(int64(numId), updatePartai)
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   partai,
	})
	
}

func (h *partaiHandler) DeletePartai(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	
	err = h.PartaiUseCase.Delete(int64(numId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   "Success Delete Partai",
	})
}
