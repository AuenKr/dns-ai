package main

import (
	"fmt"
	"os"

	"dns-server/handler"

	"github.com/joho/godotenv"
	"github.com/miekg/dns"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env file not found, using os env variable")
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
