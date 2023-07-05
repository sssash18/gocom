package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	common_http "github.com/suyash/mono-to-micro/pkg/common/http"
	"github.com/suyash/mono-to-micro/pkg/orders/application"
	"github.com/suyash/mono-to-micro/pkg/orders/domain/orders"
)

type ordersResourcre struct {
	service *application.OrdersService
	repo    *orders.Repository
}

type PostOrderRequest struct {
	ProductID orders.ProductID `json:"productID"`
	Address   PostOrderAddress `json:"address"`
}

type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"postCode"`
	Country  string `json:"country"`
}

type OrderPaidView struct{
	OrderID orders.ID `json:"orderID"`
	IsPaid bool `json:"isPaid"`
}

func AddRoutes(router *chi.Mux, service *application.OrdersService, repo *orders.Repository) {
	resource := ordersResourcre{
		service: service,
		repo:    repo,
	}
	router.Post("/order", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

func (o *ordersResourcre) Post(w http.ResponseWriter, r *http.Request) {
	var postOrder PostOrderRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}
	err = json.Unmarshal(body, &postOrder)
	if err != nil {
		render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	err = o.service.PlaceOrder(application.PlaceOrderCommand{
		OrderID:   orders.ID(uuid.New().String()),
		ProductID: postOrder.ProductID,
		Address: application.PlaceOrderCommandAddress{
			Name:     postOrder.Address.Name,
			Street:   postOrder.Address.Street,
			City:     postOrder.Address.City,
			PostCode: postOrder.Address.PostCode,
			Country:  postOrder.Address.Country,
		},
	})
	if err != nil {
		render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (o *ordersResourcre) GetPaid(w http.ResponseWriter, r *http.Request) {
	id := orders.ID(chi.URLParam(r,"id"))
	order,err := (*(o.repo)).ByID(id)
	if err != nil{
		render.Render(w,r,common_http.ErrInternal(err))
		return
	}
	order.MarkAsPaid()
	resp,err := json.Marshal(&OrderPaidView{
		OrderID: order.ID(),
		IsPaid: order.Paid(),
	})

	if err != nil {
		render.Render(w,r,common_http.ErrInternal(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
