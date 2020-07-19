package communication

import "context"

type EmailSender interface {
	Send(context.Context, *Email) (interface{}, error)
}

type Email struct {
	DataTime       string
	Sender         string
	SenderName     string
	SenderEmail    string
	RecipientEmail string
	Subject        string
	Body           string
	BodyHTML       string
}
