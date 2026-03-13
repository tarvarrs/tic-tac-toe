package repository

import (
	"errors"
	"github.com/google/uuid"
	"tic-tac-toe/internal/datasource/mapper"
	dsModel "tic-tac-toe/internal/datasource/models"
	domainModel "tic-tac-toe/internal/domain/models"
	"tic-tac-toe/internal/domain/service"
)

var _ service.GameRepository = (*GameStorage)(nil)

func (s *GameStorage) Save(cg domainModel.CurrentGame) {
	dsGame := mapper.ToDatasource(cg)
	s.db.Store(dsGame.UUID, dsGame)
}

func (s *GameStorage) Get(uuid uuid.UUID) (domainModel.CurrentGame, error) {
	value, ok := s.db.Load(uuid)
	if !ok {
		return domainModel.CurrentGame{}, errors.New("game not found")
	}
	dsGame := value.(dsModel.CurrentGame)
	return mapper.ToDomain(dsGame), nil
}
