package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("https://golang.org")

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	defer res.Body.Close()

	//r, _ := ioutil.ReadAll(res.Body)
	//fmt.Printf("Response =  %v\n", res)
	//fmt.Println(string(r))

	f, _ := os.Create("test1.html")
	defer f.Close()

	io.Copy(f, res.Body)

}
