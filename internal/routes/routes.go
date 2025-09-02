package routes

import (
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	v1 "github.com/Uttamnath64/quixzap/internal/routes/v1"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	container *storage.Container
	server    *gin.Engine
	v1        *v1.RoutesV1
	ws        *ws.WSRoutes
}

func New(container *storage.Container, server *gin.Engine) *Routes {
	return &Routes{
		container: container,
		server:    server,
		v1:        v1.New(container, server),
		ws:        ws.New(container, server),
	}
}

func (routes *Routes) Handlers() {
	// Serve static files for the main site
	routes.server.Static("/public/main", "./internal/public/main")

	// Set up API routes under /api/v1
	routes.v1.Handlers()

	// Set up WebSocket routes under /ws
	routes.ws.Handlers()
}
