package main

import (
	core "notification/Core"
	"notification/entities"
	"notification/services"
)

var logger = core.NewFileLogger()

func main() {

	order1 := entities.Order{
		ID:               1,
		UserFullName:     "Max Sullivan",
		UserId:           "09156207562",
		Price:            float64(100),
		Status:           true,
		NotificationType: entities.Sms,
	}
	order2 := entities.Order{
		ID:               2,
		UserFullName:     "Max John",
		UserId:           "0915",
		Price:            float64(8000),
		Status:           true,
		NotificationType: entities.Email,
	}
	//dependency inversion

	orderService := services.NewOrderService()
	//add error

	err, _ := orderService.CreateOrder(&order1)

	//add log

	if err != nil {
		logger.Error().Interface("entityInfo", order1).Err(err).Msg("order 1 is not valid")
	}

	err, _ = orderService.CreateOrder(&order2)

	if err != nil {
		logger.Error().Interface("entityInfo", order2).Err(err).Msg("order 2 is not valid")
	}

}
