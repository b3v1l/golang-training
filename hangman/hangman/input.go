package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (get string, err error) {
	valid := false
	for !valid {
		fmt.Println("What's your guess ? ")
		get, err = reader.ReadString('\n')
		if err != nil {
			return get, err

		}
		//remove space
		get = strings.TrimSpace(get)
		if len(get) != 1 {
			fmt.Println("One character at one! Cheater !")
			fmt.Printf("Error = %v -- len = %v\n", get, len(get))
			continue

		}
		valid = true
	}
	return get, nil
}
