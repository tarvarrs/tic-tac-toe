package http

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"

	"tic-tac-toe/internal/game"
)

type GameRequest struct {
	CurrentGame [3][3]int `json:"grid"`
}

func (req *GameRequest) ToDomain(uuid uuid.UUID) game.CurrentGame {
	return game.CurrentGame{
		UUID: uuid,
		Grid: game.Grid{Matrix: req.CurrentGame},
	}
}

type GameResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	CurrentGame [3][3]int `json:"grid"`
	Winner      string    `json:"winner,omitempty"`
}

func NewGameResponse(game game.CurrentGame, winner string) GameResponse {
	return GameResponse{
		UUID:        game.UUID,
		CurrentGame: game.Grid.Matrix,
		Winner:      winner,
	}
}

type GameHandler struct {
	Manager game.Manager
	Repo    game.GameRepository
}

func NewGameHandler(m game.Manager, r game.GameRepository) *GameHandler {
	return &GameHandler{
		Manager: m,
		Repo:    r,
	}
}

func (h *GameHandler) MakeComputerMove(c *echo.Context) error {
	gameId, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "something went wrong")
	}
	var reqCurrentState GameRequest
	if err := c.Bind(&reqCurrentState); err != nil {
		return c.String(http.StatusBadRequest, "bad request body")
	}
	domainGame := reqCurrentState.ToDomain(gameId)

	prevState, err := h.Repo.Get(gameId)
	prevGrid := prevState.Grid
	if err != nil {
		prevGrid = game.Grid{}
	}
	err = h.Manager.ValidateCurrentState(prevGrid, domainGame.Grid)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	winner, isWin := h.Manager.CheckForWin(domainGame.Grid)
	if isWin {
		switch winner {
		case game.MaximizerFigure:
			return c.JSON(http.StatusOK, NewGameResponse(domainGame, "X"))
		case game.MinimizerFigure:
			return c.JSON(http.StatusOK, NewGameResponse(domainGame, "0"))
		case game.EmptyFigure:
			return c.JSON(http.StatusOK, NewGameResponse(domainGame, "draw"))
		}
	}
	newTurn := h.Manager.GetNextTurn(domainGame.Grid)
	updatedGame := game.CurrentGame{UUID: gameId, Grid: newTurn}
	winner, isWin = h.Manager.CheckForWin(newTurn)
	if isWin {
		if winner == game.MaximizerFigure {
			return c.JSON(http.StatusOK, NewGameResponse(updatedGame, "X"))
		}
		return c.JSON(http.StatusOK, NewGameResponse(updatedGame, "0"))
	}
	h.Repo.Save(updatedGame)
	return c.JSON(http.StatusOK, NewGameResponse(updatedGame, ""))
}
