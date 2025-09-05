package repositories

import (
	"github.com/Uttamnath64/quixzap/internal/app/common/types"
	"github.com/Uttamnath64/quixzap/internal/app/models"
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"gorm.io/gorm"
)

type Avatar struct {
	container *storage.Container
}

func NewAvatar(container *storage.Container) *Avatar {
	return &Avatar{
		container: container,
	}
}

func (repo *Avatar) Get(rctx *requests.RequestContext, id uint) (*models.Avatar, error) {
	var avatar models.Avatar
	return &avatar, repo.container.Config.DB.WithContext(rctx.Ctx).Where("id = ?", id).First(&avatar).Error
}

func (repo *Avatar) GetByNameAndType(rctx *requests.RequestContext, name string, avatarType types.AvatarType) *models.Avatar {
	var avatar models.Avatar
	repo.container.Config.DB.WithContext(rctx.Ctx).Where("name = ? and type = ?", name, avatarType).First(&avatar)
	return &avatar
}

func (repo *Avatar) AvatarByTypeExists(rctx *requests.RequestContext, id uint, avatarType types.AvatarType) error {
	var count int64

	err := repo.container.Config.DB.WithContext(rctx.Ctx).Model(&models.Avatar{}).
		Where("id = ? AND type = ?", id, avatarType).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *Avatar) GetAvatarsByType(rctx *requests.RequestContext, avatarType types.AvatarType) (*[]models.Avatar, error) {
	var response []models.Avatar
	if err := repo.container.Config.DB.WithContext(rctx.Ctx).Model(&models.Avatar{}).Where("type = ?", avatarType).Scan(&response).Error; err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &response, nil
}

func (repo *Avatar) Create(rctx *requests.RequestContext, avatar models.Avatar) (uint, error) {
	return avatar.ID, repo.container.Config.DB.WithContext(rctx.Ctx).Create(&avatar).Error
}

func (repo *Avatar) Update(rctx *requests.RequestContext, id uint, payload requests.AvatarRequest) error {
	result := repo.container.Config.DB.WithContext(rctx.Ctx).Model(&models.Avatar{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name": payload.Name,
			"icon": payload.Icon,
			"type": payload.Type,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
