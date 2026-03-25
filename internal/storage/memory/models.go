package memory

import (
	"github.com/google/uuid"
	"tic-tac-toe/internal/game"
)

type GameRecord struct {
	UUID uuid.UUID
	Grid [3][3]int
}

func (r *GameRecord) ToDomain() game.CurrentGame {
	return game.CurrentGame{
		UUID: r.UUID,
		Grid: game.Grid{Matrix: r.Grid},
	}
}

func (r *GameRecord) NewGameRecord(g game.CurrentGame) GameRecord {
	return GameRecord{
		UUID: g.UUID,
		Grid: g.Grid.Matrix,
	}
}
