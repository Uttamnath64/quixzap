package repositories

import (
	"strings"

	"github.com/Uttamnath64/quixzap/app/config"
	"github.com/Uttamnath64/quixzap/app/models"
	"github.com/Uttamnath64/quixzap/app/utils/requests"
	"github.com/Uttamnath64/quixzap/app/utils/responses"
	"gorm.io/gorm"
)

type Member struct {
	mysql *config.MySQL
}

func NewMember(mysql *config.MySQL) *Member {
	return &Member{
		mysql: mysql,
	}
}

func (repo *Member) GetMemberByUsernameOrEmail(rctx *requests.RequestContext, username string, email string, member *models.Member) error {
	return repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.Member{}).
		Where("username = ? or email = ?", username, strings.ToLower(email)).First(member).Error
}

func (repo *Member) UsernameExists(rctx *requests.RequestContext, username string) error {
	var count int64

	err := repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.Member{}).
		Where("username = ?", username).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *Member) EmailExists(rctx *requests.RequestContext, email string) error {
	var count int64

	err := repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.Member{}).
		Where("email = ?", strings.ToLower(email)).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *Member) Create(rctx *requests.RequestContext, member *models.Member) (uint, error) {
	err := repo.mysql.ReadWriteDB.WithContext(rctx.Ctx).Create(member).Error
	if err != nil {
		return 0, err
	}
	return member.ID, nil
}

func (repo *Member) UpdatePasswordByEmail(rctx *requests.RequestContext, email, newPassword string) error {
	result := repo.mysql.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.Member{}).
		Where("email = ?", email).
		Update("password", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *Member) GetMember(rctx *requests.RequestContext, memberId uint, member *models.Member) error {
	if err := repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Where("id = ?", memberId).First(member).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Member) Get(rctx *requests.RequestContext, memberId uint) (*responses.MemberResponse, error) {
	var member models.Member
	var avatar models.Avatar
	var response responses.MemberResponse

	query := repo.mysql.ReadOnlyDB.WithContext(rctx.Ctx).Table(member.GetName()+" u").
		Joins("JOIN "+avatar.GetName()+" a ON a.id = u.avatar_id").Where("u.id = ?", memberId)

	err := query.Select("u.id, u.name, u.username, u.email, a.id as avatar_id, a.icon as avatar_icon").
		Scan(&response).Error

	if err != nil {
		return nil, err // Other errors
	}
	if response.Id == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &response, nil
}

func (repo *Member) Update(rctx *requests.RequestContext, memberId uint, payload requests.UpdateMember) error {
	result := repo.mysql.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.Member{}).
		Where("id = ?", memberId).
		Updates(map[string]interface{}{
			"name":      payload.Name,
			"username":  payload.Username,
			"avatar_id": payload.AvatarId,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
