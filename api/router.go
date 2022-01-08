package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/abdullohsattorov/API_Gateway/api/docs" // swag
	v1 "github.com/abdullohsattorov/API_Gateway/api/handlers/v1"
	"github.com/abdullohsattorov/API_Gateway/config"
	"github.com/abdullohsattorov/API_Gateway/pkg/logger"
	"github.com/abdullohsattorov/API_Gateway/services"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/catalogs/books", handlerV1.CreateBook)
	api.GET("/catalogs/books/:id", handlerV1.GetBook)
	api.GET("/catalogs/books", handlerV1.ListBooks)
	api.PUT("/catalogs/books/:id", handlerV1.UpdateBook)
	api.DELETE("/catalogs/books/:id", handlerV1.DeleteBook)

	api.POST("/catalogs/categories", handlerV1.CreateCategory)
	api.GET("/catalogs/categories/:id", handlerV1.GetCategory)
	api.GET("/catalogs/categories", handlerV1.ListCategories)
	api.PUT("/catalogs/categories/:id", handlerV1.UpdateCategory)
	api.DELETE("/catalogs/categories/:id", handlerV1.DeleteCategory)

	api.POST("/catalogs/authors", handlerV1.CreateAuthor)
	api.GET("/catalogs/authors/:id", handlerV1.GetAuthor)
	api.GET("/catalogs/authors", handlerV1.ListAuthors)
	api.PUT("/catalogs/authors/:id", handlerV1.UpdateAuthor)
	api.DELETE("/catalogs/authors/:id", handlerV1.DeleteAuthor)

	api.GET("/catalogs", handlerV1.List)

	api.POST("/orders", handlerV1.CreateOrder)
	api.GET("/orders/:id", handlerV1.GetOrder)
	api.GET("/orders", handlerV1.ListOrders)
	api.PUT("/orders/:id", handlerV1.UpdateOrder)
	api.DELETE("/orders/:id", handlerV1.DeleteOrder)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
