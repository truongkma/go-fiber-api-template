package utils

import (
	"fmt"

	"github.com/truongkma/go-fiber-api-template/pkg/repository"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case repository.AdminRoleName:
		// Admin credentials.
		credentials = []string{}
	case repository.UserRoleName:
		// Simple user credentials.
		credentials = []string{}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
