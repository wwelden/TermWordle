package Wordle

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
)

var (
	Green     = "\033[0;102m"
	Yellow    = "\033[0;103m"
	Red       = "\033[0;101m"
	White     = "\033[0m"
	BlackText = "\033[30m"
)

//go:embed embedded/WordList.txt
var wordListData string

//go:embed embedded/AnswerList.txt
var answerListData string

var (
	guessListOnce  sync.Once
	guessList      []string
	answerListOnce sync.Once
	answerList     []string
)

func parseList(data string) []string {
	raw := strings.Split(data, "\n")
	out := make([]string, 0, len(raw))
	for _, w := range raw {
		w = strings.TrimSpace(w)
		if w != "" {
			out = append(out, w)
		}
	}
	return out
}

// LoadGuessList returns the large 14k-word dictionary used to validate guesses.
func LoadGuessList() []string {
	guessListOnce.Do(func() {
		guessList = parseList(wordListData)
	})
	return guessList
}

// LoadAnswerList returns the curated ~2.3k-word answer pool (common English words).
func LoadAnswerList() []string {
	answerListOnce.Do(func() {
		answerList = parseList(answerListData)
	})
	return answerList
}

// ReadFile preserved for backwards compatibility. Returns the guess list.
func ReadFile() []string {
	return LoadGuessList()
}

// GetWord returns a random word from the curated answer pool.
func GetWord() string {
	words := LoadAnswerList()
	if len(words) == 0 {
		panic("answer list is empty — embedded AnswerList.txt is missing or unreadable")
	}
	return words[rand.Intn(len(words))]
}

func isValidGuess(guess string) bool {
	if len(guess) != 5 {
		return false
	}
	for _, r := range guess {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	for _, w := range LoadGuessList() {
		if w == guess {
			return true
		}
	}
	for _, w := range LoadAnswerList() {
		if w == guess {
			return true
		}
	}
	return false
}

func GetGuess() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	guess := strings.ToLower(strings.TrimSpace(text))
	if len(guess) != 5 {
		fmt.Println("Please enter a 5 letter word")
		return GetGuess()
	}
	if !isValidGuess(guess) {
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
	fmt.Print("\033[1A\r")
	fmt.Print("\033[K")
	for i := range rightSpots {
		if rightSpots[i] {
			fmt.Print(Green + BlackText + strings.ToUpper(string(guess[i])) + White)
		} else if rightLetters[i] {
			fmt.Print(Yellow + BlackText + strings.ToUpper(string(guess[i])) + White)
		} else {
			fmt.Print(White + strings.ToUpper(string(guess[i])) + White)
		}
	}
	fmt.Println()
}

func TrackLetterUsage(word string) map[rune]int {
	letterCounts := make(map[rune]int)
	for _, letter := range word {
		letterCounts[letter]++
	}
	return letterCounts
}

func CheckGuessWithDuplicates(guess, word string) []string {
	results := make([]string, len(guess))
	if len(guess) != len(word) {
		// Lengths mismatch — every position is gray to avoid index panic.
		for i := range results {
			results[i] = "gray"
		}
		return results
	}
	letterUsage := TrackLetterUsage(word)
	for i := range guess {
		if guess[i] == word[i] {
			results[i] = "green"
			letterUsage[rune(guess[i])]--
		}
	}
	for i := range guess {
		if results[i] != "" {
			continue
		}
		if letterUsage[rune(guess[i])] > 0 {
			results[i] = "yellow"
			letterUsage[rune(guess[i])]--
		} else {
			results[i] = "gray"
		}
	}
	return results
}

func ShowResultsEnhanced(guess, word string) {
	fmt.Print("\033[1A\r")
	fmt.Print("\033[K")
	results := CheckGuessWithDuplicates(guess, word)
	for i, result := range results {
		if result == "green" {
			fmt.Print(Green + BlackText + strings.ToUpper(string(guess[i])) + White)
		} else if result == "yellow" {
			fmt.Print(Yellow + BlackText + strings.ToUpper(string(guess[i])) + White)
		} else {
			fmt.Print(White + strings.ToUpper(string(guess[i])) + White)
		}
	}
	fmt.Println()
}

func ShowResultsEnhancedAI(guess, word string) {
	results := CheckGuessWithDuplicates(guess, word)
	for i, result := range results {
		if result == "green" {
			fmt.Print(Green + BlackText + strings.ToUpper(string(guess[i])) + White)
		} else if result == "yellow" {
			fmt.Print(Yellow + BlackText + strings.ToUpper(string(guess[i])) + White)
		} else {
			fmt.Print(White + strings.ToUpper(string(guess[i])) + White)
		}
	}
	fmt.Println()
}

var PlayerGuessNum = 0
var PlayerGotIt = false

func PlayerPlay(word string) {
	fmt.Println(" __        __            _ _      ")
	fmt.Println(" \\ \\      / /__  _ __ __| | | ___ ")
	fmt.Println("  \\ \\ /\\ / / _ \\| '__/ _` | |/ _ \\")
	fmt.Println("   \\ V  V / (_) | | | (_| | |  __/")
	fmt.Println("    \\_/\\_/ \\___/|_|  \\__,_|_|\\___|")
	fmt.Println("-----------------------------------")
	won := false

	for i := 0; i < 6; i++ {
		guess := strings.ToLower(GetGuess())
		PlayerGuessNum++
		ShowResultsEnhanced(guess, word)
		if DoesGuessMatch(guess, word) {
			fmt.Println("You Got It!")
			PlayerGotIt = true
			won = true
			break
		}
	}
	if !won {
		fmt.Println("You are out of guesses! The word was", word)
	}
}
