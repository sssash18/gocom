package application

import "github.com/suyash/mono-to-micro/pkg/orders/domain/orders"

type productsService interface {
}

type paymentsService interface {
}

type OrdersService struct {
	repo      *Repository
	readModel *orderReadModel
}

type PlaceOrderCommand struct {
	OrderID   orders.ID
	ProductID orders.ProductID
	Address   PlaceOrderCommandAddress
}

type PlaceOrderCommandAddress struct {
	Name     string
	Street   string
	City     string
	PostCode string
	Country  string
}

type MarkOrderAsPaidCommand struct {
}

func NewOrdersService() {

}

func (od *OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {

}

func (od *OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}

func (od *OrdersService) GetOrdersByID(id orders.ID) (orders.Order, error) {
}
