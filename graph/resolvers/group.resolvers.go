package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"fmt"
	"strconv"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/graph"
	"github.com/awanishnathpandey/leaf/graph/model"
	"github.com/awanishnathpandey/leaf/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

// Users is the resolver for the users field.
func (r *groupResolver) Users(ctx context.Context, obj *model.Group, first int64, after *int64, filter *model.UserFilter, sort *model.UserSort) (*model.UserConnection, error) {
	// Decode the cursor (if provided)
	var offset int64
	if after != nil { // Check if `after` is provided (non-nil)
		offset = *after
	}

	// Prepare sorting
	sortField := "NAME" // Default sort field
	if sort != nil {
		sortField = string(sort.Field)
	}

	sortOrder := "ASC" // Default sort order
	if sort != nil {
		sortOrder = string(sort.Order)
	}

	// Prepare filter values
	var nameFilter, emailFilter *string
	if filter != nil {
		nameFilter = filter.Name
		emailFilter = filter.Email
	}

	// Fetch users using the SQL query method for group ID
	users, err := r.DB.GetPaginatedUsersByGroupID(ctx, generated.GetPaginatedUsersByGroupIDParams{
		GroupID:     pgtype.Int8{Int64: obj.ID, Valid: true}, // Group ID from the Group object
		Limit:       int32(first),                            // Limit based on 'first' argument
		Offset:      int32(offset),                           // Offset based on 'after' cursor
		NameFilter:  nameFilter,                              // Name filter (optional)
		EmailFilter: emailFilter,                             // Email filter (optional)
		SortField:   sortField,                               // Sorting field
		SortOrder:   sortOrder,                               // Sorting order
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users for group %d: %v", obj.ID, err)
	}

	// Prepare edges and PageInfo for the connection
	edges := make([]*model.UserEdge, len(users))
	for i, user := range users {
		edges[i] = &model.UserEdge{
			Cursor: strconv.FormatInt(offset+int64(i)+1, 10), // Create cursor from index
			Node: &model.User{
				ID:              user.ID,
				Name:            user.Name,
				Email:           user.Email,
				EmailVerifiedAt: (*int64)(&user.EmailVerifiedAt.Int64),
				LastSeenAt:      user.LastSeenAt,
				CreatedAt:       user.CreatedAt,
				UpdatedAt:       user.UpdatedAt,
				DeletedAt:       (*int64)(&user.DeletedAt.Int64),
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := len(users) == int(first)

	return &model.UserConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// Folders is the resolver for the folders field.
func (r *groupResolver) Folders(ctx context.Context, obj *model.Group, first int64, after *int64, filter *model.FolderFilter, sort *model.FolderSort) (*model.FolderConnection, error) {
	// Decode the cursor (if provided)
	var offset int64
	if after != nil { // Check if `after` is provided (non-nil)
		offset = *after
	}

	// Prepare sorting
	sortField := "NAME" // Default sort field
	if sort != nil {
		sortField = string(sort.Field)
	}

	sortOrder := "ASC" // Default sort order
	if sort != nil {
		sortOrder = string(sort.Order)
	}

	// Prepare filter values
	var nameFilter, slugFilter, descriptionFilter *string
	if filter != nil {
		nameFilter = filter.Name
		slugFilter = filter.Slug
		descriptionFilter = filter.Description
	}

	// Fetch users using the SQL query method for group ID
	folders, err := r.DB.GetPaginatedFoldersByGroupID(ctx, generated.GetPaginatedFoldersByGroupIDParams{
		GroupID:           pgtype.Int8{Int64: obj.ID, Valid: true}, // Group ID from the Group object
		Limit:             int32(first),                            // Limit based on 'first' argument
		Offset:            int32(offset),                           // Offset based on 'after' cursor
		NameFilter:        nameFilter,                              // Name filter (optional)
		SlugFilter:        slugFilter,                              // Slug filter (optional)
		DescriptionFilter: descriptionFilter,                       // Slug filter (optional)
		SortField:         sortField,                               // Sorting field
		SortOrder:         sortOrder,                               // Sorting order
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch folders for group %d: %v", obj.ID, err)
	}

	// Prepare edges and PageInfo for the connection
	edges := make([]*model.FolderEdge, len(folders))
	for i, folder := range folders {
		edges[i] = &model.FolderEdge{
			Cursor: strconv.FormatInt(offset+int64(i)+1, 10), // Create cursor from index
			Node: &model.Folder{
				ID:          folder.ID,
				Name:        folder.Name,
				Slug:        folder.Slug,
				Description: folder.Description,
				CreatedAt:   folder.CreatedAt,
				UpdatedAt:   folder.UpdatedAt,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := len(folders) == int(first)

	return &model.FolderConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// Files is the resolver for the files field.
func (r *groupResolver) Files(ctx context.Context, obj *model.Group, first int64, after *int64, filter *model.FileFilter, sort *model.FileSort) (*model.FileConnection, error) {
	// Decode the cursor (if provided)
	var offset int64
	if after != nil { // Check if `after` is provided (non-nil)
		offset = *after
	}

	// Prepare sorting
	sortField := "NAME" // Default sort field
	if sort != nil {
		sortField = string(sort.Field)
	}

	sortOrder := "ASC" // Default sort order
	if sort != nil {
		sortOrder = string(sort.Order)
	}

	// Prepare filter values
	var nameFilter, slugFilter *string
	if filter != nil {
		nameFilter = filter.Name
		slugFilter = filter.Slug
	}

	// Fetch users using the SQL query method for group ID
	files, err := r.DB.GetPaginatedFilesByGroupID(ctx, generated.GetPaginatedFilesByGroupIDParams{
		GroupID:    pgtype.Int8{Int64: obj.ID, Valid: true}, // Group ID from the Group object
		Limit:      int32(first),                            // Limit based on 'first' argument
		Offset:     int32(offset),                           // Offset based on 'after' cursor
		NameFilter: nameFilter,                              // Name filter (optional)
		SlugFilter: slugFilter,                              // Slug filter (optional)
		SortField:  sortField,                               // Sorting field
		SortOrder:  sortOrder,                               // Sorting order
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch files for group %d: %v", obj.ID, err)
	}

	// Prepare edges and PageInfo for the connection
	edges := make([]*model.FileEdge, len(files))
	for i, file := range files {
		edges[i] = &model.FileEdge{
			Cursor: strconv.FormatInt(offset+int64(i)+1, 10), // Create cursor from index
			Node: &model.File{
				ID:        file.ID,
				Name:      file.Name,
				Slug:      file.Slug,
				URL:       file.Url,
				FolderID:  file.FolderID,
				CreatedAt: file.CreatedAt,
				UpdatedAt: file.UpdatedAt,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := len(files) == int(first)

	return &model.FileConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroup) (*model.Group, error) {
	// Validate input
	if err := input.Validate(); err != nil {
		// Call the reusable validation error formatter
		return nil, utils.FormatValidationErrors(err)
	}

	// Call the generated CreateGroup function with the params
	group, err := r.DB.CreateGroup(ctx, generated.CreateGroupParams{
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, err
	}

	// Return the newly created group
	return &model.Group{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
	}, nil
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, input model.UpdateGroup) (*model.Group, error) {
	// Check if the group exists
	_, err := r.DB.GetFolder(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("group not found: %w", err)
	}

	// Update the group
	err = r.DB.UpdateGroup(ctx, generated.UpdateGroupParams{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update group: %w", err)
	}

	// Fetch the updated group
	updatedGroup, err := r.DB.GetGroup(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated group: %w", err)
	}

	// Map the SQLC model to the GraphQL model
	return &model.Group{
		ID:          updatedGroup.ID,
		Name:        updatedGroup.Name,
		Description: updatedGroup.Description,
		CreatedAt:   updatedGroup.CreatedAt,
		UpdatedAt:   updatedGroup.UpdatedAt,
	}, nil
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id int64) (bool, error) {
	// Check if the group exists (optional)
	_, err := r.DB.GetGroup(ctx, id)
	if err != nil {
		return false, fmt.Errorf("group not found: %w", err)
	}

	// Attempt to delete the group
	err = r.DB.DeleteGroup(ctx, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete group: %w", err)
	}
	return true, nil
}

// AddUserToGroup is the resolver for the addUserToGroup field.
func (r *mutationResolver) AddUserToGroup(ctx context.Context, groupID int64, userID int64) (bool, error) {
	// Check if the group exists
	_, groupErr := r.DB.GetGroup(ctx, groupID)
	if groupErr != nil {
		return false, fmt.Errorf("group not found: %w", groupErr)
	}

	// Check if the user exists
	_, err := r.DB.GetUser(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("user not found: %w", err)
	}

	err = r.DB.AddUserToGroup(ctx, generated.AddUserToGroupParams{
		GroupID: groupID,
		UserID:  userID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to add user to group: %w", err)
	}
	return true, nil
}

// RemoveUserFromGroup is the resolver for the removeUserFromGroup field.
func (r *mutationResolver) RemoveUserFromGroup(ctx context.Context, groupID int64, userID int64) (bool, error) {
	// Check if the group exists
	_, groupErr := r.DB.GetGroup(ctx, groupID)
	if groupErr != nil {
		return false, fmt.Errorf("group not found: %w", groupErr)
	}

	// Check if the user exists
	_, err := r.DB.GetUser(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("user not found: %w", err)
	}

	err = r.DB.RemoveUserFromGroup(ctx, generated.RemoveUserFromGroupParams{
		GroupID: groupID,
		UserID:  userID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to remove user to group: %w", err)
	}
	return true, nil
}

// AddFolderToGroup is the resolver for the addFolderToGroup field.
func (r *mutationResolver) AddFolderToGroup(ctx context.Context, groupID int64, folderID int64) (bool, error) {
	// Check if the group exists
	_, groupErr := r.DB.GetGroup(ctx, groupID)
	if groupErr != nil {
		return false, fmt.Errorf("group not found: %w", groupErr)
	}

	// Check if the folder exists
	_, err := r.DB.GetFolder(ctx, folderID)
	if err != nil {
		return false, fmt.Errorf("folder not found: %w", err)
	}

	err = r.DB.AddFolderToGroup(ctx, generated.AddFolderToGroupParams{
		GroupID:  groupID,
		FolderID: folderID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to add folder to group: %w", err)
	}
	return true, nil
}

// RemoveFolderFromGroup is the resolver for the removeFolderFromGroup field.
func (r *mutationResolver) RemoveFolderFromGroup(ctx context.Context, groupID int64, folderID int64) (bool, error) {
	// Check if the group exists
	_, groupErr := r.DB.GetGroup(ctx, groupID)
	if groupErr != nil {
		return false, fmt.Errorf("group not found: %w", groupErr)
	}

	// Check if the folder exists
	_, err := r.DB.GetFolder(ctx, folderID)
	if err != nil {
		return false, fmt.Errorf("folder not found: %w", err)
	}

	err = r.DB.RemoveFolderFromGroup(ctx, generated.RemoveFolderFromGroupParams{
		GroupID:  groupID,
		FolderID: folderID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to remove folder to group: %w", err)
	}
	return true, nil
}

// AddFileToGroup is the resolver for the addFileToGroup field.
func (r *mutationResolver) AddFileToGroup(ctx context.Context, groupID int64, fileID int64) (bool, error) {
	// Check if the group exists
	_, groupErr := r.DB.GetGroup(ctx, groupID)
	if groupErr != nil {
		return false, fmt.Errorf("group not found: %w", groupErr)
	}

	// Check if the file exists
	_, err := r.DB.GetFile(ctx, fileID)
	if err != nil {
		return false, fmt.Errorf("file not found: %w", err)
	}

	err = r.DB.AddFileToGroup(ctx, generated.AddFileToGroupParams{
		GroupID: groupID,
		FileID:  fileID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to add file to group: %w", err)
	}
	return true, nil
}

// RemoveFileFromGroup is the resolver for the removeFileFromGroup field.
func (r *mutationResolver) RemoveFileFromGroup(ctx context.Context, groupID int64, fileID int64) (bool, error) {
	// Check if the group exists
	_, groupErr := r.DB.GetGroup(ctx, groupID)
	if groupErr != nil {
		return false, fmt.Errorf("group not found: %w", groupErr)
	}

	// Check if the folder exists
	_, err := r.DB.GetFile(ctx, fileID)
	if err != nil {
		return false, fmt.Errorf("file not found: %w", err)
	}

	err = r.DB.RemoveFileFromGroup(ctx, generated.RemoveFileFromGroupParams{
		GroupID: groupID,
		FileID:  fileID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to remove file to group: %w", err)
	}
	return true, nil
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
	// Fetch groups using sqlc
	rows, err := r.DB.ListGroups(ctx) // Assuming ListGroups is the sqlc query method
	if err != nil {
		return nil, err
	}

	// Map sqlc rows to GraphQL models
	var groups []*model.Group
	for _, row := range rows {
		groups = append(groups, &model.Group{
			ID:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			CreatedAt:   row.CreatedAt, // Or use row.CreatedAt.Time.String()
			UpdatedAt:   row.UpdatedAt, // Or use row.UpdatedAt.Time.String()
		})
	}

	return groups, nil
}

// GetGroup is the resolver for the getGroup field.
func (r *queryResolver) GetGroup(ctx context.Context, id int64) (*model.Group, error) {
	// Call the generated GetGroup query
	group, err := r.DB.GetGroup(ctx, id) // assuming input.ID is of type string
	if err != nil {
		return nil, fmt.Errorf("failed to get folder: %w", err)
	}

	// Convert the SQL result to GraphQL model
	return &model.Group{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt, // assuming you're using timestamptz
		UpdatedAt:   group.UpdatedAt, // assuming you're using timestamptz
	}, nil
}

// Group returns graph.GroupResolver implementation.
func (r *Resolver) Group() graph.GroupResolver { return &groupResolver{r} }

type groupResolver struct{ *Resolver }
