package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendVerifyEmail sends a welcome email to the registered user
func SendVerifyEmail(emailTo string, name string, verifyID string) error {
	var (
		host   = os.Getenv("SMTP_HOST") // smtp host
		sender = os.Getenv("SMTP_USER") // email sender
		pass   = os.Getenv("SMTP_PASS") // email sender password
		port   = os.Getenv("SMTP_PORT") // smtp port
	)

	// Here we do it all: connect to our server, set up a message and send it
	recipients := []string{emailTo}

	verifyURL := fmt.Sprintf("http://localhost:8080/user/verify/%s", verifyID)

	// email details and content
	msg := []byte(fmt.Sprintf("To: %s \r\n", emailTo) +
		"Subject: Welcome to Kafka example !!!\r\n" +
		"\r\n" +
		"Hi " + name + ", thank you for trying my Go Kafka example!\r\n" +
		"I hope you enjoyed this sample.\r\n" +
		fmt.Sprintf("Please following link to verify your email: %s\r\n", verifyURL))

	// SMTP authorization
	auth := smtp.PlainAuth("", sender, pass, host)

	// send mail sends an email with message
	err := smtp.SendMail(fmt.Sprintf("%s:%v", host, port), auth, sender, recipients, msg)
	if err != nil {
		return fmt.Errorf("failed to send email, %v", err.Error())
	}
	return nil
}
