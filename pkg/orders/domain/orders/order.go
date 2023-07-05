package orders

import (
	"errors"

	products "github.com/suyash/mono-to-micro/pkg/shop/domain"
)

type ID string

var ErrEmptyOrderID error = errors.New("Empty Order ID")

type Order struct {
	id      ID
	product products.Product
	address Address
	paid    bool
}

func NewOrder(id ID, product *products.Product, address *Address) (*Order, error) {
	if len(id) == 0 {
		return nil, ErrEmptyOrderID
	}
	return &Order{
		id:      id,
		product: *product,
		address: *address,
		paid:    false,
	}, nil
}

func (o *Order) ID() ID {
	return o.id
}

func (o *Order) Product() *products.Product {
	return &o.product
}

func (o *Order) Address() *Address {
	return &o.address
}

func (o *Order) Paid() bool {
	return o.paid
}

func (o *Order) MarkAsPaid() {
	o.paid = true
}
