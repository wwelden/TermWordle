package main

import (
	"fmt"
	"strings"

	"github.com/wwelden/TermWordle/AI"
	"github.com/wwelden/TermWordle/Wordle"
)

func CanYouBeatTheAI(word string) {
	Wordle.PlayerPlay(word)
	AI.AIPlay(word)
	fmt.Println("Player guesses:", Wordle.PlayerGuessNum)
	fmt.Println("AI guesses:", AI.AIGuessNum)
	if Wordle.PlayerGuessNum > AI.AIGuessNum || (AI.AIGotIt && !Wordle.PlayerGotIt) {
		fmt.Println("AI wins!")
	} else if Wordle.PlayerGotIt == AI.AIGotIt {
		fmt.Println("You tied!")
	} else {
		fmt.Println("Player wins!")
	}
}
func RunAI() {
	word := strings.ToLower(Wordle.GetWord())
	AI.AIPlay(word)
}

func main() {
	word := strings.ToLower(Wordle.GetWord())
	CanYouBeatTheAI(word)
	// RunAI()
}
