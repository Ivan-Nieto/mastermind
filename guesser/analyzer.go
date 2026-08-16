package guesser

import "mastermind/shared"
import "log"

type AnalysisData struct {
	lastScore         shared.Score
	lastScoreOutOfPos int
	lastScoreInPos    int
	lastCombinedScore int
	currScoreInPos    int
	previousGuessStr  string
}

func checkLastGuessResponse(score shared.Score) {
	if score.CorrectGuess {
		log.Fatal("New guess was requested but the last guess was correct")
	}

	if len(prevGuesses) < 2 || score.CorrectAndOutOfPosition == 0 && score.CorrectAndInPosition == 0 {
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

	if !lastGuessWasSingleDigit {
		return
	}

	lastScore := prevGuesses[numPrevGuesses-2].score
	lastScoreOutOfPos := lastScore.CorrectAndOutOfPosition
	lastScoreInPos := lastScore.CorrectAndInPosition
	lastCombinedScore := lastScoreOutOfPos + lastScoreInPos
	currScoreInPos := score.CorrectAndInPosition
	data := AnalysisData{lastScore, lastScoreOutOfPos, lastScoreInPos, lastCombinedScore, currScoreInPos, previousGuessStr}

	if lastCombinedScore > currScoreInPos {
		checkDecreasedScore(data)
	}

	// b is in the answer once or twice
	if lastCombinedScore == currScoreInPos {
		checkNoChangeInScore(data)
	}

	// b is on the left at least once
	if lastCombinedScore < currScoreInPos {
		checkIncreaseInScore(data)
	}
}

func checkDecreasedScore(data AnalysisData) {
	// Is a in the answer once?
	if data.lastCombinedScore > 1 && data.currScoreInPos <= 3 {
		savePossibleAnswer(data.previousGuessStr[0], []int{0, 1})
	}
	// Is a in the answer twice?
	if data.lastCombinedScore > 2 && data.currScoreInPos <= 2 {
		savePossibleAnswer(data.previousGuessStr[0], []int{0, 1})
	}
	// Is b in the answer once?
	if data.currScoreInPos == 1 {
		savePossibleAnswer(data.previousGuessStr[2], []int{0, 1, 2, 3})
	}
	// Is b in the answer twice?
	if data.currScoreInPos == 2 {
		if data.lastScoreInPos >= 2 {
			savePossibleAnswer(data.previousGuessStr[2], []int{2})
			savePossibleAnswer(data.previousGuessStr[2], []int{3})
		} else {
			savePossibleAnswer(data.previousGuessStr[2], []int{0, 1, 2, 3})
			savePossibleAnswer(data.previousGuessStr[2], []int{0, 1, 2, 3})
		}
	}
	// Is b in the answer thrice?
	if data.currScoreInPos == 3 {
		savePossibleAnswer(data.previousGuessStr[0], []int{2, 3})
	}
}

func checkNoChangeInScore(data AnalysisData) {
	// b was on the left if data.lastScoreOutOfPos times
	// b was on the right data.lastScoreInPos times

	allOnTheRight := data.lastScoreInPos == data.currScoreInPos && data.lastScoreOutOfPos == 0
	allOnTheLeft := data.lastScoreOutOfPos == data.currScoreInPos && data.lastScoreInPos == 0
	if allOnTheRight {
		savePossibleAnswer(data.previousGuessStr[0], []int{2, 3})
		savePossibleAnswer(data.previousGuessStr[0], []int{2, 3})
	}
	if allOnTheLeft {
		savePossibleAnswer(data.previousGuessStr[0], []int{0, 1})
		savePossibleAnswer(data.previousGuessStr[0], []int{0, 1})
	}
	if !allOnTheLeft && !allOnTheRight {
		if data.lastScoreOutOfPos > 0 {
			savePossibleAnswer(data.previousGuessStr[0], []int{0, 1})
			savePossibleAnswer(data.previousGuessStr[0], []int{2, 3})
		} else {
			savePossibleAnswer(data.previousGuessStr[0], []int{2, 3})
			savePossibleAnswer(data.previousGuessStr[0], []int{0, 1})
		}
	}
	if data.currScoreInPos == 1 {
		suspectedAnswers = suspectedAnswers[:len(suspectedAnswers)-1]
	}
}

func checkIncreaseInScore(data AnalysisData) {
	if data.lastCombinedScore == 1 {
		if data.lastScoreOutOfPos == 1 {
			//bbxx
			savePossibleAnswer(data.previousGuessStr[2], []int{0, 1})
			// if data.currScoreInPos == 3 then it's bbxb or bbbx
		} else {
			savePossibleAnswer(data.previousGuessStr[2], []int{2, 3})
			// xxbb
			// if data.currScoreInPos == 3 then it's bxbb or xbbb
		}
	}

	// TODO: Figure out if that 'x' should be an 'a'.
	if data.lastCombinedScore == 2 {
		// data.currScoreInPos must be 3
		// data.lastScoreInPos == 0 is not possible if data.currScoreInPos == 3
		if data.lastScoreInPos == 1 { // data.lastScoreOutOfPos == 1
			// Answer could be or bbxb or bbbx
			savePossibleAnswer(data.previousGuessStr[2], []int{0})
			savePossibleAnswer(data.previousGuessStr[2], []int{1})
			savePossibleAnswer(data.previousGuessStr[2], []int{2, 3})
		}
		if data.lastScoreInPos == 2 { // data.lastScoreOutOfPos == 0
			// Answer could be xbbb or bxbb
			savePossibleAnswer(data.previousGuessStr[2], []int{0, 1})
			savePossibleAnswer(data.previousGuessStr[2], []int{2})
			savePossibleAnswer(data.previousGuessStr[2], []int{3})
		}
	}
	// data.lastCombinedScore == 3 is not possible; data.currScoreInPos would need to be 4
}

func savePossibleAnswer(char byte, possibleIdexes []int) {
	suspectedAnswers = append(suspectedAnswers, SuspectedAnswerDigit{string(char), possibleIdexes})
}
