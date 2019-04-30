package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

func main() {

	http.HandleFunc("/", hello)

	appengine.Main()
}

func hello(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	dictionary := make(map[string]string)
	dictionary["message"] = "This is the main service! Welcome."

	json, _ := json.Marshal(dictionary)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
