package services

import (
	"errors"

	"github.com/Uttamnath64/quixzap/internal/app/common"
	"github.com/Uttamnath64/quixzap/internal/app/models"
	"github.com/Uttamnath64/quixzap/internal/app/repositories"
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"github.com/Uttamnath64/quixzap/internal/app/utils/responses"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	container *storage.Container
	chatRepo  repositories.ChatRepository
}

func NewChat(container *storage.Container) *Chat {
	return &Chat{
		container: container,
		chatRepo:  repositories.NewChat(container),
	}
}

func (service *Chat) Create(rctx *requests.RequestContext, uuid uuid.UUID, deviceInfo string, ip string) responses.ServiceResponse {
	// Validate portfolio
	if err := service.chatRepo.UUIDExists(rctx, uuid); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			service.container.Logger.Error("chat.service.Create-UUIDExists", "error", err.Error(), "uuid", uuid)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	}

	chat := models.Chat{
		UUID:       uuid,
		DeviceInfo: deviceInfo,
		IpAddress:  ip,
	}

	_, err := service.chatRepo.Create(rctx, &chat)
	if err != nil {
		service.container.Logger.Error("chat.service.Create-Create", "error", err.Error())
		return responses.ErrorResponse(common.StatusDatabaseError, "Unable to create chat.", err)
	}

	return responses.SuccessResponse(
		"Chat created successfully.",
		map[string]interface{}{
			"UUID": uuid.String(),
		},
	)
}
