package validation

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// IsValidUUID checks if a string is a valid UUID
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// IsValidEmail checks if a string is a valid email address
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsValidUsername checks if a username is valid (alphanumeric, 3-20 characters)
func IsValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return usernameRegex.MatchString(username)
}

// IsValidPlayer checks if a player string is valid (P1 or P2)
func IsValidPlayer(player string) bool {
	return player == "P1" || player == "P2"
}

// IsValidActionType checks if an action type is valid
func IsValidActionType(actionType string) bool {
	validTypes := []string{"hit", "block", "dodge", "special", "combo"}
	for _, validType := range validTypes {
		if strings.ToLower(actionType) == validType {
			return true
		}
	}
	return false
}

// IsValidActionCategory checks if an action category is valid
func IsValidActionCategory(category string) bool {
	validCategories := []string{"attack", "defense", "movement", "special"}
	for _, validCategory := range validCategories {
		if strings.ToLower(category) == validCategory {
			return true
		}
	}
	return false
}

// IsValidHitContext checks if a hit context is valid
func IsValidHitContext(context string) bool {
	validContexts := []string{"high", "mid", "low", "overhead", "low_profile"}
	for _, validContext := range validContexts {
		if strings.ToLower(context) == validContext {
			return true
		}
	}
	return false
}

// IsValidRole checks if a user role is valid
func IsValidRole(role string) bool {
	validRoles := []string{"player", "admin", "moderator", "spectator"}
	for _, validRole := range validRoles {
		if strings.ToLower(role) == validRole {
			return true
		}
	}
	return false
}
