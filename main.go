package main

import (
	"fmt"
)

var hider []string

func main() {
	tries := 10
	wordLen := Reader()
	hiderr(wordLen)
	fmt.Println(wordLen)        //pour les tests
	fmt.Println(len(wordLen))   //pour les tests
	inputLetter(wordLen, tries) //demande d'input lettre
}

func inputLetter(word string, try int) {
	var letter string
	fmt.Println("\nEntrez une lettre :")
	fmt.Scanln(&letter)
	if len(letter) > 1 {
		fmt.Println("Impossible, entrez une seule lettre.")
		inputLetter(word, try)
	} else {
		inWord := false
		for i := 0; i < len(word); i++ {

			if letter == string(word[i]) {
				inWord = true
			}
		}
		if inWord {
			fmt.Println("La lettre est dans le mot")
			for index, char := range word {
				if letter == string(char) {
					hider[index] = letter
				}
			}
		} else {
			fmt.Printf("La lettre %v n'est pas dans le mot\n", letter)
			if try > 0 {
				try--
				fmt.Printf("Il ne vous reste plus que %v essais \n", try)
			}

		}
		fmt.Println(hider)
	}
	inputLetter(word, try)
}

func hiderr(word string) {
	for i := 0; i < len(word); i++ {
		hider = append(hider, "_")
	}
	fmt.Println(hider)
}
