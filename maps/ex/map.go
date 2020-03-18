package main

import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Printf("Map content %v, len is %v\n", m, len(m))

	//assign values:
	m["Something"] = 10
	m["something else"] = 23
	fmt.Printf("Map content %v, len is %v\n", m, len(m))

	fmt.Printf("key=Something, value = %v\n", m["Something"])

	i := m["something else"]
	fmt.Printf("i = %v\n", i)

	//test if a key exists:
	j, o := m["polo"]
	fmt.Printf("j = %v , ok= %v\n", j, o)

	m["something new"] = 233

	if _, ok := m["something new"]; ok == true {
		fmt.Println("Value exist !")
	}
	fmt.Printf("Map content %v, len is %v\n", m, len(m))

	//delete something
	delete(m, "something new")
	fmt.Printf("Map content %v, len is %v\n", m, len(m))

	//new map copy -- not a copy
	m2 := m
	fmt.Printf("Map content %v, len is %v\n", m2, len(m2))
	m2["copy"] = 12
	fmt.Printf("Map content %v, len is %v\n", m2, len(m2))
	fmt.Printf("Map content %v, len is %v\n", m, len(m))
}
