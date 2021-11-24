package main

import "fmt"

func main() {
	wordLen := Reader()
	fmt.Println(wordLen)
	fmt.Println(len(wordLen))
	hider := "_ "
	for i := 0; i < len(wordLen); i++ {
		fmt.Print(hider)
	}
	inputLetter(wordLen)
}

func inputLetter(word string) {
	var letter string
	fmt.Println("\nEntrez une lettre :")
	fmt.Scanln(&letter)
	if len(letter) > 1 {
		fmt.Println("Impossible, entrez une seule lettre.")
		inputLetter(word)
	} else {

		for i := 0; i < len(word); i++ {

			if letter == string(word[i]) {
				fmt.Println("La lettre est dans le mot")
				break
			} else {
				fmt.Println("La lettre n'est pas dans le mot")
				break
			}
		}

	}
}
