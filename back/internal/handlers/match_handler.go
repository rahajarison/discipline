package handlers

import (
	"net/http"
	"strconv"

	"discipline/internal/services"
	"discipline/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MatchHandler struct {
	db *gorm.DB
}

func NewMatchHandler(db *gorm.DB) *MatchHandler {
	return &MatchHandler{db: db}
}

// CreateMatch creates a new match and its first round
func (h *MatchHandler) CreateMatch(c echo.Context) error {
	match := new(models.Match)
	if err := c.Bind(match); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

    // Use a transaction to create the match and the first round together
    tx := h.db.Begin()
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to start transaction"})
	}

    if err := tx.Create(match).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create match"})
	}

    if _, err := services.CreateRoundWithStartAction(tx, match.ID, 1); err != nil {
        tx.Rollback()
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create first round and initial action"})
    }

	if err := tx.Commit().Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
	}

	// Load the match with its relations for the response
	if err := h.db.Preload("Player1").Preload("Player2").Preload("Rounds").First(match, "id = ?", match.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to load match with relations"})
	}

	return c.JSON(http.StatusCreated, match)
}

// GetMatches retrieves all matches with relations
func (h *MatchHandler) GetMatches(c echo.Context) error {
	var matches []models.Match
	
	// Optional pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	
	offset := (page - 1) * limit
	
	if err := h.db.Preload("Player1").Preload("Player2").Offset(offset).Limit(limit).Find(&matches).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch matches"})
	}

	return c.JSON(http.StatusOK, matches)
}

// GetMatch retrieves a match by ID with relations
func (h *MatchHandler) GetMatch(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid match ID"})
	}

	var match models.Match
	if err := h.db.Preload("Player1").Preload("Player2").First(&match, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Match not found"})
	}

	return c.JSON(http.StatusOK, match)
}

// UpdateMatch updates a match
func (h *MatchHandler) UpdateMatch(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid match ID"})
	}

	var match models.Match
	if err := h.db.First(&match, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Match not found"})
	}

	if err := c.Bind(&match); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := h.db.Save(&match).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update match"})
	}

	return c.JSON(http.StatusOK, match)
}

// DeleteMatch deletes a match
func (h *MatchHandler) DeleteMatch(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid match ID"})
	}

	if err := h.db.Delete(&models.Match{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete match"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Match deleted successfully"})
}
