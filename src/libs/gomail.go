package libs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendEmail(email []string) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	var CONFIG_SENDER_NAME = os.Getenv("GOMAIL_SENDER_NAME")
	var CONFIG_AUTH_EMAIL = os.Getenv("GOMAIL_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("GOMAIL_PASS")

	for _, mail := range email {

		mailer := gomail.NewMessage()
		mailer.SetHeader("Subject", "Reminder Pending Payment")
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", mail)
		mailer.SetBody("text/html", "Hello, You have order pending payment order Click Link Bellow to Complete your payment")

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	log.Println("Mail sent!")

	return nil
}
