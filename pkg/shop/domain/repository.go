package products

import "errors"

var ErrNotFound error = errors.New("Product Not Found")

type Repository interface{
	Save(*Product) error
	GetProductByID(ID)(*Product,error)
}