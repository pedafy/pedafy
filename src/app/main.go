package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pedafy/pedafy/src/server"
)

func main() {
	var s server.Server

	godotenv.Load("../../.env")

	s.Start()
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
