package repositories

import (
	"github.com/Uttamnath64/quixzap/internal/app/models"
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	container *storage.Container
}

func NewChat(container *storage.Container) *Chat {
	return &Chat{
		container: container,
	}
}

func (repo *Chat) Create(rctx *requests.RequestContext, chat *models.Chat) (uint, error) {
	err := repo.container.Config.DB.WithContext(rctx.Ctx).Create(chat).Error
	if err != nil {
		return 0, err
	}
	return chat.ID, nil
}

func (repo *Chat) UUIDExists(rctx *requests.RequestContext, uuid uuid.UUID) error {
	var count int64

	err := repo.container.Config.DB.WithContext(rctx.Ctx).Model(&models.Chat{}).
		Where("uuid = ?", uuid).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
