package Wordle

import (
	"slices"
	"testing"
)

//Wordle.go Tests

func TestReadFile(t *testing.T) {
	wordList := ReadFile()
	if len(wordList) == 0 {
		t.Errorf("ReadFile returned an empty list")
	}
	if len(wordList) != 14855 {
		t.Errorf("ReadFile returned %d words, expected 10000", len(wordList))
	}
}
func TestGetWord(t *testing.T) {
	word := GetWord()
	if len(word) != 5 {
		t.Errorf("GetWord returned a word with length %d, expected 5", len(word))
	}
	wordList := ReadFile()
	if !slices.Contains(wordList, word) {
		t.Errorf("GetWord returned a word that is not in the word list")
	}
}

func TestCheckLetter(t *testing.T) {
	guess := "apple"
	word := "apply"
	letterDex := 0
	result := CheckLetter(guess, word, letterDex)
	if !result {
		t.Errorf("CheckLetter returned false, expected true")
	}
}

func TestCheckGuess(t *testing.T) {
	guess := "apple"
	word := "apply"
	results := CheckGuess(guess, word)
	if len(results) != 5 {
		t.Errorf("CheckGuess returned %d results, expected 5", len(results))
	}
}

func TestContainsLetter(t *testing.T) {
	guess := "apple"
	word := "apply"
	letterDex := 0
	result := ContainsLetter(guess, word, letterDex)
	if !result {
		t.Errorf("ContainsLetter returned false, expected true")
	}
}

func TestGuessContains(t *testing.T) {
	guess := "apple"
	word := "apply"
	results := GuessContains(guess, word)
	if len(results) != 5 {
		t.Errorf("GuessContains returned %d results, expected 5", len(results))
	}
}

func TestDoesGuessMatch(t *testing.T) {
	guess := "apply"
	word := "apply"
	result := DoesGuessMatch(guess, word)
	if !result {
		t.Errorf("DoesGuessMatch returned false, expected true")
	}
}

func TestTrackLetterUsage(t *testing.T) {
	word := "apple"
	letterCounts := TrackLetterUsage(word)
	if len(letterCounts) != 4 {
		t.Errorf("TrackLetterUsage returned %d letter counts, expected 4", len(letterCounts))
	}
}

func TestCheckGuessWithDuplicates(t *testing.T) {
	guess := "apple"
	word := "apply"
	results := CheckGuessWithDuplicates(guess, word)
	if len(results) != 5 {
		t.Errorf("CheckGuessWithDuplicates returned %d results, expected 5", len(results))
	}
}
