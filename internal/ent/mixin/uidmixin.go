package mixin

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/internal/idx"
)

var _ ent.Mixin = (*IDMixin)(nil)

// Doc holds the schema definition for the Doc entity.
type IDMixin struct {
	mixin.Schema
}

// Fields of the Doc.
func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(idx.ID("")).
			Unique().
			Immutable().
			DefaultFunc(idx.MustGetNewID),
	}
}

func (IDMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		HookID(),
	}
}

func HookID() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			if mutation.Op().Is(ent.OpCreate) {
				s := mutation.(interface {
					SetID(v idx.ID)
				})
				id, err := NextID(ctx, mutation)
				if err != nil {
					return nil, err
				}
				s.SetID(id)
			}
			return next.Mutate(ctx, mutation)
		})
	}
}

var NextID = func(ctx context.Context, mutation ent.Mutation) (idx.ID, error) {
	// test only
	return idx.ID(idx.MustGetNewID()), nil
}
