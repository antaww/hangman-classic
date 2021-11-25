package main

import (
	"bufio"
	"fmt"
	"os"
)

func Hang(try int) {
	file, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(file)
	startLine := try*7 + 1 + 1*try
	endLine := (try+1)*7 + try
	i := 1
	for scanner.Scan() {
		if i >= startLine && i <= endLine {
			fmt.Println(scanner.Text())
		}
		i++
	}
	fmt.Println()
}
