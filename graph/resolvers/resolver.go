package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/external/mail"
	"github.com/awanishnathpandey/leaf/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB          *generated.Queries
	Pool        *pgxpool.Pool
	users       []*model.User
	folders     []*model.Folder
	files       []*model.File
	groups      []*model.Group
	roles       []*model.Role
	permissions []*model.Permission
	MailService *mail.MailService
}
