package handler

import (
	"Ozinshe/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	services *service.Service
	db       *sqlx.DB
}

func NewHandler(services *service.Service, db *sqlx.DB) *Handler {
	return &Handler{
		services: services,
		db:       db}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signUp", h.signUp)
		// products.POST("/signIp", signIn)
	}

	return router
}
