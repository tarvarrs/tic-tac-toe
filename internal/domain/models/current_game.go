package models

import (
	"github.com/google/uuid"
)

type CurrentGame struct {
	UUID uuid.UUID
	Grid Grid
}
