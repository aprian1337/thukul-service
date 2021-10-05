package smtp

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type ConfigSmtpUsecase struct {
	SmtpHost         string
	SmtpPort         int
	SmtpSenderName   string
	SmtpAuthEMail    string
	SmtpAuthPassword string
}

func NewSmtpUsecase(host string, port int, senderName string, email string, password string) *ConfigSmtpUsecase {
	return &ConfigSmtpUsecase{
		SmtpHost:         host,
		SmtpPort:         port,
		SmtpSenderName:   senderName,
		SmtpAuthEMail:    email,
		SmtpAuthPassword: password,
	}
}

func (c *ConfigSmtpUsecase) SendMailSMTP(ctx context.Context, domain Domain) error {
	bcc := []string{"dwiky.dev@gmail.com"}
	mime := "\r\n" + "MIME-Version: 1.0\r\n" + "Content-Type: text/html; charset=\"utf-8\"\r\n\r\n"
	body := "From: " + c.SmtpSenderName + "\n" +
		"To: " + strings.Join(domain.MailTo, ",") + "\n" +
		"Cc: " + strings.Join(bcc, ",") + "\n" +
		"Subject: " + domain.Subject + mime +
		domain.Message

	auth := smtp.PlainAuth("", c.SmtpAuthEMail, c.SmtpAuthPassword, c.SmtpHost)
	smtpAddr := fmt.Sprintf("%s:%d", c.SmtpHost, c.SmtpPort)

	err := smtp.SendMail(smtpAddr, auth, c.SmtpAuthEMail, append(domain.MailTo, bcc...), []byte(body))
	if err != nil {
		return err
	}
	log.Println("Mail sent!")

	return nil
}
