package main

import (
	"bytes"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"gopkg.in/gomail.v2"
)

type Recipient struct {
	toEmails  []string
	ccEmails  []string
	bccEmails []string
}

func main() {
	message := "This is message body of an email"
	subject := "This is subject of an email"
	fromEmail := "<from_mail>"

	recipient := Recipient{
		toEmails:  []string{"<toEmails_1>", "<toEmails_2>"},
		ccEmails:  []string{"<ccEmails_1>", "<ccEmails_2>"},
		bccEmails: []string{"<bccEmails_1>", "<bccEmails_2>"},
	}

	attachments := []string{"<attachment_path_1>", "<attachment_path_2>"}

	SendEmailRawSES(message, subject, fromEmail, recipient, attachments)
}

// SendEmailSES sends email to specified email IDs
func SendEmailRawSES(messageBody string, subject string, fromEmail string, recipient Recipient, attachments []string) {

	// create new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Println("Error occurred while creating aws session", err)
		return
	}

	// create raw message
	msg := gomail.NewMessage()

	// set to section
	var recipients []*string
	for _, r := range recipient.toEmails {
		recipient := r
		recipients = append(recipients, &recipient)
	}

	// Set to emails
	msg.SetHeader("To", recipient.toEmails...)

	// cc mails mentioned
	if len(recipient.ccEmails) != 0 {
		// Need to add cc mail IDs also in recipient list
		for _, r := range recipient.ccEmails {
			recipient := r
			recipients = append(recipients, &recipient)
		}
		msg.SetHeader("cc", recipient.ccEmails...)
	}

	// bcc mails mentioned
	if len(recipient.bccEmails) != 0 {
		// Need to add bcc mail IDs also in recipient list
		for _, r := range recipient.bccEmails {
			recipient := r
			recipients = append(recipients, &recipient)
		}
		msg.SetHeader("bcc", recipient.bccEmails...)
	}

	// create an SES session.
	svc := ses.New(sess)

	msg.SetAddressHeader("From", fromEmail, "<name>")
	msg.SetHeader("To", recipient.toEmails...)
	msg.SetHeader("Subject", "<subject>")
	msg.SetBody("text/html", `<email_body>`)

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
		log.Println("Error sending mail - ", err)
		return
	}

	log.Println("Email sent successfully to: ", recipient.toEmails)
}
