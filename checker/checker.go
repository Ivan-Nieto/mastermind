package checker

import "mastermind/shared"

func CheckGuess(guess int, answer int) shared.Score {

	if (guess == answer) {
		return shared.Score{
			CorrectAndOutOfPosition: 0,
			CorrectAndInPossition: 4,
			CorrectGuess: true, 
		}
	}

	return shared.Score{
		CorrectAndOutOfPosition: 0,
		CorrectAndInPossition: 0,
		CorrectGuess: false, 
	}
}

