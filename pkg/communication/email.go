package communication

import (
	"context"
	"os"
)

var SELF_EMAIL = "hello@gaurav.dev"

func init() {
	emailFromEnv := os.Getenv("SELF_EMAIL")
	if len(emailFromEnv) > 0 {
		SELF_EMAIL = emailFromEnv
	}
}

type EmailClient interface {
	Send(context.Context, *Email) (interface{}, error)
}

type Email struct {
	DateTime    string
	Sender      string
	SenderName  string
	SenderEmail string
	Subject     string
	Body        string
	BodyHTML    string
}
