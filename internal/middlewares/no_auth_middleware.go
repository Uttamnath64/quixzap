package middlewares

import (
	"context"
	"time"

	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) NoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// ⏳ Create a context with timeout (e.g., 5 seconds)
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		rctx := &requests.RequestContext{
			Ctx: ctx,
		}

		c.Set("rctx", rctx)
		c.Next()
	}
}
