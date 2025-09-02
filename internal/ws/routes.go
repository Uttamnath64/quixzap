package ws

import (
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/Uttamnath64/quixzap/internal/middlewares"
	"github.com/Uttamnath64/quixzap/internal/ws"
	"github.com/gin-gonic/gin"
)

type WSRoutes struct {
	container *storage.Container
	rGroup    *gin.RouterGroup
}

func New(container *storage.Container, server *gin.Engine) *WSRoutes {
	return &WSRoutes{
		container: container,
		rGroup:    server.Group("/ws"),
	}
}

func (routes *WSRoutes) Handlers() {
	hub := ws.NewHub(routes.container)
	middle := middlewares.New(routes.container)

	// WebSocket endpoints
	routes.rGroup.GET("/user/:customer_id/:session_id", middle.SessionAuthMiddleware(), hub.UserWebSocket)
	routes.rGroup.GET("/admin/:admin_id", middle.AdminAuthMiddleware(), hub.AdminWebSocket)
}
