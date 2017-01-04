package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"bufio"
	"os"
	"encoding/json"
	"strconv"
)

var dir string

type Result struct {
	Title string `json:"title"`
	Lines string `json:"lines,omitempty"`
	Error string `json:"error,omitempty"`
}

func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

func CountFile(book string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	path := dir + "/" + book

	f, err := os.Open(path)
	if err != nil {
		result, _ := json.Marshal(Result{book, "", "book not found"})
		io.WriteString(w, string(result))
		return
	}
	defer f.Close()
	lines, err := CountLines(f)
	if err != nil {
		result, _ := json.Marshal(Result{book, "", "can't read lines"})
		io.WriteString(w, string(result))
		return
	}

	result, _ := json.Marshal(Result{book, strconv.Itoa(lines), ""})
	io.WriteString(w, string(result))
}

func books(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := params["book"]
	CountFile(book, w)
}

func main() {
	dir = "./21_line_counting_http_service/books"

	router := mux.NewRouter()
	router.HandleFunc("/books/{book:[a-z]+.txt}", books).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
