package services

import (
	"discipline/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

// GetUsers retrieves all users with pagination
func (s *UserService) GetUsers(page, limit int) ([]models.User, error) {
	var users []models.User
	
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	
	offset := (page - 1) * limit
	
	err := s.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id uuid.UUID) error {
	return s.db.Delete(&models.User{}, "id = ?", id).Error
}

// GetMatchesByUser retrieves all matches for a user
func (s *UserService) GetMatchesByUser(userID uuid.UUID) ([]models.Match, error) {
	var matches []models.Match
	err := s.db.Preload("Player1").Preload("Player2").
		Where("player1_id = ? OR player2_id = ?", userID, userID).
		Find(&matches).Error
	return matches, err
}
