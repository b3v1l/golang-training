package hangman

import "fmt"

func Welcome() {

	fmt.Println(`

██╗  ██╗ █████╗ ███╗   ███╗ ██████╗  ███╗   ███╗ █████╗ ███╗    ██╗
██║  ██║██╔══██╗████╗ ████║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║
███████║███████║██╔████╔██║██║  ███╗██╔████╔██║███████║██╔██╗██║
██╔══██ ██╔══██║██║╚██╔╝██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║ 
██║  ██║██║  ██║██║ ╚═╝ ██║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝ ╚═════╝  ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)

}

func drawTurns(i int) {

	var draw string
	switch i {
	case 0:
		draw = `
		
    ____
   |    |
   |    o
   |   /|\
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
`
	case 1:
		draw = `
			


    ____
   |    |
   |    o
   |   /|\
   |    |
   |
  _|_
 |   |______
 |          |
 |__________|
`
	case 2:
		draw = `

    ____
   |    |
   |    o
   |    |
   |    |
   |
  _|_
 |   |______
 |          |
 |__________|
 `
	case 3:
		draw = `
    ____
   |    |
   |    o
   |
   |
   |
  _|_
 |   |______
 |          |
 |__________|
`
	case 4:
		draw = `
   ____
  |    |
  |
  |
  |
  |
 _|_
|   |______
|          |
|__________|
`
	case 5:
		draw = `
    ____
   |
   |
   |
   |
   |
  _|_
 |   |______
 |          |
 |__________|
`
	case 6:
		draw = `

   |
   |
   |
   |
   |
  _|_
 |   |______
 |          |
 |__________|
`

	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
`

	case 8:
		draw = ``

	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "Good guess":
		fmt.Print("Good guess")
	case "alreadyGuessed":
		fmt.Printf("Letter %v already used ...", guess)

	case "Bad Guess":
		fmt.Println("Bad guess ... :(")

	case "lose":
		fmt.Print("You lose ! Word was ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You Win, the word was :")
		drawLetters(g.Letters)

	}

}

func drawLetters(l []string) {
	for _, val := range l {
		fmt.Printf(" %v", val)
	}
	fmt.Println()
}
