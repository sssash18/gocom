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
	repo products.Repository
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

func (resource *productsResource) Get(w http.ResponseWriter, r *http.Request) {
	product, err := resource.repo.GetProductByID(products.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	price := product.Price()
	productView := ProductView{
		ID:          product.ID(),
		Name:        product.Name(),
		Description: product.Description(),
		PriceView:   *priceViewFromPrice(&price),
	}
	resp, err := json.Marshal(productView)
	if err != nil {
		render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	w.Write(resp)
}

func AddRoutes(router *chi.Mux, repo *products.Repository) {
	resource := productsResource{
		repo: *repo,
	}
	router.Get("/products/{id}", resource.Get)
}
