package service

import "context"

type EmailSender interface {
	SendMail(ctx context.Context, url string, email string) error
}