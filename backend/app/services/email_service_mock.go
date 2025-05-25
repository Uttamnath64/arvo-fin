package services

import (
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type TestEmail struct {
	SMTPHost    string
	SMTPPort    int
	SenderEmail string
	SenderPass  string
	IsLive      bool
}

// NewEmailService initializes the email service
func NewTestEmail(container *storage.Container) *TestEmail {
	return &TestEmail{
		SMTPHost:    container.Env.Server.Smtp.Host,
		SMTPPort:    container.Env.Server.Smtp.Port,
		SenderEmail: container.Env.Server.Smtp.Email,
		SenderPass:  container.Env.Server.Smtp.Password,
		IsLive:      container.Env.Server.IsLive,
	}
}

// SendEmail sends an email with optional attachments
func (service *TestEmail) SendEmail(to, subject, templateFile string, data map[string]string, attachments []string) error {
	return nil
}
