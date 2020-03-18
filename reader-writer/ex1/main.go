package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	r, err := os.Open("test.txt")

	//r := strings.NewReader("This is a reader test !")
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Printf("Error in reader: %v\n", err)
		return
	}

	fmt.Println(string(buf))
}
