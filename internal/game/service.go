package game

import (
	"errors"
	"math"
)

var (
	ErrMoreThanOneTurn      = errors.New("new field contains more than one new turn")
	ErrPreviousTurnsChanged = errors.New("previous turns have been changed")
	ErrNoNewTurns           = errors.New("no new turns")
)

type Manager interface {
	GetNextTurn(Grid) Grid
	ValidateCurrentState(oldGrid Grid, newGrid Grid) error
	CheckForWin(Grid) (int, bool)
}

type GameService struct {
	repo GameRepository
}

func NewGameService() *GameService {
	return &GameService{}
}

var _ Manager = (*GameService)(nil)

func (s *GameService) GetNextTurn(grid Grid) Grid {
	bestResult := math.MaxInt64
	var bestTurn Grid
	for _, state := range s.actions(grid) {
		currentResult := s.minimax(state)
		if currentResult < bestResult {
			bestResult = currentResult
			bestTurn = state
		}
	}
	return bestTurn
}

func (s *GameService) player(grid Grid) int {
	turnsCount := 0
	for i := range 3 {
		for j := range 3 {
			if grid.Matrix[i][j] != EmptyFigure {
				turnsCount++
			}
		}
	}
	if turnsCount%2 == 0 {
		return MaximizerFigure
	}
	return MinimizerFigure
}

func (s *GameService) actions(grid Grid) []Grid {
	allPossibleStates := make([]Grid, 0)
	currentTurnFigure := s.player(grid)
	for i := range 3 {
		for j := range 3 {
			if grid.Matrix[i][j] == EmptyFigure {
				newGrid := grid
				newGrid.Matrix[i][j] = currentTurnFigure
				allPossibleStates = append(allPossibleStates, newGrid)
			}
		}
	}
	return allPossibleStates
}

func (s *GameService) minimax(grid Grid) int {
	winner, isTerminal := s.CheckForWin(grid)
	if isTerminal {
		return winner
	}
	if s.player(grid) == MaximizerFigure {
		value := math.MinInt64
		for _, a := range s.actions(grid) {
			value = max(value, s.minimax(a))
		}
		return value
	}
	if s.player(grid) == MinimizerFigure {
		value := math.MaxInt64
		for _, a := range s.actions(grid) {
			value = min(value, s.minimax(a))
		}
		return value
	}
	return winner
}

func (s *GameService) ValidateCurrentState(oldGrid Grid, newGrid Grid) error {
	newTurns := 0
	for i := range 3 {
		for j := range 3 {
			if oldGrid.Matrix[i][j] != newGrid.Matrix[i][j] {
				if oldGrid.Matrix[i][j] != EmptyFigure {
					return ErrPreviousTurnsChanged
				}
				newTurns++
			}
		}
	}
	if newTurns > 1 {
		return ErrMoreThanOneTurn
	} else if newTurns == 0 {
		return ErrNoNewTurns
	}
	return nil
}

func (s *GameService) CheckForWin(grid Grid) (int, bool) {
	leftDiagonalSum := 0
	rightDiagonalSum := 0
	hasEmptyCells := false
	for i := range 3 {
		horizontalSum := 0
		verticalSum := 0
		for j := range 3 {
			if grid.Matrix[i][j] == EmptyFigure {
				hasEmptyCells = true
			}
			horizontalSum += grid.Matrix[i][j]
			verticalSum += grid.Matrix[j][i]
		}
		if horizontalSum == WinScoreMax || verticalSum == WinScoreMax {
			return MaximizerFigure, true
		} else if horizontalSum == WinScoreMin || verticalSum == WinScoreMin {
			return MinimizerFigure, true
		}
		leftDiagonalSum += grid.Matrix[i][2-i]
		rightDiagonalSum += grid.Matrix[i][i]
	}
	if leftDiagonalSum == WinScoreMax || rightDiagonalSum == WinScoreMax {
		return MaximizerFigure, true
	} else if leftDiagonalSum == WinScoreMin || rightDiagonalSum == WinScoreMin {
		return MinimizerFigure, true
	}
	if hasEmptyCells {
		return EmptyFigure, false
	}
	return EmptyFigure, true
}
