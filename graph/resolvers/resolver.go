package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB      *generated.Queries
	users   []*model.User
	folders []*model.Folder
}
