package cmd

import (
	"log"
	"net"
	"time"
)

func WaitForService(host string){
	log.Printf("Waiting for the %s",host)
	for {
		log.Printf("Testing the connection to %s",host)
		conn,err := net.Dial("tcp", host)
		if err == nil {
			_ = conn.Close()
			log.Printf("%s is up!",host)
			return
		}
		time.Sleep(time.Millisecond*5)
	}
}