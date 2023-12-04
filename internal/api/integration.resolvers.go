package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
)

// CreateIntegration is the resolver for the createIntegration field.
func (r *mutationResolver) CreateIntegration(ctx context.Context, input generated.CreateIntegrationInput) (*IntegrationCreatePayload, error) {
	// TODO - add permissions checks
	i, err := r.client.Integration.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to create integration", "error", err)
		return nil, ErrInternalServerError
	}

	return &IntegrationCreatePayload{Integration: i}, nil
}

// UpdateIntegration is the resolver for the updateIntegration field.
func (r *mutationResolver) UpdateIntegration(ctx context.Context, id string, input generated.UpdateIntegrationInput) (*IntegrationUpdatePayload, error) {
	// TODO - add permissions checks

	i, err := r.client.Integration.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get integration", "error", err)
		return nil, ErrInternalServerError
	}

	i, err = i.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to update integration", "error", err)
		return nil, ErrInternalServerError
	}

	return &IntegrationUpdatePayload{Integration: i}, nil
}

// DeleteIntegration is the resolver for the deleteIntegration field.
func (r *mutationResolver) DeleteIntegration(ctx context.Context, id string) (*IntegrationDeletePayload, error) {
	// TODO - add permissions checks

	if err := r.client.Integration.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to delete integration", "error", err)
		return nil, err
	}

	return &IntegrationDeletePayload{DeletedID: id}, nil
}

// Integration is the resolver for the integration field.
func (r *queryResolver) Integration(ctx context.Context, id string) (*generated.Integration, error) {
	// TODO - add permissions checks

	i, err := r.client.Integration.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get integration", "error", err)
		return nil, ErrInternalServerError
	}

	return i, nil
}
