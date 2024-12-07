package cronmanager

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/internal/cronmanager/jobs"
)

// JobRegistry holds a map of job slugs to functions
type JobRegistry struct {
	Jobs map[string]func()
}

// NewJobRegistry initializes the job registry with predefined jobs
func NewJobRegistry(queries *generated.Queries) *JobRegistry {
	return &JobRegistry{
		Jobs: map[string]func(){
			"sync_users": func() {
				ctx := context.Background()  // Create a background context
				jobs.SyncUsers(ctx, queries) // Call job function
			},
			"clean_audit_logs": func() {
				ctx := context.Background()       // Create a background context
				jobs.CleanAuditLogs(ctx, queries) // Pass context and db queries to CleanAuditLogs
			},
			"push_notifications": func() {
				ctx := context.Background()          // Create a background context
				jobs.PushNotifications(ctx, queries) // Call job function
			},
		},
	}
}

// RegisterJob registers a new job with a given slug and function.
func (jr *JobRegistry) RegisterJob(slug string, jobFunc func()) {
	if _, exists := jr.Jobs[slug]; exists {
		log.Warn().Msgf("Job with slug %s already exists, overwriting.", slug)
	}
	jr.Jobs[slug] = jobFunc
	// log.Info().Msgf("Registered job with Slug: %s", slug)
}

// GetJobFunction retrieves the function associated with a given job slug.
func (jr *JobRegistry) GetJobFunction(slug string) func() {
	jobFunc, exists := jr.Jobs[slug]
	if !exists {
		log.Warn().Msgf("Job with slug %s not found", slug)
		return nil
	}
	return jobFunc
}

// RunJob runs a registered job by its slug.
func (jr *JobRegistry) RunJob(slug string) {
	jobFunc := jr.GetJobFunction(slug)
	if jobFunc != nil {
		// log.Info().Msgf("Running job with Slug: %s", slug)
		jobFunc() // Execute the registered job
	} else {
		log.Warn().Msgf("Job with Slug: %s not registered", slug)
	}
}
