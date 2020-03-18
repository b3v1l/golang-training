package main

import "fmt"

func sumFun(x, y, z int) (sum int) {

	sum = x + y + z
	return sum
}

func noParams() {

	fmt.Printf("name = %s, age = %d, email = %s\n", "polo", 10, "polo@golang.com")
}

func main() {
	noParams()
	sum := sumFun(12, 34, 45)
	fmt.Printf("Addition result = %d\n", sum)
}
