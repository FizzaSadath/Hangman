package main

import (
	"bufio"
	"fmt"
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
	fmt.Println(allowedWords)
	return "elephant"
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
