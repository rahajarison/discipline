package models

import (
	"time"
	"github.com/google/uuid"
)

type Action struct {
	Base
	Type       string    `gorm:"not null" json:"type"`
	Category   string    `gorm:"not null" json:"category"`
	HitContext string    `gorm:"not null" json:"hit_context"`
	Player     string    `gorm:"not null" json:"player"` // P1 or P2
	Timestamp  time.Time `gorm:"not null" json:"timestamp"`
	RoundID    uuid.UUID `gorm:"type:uuid;not null" json:"round_id"`
	
	Round      Round     `gorm:"foreignKey:RoundID" json:"round"`
} 