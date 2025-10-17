package models

import (
	"time"
	"github.com/google/uuid"
)

type Action struct {
	Base
	Type       string    `gorm:"not null"`
	Category   string    `gorm:"not null"`
	HitContext string    `gorm:"not null"`
	Player     string    `gorm:"not null"` // P1 or P2
	Timestamp  time.Time `gorm:"not null"`
	RoundID    uuid.UUID `gorm:"type:uuid;not null"`
	
	Round      Round     `gorm:"foreignKey:RoundID"`
} 