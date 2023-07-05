package products

import (
	products "github.com/suyash/mono-to-micro/pkg/shop/domain"
)

type MemoryRepository struct {
	products []products.Product
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		products: make([]products.Product, 0),
	}
}

func (m *MemoryRepository) Save(product *products.Product) error {
	for i, p := range m.products {
		if p.ID() == product.ID() {
			m.products[i] = *product
			return nil
		}
	}
	m.products = append(m.products, *product)
	return nil
}

func (m *MemoryRepository) GetProductByID(id products.ID) (*products.Product, error) {
	for _, p := range m.products {
		if id == p.ID() {
			return &p, nil
		}
	}
	return nil, products.ErrNotFound 
}

func (m *MemoryRepository) AllProducts() ([]products.Product, error) {
	return m.products, nil
}
