package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("words.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
}
