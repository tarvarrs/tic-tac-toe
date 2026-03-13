package service

import (
	"errors"
	"math"
	"tic-tac-toe/internal/domain/models"
)

var (
	ErrMoreThanOneTurn      = errors.New("new field contains more than one new turn")
	ErrPreviousTurnsChanged = errors.New("previous turns have been changed")
	ErrNoNewTurns           = errors.New("no new turns")
)

var _ Manager = (*GameService)(nil)

type GameService struct {
	r GameRepository
}

func (s *GameService) GetNextTurn(grid models.Grid) models.Grid {
	bestResult := math.MaxInt64
	var bestTurn models.Grid
	for _, state := range s.actions(grid) {
		currentResult := s.minimax(state)
		if currentResult < bestResult {
			bestResult = currentResult
			bestTurn = state
		}
	}
	return bestTurn
}

func (s *GameService) player(grid models.Grid) int {
	turnsCount := 0
	for i := range 3 {
		for j := range 3 {
			if grid.Matrix[i][j] != models.EmptyFig {
				turnsCount++
			}
		}
	}
	if turnsCount%2 == 0 {
		return models.MaximizerFig
	}
	return models.MinimizerFig
}

func (s *GameService) actions(grid models.Grid) []models.Grid {
	allPossibleStates := make([]models.Grid, 0)
	currentTurnFigure := s.player(grid)
	for i := range 3 {
		for j := range 3 {
			if grid.Matrix[i][j] == models.EmptyFig {
				newGrid := grid
				newGrid.Matrix[i][j] = currentTurnFigure
				allPossibleStates = append(allPossibleStates, newGrid)
			}
		}
	}
	return allPossibleStates
}

func (s *GameService) minimax(grid models.Grid) int {
	winner, isTerminal := s.CheckForWin(grid)
	if isTerminal {
		return winner
	}
	if s.player(grid) == models.MaximizerFig {
		value := math.MinInt64
		for _, a := range s.actions(grid) {
			value = max(value, s.minimax(a))
		}
		return value
	}
	if s.player(grid) == models.MinimizerFig {
		value := math.MaxInt64
		for _, a := range s.actions(grid) {
			value = min(value, s.minimax(a))
		}
		return value
	}
	return winner
}

func (s *GameService) ValidateCurrentState(oldGrid models.Grid, newGrid models.Grid) error {
	newTurns := 0
	for i := range 3 {
		for j := range 3 {
			if oldGrid.Matrix[i][j] != newGrid.Matrix[i][j] {
				if oldGrid.Matrix[i][j] != models.EmptyFig {
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

func (s *GameService) CheckForWin(grid models.Grid) (int, bool) {
	leftDiagonalSum := 0
	rightDiagonalSum := 0
	hasEmptyCells := false
	for i := range 3 {
		horizontalSum := 0
		verticalSum := 0
		for j := range 3 {
			if grid.Matrix[i][j] == models.EmptyFig {
				hasEmptyCells = true
			}
			horizontalSum += grid.Matrix[i][j]
			verticalSum += grid.Matrix[j][i]
		}
		if horizontalSum == models.WinScoreMax || verticalSum == models.WinScoreMax {
			return models.MaximizerFig, true
		} else if horizontalSum == models.WinScoreMin || verticalSum == models.WinScoreMin {
			return models.MinimizerFig, true
		}
		leftDiagonalSum += grid.Matrix[i][2-i]
		rightDiagonalSum += grid.Matrix[i][i]
	}
	if leftDiagonalSum == models.WinScoreMax || rightDiagonalSum == models.WinScoreMax {
		return models.MaximizerFig, true
	} else if leftDiagonalSum == models.WinScoreMin || rightDiagonalSum == models.WinScoreMin {
		return models.MinimizerFig, true
	}
	if hasEmptyCells {
		return models.EmptyFig, false
	}
	return models.EmptyFig, true
}
