package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Cooker interface {
	Cook()
}

type Chef struct {
	Person
	Stars int
}

func (c Chef) Cook() {
	fmt.Printf("i'm a Cook and i've got %v stars \n", c.Stars)
}

func processPerson(i interface{}) {
	switch p := i.(type) {
	case Person:
		fmt.Printf("We have a Person %v\n", p)
	case Chef:
		fmt.Printf("The person is a Chef : %v\n", p)
		p.Cook()
	}
}

func main() {

	var x interface{} = Person{"Bob", 45}
	fmt.Printf("x type=%T, data=%v\n", x, x)

	s, ok := x.(string)
	fmt.Printf("Person as a string  ok ? =%v value='%v'\n", ok, s)

	processPerson(x)

	x = Chef{
		Stars: 4,
		Person: Person{
			Name: "polo",
			Age:  32,
		},
	}
	processPerson(x)
}
