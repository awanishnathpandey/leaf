package cronmanager

import (
	"log"

	"github.com/awanishnathpandey/leaf/internal/cronmanager/jobs"
)

// JobRegistry holds a map of job slugs to functions
type JobRegistry struct {
	jobMap map[string]func()
}

// NewJobRegistry initializes the job registry with predefined jobs
func NewJobRegistry() *JobRegistry {
	return &JobRegistry{
		jobMap: map[string]func(){
			"sync_users":         jobs.SyncUsers,
			"clean_audit_logs":   jobs.CleanAuditLogs,
			"push_notifications": jobs.PushNotifications,
		},
	}
}

// AddJob dynamically adds a new job to the registry
func (r *JobRegistry) AddJob(slug string, jobFunc func()) {
	r.jobMap[slug] = jobFunc
}

// RemoveJob removes a job from the registry
func (r *JobRegistry) RemoveJob(slug string) {
	delete(r.jobMap, slug)
}

// GetJobFunction returns the function for a given job slug, or nil if not found
func (r *JobRegistry) GetJobFunction(slug string) func() {
	return r.jobMap[slug]
}

// RunJob will execute the job function for the given slug if it exists
func (r *JobRegistry) RunJob(slug string) {
	jobFunc := r.GetJobFunction(slug)

	if jobFunc == nil {
		log.Printf("Unknown cron job with Slug: %s", slug)
		return
	}

	log.Printf("Running cron job with Slug: %s", slug)
	jobFunc()
}
