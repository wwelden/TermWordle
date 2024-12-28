package main

import (
	"fmt"
	"strings"
)

func main() {
	word := strings.ToLower(GetWord())
	fmt.Println(" __        __            _ _      ")
	fmt.Println(" \\ \\      / /__  _ __ __| | | ___ ")
	fmt.Println("  \\ \\ /\\ / / _ \\| '__/ _` | |/ _ \\")
	fmt.Println("   \\ V  V / (_) | | | (_| | |  __/")
	fmt.Println("    \\_/\\_/ \\___/|_|  \\__,_|_|\\___|")
	fmt.Println("-----------------------------------")
	won := false
	for i := 0; i < 6; i++ {
		guess := strings.ToLower(GetGuess())
		ShowResults(guess, word)
		if DoesGuessMatch(guess, word) {
			fmt.Println("You win!")
			won = true
			break
		}
	}
	if !won {
		fmt.Println("You are out of guesses! The word was", word)
	}
}
