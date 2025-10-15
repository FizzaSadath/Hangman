package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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
		if word == strings.ToLower(word) {
			allowedWords = append(allowedWords, word)
		}
	}
	//fmt.Println(len(allowedWords))
	randomNo := rand.Intn(len(allowedWords))
	//fmt.Println(randomNo)
	return allowedWords[randomNo]
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
