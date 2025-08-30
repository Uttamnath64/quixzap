package services

import (
	"github.com/Uttamnath64/quick-connect/app/utils/requests"
	"github.com/Uttamnath64/quick-connect/app/utils/responses"
	"github.com/Uttamnath64/quick-connect/pkg/validater"
	"github.com/google/uuid"
)

var (
	Validate *validater.Validater
)

type ChatService interface {
	Create(rctx *requests.RequestContext, uuid uuid.UUID, deviceInfo string, ip string) responses.ServiceResponse
}
