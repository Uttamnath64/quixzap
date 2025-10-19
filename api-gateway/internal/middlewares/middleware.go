package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/Uttamnath64/quixzap/app/common"
	"github.com/Uttamnath64/quixzap/app/common/types"
	"github.com/Uttamnath64/quixzap/app/utils/requests"
	"github.com/Uttamnath64/quixzap/app/utils/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Middleware struct {
	appCtx *appcontext.AppContext
}

func New(appCtx *appcontext.AppContext) *Middleware {
	return &Middleware{
		appCtx: appCtx,
	}
}

func (m *Middleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Missing access token.",
			})
			c.Abort()
			return
		}

		// remove bearer
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return m.appCtx.Env.Auth.AccessPublicKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid or expired access token.",
			})
			c.Abort()
			return
		}

		// Check if token claims exist and have the expected format
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			m.appCtx.Logger.Error("middleware-claims-format", "token", token.Raw)
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid token claims format.",
			})
			c.Abort()
			return
		}

		// ✅ Check token expiration manually
		exp, ok := claims["exp"].(float64)
		if !ok || int64(exp) < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Access token has expired.",
			})
			c.Abort()
			return
		}

		sessionIDFloat, ok := claims[string(common.CtxSessionID)].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid session ID in token.",
			})
			c.Abort()
			return
		}

		// check Session
		_, err = m.appCtx.Redis.GetValue(fmt.Sprintf("user_sessions:%d", uint(sessionIDFloat)))
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Session expired or logged out",
			})
			c.Abort()
			return
		}

		userIDFloat, ok := claims[string(common.CtxUserID)].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user ID in token.",
			})
			c.Abort()
			return
		}

		userTypeFloat, ok := claims[string(common.CtxUserType)].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user type in token.",
			})
			c.Abort()
			return
		}

		rctx := &requests.RequestContext{
			Ctx:       ctx,
			UserID:    uint(userIDFloat),
			UserType:  types.UserType(int8(userTypeFloat)),
			SessionID: uint(sessionIDFloat),
		}

		c.Set("rctx", rctx)
		c.Next()
	}
}
