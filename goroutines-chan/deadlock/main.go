package main

import "time"

func hello(c chan string) {

	c <- "hellooooooo"
}

func main() {

	c := make(chan string)
	go hello(c)
	c <- "block the chan with main!"
	time.Sleep(10 * time.Second)

}
