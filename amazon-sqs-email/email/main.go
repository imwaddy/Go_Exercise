package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"gopkg.in/gomail.v2"
)

func main() {
	toEmails := []string{"mayur.wadekar@pb.com"}
	ccEmails := []string{}  //"email@gmail.com"}
	bccEmails := []string{} //"email@gmail.com"}
	attachments := []string{}
	fromEmail := "shipping.mypost@xyz.com"

	// new session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println("Error occurred while creating aws session", err)
		return
	}

	// create recipients
	var recipients []*string
	for _, r := range toEmails {
		recipient := r
		recipients = append(recipients, &recipient)
	}

	// Create an SES session.
	svc := ses.New(sess)

	htmlBody, err := os.ReadFile("template.html")
	if err != nil {
		fmt.Println("Error reading template:", err)
		return
	}

	// substitute template tokens
	replacer := strings.NewReplacer(
		"{{CompanyName}}", "My Company",
		"{{CompanyDomain}}", "mycompany.com",
		"{{ProductName}}", "My Product",
		"{{SupportEmail}}", "support@mycompany.com",
		"{{ProductDomain}}", "app.mycompany.com",
		"{{EcommerceGuideURL}}", "https://mycompany.com/guides/ecommerce",
		"{{FirstName}}", "User",
		"{{EmailAddress}}", toEmails[0],
		"{{WelcomeLink}}", "https://app.mycompany.com/welcome",
	)
	body := replacer.Replace(string(htmlBody))

	// create raw message
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", fromEmail, "Australia")
	msg.SetHeader("To", toEmails...)
	msg.SetHeader("Subject", "Sample subject")
	msg.SetBody("text/html", body)

	// cc mails mentioned
	if len(ccEmails) != 0 {
		// Need to add cc mail IDs also in recipient list
		for _, r := range ccEmails {
			recipient := r
			recipients = append(recipients, &recipient)
		}
		msg.SetHeader("cc", ccEmails...)
	}

	// As per documentation need to add CC and BCC emails with To field of email. Due to only single destination list.

	// bcc mails mentioned
	if len(bccEmails) != 0 {
		// Need to add bcc mail IDs also in recipient list
		for _, r := range bccEmails {
			recipient := r
			recipients = append(recipients, &recipient)
		}
		msg.SetHeader("bcc", bccEmails...)
	}

	// If attachments exists
	if len(attachments) != 0 {
		for _, f := range attachments {
			msg.Attach(f)
		}
	}

	// create a new buffer to add raw data
	var emailRaw bytes.Buffer
	msg.WriteTo(&emailRaw)

	// create new raw message
	message := ses.RawMessage{Data: emailRaw.Bytes()}

	input := &ses.SendRawEmailInput{Source: &fromEmail, Destinations: recipients, RawMessage: &message}

	// send raw email
	_, err = svc.SendRawEmail(input)
	if err != nil {
		fmt.Println("Error ", err)
	}
}
