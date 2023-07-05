package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	common_http "github.com/suyash/mono-to-micro/pkg/common/http"
	"github.com/suyash/mono-to-micro/pkg/common/price"
	products "github.com/suyash/mono-to-micro/pkg/shop/domain"
)

type productsResource struct {
	readModel productReadModel
}

type productReadModel interface {
	AllProducts() ([]products.Product, error)
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

type ProductView struct {
	ID          products.ID `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	PriceView   PriceView   `json:"price"`
}

func priceViewFromPrice(p *price.Price) *PriceView {
	return &PriceView{
		Cents:    p.Cents(),
		Currency: p.Currency(),
	}
}

func AddRoutes(router *chi.Mux, productReadModel *productReadModel) {
	resource := productsResource{
		readModel: *productReadModel,
	}
	router.Get("/products", resource.GetAll)
}

func (p *productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.readModel.AllProducts()
	if err != nil {
		render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	var productsView []ProductView 
	for _,product := range products{
		price := product.Price()
		productsView = append(productsView, ProductView{
			ID: product.ID(),
			Name: product.Name(),
			Description: product.Description(),
			PriceView: *priceViewFromPrice(&price),
		})
	}
	resp,err := json.Marshal(productsView)
	if err != nil{
		render.Render(w,r,common_http.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type","application/json")
	w.Write(resp)
}
