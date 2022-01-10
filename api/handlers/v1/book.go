package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/abdullohsattorov/API_Gateway/genproto/catalog_service"
	l "github.com/abdullohsattorov/API_Gateway/pkg/logger"
	"github.com/abdullohsattorov/API_Gateway/pkg/utils"
)

// CreateBook ...
// @Summary CreateBook
// @Description This API for creating a new book
// @Tags book
// @Accept  json
// @Produce  json
// @Param category request body models.CreateBook true "bookCreateRequest"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/books/ [post]
func (h *handlerV1) CreateBook(c *gin.Context) {
	var (
		body        pb.Book
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	fmt.Println(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().CreateBook(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create book", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetBook ...
// @Summary GetBook
// @Description This API for getting book detail
// @Tags book
// @Accept  json
// @Produce  json
// @Param id path string true "BookId"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/books/{id} [get]
func (h *handlerV1) GetBook(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().GetBook(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get book", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListBooks ...
// @Summary ListBooks
// @Description This API for getting list of books
// @Tags book
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Param authors query string false "Limit"
// @Param category query string false "Limit"
// @Success 200 {object} models.ListBooks
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/books [get]
func (h *handlerV1) ListBooks(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().ListBook(
		ctx, &pb.ListBookReq{
			Limit:   params.Limit,
			Page:    params.Page,
			Filters: params.Filters,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateBook ...
// @Summary UpdateBook
// @Description This API for updating book
// @Tags book
// @Accept  json
// @Produce  json
// @Param id path string true "BookId"
// @Param User request body models.UpdateBook true "bookUpdateRequest"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/books/{id} [put]
func (h *handlerV1) UpdateBook(c *gin.Context) {
	var (
		body        pb.Book
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.BookId = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().UpdateBook(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update book", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteBook ...
// @Summary DeleteBook
// @Description This API for deleting book
// @Tags book
// @Accept  json
// @Produce  json
// @Param id path string true "BookId"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/books/{id} [delete]
func (h *handlerV1) DeleteBook(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().DeleteBook(ctx, &pb.ByIdReq{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete book", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
