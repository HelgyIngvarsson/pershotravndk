package utils

import (
	"log"
	"net/smtp"
)

func SendMessage(recipients []string, msg string) error {
	auth := smtp.PlainAuth(
		"",
		"olehgavrilin@gmail.com",
		"01Ltrf,hm1996",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"olehgavrilin@gmail.com",
		recipients,
		[]byte(msg),
	)
	if err != nil {
		log.Print(err)
	}
	return nil
}
