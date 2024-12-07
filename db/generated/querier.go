// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package generated

import (
	"context"
)

type Querier interface {
	AddFileToGroup(ctx context.Context, arg AddFileToGroupParams) error
	AddFolderToGroup(ctx context.Context, arg AddFolderToGroupParams) error
	AddPermissionToRole(ctx context.Context, arg AddPermissionToRoleParams) error
	AddRoleToUser(ctx context.Context, arg AddRoleToUserParams) error
	AddUserToGroup(ctx context.Context, arg AddUserToGroupParams) error
	CheckHealth(ctx context.Context) error
	CreateAuditLog(ctx context.Context, arg CreateAuditLogParams) error
	CreateCronJobLog(ctx context.Context, cronSlug string) (int64, error)
	CreateFile(ctx context.Context, arg CreateFileParams) (File, error)
	CreateFolder(ctx context.Context, arg CreateFolderParams) (Folder, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreatePasswordReset(ctx context.Context, arg CreatePasswordResetParams) (PasswordReset, error)
	CreatePermission(ctx context.Context, arg CreatePermissionParams) (Role, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteFile(ctx context.Context, id int64) error
	DeleteFilesByIDs(ctx context.Context, dollar_1 []int64) error
	DeleteFolder(ctx context.Context, id int64) error
	DeleteFoldersByIDs(ctx context.Context, dollar_1 []int64) error
	DeleteGroup(ctx context.Context, id int64) error
	DeleteGroupsByIDs(ctx context.Context, dollar_1 []int64) error
	DeletePasswordResetbyUserID(ctx context.Context, userID int64) error
	DeletePermission(ctx context.Context, id int64) error
	DeletePermissionsByIDs(ctx context.Context, dollar_1 []int64) error
	DeleteRole(ctx context.Context, id int64) error
	DeleteRolesByIDs(ctx context.Context, dollar_1 []int64) error
	DeleteUser(ctx context.Context, id int64) error
	DeleteUsersByIDs(ctx context.Context, dollar_1 []int64) error
	GetDashboardKPICount(ctx context.Context) (GetDashboardKPICountRow, error)
	GetEmailTemplateByName(ctx context.Context, name string) (EmailTemplate, error)
	GetFile(ctx context.Context, id int64) (File, error)
	// Join the folders table to get folder details
	// Join group_users to get the groups the user belongs to
	// Join group_files to get the files directly associated with the user's groups
	// Join group_folders to get the folders associated with the user's groups via the pivot table
	GetFilesAndFoldersByUser(ctx context.Context, arg GetFilesAndFoldersByUserParams) ([]GetFilesAndFoldersByUserRow, error)
	// Join the folders table to get folder details
	// Join group_users to get the groups the user belongs to
	// Join group_files to get the files directly associated with the user's groups
	// Join group_folders to get the folders associated with the user's groups via the pivot table
	GetFilesAndFoldersByUserBB(ctx context.Context, arg GetFilesAndFoldersByUserBBParams) ([]GetFilesAndFoldersByUserBBRow, error)
	GetFilesByFolder(ctx context.Context, folderID int64) ([]File, error)
	GetFilesByFolderID(ctx context.Context, folderID int64) ([]File, error)
	GetFilesByGroupID(ctx context.Context, groupID int64) ([]File, error)
	GetFilesByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error)
	GetFolder(ctx context.Context, id int64) (Folder, error)
	GetFoldersByGroupID(ctx context.Context, groupID int64) ([]Folder, error)
	GetFoldersByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error)
	GetGroup(ctx context.Context, id int64) (GetGroupRow, error)
	GetGroupsByFileID(ctx context.Context, fileID int64) ([]Group, error)
	GetGroupsByFolderID(ctx context.Context, folderID int64) ([]Group, error)
	GetGroupsByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error)
	GetGroupsByUserID(ctx context.Context, userID int64) ([]Group, error)
	GetPaginatedFilesByFolderID(ctx context.Context, arg GetPaginatedFilesByFolderIDParams) ([]File, error)
	GetPaginatedFilesByFolderIDCount(ctx context.Context, arg GetPaginatedFilesByFolderIDCountParams) (int64, error)
	GetPaginatedFilesByGroupID(ctx context.Context, arg GetPaginatedFilesByGroupIDParams) ([]GetPaginatedFilesByGroupIDRow, error)
	GetPaginatedFilesByGroupIDCount(ctx context.Context, arg GetPaginatedFilesByGroupIDCountParams) (int64, error)
	GetPaginatedFoldersByGroupID(ctx context.Context, arg GetPaginatedFoldersByGroupIDParams) ([]GetPaginatedFoldersByGroupIDRow, error)
	GetPaginatedFoldersByGroupIDCount(ctx context.Context, arg GetPaginatedFoldersByGroupIDCountParams) (int64, error)
	GetPaginatedGroupsByFileID(ctx context.Context, arg GetPaginatedGroupsByFileIDParams) ([]GetPaginatedGroupsByFileIDRow, error)
	GetPaginatedGroupsByFileIDCount(ctx context.Context, arg GetPaginatedGroupsByFileIDCountParams) (int64, error)
	GetPaginatedGroupsByFolderID(ctx context.Context, arg GetPaginatedGroupsByFolderIDParams) ([]GetPaginatedGroupsByFolderIDRow, error)
	GetPaginatedGroupsByFolderIDCount(ctx context.Context, arg GetPaginatedGroupsByFolderIDCountParams) (int64, error)
	GetPaginatedGroupsByUserID(ctx context.Context, arg GetPaginatedGroupsByUserIDParams) ([]GetPaginatedGroupsByUserIDRow, error)
	GetPaginatedGroupsByUserIDCount(ctx context.Context, arg GetPaginatedGroupsByUserIDCountParams) (int64, error)
	GetPaginatedPermissionsByRoleID(ctx context.Context, arg GetPaginatedPermissionsByRoleIDParams) ([]GetPaginatedPermissionsByRoleIDRow, error)
	GetPaginatedPermissionsByRoleIDCount(ctx context.Context, arg GetPaginatedPermissionsByRoleIDCountParams) (int64, error)
	GetPaginatedRolesByPermissionID(ctx context.Context, arg GetPaginatedRolesByPermissionIDParams) ([]GetPaginatedRolesByPermissionIDRow, error)
	GetPaginatedRolesByPermissionIDCount(ctx context.Context, arg GetPaginatedRolesByPermissionIDCountParams) (int64, error)
	GetPaginatedRolesByUserID(ctx context.Context, arg GetPaginatedRolesByUserIDParams) ([]GetPaginatedRolesByUserIDRow, error)
	GetPaginatedRolesByUserIDCount(ctx context.Context, arg GetPaginatedRolesByUserIDCountParams) (int64, error)
	GetPaginatedUsersByGroupID(ctx context.Context, arg GetPaginatedUsersByGroupIDParams) ([]GetPaginatedUsersByGroupIDRow, error)
	GetPaginatedUsersByGroupIDCount(ctx context.Context, arg GetPaginatedUsersByGroupIDCountParams) (int64, error)
	GetPaginatedUsersByRoleID(ctx context.Context, arg GetPaginatedUsersByRoleIDParams) ([]GetPaginatedUsersByRoleIDRow, error)
	GetPaginatedUsersByRoleIDCount(ctx context.Context, arg GetPaginatedUsersByRoleIDCountParams) (int64, error)
	GetPasswordResetbyUserID(ctx context.Context, userID int64) (PasswordReset, error)
	GetPermission(ctx context.Context, id int64) (GetPermissionRow, error)
	GetPermissionsByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error)
	GetPermissionsByRoleID(ctx context.Context, roleID int64) ([]Permission, error)
	GetRole(ctx context.Context, id int64) (GetRoleRow, error)
	GetRolesByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error)
	GetRolesByPermissionID(ctx context.Context, permissionID int64) ([]Role, error)
	GetRolesByUserID(ctx context.Context, userID int64) ([]Role, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserID(ctx context.Context, id int64) (int64, error)
	GetUserPermissions(ctx context.Context, userID int64) ([]string, error)
	GetUsersByGroupID(ctx context.Context, groupID int64) ([]GetUsersByGroupIDRow, error)
	GetUsersByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error)
	GetUsersByRoleID(ctx context.Context, roleID int64) ([]GetUsersByRoleIDRow, error)
	ListCronJobs(ctx context.Context) ([]CronJob, error)
	ListFiles(ctx context.Context) ([]File, error)
	ListFolders(ctx context.Context) ([]Folder, error)
	ListGroups(ctx context.Context) ([]Group, error)
	ListPermissions(ctx context.Context) ([]Permission, error)
	ListRoles(ctx context.Context) ([]Role, error)
	ListUsers(ctx context.Context) ([]User, error)
	PaginatedFiles(ctx context.Context, arg PaginatedFilesParams) ([]File, error)
	PaginatedFilesCount(ctx context.Context, arg PaginatedFilesCountParams) (int64, error)
	PaginatedFolders(ctx context.Context, arg PaginatedFoldersParams) ([]Folder, error)
	PaginatedFoldersCount(ctx context.Context, arg PaginatedFoldersCountParams) (int64, error)
	PaginatedGroups(ctx context.Context, arg PaginatedGroupsParams) ([]Group, error)
	PaginatedGroupsCount(ctx context.Context, arg PaginatedGroupsCountParams) (int64, error)
	PaginatedPermissions(ctx context.Context, arg PaginatedPermissionsParams) ([]Permission, error)
	PaginatedPermissionsCount(ctx context.Context, arg PaginatedPermissionsCountParams) (int64, error)
	PaginatedRoles(ctx context.Context, arg PaginatedRolesParams) ([]Role, error)
	PaginatedRolesCount(ctx context.Context, arg PaginatedRolesCountParams) (int64, error)
	PaginatedUsers(ctx context.Context, arg PaginatedUsersParams) ([]User, error)
	PaginatedUsersCount(ctx context.Context, arg PaginatedUsersCountParams) (int64, error)
	RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error)
	RemoveFileFromGroup(ctx context.Context, arg RemoveFileFromGroupParams) error
	RemoveFolderFromGroup(ctx context.Context, arg RemoveFolderFromGroupParams) error
	RemovePermissionFromRole(ctx context.Context, arg RemovePermissionFromRoleParams) error
	RemoveRoleFromUser(ctx context.Context, arg RemoveRoleFromUserParams) error
	RemoveUserFromGroup(ctx context.Context, arg RemoveUserFromGroupParams) error
	UpdateCronJobLogFailed(ctx context.Context, arg UpdateCronJobLogFailedParams) error
	UpdateCronJobLogSuccess(ctx context.Context, arg UpdateCronJobLogSuccessParams) error
	UpdateFile(ctx context.Context, arg UpdateFileParams) (File, error)
	UpdateFolder(ctx context.Context, arg UpdateFolderParams) (Folder, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (Group, error)
	UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
	UpdateUserEmailVerifiedAt(ctx context.Context, id int64) error
	UpdateUserLastSeenAt(ctx context.Context, id int64) error
	UpdateUserLastSeenAtByEmail(ctx context.Context, email string) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
