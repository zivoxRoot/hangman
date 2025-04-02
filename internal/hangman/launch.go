package hangman

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/zivoxRoot/hangman/internal/colorer"
)

// Declare and initialize necessary values.
var color = colorer.NewColorer()
var status string
var playerQuits bool
var needHelp bool

// PlayGame launches and handles a new game.
func PlayGame(wordlistPath string) {

	// Initialize a hangman.
	var hangman, usedWordlist = newHangman(wordlistPath)
	status = color.Green() + "Using a random word from : " + usedWordlist + color.Reset()

	// Main loop.
	for {

		// Update the front-end.
		updateFrontend(hangman)

		// Prompt the user and get his choice.
		fmt.Print("\n> ")

		var choice string
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Error getting the user input :", err)
			os.Exit(1)
		}

		// Check the user input.
		if choice == "quit" || choice == "exit" {
			playerQuits = true
		} else if choice == "help" {
			needHelp = true
		} else if len(choice) == 1 {

			result := hangman.tryLetter(choice)

			if result {
				status = color.Cyan() + choice + " -> success!" + color.Reset()
			} else {
				status = color.Red() + choice + " -> fail..." + color.Reset()
			}

		} else {
			status = color.Green() + "You should only enter 1 letter" + color.Reset()
		}
	}
}

func updateFrontend(hangman *hangman) {

	// Clear the screen.
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	// Greet and display hints.
	fmt.Println(greet())
	fmt.Println(hints())

	// Print the hangman.
	fmt.Println(printHangmanState(hangman.TriesLeft))

	// Print the partially found word.
	if hangman.TriesLeft == 0 {
		fmt.Printf("\n%vWord%v : %v\n\n", color.Bold(), color.Reset(), hangman.printFullWord())
	} else {
		fmt.Printf("\n%vWord%v : %v\n\n", color.Bold(), color.Reset(), hangman.printRightLetters())
	}

	// Print the already found letters.
	fmt.Printf("%vLetters tried%v :", color.Bold(), color.Reset())
	for _, i := range hangman.LettersDone {

		if strings.Contains(hangman.GoalWord, i) {
			fmt.Printf(" %v%v%v", color.Cyan(), i, color.Reset())
		} else {
			fmt.Print(" ", i)
		}
	}
	fmt.Print("\n")

	// Check if the player need help.
	if needHelp {
		status = printHelp()
		needHelp = false
	}

	// Check if the player wants to quit.
	if playerQuits {
		status = color.Green() + "See you later :)" + color.Reset()
		defer os.Exit(0)
	}

	// Check if the player won the game.
	if hangman.isWordFound() {
		fmt.Printf("\n%vYou won, congratulations !%v\n", color.Cyan(), color.Reset())
		os.Exit(0)
	}

	// Check if the player lost the game.
	if hangman.TriesLeft == 0 {
		fmt.Printf("\n%vYou unfortunately lost...%v\n", color.Red(), color.Reset())
		os.Exit(0)
	}

	// Print the status and reset it.
	fmt.Printf("\n%v\n", status)
	status = ""
}
