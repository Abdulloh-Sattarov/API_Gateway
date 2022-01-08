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

// CreateAuthor ...
// @Summary CreateAuthor
// @Description This API for creating a new author
// @Tags author
// @Accept json
// @Produce json
// @Param author request body models.CUAuthor true "authorCreateRequest"
// @Success 200 {object} models.Author
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/authors/ [post]
func (h *handlerV1) CreateAuthor(c *gin.Context) {
	var (
		body        pb.Author
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

	response, err := h.serviceManager.CatalogService().CreateAuthor(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create author", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetAuthor ...
// @Summary GetAuthor
// @Description This API for getting author detail
// @Tags author
// @Accept json
// @Produce json
// @Param id path string true "AuthorId"
// @Success 200 {object} models.Author
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/authors/{id} [get]
func (h *handlerV1) GetAuthor(c *gin.Context) {
	var jspbMarhsal protojson.MarshalOptions
	jspbMarhsal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().GetAuthor(ctx, &pb.ByIdReq{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get author", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListAuthors ...
// @Summary ListAuthors
// @Description This API for getting list of authors
// @Tags author
// @Accept json
// @Produce json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListAuthors
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/authors [get]
func (h *handlerV1) ListAuthors(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to pars query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().ListAuthor(
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
		h.log.Error("failed to list authors", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateAuthor ...
// @Summary UpdateAuthor
// @Description This API for updating author
// @Tags author
// @Accept json
// @Produce json
// @Param id path string true "AuthorId"
// @Param User request body models.CUAuthor true "BookUpdateRequest"
// @Success 200 {object} models.Author
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/authors/{id} [put]
func (h *handlerV1) UpdateAuthor(c *gin.Context) {
	var (
		body        pb.Author
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
	body.AuthorId = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().UpdateAuthor(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update author", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteAuthor ...
// @Summary DeleteAuthor
// @Description This API for deleting the author
// @Tags author
// @Accept json
// @Produce json
// @Param id path string true "AuthorId"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/catalogs/authors/{id} [delete]
func (h *handlerV1) DeleteAuthor(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().DeleteAuthor(ctx, &pb.ByIdReq{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete author", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
