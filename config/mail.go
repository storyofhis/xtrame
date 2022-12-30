package config

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func ConnectMail() error {
	Host := os.Getenv("CONFIG_SMTP_HOST")
	// Port := os.Getenv("CONFIG_SMTP_PORT")
	Sender := "PT. Makmur Subur Jaya <maulaizzaazizi@gmail.com>"
	Email := os.Getenv("CONFIG_AUTH_EMAIL")
	Pass := os.Getenv("CONFIG_AUTH_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", Sender)
	mailer.SetHeader("To", "azizi.maula@gmail.com", "maulaizzaazizi@gmail.com")
	mailer.SetAddressHeader("Cc", "maulaizzaazizi@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")

	dialer := gomail.NewDialer(Host, 587, Email, Pass)

	// body := "From: " + sender + "\n" +
	// 	"To: " + strings.Join(to, ",") + "\n" +
	// 	"Cc: " + strings.Join(cc, ",") + "\n" +
	// 	"Subject: " + subject + "\n\n" +
	// 	message

	// auth := smtp.PlainAuth("", email, pass, host)
	// smtpAddr := fmt.Sprintf(
	// 	"%s:%d", host, port,
	// )

	// err := smtp.SendMail(smtpAddr, auth, email, append(to, cc...), []byte(body))

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
