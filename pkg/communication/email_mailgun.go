package communication

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mailgun/mailgun-go/v4"
)

func NewMailgunEmailClient() EmailClient {
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))
	return &mailgunEmailService{mg: mg}
}

type mailgunEmailService struct {
	mg *mailgun.MailgunImpl
}

func (m *mailgunEmailService) Send(ctx context.Context, email *Email) (interface{}, error) {
	emailBody := `Sender:  %s <%s>
Date: %s
Message: 
%s
`
	emailBody = fmt.Sprintf(emailBody, email.SenderName, email.SenderEmail, email.DateTime, email.Body)

	message := m.mg.NewMessage(email.SenderEmail, email.Subject, emailBody, SELF_EMAIL)
	resp, id, err := m.mg.Send(ctx, message)

	if err != nil {
		log.Println(err, resp)
	}
	log.Printf("ID: %s Resp: %s\n", id, resp)
	return id, err
}
