package models

import "github.com/google/uuid"

type Report struct {
	Base
	MatchID   uuid.UUID `gorm:"type:uuid;not null" json:"match_id"`
	Type      string    `gorm:"not null" json:"type"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	AuthorID  uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`
	
	Match     Match     `gorm:"foreignKey:MatchID" json:"match"`
	Author    User      `gorm:"foreignKey:AuthorID" json:"author"`
} 