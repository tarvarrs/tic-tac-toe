package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"

	"tic-tac-toe/internal/domain/models"
	"tic-tac-toe/internal/domain/service"
	"tic-tac-toe/internal/web/mapper"
	dto "tic-tac-toe/internal/web/models"
)

type GameHandler struct {
	Manager service.Manager
	Repo    service.GameRepository
}

func NewGameHandler(m service.Manager, r service.GameRepository) *GameHandler {
	return &GameHandler{
		Manager: m,
		Repo:    r,
	}
}

func (h *GameHandler) MakeComputerMove(c *echo.Context) error {
	id := c.Param("uuid")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "no game with such id")
	}
	var reqCurrentState dto.GameRequest
	err = c.Bind(&reqCurrentState)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	currentState := mapper.WebToDomain(uuid, reqCurrentState)
	prevState, err := h.Repo.Get(uuid)
	if err != nil {
		var emptyGrid models.Grid
		err = h.Manager.ValidateCurrentState(emptyGrid, currentState.Grid)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
		}
	}
	err = h.Manager.ValidateCurrentState(prevState.Grid, currentState.Grid)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	winner, isWin := h.Manager.CheckForWin(currentState.Grid)
	if isWin {
		if winner == models.MaximizerFigure {
			return c.JSON(http.StatusOK, mapper.DomainToResponse(currentState, "X"))
		} else if winner == models.MinimizerFigure {
			return c.JSON(http.StatusOK, mapper.DomainToResponse(currentState, "0"))
		}
		return c.JSON(http.StatusOK, mapper.DomainToResponse(currentState, "draw"))
	}
	newTurn := h.Manager.GetNextTurn(currentState.Grid)
	newGame := models.CurrentGame{UUID: uuid, Grid: newTurn}
	winner, isWin = h.Manager.CheckForWin(newTurn)
	if isWin {
		if winner == models.MaximizerFigure {
			return c.JSON(http.StatusOK, mapper.DomainToResponse(newGame, "X"))
		}
		return c.JSON(http.StatusOK, mapper.DomainToResponse(newGame, "0"))
	}
	h.Repo.Save(newGame)
	return c.JSON(http.StatusOK, mapper.DomainToResponse(newGame, ""))
}
