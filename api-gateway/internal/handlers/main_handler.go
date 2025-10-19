package handlers

import (
	"net/http"

	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/Uttamnath64/quixzap/app/services"
	"github.com/Uttamnath64/quixzap/app/utils/requests"
	"github.com/Uttamnath64/quixzap/app/utils/responses"
	"github.com/gin-gonic/gin"
)

type Main struct {
	appCtx      *appcontext.AppContext
	authService services.AuthService
}

func NewMain(appCtx *appcontext.AppContext) *Main {
	return &Main{
		appCtx: appCtx,
		// authService: services.NewAuth(container),
	}
}

func (h *Main) Register(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.Register
	if !bindAndValidateJson(c, &payload) {
		return
	}

	serviceResponse := h.authService.Register(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (h *Main) Login(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.Login
	if !bindAndValidateJson(c, &payload) {
		return
	}

	serviceResponse := h.authService.Login(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (h *Main) GetToken(c *gin.Context) {
	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.Token

	if !bindAndValidateJson(c, &payload) {
		return
	}

	serviceResponse := h.authService.GetToken(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (h *Main) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: "API is running smoothly.",
	})
}

func (h *Main) SendOTP(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.SentOTP
	if !bindAndValidateJson(c, &payload) {
		return
	}

	serviceResponse := h.authService.SendOTP(rctx, payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (h *Main) ResetPassword(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.ResetPassword
	if !bindAndValidateJson(c, &payload) {
		return
	}

	// Reset password
	serviceResponse := h.authService.ResetPassword(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (h *Main) Token(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.Token
	if !bindAndValidateJson(c, &payload) {
		return
	}

	// Get token
	serviceResponse := h.authService.GetToken(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})

}
