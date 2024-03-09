package externalServices

import (
	"fmt"

	"github.com/rs/zerolog"
)

type SmsService struct {
}

// func (s *SmsService) SendMessage(order *entities.Order) {
// 	fmt.Printf("Sms Sent: %v\n", order)
// }

// add 4
func (s *SmsService) SendNotify(receiver string, message string) error {
	//add error
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("Sms Sent , Receiver : %s Message : %s\n", receiver, message)

	//add logger
	logger.Info().Str("notifier", "SmsService").
		Dict("messageInfo", zerolog.Dict().
			Str("receiver", receiver).Str("message", message)).
		Msg("message sent ")
	return nil
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
