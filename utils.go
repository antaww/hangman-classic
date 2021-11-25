package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Reader() string {

	var word string
	file, _ := os.Open("words.txt")
	scanner := bufio.NewScanner(file)
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(Counter())
	i := 0
	for scanner.Scan() {
		if i == randomNumber {
			word = scanner.Text()
		}
		i++
	}
	return word
}

func Counter() int {
	file, _ := os.Open("words.txt")
	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() { //while \n available
		count++
	}
	return count
}

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func Input() {
	var input string
	fmt.Println("Appuyez sur entrée pour continuer...")
	fmt.Scanln(&input)
	if input == "\n" {
		strings.Replace(input, "\r\n", "", -1)
	}
}

func Close() {
	os.Exit(0)
}

// func Version() string {
// 	Clear()
// 	var open string
// 	var a int
// 	fmt.Println("1 - Version classique")
// 	fmt.Println("2 - Version histoire")
// 	fmt.Scanln(&a)
// 	if a == 1 {
// 		open = "words.txt"
// 	} else if a == 2 {
// 		open = "histoire.txt"
// 	} else {
// 		fmt.Println("Veuillez entrer un choix valable.")
// 		Input()
// 		Version()
// 	}
// 	return open
// }
