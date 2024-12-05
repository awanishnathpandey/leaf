package utils

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/internal/config"
	"github.com/rs/zerolog/log"
)

var (
	// Cache to store user permissions with a read-write mutex for thread safety
	userPermissionsCache  = make(map[int64][]string)
	permissionsTimestamps = make(map[int64]time.Time) // Map to store the timestamp of when permissions were last fetched
	cacheMutex            = &sync.RWMutex{}
	cacheExpiry           time.Duration // Cache expiration time, to be set in init
	cacheMaxSize          int           // Max size of the cache, to be set in init
	cacheCleanupInterval  time.Duration // Cache cleanup interval, to be set in init
	cleanupCancelChan     = make(chan struct{})
)

func InitializePermissionCache() {
	// Initialize cache settings at runtime
	cacheExpiry = config.GetCacheExpiry()   // Cache expiration time
	cacheMaxSize = config.GetCacheMaxSize() // Max size of the cache
	cacheCleanupInterval = config.GetCacheCleanupInterval()
	// Start the automatic cleanup scheduler
	startCacheCleanupScheduler()
	log.Info().Msg("Started permissions cache with cleanup scheduler")
}

// CheckUserPermissions verifies if the user has the required permissions for an action,
// with added caching to improve performance by reducing database queries.
func CheckUserPermissions(ctx context.Context, requiredPermissions []string, queries *generated.Queries) error {
	// Retrieve userID from context
	userID, ok := ctx.Value("userID").(int64)
	if !ok {
		return errors.New("user ID not found in context")
	}

	// Check cache first
	permissions, found := getCachedPermissions(userID)
	if found {
		// If found in cache, use the cached permissions
		if hasRequiredPermissions(permissions, requiredPermissions) {
			return nil
		}
	}

	// If not found in cache, fetch from database and update cache
	userPermissions, err := queries.GetUserPermissions(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to fetch user permissions: %w", err)
	}

	// Cache the fetched permissions with a timestamp
	cachePermissions(userID, userPermissions)

	// Check permissions
	if hasRequiredPermissions(userPermissions, requiredPermissions) {
		return nil
	}

	// Return error if permissions are insufficient
	return errors.New("insufficient permissions")
}

// getCachedPermissions retrieves cached permissions if they are still valid.
func getCachedPermissions(userID int64) ([]string, bool) {
	cacheMutex.RLock() // Read lock for cache
	defer cacheMutex.RUnlock()

	permissions, found := userPermissionsCache[userID]
	if !found {
		return nil, false
	}

	// If cache is stale (expired), return false
	if time.Since(permissionsTimestamp(userID)) > cacheExpiry {
		return nil, false
	}

	return permissions, true
}

// cachePermissions stores the permissions in cache with a timestamp.
func cachePermissions(userID int64, permissions []string) {
	cacheMutex.Lock() // Write lock for cache
	defer cacheMutex.Unlock()

	// Check if the cache size exceeds the limit
	if len(userPermissionsCache) >= cacheMaxSize {
		evictExpiredCacheEntries() // Evict expired entries if the cache is full
	}

	// Store the permissions and update the timestamp
	userPermissionsCache[userID] = permissions
	permissionsTimestamps[userID] = time.Now() // Update the timestamp for this user
}

// evictExpiredCacheEntries removes expired entries from the cache.
func evictExpiredCacheEntries() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	for userID, timestamp := range permissionsTimestamps {
		if time.Since(timestamp) > cacheExpiry {
			// Remove the expired entry from both cache maps
			delete(userPermissionsCache, userID)
			delete(permissionsTimestamps, userID)
		}
	}
	log.Info().Msg("Permission Cache cleanup complete")

	// If the cache still exceeds the max size, we can perform additional logic here.
	// This could be either random eviction or eviction based on some custom policy (e.g., LRU).
}

// permissionsTimestamp retrieves the timestamp of when the permissions were last updated.
func permissionsTimestamp(userID int64) time.Time {
	cacheMutex.RLock() // Read lock for cache
	defer cacheMutex.RUnlock()

	// Retrieve the timestamp from the map (if not found, return a zero value)
	return permissionsTimestamps[userID]
}

// hasRequiredPermissions verifies if any required permission exists in the user's permissions.
func hasRequiredPermissions(userPermissions, requiredPermissions []string) bool {
	// Convert required permissions to a map for efficient lookup
	requiredSet := make(map[string]struct{}, len(requiredPermissions))
	for _, perm := range requiredPermissions {
		requiredSet[perm] = struct{}{}
	}

	// Check for an intersection between userPermissions and requiredPermissions
	for _, perm := range userPermissions {
		if _, exists := requiredSet[perm]; exists {
			return true
		}
	}

	return false
}

// startCacheCleanupScheduler starts a background goroutine to enforce periodic cache cleanup.
func startCacheCleanupScheduler() {
	go func() {
		ticker := time.NewTicker(cacheCleanupInterval)
		defer ticker.Stop()

		for {
			select {
			case <-cleanupCancelChan: // Listen for stop signal
				log.Info().Msg("Permission cache cleanup scheduler stopped")
				return
			case <-ticker.C:
				evictExpiredCacheEntries() // Perform cleanup
			}
		}
	}()
}

// StopPermissionCacheCleanup sends a signal to stop the cleanup process.
func StopPermissionCacheCleanup() {
	if cleanupCancelChan != nil {
		close(cleanupCancelChan) // This will stop the cleanup goroutine
	}
}
