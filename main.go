package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func getRune(scanner *bufio.Scanner, prompt string) (c rune) {
	fmt.Printf(prompt)
	for scanner.Scan() {
		input := scanner.Bytes()
		if utf8.RuneCount(input) != 1 {
			fmt.Println("Please enter a single character.")
			fmt.Printf(prompt)
			continue
		}
		c, _ = utf8.DecodeRune(input)
		break
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		char := getRune(scanner, "Enter a character: ")
		fmt.Printf("You entered: %c\n", char)
	}
}
