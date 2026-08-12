package main

import "mastermind/shared"
import "mastermind/guesser"
import "mastermind/checker"
import "fmt"

func main() {
	maxGuesses := 20
	answer := "1100"
	guess := guesser.MakeGuess(shared.Score{})
	currScore := checker.CheckGuess(guess, answer)

	i := 1
	for i := 1; i < maxGuesses; i += 1 {
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
