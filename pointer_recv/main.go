package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title      string
	Text       string
	plublished bool
}

//value receiver function

func (p Post) Overview() string {
	return fmt.Sprintf("%v - %v", p.Title, p.Text[:10])
}

func (p *Post) Published() bool {
	return p.plublished
}

func (p *Post) Publish() {
	p.plublished = true
}

func (p *Post) UnPublish() {
	p.plublished = false
}

func UpperTitle(p *Post) {
	p.Title = strings.ToUpper(p.Title)
}

func main() {

	p := Post{
		Title: "Test",
		Text:  `this is just a test`,
	}
	fmt.Printf("%v\n", p.Overview())
	//	fmt.Println(p)i
	fmt.Printf("Published ?: %v\n", p.plublished)
	p.Publish()
	fmt.Printf("Published ?: %v\n", p.plublished)

	//pointer de structure
	pointerStruct := &Post{
		Title: "Pointer",
		Text:  "Blablablablablabla",
	}
	UpperTitle(pointerStruct)
	fmt.Println(pointerStruct.Overview())
	fmt.Printf("%v\n", &pointerStruct)
	fmt.Printf("%v\n", *pointerStruct)
	pointerStruct.Publish()
	fmt.Printf("Published : %v\n", pointerStruct.plublished)
	fmt.Printf("%v\n", *pointerStruct)
}
