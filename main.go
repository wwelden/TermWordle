package main

import (
	"fmt"
	"strings"
)

func CanYouBeatTheAI(word string) {
	PlayerPlay(word)
	AIPlay(word)
	fmt.Println("Player guesses:", playerGuessNum)
	fmt.Println("AI guesses:", AIGuessNum)
	if playerGuessNum > AIGuessNum || (AIGotIt && !playerGotIt) {
		fmt.Println("AI wins!")
	} else if playerGuessNum == AIGuessNum {
		fmt.Println("You tied!")
	} else {
		fmt.Println("Player wins!")
	}
}

func main() {
	word := strings.ToLower(GetWord())
	CanYouBeatTheAI(word)
}
