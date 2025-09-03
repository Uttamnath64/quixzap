package handlers

import (
	"net/http"

	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"github.com/Uttamnath64/quixzap/internal/app/utils/responses"
	"github.com/gin-gonic/gin"
)

type Main struct {
	container   *storage.Container
	authService services.AuthService
}

func NewMain(container *storage.Container) *Main {
	return &Main{
		container:   container,
		authService: services.NewAuth(container),
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

func (h *Main) RefreshToken(c *gin.Context) {
	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.Token

	if !bindAndValidateJson(c, &payload) {
		return
	}

	serviceResponse := h.authService.RefreshToken(rctx, payload, c.Request.UserAgent(), c.ClientIP())
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
