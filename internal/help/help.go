package help

// DefaultHelp returns the default help message.
func DefaultHelp() string {
	helpMessage := `
hangman - A hangman game right in your terminal

USAGE:
    hangman [OPTIONS]
    hangman                             start a new hangman game using the default wordlist
    hangman -w /path/to/wordlist.txt    start a new hangman game using your own wordlist

OPTIONS:
    -h, --help                          display this help message and exit
    -v, --version                       display the version of this program
	`
	return helpMessage
}

// PrintVersion returns the version of this program.
func PrintVersion() string {
	return "version 0.0.0"
}

// UnknownCommand notifies the user that the entered command doesn't exist.
func UnknownCommand(command string) string {
	return "Unknown command :" + command + "\nTry 'hangman --help' for more infos"
}
