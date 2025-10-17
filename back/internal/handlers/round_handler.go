package handlers

import (
	"net/http"

	"discipline/internal/services"
	"discipline/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RoundHandler struct {
	db *gorm.DB
}

func NewRoundHandler(db *gorm.DB) *RoundHandler {
	return &RoundHandler{db: db}
}

// CreateRound creates a new round and its initial action
func (h *RoundHandler) CreateRound(c echo.Context) error {
    // matchId from route, round number from body
    matchID, err := uuid.Parse(c.Param("matchId"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid match ID"})
    }

    payload := new(struct{ Number int `json:"number"` })
    if err := c.Bind(payload); err != nil || payload.Number == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }

    tx := h.db.Begin()
    if tx.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to start transaction"})
    }

    round, svcErr := services.CreateRoundWithStartAction(tx, matchID, payload.Number)
    if svcErr != nil {
        tx.Rollback()
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create round and initial action"})
    }

    if err := tx.Commit().Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
    }

    return c.JSON(http.StatusCreated, round)
}

// GetRounds retrieves all rounds for a match
func (h *RoundHandler) GetRounds(c echo.Context) error {
	matchID, err := uuid.Parse(c.Param("matchId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid match ID"})
	}

	var rounds []models.Round
	if err := h.db.Where("match_id = ?", matchID).Order("number").Find(&rounds).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch rounds"})
	}

	return c.JSON(http.StatusOK, rounds)
}

// GetRound retrieves a round by ID
func (h *RoundHandler) GetRound(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid round ID"})
	}

	var round models.Round
	if err := h.db.Preload("Match").First(&round, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Round not found"})
	}

	return c.JSON(http.StatusOK, round)
}

// UpdateRound updates a round
func (h *RoundHandler) UpdateRound(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid round ID"})
	}

	var round models.Round
	if err := h.db.First(&round, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Round not found"})
	}

	if err := c.Bind(&round); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := h.db.Save(&round).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update round"})
	}

	return c.JSON(http.StatusOK, round)
}

// DeleteRound deletes a round
func (h *RoundHandler) DeleteRound(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid round ID"})
	}

	if err := h.db.Delete(&models.Round{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete round"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Round deleted successfully"})
}
