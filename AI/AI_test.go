package AI

import (
	"testing"

	"github.com/wwelden/TermWordle/Wordle"
)

//AI.go Tests

func TestGetContainedLetters(t *testing.T) {
	guess := "apple"
	word := "apply"
	containedLetters := GetContainedLetters(guess, word)
	if containedLetters != "appl" {
		t.Errorf("GetContainedLetters returned %s, expected appl", containedLetters)
	}
}

func TestGetAllWordsThatContain(t *testing.T) {
	guess := "apple"
	word := "apply"
	wordList := Wordle.ReadFile()
	matches := GetAllWordsThatContain(guess, word, wordList)
	if len(matches) != 197 {
		t.Errorf("GetAllWordsThatContain returned %d matches, expected 197", len(matches))
	}
}

func TestGetAllWordsThatMatch(t *testing.T) {
	guess := "apple"
	word := "apply"
	wordList := Wordle.ReadFile()
	matches := GetAllWordsThatMatch(guess, word, wordList)
	if len(matches) != 2 {
		t.Errorf("GetAllWordsThatMatch returned %d matches, expected 2", len(matches))
	}
}

func TestGetAllWordsWithOutLetters(t *testing.T) {
	guess := "apple"
	word := "apply"
	wordList := Wordle.ReadFile()
	matches := GetAllWordsWithOutLetters(guess, word, wordList)
	if len(matches) != 12593 {
		t.Errorf("GetAllWordsWithOutLetters returned %d matches, expected 12593", len(matches))
	}
}

func TestFindSubSet(t *testing.T) {
	matches := []string{"apple", "apply", "apples", "applesauce"}
	contained := []string{"apple", "apples"}
	subset := findSubSet(matches, contained)
	if len(subset) != 2 {
		t.Errorf("FindSubSet returned %d subset, expected 2", len(subset))
	}
}

func TestFirstGuess(t *testing.T) {
	wordList := Wordle.ReadFile()
	guess := FirstGuess(wordList)
	if len(guess) != 5 {
		t.Errorf("FirstGuess returned a guess with length %d, expected 5", len(guess))
	}
}

func TestCompete(t *testing.T) {
	word := "apple"
	wordList := Wordle.ReadFile()
	guess := Compete(word, wordList)
	if len(guess) != 5 {
		t.Errorf("Compete returned a guess with length %d, expected 5", len(guess))
	}
}
