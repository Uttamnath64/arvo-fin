package services

import (
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type AuthService struct {
	container   *storage.Container
	authService *services.AuthService
}

func NewAuthService(container *storage.Container) *AuthService {
	return &AuthService{
		container:   container,
		authService: services.NewAuthService(container),
	}
}

func (service *AuthService) Login(payload requests.LoginRequest, ip string) responses.ServiceResponse {
	return service.authService.Login(payload, ip)
}

func (service *AuthService) Register(payload requests.RegisterRequest, ip string) responses.ServiceResponse {
	return service.authService.Register(payload, ip)
}
