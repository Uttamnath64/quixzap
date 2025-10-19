package middlewares

import (
	"context"
	"time"

	"github.com/Uttamnath64/quixzap/app/utils/requests"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) NoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		rctx := &requests.RequestContext{
			Ctx: ctx,
		}

		c.Set("rctx", rctx)
		c.Next()
	}
}
