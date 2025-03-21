package models

import "github.com/google/uuid"

type Round struct {
	Base
	MatchID uuid.UUID `gorm:"type:uuid;not null"`
	Number  int       `gorm:"not null"`
	Actions []Action  `gorm:"foreignKey:RoundID"`
	
	Match   Match     `gorm:"foreignKey:MatchID"`
} 