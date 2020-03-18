package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func readFile(filename string) (string, error) {

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	if len(dat) == 0 {

		return "", errors.New("empty file")
	}

	return string(dat), nil

}

func main() {

	dat, err := readFile("test.txt")
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}
	fmt.Println(dat)

}
