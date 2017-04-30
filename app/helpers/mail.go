package helpers

import (
	"fmt"

	"gopkg.in/mailgun/mailgun-go.v1"
)

// SendEmail ...
func SendEmail(from, to, subject, text string) error {
	domain := "mitty.co"
	apiKey := "key-04584e6b040b239e20234768c1209c77"
	publicAPIKey := "pubkey-4918c3ad10a37886f15a84d0def08f86"
	mg := mailgun.NewMailgun(domain, apiKey, publicAPIKey)
	message := mailgun.NewMessage(
		from,
		subject,
		text,
		to)
	resp, id, err := mg.Send(message)
	if err != nil {
		return err
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
