package orders

import (
	"errors"

	"github.com/suyash/mono-to-micro/pkg/common/price"
)

var (
	ErrEmptyID   error = errors.New("Empty product ID")
	ErrEmptyName error = errors.New("Empty product Name")
)

type ProductID string

type Product struct {
	id          ProductID
	name        string
	description string
	price       price.Price
}

func (p *Product) ID() ProductID {
	return p.id
}

func (p *Product) Name() string{
	return p.name
}

func (p *Product) Description() string{
	return p.description
}

func (p *Product) Price() price.Price {
	return p.price
}

func NewProduct(id string, name string, description string, price price.Price) (*Product, error) {
	if len(id) == 0 {
		return nil, ErrEmptyID
	}
	if len(name) == 0 {
		return nil, ErrEmptyName
	}
	return &Product{
		id:          ProductID(id),
		name:        name,
		description: description,
		price:       price,
	}, nil
}
