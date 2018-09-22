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

func getMovementComponent(scanner *bufio.Scanner, prompt string, min rune, max rune) rune {
	for {
		r := getRune(scanner, prompt)
		if r >= min && r <= max {
			return r
		}
		fmt.Printf("You must enter a letter from %c to %c\n", min, max)
	}
}

func main() {
	board := NewBoard()
	clearScreen := "\x1b[2J\x1b[1;1H"
	fmt.Printf("%s\n%s\n", clearScreen, board)
	scanner := bufio.NewScanner(os.Stdin)
	for board.VictoryState() == Ongoing {
		numMovements := board.NumLegalMovements()
		if numMovements == 1 {
			fmt.Print("There is 1 legal move.\n")
		} else {
			fmt.Printf("There are %d legal moves.\n\n", board.NumLegalMovements())
		}
		origin := getMovementComponent(
			scanner,
			"Enter position to move FROM (labelled e-t): ",
			MinMovementOrigin,
			MaxMovementOrigin)
		destination := getMovementComponent(
			scanner,
			"\nEnter position to move TO (labelled a-m): ",
			MinMovementDestination,
			MaxMovementDestination)
		movement := Movement{origin, destination}
		if board.Permits(movement) {
			board.Execute(movement)
			fmt.Printf("%s\n%s\n", clearScreen, board)
		} else {
			fmt.Println("That move is not permitted, try again!")
		}
	}
	switch board.VictoryState() {
	case Won:
		fmt.Printf("%s\n%s\nYou won! Hooray!\n", clearScreen, board)
	case Lost:
		fmt.Printf("%s\n%s\nThere are no legal moves left. You lost :(\n", clearScreen, board)
	}
}
