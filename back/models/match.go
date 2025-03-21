package models

import (
	"time"
	"github.com/google/uuid"
)

type Match struct {
	Base
	Player1ID  uuid.UUID `gorm:"type:uuid;not null"`
	Player2ID  uuid.UUID `gorm:"type:uuid;not null"`
	Date       time.Time
	ReplayURL  string
	Notes      string    `gorm:"type:text"`
	Rounds     []Round   `gorm:"foreignKey:MatchID"`
	Reports    []Report  `gorm:"foreignKey:MatchID"`
	// Objectives []Objective `gorm:"foreignKey:MatchID"`
	
	Player1    User      `gorm:"foreignKey:Player1ID"`
	Player2    User      `gorm:"foreignKey:Player2ID"`
} 