package main

import (
	"fmt"
	"net/http/httptest"
	"net/http"
	"time"
	"math/rand"
	"log"
	"io/ioutil"
)


func main() {
	server := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		sleep := rand.Intn(10)
		time.Sleep(time.Duration(sleep) * time.Second)
		fmt.Fprintln(w, "Done!")
	}))
	defer server.Close()

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(server.URL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
