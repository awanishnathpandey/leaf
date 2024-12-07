package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/rs/zerolog/log"
)

func PushNotifications(ctx context.Context, queries *generated.Queries) {
	// Add logic
	log.Info().Msgf("Executing cron job push notifications")

	// Step 1: Start logging the cron job
	logID, err := queries.CreateCronJobLog(ctx, "push_notifications")
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert cron job log")
		return
	}
	// Add logic
	log.Info().Msgf("Started cron job: push_notifications with ID: %d", logID)
	// Step 2: Execute the logic for triggering push notifications
	// Simulate  logic (replace with actual logic)
	time.Sleep(2 * time.Second)
	affectedRecords := int64(0)
	errMsg := ""
	defer func() {
		// Step 3: Handle success or failure and update the job log
		if err != nil {
			// If there was an error, update the status to "failed"
			errMsg = fmt.Sprintf("Error occurred during triggering push notifications: %v", err)
			if err := queries.UpdateCronJobLogFailed(ctx, generated.UpdateCronJobLogFailedParams{
				ID:      logID,
				Message: errMsg,
			}); err != nil {
				log.Error().Err(err).Msg("Failed to update cron job log to failed")
			} else {
				log.Error().Msg(errMsg)
			}
		} else {
			// If successful, update the status to "success"
			affectedRecords = 100 // Example: Assume 100 records were cleaned (replace with actual count)
			if err := queries.UpdateCronJobLogSuccess(ctx, generated.UpdateCronJobLogSuccessParams{
				ID:              logID,
				AffectedRecords: affectedRecords,
			}); err != nil {
				log.Error().Err(err).Msg("Failed to update cron job log to success")
			} else {
				log.Info().Msgf("Cron job: push_notifications completed successfully. Affected records: %d", affectedRecords)
			}
		}
	}()

	// Step 4: Trigger push notifications (your actual push notifications logic here)
	// Example: triggerPushNotificationsLogic()
	// affectedRecords = triggerPushNotificationsLogic() // Replace this line with actual cleanup code

	// Simulate success or failure for this example
	// err = nil // Uncomment this line if you want to simulate success
	// err = fmt.Errorf("simulated error") // Uncomment this line if you want to simulate failure
}
