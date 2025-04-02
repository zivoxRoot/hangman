package hangman

// greet returns the hangman greeter.
func greet() string {
return `
 __    __
| $$  | $$  ______   _______    ______   ______ ____    ______   _______
| $$  | $$  ______   _______    ______   ______ ____    ______   _______
| $$    $$  \$$$$$$\| $$$$$$$\|  $$$$$$\| $$$$$$\$$$$\  \$$$$$$\| $$$$$$$\
| $$$$$$$$ /      $$| $$  | $$| $$  | $$| $$ | $$ | $$ /      $$| $$  | $$
| $$  | $$|  $$$$$$$| $$  | $$| $$__| $$| $$ | $$ | $$|  $$$$$$$| $$  | $$
| $$  | $$ \$$    $$| $$  | $$ \$$    $$| $$ | $$ | $$ \$$    $$| $$  | $$
 \$$   \$$  \$$$$$$$ \$$   \$$ _\$$$$$$$ \$$  \$$  \$$  \$$$$$$$ \$$   \$$
                              |  \__| $$
                               \$$    $$
                                \$$$$$$
`
}

// hints returns the hint string.
func hints() string {
	return "Quit with 'quit', get help with 'help'\n"
}

// help returns the in game help message.
func printHelp() string {
	helpMessage := `
You have to find the word with dots '••••' using a letter, if your letter is right
it will appear in the word, else the hangman goes one step further. You win by finding
the hidden word, and lose by having a full hangman.
	`
	return helpMessage
}

// printHangmanState returns the hangman at his current state.
func printHangmanState(state int) string {
	var hangman string

	switch state {
	case 6:
		hangman = `





		┗━━━━━━━
		`
	case 5:
		hangman = `

		┃
		┃
		┃
		┃
		┗━━━━━━━
		`
	case 4:
		hangman = `
		┏━━━━━┓
		┃
		┃
		┃
		┃
		┗━━━━━━━
		`
	case 3:
		hangman = `
		┏━━━━━┓
		┃     ┃
		┃
		┃
		┃
		┗━━━━━━━
		`
	case 2:
		hangman = `
		┏━━━━━┓
		┃     ┃
		┃     O
		┃
		┃
		┗━━━━━━━
		`
	case 1:
		hangman = `
		┏━━━━━┓
		┃     ┃
		┃     O
		┃    /|\
		┃
		┗━━━━━━━
		`
	case 0:
		hangman = `
		┏━━━━━┓
		┃     ┃
		┃     O
		┃    /|\
		┃    / \
		┗━━━━━━━

		GAME OVER !!!
		`
	}
	return hangman
}
