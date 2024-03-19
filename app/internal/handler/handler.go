package handler

import (
	"app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/create", h.userCreate)
			user.GET("/info/:id", h.userInfo)
			user.POST("/quest", h.userQuest)
		}

		quest := api.Group("/quest")
		{
			quest.POST("/create", h.questCreate)
		}
	}

	return router
}
