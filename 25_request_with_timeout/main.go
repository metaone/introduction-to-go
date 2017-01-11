package main

import (
	"fmt"
	"net/http"
	"time"
	"net/http/httptest"
	"math/rand"
	"log"
	"io/ioutil"
)


func main() {
	maxSleep := 10

	server := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		sleep := rand.Intn(maxSleep)
		time.Sleep(time.Duration(sleep) * time.Second)
		fmt.Fprintf(w, "Done after %v!\n", sleep)
	}))
	defer server.Close()

	chain := make(chan string)
	for range [5]int{} {
		go func() {
			client := http.Client{Timeout: time.Duration(maxSleep) * time.Second}
			resp, err := client.Get(server.URL)
			defer resp.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			chain <- string(body)
		}()
	}

	for {
		select {
		case <- time.After(time.Duration(5) * time.Second):
			fmt.Print("stop after 5 seconds!")
			return
		case body := <- chain:
			fmt.Print(body)
		}
	}
}
