package libs

import (
	"log"
	"os"

	"github.com/AnggaArdhinata/indochat/src/models"
	"gopkg.in/gomail.v2"
)

func SendEmail(payload []models.OrderPayload) error {

	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	var CONFIG_SENDER_NAME = os.Getenv("GOMAIL_SENDER_NAME")
	var CONFIG_AUTH_EMAIL = os.Getenv("GOMAIL_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("GOMAIL_PASS")

	for _, mail := range payload {

		mailer := gomail.NewMessage()
		mailer.SetHeader("Subject", "Reminder Pending Payment")
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", mail.Email)
		mailer.SetBody("text/html", "Hello, "+mail.Name+", you have order: "+mail.Product)

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	log.Println("Mail sent!")

	return nil
}
