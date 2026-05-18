package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/wwelden/TermWordle/AI"
	"github.com/wwelden/TermWordle/Wordle"
)

func CanYouBeatTheAI(word string) {
	Wordle.PlayerPlay(word)
	AI.AIPlay(word)
	fmt.Println("Player guesses:", Wordle.PlayerGuessNum)
	fmt.Println("AI guesses:", AI.AIGuessNum)
	switch {
	case Wordle.PlayerGotIt && !AI.AIGotIt:
		fmt.Println("Player wins!")
	case AI.AIGotIt && !Wordle.PlayerGotIt:
		fmt.Println("AI wins!")
	case !Wordle.PlayerGotIt && !AI.AIGotIt:
		fmt.Println("Nobody got it — it's a tie.")
	case Wordle.PlayerGuessNum < AI.AIGuessNum:
		fmt.Println("Player wins!")
	case AI.AIGuessNum < Wordle.PlayerGuessNum:
		fmt.Println("AI wins!")
	default:
		fmt.Println("You tied!")
	}
}

func main() {
	mode := flag.String("mode", "vs", "game mode: vs (player vs AI), player, or ai")
	flag.Parse()

	word := strings.ToLower(Wordle.GetWord())

	switch *mode {
	case "vs":
		CanYouBeatTheAI(word)
	case "player":
		Wordle.PlayerPlay(word)
	case "ai":
		AI.AIPlay(word)
	default:
		fmt.Fprintf(os.Stderr, "unknown mode %q (expected: vs, player, ai)\n", *mode)
		os.Exit(2)
	}
}
