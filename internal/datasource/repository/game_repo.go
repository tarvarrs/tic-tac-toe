package repository

import (
	"errors"
	"github.com/google/uuid"
	"tic-tac-toe/internal/datasource/mapper"
	datasourceModel "tic-tac-toe/internal/datasource/models"
	domainModel "tic-tac-toe/internal/domain/models"
	"tic-tac-toe/internal/domain/service"
)

var _ service.GameRepository = (*GameStorage)(nil)

func (s *GameStorage) Save(cg domainModel.CurrentGame) {
	datasourceGame := mapper.ToDatasource(cg)
	s.db.Store(datasourceGame.UUID, datasourceGame)
}

func (s *GameStorage) Get(uuid uuid.UUID) (domainModel.CurrentGame, error) {
	value, ok := s.db.Load(uuid)
	if !ok {
		return domainModel.CurrentGame{}, errors.New("game not found")
	}
	datasourceGame := value.(datasourceModel.CurrentGame)
	return mapper.ToDomain(datasourceGame), nil
}
