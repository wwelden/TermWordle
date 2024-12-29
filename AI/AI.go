package AI

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/wwelden/TermWordle/Wordle"
)

func GetContainedLetters(guess string, word string) string {
	contains := Wordle.GuessContains(guess, word)
	containedLetters := ""
	for i := range contains {
		if contains[i] {
			containedLetters += string(guess[i])
		}
	}
	return containedLetters
}

func GetAllWordsThatContain(guess string, word string, wordList []string) []string {
	contains := GetContainedLetters(guess, word)
	matches := []string{}

	for _, testWord := range wordList {
		hasAllLetters := true
		for _, letter := range contains {
			if !strings.Contains(testWord, string(letter)) {
				hasAllLetters = false
				break
			}
		}
		if hasAllLetters {
			matches = append(matches, testWord)
		}
	}

	return matches
}
func GetAllWordsThatMatch(guess string, word string, wordList []string) []string {
	letters := Wordle.CheckGuess(guess, word)
	matches := []string{}

	for _, testWord := range wordList {
		isMatch := true
		for i := range letters {
			if letters[i] {
				if testWord[i] != guess[i] {
					isMatch = false
					break
				}
			}
		}
		if isMatch {
			matches = append(matches, testWord)
		}
	}
	return matches
}

func GetAllWordsWithOutLetters(guess string, word string, wordList []string) []string {
	letters := Wordle.CheckGuess(guess, word)
	matches := []string{}

	for _, testWord := range wordList {
		isMatch := true
		for i := range letters {
			if letters[i] {
				if testWord[i] == guess[i] {
					isMatch = false
					break
				}
			}
		}
		if isMatch {
			matches = append(matches, testWord)
		}
	}
	return matches
}

func findSubSet(matches []string, contained []string) []string {
	subset := []string{}
	for _, match := range matches {
		for _, contain := range contained {
			if match == contain {
				subset = append(subset, match)
				break
			}
		}
	}
	return subset
}

func FirstGuess(wordList []string) string {
	length := len(wordList)
	randomIndex := rand.Intn(length)
	return wordList[randomIndex]
}

func Compete(word string, wordList []string) string {
	guess := FirstGuess(wordList)
	matches := GetAllWordsThatMatch(guess, word, wordList)
	contained := GetAllWordsThatContain(guess, word, wordList)
	subset := findSubSet(matches, contained)
	return FirstGuess(subset)
}

func ShowResults2(guess string, word string) {
	rightSpots := Wordle.CheckGuess(guess, word)
	rightLetters := Wordle.GuessContains(guess, word)
	for i := range rightSpots {
		if rightSpots[i] {
			fmt.Print(Wordle.Green + strings.ToUpper(string(guess[i])) + Wordle.White)
		} else if rightLetters[i] {
			fmt.Print(Wordle.Yellow + strings.ToUpper(string(guess[i])) + Wordle.White)
		} else {
			fmt.Print(Wordle.White + strings.ToUpper(string(guess[i])) + Wordle.White)
		}
	}
	fmt.Println()
}

var AIGuessNum = 0
var AIGotIt = false

func CompeteLoop(word string, wordList []string) {
	for i := 0; i < 6; i++ {
		guess := Compete(word, wordList)
		// ShowResults2(guess, word)
		Wordle.ShowResultsEnhancedAI(guess, word)
		AIGuessNum++
		if Wordle.DoesGuessMatch(guess, word) {
			fmt.Println("AI Got It!")
			AIGotIt = true
			break
		}
		matches := GetAllWordsThatMatch(guess, word, wordList)
		contained := GetAllWordsThatContain(guess, word, wordList)
		wordList = findSubSet(matches, contained)
	}
}

func AIPlay(word string) {
	// fmt.Println(" __        __            _ _      ")
	// fmt.Println(" \\ \\      / /__  _ __ __| | | ___ ")
	// fmt.Println("  \\ \\ /\\ / / _ \\| '__/ _` | |/ _ \\")
	// fmt.Println("   \\ V  V / (_) | | | (_| | |  __/")
	// fmt.Println("    \\_/\\_/ \\___/|_|  \\__,_|_|\\___|")
	// fmt.Println("-----------------------------------")
	fmt.Println("AI's turn")
	wordList := Wordle.ReadFile()
	CompeteLoop(word, wordList)
}