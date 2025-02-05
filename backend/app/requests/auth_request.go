package requests

type LoginRequest struct {
	UsernameEmail string `json:"usernameEmail" binding:"required"`
	Password      string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Username     string `json:"Username" binding:"required"`
	MobileNumber string `json:"mobileNumber" binding:"required"`
	Password     string `json:"password" binding:"required"`
	OTP          string `json:"otp" binding:"required"`
}

type SentOTPRequest struct {
	Email string `json:"email" binding:"required"`
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

func (r *RegisterRequest) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if err := Validate.IsValidUsername(r.Username); err != nil {
		return err
	}
	if err := Validate.IsValidEmail(r.Email); err != nil {
		return err
	}
	if err := Validate.IsValidMobileNumber(r.MobileNumber); err != nil {
		return err
	}
	if err := Validate.IsValidPassword(r.Password); err != nil {
		return err
	}
	if err := Validate.IsValidOTP(r.OTP); err != nil {
		return err
	}
	return nil
}

func (r *SentOTPRequest) IsValid() error {
	if err := Validate.IsValidEmail(r.Email); err != nil {
		return err
	}
	return nil
}
