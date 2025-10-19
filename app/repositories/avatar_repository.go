package repositories

import (
	"github.com/Uttamnath64/quixzap/app/common/types"
	"github.com/Uttamnath64/quixzap/app/config"
	"github.com/Uttamnath64/quixzap/app/models"
	"github.com/Uttamnath64/quixzap/app/utils/requests"
	"gorm.io/gorm"
)

type Avatar struct {
	mysql *config.Mysql
}

func NewAvatar(mysql *config.Mysql) *Avatar {
	return &Avatar{
		mysql: mysql,
	}
}

func (repo *Avatar) Get(rctx *requests.RequestContext, id uint) (*models.Avatar, error) {
	var avatar models.Avatar
	return &avatar, repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Where("id = ?", id).First(&avatar).Error
}

func (repo *Avatar) GetByNameAndType(rctx *requests.RequestContext, name string, avatarType types.AvatarType) *models.Avatar {
	var avatar models.Avatar
	repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Where("name = ? and type = ?", name, avatarType).First(&avatar)
	return &avatar
}

func (repo *Avatar) AvatarByTypeExists(rctx *requests.RequestContext, id uint, avatarType types.AvatarType) error {
	var count int64

	err := repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.Avatar{}).
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
	if err := repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.Avatar{}).Where("type = ?", avatarType).Scan(&response).Error; err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &response, nil
}

func (repo *Avatar) Create(rctx *requests.RequestContext, avatar models.Avatar) (uint, error) {
	return avatar.ID, repo.mysql.ReadWriteDB.WithContext(rctx.Ctx).Create(&avatar).Error
}

func (repo *Avatar) Update(rctx *requests.RequestContext, id uint, payload requests.AvatarRequest) error {
	result := repo.mysql.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.Avatar{}).
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
