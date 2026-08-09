package checker

import "fmt"
import "mastermind/shared"
import "sort"

func CheckGuess(guess string, answer string) shared.Score {
	if guess == answer {
		return shared.Score{
			CorrectAndOutOfPosition: 0,
			CorrectAndInPossition:   4,
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
		CorrectAndInPossition:   numInPosition,
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
	numOutOfPosition := 0
	hideRune := rune(-1)

	// Remove all the in position matches.
	for i := 0; i < len(answer); i += 1 {
		if answer[i] == guess[i] {
			answerCopy[i] = hideRune
		}
	}

	for i := 0; i < len(answer); i += 1 {
		index, found := sort.Find(len(answerCopy), func(j int) int {
			if answerCopy[j] == guess[i] {
				return j
			} else {
				return -1
			}
		})

		if index > 0 || found {
			numOutOfPosition += 1
			answerCopy[index] = hideRune
		}
	}

	return numOutOfPosition
}
