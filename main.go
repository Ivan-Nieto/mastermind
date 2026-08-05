package main

import "mastermind/shared"
import "mastermind/guesser"
import "mastermind/checker"
import "fmt"

func main() {

	maxGuesses := 20
	tempCode := 1010
	guess := guesser.MakeGuess(shared.Score{})
	currScore := checker.CheckGuess(guess, tempCode)
	
	i := 1
	for i := 1; i < maxGuesses; i += 1 {
    if (currScore.CorrectGuess) {
			break
		}

		guess = guesser.MakeGuess(currScore)
		currScore = checker.CheckGuess(guess, tempCode)
	}

	if (currScore.CorrectGuess) {
		fmt.Printf("The correct number '%d' was guessed in %d guesses\n", tempCode, i)
	} else {
		fmt.Printf("The number'%d' was not found :(\n", tempCode)
	}
}

