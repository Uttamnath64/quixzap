package repositories

import (
	"github.com/Uttamnath64/quixzap/internal/app/common/types"
	"github.com/Uttamnath64/quixzap/internal/app/models"
	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"github.com/Uttamnath64/quixzap/internal/app/utils/responses"
)

type AuthRepository interface {
	GetSessionByUser(rctx *requests.RequestContext, userId uint, userType types.UserType, signedToken string) (*models.Session, error)
	GetSessionByRefreshToken(rctx *requests.RequestContext, refreshToken string, userType types.UserType) (*models.Session, error)
	CreateSession(rctx *requests.RequestContext, session *models.Session) (uint, error)
	UpdateSession(rctx *requests.RequestContext, id uint, refreshToken string, expiresAt int64) error
	DeleteSession(rctx *requests.RequestContext, sessionID uint) error
}

type AvatarRepository interface {
	Get(rctx *requests.RequestContext, id uint) (*models.Avatar, error)
	GetByNameAndType(rctx *requests.RequestContext, name string, avatarType types.AvatarType) *models.Avatar
	AvatarByTypeExists(rctx *requests.RequestContext, id uint, avatarType types.AvatarType) error
	GetAvatarsByType(rctx *requests.RequestContext, avatarType types.AvatarType) (*[]models.Avatar, error)
	Create(rctx *requests.RequestContext, payload models.Avatar) (uint, error)
	Update(rctx *requests.RequestContext, id uint, payload requests.AvatarRequest) error
}

type MemberRepository interface {
	GetMemberByUsernameOrEmail(rctx *requests.RequestContext, username string, email string, member *models.Member) error
	UsernameExists(rctx *requests.RequestContext, username string) error
	EmailExists(rctx *requests.RequestContext, email string) error
	Create(rctx *requests.RequestContext, member *models.Member) (uint, error)
	UpdatePasswordByEmail(rctx *requests.RequestContext, email, newPassword string) error
	GetMember(rctx *requests.RequestContext, memberId uint, member *models.Member) error
	Get(rctx *requests.RequestContext, memberId uint) (*responses.MemberResponse, error)
	Update(rctx *requests.RequestContext, memberId uint, payload requests.UpdateMember) error
}
