package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"fmt"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/graph/model"
	"github.com/awanishnathpandey/leaf/internal/middleware"
	"github.com/awanishnathpandey/leaf/internal/utils"
)

// DeleteAuditLog is the resolver for the deleteAuditLog field.
func (r *mutationResolver) DeleteAuditLog(ctx context.Context, id int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "delete_log"}

	// Check if the user has the required permissions
	if err := middleware.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
	// Check if the audit log exists (optional)
	_, err := r.DB.GetAuditLog(ctx, id)
	if err != nil {
		return false, fmt.Errorf("log not found: %w", err)
	}

	// Attempt to delete the audit log
	err = r.DB.DeleteAuditLog(ctx, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete log: %w", err)
	}
	return true, nil
}

// DeleteAuditLogs is the resolver for the deleteAuditLogs field.
func (r *mutationResolver) DeleteAuditLogs(ctx context.Context, ids []int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "delete_log"}

	// Check if the user has the required permissions
	if err := middleware.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
	// Validate that all IDs exist
	existingFiles, err := r.DB.GetAuditLogsByIDs(ctx, ids)
	if err != nil {
		return false, fmt.Errorf("failed to fetch audit logs: %w", err)
	}
	if len(existingFiles) != len(ids) {
		return false, fmt.Errorf("validation failed: some audit logs do not exist")
	}

	// Proceed to delete the audit logs
	err = r.DB.DeleteAuditLogsByIDs(ctx, ids)
	if err != nil {
		return false, fmt.Errorf("failed to delete audit logs: %w", err)
	}

	// All files successfully deleted
	return true, nil
}

// AuditLogs is the resolver for the auditLogs field.
func (r *queryResolver) AuditLogs(ctx context.Context, first int64, after *int64, filter *model.AuditLogFilter, sort *model.AuditLogSort) (*model.AuditLogConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_log"}

	// Check if the user has the required permissions
	if err := middleware.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}

	// Prepare sorting
	sortField := "TIMESTAMP" // Default sort field
	sortOrder := "DESC"      // Default sort order
	if sort != nil {
		// Prepare sorting using the utility
		sortField, sortOrder = utils.PrepareSorting("TIMESTAMP", "ASC", string(sort.Field), string(sort.Order))
	}

	// Calculate pagination and sorting
	offset, first := utils.PreparePaginationParams(after, first)

	// Prepare filter values
	var tableNameFilter, actorFilter, ipAddressFilter, actionFilter, recordKeyFilter, descriptionFilter *string
	if filter != nil {
		tableNameFilter = filter.TableName
		actorFilter = filter.Actor
		ipAddressFilter = filter.IPAddress
		actionFilter = filter.Action
		recordKeyFilter = filter.RecordKey
		descriptionFilter = filter.Description
	}

	auditLogs, err := r.DB.GetPaginatedAuditLogs(ctx, generated.GetPaginatedAuditLogsParams{
		Limit:             int32(first),  // Limit based on 'first' argument
		Offset:            int32(offset), // Offset based on 'after' cursor
		TableNameFilter:   tableNameFilter,
		ActorFilter:       actorFilter,
		IpAddressFilter:   ipAddressFilter,
		ActionFilter:      actionFilter,      // action filter (optional)
		RecordKeyFilter:   recordKeyFilter,   // Record key filter (optional)
		DescriptionFilter: descriptionFilter, // Description filter (optional)
		SortField:         sortField,         // Sorting field
		SortOrder:         sortOrder,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to fetch audit logs: %v", err)
	}

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.GetPaginatedAuditLogsCount(ctx, generated.GetPaginatedAuditLogsCountParams{
		TableNameFilter:   tableNameFilter,
		ActorFilter:       actorFilter,
		IpAddressFilter:   ipAddressFilter,
		ActionFilter:      actionFilter,      // action filter (optional)
		RecordKeyFilter:   recordKeyFilter,   // Record key filter (optional)
		DescriptionFilter: descriptionFilter, // Description filter (optional)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query audit logs: %v", err)
	}
	// Prepare edges and PageInfo for the connection
	edges := make([]*model.AuditLogEdge, len(auditLogs))
	for i, auditLog := range auditLogs {
		edges[i] = &model.AuditLogEdge{
			Cursor: utils.GenerateCursor(offset, int64(i)), // Create cursor from index
			Node: &model.AuditLog{
				ID:          auditLog.ID,
				TableName:   auditLog.TableName,
				Actor:       auditLog.Actor,
				Action:      auditLog.Action,
				RecordKey:   auditLog.RecordKey,
				IPAddress:   auditLog.IpAddress,
				Description: auditLog.Description,
				Timestamp:   auditLog.Timestamp,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := utils.CalculateHasNextPage(offset, int64(len(auditLogs)), totalCount)

	return &model.AuditLogConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// GetAuditLog is the resolver for the getAuditLog field.
func (r *queryResolver) GetAuditLog(ctx context.Context, id int64) (*model.AuditLog, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_log"}

	// Check if the user has the required permissions
	if err := middleware.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Call the generated GetAuditLog query
	auditLog, err := r.DB.GetAuditLog(ctx, id) // assuming input.ID is of type string
	if err != nil {
		return nil, fmt.Errorf("failed to get audit log: %w", err)
	}

	// Convert the SQL result to GraphQL model
	return &model.AuditLog{
		ID:          auditLog.ID,
		TableName:   auditLog.TableName,
		Actor:       auditLog.Actor,
		Action:      auditLog.Action,
		RecordKey:   auditLog.RecordKey,
		IPAddress:   auditLog.IpAddress,
		Description: auditLog.Description,
		Timestamp:   auditLog.Timestamp,
	}, nil
}
