package services

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/go-gomail/gomail"
)

type EmailService struct {
	SMTPHost    string
	SMTPPort    int
	SenderEmail string
	SenderPass  string
	IsLive      bool
}

// NewEmailService initializes the email service
func NewEmailService(container *storage.Container) *EmailService {
	return &EmailService{
		SMTPHost:    container.Env.Server.Smtp.Host,
		SMTPPort:    container.Env.Server.Smtp.Port,
		SenderEmail: container.Env.Server.Smtp.Email,
		SenderPass:  container.Env.Server.Smtp.Password,
		IsLive:      container.Env.Server.IsLive,
	}
}

// SendEmail sends an email with optional attachments
func (service *EmailService) SendEmail(to, subject, body string, attachments []string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", service.SenderEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	// Attach files if provided
	for _, attachment := range attachments {
		mailer.Attach(attachment)
	}

	dialer := gomail.NewDialer(service.SMTPHost, service.SMTPPort, service.SenderEmail, service.SenderPass)

	if service.IsLive {
		dialer.TLSConfig = &tls.Config{ServerName: service.SMTPHost}
	} else {
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Skip verification (for testing)
	}

	// Send the email
	if err := dialer.DialAndSend(mailer); err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	fmt.Println("✅ Email sent successfully to", to)
	return nil
}
