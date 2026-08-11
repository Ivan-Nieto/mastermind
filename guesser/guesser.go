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

type SuspectedAnswerDigit struct {
	guess    string
	position []int
}

var suspectedAnswers = []SuspectedAnswerDigit{}
var prevGuesses = []PrevGuess{}

func saveScore(score shared.Score) {
	if len(prevGuesses) > 0 {
		prevGuesses[len(prevGuesses)-1].score = score
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

	if len(prevGuesses) < 2 || score.CorrectAndOutOfPosition == 0 && score.CorrectAndInPossition == 0 {
		// Nothing to do; the previous answer didn't get any right
		return
	}

	numPrevGuesses := len(prevGuesses)
	previousGuess := prevGuesses[numPrevGuesses-1]
	previousGuessStr := previousGuess.guess
	lastGuessWasSingleDigit := true
	for x := range previousGuessStr {
		if previousGuessStr[x] != previousGuessStr[0] {
			lastGuessWasSingleDigit = false
		}
	}

	lastScore := prevGuesses[numPrevGuesses-2].score
	lastScoreOutOfPos := lastScore.CorrectAndOutOfPosition
	lastScoreInPos := lastScore.CorrectAndInPosition
	currScoreOutOfPos := score.CorrectAndOutOfPosition
	lastCombinedScore := lastScoreOutOfPosition + lastScoreInPos
	currScoreInPos := score.CorrectAndInPosition

	if lastCombinedScore > currScoreInPos {
		// Is a in the answer once?
		if lastCombinedScore > 1 && currScoreInPosition <= 3 {
		}
		// Is a in the answer twice?
		if lastCombinedScore > 2 && currScoreInPosition <= 2 {
		}
		// Is b in the answer once?
		if currScoreInPos == 1 {
		}
		// Is b in the answer twice?
		if currScoreInPos == 2 {
		}
		// Is b in the answer thrice?
		if currScoreInPos == 3 {
		}
	}

	// b is in the answer once or twice
	if lastCombinedScore == currScoreinPos {
		// b was on the left if lastScoreOutOfPos times
		// b was on the right lastScoreInPos times

		allOnTheRight := lastScoreInPos == currScoreInPos && lastScoreOutOfPos == 0
		allOnTheLeft := lastScoreOutOfPos == currScoreInPos && lastScoreInPos == 0
		if allOnTheRight {
			suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(previousGuessStr[0]), []int{2, 3}})
			suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(previousGuessStr[0]), []int{2, 3}})
		}
		if allOnTheLeft {
			suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(previousGuessStr[0]), []int{0, 1}})
			suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(previousGuessStr[0]), []int{0, 1}})
		}
		if !allOnTheLeft && !allOnTheRight {
			if lastScoreOutOfPos {
				suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(previousGuessStr[0]), []int{0, 1}})
				suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(previousGuessStr[0]), []int{2, 3}})
			} else {
				suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(previousGuessStr[0]), []int{2, 3}})
				suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(previousGuessStr[0]), []int{0, 1}})
			}
		}
		if currScoreInPos == 1 {
			// TODO: Remove the last item that was added to the suspectedAnswers slice.
		}
	}

	// b is on the left at least once
	if lastCombinedScore < currScoreInPos {
		if lastCombinedScore == 1 {
			if lastScoreOutOfPos == 1 {
				//bbxx
				// if currScoreInPos == 3 then it's bbxb or bbbx
			} else {
				// xxbb
				// if currScoreInPos == 3 then it's bxbb or xbbb
			}
		}

		// TODO: Figure out if that 'x' should be an 'a'.
		if lastCombinedScore == 2 {
			// currScoreInPos must be 3
			// lastScoreInPos == 0 is not possible if currScoreInPos == 3
			if lastScoreInPos == 1 { // lastScoreOutOfPos == 1
				// Answer could be or bbxb or bbbx
			}
			if lastScoreInPos == 2 { // lastScoreOutOfPos == 0
				// Answer could be xbbb or bxbb
			}
		}
		// lastCombinedScore == 3 is not possible; currScoreInPos would need to be 4
	}
}

func getNewGuessPair() string {
	isFirstGuess := len(prevGuesses) == 0
	if isFirstGuess {
		return "0011"
	}

	lastGuess := prevGuesses[len(prevGuesses)-1].guess
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
