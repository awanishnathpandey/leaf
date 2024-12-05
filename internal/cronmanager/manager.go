package cronmanager

import (
	"context"
	"sync"
	"time"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/internal/config"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type CronManager struct {
	CronScheduler   *cron.Cron
	DB              *generated.Queries // Assuming you're using sqlc for database queries
	Jobs            map[string]cron.EntryID
	monitorInterval time.Duration
	JobRegistry     *JobRegistry // Registry for job functions
	mu              sync.Mutex   // Mutex for protecting access to JobRegistry
}

// NewCronManager initializes the CronManager
func NewCronManager(db *generated.Queries) *CronManager {
	return &CronManager{
		CronScheduler: cron.New(),
		DB:            db,
		Jobs:          make(map[string]cron.EntryID),
		JobRegistry:   NewJobRegistry(),
	}
}

// Start will initialize and start the cron jobs
func (cm *CronManager) Start() {
	cm.monitorInterval = config.GetCronJobMonitorInterval()
	log.Info().Msg("All Cron jobs started")
	// Start the cron scheduler
	cm.CronScheduler.Start()

	// Immediately check and add cron jobs before starting the ticker
	cm.checkAndUpdateCronJobs()

	// Start monitoring cron jobs for updates
	go cm.monitorCronJobs()

	// Keep the main function alive
	select {}
}

func (cm *CronManager) Stop() {
	// Stop the cron scheduler from running further jobs
	cm.CronScheduler.Stop()

	// Optionally, you can remove jobs if you want a complete shutdown
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for _, entryID := range cm.Jobs {
		cm.CronScheduler.Remove(entryID)
	}
	log.Info().Msg("All Cron jobs stopped")
}

// MonitorCronJobs checks for updates to active column and schedules in the database
func (cm *CronManager) monitorCronJobs() {
	ticker := time.NewTicker(cm.monitorInterval)
	defer ticker.Stop()

	// Log the state of cm.Jobs at the start of each tick
	// log.Debug().Msgf("Jobs in the map at the start of monitoring: %v", cm.Jobs)

	// Use for range to read from the ticker channel
	for range ticker.C {
		cm.checkAndUpdateCronJobs() // Call the function on each tick
	}
}

func (cm *CronManager) checkAndUpdateCronJobs() {
	// log.Debug().Msgf("Jobs in the map at the start of monitoring: %v", cm.Jobs)
	ctx := context.Background()

	// Fetch active cron jobs from the database
	cronJobs, err := cm.DB.ListActiveCronJobs(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch active cron jobs")
		return
	}

	// log.Debug().Msgf("Fetched cron jobs: %v", cronJobs) // Debugging step

	// Mutex to lock the Jobs map and prevent race conditions
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Iterate over the fetched cron jobs to manage them
	for _, job := range cronJobs {
		// Check if the job is active and the value is valid
		if job.Active.Valid && job.Active.Bool {
			// The job is active, check if it's already in the map
			if entryID, exists := cm.Jobs[job.Slug]; exists {
				// The job exists, check if the schedule has changed
				// log.Debug().Msgf("Job %s exists, checking schedule update", job.Slug)
				if cm.isScheduleUpdated(entryID, job.Schedule) {
					// log.Debug().Msgf("Schedule for job %s has changed, updating job", job.Slug)
					cm.updateJob(entryID, job.Schedule)
				} else {
					// log.Debug().Msgf("Schedule for job %s has not changed, no update required", job.Slug)
				}
			} else {
				// New job, add it to the scheduler
				// log.Debug().Msgf("Adding new job with Slug: %s", job.Slug)
				cm.addCronJob(job)
			}
		} else {
			// The job is no longer active, remove it from the scheduler if it exists
			if entryID, exists := cm.Jobs[job.Slug]; exists {
				// log.Debug().Msgf("Stopping cron job with Slug: %s as it is no longer active", job.Slug)
				cm.CronScheduler.Remove(entryID)
				delete(cm.Jobs, job.Slug)
				log.Debug().Msgf("Stopped cron job with Slug: %s", job.Slug)
			}
		}
	}

	// log.Debug().Msgf("Jobs in the map after processing: %v", cm.Jobs)
}

// Helper function to check if the schedule has changed for an existing job
func (cm *CronManager) isScheduleUpdated(entryID cron.EntryID, newSchedule string) bool {
	// Parse the new schedule string into a cron.Schedule using robfig/cron package
	parsedSchedule, err := cron.ParseStandard(newSchedule)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to parse new schedule: %s", newSchedule)
		return false // If parsing fails, do not update the schedule
	}

	// Retrieve the existing entry from the cron scheduler
	entry := cm.CronScheduler.Entry(entryID)

	// Compare the next scheduled run time of both schedules
	oldNextRun := entry.Schedule.Next(time.Now())
	newNextRun := parsedSchedule.Next(time.Now())

	// If the next run times are different, then the schedules are different
	return oldNextRun != newNextRun
}

// Add a new cron job to the scheduler
func (cm *CronManager) addCronJob(job generated.CronJob) {
	// log.Debug().Msgf("Jobs in the map before update: %v", cm.Jobs)
	entryID, err := cm.CronScheduler.AddFunc(job.Schedule, func() {
		// log.Debug().Msgf("Running cron job with Slug: %s", job.Slug)
		cm.runCronJob(job)
	})
	if err != nil {
		log.Error().Err(err).Msg("Error adding cron job")
		return
	}
	// Store the cron job in the map
	// log.Debug().Msgf("Adding cron job with Slug: %s", job.Slug)
	cm.Jobs[job.Slug] = entryID
	log.Info().Msgf("Started cron job with Slug: %s", job.Slug)
	// log.Debug().Msgf("Jobs in the map after adding: %v", cm.Jobs) // Debugging step
}

// Update an existing cron job
func (cm *CronManager) updateJob(entryID cron.EntryID, schedule string) {
	// log.Debug().Msgf("Jobs in the map before update: %v", cm.Jobs)
	// Remove the old job
	cm.CronScheduler.Remove(entryID)

	// Add the updated job with the new schedule
	_, err := cm.CronScheduler.AddFunc(schedule, func() {
		// Placeholder for real job data; adapt as needed
		cm.runCronJob(generated.CronJob{ID: 1, Schedule: schedule})
	})
	if err != nil {
		log.Error().Err(err).Msg("Error updating cron job")
	}
}

// Run the cron job based on slug
func (cm *CronManager) runCronJob(job generated.CronJob) {
	log.Info().Msgf("Running cron job with Slug: %s", job.Slug)

	// Lock access to JobRegistry during execution to ensure thread-safety
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Look up and run the job function using the JobRegistry
	jobFunction := cm.JobRegistry.GetJobFunction(job.Slug)
	if jobFunction != nil {
		jobFunction() // Execute the job's function
	} else {
		log.Info().Msgf("runcronjob: unknown cron job with Slug: %s", job.Slug)
	}
}