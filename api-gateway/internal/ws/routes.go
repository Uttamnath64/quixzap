package ws

// type WSRoutes struct {
// 	container *storage.Container
// 	rGroup    *gin.RouterGroup
// }

// func New(container *storage.Container, server *gin.Engine) *WSRoutes {
// 	return &WSRoutes{
// 		container: container,
// 		rGroup:    server.Group("/ws"),
// 	}
// }

// func (routes *WSRoutes) Handlers() {
// 	hub := NewHub(routes.container)
// 	middle := middlewares.New(routes.container)
// 	routes.rGroup.GET("/user/:customer_id/:session_id", middle.Middleware(), hub.UserWebSocket)
// 	routes.rGroup.GET("/admin/:admin_id", middle.Middleware(), hub.AdminWebSocket)
// }
