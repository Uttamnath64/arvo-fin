package validater

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// regex
var (
	Name         *regexp.Regexp
	Email        *regexp.Regexp
	Username     *regexp.Regexp
	MobileNumber *regexp.Regexp
	OTP          *regexp.Regexp
)

type Validater struct{}

func New() *Validater {
	Email = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&’*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
	Username = regexp.MustCompile("^[A-Za-z0-9_.]{3,15}$")
	Name = regexp.MustCompile("^[A-Za-z ]{3,30}$")
	OTP = regexp.MustCompile("^[0-9]{6}$")
	MobileNumber = regexp.MustCompile("^[0-9]{10}$")
	return &Validater{}
}

func (v *Validater) IsValidEmail(data string) error {
	if !Email.MatchString(data) {
		return errors.New("Invalid email address!")
	}
	return nil
}

func (v *Validater) IsValidUsername(data string) error {
	if !Username.MatchString(data) {
		return errors.New("Invalid username!")
	}
	return nil
}

func (v *Validater) IsValidName(data string) error {
	if !Name.MatchString(data) {
		return errors.New("Invalid name!")
	}
	return nil
}

func (v *Validater) IsValidPassword(data string) error {

	if len(data) < 6 || len(data) > 20 {
		return errors.New("Password must be 6 to 20 characters long!")
	}

	hasDigit := false
	hasLower := false
	hasUpper := false
	hasSpecial := false

	for _, char := range data {
		switch {
		case '0' <= char && char <= '9':
			hasDigit = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' || char == '&' || char == '*':
			hasSpecial = true
		}
	}

	if !(hasDigit && hasLower && hasUpper && hasSpecial) {
		return errors.New("Password must include upper & lower case, a number, and a special character")
	}

	return nil
}

func (v *Validater) IsValidOTP(data string) error {
	if !OTP.MatchString(data) {
		return errors.New("Invalid otp!")
	}
	return nil
}

func (v *Validater) IsValidMobileNumber(data string) error {
	if !MobileNumber.MatchString(data) {
		return errors.New("Invalid mobile number!")
	}
	return nil
}

func (v *Validater) IsValidID(id uint) bool {
	return id > 0
}

func (v *Validater) HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("could not hash password " + err.Error())
	}
	return string(hashPassword), nil
}

func (v *Validater) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
