package main

import (
	"fmt"

	"training.go/struct/player"
)

func main() {

	var p player.Player
	p.Age = 12
	p.Name = "polo"
	a := player.Avatar{"http://whatver.test"}

	p2 := player.Player{
		Name: "Test",
		Age:  11,
		Avatar: player.Avatar{
			Url: "http://new_Avatar.com",
		},
	}

	p4 := player.New("Perin")

	fmt.Printf("%v\n", p)
	fmt.Printf("Player avatar: %v\n", a)
	fmt.Printf("player 2 =%v\n", p2)
	fmt.Printf("player 4 =%v\n", p4)

}
