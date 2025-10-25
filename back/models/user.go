package models

type User struct {
	Base
	Username     string    `gorm:"uniqueIndex;not null" json:"username"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Role         string    `gorm:"not null" json:"role"`
	PasswordHash string    `gorm:"not null" json:"password_hash"`
	MatchesPlayer1     []Match   `gorm:"foreignKey:Player1ID" json:"matches_player1"`
	MatchesPlayer2     []Match   `gorm:"foreignKey:Player2ID" json:"matches_player2"`
	Reports      []Report  `gorm:"foreignKey:AuthorID" json:"reports"`
} 