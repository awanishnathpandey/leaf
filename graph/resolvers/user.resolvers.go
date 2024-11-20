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
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
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
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
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
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	// Check if the user exists
	_, err := r.DB.GetUser(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("folder not found: %w", err)
	}

	// Update the user
	err = r.DB.UpdateUser(ctx, generated.UpdateUserParams{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Fetch the updated user
	updatedUser, err := r.DB.GetUser(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated folder: %w", err)
	}

	// Map the SQLC model to the GraphQL model
	return &model.User{
		ID:              updatedUser.ID,
		Name:            updatedUser.Name,
		Email:           updatedUser.Email,
		EmailVerifiedAt: (*int64)(&updatedUser.EmailVerifiedAt.Int64),
		LastSeenAt:      updatedUser.LastSeenAt,
		CreatedAt:       updatedUser.CreatedAt,
		UpdatedAt:       updatedUser.UpdatedAt,
		DeletedAt:       (*int64)(&updatedUser.DeletedAt.Int64),
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (bool, error) {
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
		return false, fmt.Errorf("failed to delete user: %w", err)
	}
	return true, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, first int64, after *string, filter *model.UserFilter, sort *model.UserSort) (*model.UserConnection, error) {
	// Decode the cursor (if provided)
	var offset int
	if after != nil {
		cursor, err := strconv.Atoi(*after) // Convert string cursor to int
		if err != nil {
			return nil, fmt.Errorf("invalid cursor: %v", err)
		}
		offset = cursor
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

	// Prepare edges and PageInfo
	edges := make([]*model.UserEdge, len(users))
	for i, user := range users {
		edges[i] = &model.UserEdge{
			Cursor: strconv.Itoa(offset + i + 1), // Create cursor from index
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

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id int64) (*model.User, error) {
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
func (r *userResolver) Groups(ctx context.Context, obj *model.User) ([]*model.Group, error) {
	groups, err := r.DB.GetGroupsByUserID(ctx, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups for user %d: %w", obj.ID, err)
	}

	var result []*model.Group
	for _, group := range groups {
		result = append(result, &model.Group{
			ID:          group.ID,
			Name:        group.Name,
			Description: group.Description,
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
		})
	}
	return result, nil
}

// Roles is the resolver for the roles field.
func (r *userResolver) Roles(ctx context.Context, obj *model.User) ([]*model.Role, error) {
	roles, err := r.DB.GetRolesByUserID(ctx, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch roles for user %d: %w", obj.ID, err)
	}

	var result []*model.Role
	for _, role := range roles {
		result = append(result, &model.Role{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		})
	}
	return result, nil
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
