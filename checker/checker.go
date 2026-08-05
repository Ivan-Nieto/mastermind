package checker

import "strconv"
import "mastermind/shared"

func CheckGuess(guess int, answer int) shared.Score {

	if guess == answer {
		return shared.Score{
			CorrectAndOutOfPosition: 0,
			CorrectAndInPossition:   4,
			CorrectGuess:            true,
		}
	}

	answerStr := strconv.Itoa(answer)
	answerRunes := []rune(answerStr)
	guessStr := strconv.Itoa(guess)
	guessRunes := []rune(guessStr)
	numInPosition := CountInPosition(guessRunes, answerRunes)
	numOutOfPosition := CountOutOfPosition(guessRunes, answerRunes)

	return shared.Score{
		CorrectAndOutOfPosition: numOutOfPosition,
		CorrectAndInPossition:   numInPosition,
		CorrectGuess:            numInPosition == len(answerStr) && len(guessStr) == len(answerStr),
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

	return 0
}
