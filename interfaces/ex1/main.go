package main

import "fmt"

type Instrumenter interface {
	Play()
}

type Ampli interface {
	Connected(amp string)
}

type Guitar struct {
	Strings int
}

type Piano struct {
	keys int
}

func (g Guitar) Connected(amp string) {
	fmt.Printf("Guitar connected on %v amplificator\n", amp)
}

func (p Piano) Play() {
	fmt.Printf("dododododododoododo with %v keys\n", p.keys)
}

func (g Guitar) Play() {

	fmt.Printf("tadadadadadadada using %v strings\n", g.Strings)
}

func main() {

	var instr Instrumenter
	instr = &Guitar{5}
	instr.Play()
	instr = &Piano{7}
	instr.Play()

	var con Ampli
	con = &Guitar{12}
	con.Connected("marshall")

}
