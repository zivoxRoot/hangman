package main

import (
	"fmt"
	"os"

	"github.com/zivoxRoot/hangman/internal/hangman"
	"github.com/zivoxRoot/hangman/internal/help"
)

var config = hangman.Config{
	ShowOneLetter: false,
	ShowBanner:    true,
	ShowHints:     true,
	ShowColors:    true,
}

func main() {
	args := os.Args

	if len(args) == 1 {
		hangman.PlayGame("", config)
	}

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			fmt.Println(help.DefaultHelp())
		} else if args[1] == "-v" || args[1] == "--version" {
			fmt.Println(help.PrintVersion())
		} else {
			fmt.Println(help.UnknownCommand(args[1]))
		}
	}

	if len(args) == 3 {
		if args[1] == "-w" || args[1] == "--wordlist" {
			hangman.PlayGame(args[2], config)
		} else {
			fmt.Println(help.UnknownCommand(args[1]))
		}
	}
}
