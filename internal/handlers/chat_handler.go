package handlers

import (
	"net/http"

	"github.com/Uttamnath64/quixzap/internal/app/common/types"
	"github.com/Uttamnath64/quixzap/internal/app/services"
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/Uttamnath64/quixzap/internal/app/utils/responses"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Chat struct {
	container   *storage.Container
	chatService services.ChatService
}

func NewChat(container *storage.Container) *Chat {
	return &Chat{
		container:   container,
		chatService: services.NewChat(container),
	}
}

func (handler *Chat) Create(c *gin.Context) {
	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}
	rctx.UserType = types.UserTypeUser

	serviceResponse := handler.chatService.Create(rctx, uuid.New(), c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}
