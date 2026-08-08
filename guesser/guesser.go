package guesser

import "mastermind/shared"
import "fmt"

func MakeGuess(prevResponse shared.Score) int {

	guess := 1001

	fmt.Printf("Making guess %d\n", guess)
	return guess
}
