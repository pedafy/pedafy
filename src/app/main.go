package main

import (
	"github.com/joho/godotenv"
	"github.com/pedafy/pedafy/src/server"
	"google.golang.org/appengine"
)

func main() {
	var s server.Server

	godotenv.Load("../../.env")

	s.RegisterHandlers()
	appengine.Main()
}
