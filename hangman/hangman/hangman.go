package hangman

import (
	"fmt"
	"strings"
)

type Game struct {
	State        string   // Game state
	Letters      []string // slice - word's letters to find
	FoundLetters []string // good guesses
	UsedLetters  []string // Ued letters
	TurnsLeft    int      // remaining attemps
}

func New(turns int, word string) (*Game, error) {

	if len(word) < 3 {
		return nil, fmt.Errorf("Word must be minimum 3 chars long")
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))

	for i := 0; i < len(letters); i++ {
		found[i] = "_ "
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		// empty slice
		UsedLetters: []string{},
		TurnsLeft:   turns,
	}
	return g, nil
}
func (g *Game) MakeAGuess(guess string) {

	guess = strings.ToUpper(guess)

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "GoodGuess"
		g.RevealLetters(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}

	} else {
		g.State = "Bad Guess"
		g.LoseTurn(guess)
		if g.TurnsLeft == 0 {
			g.State = "lose"
			//fmt.Println("Game Over")
			//os.Exit(0)
		}
		//fmt.Printf("Turn left = %v", g.TurnsLeft)

	}
	//g.TurnsLeft -= 1
	//fmt.Printf("Turns = %v\n", g.TurnsLeft)
}

func (g *Game) RevealLetters(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)

	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}

	}
}

func (g *Game) LoseTurn(guess string) {

	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

func hasWon(letters []string, FoundLetters []string) bool {
	for i := range letters {
		if letters[i] != FoundLetters[i] {
			return false
		}
	}
	return true
}

func letterInWord(guess string, letters []string) bool {
	for _, v := range letters {
		if v == guess {
			return true
		}
	}
	return false
}
