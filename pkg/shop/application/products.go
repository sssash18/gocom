package application

import (
	"github.com/suyash/mono-to-micro/pkg/common/price"
	products "github.com/suyash/mono-to-micro/pkg/shop/domain"
)

type productReadModel interface {
	AllProducts() ([]products.Product, error)
}

type ProductsService struct {
	repo      products.Repository
	readModel productReadModel
}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCents    int
	PriceCurrency string
}

func  NewProductsService(repo *products.Repository, readModel *productReadModel) *ProductsService {
	return &ProductsService{
		repo:      *repo,
		readModel: *readModel,
	}
}

func (s *ProductsService) AllProducts() ([]products.Product, error) {
	return s.readModel.AllProducts()
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {
	price, err := price.NewPrice((uint)(cmd.PriceCents), cmd.PriceCurrency)
	if err != nil {
		return err
	}
	product, err := products.NewProduct(cmd.ID, cmd.Name, cmd.Description, *price)
	if err != nil {
		return err
	}
	err = s.repo.Save(product)
	if err != nil {
		return err
	}
	return nil
}
