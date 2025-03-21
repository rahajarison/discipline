package models

type User struct {
	Base
	Username     string    `gorm:"uniqueIndex;not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	Role         string    `gorm:"not null"`
	PasswordHash string    `gorm:"not null"`
	MatchesPlayer1     []Match   `gorm:"foreignKey:Player1ID"` 
	MatchesPlayer2     []Match   `gorm:"foreignKey:Player2ID"` 
	Reports      []Report  `gorm:"foreignKey:AuthorID"`
} 