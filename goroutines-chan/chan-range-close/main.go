package main

import (
	"fmt"
	"time"
)

func reader(c chan string) {
	fmt.Println("Start reading")
	defer fmt.Println("stop reading")

	for n := range c {
		fmt.Println(n)
	}
}

func main() {

	c := make(chan string)
	go reader(c)

	c <- "bob"
	c <- "test"
	close(c)
	time.Sleep(5 * time.Second)
}
