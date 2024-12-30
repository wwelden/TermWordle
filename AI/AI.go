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

func WordHasLetter(word string, letter string) bool {
	return strings.Contains(word, letter)
}
func GetLettersNotInWord(word string, guess string) string {
	letters := ""
	for _, letter := range guess {
		if !strings.Contains(word, string(letter)) {
			letters += string(letter)
		}
	}
	return letters
}

func GetAllWordsWithOutLetters(guess string, word string, wordList []string) []string {
	letters := GetLettersNotInWord(word, guess)
	matches := []string{}

	for _, testWord := range wordList {
		if !WordHasLetter(testWord, letters) {
			matches = append(matches, testWord)
		}
	}
	return matches
}

func GetLettersInWord(guess string, word string) string {
	letters := ""
	used := make(map[rune]bool) // Track letters already added
	for _, letter := range guess {
		if strings.Contains(word, string(letter)) && !used[letter] {
			letters += string(letter)
			used[letter] = true // Mark letter as used
		}
	}
	return letters
}

func GetYellowLetters(guess string, word string) string {
	letters := ""
	used := make(map[rune]bool) // Track letters already added

	for _, letter := range guess {
		// If the letter is in the word and not already added to the result
		if strings.Contains(word, string(letter)) {
			letters += string(letter)
			used[letter] = true
		}
	}
	return letters
}

func GetAllWordsWithYellowLetters(guess string, word string, wordList []string) []string {
	yellowLetters := GetYellowLetters(guess, word) // Get all yellow letters
	matches := []string{}

	for _, testWord := range wordList {
		isValid := true
		for _, letter := range yellowLetters {
			if !strings.Contains(testWord, string(letter)) {
				isValid = false
				break
			}
		}
		// Ensure the word contains all yellow letters
		if isValid {
			matches = append(matches, testWord)
		}
	}
	return matches
}

func GetAllCorrectLetters(guess string, word string) string {
	letters := ""
	for i := range guess {
		if guess[i] == word[i] {
			letters += string(guess[i])
		}
	}
	return letters
}

func GetAllWordsWithGreenLetters(guess string, word string, wordList []string) []string {
	letters := GetAllCorrectLetters(guess, word)
	matches := []string{}

	for _, testWord := range wordList {
		if WordHasLetter(testWord, letters) {
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

//get all the word with letters in the right places
//get all the words that contain the letters in the wrong places
//get all the words that do not contains the letters the word does not contain
//find all of the words that are in all 3 lists

func getRightPlaces(guess string, answer string) string {
	ret := ""
	for i, letter := range answer {
		if letter == rune(guess[i]) {
			ret += string(letter)
		} else {
			ret += "-"
		}
	}
	return ret
}

func rightPlaces(guess string, answer string, wordlist []string) []string {
	results := []string{}
	greenPos := getRightPlaces(guess, answer)
	for _, testWord := range wordlist {
		testGreen := getRightPlaces(testWord, answer)
		for i := 0; i < 5; i++ {
			if testGreen[i] == greenPos[i] || greenPos[i] == '-' {
				results = append(results, testWord)
			}
		}
	}
	return results
}

func getWrongPlaces(guess string, answer string) string {
	ret := ""
	for _, letter := range guess {
		if strings.Contains(answer, string(letter)) {
			ret += string(letter)

		}
	}
	return ret
}
func hasLetters(input string, word string) bool {
	for _, letter := range input {
		if !strings.Contains(word, string(letter)) {
			return false
		}
	}
	return true
}

func wrongPlaces(guess string, answer string, wordlist []string) []string {
	results := []string{}
	mustContain := getWrongPlaces(guess, answer)
	for _, testWord := range wordlist {
		if hasLetters(mustContain, testWord) {
			results = append(results, testWord)
		}
	}
	return results
}
func notHaveLetters(guess string, answer string) string {
	ret := ""
	for _, letter := range guess {
		if !strings.Contains(answer, string(letter)) {
			ret += string(letter)
		}
	}
	return ret
}

func getWordsWithoutLetters(guess string, answer string, wordlist []string) []string {
	results := []string{}
	input := notHaveLetters(guess, answer)
	for _, word := range wordlist {
		if !hasLetters(input, word) {
			results = append(results, word)
		}
	}
	return results
}

func subSetEnhanced(matches []string, contained []string, wordWithOut []string) []string {
	subset := []string{}
	for _, match := range matches {
		inContained := false
		inWithout := false
		for _, contain := range contained {
			if match == contain {
				inContained = true
				break
			}
		}
		for _, without := range wordWithOut {
			if match == without {
				inWithout = true
				break
			}
		}
		if inContained && inWithout {
			subset = append(subset, match)
		}
	}
	return subset
}

func FirstGuess(wordList []string) string {
	length := len(wordList)
	if length == 0 {
		return "" // Return an empty string if the word list is empty
	}
	randomIndex := rand.Intn(length)
	return wordList[randomIndex]
}

func Compete(word string, wordList []string) string {
	guess := FirstGuess(wordList)
	matches := GetAllWordsThatMatch(guess, word, wordList)
	contained := GetAllWordsThatContain(guess, word, wordList)
	subset := findSubSet(matches, contained)
	// wordWithOut := GetAllWordsWithOutLetters(guess, word, wordList)
	// subset2 := findSubSet2(matches, contained, wordWithOut)
	return FirstGuess(subset)
}
func CompeteEnhanced(word string, wordList []string, firstGuessFunc func([]string) string) string {
	guess := firstGuessFunc(wordList) // Use the injected function
	greenMatches := rightPlaces(guess, word, wordList)
	yellowMatches := wrongPlaces(guess, word, wordList)
	wordWithOut := getWordsWithoutLetters(guess, word, wordList)
	subset := subSetEnhanced(greenMatches, yellowMatches, wordWithOut)
	return firstGuessFunc(subset)
}

func ShowResults2(guess string, word string) {
	rightSpots := Wordle.CheckGuess(guess, word)
	rightLetters := Wordle.GuessContains(guess, word)
	for i := range rightSpots {
		if rightSpots[i] {
			fmt.Print(Wordle.Green + "\033[30m" + strings.ToUpper(string(guess[i])) + Wordle.White)
		} else if rightLetters[i] {
			fmt.Print(Wordle.Yellow + "\033[30m" + strings.ToUpper(string(guess[i])) + Wordle.White)
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
		Wordle.ShowResultsEnhancedAI(guess, word)
		AIGuessNum++
		if Wordle.DoesGuessMatch(guess, word) {
			fmt.Println("AI Got It!")
			AIGotIt = true
			break
		}
		matches := GetAllWordsThatMatch(guess, word, wordList)
		contained := GetAllWordsThatContain(guess, word, wordList)
		// wordWithOut := GetAllWordsWithOutLetters(guess, word, wordList)
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
