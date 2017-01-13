package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

var reverseHost string = "http://www.google.com:80"

func main() {
	http.HandleFunc("/", proxyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		resp, err := http.Get(reverseHost + r.RequestURI)
		if err != nil {
			log.Panic(err)
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Panic(err)
		}
		resp.Body.Close()

		w.Write(data)
	default:
		w.Write([]byte(r.Method + " is not implemented yet!"))
	}
}
