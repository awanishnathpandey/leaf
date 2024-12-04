package middleware

import (
	"context"
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/internal/config"
)

// Global variables for worker management and audit log queue
var (
	// queries       *generated.Queries    // Holds the database queries object
	auditLogQueue       chan AuditLogEntry    // Buffered channel for audit log entries
	stopAuditLogWorkers = make(chan struct{}) // Channel to signal workers to stop
	alwg                sync.WaitGroup        // WaitGroup for worker synchronization
	numAuditWorkers     int                   // Number of worker goroutines for logging
)

// AuditLogEntry represents an audit log entry
type AuditLogEntry struct {
	Ctx         context.Context
	TableName   string
	Action      string
	RecordKey   string
	Description string
}

// // Initialize the queries object once
// func InitializeQueries(q *generated.Queries) {
// 	queries = q
// }

// StartWorkerPool initializes the worker pool for handling audit log insertion
func StartAuditWorkerPool() {
	numAuditWorkers = config.GetAuditWorkerPoolSize() // Number of workers
	auditLogQueue = make(chan AuditLogEntry, config.GetAuditLogQueueSize())

	// Start worker goroutines
	for i := 0; i < numAuditWorkers; i++ {
		alwg.Add(1)
		go auditWorker()
	}
	log.Info().Msgf("%d worker(s) started for processing audit logs", numAuditWorkers)
}

// StopWorkerPool gracefully stops all workers
func StopAuditWorkerPool() {
	close(stopAuditLogWorkers)
	alwg.Wait()
	log.Info().Msg("All audit log workers stopped")
}

// Worker goroutine processes audit log entries in the queue
func auditWorker() {
	defer alwg.Done()
	for {
		select {
		case logEntry := <-auditLogQueue:
			err := InsertAuditLog(logEntry.Ctx, logEntry)
			if err != nil {
				log.Error().Err(err).Msg("Failed to insert audit log")
			}
		case <-stopAuditLogWorkers:
			return
		}
	}
}

// Insert the audit log asynchronously
func InsertAuditLog(ctx context.Context, logEntry AuditLogEntry) error {

	// Insert the audit log into the database using generated queries
	err := queries.CreateAuditLog(ctx, generated.CreateAuditLogParams{
		TableName:   logEntry.TableName,
		Actor:       ctx.Value("userEmail").(string),
		Action:      logEntry.Action,
		RecordKey:   logEntry.RecordKey,
		Description: logEntry.Description,
		IpAddress:   ctx.Value("userIpAddress").(string),
	})
	if err != nil {
		return fmt.Errorf("failed to insert audit log for user %s: %w", ctx.Value("userEmail").(string), err)
	}

	return nil
}
