package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input generated.CreateUserInput) (*UserCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input generated.UpdateUserInput) (*UserUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*UserDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*generated.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
