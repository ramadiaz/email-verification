package main

import (
	"email-verification/routers"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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
	smtpPort := os.Getenv("SMTP_PORT")
	port := os.Getenv("PORT")

	i, err := strconv.Atoi(smtpPort)
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

	r := gin.Default()

	api := r.Group("/api")

	routers.CompRouters(api)

	r.Run(":" + port)
}
