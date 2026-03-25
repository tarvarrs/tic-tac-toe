package game

import (
	"github.com/google/uuid"
)

const (
	MaximizerFigure = 1
	MinimizerFigure = -1
	EmptyFigure     = 0
	WinScoreMax     = 3
	WinScoreMin     = -3
)

type CurrentGame struct {
	UUID uuid.UUID
	Grid Grid
}

type Grid struct {
	Matrix [3][3]int
}
