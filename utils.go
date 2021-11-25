package main

import (
	"bufio"
	"math/rand"
	"os"
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
