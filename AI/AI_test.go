package AI

import (
	"fmt"
	"testing"

	"github.com/wwelden/TermWordle/Wordle"
)

//AI.go Tests

func TestGetContainedLetters(t *testing.T) {
	guess := "apple"
	word := "apply"
	containedLetters := GetContainedLetters(guess, word)
	if containedLetters != "appl" {
		t.Errorf("GetContainedLetters returned %s, expected appl", containedLetters)
	}
}

func TestGetAllWordsThatContain(t *testing.T) {
	guess := "apple"
	word := "apply"
	wordList := Wordle.ReadFile()
	matches := GetAllWordsThatContain(guess, word, wordList)
	if len(matches) != 197 {
		t.Errorf("GetAllWordsThatContain returned %d matches, expected 197", len(matches))
	}
}

func TestGetAllWordsThatMatch(t *testing.T) {
	guess := "apple"
	word := "apply"
	wordList := Wordle.ReadFile()
	matches := GetAllWordsThatMatch(guess, word, wordList)
	if len(matches) != 2 {
		t.Errorf("GetAllWordsThatMatch returned %d matches, expected 2", len(matches))
	}
}

func TestGetAllWordsWithOutLetters(t *testing.T) {
	tests := []struct {
		name     string
		guess    string
		word     string
		wordList []string
		expected []string // The expected matches after filtering
	}{
		{
			name:     "Basic functionality",
			guess:    "apple",
			word:     "apply",
			wordList: []string{"apple", "apply", "apples", "applesauce"},
			expected: []string{"apply"}, // "apply" is the only word without "e"
		},
		{
			name:     "No matching words",
			guess:    "apple",
			word:     "apply",
			wordList: []string{"apples", "applesauce"},
			expected: []string{}, // All words contain letters not in "apply"
		},
		{
			name:     "All words match",
			guess:    "brick",
			word:     "grape",
			wordList: []string{"grape", "grape", "grape"},
			expected: []string{"grape", "grape", "grape"}, // All words are valid since they match
		},
		{
			name:     "Mixed valid and invalid words",
			guess:    "plane",
			word:     "apple",
			wordList: []string{"plane", "plate", "apple", "apply", "plead"},
			expected: []string{"apple", "apply", "plead", "plate"}, // Only "apple" and "apply" contain valid letters
		},
		{
			name:     "Empty word list",
			guess:    "apple",
			word:     "apply",
			wordList: []string{},
			expected: []string{}, // No words to filter
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matches := GetAllWordsWithOutLetters(test.guess, test.word, test.wordList)

			// Check the length of the results
			if len(matches) != len(test.expected) {
				t.Errorf("For %s: expected %d matches, but got %d", test.name, len(test.expected), len(matches))
			}

			// Check the content of the results
			for _, expectedWord := range test.expected {
				found := false
				for _, match := range matches {
					if match == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("For %s: expected word %s not found in results %v", test.name, expectedWord, matches)
				}
			}

			// Ensure no unexpected words are included
			for _, match := range matches {
				found := false
				for _, expectedWord := range test.expected {
					if match == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("For %s: unexpected word %s found in results %v", test.name, match, matches)
				}
			}
		})
	}
}

func TestFindSubSet(t *testing.T) {
	matches := []string{"apple", "apply", "apples", "applesauce"}
	contained := []string{"apple", "apples"}
	subset := findSubSet(matches, contained)
	if len(subset) != 2 {
		t.Errorf("FindSubSet returned %d subset, expected 2", len(subset))
	}
}

func TestFirstGuess(t *testing.T) {
	wordList := Wordle.ReadFile()
	guess := FirstGuess(wordList)
	if len(guess) != 5 {
		t.Errorf("FirstGuess returned a guess with length %d, expected 5", len(guess))
	}
}

func TestFindSubSet2(t *testing.T) {
	matches := []string{"apple", "apply", "apples", "applesauce"}
	contained := []string{"apple", "apples"}
	wordWithOut := []string{"apple"}
	subset := findSubSet2(matches, contained, wordWithOut)
	if len(subset) != 1 {
		t.Errorf("FindSubSet2 returned %d subset, expected 1", len(subset))
	}
}

func TestWordHasLetter(t *testing.T) {
	word := "apple"
	letter := "a"
	hasLetter := WordHasLetter(word, letter)
	if !hasLetter {
		t.Errorf("WordHasLetter returned %t, expected true", hasLetter)
	}
}

func TestGetLettersNotInWord(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		guess    string
		expected string
	}{
		{
			name:     "Basic case with one unmatched letter",
			word:     "apple",
			guess:    "apply",
			expected: "y", // "y" is not in "apple"
		},
		{
			name:     "All letters unmatched",
			word:     "apple",
			guess:    "brick",
			expected: "brick", // All letters in "brick" are not in "apple"
		},
		{
			name:     "No unmatched letters",
			word:     "apple",
			guess:    "apple",
			expected: "", // All letters in "guess" are in "word"
		},
		{
			name:     "Repeated unmatched letters",
			word:     "apple",
			guess:    "applyy",
			expected: "yy", // Only one "y" should be returned
		},
		{
			name:     "Mixed matches and unmatched letters",
			word:     "apple",
			guess:    "apric",
			expected: "ric", // "r", "c", and "i" are not in "apple"
		},
		{
			name:     "Empty guess string",
			word:     "apple",
			guess:    "",
			expected: "", // No letters to match
		},
		{
			name:     "Empty word string",
			word:     "",
			guess:    "apple",
			expected: "apple", // All letters in "guess" are unmatched
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetLettersNotInWord(test.word, test.guess)

			// Compare the result with the expected output
			if result != test.expected {
				t.Errorf("For word=%q and guess=%q, expected %q but got %q",
					test.word, test.guess, test.expected, result)
			}
		})
	}
}

func TestGetLettersInWord(t *testing.T) {
	tests := []struct {
		guess    string
		word     string
		expected string
	}{
		{"paper", "apply", "pa"},
		{"hello", "world", "lo"},
		{"abc", "xyz", ""},
		{"repeat", "tape", "epat"},
		{"AaBb", "aabb", "ab"}, // Case-sensitive
	}

	for _, test := range tests {
		result := GetLettersInWord(test.guess, test.word)
		if result != test.expected {
			t.Errorf("For guess %q and word %q, expected %q but got %q",
				test.guess, test.word, test.expected, result)
		}
	}
}

func TestGetAllWordsWithYellowLetters(t *testing.T) {
	tests := []struct {
		guess    string
		word     string
		wordList []string
		expected []string
	}{
		{
			guess:    "apple",
			word:     "peach",
			wordList: []string{"peach", "grape", "plane", "apple"},
			expected: []string{"peach", "grape", "plane"}, // "p", "e", "a" are yellow letters
		},
		{
			guess:    "table",
			word:     "plant",
			wordList: []string{"plant", "plate", "blame", "table"},
			expected: []string{"plant", "plate", "blame"}, // "t", "a", "l" are yellow letters
		},
		{
			guess:    "stone",
			word:     "notes",
			wordList: []string{"notes", "tones", "stone", "nest"},
			expected: []string{"notes", "tones", "nest"}, // "s", "t", "o", "n", "e" are yellow letters
		},
		{
			guess:    "apple",
			word:     "apple",
			wordList: []string{"apple", "ample", "apply"},
			expected: []string{"apple", "ample", "apply"}, // All letters are yellow
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("guess=%s, word=%s", test.guess, test.word), func(t *testing.T) {
			matches := GetAllWordsWithYellowLetters(test.guess, test.word, test.wordList)

			// Validate the results
			if len(matches) != len(test.expected) {
				t.Errorf("Expected %d matches, got %d", len(test.expected), len(matches))
			}

			for _, expectedWord := range test.expected {
				found := false
				for _, match := range matches {
					if match == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %s not found in matches: %v", expectedWord, matches)
				}
			}
		})
	}
}

func TestGetAllWordsWithGreenLetters(t *testing.T) {
	guess := "paper"
	word := "apply"
	wordList := []string{"apple", "food", "house", "housework", "housework"}
	matches := GetAllWordsWithGreenLetters(guess, word, wordList)

	// Define the expected result
	expected := []string{"apple"}

	// Check the length of the result
	if len(matches) != len(expected) {
		t.Errorf("GetAllWordsWithGreenLetters returned %d matches, expected %d", len(matches), len(expected))
	}

	// Optionally, check the content of the result
	for i, match := range matches {
		if match != expected[i] {
			t.Errorf("Mismatch at index %d: got %s, expected %s", i, match, expected[i])
		}
	}
}

func TestCompete(t *testing.T) {
	word := "apple"
	wordList := Wordle.ReadFile()
	guess := Compete(word, wordList)
	if len(guess) != 5 {
		t.Errorf("Compete returned a guess with length %d, expected 5", len(guess))
	}
}

func TestCompeteEnhanced(t *testing.T) {
	tests := []struct {
		name       string
		word       string
		wordList   []string
		firstGuess string
		expected   []string // Expected filtered words at the end of the test
	}{
		{
			name:       "No common letters",
			word:       "apple",
			wordList:   []string{"grape", "brick", "table", "apple", "peach"},
			firstGuess: "brick",
			expected:   []string{"apple", "peach"}, // Only words with no letters in common with "brick"
		},
		{
			name:       "Some letters in correct positions",
			word:       "apple",
			wordList:   []string{"apple", "ample", "angle", "apply", "baker"},
			firstGuess: "ample",
			expected:   []string{"apple"}, // Words matching "a", "p", and "l" positions
		},
		{
			name:       "All letters in wrong positions",
			word:       "apple",
			wordList:   []string{"apple", "apply", "peach", "plane", "baker"},
			firstGuess: "plane",
			expected:   []string{"apple", "apply"}, // Words containing "p", "l", "a", "e" but in wrong spots
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Mock fstGuess to return the specified first guess
			mockFstGuess := func(wordList []string) string {
				return test.firstGuess
			}

			// Run CompeteEnhanced with the mocked fstGuess function
			guess := CompeteEnhanced(test.word, test.wordList, mockFstGuess)

			// Validate the final filtered word list
			filteredWords := GetAllWordsThatMatch(guess, test.word, test.wordList)
			if len(filteredWords) != len(test.expected) {
				t.Errorf("Expected %d words, got %d: %v", len(test.expected), len(filteredWords), filteredWords)
			}

			for _, expectedWord := range test.expected {
				found := false
				for _, filteredWord := range filteredWords {
					if filteredWord == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %s was not in filtered results: %v", expectedWord, filteredWords)
				}
			}
		})
	}
}
