package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateIntegration is the resolver for the createIntegration field.
func (r *mutationResolver) CreateIntegration(ctx context.Context, input generated.CreateIntegrationInput) (*IntegrationCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateIntegration - createIntegration"))
}

// UpdateIntegration is the resolver for the updateIntegration field.
func (r *mutationResolver) UpdateIntegration(ctx context.Context, id string, input generated.UpdateIntegrationInput) (*IntegrationUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateIntegration - updateIntegration"))
}

// DeleteIntegration is the resolver for the deleteIntegration field.
func (r *mutationResolver) DeleteIntegration(ctx context.Context, id string) (*IntegrationDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteIntegration - deleteIntegration"))
}

// Integration is the resolver for the integration field.
func (r *queryResolver) Integration(ctx context.Context, id string) (*generated.Integration, error) {
	panic(fmt.Errorf("not implemented: Integration - integration"))
}
