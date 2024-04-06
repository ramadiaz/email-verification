package services

import (
	"email-verification/repositories"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type CompServices interface {
	EmailSend(destination string) error
}

type compServices struct {
	repo repositories.CompRepositories
}

func NewServices(r repositories.CompRepositories) *compServices {
	return &compServices{
		repo: r,
	}
}

func (s *compServices) EmailSend(destination string) error {
	randomNumber := rand.Intn(900000) + 100000

	token := strconv.Itoa(randomNumber)


	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	server := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")

	i, err := strconv.Atoi(smtpPort)
	if err != nil{
		return err
	}

	message := fmt.Sprintf("This is the <b>VERIFICATION CODE: %s</b> message body.", token)

	// Set up the email message``
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", destination)
	m.SetHeader("Subject", "Testing SMTP Email!")
	m.SetBody("text/html", message)

	// Create a new SMTP client session
	d := gomail.NewDialer(server, i, email, password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}