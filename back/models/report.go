package models

import "github.com/google/uuid"

type Report struct {
	Base
	MatchID   uuid.UUID `gorm:"type:uuid;not null"`
	Type      string    `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	AuthorID  uuid.UUID `gorm:"type:uuid;not null"`
	
	Match     Match     `gorm:"foreignKey:MatchID"`
	Author    User      `gorm:"foreignKey:AuthorID"`
} 