package models

import "github.com/google/uuid"

type Round struct {
	Base
	MatchID uuid.UUID `gorm:"type:uuid;not null" json:"match_id"`
	Number  int       `gorm:"not null" json:"number"`
	Actions []Action  `gorm:"foreignKey:RoundID" json:"actions"`
	
	Match   Match     `gorm:"foreignKey:MatchID" json:"match"`
} 