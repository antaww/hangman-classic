package main

import "fmt"

func Load() {
	fmt.Println(" _    _          _   _  _____ __  __          _   _ ")
	fmt.Println("| |  | |   /\\   | \\ | |/ ____|  \\/  |   /\\   | \\ | |")
	fmt.Println("| |__| |  /  \\  |  \\| | |  __| \\  / |  /  \\  |  \\| |")
	fmt.Println("|  __  | / /\\ \\ | . ` | | |_ | |\\/| | / /\\ \\ | . ` |")
	fmt.Println("| |  | |/ ____ \\| |\\  | |__| | |  | |/ ____ \\| |\\  |")
	fmt.Println("|_|  |_/_/    \\_\\_| \\_|\\_____|_|  |_/_/    \\_\\_| \\_|")
}

func Rules() {
	fmt.Println("Bienvenu dans le Hangman !")
	fmt.Println("Les règles sont celles d'un pendu classique,")
	fmt.Println("à l'exception de quelques petits détails :")
	fmt.Println("\t- Un système d'essais limités est implémenté (en plus du dessin classique)")
	fmt.Println("\t- Un mot entier peut être proposé, en cas d'échec vous perdrez 2 essais")
	fmt.Println("\t- Si vous entrez deux fois la même réponse, vous ne risquez rien")
	fmt.Println("\t- Un mot PEUT comporter un chiffre (rare), par conséquent ceux-ci ne sont pas désactivés")
	fmt.Println("\t- Pas toutes les lettres similaires sont proposées au début du jeu")
}
