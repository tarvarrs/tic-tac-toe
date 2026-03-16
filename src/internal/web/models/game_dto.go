package models

import (
	"github.com/google/uuid"
)

type GameRequest struct {
	CurrentGame [3][3]int `json:"grid"`
}

type GameResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	CurrentGame [3][3]int `json:"grid"`
	Winner      string    `json:"winner,omitempty"`
}
