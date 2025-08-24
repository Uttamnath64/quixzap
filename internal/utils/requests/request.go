package requests

import "github.com/Uttamnath64/quick-connect/pkg/validater"

var (
	Validate *validater.Validater
)

func NewResponse() {
	Validate = validater.New()
}
