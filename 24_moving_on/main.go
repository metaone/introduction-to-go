package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)


func main() {
	input := []byte(`[
		"https://raw.githubusercontent.com/metaone/introduction-to-go/master/01_short_declaration_syntax/main.go",
		"https://raw.githubusercontent.com/metaone/introduction-to-go/master/02_slice_exercises/main.go",
		"https://raw.githubusercontent.com/metaone/introduction-to-go/master/03_slice_initialisation/main.go",
		"https://raw.githubusercontent.com/metaone/introduction-to-go/master/04_subslices/main.go",
		"https://raw.githubusercontent.com/metaone/introduction-to-go/master/05_inserting_values_into_a_map/main.go"
	]`)

	urls := []string{}
	err := json.Unmarshal(input, &urls)
	if err != nil {
		log.Fatal(err)
	}

	chain := make(chan string)
	for _, url := range urls {
		go func(url string) {
			fmt.Print("start" + url + "\n")
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Print("end" + url + "\n")
			chain <- string(body)
		}(url)
	}

	fmt.Print(<- chain)
}
