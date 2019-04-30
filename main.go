package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {

	http.HandleFunc("/", home)

	appengine.Main()
}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Coming soon ...")
}
