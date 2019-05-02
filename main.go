package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func main() {

	http.HandleFunc("/", home)

	appengine.Main()
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	var resp *http.Response
	var err error

	if appengine.IsDevAppServer() {
		resp, err = client.Get("http://localhost:9001")
	} else {
		resp, err = client.Get("https://api.pedafy.com/tig/v1/")
	}

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}
	fmt.Fprint(w, "<h1>Welcome, work in progress ...</h1><br>Testing APIs: "+string(responseBody))
}
