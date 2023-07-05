package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting the shop microservice...")

	ctx := cmd.Context()

	r, closeFn := createShopMicroservice()

	server := &http.Server{
		Addr:    os.Getenv("SHOP_PRODUCT_SERVICE_BIND_ADDR"),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
		<-ctx.Done()
		log.Println("Closing shop microservice")
		if err := server.Close(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func createShopMicroservice() (router *chi.Mux, closeFn func()) {

}
