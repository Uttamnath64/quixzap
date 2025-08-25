package requests

import (
	"context"

	"github.com/Uttamnath64/quick-connect/app/common/types"
)

type RequestContext struct {
	Ctx       context.Context
	UserID    uint
	UserType  types.UserType
	SessionID uint
}
