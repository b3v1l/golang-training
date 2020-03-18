package hangman

import "testing"

func TestLetterinWord(t *testing.T) {

	word := []string{"b", "o", "b"}
	guess := "b"
	hasletter := letterInWord(guess, word)
	if !hasletter {
		t.Errorf("word %s contains letter %s. got = %v", word, guess, hasletter)
	}

}

func TestLetterNotinWord(t *testing.T) {

	word := []string{"b", "b"}
	guess := "a"
	hasletter := letterInWord(guess, word)
	if hasletter {
		t.Errorf("word %s contains letter %s. got = %v", word, guess, hasletter)
	}

}

func TestInvalidWord(t *testing.T) {

	_, err := New(3, " ")
	if err == nil {
		t.Errorf("Word must not be empty, error must be returned")

	}

}
