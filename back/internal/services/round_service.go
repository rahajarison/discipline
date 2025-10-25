package services

import (
	"time"

	"discipline/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateRoundWithStartAction creates a round and inserts the initial "Round start" system event action within the provided transaction.
func CreateRoundWithStartAction(tx *gorm.DB, matchID uuid.UUID, roundNumber int) (*models.Round, error) {
	// Create round
	round := &models.Round{
		MatchID: matchID,
		Number:  roundNumber,
	}
	if err := tx.Create(round).Error; err != nil {
		return nil, err
	}

	// Create initial action for the round
	startAction := &models.Action{
		Type:       "Event",
		Category:   "System",
		HitContext: "Round start",
		Player:     "N/A",
		Timestamp:  time.Now().UTC(),
		RoundID:    round.ID,
	}
	if err := tx.Create(startAction).Error; err != nil {
		return nil, err
	}

	return round, nil
}



