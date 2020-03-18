package main

import "fmt"

func main() {

	var test string = "test"
	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -1
	nums = append(nums, 2)
	nums = append(nums, 123)
	//nums[4] = 224
	fmt.Printf("%v len(%d) cap(%d)\n", nums, len(nums), cap(nums))

	newt := []string{"g", "o", "l", "a", "n", "g"}
	newt = append(newt, "!")
	fmt.Printf("%v (len(%d)) (cat(%v))\n", newt, len(newt), cap(newt))

	sub1 := newt[0:2]
	fmt.Println(sub1)
}
