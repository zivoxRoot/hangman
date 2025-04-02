package hangman

import (
	"math/rand"
	"strings"

	"fmt"
	"os"
)

// Hangman is a hangman game instance.
type hangman struct {
	TriesLeft    int
	GoalWord     string
	LettersDone  []string
	RightLetters []string
}

// NewHangman initializes a new hangman game. It takes a path to a custom worlist, checks if it exists, else use the default wordlist. It returns a pointer to the Hangman and the used wordlist.
func newHangman(wordlistPath string) (*hangman, string) {

	// usedWordlist is the used wordlist between the one given by the user and the default one.
	var usedWordlist string

	// Get a random word from the wordlist path.
	randomWord, err := getGoalWord(wordlistPath)
	usedWordlist = wordlistPath
	if err != nil {

		// Check the default wordlist exists.
		wordlistPath, err := getDefaultWordlist()
		if err != nil {
			fmt.Println("Could not get the default wordlist :", err)
			os.Exit(1)
		}

		// Get a random word from the default wordlist.
		randomWord, err = getGoalWord(wordlistPath)
		usedWordlist = wordlistPath
		if err != nil {
			fmt.Println("Error :", err)
			os.Exit(1)
		}
	}

	hangman := hangman{
		7,
		randomWord,
		make([]string, 0),
		make([]string, 0),
	}

	return &hangman, usedWordlist
}

func getDefaultWordlist() (string, error) {
	// Get the user home directory.
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	// Check the wordlist exists.
	wordlist := configDir + "/hangman/words.txt"
	if _, err := os.Stat(wordlist); err != nil {
		return "", err
	}

	return wordlist, nil
}

// getGoalWord get a random word in a given wordlist.
func getGoalWord(filePath string) (string, error) {

	// Open the wordlist.
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Split the content into a slice.
	contentSlice := strings.Split(string(content), "\n")

	// Return a random word.
	random := rand.Intn(len(contentSlice) - 1)
	return contentSlice[random], nil
}

// printRightLetters returns the found letters in the goal word and print dots instead of the secret letters.
func (h *hangman) printRightLetters() string {

	var output string

	for _, i := range h.GoalWord {

		if strings.Contains(strings.Join(h.LettersDone, " "), string(i)) {
			output = output + string(i)
		} else {
			output = output + "•"
		}

	}

	return output
}

// printFullWord returns the full word with the missing letters in red.
func (h *hangman) printFullWord() string {

	var output string

	for _, i := range h.GoalWord {

		if strings.Contains(strings.Join(h.LettersDone, " "), string(i)) {
			output = output + string(i)
		} else {
			output = output + color.Red() + string(i) + color.Reset()
		}

	}

	return output
}

// TryLetter compares a letter to the actual word and return the result.
func (h *hangman) tryLetter(letter string) bool {

	h.LettersDone = append(h.LettersDone, letter)

	if strings.Contains(h.GoalWord, letter) {
		h.RightLetters = append(h.RightLetters, letter)
		return true
	}

	h.TriesLeft--
	return false
}

// IsWordFound checks if the player has found the entire word.
func (h *hangman) isWordFound() bool {
	for _, i := range h.GoalWord {
		if strings.Contains(strings.Join(h.LettersDone, " "), string(i)) {
			continue
		} else {
			return false
		}
	}
	return true
}
