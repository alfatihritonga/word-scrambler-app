package main

import (
	"fmt"

	scrambler "github.com/alfatihritonga/word-scrambler-go"
)

func main() {
	fmt.Print("Input a word: ")

	var input string
	fmt.Scanln(&input)

	if input == "" {
		fmt.Println("The word cannot be empty!")
		return
	}

	scrambled := scrambler.ScrambleWord(input)
	fmt.Printf("The result of the word shuffle : %s\n", scrambled)
}
