# Hangman

A hangman game right in your terminal !

## Features

- Take a random word from a wordlist
- Choose between the default wordlist or give your own
- Automatic color support for 16 colors, 256 colors, and truecolor

## Installation

1. Installation script :

> Be careful when downloading scripts from the internet, always read them before ! Read the script [here](https://github.com/zivoxRoot/hangman/blob/main/install.sh)

```Bash
curl -L https://raw.githubusercontent.com/zivoxRoot/hangman/refs/heads/main/install.sh | bash
```

2. Build it from source :

> Dependency : you will need Go to compile the program, [Golang](https://go.dev/)

> if you don't have Go installed, you can copy the already compiled binary and move it to its proper place :

With golang installed :

```bash
git clone https://github.com/zivoxRoot/hangman.git
cd hangman
go build -o ./hangman ./cmd/main.go
sudo mv hangman /usr/local/bin
```

Without golang installed :

```bash
git clone https://github.com/zivoxRoot/hangman.git
cd hangman
sudo mv bin/linux/hangman /usr/local/bin
```

**Copy the default wordlist :**

```Bash
mkdir -p $HOME/.config/hangman
cp assets/words.txt $HOME/.config/hangman/
```

## Usage

Start the game :

```bash
hangman
```

Start the game with your own wordlist :

```Bash
hangman -w /path/to/your/wordlist.txt
```

