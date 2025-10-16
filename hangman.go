package main

import (
	"bufio"
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

	return state
}
func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
	fmt.Println(NewGame(getSecretWord("/usr/share/dict/words")))
}
