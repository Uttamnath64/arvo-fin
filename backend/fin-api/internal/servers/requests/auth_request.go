package requests

type LoginRequest struct {
	UsernameEmail string `json:"usernameEmail" binding:"required"`
	Password      string `json:"password" binding:"required"`
}

func (r *LoginRequest) IsValid() error {
	emailErr := Validate.IsValidEmail(r.UsernameEmail)
	usernameErr := Validate.IsValidUsername(r.UsernameEmail)
	if emailErr != nil && usernameErr != nil {
		if emailErr != nil {
			return emailErr
		} else {
			return usernameErr
		}
	}

	if passErr := Validate.IsValidPassword(r.Password); passErr != nil {
		return passErr
	}
	return nil
}
