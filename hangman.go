package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

type Game struct {
	secretWord     string
	chancesLeft    uint
	guesses        []byte
	correctGuesses []byte
}

func NewGame(secretWord string) Game {
	return Game{
		secretWord:     secretWord,
		chancesLeft:    7,
		guesses:        []byte{},
		correctGuesses: []byte{},
	}
}

func isAllLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getSecretWord(wordFileName string) string {
	var allowedWords []string
	wordFile, err := os.Open(wordFileName)
	if err != nil {
		errMessage := fmt.Sprintf("Can't open file %s : %v\n", wordFileName, err)
		panic(errMessage)
	}
	defer wordFile.Close()
	scanner := bufio.NewScanner(wordFile)
	for scanner.Scan() {
		word := scanner.Text()
		if word == strings.ToLower(word) && isAllLetters(word) && len(word) > 5 {
			{
				allowedWords = append(allowedWords, word)
			}
		}

	}
	randomNo := rand.Intn(len(allowedWords))
	return allowedWords[randomNo]
}
func playTurn(state Game, guess byte) Game {
	isContainByte := strings.ContainsRune(state.secretWord, rune(guess))
	isAlreadyGuessed := bytes.Contains(state.guesses, []byte{guess})

	if state.chancesLeft > 0 && isContainByte && !isAlreadyGuessed {
		state = Game{
			secretWord:     state.secretWord,
			chancesLeft:    state.chancesLeft,
			guesses:        append(state.guesses, guess),
			correctGuesses: append(state.correctGuesses, guess),
		}
	}
	if state.chancesLeft > 0 && !isContainByte && !isAlreadyGuessed {
		state = Game{
			secretWord:     state.secretWord,
			chancesLeft:    state.chancesLeft - 1,
			guesses:        append(state.guesses, guess),
			correctGuesses: state.correctGuesses,
		}
	}

	return state
}
func getUserInput() byte {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter a character: ")
		b, _ := reader.ReadByte()
		reader.ReadByte()

		if b >= 'a' && b <= 'z' {
			return b
		}
		fmt.Println("Invalid input! Please enter a lowercase letter (a-z).")
	}
}

func displayProgress(state Game) string {
	var result strings.Builder

	for i := 0; i < len(state.secretWord); i++ {
		letter := state.secretWord[i]
		found := false
		for _, guessed := range state.correctGuesses {
			if guessed == letter {
				found = true
				break
			}
		}
		if found {
			fmt.Fprintf(&result, "%c ", letter)
		} else {
			result.WriteString("_ ")
		}
	}

	return result.String()
}

func hasWon(state Game) bool {
	letters := make(map[byte]bool)
	for i := 0; i < len(state.secretWord); i++ {
		letters[state.secretWord[i]] = true
	}
	for letter := range letters {
		found := false
		for _, guessed := range state.correctGuesses {
			if guessed == letter {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func hasLost(state Game) bool {
	if state.chancesLeft == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	state := NewGame(getSecretWord("/usr/share/dict/words"))
	fmt.Println("Welcome to Hangman!")
	//fmt.Println(state.secretWord)
	fmt.Println(strings.Repeat("_ ", len(state.secretWord)))
	fmt.Println("Chances left: ", state.chancesLeft)
	for state.chancesLeft > 0 {
		guess := getUserInput()
		state = playTurn(state, guess)
		progress := displayProgress(state)
		fmt.Println(progress)
		fmt.Println("Chances left: ", state.chancesLeft)
		fmt.Println("Guessed letters: ", string(state.guesses))
		fmt.Println()
		if hasWon(state) {
			fmt.Println("\nCongratulations! You won")
			break
		}
	}
	if hasLost(state) {
		fmt.Println("You lost! The word was: ", state.secretWord)
	}

}
