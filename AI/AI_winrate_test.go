package AI

import (
	"io"
	"math/rand"
	"os"
	"testing"

	"github.com/wwelden/TermWordle/Wordle"
)

// runAIGame plays one game silently and returns (won, guesses).
func runAIGame(target string) (bool, int) {
	ResetState()
	wordList := Wordle.LoadAnswerList()
	guess := FirstGuess(wordList)
	for i := 0; i < 6; i++ {
		if guess == "" {
			break
		}
		AIGuessNum++
		if Wordle.DoesGuessMatch(guess, target) {
			AIGotIt = true
			return true, AIGuessNum
		}
		wordList = filterCandidates(guess, target, wordList)
		guess = FirstGuess(wordList)
	}
	return false, AIGuessNum
}

func TestAIWinRate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping AI win-rate harness in short mode")
	}

	// Silence the AI's stdout chatter — CompeteLoop's fmt.Println would otherwise spam.
	stdout := os.Stdout
	devnull, _ := os.Open(os.DevNull)
	os.Stdout = devnull
	defer func() {
		os.Stdout = stdout
		_ = devnull.Close()
		_, _ = io.Copy(io.Discard, devnull)
	}()

	answers := Wordle.LoadAnswerList()
	rng := rand.New(rand.NewSource(42))

	const trials = 200
	wins := 0
	totalGuessesOnWin := 0
	for i := 0; i < trials; i++ {
		target := answers[rng.Intn(len(answers))]
		won, guesses := runAIGame(target)
		if won {
			wins++
			totalGuessesOnWin += guesses
		}
	}

	winRate := float64(wins) / float64(trials) * 100
	var avg float64
	if wins > 0 {
		avg = float64(totalGuessesOnWin) / float64(wins)
	}
	t.Logf("AI played %d games, won %d (%.1f%%), avg guesses on win: %.2f", trials, wins, winRate, avg)

	if winRate < 90.0 {
		t.Errorf("AI win rate %.1f%% is below 90%% — AI filter logic may be broken", winRate)
	}
}
