package main

import "fmt"

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Width * r.Height
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect ==> %v ==> %v\n", r.Width, r.Height)
}

func (r Rect) Double() {
	r.Width *= 2
	r.Height *= 2
	fmt.Printf("Double = %v\n", r)
}

func main() {

	r := Rect{20, 10}
	fmt.Printf("Area = %v\n", r.Area())
	fmt.Println(r)
	r.Double()
	fmt.Println(r)
}
