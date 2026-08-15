package checker

import "fmt"
import "mastermind/shared"

func CheckGuess(guess string, answer string) shared.Score {
	if guess == answer {
		return shared.Score{
			CorrectAndOutOfPosition: 0,
			CorrectAndInPosition:    4,
			CorrectGuess:            true,
		}
	}

	answerRunes := []rune(answer)
	guessRunes := []rune(guess)
	numInPosition := CountInPosition(guessRunes, answerRunes)
	numOutOfPosition := CountOutOfPosition(guessRunes, answerRunes)

	fmt.Println("****************************************")
	fmt.Printf("Ran checker for guess='%s' and answer='%s'\n", guess, answer)
	fmt.Printf("    CorrectAndOutOfPosition: %d\n", numOutOfPosition)
	fmt.Printf("    CorrectAndInPosition: %d\n", numInPosition)
	fmt.Println("****************************************\n")

	return shared.Score{
		CorrectAndOutOfPosition: numOutOfPosition,
		CorrectAndInPosition:    numInPosition,
		CorrectGuess:            numInPosition == len(answer) && len(guess) == len(answer),
	}
}

func CountInPosition(guess []rune, answer []rune) int {
	numInPosition := 0
	for i := 0; i < len(answer); i += 1 {
		if guess[i] == answer[i] {
			numInPosition += 1
		}
	}

	return numInPosition
}

func CountOutOfPosition(guess []rune, answer []rune) int {
	answerCopy := answer // This should be a clone!
	guessCopy := guess
	numOutOfPosition := 0
	hideRune := rune(9)

	// Remove all the in position matches.
	for i := 0; i < len(answer); i += 1 {
		if answer[i] == guessCopy[i] {
			answerCopy[i] = '9' // hideRune
		}
	}

	for i := 0; i < len(answerCopy); i += 1 {
		foundIndex := 9
		for j, val := range guessCopy {
			if val == answerCopy[i] && val != hideRune {
				foundIndex = j
				break
			}
		}

		if foundIndex != 9 {
			numOutOfPosition += 1
			guessCopy[foundIndex] = hideRune
		}
	}

	return numOutOfPosition
}
