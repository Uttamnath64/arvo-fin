package responses

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
)

type AccountTypeResponse struct {
	Type commonType.AccountType `json:"type"`
	Name string                 `json:"name"`
	Icon *models.Avatar         `json:"icon"`
}
