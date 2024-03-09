package externalServices

import (
	"fmt"
	core "notification/Core"

	"github.com/rs/zerolog"
)

var logger = core.NewFileLogger()

type EmailService struct {
}

//	func (e *EmailService) SendMessage(order *entities.Order) {
//		fmt.Printf("Email sent: %v\n", order)
//	}
//
// add 3 change output
func (e *EmailService) SendNotify(receiver string, message string) error {
	//add error
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("Email sent , receiver : %s , message : %s\n", receiver, message)

	//add logger
	logger.Info().Str("notifier", "EmailService").
		Dict("messageInfo", zerolog.Dict().
			Str("receiver", receiver).Str("message", message)).
		Msg("message sent ")

	return nil
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
