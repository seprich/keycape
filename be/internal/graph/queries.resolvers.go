package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"
	"fmt"

	"github.com/seprich/keycape/internal/graph/model"
)

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	fake := model.User{ID: "1234", Name: "Blake Dance"}
	//panic(fmt.Errorf("not implemented: GetUser - getUser"))
	return &fake, nil
}

// GetTodo is the resolver for the getTodo field.
func (r *queryResolver) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: GetTodo - getTodo"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
