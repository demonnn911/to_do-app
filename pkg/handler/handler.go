package handler

import (
	"todo-app/pkg/service"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}
	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			itemsLinkedLists := lists.Group(":id/items")
			{
				itemsLinkedLists.POST("/", h.createItem)
				itemsLinkedLists.GET("/", h.getAllItems)
			}

		}

		itemsIndependence := api.Group("/items")
		{
			itemsIndependence.GET("/:id", h.getItemById)
			itemsIndependence.PUT("/:id", h.updateItem)
			itemsIndependence.DELETE("/:id", h.deleteItem)
		}

	}
	return router

}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		services: serv,
	}
}
