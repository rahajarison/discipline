package api

import (
	"discipline/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo, userHandler *handlers.UserHandler, matchHandler *handlers.MatchHandler, roundHandler *handlers.RoundHandler, actionHandler *handlers.ActionHandler) {
	// Global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// API v1 group
	v1 := e.Group("/api/v1")

	// Health check
	v1.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Users routes
	users := v1.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.GET("", userHandler.GetUsers)
	users.GET("/:id", userHandler.GetUser)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)
	users.GET("/:userId/matches", userHandler.GetMatchesByUser)

	// Matches routes
	matches := v1.Group("/matches")
	matches.POST("", matchHandler.CreateMatch)
	matches.GET("", matchHandler.GetMatches)
	matches.GET("/:id", matchHandler.GetMatch)
	matches.PUT("/:id", matchHandler.UpdateMatch)
	matches.DELETE("/:id", matchHandler.DeleteMatch)

	// Rounds routes (nested under matches)
	rounds := matches.Group("/:matchId/rounds")
	rounds.POST("", roundHandler.CreateRound)
	rounds.GET("", roundHandler.GetRounds)
	rounds.GET("/:id", roundHandler.GetRound)
	rounds.PUT("/:id", roundHandler.UpdateRound)
	rounds.DELETE("/:id", roundHandler.DeleteRound)

	// Actions routes (nested under rounds)
	actions := rounds.Group("/:roundId/actions")
	actions.POST("", actionHandler.CreateAction)
	actions.GET("", actionHandler.GetActions)
	actions.GET("/:id", actionHandler.GetAction)
	actions.PUT("/:id", actionHandler.UpdateAction)
	actions.DELETE("/:id", actionHandler.DeleteAction)
	actions.GET("/player/:player", actionHandler.GetActionsByPlayer)

	// Alternative routes for direct access (if needed)
	// Direct rounds access
	directRounds := v1.Group("/rounds")
	directRounds.GET("/:id", roundHandler.GetRound)
	directRounds.PUT("/:id", roundHandler.UpdateRound)
	directRounds.DELETE("/:id", roundHandler.DeleteRound)

	// Direct actions access
	directActions := v1.Group("/actions")
	directActions.GET("/:id", actionHandler.GetAction)
	directActions.PUT("/:id", actionHandler.UpdateAction)
	directActions.DELETE("/:id", actionHandler.DeleteAction)
}
