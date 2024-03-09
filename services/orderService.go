package services

import (
	"errors"
	"fmt"
	core "notification/Core"
	"notification/entities"
	"notification/externalServices"
)

var logger = core.NewFileLogger()

// Dependency Inversion depend on interface
type OrderService struct {
	externalServices.Notifier
}

//Add errors and logs here

func (orderService *OrderService) CreateOrder(order *entities.Order) (error, *entities.Order) {
	//Add errors
	if !order.Status {
		return errors.New("order is not valid"), order
	}
	if order.Price < 1000 {
		return errors.New("order price is not valid"), order
	}

	fmt.Printf("Order created: %v\n", order)

	//Add log
	logger.Info().Interface("order", order).Msg("Order created")

	orderService.Notifier = externalServices.NewNotifier(order.NotificationType)

	//Add log

	logger.Info().Msgf("Notifier Founded: %T", orderService.Notifier)

	//add err
	err := orderService.SendNotify(order.UserId, "Order created")

	return err, order

}

// dependency to notifier is Created when Create the OrderService
func NewOrderService() *OrderService {
	return &OrderService{}
}
