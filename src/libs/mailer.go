package libs

import (
	"log"
	"os"
	"strconv"

	"github.com/AnggaArdhinata/indochat/src/models"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

func SendEmail(payload []models.OrderPayload) error {

	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	var CONFIG_AUTH_EMAIL = os.Getenv("GOMAIL_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("GOMAIL_PASS")

	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      os.Getenv("COMPANY_NAME"),
			Link:      os.Getenv("BASE_URL"),
			Logo:      "https://mailtrap.io/wp-content/uploads/2019/08/How-to-Send-and-Receive-Emails-with-Go_Featured-Image.png",
			Copyright: "Copyright © 2023 Ardhinata Corp. All rights reserved.",
		},
	}

	for _, mail := range payload {

		price := strconv.Itoa(mail.Price)

		id := strconv.Itoa(mail.Id)

		emailBody, err := h.GenerateHTML(hermes.Email{
			Body: hermes.Body{
				Name: mail.Name,
				Intros: []string{
					"Thankyou for purchasing our product, here are the details of your order.",
				},
				Dictionary: []hermes.Entry{
					{Key: "Product", Value: mail.Product},
					{Key: "Description", Value: mail.Description},
					{Key: "Price", Value: "Rp. " + price},
				},
				Actions: []hermes.Action{
					{
						Instructions: "To confirm your payment, please click here:",
						Button: hermes.Button{
							Color: "#22BC66", // Optional action button color
							Text:  "Confirm your payment",
							Link:  os.Getenv("BASE_URL")+"api/v1/order/verify/" + id,
						},
					},
				},
				Outros: []string{
					"Need help, or have questions? Just reply to this email, we'd love to help.",
				},
			},
		})
		if err != nil {
			return err
		}

		mailer := gomail.NewMessage()
		mailer.SetHeader("Subject", "[REMINDER] You Have Pending Payment !")
		mailer.SetHeader("From", os.Getenv("GOMAIL_SENDER_NAME"))
		mailer.SetHeader("To", mail.Email)
		mailer.SetBody("text/html", emailBody)

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err = dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent to: " + mail.Email)
	}

	return nil

}
