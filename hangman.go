package main

import (
	"fmt"
)

func getSecretWord(wordFile string) string {
	// words, err:=os.ReadFile(wordFile)
	// if err!=nil{
	// 	fmt.Println("Error opening file: %v",err)
	// }
	return "elephant"

}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
