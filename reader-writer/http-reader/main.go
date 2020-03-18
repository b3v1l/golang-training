package main

import (
	"fmt"
	"net/http"
)

func main() {

	req, err := http.Get("https://golang.org")

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("Test %v\n", req)
}
