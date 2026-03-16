package mapper

import (
	"github.com/google/uuid"
	domainModel "tic-tac-toe/internal/domain/models"
	webDTO "tic-tac-toe/internal/web/models"
)

func DomainToResponse(domainModel domainModel.CurrentGame, winner string) webDTO.GameResponse {
	return webDTO.GameResponse{
		UUID:        domainModel.UUID,
		CurrentGame: domainModel.Grid.Matrix,
		Winner:      winner,
	}
}

func WebToDomain(uuid uuid.UUID, webDTO webDTO.GameRequest) domainModel.CurrentGame {
	return domainModel.CurrentGame{
		UUID: uuid,
		Grid: domainModel.Grid{Matrix: webDTO.CurrentGame},
	}
}
