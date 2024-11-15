package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/awanishnathpandey/leaf/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB      *pgxpool.Pool
	users   []*model.User
	folders []*model.Folder
}
