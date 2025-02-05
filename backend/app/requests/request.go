package requests

import "github.com/Uttamnath64/arvo-fin/pkg/validater"

var (
	Validate *validater.Validater
)

func NewResponse() {
	Validate = validater.New()
}
