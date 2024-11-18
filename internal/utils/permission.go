package utils

import (
	"context"
	"fmt"

	"github.com/awanishnathpandey/leaf/db/generated"
)

// CheckUserPermissions is a utility function that fetches user permissions from the database
// and checks if the user has the required permissions for an action.
func CheckUserPermissions(ctx context.Context, requiredPermissions []string, queries *generated.Queries) error {
	// Retrieve userID from context
	userID, ok := ctx.Value("userID").(int64) // Access the user ID from the context
	if !ok {
		return fmt.Errorf("user ID not found in context")
	}
	fmt.Println("Retrieved userID:", userID)

	// Fetch user permissions from the database
	userPermissions, err := queries.GetUserPermissions(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to fetch user permissions: %v", err)
	}

	// Log the retrieved permissions for debugging
	fmt.Println("User Permissions: ", userPermissions)

	// Check if the user has the required permissions
	if !hasPermissions(userPermissions, requiredPermissions) {
		return fmt.Errorf("insufficient permissions")
	}

	return nil
}

// hasPermissions checks if the user has anyone of the required permissions.
func hasPermissions(userPermissions []string, requiredPermissions []string) bool {
	permissionMap := make(map[string]bool)
	for _, p := range userPermissions {
		permissionMap[p] = true
	}

	// Check if any of the required permissions are present in the user's permissions
	for _, required := range requiredPermissions {
		if permissionMap[required] {
			return true // At least one permission matches
		}
	}

	// If no match was found, return false
	return false
}
