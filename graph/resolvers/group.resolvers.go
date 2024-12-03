package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"fmt"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/graph"
	"github.com/awanishnathpandey/leaf/graph/model"
	"github.com/awanishnathpandey/leaf/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

// Users is the resolver for the users field.
func (r *groupResolver) Users(ctx context.Context, obj *model.Group, first int64, after *int64, filter *model.UserFilter, sort *model.UserSort) (*model.UserConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_user"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}

	// Prepare sorting
	sortField := "NAME" // Default sort field
	sortOrder := "ASC"  // Default sort order
	if sort != nil {
		// Prepare sorting using the utility
		sortField, sortOrder = utils.PrepareSorting("NAME", "ASC", string(sort.Field), string(sort.Order))
	}

	// Calculate pagination and sorting
	offset, first := utils.PreparePaginationParams(after, first)

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

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.GetPaginatedUsersByGroupIDCount(ctx, generated.GetPaginatedUsersByGroupIDCountParams{
		GroupID:     pgtype.Int8{Int64: obj.ID, Valid: true},
		NameFilter:  nameFilter,
		EmailFilter: emailFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query users count for group %d: %v", obj.ID, err)
	}

	// Prepare edges and PageInfo for the connection
	edges := make([]*model.UserEdge, len(users))
	for i, user := range users {
		edges[i] = &model.UserEdge{
			Cursor: utils.GenerateCursor(offset, int64(i)), // Create cursor from index
			Node: &model.User{
				ID:              user.ID,
				FirstName:       user.FirstName,
				LastName:        user.LastName,
				Email:           user.Email,
				JobTitle:        &user.JobTitle.String,
				LineOfBusiness:  &user.LineOfBusiness.String,
				LineManager:     &user.LineManager.String,
				EmailVerifiedAt: (*int64)(&user.EmailVerifiedAt.Int64),
				LastSeenAt:      user.LastSeenAt,
				CreatedAt:       user.CreatedAt,
				UpdatedAt:       user.UpdatedAt,
				DeletedAt:       (*int64)(&user.DeletedAt.Int64),
				CreatedBy:       user.CreatedBy,
				UpdatedBy:       user.UpdatedBy,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := utils.CalculateHasNextPage(offset, int64(len(users)), totalCount)

	return &model.UserConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// Folders is the resolver for the folders field.
func (r *groupResolver) Folders(ctx context.Context, obj *model.Group, first int64, after *int64, filter *model.FolderFilter, sort *model.FolderSort) (*model.FolderConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_folder"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}

	// Prepare sorting
	sortField := "NAME" // Default sort field
	sortOrder := "ASC"  // Default sort order
	if sort != nil {
		// Prepare sorting using the utility
		sortField, sortOrder = utils.PrepareSorting("NAME", "ASC", string(sort.Field), string(sort.Order))
	}

	// Calculate pagination and sorting
	offset, first := utils.PreparePaginationParams(after, first)

	// Prepare filter values
	var nameFilter, slugFilter, descriptionFilter *string
	if filter != nil {
		nameFilter = filter.Name
		slugFilter = filter.Slug
		descriptionFilter = filter.Description
	}

	// Fetch folders using the SQL query method for group ID
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

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.GetPaginatedFoldersByGroupIDCount(ctx, generated.GetPaginatedFoldersByGroupIDCountParams{
		GroupID:           pgtype.Int8{Int64: obj.ID, Valid: true},
		NameFilter:        nameFilter,
		SlugFilter:        slugFilter,
		DescriptionFilter: descriptionFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query folders count for group %d: %v", obj.ID, err)
	}

	// Prepare edges and PageInfo for the connection
	edges := make([]*model.FolderEdge, len(folders))
	for i, folder := range folders {
		edges[i] = &model.FolderEdge{
			Cursor: utils.GenerateCursor(offset, int64(i)), // Create cursor from index
			Node: &model.Folder{
				ID:          folder.ID,
				Name:        folder.Name,
				Slug:        folder.Slug,
				Description: folder.Description,
				CreatedAt:   folder.CreatedAt,
				UpdatedAt:   folder.UpdatedAt,
				CreatedBy:   folder.CreatedBy,
				UpdatedBy:   folder.UpdatedBy,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := utils.CalculateHasNextPage(offset, int64(len(folders)), totalCount)

	return &model.FolderConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// Files is the resolver for the files field.
func (r *groupResolver) Files(ctx context.Context, obj *model.Group, first int64, after *int64, filter *model.FileFilter, sort *model.FileSort) (*model.FileConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}

	// Prepare sorting
	sortField := "NAME" // Default sort field
	sortOrder := "ASC"  // Default sort order
	if sort != nil {
		// Prepare sorting using the utility
		sortField, sortOrder = utils.PrepareSorting("NAME", "ASC", string(sort.Field), string(sort.Order))
	}

	// Calculate pagination and sorting
	offset, first := utils.PreparePaginationParams(after, first)

	// Prepare filter values
	var nameFilter, slugFilter *string
	if filter != nil {
		nameFilter = filter.Name
		slugFilter = filter.Slug
	}

	// Fetch files using the SQL query method for group ID
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

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.GetPaginatedFilesByGroupIDCount(ctx, generated.GetPaginatedFilesByGroupIDCountParams{
		GroupID:    pgtype.Int8{Int64: obj.ID, Valid: true},
		NameFilter: nameFilter,
		SlugFilter: slugFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query files count for group %d: %v", obj.ID, err)
	}

	// Prepare edges and PageInfo for the connection
	edges := make([]*model.FileEdge, len(files))
	for i, file := range files {
		edges[i] = &model.FileEdge{
			Cursor: utils.GenerateCursor(offset, int64(i)), // Create cursor from index
			Node: &model.File{
				ID:           file.ID,
				Name:         file.Name,
				Slug:         file.Slug,
				FilePath:     file.FilePath,
				FileBytes:    file.FileBytes,
				AutoDownload: file.AutoDownload,
				FolderID:     file.FolderID,
				CreatedAt:    file.CreatedAt,
				UpdatedAt:    file.UpdatedAt,
				CreatedBy:    file.CreatedBy,
				UpdatedBy:    file.UpdatedBy,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := utils.CalculateHasNextPage(offset, int64(len(files)), totalCount)

	return &model.FileConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroup) (*model.Group, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "create_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Validate input
	if err := input.Validate(); err != nil {
		// Call the reusable validation error formatter
		return nil, utils.FormatValidationErrors(err)
	}

	// Call the generated CreateGroup function with the params
	group, err := r.DB.CreateGroup(ctx, generated.CreateGroupParams{
		Name:        input.Name,
		Description: input.Description,
		CreatedBy:   ctx.Value("userEmail").(string),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	// Return the newly created group
	return &model.Group{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
		CreatedBy:   group.CreatedBy,
		UpdatedBy:   group.UpdatedBy,
	}, nil
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, input model.UpdateGroup) (*model.Group, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Check if the group exists
	_, err := r.DB.GetFolder(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("group not found: %w", err)
	}

	// Call the sqlc generated query to update the group in the database
	group, err := r.DB.UpdateGroup(ctx, generated.UpdateGroupParams{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		UpdatedBy:   ctx.Value("userEmail").(string),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update group: %w", err)
	}

	// Map the SQLC model to the GraphQL model
	return &model.Group{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
		CreatedBy:   group.CreatedBy,
		UpdatedBy:   group.UpdatedBy,
	}, nil
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "delete_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
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

// DeleteGroups is the resolver for the deleteGroups field.
func (r *mutationResolver) DeleteGroups(ctx context.Context, ids []int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "delete_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
	// Validate that all IDs exist
	existingFiles, err := r.DB.GetGroupsByIDs(ctx, ids)
	if err != nil {
		return false, fmt.Errorf("failed to fetch groups: %w", err)
	}
	if len(existingFiles) != len(ids) {
		return false, fmt.Errorf("validation failed: some groups do not exist")
	}

	// Proceed to delete the files
	err = r.DB.DeleteGroupsByIDs(ctx, ids)
	if err != nil {
		return false, fmt.Errorf("failed to delete groups: %w", err)
	}

	// All files successfully deleted
	return true, nil
}

// AddUserToGroup is the resolver for the addUserToGroup field.
func (r *mutationResolver) AddUserToGroup(ctx context.Context, groupID int64, userID int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_user", "update_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
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
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_user", "update_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
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
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_folder", "update_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
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
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_folder", "update_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
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
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_file", "update_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
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
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_file", "update_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
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
func (r *queryResolver) Groups(ctx context.Context, first int64, after *int64, filter *model.GroupFilter, sort *model.GroupSort) (*model.GroupConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}

	// Prepare sorting
	sortField := "NAME" // Default sort field
	sortOrder := "ASC"  // Default sort order
	if sort != nil {
		// Prepare sorting using the utility
		sortField, sortOrder = utils.PrepareSorting("NAME", "ASC", string(sort.Field), string(sort.Order))
	}

	// Calculate pagination and sorting
	offset, first := utils.PreparePaginationParams(after, first)

	// Prepare filter values
	var nameFilter, descriptionFilter *string
	if filter != nil {
		nameFilter = filter.Name
		descriptionFilter = filter.Description
	}
	// Fetch groups using sqlc
	groups, err := r.DB.PaginatedGroups(ctx, generated.PaginatedGroupsParams{
		Limit:             int32(first),
		Offset:            int32(offset),
		NameFilter:        nameFilter,
		DescriptionFilter: descriptionFilter,
		SortField:         sortField,
		SortOrder:         sortOrder,
	}) // Assuming ListGroups is the sqlc query method
	if err != nil {
		return nil, fmt.Errorf("failed to query groups: %v", err)
	}

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.PaginatedGroupsCount(ctx, generated.PaginatedGroupsCountParams{
		NameFilter:        nameFilter,
		DescriptionFilter: descriptionFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query groups count: %v", err)
	}

	// Prepare edges and PageInfo
	edges := make([]*model.GroupEdge, len(groups))
	for i, group := range groups {
		edges[i] = &model.GroupEdge{
			Cursor: utils.GenerateCursor(offset, int64(i)), // Create cursor from index
			Node: &model.Group{
				ID:          group.ID,
				Name:        group.Name,
				Description: group.Description,
				CreatedAt:   group.CreatedAt,
				UpdatedAt:   group.UpdatedAt,
				CreatedBy:   group.CreatedBy,
				UpdatedBy:   group.UpdatedBy,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := utils.CalculateHasNextPage(offset, int64(len(groups)), totalCount)

	return &model.GroupConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// GetGroup is the resolver for the getGroup field.
func (r *queryResolver) GetGroup(ctx context.Context, id int64) (*model.Group, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_group"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
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
