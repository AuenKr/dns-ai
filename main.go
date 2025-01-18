package main

import (
	"dns-server/handler"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/miekg/dns"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading env")
	}

	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		PORT = "3000"
	}

	fmt.Println("starting dns server at", PORT)
	server := &dns.Server{
		Addr:    ":" + PORT,
		Net:     "udp",
		Handler: &handler.DNSHandler{},
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server Fail to start", err)
	}
}
