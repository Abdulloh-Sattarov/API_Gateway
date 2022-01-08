package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/abdullohsattorov/API_Gateway/genproto/catalog_service"
	l "github.com/abdullohsattorov/API_Gateway/pkg/logger"
	"github.com/abdullohsattorov/API_Gateway/pkg/utils"
)

// CreateCategory ...
// @Summary CreateCategory
// @Description This API for creating a new category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category request body models.CreateCategory true "categoryCreateRequest"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/categories/ [post]
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		body        pb.Category
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().CreateCategory(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create category", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetCategory ...
// @Summary GetCategory
// @Description This API for getting category detail
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "CategoryId"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/categories/{id} [get]
func (h *handlerV1) GetCategory(c *gin.Context) {
	var jspbMarhsal protojson.MarshalOptions
	jspbMarhsal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().GetCategory(ctx, &pb.ByIdReq{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get category", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListCategories ...
// @Summary ListCategories
// @Description This API for getting list of categories
// @Tags category
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListCategories
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/categories [get]
func (h *handlerV1) ListCategories(c *gin.Context) {
	queryParam := c.Request.URL.Query()
	params, errStr := utils.ParseQueryParams(queryParam)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().ListCategory(
		ctx,
		&pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list category", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateCategory ...
// @Summary UpdateCategory
// @Description This API for updating category
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "CategoryId"
// @Param User request body models.UpdateCategory true "categoryUpdateRequest"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/categories/{id} [put]
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var (
		body        pb.Category
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
	body.CategoryId = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().UpdateCategory(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update category", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteCategory ...
// @Summary DeleteCategory
// @Description This API for deleting category
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "CategoryId"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/categories/{id} [delete]
func (h *handlerV1) DeleteCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().DeleteCategory(ctx, &pb.ByIdReq{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete category", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
