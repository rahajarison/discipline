package handlers

import (
	"net/http"

	"discipline/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ActionHandler struct {
	db *gorm.DB
}

func NewActionHandler(db *gorm.DB) *ActionHandler {
	return &ActionHandler{db: db}
}

// CreateAction creates a new action
func (h *ActionHandler) CreateAction(c echo.Context) error {
	action := new(models.Action)
	if err := c.Bind(action); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := h.db.Create(action).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create action"})
	}

	return c.JSON(http.StatusCreated, action)
}

// GetActions retrieves all actions for a round
func (h *ActionHandler) GetActions(c echo.Context) error {
	roundID, err := uuid.Parse(c.Param("roundId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid round ID"})
	}

	var actions []models.Action
	if err := h.db.Where("round_id = ?", roundID).Order("timestamp").Find(&actions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch actions"})
	}

	return c.JSON(http.StatusOK, actions)
}

// GetAction retrieves an action by ID
func (h *ActionHandler) GetAction(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid action ID"})
	}

	var action models.Action
	if err := h.db.Preload("Round").First(&action, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Action not found"})
	}

	return c.JSON(http.StatusOK, action)
}

// UpdateAction updates an action
func (h *ActionHandler) UpdateAction(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid action ID"})
	}

	var action models.Action
	if err := h.db.First(&action, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Action not found"})
	}

	if err := c.Bind(&action); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := h.db.Save(&action).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update action"})
	}

	return c.JSON(http.StatusOK, action)
}

// DeleteAction deletes an action
func (h *ActionHandler) DeleteAction(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid action ID"})
	}

	if err := h.db.Delete(&models.Action{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete action"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Action deleted successfully"})
}

// GetActionsByPlayer retrieves all actions for a player in a round
func (h *ActionHandler) GetActionsByPlayer(c echo.Context) error {
	roundID, err := uuid.Parse(c.Param("roundId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid round ID"})
	}

	player := c.Param("player")
	if player != "P1" && player != "P2" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Player must be P1 or P2"})
	}

	var actions []models.Action
	if err := h.db.Where("round_id = ? AND player = ?", roundID, player).Order("timestamp").Find(&actions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch actions"})
	}

	return c.JSON(http.StatusOK, actions)
}
