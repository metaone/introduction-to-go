package main

import (
	"fmt"
	"encoding/json"
	"log"
	"flag"
	"sync"
	"net/http"
	"bufio"
	"strings"
	"io"
)


func main() {
	chunk := flag.Int("chunk", 2, "Download chunk size")
	flag.Parse()

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

	channel := make(chan string, *chunk)
	var wait sync.WaitGroup
	for _, url := range urls {
		wait.Add(1)
		go func(url string) {
			defer wait.Done()
			channel <- downloadUrl(url)
		}(url)

		result := <- channel
		fmt.Print(result, "\n\n\n")
	}
	wait.Wait()
}

func downloadUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return readBody(resp.Body)
}

func readBody(body io.Reader) string {
	sc := bufio.NewScanner(body)
	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}

	return strings.Join(lines, "\n")
}
