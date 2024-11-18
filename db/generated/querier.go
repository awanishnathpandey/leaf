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
	AddUserToGroup(ctx context.Context, arg AddUserToGroupParams) error
	CheckHealth(ctx context.Context) error
	CreateFile(ctx context.Context, arg CreateFileParams) (File, error)
	CreateFolder(ctx context.Context, arg CreateFolderParams) (Folder, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteFile(ctx context.Context, id int64) error
	DeleteFolder(ctx context.Context, id int64) error
	DeleteGroup(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetFile(ctx context.Context, id int64) (File, error)
	GetFilesByFolder(ctx context.Context, folderID int64) ([]File, error)
	GetFilesByGroupID(ctx context.Context, groupID int64) ([]File, error)
	GetFolder(ctx context.Context, id int64) (Folder, error)
	GetFoldersByGroupID(ctx context.Context, groupID int64) ([]Folder, error)
	GetGroup(ctx context.Context, id int64) (Group, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUsersByGroupID(ctx context.Context, groupID int64) ([]GetUsersByGroupIDRow, error)
	ListFiles(ctx context.Context) ([]File, error)
	ListFolders(ctx context.Context) ([]Folder, error)
	ListGroups(ctx context.Context) ([]Group, error)
	ListUsers(ctx context.Context) ([]User, error)
	RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error)
	RemoveFileFromGroup(ctx context.Context, arg RemoveFileFromGroupParams) error
	RemoveFolderFromGroup(ctx context.Context, arg RemoveFolderFromGroupParams) error
	RemoveUserFromGroup(ctx context.Context, arg RemoveUserFromGroupParams) error
	UpdateFile(ctx context.Context, arg UpdateFileParams) (File, error)
	UpdateFolder(ctx context.Context, arg UpdateFolderParams) error
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateUserEmailVerifiedAt(ctx context.Context, id int64) error
	UpdateUserLastSeenAt(ctx context.Context, id int64) error
}

var _ Querier = (*Queries)(nil)
