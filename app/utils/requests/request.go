package requests

import "github.com/Uttamnath64/quixzap/pkg/validater"

var (
	Validate *validater.Validater
)

func NewResponse() {
	Validate = validater.New()
}
