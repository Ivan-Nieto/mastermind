package guesser

import "mastermind/shared"
import (
	"fmt"
	"log"
	"strconv"
)

type PrevGuess struct {
	guess string
	score shared.Score
}


var prevGuesses = []PrevGuess{}

func saveScore(score shared.Score) {
	if len(prevGuesses) > 0 {
		prevGuesses[len(prevGuesses) - 1].score = score
	}
}

// var prevGuesses = []string{}
var confirmedPosition = []int{9, 9, 9, 9}

func MakeGuess(prevResponseScore shared.Score) string {
	saveScore(prevResponseScore)
	guess := getNewGuessPair()

	checkLastGuessResponse(prevResponseScore)

	fmt.Printf("Making guess '%s', numPreviousGuessesu=%d\n", guess, len(prevGuesses))
	prevGuesses = append(prevGuesses, PrevGuess{guess, shared.EmptyScore})
	return guess
}

func checkLastGuessResponse(score shared.Score) {
	if score.CorrectGuess {
		log.Fatal("New guess was requested but the last guess was correct")
	}

	if score.CorrectAndOutOfPosition == 0 && score.CorrectAndInPossition == 0 {
		// Nothing to do; the previous answer didn't get any right
		return
	}

	numPrevGuesses := len(prevGuesses)
	previousGuess := prevGuesses[numPrevGuesses - 1]
	previousGuessStr := previousGuess.guess
  lastGuessWasSingleDigit := true
	for x := range previousGuessStr {
		if previousGuessStr[x] != previousGuessStr[0] {
			lastGuessWasSingleDigit = false
		}
	}

	lastScore := prevGuesses[numPrevGuesses - 2].score
	lastScoreOutOfPos := lastScore.CorrectAndOutOfPosition
	lastScoreInPos := lastScore.CorrectAndInPosition
	currScoreOutOfPos := score.CorrectAndOutOfPosition
	currScoreInPos := score.CorrectAndInPosition

	if currScoreOutOfPos > 0 {
		// Might not be necessary to check len > 2 given single
		// digit guesses are never the first guess.
		if lastGuessWasSingleDigit && numPrevGuesses > 2 {
			if lastScoreOutOfPos == 1) {
				// Means one of these digits is on the left hand side.
				// If currScoreOutOfPos is = 2; that means the answer is bbxx
				if currScoreOutOfPos == 2 {
					confirmedPosition[0] = previousGuessStr[0]
					confirmedPosition[1] = previousGuessStr[0]
				} else {
					// We don't know exactly which side just that it's on the first two digits
				}
			} else if lastScoreOutOfPos === 2 {
				// Means the first two digits of the answer are the previous guesses digit;
				// If the currentScoreOutOfPos is 3 then the is something like bbax, bbxa, aabx, aaxb
			} else if lastScoreOutOfPos === 3 {
				// Means one side of the aabb guess is on the wrong side; check the curretn score
				// to figure out which one it was (a or b). The other has a single digit on that
				// side, i.e the answer is one of the following; bbax, bbxa, aabx, aaxb
			}
		}
	}
}

func getNewGuessPair() string {
	isFirstGuess := len(prevGuesses) == 0
	if isFirstGuess {
		return "0011"
	}

	lastGuess := prevGuesses[len(prevGuesses) - 1].guess
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
	lastGuess := prevGuesses[len(prevGuesses)-1].guess
	wasInitialPairGuess := lastGuess[0] != lastGuess[3]
	wasInitialPairGuess = wasInitialPairGuess && lastGuess[0] == lastGuess[1]
	return wasInitialPairGuess && lastGuess[2] == lastGuess[3]
}
