package main

import (
	"fmt"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/http"
	"tic-tac-toe/internal/storage/memory"

	"context"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				memory.NewGameStorage,
				fx.As(new(game.GameRepository)),
			),
			fx.Annotate(
				game.NewGameService,
				fx.As(new(game.Manager)),
			),
			http.NewGameHandler,
			echo.New,
		),
		fx.Invoke(registerRoutes, startServer),
	).Run()
}

func registerRoutes(e *echo.Echo, h *http.GameHandler) {
	e.Use(middleware.CORS("*"))
	e.File("/openapi.yaml", "openapi.yaml")
	e.GET("/swagger", http.ProvideSwagger)

	e.POST("/game/:uuid", h.MakeComputerMove)
}

func startServer(lc fx.Lifecycle, e *echo.Echo) {
	ctx, cancel := context.WithCancel(context.Background())

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			sc := echo.StartConfig{
				Address: ":8080",
			}
			go func() {
				if err := sc.Start(ctx, e); err != nil {
					e.Logger.Error(fmt.Sprintf("%v", err))
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})
}
