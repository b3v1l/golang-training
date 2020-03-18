package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	color string
}

type Cat struct {
	jumpheight int
}

func (d Dog) Speak() string {
	return "woof"

}

func (c Cat) Speak() string {
	return "Miawww"
}

func main() {

	animals := []Animal{
		Dog{"fire"},
		Cat{12},
	}

	for _, v := range animals {
		fmt.Println(v.Speak())
		fmt.Printf("Type of animal %T\n", v)

		//		t, ok := v.(Dog)
		//fmt.Printf("Type assertion value=%v, ok=%v\n", t, ok)
		//if ok {
		if t, ok := v.(Dog); ok {
			fmt.Printf("This a dog , color is %v\n", t.color)

		} else {
			fmt.Println("Not a dog !")
		}

	}

	fmt.Println("-------------------")
	for _, a := range animals {
		describeAnimal(a)
	}

}

func describeAnimal(a Animal) {

	switch v := a.(type) {
	case Dog:
		fmt.Printf("This is a Dooooooooog ! Color= %v\n", v.color)

	case Cat:
		fmt.Printf("This is a cat ! jump= %v\n", v.jumpheight)

	}

}
