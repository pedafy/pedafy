package main

import (
	"github.com/pedafy/pedafy/src/server"
	"google.golang.org/appengine"
)

func main() {
	var s server.Server

	s.RegisterHandlers()
	appengine.Main()
}
