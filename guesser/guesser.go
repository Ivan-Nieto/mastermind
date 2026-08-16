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

var confirmedPosition = []int{9, 9, 9, 9}

func MakeGuess(prevResponseScore shared.Score) string {
	saveScore(prevResponseScore)
	guess := getNewGuessPair()

	checkLastGuessResponse(prevResponseScore)

	fmt.Print("Current suspected answer: ", suspectedAnswers, "\n")
	fmt.Printf("\nMaking guess '%s', numPreviousGuessesu=%d\n", guess, len(prevGuesses))
	prevGuesses = append(prevGuesses, PrevGuess{guess, shared.EmptyScore})
	return guess
}

func getNewGuessPair() string {
	isFirstGuess := len(prevGuesses) == 0
	if isFirstGuess {
		return "0011"
	}

	lastScore := prevGuesses[len(prevGuesses)-1].score
	lastGuess := prevGuesses[len(prevGuesses)-1].guess
	wasInitialPairGuess := getWasInitialPairGuess()
	lastGuessLastRune, err := strconv.Atoi(string(lastGuess[3]))
	lastGuessTotalyWrong := lastScore.CorrectAndOutOfPosition == 0 && lastScore.CorrectAndInPosition == 0

	if err != nil {
		log.Fatal("Failed to convert last guess rune to int")
	}

	if wasInitialPairGuess && !lastGuessTotalyWrong {
		return fmt.Sprintf("%d%d%d%d", lastGuessLastRune, lastGuessLastRune, lastGuessLastRune, lastGuessLastRune)
	}

	a := lastGuessLastRune + 1
	b := lastGuessLastRune + 2
	s := fmt.Sprintf("%d%d%d%d", a, a, b, b)

	if a > 6 || b > 6 {
		log.Fatal("Was unable to find an answer")
	}

	return s
}

func getWasInitialPairGuess() bool {
	lastGuess := prevGuesses[len(prevGuesses)-1].guess
	wasInitialPairGuess := lastGuess[0] != lastGuess[3]
	wasInitialPairGuess = wasInitialPairGuess && lastGuess[0] == lastGuess[1]
	return wasInitialPairGuess && lastGuess[2] == lastGuess[3]
}
