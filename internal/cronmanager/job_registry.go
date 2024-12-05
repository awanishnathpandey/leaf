package cronmanager

import (
	"github.com/rs/zerolog/log"

	"github.com/awanishnathpandey/leaf/internal/cronmanager/jobs"
)

// JobRegistry holds a map of job slugs to functions
type JobRegistry struct {
	Jobs map[string]func()
}

// NewJobRegistry initializes the job registry with predefined jobs
func NewJobRegistry() *JobRegistry {
	return &JobRegistry{
		Jobs: map[string]func(){
			"sync_users":         jobs.SyncUsers,
			"clean_audit_logs":   jobs.CleanAuditLogs,
			"push_notifications": jobs.PushNotifications,
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
