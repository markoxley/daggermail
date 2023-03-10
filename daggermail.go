package daggermail

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

type MailMessage struct {
	Recipients []string
	BCC        []string
	CC         []string
	Subject    string
	Message    string
}

func New(recipients []string, subject, message string) (*MailMessage, error) {
	if len(recipients) == 0 {
		return nil, errors.New("missing recipients")
	}
	return &MailMessage{
		Recipients: recipients,
		Subject:    subject,
		Message:    message,
	}, nil
}

func SendMail(msg *MailMessage) error {
	if configuration == nil {
		return errors.New("daggermail not configured")
	}
	if msg == nil {
		return errors.New("missing mail message")
	}
	if len(msg.Recipients) == 0 {
		return errors.New("missing recipients")
	}
	auth := smtp.PlainAuth(configuration.Identity, configuration.User, configuration.Password, configuration.Host)

	to := msg.Recipients

	message :=
		"Date: " + time.Now().Format(time.RFC1123Z) + "\r\n" +
			"From: " + configuration.From + "\r\n" +
			"Sender: " + configuration.From + "\r\n" +
			"To: " + strings.Join(msg.Recipients, ",") + "\r\n"
	if len(msg.BCC) > 0 {
		message += "BCC: " + strings.Join(msg.BCC, ",") + "\r\n"
	}
	if len(msg.CC) > 0 {
		message += "CC: " + strings.Join(msg.CC, ",") + "\r\n"
	}

	message += "Subject: " + msg.Subject + "\r\n" +

		"\r\n" + msg.Message + "\r\n"
	sendHost := fmt.Sprintf("%s:%d", configuration.Host, configuration.Port)
	err := smtp.SendMail(sendHost, auth, configuration.From, to, []byte(message))
	return err
}
