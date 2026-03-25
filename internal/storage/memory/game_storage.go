package memory

import (
	"errors"
	"sync"
	"tic-tac-toe/internal/game"

	"github.com/google/uuid"
)

type GameStorage struct {
	db sync.Map
}

func NewGameStorage() *GameStorage {
	return &GameStorage{}
}

var _ game.GameRepository = (*GameStorage)(nil)

func (s *GameStorage) Save(cg game.CurrentGame) {
	var newGame = GameRecord{}
	newGame = newGame.NewGameRecord(cg)
	s.db.Store(newGame.UUID, newGame)
}

func (s *GameStorage) Get(uuid uuid.UUID) (game.CurrentGame, error) {
	value, ok := s.db.Load(uuid)
	if !ok {
		return game.CurrentGame{}, errors.New("game not found")
	}
	newGame := value.(GameRecord)
	return newGame.ToDomain(), nil
}
