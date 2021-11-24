package main

import "fmt"


wordLen := Reader()

func main() {

	fmt.Println(wordLen)
	fmt.Println(len(wordLen))
	hider := "_ "
	for i := 0; i < len(wordLen); i++ {
		fmt.Print(hider)
	}
}

func inputLetter() {
	var letter string
	fmt.Scanln(letter)
	if len(letter) > 1 {
		fmt.Println("Impossible, entrez une seule lettre.")
	} else {
		for i := 0; i < len(wordLen)-1; i++ {
			if letter == wordLen[i] {
				fmt.Println("La lettre est dans le mot")
				break
			} else {
				("La lettre n'est pas dans le mot")
			}
		}
	
	}
}

