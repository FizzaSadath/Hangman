package main

import (
	"fmt"
	"unicode"
)

func isAllLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func getSecretWord(wordFileName string) string {
	// var allowedWords []string
	// wordFile, err := os.Open(wordFileName)
	// if err != nil {
	// 	errMessage := fmt.Sprintf("Can't open file %s : %v\n", wordFileName, err)
	// 	panic(errMessage)
	// }
	// defer wordFile.Close()
	// scanner := bufio.NewScanner(wordFile)
	// for scanner.Scan() {
	// 	word := scanner.Text()
	// 	if word == strings.ToLower(word) && isAllLetters(word) {
	// 		allowedWords = append(allowedWords, word)
	// 	}
	// }
	// randomNo := rand.Intn(len(allowedWords))
	//return allowedWords[randomNo]
	return "name's"
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
