package game

import (
	"github.com/google/uuid"
)

type GameRepository interface {
	Save(cg CurrentGame)
	Get(uuid uuid.UUID) (CurrentGame, error)
}
