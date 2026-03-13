package repository

import (
	"sync"
)

type GameStorage struct {
	db sync.Map
}
