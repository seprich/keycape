package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"

	"github.com/seprich/keycape/internal/graph/model"
)

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, id string, name string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: AddUser - addUser"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
