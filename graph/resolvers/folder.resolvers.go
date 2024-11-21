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

// Groups is the resolver for the groups field.
func (r *folderResolver) Groups(ctx context.Context, obj *model.Folder, first int64, after *int64, filter *model.GroupFilter, sort *model.GroupSort) (*model.GroupConnection, error) {
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
	groups, err := r.DB.GetPaginatedGroupsByFolderID(ctx, generated.GetPaginatedGroupsByFolderIDParams{
		FolderID:          pgtype.Int8{Int64: obj.ID, Valid: true}, // Group ID from the Group object
		Limit:             int32(first),                            // Limit based on 'first' argument
		Offset:            int32(offset),                           // Offset based on 'after' cursor
		NameFilter:        nameFilter,                              // Name filter (optional)
		DescriptionFilter: descriptionFilter,                       // Email filter (optional)
		SortField:         sortField,                               // Sorting field
		SortOrder:         sortOrder,                               // Sorting order
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups for folder %d: %v", obj.ID, err)
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
	hasNextPage := len(groups) == int(first)

	return &model.GroupConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// Files is the resolver for the files field.
func (r *folderResolver) Files(ctx context.Context, obj *model.Folder, first int64, after *int64, filter *model.FileFilter, sort *model.FileSort) (*model.FileConnection, error) {
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

	// Fetch groups using the SQL query method for folder ID
	files, err := r.DB.GetPaginatedFilesByFolderID(ctx, generated.GetPaginatedFilesByFolderIDParams{
		FolderID:   pgtype.Int8{Int64: obj.ID, Valid: true}, // Group ID from the Group object
		Limit:      int32(first),                            // Limit based on 'first' argument
		Offset:     int32(offset),                           // Offset based on 'after' cursor
		NameFilter: nameFilter,                              // Name filter (optional)
		SlugFilter: slugFilter,                              // Email filter (optional)
		SortField:  sortField,                               // Sorting field
		SortOrder:  sortOrder,                               // Sorting order
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch files for folder %d: %v", obj.ID, err)
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

// CreateFolder is the resolver for the createFolder field.
func (r *mutationResolver) CreateFolder(ctx context.Context, input model.CreateFolder) (*model.Folder, error) {
	// Validate input
	if err := input.Validate(); err != nil {
		// Call the reusable validation error formatter
		return nil, utils.FormatValidationErrors(err)
	}

	// Call the generated CreateFolder function with the params
	folder, err := r.DB.CreateFolder(ctx, generated.CreateFolderParams{
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
	})
	if err != nil {
		return nil, err
	}

	// Return the newly created folder
	return &model.Folder{
		ID:          folder.ID,
		Name:        folder.Name,
		Slug:        folder.Slug,
		Description: folder.Description,
		CreatedAt:   folder.CreatedAt,
		UpdatedAt:   folder.UpdatedAt,
	}, nil
}

// UpdateFolder is the resolver for the UpdateFolder field.
func (r *mutationResolver) UpdateFolder(ctx context.Context, input model.UpdateFolder) (*model.Folder, error) {
	// Check if the folder exists
	_, err := r.DB.GetFolder(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("folder not found: %w", err)
	}

	// Update the folder
	err = r.DB.UpdateFolder(ctx, generated.UpdateFolderParams{
		ID:          input.ID,
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update folder: %w", err)
	}

	// Fetch the updated folder
	updatedFolder, err := r.DB.GetFolder(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated folder: %w", err)
	}

	// Map the SQLC model to the GraphQL model
	return &model.Folder{
		ID:          updatedFolder.ID,
		Name:        updatedFolder.Name,
		Slug:        updatedFolder.Slug,
		Description: updatedFolder.Description,
		CreatedAt:   updatedFolder.CreatedAt,
		UpdatedAt:   updatedFolder.UpdatedAt,
	}, nil
}

// DeleteFolder is the resolver for the deleteFolder field.
func (r *mutationResolver) DeleteFolder(ctx context.Context, id int64) (bool, error) {
	// Check if the folder exists (optional)
	_, err := r.DB.GetFolder(ctx, id)
	if err != nil {
		return false, fmt.Errorf("folder not found: %w", err)
	}

	// Attempt to delete the folder
	err = r.DB.DeleteFolder(ctx, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete folder: %w", err)
	}
	return true, nil
}

// Folders is the resolver for the folders field.
func (r *queryResolver) Folders(ctx context.Context, first int64, after *int64, filter *model.FolderFilter, sort *model.FolderSort) (*model.FolderConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_folder"}

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
	var nameFilter, slugFilter, descriptionFilter *string
	if filter != nil {
		nameFilter = filter.Name
		slugFilter = filter.Slug
		descriptionFilter = filter.Description
	}
	// Fetch users using sqlc
	folders, err := r.DB.PaginatedFolders(ctx, generated.PaginatedFoldersParams{
		Limit:             int32(first),
		Offset:            int32(offset),
		NameFilter:        nameFilter,
		SlugFilter:        slugFilter,
		DescriptionFilter: descriptionFilter,
		SortField:         sortField,
		SortOrder:         sortOrder,
	}) // Assuming ListFiles is the sqlc query method
	if err != nil {
		return nil, fmt.Errorf("failed to query files: %v", err)
	}

	// Prepare edges and PageInfo
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

// GetFolder is the resolver for the getFolder field.
func (r *queryResolver) GetFolder(ctx context.Context, id int64) (*model.Folder, error) {
	// Call the generated GetFolder query
	folder, err := r.DB.GetFolder(ctx, id) // assuming input.ID is of type string
	if err != nil {
		return nil, fmt.Errorf("failed to get folder: %w", err)
	}

	// Convert the SQL result to GraphQL model
	return &model.Folder{
		ID:          folder.ID,
		Name:        folder.Name,
		Slug:        folder.Slug,
		Description: folder.Description,
		CreatedAt:   folder.CreatedAt, // assuming you're using timestamptz
		UpdatedAt:   folder.UpdatedAt, // assuming you're using timestamptz
	}, nil
}

// Folder returns graph.FolderResolver implementation.
func (r *Resolver) Folder() graph.FolderResolver { return &folderResolver{r} }

type folderResolver struct{ *Resolver }
