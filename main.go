package main

import (
	"fmt"
	"math/rand"
	"time"
)

var hider []string
var used []string

func main() {
	Clear()
	Load()
	Input()
	Clear()
	tries := 0
	wordLen := Reader()
	hiderr(wordLen)
	Reveal(wordLen)
	// fmt.Println(wordLen)      //pour les tests
	// fmt.Println(len(wordLen)) //pour les tests

	inputLetter(wordLen, tries)
}

func inputLetter(word string, try int) {
	Hang(try)
	var letter string
	fmt.Println(letter)

	fmt.Println("Lettres utilisées : ", used)
	fmt.Println("\nEntrez une lettre :")
	fmt.Scanln(&letter)
	used = append(used, letter)

	if letter == word {
		fmt.Println("Bravo vous avez réussi !")
		Input()
		Close()
	}
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
			fmt.Printf("La lettre %v est dans le mot\n", letter)
			for index, char := range word {
				if letter == string(char) {
					hider[index] = letter
				}
			}
		} else {
			fmt.Printf("La lettre %v n'est pas dans le mot\n", letter)
			if try < 9 {
				try++
				fmt.Printf("Il ne vous reste plus que %v essais \n", 10-try)
			} else {
				fmt.Println("Nombre d'essais insuffisant...")
				fmt.Printf("Le mot était %v\n", word)
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
	hider = Reveal(word)
	fmt.Println(hider)

}

func Reveal(word string) []string {
	random := len(word)/2 - 1
	for i := 0; i < random; i++ {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(len(word) - 1)
		hider[randomNumber] = string(word[randomNumber])
	}
	return hider
}
