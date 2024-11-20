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
)

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input model.CreateRole) (*model.Role, error) {
	// Validate input
	if err := input.Validate(); err != nil {
		// Call the reusable validation error formatter
		return nil, utils.FormatValidationErrors(err)
	}

	// Call the generated CreateRole function with the params
	role, err := r.DB.CreateRole(ctx, generated.CreateRoleParams{
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, err
	}

	// Return the newly created role
	return &model.Role{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, input model.UpdateRole) (*model.Role, error) {
	// Check if the role exists
	_, err := r.DB.GetRole(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("role not found: %w", err)
	}

	// Update the group
	err = r.DB.UpdateRole(ctx, generated.UpdateRoleParams{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	// Fetch the updated role
	updatedRole, err := r.DB.GetRole(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated role: %w", err)
	}

	// Map the SQLC model to the GraphQL model
	return &model.Role{
		ID:          updatedRole.ID,
		Name:        updatedRole.Name,
		Description: updatedRole.Description,
		CreatedAt:   updatedRole.CreatedAt,
		UpdatedAt:   updatedRole.UpdatedAt,
	}, nil
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, id int64) (bool, error) {
	// Check if the role exists (optional)
	_, err := r.DB.GetRole(ctx, id)
	if err != nil {
		return false, fmt.Errorf("role not found: %w", err)
	}

	// Attempt to delete the role
	err = r.DB.DeleteRole(ctx, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete role: %w", err)
	}
	return true, nil
}

// CreatePermission is the resolver for the createPermission field.
func (r *mutationResolver) CreatePermission(ctx context.Context, input model.CreatePermission) (*model.Permission, error) {
	// Validate input
	if err := input.Validate(); err != nil {
		// Call the reusable validation error formatter
		return nil, utils.FormatValidationErrors(err)
	}

	// Call the generated CreatePermission function with the params
	permission, err := r.DB.CreatePermission(ctx, generated.CreatePermissionParams{
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, err
	}

	// Return the newly created permission
	return &model.Permission{
		ID:          permission.ID,
		Name:        permission.Name,
		Description: permission.Description,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}, nil
}

// UpdatePermission is the resolver for the updatePermission field.
func (r *mutationResolver) UpdatePermission(ctx context.Context, input model.UpdatePermission) (*model.Permission, error) {
	// Check if the permission exists
	_, err := r.DB.GetPermission(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("permission not found: %w", err)
	}

	// Update the permission
	err = r.DB.UpdatePermission(ctx, generated.UpdatePermissionParams{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update permission: %w", err)
	}

	// Fetch the updated permission
	updatedPermission, err := r.DB.GetPermission(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated permission: %w", err)
	}

	// Map the SQLC model to the GraphQL model
	return &model.Permission{
		ID:          updatedPermission.ID,
		Name:        updatedPermission.Name,
		Description: updatedPermission.Description,
		CreatedAt:   updatedPermission.CreatedAt,
		UpdatedAt:   updatedPermission.UpdatedAt,
	}, nil
}

// DeletePermission is the resolver for the deletePermission field.
func (r *mutationResolver) DeletePermission(ctx context.Context, id int64) (bool, error) {
	// Check if the permission exists (optional)
	_, err := r.DB.GetPermission(ctx, id)
	if err != nil {
		return false, fmt.Errorf("permission not found: %w", err)
	}

	// Attempt to delete the permission
	err = r.DB.DeletePermission(ctx, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete permission: %w", err)
	}
	return true, nil
}

// AddRoleToUser is the resolver for the addRoleToUser field.
func (r *mutationResolver) AddRoleToUser(ctx context.Context, roleID int64, userID int64) (bool, error) {
	// Check if the role exists
	_, roleErr := r.DB.GetRole(ctx, roleID)
	if roleErr != nil {
		return false, fmt.Errorf("role not found: %w", roleErr)
	}

	// Check if the user exists
	_, err := r.DB.GetUser(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("user not found: %w", err)
	}

	err = r.DB.AddRoleToUser(ctx, generated.AddRoleToUserParams{
		RoleID: roleID,
		UserID: userID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to add role to user: %w", err)
	}
	return true, nil
}

// RemoveRoleFromUser is the resolver for the removeRoleFromUser field.
func (r *mutationResolver) RemoveRoleFromUser(ctx context.Context, roleID int64, userID int64) (bool, error) {
	// Check if the role exists
	_, roleErr := r.DB.GetRole(ctx, roleID)
	if roleErr != nil {
		return false, fmt.Errorf("role not found: %w", roleErr)
	}

	// Check if the user exists
	_, err := r.DB.GetUser(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("user not found: %w", err)
	}

	err = r.DB.RemoveRoleFromUser(ctx, generated.RemoveRoleFromUserParams{
		RoleID: roleID,
		UserID: userID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to remove role from user: %w", err)
	}
	return true, nil
}

// AddPermissionToRole is the resolver for the addPermissionToRole field.
func (r *mutationResolver) AddPermissionToRole(ctx context.Context, roleID int64, permissionID int64) (bool, error) {
	// Check if the role exists
	_, roleErr := r.DB.GetRole(ctx, roleID)
	if roleErr != nil {
		return false, fmt.Errorf("role not found: %w", roleErr)
	}

	// Check if the permission exists
	_, err := r.DB.GetPermission(ctx, permissionID)
	if err != nil {
		return false, fmt.Errorf("permission not found: %w", err)
	}

	err = r.DB.AddPermissionToRole(ctx, generated.AddPermissionToRoleParams{
		RoleID:       roleID,
		PermissionID: permissionID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to add permission to role: %w", err)
	}
	return true, nil
}

// RemovePermissionFromRole is the resolver for the removePermissionFromRole field.
func (r *mutationResolver) RemovePermissionFromRole(ctx context.Context, roleID int64, permissionID int64) (bool, error) {
	// Check if the role exists
	_, roleErr := r.DB.GetRole(ctx, roleID)
	if roleErr != nil {
		return false, fmt.Errorf("role not found: %w", roleErr)
	}

	// Check if the permission exists
	_, err := r.DB.GetPermission(ctx, permissionID)
	if err != nil {
		return false, fmt.Errorf("permission not found: %w", err)
	}

	err = r.DB.RemovePermissionFromRole(ctx, generated.RemovePermissionFromRoleParams{
		RoleID:       roleID,
		PermissionID: permissionID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to remove permission from role: %w", err)
	}
	return true, nil
}

// Roles is the resolver for the roles field.
func (r *permissionResolver) Roles(ctx context.Context, obj *model.Permission) ([]*model.Role, error) {
	roles, err := r.DB.GetRolesByPermissionID(ctx, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch roles for permission %d: %w", obj.ID, err)
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

// Roles is the resolver for the roles field.
func (r *queryResolver) Roles(ctx context.Context) ([]*model.Role, error) {
	// Fetch roles using sqlc
	rows, err := r.DB.ListRoles(ctx) // Assuming ListRoles is the sqlc query method
	if err != nil {
		return nil, err
	}

	// Map sqlc rows to GraphQL models
	var roles []*model.Role
	for _, row := range rows {
		roles = append(roles, &model.Role{
			ID:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			CreatedAt:   row.CreatedAt, // Or use row.CreatedAt.Time.String()
			UpdatedAt:   row.UpdatedAt, // Or use row.UpdatedAt.Time.String()
		})
	}

	return roles, nil
}

// Permissions is the resolver for the permissions field.
func (r *queryResolver) Permissions(ctx context.Context) ([]*model.Permission, error) {
	// Fetch permissions using sqlc
	rows, err := r.DB.ListPermissions(ctx) // Assuming ListPermissions is the sqlc query method
	if err != nil {
		return nil, err
	}

	// Map sqlc rows to GraphQL models
	var permissions []*model.Permission
	for _, row := range rows {
		permissions = append(permissions, &model.Permission{
			ID:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			CreatedAt:   row.CreatedAt, // Or use row.CreatedAt.Time.String()
			UpdatedAt:   row.UpdatedAt, // Or use row.UpdatedAt.Time.String()
		})
	}

	return permissions, nil
}

// GetRole is the resolver for the getRole field.
func (r *queryResolver) GetRole(ctx context.Context, id int64) (*model.Role, error) {
	// Call the generated GetRole query
	role, err := r.DB.GetRole(ctx, id) // assuming input.ID is of type string
	if err != nil {
		return nil, fmt.Errorf("failed to get role: %w", err)
	}

	// Convert the SQL result to GraphQL model
	return &model.Role{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt, // assuming you're using timestamptz
		UpdatedAt:   role.UpdatedAt, // assuming you're using timestamptz
	}, nil
}

// GetPermission is the resolver for the getPermission field.
func (r *queryResolver) GetPermission(ctx context.Context, id int64) (*model.Permission, error) {
	// Call the generated GetPermission query
	permission, err := r.DB.GetPermission(ctx, id) // assuming input.ID is of type string
	if err != nil {
		return nil, fmt.Errorf("failed to get permission: %w", err)
	}

	// Convert the SQL result to GraphQL model
	return &model.Permission{
		ID:          permission.ID,
		Name:        permission.Name,
		Description: permission.Description,
		CreatedAt:   permission.CreatedAt, // assuming you're using timestamptz
		UpdatedAt:   permission.UpdatedAt, // assuming you're using timestamptz
	}, nil
}

// Permissions is the resolver for the permissions field.
func (r *roleResolver) Permissions(ctx context.Context, obj *model.Role) ([]*model.Permission, error) {
	permissions, err := r.DB.GetPermissionsByRoleID(ctx, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch permissions for role %d: %w", obj.ID, err)
	}

	var result []*model.Permission
	for _, role := range permissions {
		result = append(result, &model.Permission{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		})
	}
	return result, nil
}

// Users is the resolver for the users field.
func (r *roleResolver) Users(ctx context.Context, obj *model.Role) ([]*model.User, error) {
	users, err := r.DB.GetUsersByRoleID(ctx, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users for role %d: %w", obj.ID, err)
	}

	var result []*model.User
	for _, user := range users {
		result = append(result, &model.User{
			ID:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			EmailVerifiedAt: (*int64)(&user.EmailVerifiedAt.Int64),
			LastSeenAt:      user.LastSeenAt,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
			DeletedAt:       (*int64)(&user.DeletedAt.Int64),
		})
	}
	return result, nil
}

// Permission returns graph.PermissionResolver implementation.
func (r *Resolver) Permission() graph.PermissionResolver { return &permissionResolver{r} }

// Role returns graph.RoleResolver implementation.
func (r *Resolver) Role() graph.RoleResolver { return &roleResolver{r} }

type permissionResolver struct{ *Resolver }
type roleResolver struct{ *Resolver }