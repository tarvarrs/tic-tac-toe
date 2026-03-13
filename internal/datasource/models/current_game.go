package models

import (
	"github.com/google/uuid"
)

type CurrentGame struct {
	UUID uuid.UUID
	Grid [3][3]int
}
