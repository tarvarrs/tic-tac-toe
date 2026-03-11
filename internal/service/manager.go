package service

import (
	"github.com/google/uuid"
	"tic-tac-toe/internal/models"
)

type Manager interface {
	GetNextTurn(models.Grid) models.Grid                            // Метод получения следующего хода текущей игры алгоритмом «Минимакс»
	ValidateCurrentState(uuid uuid.UUID, newGrid models.Grid) error // Метод валидации игрового поля текущей игры (проверь, что не изменены предыдущие ходы)
	CheckForWin(models.Grid) (int, bool)                            // Метод проверки окончания игры
}
