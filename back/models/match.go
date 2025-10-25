package models

import (
	"time"
	"github.com/google/uuid"
)

type Match struct {
	Base
	Player1ID  uuid.UUID `gorm:"type:uuid;not null" json:"player1_id"`
	Player2ID  uuid.UUID `gorm:"type:uuid;not null" json:"player2_id"`
	Date       time.Time `json:"date"`
	ReplayURL  string    `json:"replay_url"`
	Notes      string    `gorm:"type:text" json:"notes"`
	Rounds     []Round   `gorm:"foreignKey:MatchID" json:"rounds"`
	Reports    []Report  `gorm:"foreignKey:MatchID" json:"reports"`
	// Objectives []Objective `gorm:"foreignKey:MatchID"`
	
	Player1    User      `gorm:"foreignKey:Player1ID" json:"player1"`
	Player2    User      `gorm:"foreignKey:Player2ID" json:"player2"`
} 