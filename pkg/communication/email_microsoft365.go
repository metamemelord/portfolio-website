package communication

import (
	"context"
	"fmt"
	"log"
	"os"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/me"
	msGraphModels "github.com/microsoftgraph/msgraph-sdk-go/models"
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

	auth, err := a.NewAzureIdentityAuthenticationProvider(cred)

	if err != nil {
		log.Println("Error cannot initialize the microsoft 365 client", err)
		return &microsoft365EmailService{}
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)

	if err != nil {
		log.Println("Error cannot initialize the microsoft 365 client", err)
		return &microsoft365EmailService{}
	}

	client := msgraphsdk.NewGraphServiceClient(adapter)

	return &microsoft365EmailService{mgGraphClient: client}
}

type microsoft365EmailService struct {
	mgGraphClient *msgraphsdk.GraphServiceClient
}

func (m *microsoft365EmailService) Send(ctx context.Context, email *Email) (interface{}, error) {
	emailBody := `<div style="color:#454545;margin:0;padding:0;"><h1>Hi <b>%s</b>, thanks for writing to me!</h1>I have received your message<p style="background:#efefef;color:#000;font-size:1rem;"><tt>%s</tt></p>at %s.<br><br>I will get back to you soon!<br><br><div>Regards,<br><b>Gaurav Saini<br>(778) 858-3884<br>https://gaurav.dev</b></div></div>`
	emailBody = fmt.Sprintf(emailBody, email.SenderName, email.Body, email.DateTime)

	var message msGraphModels.Messageable = msGraphModels.NewMessage()
	subject := "Hola from Gaurav!"
	message.SetSubject(&subject)

	body := msGraphModels.NewItemBody()
	contentType := msGraphModels.HTML_BODYTYPE
	body.SetContentType(&contentType)
	body.SetContent(&emailBody)
	message.SetBody(body)

	mainRecipient := msGraphModels.NewRecipient()
	mainRecipient.SetEmailAddress(prepareRecipientEmailAddress(email.SenderName, email.SenderEmail))
	message.SetToRecipients([]msGraphModels.Recipientable{mainRecipient})

	selfRecipient := msGraphModels.NewRecipient()
	selfRecipient.SetEmailAddress(prepareRecipientEmailAddress("Gaurav Saini", SELF_EMAIL))
	message.SetBccRecipients([]msGraphModels.Recipientable{selfRecipient})

	saveToSentItems := false
	sendEmailRequest := me.NewSendMailPostRequestBody()
	sendEmailRequest.SetMessage(message)
	sendEmailRequest.SetSaveToSentItems(&saveToSentItems)
	return "", m.mgGraphClient.UsersById(MS_GRAPH_SELF_USER_ID).
		SendMail().Post(ctx, sendEmailRequest, nil)
}

func prepareRecipientEmailAddress(name, email string) msGraphModels.EmailAddressable {
	emailAddress := msGraphModels.NewEmailAddress()
	emailAddress.SetName(&name)
	emailAddress.SetAddress(&email)
	return emailAddress
}
