package requests

import (
	"context"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type RequestContext struct {
	Ctx       context.Context
	UserID    uint
	UserType  commonType.UserType
	SessionID uint
}
