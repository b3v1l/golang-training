package main

import "fmt"

func updateVal(val string) {
	val = "something"
}

func updatePtr(ptr *string) {
	*ptr = "value"
}

func main() {

	var i int = 1
	var p *int = &i

	fmt.Printf("i = %v\n", i)
	fmt.Printf("p = %v\n", p)
	fmt.Printf("*p = %v\n", *p)

	var s string = "polo"
	var sPtr *string = &s
	s2 := *sPtr

	fmt.Printf("s= %v\n", s)
	fmt.Printf("sPtr= %v\n", sPtr)
	fmt.Printf("*sPtr= %v\n", *sPtr)
	fmt.Printf("s2= %v\n", s2)

	//deferencement:

	*sPtr = "Perin"
	fmt.Println("Deferencement and update\n")
	fmt.Printf("s= %v\n", s)
	fmt.Printf("sPtr= %v\n", sPtr)
	fmt.Printf("*sPtr= %v\n", *sPtr)
	fmt.Printf("s2= %v\n", s2)

	updatePtr(&s)

	fmt.Println("Update")
	fmt.Printf("s= %v\n", s)
	fmt.Printf("sPtr= %v\n", sPtr)
	fmt.Printf("*sPtr= %v\n", *sPtr)
	fmt.Printf("s2= %v\n", s2)

}
