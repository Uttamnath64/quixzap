package v1

import (
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/gin-gonic/gin"
)

type RoutesV1 struct {
	container *storage.Container
	rGroup    *gin.RouterGroup // Router group for /api/v1
}

func New(container *storage.Container, server *gin.Engine) *RoutesV1 {
	return &RoutesV1{
		container: container,
		rGroup:    server.Group("/api/v1"),
	}
}

func (routes *RoutesV1) Handlers() {
	// Register all v1 route groups
	routes.MainRoutes()
	// routes.CustomerRoutes()
	// routes.AdminRoutes()
	// routes.WidgetRoutes()
	// routes.SuperAdminRoutes()
}
