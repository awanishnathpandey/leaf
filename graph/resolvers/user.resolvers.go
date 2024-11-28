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

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "create_user"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Validate input
	if err := input.Validate(); err != nil {
		// Call the reusable validation error formatter
		return nil, utils.FormatValidationErrors(err)
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Call the generated CreateUser function with the params
	user, err := r.DB.CreateUser(ctx, generated.CreateUserParams{
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashedPassword,
		CreatedBy: ctx.Value("userEmail").(string),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Return the newly created user
	return &model.User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		EmailVerifiedAt: (*int64)(&user.EmailVerifiedAt.Int64),
		LastSeenAt:      user.LastSeenAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       (*int64)(&user.DeletedAt.Int64),
		CreatedBy:       user.CreatedBy,
		UpdatedBy:       user.UpdatedBy,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "update_user"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Check if the user exists
	_, err := r.DB.GetUser(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Call the sqlc generated query to update the user in the database
	user, err := r.DB.UpdateUser(ctx, generated.UpdateUserParams{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		UpdatedBy: ctx.Value("userEmail").(string),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Map the SQLC model to the GraphQL model
	return &model.User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		EmailVerifiedAt: (*int64)(&user.EmailVerifiedAt.Int64),
		LastSeenAt:      user.LastSeenAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       (*int64)(&user.DeletedAt.Int64),
		CreatedBy:       user.CreatedBy,
		UpdatedBy:       user.UpdatedBy,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "delete_user"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
	// Check if the user exists (optional)
	_, err := r.DB.GetUser(ctx, id)
	if err != nil {
		return false, fmt.Errorf("user not found: %w", err)
	}

	// Attempt to delete the user
	err = r.DB.DeleteUser(ctx, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}
	return true, nil
}

// DeleteUsers is the resolver for the deleteUsers field.
func (r *mutationResolver) DeleteUsers(ctx context.Context, ids []int64) (bool, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "delete_user"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return false, err
	}
	// Validate that all IDs exist
	existingFiles, err := r.DB.GetUsersByIDs(ctx, ids)
	if err != nil {
		return false, fmt.Errorf("failed to fetch users: %w", err)
	}
	if len(existingFiles) != len(ids) {
		return false, fmt.Errorf("validation failed: some users do not exist")
	}

	// Proceed to delete the files
	err = r.DB.DeleteUsersByIDs(ctx, ids)
	if err != nil {
		return false, fmt.Errorf("failed to delete users: %w", err)
	}

	// All files successfully deleted
	return true, nil
}

// UpdateUserEmailVerifiedAt is the resolver for the UpdateUserEmailVerifiedAt field.
func (r *mutationResolver) UpdateUserEmailVerifiedAt(ctx context.Context, id int64) (bool, error) {
	// Check if the user exists (optional)
	_, err := r.DB.GetUser(ctx, id)
	if err != nil {
		return false, fmt.Errorf("user not found: %w", err)
	}

	// Attempt to update the user email verified at
	err = r.DB.UpdateUserEmailVerifiedAt(ctx, id)
	if err != nil {
		return false, fmt.Errorf("failed to update user email verified at: %w", err)
	}
	return true, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, first int64, after *int64, filter *model.UserFilter, sort *model.UserSort) (*model.UserConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_user"}

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
	var nameFilter, emailFilter *string
	if filter != nil {
		nameFilter = filter.Name
		emailFilter = filter.Email
	}
	// Fetch users using sqlc
	users, err := r.DB.PaginatedUsers(ctx, generated.PaginatedUsersParams{
		Limit:       int32(first),
		Offset:      int32(offset),
		NameFilter:  nameFilter,
		EmailFilter: emailFilter,
		SortField:   sortField,
		SortOrder:   sortOrder,
	}) // Assuming ListUsers is the sqlc query method
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %v", err)
	}

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.PaginatedUsersCount(ctx, generated.PaginatedUsersCountParams{
		NameFilter:  nameFilter,
		EmailFilter: emailFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query users count: %v", err)
	}

	// Prepare edges and PageInfo
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
				CreatedBy:       user.CreatedBy,
				UpdatedBy:       user.UpdatedBy,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := offset+int64(len(users)) < totalCount

	return &model.UserConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id int64) (*model.User, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_user"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Call the generated GetUser query
	user, err := r.DB.GetUser(ctx, id) // assuming input.ID is of type string
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Convert the SQL result to GraphQL model
	return &model.User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		EmailVerifiedAt: (*int64)(&user.EmailVerifiedAt.Int64),
		LastSeenAt:      user.LastSeenAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       (*int64)(&user.DeletedAt.Int64),
	}, nil
}

// GetUserByEmail is the resolver for the getUserByEmail field.
func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_user"}

	// Check if the user has the required permissions
	if err := utils.CheckUserPermissions(ctx, requiredPermissions, r.DB); err != nil {
		return nil, err
	}
	// Call the generated GetUserByEmail query
	user, err := r.DB.GetUserByEmail(ctx, email) // assuming email is of type string
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Convert the SQL result to GraphQL model
	return &model.User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		EmailVerifiedAt: (*int64)(&user.EmailVerifiedAt.Int64),
		LastSeenAt:      user.LastSeenAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       (*int64)(&user.DeletedAt.Int64),
	}, nil
}

// Groups is the resolver for the groups field.
func (r *userResolver) Groups(ctx context.Context, obj *model.User, first int64, after *int64, filter *model.GroupFilter, sort *model.GroupSort) (*model.GroupConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_group"}

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

	// Fetch groups using the SQL query method for role ID
	groups, err := r.DB.GetPaginatedGroupsByUserID(ctx, generated.GetPaginatedGroupsByUserIDParams{
		UserID:            pgtype.Int8{Int64: obj.ID, Valid: true}, // Group ID from the Group object
		Limit:             int32(first),                            // Limit based on 'first' argument
		Offset:            int32(offset),                           // Offset based on 'after' cursor
		NameFilter:        nameFilter,                              // Name filter (optional)
		DescriptionFilter: descriptionFilter,                       // Email filter (optional)
		SortField:         sortField,                               // Sorting field
		SortOrder:         sortOrder,                               // Sorting order
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups for user %d: %v", obj.ID, err)
	}

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.GetPaginatedGroupsByUserIDCount(ctx, generated.GetPaginatedGroupsByUserIDCountParams{
		UserID:            pgtype.Int8{Int64: obj.ID, Valid: true},
		NameFilter:        nameFilter,
		DescriptionFilter: descriptionFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query groups count for user %d: %v", obj.ID, err)
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
				CreatedBy:   group.CreatedBy,
				UpdatedBy:   group.UpdatedBy,
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

// Roles is the resolver for the roles field.
func (r *userResolver) Roles(ctx context.Context, obj *model.User, first int64, after *int64, filter *model.RoleFilter, sort *model.RoleSort) (*model.RoleConnection, error) {
	// Define the required permissions for this action
	requiredPermissions := []string{"all", "read_role"}

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

	// Fetch roles using the SQL query method for role ID
	roles, err := r.DB.GetPaginatedRolesByUserID(ctx, generated.GetPaginatedRolesByUserIDParams{
		UserID:            pgtype.Int8{Int64: obj.ID, Valid: true}, // Group ID from the Group object
		Limit:             int32(first),                            // Limit based on 'first' argument
		Offset:            int32(offset),                           // Offset based on 'after' cursor
		NameFilter:        nameFilter,                              // Name filter (optional)
		DescriptionFilter: descriptionFilter,                       // Description filter (optional)
		SortField:         sortField,                               // Sorting field
		SortOrder:         sortOrder,                               // Sorting order
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch roles for user %d: %v", obj.ID, err)
	}

	// Fetch filtered count using sqlc
	totalCount, err := r.DB.GetPaginatedRolesByUserIDCount(ctx, generated.GetPaginatedRolesByUserIDCountParams{
		UserID:            pgtype.Int8{Int64: obj.ID, Valid: true},
		NameFilter:        nameFilter,
		DescriptionFilter: descriptionFilter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query roles count for user %d: %v", obj.ID, err)
	}

	// Prepare edges and PageInfo for the connection
	edges := make([]*model.RoleEdge, len(roles))
	for i, role := range roles {
		edges[i] = &model.RoleEdge{
			Cursor: strconv.FormatInt(offset+int64(i)+1, 10), // Create cursor from index
			Node: &model.Role{
				ID:          role.ID,
				Name:        role.Name,
				Description: role.Description,
				CreatedAt:   role.CreatedAt,
				UpdatedAt:   role.UpdatedAt,
				CreatedBy:   role.CreatedBy,
				UpdatedBy:   role.UpdatedBy,
			},
		}
	}

	// Calculate hasNextPage
	hasNextPage := offset+int64(len(roles)) < totalCount

	return &model.RoleConnection{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: offset > 0,
		},
	}, nil
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
