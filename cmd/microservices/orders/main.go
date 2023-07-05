package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting the Orders microservice...")
	r, closeFn := createOrderMicroservice()
	defer closeFn()

	ctx := cmd.Context()
	server := http.Server{
		Addr:    os.Getenv("SHOP_ORDER_SERVICE_BINDING_ADDRESS"),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}

		<-ctx.Done()
		log.Println("Closing Order Microservice")

		if err := server.Close(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func createOrderMicroservice() (router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	shopHTTPClient := order_infra_client.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))
	r := cmd.NewRouter()
	return r
}
