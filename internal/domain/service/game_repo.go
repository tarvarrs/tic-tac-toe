package service

import (
	"github.com/google/uuid"
	"tic-tac-toe/internal/domain/models"
)

type GameRepository interface {
	Save(cg models.CurrentGame)
	Get(uuid uuid.UUID) (models.CurrentGame, error)
}
