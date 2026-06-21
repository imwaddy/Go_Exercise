package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
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

	SendEmailSES(message, subject, fromEmail, recipient)
}

// SendEmailSES sends email to specified email IDs
func SendEmailSES(messageBody string, subject string, fromEmail string, recipient Recipient) {

	// create new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Println("Error occurred while creating aws session", err)
		return
	}

	// set to section
	var recipients []*string
	for _, r := range recipient.toEmails {
		recipient := r
		recipients = append(recipients, &recipient)
	}

	// set cc section
	var ccRecipients []*string
	if len(recipient.ccEmails) > 0 {
		for _, r := range recipient.ccEmails {
			ccrecipient := r
			ccRecipients = append(ccRecipients, &ccrecipient)
		}
	}

	// set bcc section
	var bccRecipients []*string
	if len(recipient.bccEmails) > 0 {
		for _, r := range recipient.bccEmails {
			bccrecipient := r
			recipients = append(recipients, &bccrecipient)
		}
	}

	// create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{

		// Set destination emails
		Destination: &ses.Destination{
			CcAddresses:  ccRecipients,
			ToAddresses:  recipients,
			BccAddresses: bccRecipients,
		},

		// Set email message and subject
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(messageBody),
				},
			},

			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},

		// send from email
		Source: aws.String(fromEmail),
	}

	// Call AWS send email function which internally calls to SES API
	_, err = svc.SendEmail(input)
	if err != nil {
		log.Println("Error sending mail - ", err)
		return
	}

	log.Println("Email sent successfully to: ", recipient.toEmails)
}
