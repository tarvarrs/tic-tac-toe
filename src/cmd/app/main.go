package main

import (
	"fmt"
	"tic-tac-toe/internal/datasource/repository"
	"tic-tac-toe/internal/domain/service"
	"tic-tac-toe/internal/web/handler"

	"context"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				repository.NewGameStorage,
				fx.As(new(service.GameRepository)),
			),
			fx.Annotate(
				service.NewGameService,
				fx.As(new(service.Manager)),
			),
			handler.NewGameHandler,
			echo.New,
		),
		fx.Invoke(registerRoutes, startServer),
	).Run()
}

func registerRoutes(e *echo.Echo, h *handler.GameHandler) {
	e.POST("/game/:uuid", h.MakeComputerMove)
	e.Use(middleware.CORS("*"))
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
