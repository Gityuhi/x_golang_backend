package email

import (
	"context"
	"fmt"
	"net/smtp"
	"x_golang_api/internal/domain/service"
)

type smtpSender struct {
	host     string
	port     string
	username string
	password string
	from     string
}

func NewSendEmail(host, port, username, password, from string) service.EmailSender {
	return &smtpSender{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (s *smtpSender) SendMail(ctx context.Context, activationURL string, toEmail string) error {
	msg := []byte("To: " + toEmail + "\r\n" +
		"From: " + s.from + "\r\n" +
		"Subject: アカウントを有効化してください\r\n" +
		"\r\n" +
		"以下のリンクをクリックして、アカウントの有効化を完了してください。\r\n" +
		activationURL)

	client, err := smtp.Dial(s.host + ":" + s.port)
	if err != nil {
		return fmt.Errorf("failed to dial smtp server: %w", err)
	}
	defer client.Close()

	if err := client.Mail(s.from); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	if err := client.Rcpt(toEmail); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to start data transmission: %w", err)
	}

	_, err = wc.Write(msg)
	if err != nil {
		return fmt.Errorf("failed to write message body: %w", err)
	}

	err = wc.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	client.Quit()

	return nil
}