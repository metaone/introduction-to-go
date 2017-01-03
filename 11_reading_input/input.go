package main

import (
	"bufio"
	"io"
	"log"
	"strings"
	"fmt"
	"os"
)

// ReadAll reads all the lines of text from r and returns
// all the data read as a string
func ReadAll(r io.Reader) string {
	sc := bufio.NewScanner(r)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}
	return strings.Join(lines, "\n")
}

func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/11_reading_input/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(ReadAll(file))
}
