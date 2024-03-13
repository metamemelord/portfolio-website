package communication

import (
	"context"
	"fmt"
	"log"
	"os"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

var MS_GRAPH_SELF_USER_ID = ""

func init() {
	MS_GRAPH_SELF_USER_ID = os.Getenv("MS_GRAPH_SELF_USER_ID")
}

func NewMicrosoft365EmailClient() EmailClient {
	cred, err := azidentity.NewClientSecretCredential(
		os.Getenv("MS_TENANT_ID"),
		os.Getenv("MS_CLIENT_ID"),
		os.Getenv("MS_EMAIL_KEY"),
		nil,
	)

	if err != nil {
		log.Println("Error cannot initialize the microsoft 365 client: ", err)
		return &microsoft365EmailService{}
	}

	cl, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, nil)

	if err != nil {
		log.Println("Error cannot initialize the microsoft 365 client: ", err)
		return &microsoft365EmailService{}
	}

	return &microsoft365EmailService{msGraphClient: cl}
}

type microsoft365EmailService struct {
	msGraphClient *msgraphsdk.GraphServiceClient
}

func (m *microsoft365EmailService) Send(ctx context.Context, email *Email) (interface{}, error) {
	emailBody := `<div style="color:#454545;margin:0;padding:0;"><h1>Hi <b>%s</b>, thanks for writing to me!</h1>I have received your message<p style="background:#efefef;color:#000;font-size:1rem;"><tt>%s</tt></p>at %s.<br><br>I will get back to you soon!<br><br><div>Regards,<br><b>Gaurav Saini<br>(778) 858-3884<br>https://gaurav.dev</b></div></div>`
	emailBody = fmt.Sprintf(emailBody, email.SenderName, email.Body, email.DateTime)
	subject := "Hola from Gaurav!"
	contentType := models.HTML_BODYTYPE

	// Create a new message
	message := models.NewMessage()
	message.SetSubject(&subject)

	messageBody := models.NewItemBody()
	messageBody.SetContent(&emailBody)
	messageBody.SetContentType(&contentType)
	message.SetBody(messageBody)

	mainRecipient := models.NewRecipient()
	mainRecipient.SetEmailAddress(prepareRecipientEmailAddress(email.SenderName, email.SenderEmail))
	message.SetToRecipients([]models.Recipientable{mainRecipient})

	selfRecipient := models.NewRecipient()
	selfRecipient.SetEmailAddress(prepareRecipientEmailAddress("Gaurav Saini", SELF_EMAIL))
	message.SetBccRecipients([]models.Recipientable{selfRecipient})

	sendMailBody := users.NewItemSendMailPostRequestBody()
	sendMailBody.SetMessage(message)

	saveToSentItems := false
	sendMailBody.SetSaveToSentItems(&saveToSentItems)

	return "", m.msGraphClient.Users().ByUserId(MS_GRAPH_SELF_USER_ID).SendMail().Post(ctx, sendMailBody, nil)
}

func prepareRecipientEmailAddress(name, email string) models.EmailAddressable {
	emailAddress := models.NewEmailAddress()
	emailAddress.SetName(&name)
	emailAddress.SetAddress(&email)
	return emailAddress
}
