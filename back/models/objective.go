package models

import "github.com/google/uuid"

type Objective struct {
	Base
	MatchID     uuid.UUID `gorm:"type:uuid;not null"`
	Description string    `gorm:"type:text;not null"`
	Status      string    `gorm:"not null"`
	
	Match       Match     `gorm:"foreignKey:MatchID"`
} 