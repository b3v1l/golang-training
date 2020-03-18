package main

import (
	"fmt"
	"os"

	dictionnary "training.go/hangman/dictionary"
	"training.go/hangman/hangman"
)

func main() {

	dictionnary.Load("words.txt")
	w := dictionnary.PickWord()

	g := hangman.New(8, w)
	//fmt.Printf("%v\n", g)
	//fmt.Printf("%v\n", g.Letters)

	hangman.Welcome()
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lose":
			os.Exit(0)
		}
		s, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		guess = s
		g.MakeAGuess(guess)
		fmt.Println(s)
	}
}
