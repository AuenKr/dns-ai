package main

import (
	"dns-server/handler"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/miekg/dns"
)

var PORT string = ":3000"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading env")
	}

	fmt.Println("starting dns server at", PORT)
	server := &dns.Server{
		Addr:    PORT,
		Net:     "udp",
		Handler: &handler.DNSHandler{},
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server Fail to start", err)
	}
}
