package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getRune(prompt string) (c rune) {
	scanner := bufio.NewScanner(os.Stdin)
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
	for {
		char := getRune("Enter a character: ")
		fmt.Printf("You entered: %c\n", char)
	}
}
