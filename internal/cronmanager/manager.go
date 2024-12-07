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
		JobRegistry:   NewJobRegistry(db),
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
	cronJobs, err := cm.DB.ListCronJobs(ctx)
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
		// log.Debug().Msgf("Processing job - Slug: %s, Active.Valid: %v, Active.Bool: %v", job.Slug, job.Active.Valid, job.Active.Bool)

		if job.Active.Valid && job.Active.Bool {
			// log.Debug().Msgf("Active job detected: %s", job.Slug)

			if entryID, exists := cm.Jobs[job.Slug]; exists {
				if cm.isScheduleUpdated(entryID, job.Schedule) {
					log.Debug().Msgf("Updating job: %s", job.Slug)
					cm.updateJob(entryID, job)
				} else {
					// log.Debug().Msgf("No update needed for job: %s", job.Slug)
				}
			} else {
				// log.Info().Msgf("Adding new job: %s", job.Slug)
				cm.addCronJob(job)
			}
		} else {
			// log.Debug().Msgf("Inactive or invalid job detected: Slug=%s", job.Slug)

			if job.Slug == "" {
				log.Error().Msg("Encountered job with empty Slug, skipping removal")
				continue
			}

			entryID, exists := cm.Jobs[job.Slug]
			if exists {
				// log.Debug().Msgf("Removing job: %s", job.Slug)
				cm.CronScheduler.Remove(entryID)
				delete(cm.Jobs, job.Slug)
				log.Info().Msgf("Removed job: %s", job.Slug)
			} else {
				// log.Warn().Msgf("Job not found in active jobs map: %s", job.Slug)
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
	if entry.Schedule == nil {
		log.Error().Msgf("Invalid entry schedule for ID: %v", entryID)
		return true // Force update if the entry is invalid
	}

	fixedTime := time.Now()
	oldNextRun := entry.Schedule.Next(fixedTime)
	newNextRun := parsedSchedule.Next(fixedTime)
	// Define a time tolerance (e.g., 1 second)
	tolerance := time.Second

	// Calculate the absolute difference between the next run times
	timeDiff := oldNextRun.Sub(newNextRun)
	if timeDiff < 0 {
		timeDiff = -timeDiff
	}

	// log.Debug().Msgf("Old Next Run: %v, New Next Run: %v, Time Difference: %v", oldNextRun, newNextRun, timeDiff)

	// If the difference exceeds the tolerance, the schedules are considered different
	return timeDiff > tolerance
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
func (cm *CronManager) updateJob(entryID cron.EntryID, job generated.CronJob) {
	// log.Debug().Msgf("Jobs in the map before update: %v", cm.Jobs)
	// Remove the old job
	cm.CronScheduler.Remove(entryID)

	// Add the updated job with the new schedule
	entryID, err := cm.CronScheduler.AddFunc(job.Schedule, func() {
		cm.runCronJob(job)
	})
	if err != nil {
		log.Error().Err(err).Msg("Error updating cron job")
		return
	}

	// Update the Jobs map with the new entry ID
	cm.Jobs[job.Slug] = entryID
	log.Info().Msgf("Updated cron job with Slug: %s", job.Slug)
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
