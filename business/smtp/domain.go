package smtp

import "context"

type Domain struct {
	MailTo  []string
	Subject string
	Message string
}

type Usecase interface {
	SendMailSMTP(ctx context.Context, domain Domain) error
}
