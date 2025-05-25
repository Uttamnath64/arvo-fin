package services

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"

	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/go-gomail/gomail"
)

type Email struct {
	SMTPHost    string
	SMTPPort    int
	SenderEmail string
	SenderPass  string
	IsLive      bool
}

// NewEmailService initializes the email service
func NewEmail(container *storage.Container) *Email {
	return &Email{
		SMTPHost:    container.Env.Server.Smtp.Host,
		SMTPPort:    container.Env.Server.Smtp.Port,
		SenderEmail: container.Env.Server.Smtp.Email,
		SenderPass:  container.Env.Server.Smtp.Password,
		IsLive:      container.Env.Server.IsLive,
	}
}

// SendEmail sends an email with optional attachments
func (service *Email) SendEmail(to, subject, templateFile string, data map[string]string, attachments []string) error {

	// Load email template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("error loading email template: %v", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return fmt.Errorf("error executing email template: %v", err)
	}

	// Set up the email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", service.SenderEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body.String())

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
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
