package intraprocess

import (
	"github.com/suyash/mono-to-micro/pkg/orders/application"
)

type OrdersInterface struct {
	service *application.OrdersService
}

func NewOrderInterface(service *application.OrdersService) *OrdersInterface {
	return &OrdersInterface{
		service: service ,
	}
}

func (oi *OrdersInterface) MarkOrderAsPaid(orderID string){
	return oi.service.MarkOrderAsPaid(application.MarkOrderAsPaidCommand{
		OrderID: orders.ID(orderID),
	})
}
