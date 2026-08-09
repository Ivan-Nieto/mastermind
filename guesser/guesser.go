package guesser

import "mastermind/shared"
import (
	"fmt"
	"log"
	"strconv"
)

var prevGuesses = []string{}
var confirmedPosition = []int{9, 9, 9, 9}

func MakeGuess(prevResponse shared.Score) string {
	guess := getNewGuessPair()

	fmt.Printf("Making guess '%s', numPreviousGuesses=%d\n", guess, len(prevGuesses))
	prevGuesses = append(prevGuesses, guess)
	return guess
}

func getNewGuessPair() string {
	isFirstGuess := len(prevGuesses) == 0
	if isFirstGuess {
		return "0011"
	}

	lastGuess := prevGuesses[len(prevGuesses)-1]
	wasInitialPairGuess := getWasInitialPairGuess()
	lastGuessLastRune, err := strconv.Atoi(string(lastGuess[3]))

	if err != nil {
		log.Fatal("Failed to convert last guess rune to int")
	}

	if wasInitialPairGuess {
		return fmt.Sprintf("%d%d%d%d", lastGuessLastRune, lastGuessLastRune, lastGuessLastRune, lastGuessLastRune)
	}

	a := lastGuessLastRune + 1
	b := lastGuessLastRune + 2
	s := fmt.Sprintf("%d%d%d%d", a, a, b, b)

	return s
}

func getWasInitialPairGuess() bool {
	lastGuess := prevGuesses[len(prevGuesses)-1]
	wasInitialPairGuess := lastGuess[0] != lastGuess[3]
	wasInitialPairGuess = wasInitialPairGuess && lastGuess[0] == lastGuess[1]
	return wasInitialPairGuess && lastGuess[2] == lastGuess[3]
}
