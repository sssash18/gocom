package application

import (
	"errors"

	"github.com/google/uuid"
	"github.com/suyash/mono-to-micro/pkg/common/price"
	"github.com/suyash/mono-to-micro/pkg/orders/domain/orders"
)

type productsService interface {
	ProductsByID(id orders.ProductID) (*orders.Product, error)
}

type paymentsService interface {
	InitializeOrderPayment(id orders.ID, price price.Price) error
}

type OrdersService struct {
	productsService productsService
	paymentsService paymentsService
	repo            orders.Repository
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
	OrderID orders.ID
}

func NewOrdersService(productsService productsService, paymentsService paymentsService, repo orders.Repository) *OrdersService {
	return &OrdersService{
		productsService: productsService,
		paymentsService: paymentsService,
		repo:            repo,
	}
}

func (od *OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {
	address, err := orders.NewAddress(cmd.Address.Name, cmd.Address.Street, cmd.Address.City, cmd.Address.PostCode, cmd.Address.Country)
	if err != nil {
		return err
	}

	product, err := od.productsService.ProductsByID(cmd.ProductID)
	if err != nil {
		return err
	}

	order, err := orders.NewOrder(orders.ID(uuid.New().String()), product, address)
	if err != nil {
		return err
	}

	err = od.repo.Save(order)
	if err != nil {
		return err
	}

	err = od.paymentsService.InitializeOrderPayment(order.ID(), order.Product().Price())
	if err != nil {
		return err
	}

	return nil
}

func (od *OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {
	o, err := od.repo.ByID(cmd.OrderID)
	if err != nil {
		return err
	}

	o.MarkAsPaid()

	err = od.repo.Save(o)
	if err != nil {
		return err
	}

	return nil

}

func (od *OrdersService) GetOrdersByID(id orders.ID) (*orders.Order, error) {
	o, err := od.repo.ByID(id)
	if err != nil {
		return nil, errors.New("cannot get order" + err.Error())
	}
	return o, nil
}
