package main

import (
	"log"
	"os"
)

func main(){
	log.Println("Starting Payments microservice...")
	ctx := cmd.Context()

	paymentsInterface := createPaymentsInterface()

	if err := paymentsInterface.Run(ctx); err != nil{
		panic(err)
	}

}

func createPaymentsInterface(){
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	
}