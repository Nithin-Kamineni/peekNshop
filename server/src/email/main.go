package main

import (
	"log"
	"net/smtp"
)

func main() {

	// Sender data.
	from := "flimentrpg@gmail.com"
	password := "rampagerpg"
	to := []string{"nithinkamineni1@mail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("My super secret message.")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}
