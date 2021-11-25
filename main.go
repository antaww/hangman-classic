package main

import (
	"fmt"
)

func main() {
	wordLen := Reader()
	fmt.Println(wordLen)      //pour les tests
	fmt.Println(len(wordLen)) //pour les tests
	hider(wordLen)            //affiche le mot sous forme cachée
	inputLetter(wordLen)      //demande d'input lettre
}

func inputLetter(word string) {
	var letter string
	fmt.Println("\nEntrez une lettre :")
	fmt.Scanln(&letter)
	if len(letter) > 1 {
		fmt.Println("Impossible, entrez une seule lettre.")
		inputLetter(word)
	} else {
		inWord := false
		for i := 0; i < len(word); i++ {

			if letter == string(word[i]) {
				inWord = true
			}
		}
		if inWord {
			fmt.Println("La lettre est dans le mot")
		} else {
			fmt.Println("La lettre n'est pas dans le mot")
		}

	}
}

func hider(word string) {
	var hidden []string
	for i := 0; i < len(word); i++ {
		hidden = append(hidden, "_")
	}
	fmt.Println(hidden)
}

func Replace(word string) {

}
