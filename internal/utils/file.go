package utils

import "time"

// Helper function to determine if a file is new (updated in the last 7 days)
func IsFileNew(updatedAt int64, durationInDays int) bool {
	// Convert the updated_at to a time.Time
	fileUpdatedAt := time.Unix(updatedAt, 0)
	// Get the threshold (current time - 7 days)
	threshold := time.Now().AddDate(0, 0, -durationInDays)
	// Return true if the file was updated in the last 7 days
	return fileUpdatedAt.After(threshold)
}
