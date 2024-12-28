package main

import (
	"fmt"
	"strings"
)

func TrackLetterUsage(word string) map[rune]int {
	letterCounts := make(map[rune]int)
	for _, letter := range word {
		letterCounts[letter]++
	}
	return letterCounts
}

func CheckGuessWithDuplicates(guess, word string) []string {
	results := make([]string, len(guess))
	letterUsage := TrackLetterUsage(word)

	// First pass: Mark greens
	for i := range guess {
		if guess[i] == word[i] {
			results[i] = "green"
			letterUsage[rune(guess[i])]--
		}
	}

	// Second pass: Mark yellows
	for i := range guess {
		if results[i] != "" { // Skip already marked greens
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
			fmt.Print(Green + "\033[30m" + strings.ToUpper(string(guess[i])) + White)
		} else if result == "yellow" {
			fmt.Print(Yellow + "\033[30m" + strings.ToUpper(string(guess[i])) + White)
		} else {
			fmt.Print(White + strings.ToUpper(string(guess[i])) + White)
		}
	}
	fmt.Println()
}

func ShowResultsEnhancedAI(guess, word string) {
	// fmt.Print("\033[1A\r")
	// fmt.Print("\033[K")
	results := CheckGuessWithDuplicates(guess, word)
	for i, result := range results {
		if result == "green" {
			fmt.Print(Green + "\033[30m" + strings.ToUpper(string(guess[i])) + White)
		} else if result == "yellow" {
			fmt.Print(Yellow + "\033[30m" + strings.ToUpper(string(guess[i])) + White)
		} else {
			fmt.Print(White + strings.ToUpper(string(guess[i])) + White)
		}
	}
	fmt.Println()
}
