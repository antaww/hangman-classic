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
	randomNumber := rand.Intn(36)
	i := 0
	for scanner.Scan() {
		if i == randomNumber {
			word = scanner.Text()
		}
		i++
	}
	return word
}
