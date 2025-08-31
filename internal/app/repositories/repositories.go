package repositories

import (
	"github.com/Uttamnath64/quick-connect/internal/app/common/types"
	"github.com/Uttamnath64/quick-connect/internal/app/models"
	"github.com/Uttamnath64/quick-connect/internal/app/utils/requests"
	"github.com/google/uuid"
)

type AuthRepository interface {
	GetSessionByRefreshToken(rctx *requests.RequestContext, refreshToken string) (*models.Session, error)
	CreateSession(rctx *requests.RequestContext, session *models.Session) (uint, error)
	UpdateSession(rctx *requests.RequestContext, id uint, refreshToken string, expiresAt int64) error
	DeleteSession(rctx *requests.RequestContext, sessionID uint) error
}

type AdminRepository interface {
	GetList(rctx *requests.RequestContext, userId uint) (*[]models.Admin, error)
	Get(rctx *requests.RequestContext, id, userId uint, userType types.UserType) (*models.Admin, error)
	Create(rctx *requests.RequestContext, portfolio models.Admin) error
	Update(rctx *requests.RequestContext, id, userId uint, payload requests.AdminRequest) error
	Delete(rctx *requests.RequestContext, id, userId uint) error
}

type UserRepository interface {
	// Create(rctx *requests.RequestContext, user *models.User) (uint, error)
	UpdatePasswordByEmail(rctx *requests.RequestContext, email, newPassword string) error
	// Get(rctx *requests.RequestContext, userId uint) (*models.User, error)
	Block(rctx *requests.RequestContext, userId uint) error
}

type ChatRepository interface {
	Create(rctx *requests.RequestContext, user *models.Chat) (uint, error)
	UUIDExists(rctx *requests.RequestContext, uuid uuid.UUID) error
	// GetAll(rctx *requests.RequestContext, email, newPassword string) error
}
