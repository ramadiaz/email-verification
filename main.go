package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	server := os.Getenv("SMTP_SERVER")
	port := os.Getenv("SMTP_PORT")

	i, err := strconv.Atoi(port)
	if err != nil{
		panic(err)
	}


	// Set up the email message
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", "ramadiaz221@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "This is the <b>HAII</b> message body.")

	// Create a new SMTP client session
	d := gomail.NewDialer(server, i, email, password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
}
