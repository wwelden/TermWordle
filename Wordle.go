package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	Green  = "\033[0;102m"
	Yellow = "\033[0;103m"
	Red    = "\033[0;101m"
	White  = "\033[0m"
)

func ReadFile() []string {
	inputFile, err := os.ReadFile("WordList.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(inputFile), "\n")
}

func GetWord() string {
	words := ReadFile()
	length := len(words)
	randomIndex := rand.Intn(length)
	return words[randomIndex]
}

func GetGuess() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	guess := strings.ToLower(strings.TrimSpace(text))
	if len(guess) != 5 {
		fmt.Println("Please enter a 5 letter word")
		return GetGuess()
	}

	wordList := ReadFile()
	isValidWord := false
	for _, word := range wordList {
		if strings.TrimSpace(word) == guess {
			isValidWord = true
			break
		}
	}

	if !isValidWord {
		fmt.Println("Not a valid word")
		return GetGuess()
	}

	return guess
}

func CheckLetter(guess string, word string, letterDex int) bool {
	return guess[letterDex] == word[letterDex]
}

func CheckGuess(guess string, word string) []bool {
	results := []bool{}
	for i := range guess {
		results = append(results, CheckLetter(guess, word, i))
	}
	return results
}

func ContainsLetter(guess string, word string, letterDex int) bool {
	return strings.Contains(word, string(guess[letterDex]))
}

func GuessContains(guess string, word string) []bool {
	results := []bool{}
	for i := range guess {
		results = append(results, ContainsLetter(guess, word, i))
	}
	return results
}

func DoesGuessMatch(guess string, word string) bool {
	return guess == word
}

func ShowResults(guess string, word string) {
	rightSpots := CheckGuess(guess, word)
	rightLetters := GuessContains(guess, word)
	// Move cursor up one line to overwrite the input
	fmt.Print("\033[1A\r")
	// Clear the current line
	fmt.Print("\033[K")
	for i := range rightSpots {
		if rightSpots[i] {
			fmt.Printf(Green + strings.ToUpper(string(guess[i])) + White)
			// emoji += "🟩"
		} else if rightLetters[i] {
			fmt.Printf(Yellow + strings.ToUpper(string(guess[i])) + White)
			// emoji += "🟨"
		} else {
			fmt.Printf(White + strings.ToUpper(string(guess[i])) + White)
			// emoji += "⬜"
		}
	}
	fmt.Println()
	// fmt.Println(emoji)
}
