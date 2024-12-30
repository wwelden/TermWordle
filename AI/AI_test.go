package AI

import (
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
		name     string
		guess    string
		word     string
		wordList []string
		expected []string
	}{
		{
			name:     "P E A",
			guess:    "apple",
			word:     "peach",
			wordList: []string{"peach", "grape", "plane", "apple", "brick"},
			expected: []string{"peach", "grape", "plane", "apple"}, // "p", "e", "a" are yellow letters
		},
		{
			name:     "T A L",
			guess:    "table",
			word:     "plant",
			wordList: []string{"plant", "plate", "blame", "table"},
			expected: []string{"plant", "plate", "table"}, // "t", "a", "l" are yellow letters
		},
		{
			name:     "S T O N E",
			guess:    "stone",
			word:     "notes",
			wordList: []string{"notes", "tones", "stone", "nests"},
			expected: []string{"notes", "tones", "stone"}, // "s", "t", "o", "n", "e" are yellow letters
		},
		{
			name:     "All Letters are yellow",
			guess:    "apple",
			word:     "apple",
			wordList: []string{"apple", "ample", "apply"},
			expected: []string{"apple", "ample"}, // All letters are yellow
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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
					t.Errorf("Expected word %s not found in matches: %v", test.expected, matches)
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

//

func TestGetRightPlaces(t *testing.T) {
	tests := []struct {
		name     string
		guess    string
		answer   string
		expected string
	}{
		{
			name:     "test 1",
			guess:    "plant",
			answer:   "pleat",
			expected: "pl--t",
		},
		{
			name:     "test 2",
			guess:    "apple",
			answer:   "spade",
			expected: "-p--e",
		},
		{
			name:     "test 3",
			guess:    "optic",
			answer:   "spoil",
			expected: "-p-i-",
		},
		{
			name:     "test 4",
			guess:    "apple",
			answer:   "xxxxx",
			expected: "-----",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := getRightPlaces(test.guess, test.answer)

			for _, expectedWord := range test.expected {
				found := false
				for _, result := range results {
					if result == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %v was not in filtered results: %v", expectedWord, results)
				}
			}
		})

	}
}

func TestRightPlaces(t *testing.T) {
	tests := []struct {
		name     string
		guess    string
		answer   string
		wordlist []string
		expected []string
	}{
		{
			name:     "test 1",
			guess:    "plant",
			answer:   "pleat",
			wordlist: []string{"platt", "plast", "pluot", "plait", "pleat", "plant", "abuzz", "abyss", "above", "alamo", "album"},
			expected: []string{"platt", "plast", "pluot", "plait", "pleat", "plant"},
		},
		{
			name:     "test 2",
			guess:    "apple",
			answer:   "spade",
			wordlist: []string{"apple", "spade", "spare", "apply", "april"},
			expected: []string{"apple", "spade", "spare"},
		},
		{
			name:     "test 3",
			guess:    "brick",
			answer:   "break",
			wordlist: []string{"brick", "brink", "brake", "break", "bread", "beach", "bench"},
			expected: []string{"brick", "brink", "brake", "break", "bread"},
		},
		{
			name:     "test 4",
			guess:    "stain",
			answer:   "steam",
			wordlist: []string{"stain", "stand", "stamp", "steam", "steal", "stack", "start"},
			expected: []string{"stain", "stand", "stamp", "steam", "steal"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := rightPlaces(test.guess, test.answer, test.wordlist)

			for _, expectedWord := range test.expected {
				found := false
				for _, result := range results {
					if result == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %v was not in filtered results: %v", expectedWord, results)
				}
			}
		})

	}
}

func TestGetWrongPlaces(t *testing.T) {
	tests := []struct {
		name     string
		guess    string
		answer   string
		expected string
	}{
		{
			name:     "test 1",
			guess:    "plant",
			answer:   "pleat",
			expected: "plat",
		},
		{
			name:     "test 2",
			guess:    "apple",
			answer:   "spade",
			expected: "ape",
		},
		{
			name:     "test 3",
			guess:    "optic",
			answer:   "spoil",
			expected: "poi",
		},
		{
			name:     "test 4",
			guess:    "apple",
			answer:   "xxxxx",
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := getWrongPlaces(test.guess, test.answer)

			for _, expectedWord := range test.expected {
				found := false
				for _, result := range results {
					if result == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %v was not in filtered results: %v", expectedWord, results)
				}
			}
		})

	}
}

func TestHasLetters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		word     string
		expected bool
	}{
		{
			name:     "test 1",
			input:    "elat",
			word:     "pleat",
			expected: true,
		},
		{
			name:     "test 2",
			input:    "eds",
			word:     "spade",
			expected: true,
		},
		{
			name:     "test 3",
			input:    "oil",
			word:     "spoil",
			expected: true,
		},
		{
			name:     "test 4",
			input:    "apple",
			word:     "xxxxx",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := hasLetters(test.input, test.word)

			if results != test.expected {
				t.Errorf("Expected word %v was not in filtered results: %v", test.expected, results)
			}
		})

	}
}

func TestWrongPlaces(t *testing.T) {
	tests := []struct {
		name     string
		guess    string
		answer   string
		wordlist []string
		expected []string
	}{
		{
			name:     "test 1 - letter in wrong position",
			guess:    "plant",
			answer:   "pleat",
			wordlist: []string{"platt", "plast", "pluot", "plait", "pleat", "plant"},
			expected: []string{"pleat"}, // Only pleat has 'a' in wrong position
		},
		{
			name:     "test 2 - multiple letters in wrong positions",
			guess:    "apple",
			answer:   "spade",
			wordlist: []string{"apple", "spade", "spare", "apply", "april"},
			expected: []string{"spade"}, // Only spade has both 'a' and 'p' in wrong positions
		},
		{
			name:     "test 3 - no letters in wrong positions",
			guess:    "brick",
			answer:   "pleat",
			wordlist: []string{"pleat", "plant", "spare", "brick"},
			expected: []string{}, // No words have letters in wrong positions
		},
		{
			name:     "test 4 - repeated letters",
			guess:    "speed",
			answer:   "spade",
			wordlist: []string{"speed", "spade", "shade", "space"},
			expected: []string{"spade"}, // Only spade has 'e' in wrong position
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := wrongPlaces(test.guess, test.answer, test.wordlist)

			for _, expectedWord := range test.expected {
				found := false
				for _, result := range results {
					if result == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %v was not in filtered results: %v", expectedWord, results)
				}
			}
		})

	}
}

func TestNotHaveLetters(t *testing.T) {
	tests := []struct {
		name     string
		guess    string
		answer   string
		expected string
	}{
		{
			name:     "test 1",
			guess:    "plant",
			answer:   "pleat",
			expected: "n",
		},
		{
			name:     "test 2",
			guess:    "apple",
			answer:   "spade",
			expected: "l",
		},
		{
			name:     "test 3",
			guess:    "optic",
			answer:   "spoil",
			expected: "tc",
		},
		{
			name:     "test 4",
			guess:    "apple",
			answer:   "xxxxx",
			expected: "apple",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := notHaveLetters(test.guess, test.answer)

			for _, expectedWord := range test.expected {
				found := false
				for _, result := range results {
					if result == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %v was not in filtered results: %v", expectedWord, results)
				}
			}
		})

	}
}

func TestGetWordsWithoutLetters(t *testing.T) {
	tests := []struct {
		name     string
		guess    string
		answer   string
		wordlist []string
		expected []string
	}{
		{
			name:     "test 1 - filter words with n",
			guess:    "plant",
			answer:   "pleat",
			wordlist: []string{"platt", "plast", "pluot", "plait", "pleat", "plant"},
			expected: []string{"pleat"}, // Only word without 'n'
		},
		{
			name:     "test 2 - filter words with l",
			guess:    "apple",
			answer:   "spade",
			wordlist: []string{"apple", "spade", "spare", "apply", "april"},
			expected: []string{"spade"}, // Only word without 'l'
		},
		{
			name:     "test 3 - filter words with t,c",
			guess:    "optic",
			answer:   "spoil",
			wordlist: []string{"trick", "spoil", "spear", "topic"},
			expected: []string{"spoil"}, // Only word without 't' and 'c'
		},
		{
			name:     "test 4 - filter all letters",
			guess:    "apple",
			answer:   "xxxxx",
			wordlist: []string{"apple", "bring", "crown", "drink"},
			expected: []string{"bring", "crown", "drink"}, // Words without any letters from 'apple'
		},
		{
			name:     "test 5 - no common letters",
			guess:    "brick",
			answer:   "apple",
			wordlist: []string{"grape", "brick", "table", "apple", "peach"},
			expected: []string{"apple", "peach"}, // Only words with no letters in common with "brick"
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := getWordsWithoutLetters(test.guess, test.answer, test.wordlist)

			for _, expectedWord := range test.expected {
				found := false
				for _, result := range results {
					if result == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %v was not in filtered results: %v", expectedWord, results)
				}
			}
		})

	}
}

func TestSubsetEnhanced(t *testing.T) {
	tests := []struct {
		name           string
		greenMatches   []string
		yellowMatches  []string
		wordWithOut    []string
		expected       []string // Expected filtered words at the end of the test
		expectedLength int
	}{
		{
			name:           "All matches overlap",
			greenMatches:   []string{"plant", "pleat", "plait"},
			yellowMatches:  []string{"plant", "pleat", "plait", "plain"},
			wordWithOut:    []string{"plant", "pleat", "plait", "pluot"},
			expected:       []string{"plant", "pleat", "plait"},
			expectedLength: 3,
		},
		{
			name:           "No matches overlap",
			greenMatches:   []string{"plant", "pleat"},
			yellowMatches:  []string{"spare", "spade"},
			wordWithOut:    []string{"brick", "break"},
			expected:       []string{},
			expectedLength: 0,
		},
		{
			name:           "Partial overlap",
			greenMatches:   []string{"spare", "spade", "space"},
			yellowMatches:  []string{"spade", "space", "speak"},
			wordWithOut:    []string{"space", "spade", "spark"},
			expected:       []string{"spade", "space"},
			expectedLength: 2,
		},
		{
			name:           "Multiple matches from larger lists",
			greenMatches:   []string{"plant", "pleat", "plait", "plain", "plate", "place", "plane", "plank", "plaza", "plage"},
			yellowMatches:  []string{"spare", "spade", "space", "speak", "spear", "pleat", "plait", "plain", "plate", "place"},
			wordWithOut:    []string{"brick", "break", "bread", "pleat", "plait", "plain", "plate", "place", "bring", "broad"},
			expected:       []string{"pleat", "plait", "plain", "plate", "place"},
			expectedLength: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := subsetEnhanced(test.greenMatches, test.yellowMatches, test.wordWithOut)

			if len(results) != test.expectedLength {
				t.Errorf("FindSubSet2 returned %d subset, expected 1", len(results)) // test not finished
			}
			for _, expectedWord := range test.expected {
				found := false
				for _, result := range results {
					if result == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %s was not in filtered results: %v", expectedWord, results)
				}
			}
		})

	}
}

func TestSubSetEmptyLists(t *testing.T) {
	tests := []struct {
		name           string
		greenMatches   []string
		yellowMatches  []string
		wordWithOut    []string
		expected       []string // Expected filtered words at the end of the test
		expectedLength int
	}{
		{
			name:           "All matches overlap",
			greenMatches:   []string{"plant", "pleat", "plait"},
			yellowMatches:  []string{"plant", "pleat", "plait", "plain"},
			wordWithOut:    []string{"plant", "pleat", "plait", "pluot"},
			expected:       []string{"plant", "pleat", "plait"},
			expectedLength: 3,
		},
		{
			name:           "No matches overlap",
			greenMatches:   []string{"plant", "pleat"},
			yellowMatches:  []string{"spare", "spade"},
			wordWithOut:    []string{"brick", "break"},
			expected:       []string{},
			expectedLength: 0,
		},
		{
			name:           "Partial overlap",
			greenMatches:   []string{"spare", "spade", "space"},
			yellowMatches:  []string{"spade", "space", "speak"},
			wordWithOut:    []string{"space", "spade", "spark"},
			expected:       []string{"spade", "space"},
			expectedLength: 2,
		},
		{
			name:           "Multiple matches from larger lists",
			greenMatches:   []string{"plant", "pleat", "plait", "plain", "plate", "place", "plane", "plank", "plaza", "plage"},
			yellowMatches:  []string{"spare", "spade", "space", "speak", "spear", "pleat", "plait", "plain", "plate", "place"},
			wordWithOut:    []string{"brick", "break", "bread", "pleat", "plait", "plain", "plate", "place", "bring", "broad"},
			expected:       []string{"pleat", "plait", "plain", "plate", "place"},
			expectedLength: 5,
		},
		{
			name:           "Empty green matches",
			greenMatches:   []string{},
			yellowMatches:  []string{"spare", "spade", "space"},
			wordWithOut:    []string{"spare", "spade", "space"},
			expected:       []string{"spare", "spade", "space"},
			expectedLength: 3,
		},
		{
			name:           "Empty yellow matches",
			greenMatches:   []string{"spare", "spade", "space"},
			yellowMatches:  []string{},
			wordWithOut:    []string{"spare", "spade", "space"},
			expected:       []string{"spare", "spade", "space"},
			expectedLength: 3,
		},
		{
			name:           "Empty without matches",
			greenMatches:   []string{"spare", "spade", "space"},
			yellowMatches:  []string{"spare", "spade", "space"},
			wordWithOut:    []string{},
			expected:       []string{"spare", "spade", "space"},
			expectedLength: 3,
		},
		{
			name:           "All empty lists",
			greenMatches:   []string{},
			yellowMatches:  []string{},
			wordWithOut:    []string{},
			expected:       []string{},
			expectedLength: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := SubSetEmptyLists(test.greenMatches, test.yellowMatches, test.wordWithOut)

			if len(results) != test.expectedLength {
				t.Errorf("FindSubSet2 returned %d subset, expected 1", len(results)) // test not finished
			}
			for _, expectedWord := range test.expected {
				found := false
				for _, result := range results {
					if result == expectedWord {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected word %s was not in filtered results: %v", expectedWord, results)
				}
			}
		})

	}
}

func TestCompeteEnhanced(t *testing.T) {
	tests := []struct {
		name       string
		word       string
		wordList   []string
		firstGuess string
		expected   string // Expected final guess
	}{
		{
			name:       "No common letters",
			word:       "apple",
			wordList:   []string{"grape", "brick", "table", "apple", "peach"},
			firstGuess: "grape",
			expected:   "apple", // Should find the target word
		},
		{
			name:       "Repeated letters in word",
			word:       "bloom",
			wordList:   []string{"bloom", "roomy", "booms", "gloom", "broom"},
			firstGuess: "bloom",
			expected:   "bloom", // Should handle repeated letters correctly
		},
		{
			name:       "Case sensitivity",
			word:       "Apple",
			wordList:   []string{"grape", "brick", "table", "apple", "peach"},
			firstGuess: "grape",
			expected:   "apple", // Function should be case insensitive
		},
		{
			name:       "Shuffled word list",
			word:       "apple",
			wordList:   []string{"peach", "table", "grape", "brick", "apple"},
			firstGuess: "peach",
			expected:   "grape", // Should work with randomized list order
		},
		{
			name:       "Invalid first guess",
			word:       "apple",
			wordList:   []string{"grape", "brick", "table", "apple", "peach"},
			firstGuess: "grape", // First word in the list
			expected:   "apple", // Should still find the target word
		},
		// {
		// 	name:       "Multiple words with no matches",
		// 	word:       "apple",
		// 	wordList:   []string{"elder", "bowed", "boned", "donut", "pound", "hundo"},
		// 	firstGuess: "elder",
		// 	expected:   "boned", // Should find the first word that matches
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Define a custom first guess function that returns the first word in the list
			firstGuessFunc := func(wordList []string) string {
				if len(wordList) == 0 {
					return ""
				}
				return wordList[0]
			}

			// Run CompeteEnhanced with the custom first guess function
			result := CompeteEnhanced(test.word, test.wordList, firstGuessFunc)

			if result != test.expected {
				t.Errorf("CompeteEnhanced(%q, %v, firstGuessFunc) = %q, want %q",
					test.word, test.wordList, result, test.expected)
			}
		})
	}
}
