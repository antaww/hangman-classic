package main

import (
	"bufio"
	"fmt"
)

var hider []string
var BR *bufio.Reader

func main() {
	Clear()
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
			if try > 1 {
				try--
				fmt.Printf("Il ne vous reste plus que %v essais \n", try)
			} else {
				fmt.Println("Nombre d'essais insuffisant...")
				fmt.Println("Fermeture du jeu...")
				Close()
			}

		}
		Input()
		Clear()
		fmt.Println(hider)
	}
	stillHider := false
	for index, _ := range word {
		if hider[index] == "_" {
			stillHider = true
			inputLetter(word, try)
			break
		}
	}
	if stillHider == false {
		fmt.Println("Bravo vous avez réussi !")
		Input()
		Close()
	}

}

func hiderr(word string) {
	for i := 0; i < len(word); i++ {
		hider = append(hider, "_")
	}
	fmt.Println(hider)
}
