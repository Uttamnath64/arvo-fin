package requests

import (
	"errors"
	"strings"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type AvatarRequest struct {
	Name string                `json:"name" binding:"required"`
	Icon string                `json:"icon" binding:"required"`
	Type commonType.AvatarType `json:"type" binding:"required"`
}

func (r AvatarRequest) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if strings.TrimSpace(r.Icon) == "" {
		return errors.New("invalid icon")
	}
	if !r.Type.IsValid() {
		return errors.New("invalid avatar type")
	}

	return nil
}
