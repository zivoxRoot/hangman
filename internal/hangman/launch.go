package hangman

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/zivoxRoot/hangman/internal/colorer"
)

// Player represents a player instance.
type Player struct {
	status      string
	playerQuits bool
	needHelp    bool
}

// Declare and initialize necessary values.
var color = colorer.NewColorer()
var player = Player{
	status:      "",
	playerQuits: false,
	needHelp:    false,
}

// PlayGame launches and handles a new game.
func PlayGame(wordlistPath string) {

	// Initialize a hangman.
	var hangman, usedWordlist = newHangman(wordlistPath)
	player.status = color.Green() + "Using a random word from : " + usedWordlist + color.Reset()

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
			player.playerQuits = true
		} else if choice == "help" {
			player.needHelp = true
		} else if len(choice) == 1 {

			// Check if the letter was already entered.
			if strings.Contains(strings.Join(hangman.LettersDone, " "), choice) {
				player.status = color.Green() + "Are you stupid ? You entered '" + choice + "' twice..." + color.Reset()
				continue
			}

			result := hangman.tryLetter(choice)

			if result {
				player.status = color.Cyan() + choice + " -> success!" + color.Reset()
			} else {
				player.status = color.Red() + choice + " -> fail..." + color.Reset()
			}

		} else {
			player.status = color.Green() + "You should only enter 1 letter" + color.Reset()
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
	if player.needHelp {
		player.status = printHelp()
		player.needHelp = false
	}

	// Check if the player wants to quit.
	if player.playerQuits {
		player.status = color.Green() + "See you later :)" + color.Reset()
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
	fmt.Printf("\n%v\n", player.status)
	player.status = ""
}
