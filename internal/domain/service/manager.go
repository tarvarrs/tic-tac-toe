package service

import (
	"tic-tac-toe/internal/domain/models"
)

type Manager interface {
	GetNextTurn(models.Grid) models.Grid
	ValidateCurrentState(oldGrid models.Grid, newGrid models.Grid) error
	CheckForWin(models.Grid) (int, bool)
}
