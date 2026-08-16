package main

import "mastermind/shared"
import "mastermind/guesser"
import "mastermind/checker"
import "fmt"

func main() {
	maxGuesses := 10
	answer := "0123"
	guess := guesser.MakeGuess(shared.Score{})
	currScore := checker.CheckGuess(guess, answer)

	i := 2
	for i := 2; i < maxGuesses; i += 1 {
		if currScore.CorrectGuess {
			break
		}

		guess = guesser.MakeGuess(currScore)
		currScore = checker.CheckGuess(guess, answer)
	}

	if currScore.CorrectGuess {
		fmt.Printf("The correct number '%s' was guessed in %d guesses\n", answer, i)
	} else {
		fmt.Printf("The number '%s' was not found :(\n", answer)
	}
}
