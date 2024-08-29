package mailer

import (
	"log"
	"main/pkg/globals"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Config struct {
	Subject string
	Body    string
	To      string
}

func Send(config Config) (bool, error) {
	to := mail.NewEmail("alfashop", config.To)
	from := mail.NewEmail("alfashop", "ucha1bokeria@gmail.com")
	message := mail.NewSingleEmail(from, config.Subject, to, "", config.Body)
	client := sendgrid.NewSendClient(globals.Env.SENDGRID_API_KEY)
	_, err := client.Send(message)

	// log.Println(err, config)
	if err != nil {
		log.Println(err, config)
		return false, err
	} else {
		return true, nil
	}
}
