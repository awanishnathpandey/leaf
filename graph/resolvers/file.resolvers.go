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

// Folder is the resolver for the folder field.
func (r *fileResolver) Folder(ctx context.Context, obj *model.File) (*model.Folder, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_folder", "read_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Fetch the folder by folderID
	folder, err := r.DB.GetFolder(ctx, obj.FolderID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch folder: %w", err)
	}

	// Map the folder data to the GraphQL model
	return &model.Folder{
		ID:          folder.ID,
		Name:        folder.Name,
		Slug:        folder.Slug,
		Description: folder.Description,
		CreatedAt:   folder.CreatedAt,
		UpdatedAt:   folder.UpdatedAt,
	}, nil
}

// Groups is the resolver for the groups field.
func (r *fileResolver) Groups(ctx context.Context, obj *model.File, first int64, after *int64, filter *model.GroupFilter, sort *model.GroupSort) (*model.GroupConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_group", "read_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
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
	var nameFilter, descriptionFilter *string
	if filter != nil {
		nameFilter = filter.Name
		descriptionFilter = filter.Description
	}

	// Fetch groups using the SQL query method for rolder ID
	groups, err := r.DB.GetPaginatedGroupsByFileID(ctx, generated.GetPaginatedGroupsByFileIDParams{
		FileID:            pgtype.Int8{Int64: obj.ID, Valid: true}, // Group ID from the Group object
		Limit:             int32(first),                            // Limit based on 'first' argument
		Offset:            int32(offset),                           // Offset based on 'after' cursor
		NameFilter:        nameFilter,                              // Name filter (optional)
		DescriptionFilter: descriptionFilter,                       // Email filter (optional)
		SortField:         sortField,                               // Sorting field
		SortOrder:         sortOrder,                               // Sorting order
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups for file %d: %v", obj.ID, err)
	}

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.GetPaginatedGroupsByFileIDCount(ctx, generated.GetPaginatedGroupsByFileIDCountParams{
		FileID:            pgtype.Int8{Int64: obj.ID, Valid: true},
		NameFilter:        nameFilter,
		DescriptionFilter: descriptionFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query groups count for file %d: %v", obj.ID, err)
	}

	// Prepare edges and PageInfo for the connection
	edges := make([]*model.GroupEdge, len(groups))
	for i, group := range groups {
		edges[i] = &model.GroupEdge{
			Cursor: strconv.FormatInt(offset+int64(i)+1, 10), // Create cursor from index
			Node: &model.Group{
				ID:          group.ID,
				Name:        group.Name,
				Description: group.Description,
				CreatedAt:   group.CreatedAt,
				UpdatedAt:   group.UpdatedAt,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := offset+int64(len(groups)) < totalCount

	return &model.GroupConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// CreateFile is the resolver for the createFile field.
func (r *mutationResolver) CreateFile(ctx context.Context, input model.CreateFile) (*model.File, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "create_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Validate input
	if err := input.Validate(); err != nil {
		// Call the reusable validation error formatter
		return nil, utils.FormatValidationErrors(err)
	}

	// Call the sqlc generated query to insert the file into the database
	file, err := r.DB.CreateFile(ctx, generated.CreateFileParams{
		Name:     input.Name,
		Slug:     input.Slug,
		Url:      input.URL,
		FolderID: input.FolderID, // Ensure the Folder ID is passed correctly
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	// Map the result from sqlc to the GraphQL model
	return &model.File{
		ID:        file.ID,
		Name:      file.Name,
		Slug:      file.Slug,
		URL:       file.Url,
		FolderID:  file.FolderID,
		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
	}, nil
}

// UpdateFile is the resolver for the updateFile field.
func (r *mutationResolver) UpdateFile(ctx context.Context, input model.UpdateFile) (*model.File, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Check if the file exists
	_, err := r.DB.GetFile(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("file not found: %w", err)
	}
	// Call the sqlc generated query to update the file in the database
	file, err := r.DB.UpdateFile(ctx, generated.UpdateFileParams{
		ID:   input.ID,
		Name: input.Name,
		Slug: input.Slug,
		Url:  input.URL,
	})
	if err != nil {
		return nil, err
	}

	// Map the result from sqlc to the GraphQL model
	return &model.File{
		ID:        file.ID,
		Name:      file.Name,
		Slug:      file.Slug,
		URL:       file.Url,
		FolderID:  file.FolderID,
		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
	}, nil
}

// DeleteFile is the resolver for the deleteFile field.
func (r *mutationResolver) DeleteFile(ctx context.Context, id int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "delete_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
	// Check if the file exists (optional)
	_, err := r.DB.GetFile(ctx, id)
	if err != nil {
		return false, fmt.Errorf("file not found: %w", err)
	}

	// Attempt to delete the file
	err = r.DB.DeleteFile(ctx, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete file: %w", err)
	}
	return true, nil
}

// DeleteFiles is the resolver for the deleteFiles field.
func (r *mutationResolver) DeleteFiles(ctx context.Context, ids []int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "delete_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}

	// Validate that all IDs exist
	existingFiles, err := r.DB.GetFilesByIDs(ctx, ids)
	if err != nil {
		return false, fmt.Errorf("failed to fetch files: %w", err)
	}
	if len(existingFiles) != len(ids) {
		return false, fmt.Errorf("validation failed: some files do not exist")
	}

	// Proceed to delete the files
	err = r.DB.DeleteFilesByIDs(ctx, ids)
	if err != nil {
		return false, fmt.Errorf("failed to delete files: %w", err)
	}

	// All files successfully deleted
	return true, nil
}

// Files is the resolver for the files field.
func (r *queryResolver) Files(ctx context.Context, first int64, after *int64, filter *model.FileFilter, sort *model.FileSort) (*model.FileConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
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

	// Prepare sorting
	sortOrder := "ASC" // Default sort field
	if sort != nil {
		sortOrder = string(sort.Order)
	}

	// Prepare filter values
	var nameFilter, slugFilter *string
	if filter != nil {
		nameFilter = filter.Name
		slugFilter = filter.Slug
	}
	// Fetch files using sqlc
	files, err := r.DB.PaginatedFiles(ctx, generated.PaginatedFilesParams{
		Limit:      int32(first),
		Offset:     int32(offset),
		NameFilter: nameFilter,
		SlugFilter: slugFilter,
		SortField:  sortField,
		SortOrder:  sortOrder,
	}) // Assuming ListFiles is the sqlc query method
	if err != nil {
		return nil, fmt.Errorf("failed to query files: %v", err)
	}

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.PaginatedFilesCount(ctx, generated.PaginatedFilesCountParams{
		NameFilter: nameFilter,
		SlugFilter: slugFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query files count: %v", err)
	}

	// Prepare edges and PageInfo
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
	hasNextPage := offset+int64(len(files)) < totalCount

	return &model.FileConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// GetFile is the resolver for the getFile field.
func (r *queryResolver) GetFile(ctx context.Context, id int64) (*model.File, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_file"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Call the generated GetFile query
	file, err := r.DB.GetFile(ctx, id) // assuming input.ID is of type string
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}

	// Convert the SQL result to GraphQL model
	return &model.File{
		ID:        file.ID,
		Name:      file.Name,
		Slug:      file.Slug,
		URL:       file.Url,
		FolderID:  file.FolderID,
		CreatedAt: file.CreatedAt, // assuming you're using timestamptz
		UpdatedAt: file.UpdatedAt, // assuming you're using timestamptz
	}, nil
}

// GetFilesByFolder is the resolver for the getFilesByFolder field.
func (r *queryResolver) GetFilesByFolder(ctx context.Context, folderID int64) ([]*model.File, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_file", "read_folder"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	files, err := r.DB.GetFilesByFolder(ctx, folderID)
	if err != nil {
		return nil, err
	}
	var result []*model.File
	for _, file := range files {
		result = append(result, &model.File{
			ID:        file.ID,
			Name:      file.Name,
			Slug:      file.Slug,
			URL:       file.Url,
			FolderID:  file.FolderID,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
		})
	}
	return result, nil
}

// File returns graph.FileResolver implementation.
func (r *Resolver) File() graph.FileResolver { return &fileResolver{r} }

type fileResolver struct{ *Resolver }
