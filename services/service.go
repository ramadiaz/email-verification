package services

import (
	"email-verification/repositories"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type CompServices interface {
	TokenSend(destination string) error
	TokenVerify(email string, token string) error 
	RegistUser(email string, password string) error
}

type compServices struct {
	repo repositories.CompRepositories
}

func NewServices(r repositories.CompRepositories) *compServices {
	return &compServices{
		repo: r,
	}
}

func (s *compServices) TokenSend(destination string) error {
	randomNumber := rand.Intn(900000) + 100000

	token := strconv.Itoa(randomNumber)

	err := s.repo.InsertToken(destination, token)
	if err != nil{
		return err
	}

	err = godotenv.Load()
	if err != nil {
	  return err
	}
  
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	server := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")

	i, err := strconv.Atoi(smtpPort)
	if err != nil{
		return err
	}

	message := fmt.Sprintf(
		`<html>
			<head>
				<title>Email Verification</title>
			</head>
			<body>
				<p>Dear User,</p>
				<p>
					Thank you for registering with our platform. To complete your registration, please use the following verification code:
				</p>
				<p style="font-size: 24px; font-weight: bold;">%s</p>
				<p>
					This code will expire in 24 hours. If you did not request this verification, please ignore this email.
				</p>
				<p>Best regards,<br>Your Company Team</p>
			</body>
		</html>`,
		token,
	)
	

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

func (s *compServices) TokenVerify(email string, token string) error {
	data, err := s.repo.GetUser(email)
	if err != nil {
		return err
	}

	if (data.Token != token) {
		return errors.New("invalid or expired verification code")
	}

	err = s.repo.VerifyEmail(email)
	if err != nil {
		return err
	}

	return nil
}

func (s *compServices) RegistUser(email string, password string) error {
	err := s.repo.RegistUser(email, password)
	if err != nil {
		return err
	}

	return nil
}