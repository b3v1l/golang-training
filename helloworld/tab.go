package main

import "fmt"

func main() {

	var names [3]string
	//names :=  [3] string
	fmt.Printf("names=%v (len=%v)\n", names, len(names))

	names[0] = "polo"
	names[1] = "test1"
	names[2] = "test2"
	//	names[4] = "test4"
	fmt.Printf("names=%v (len=%v)\n", names, len(names))

}
