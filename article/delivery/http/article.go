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

type articleHandler struct {
	domain.ArticleUseCase
}

func NewArticleHandler(e *echo.Echo, au domain.ArticleUseCase) {
	handler := &articleHandler{au}
	g := e.Group("/api/v1")
	g.GET("/articles", components.TokenMiddleware(handler.FetchArticle))
	g.GET("/articles/:id", components.TokenMiddleware(handler.GetByIDArticle))
	g.POST("/articles/post", components.TokenMiddleware(handler.CreateArticle))
	g.PUT("/articles/update/:id", components.TokenMiddleware(handler.UpdateArticle))
	g.DELETE("/articles/delete/:id", components.TokenMiddleware(handler.DeleteArticle))
}

func (h *articleHandler) FetchArticle(c echo.Context) error {
	articles, err := h.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   articles,
	})
}

func (h *articleHandler) GetByIDArticle(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	article, err := h.ArticleUseCase.GetByID(int64(numId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   article,
	})
}

func (h *articleHandler) CreateArticle(c echo.Context) error {
	var newArticle domain.Article
	cld := components.CloudinaryUploadConfig()
	title := c.FormValue("title")
	author := c.FormValue("author")
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
	
	description := c.FormValue("description")
	
	newArticle = domain.Article{
		Title:       title,
		Author:      author,
		Image:       scrURL,
		Description: description,
	}
	
	article, err := h.ArticleUseCase.Store(newArticle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   article,
	})
}

func (h *articleHandler) UpdateArticle(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	
	var updateArticle domain.Article
	_, err = h.ArticleUseCase.Update(int64(numId), updateArticle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	cld := components.CloudinaryUploadConfig()
	title := c.FormValue("title")
	author := c.FormValue("author")
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
	
	description := c.FormValue("description")
	
	updateArticle = domain.Article{
		Title:       title,
		Author:      author,
		Image:       scrURL,
		Description: description,
	}
	
	article, _ := h.ArticleUseCase.Update(int64(numId), updateArticle)
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":   200,
		"Status": http.StatusText(http.StatusOK),
		"data":   article,
	})
}

func (h *articleHandler) DeleteArticle(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err = h.ArticleUseCase.Delete(int64(numId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"Code":    200,
		"Status":  http.StatusText(http.StatusOK),
		"message": "success delete",
	})
}
