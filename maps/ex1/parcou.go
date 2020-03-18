package main

import "fmt"

func main() {

	m := map[string]int{
		"polo":    10,
		"polette": 15,
		"perin":   45,
		"test":    2,
	}

	for i, _ := range m {
		fmt.Printf("keys = %v\n", i)
	}

	c := 1
	for i := range m {
		fmt.Printf("keys = %v\n", i)
		m[i] = c
		c++
	}
	fmt.Println(m)
}
