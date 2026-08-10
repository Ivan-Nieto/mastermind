package shared

type Score struct {
	CorrectAndOutOfPosition int
	CorrectAndInPossition   int
	CorrectGuess            bool
}

var EmptyScore = Score{0, 0, false}
