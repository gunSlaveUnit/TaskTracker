package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gunSlaveUnit/TaskTracker/pkg/service"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/0.1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.SignUp)
			auth.POST("/sign-in", h.SignIn)
			auth.POST("/sign-out", h.SignOut)
		}

		tasks := api.Group("/tasks")
		{
			tasks.GET("/", h.All)
			tasks.GET("/:id", h.GetBy)
			tasks.POST("/", h.Create)
			tasks.PUT("/:id", h.UpdateBy)
			tasks.DELETE("/:id", h.DeleteBy)
		}
	}

	return router
}
