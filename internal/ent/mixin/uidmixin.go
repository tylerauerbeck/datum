package mixin

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/internal/nanox"
	"github.com/ogen-go/ogen"
)

// IDMixin holds the schema definition for the ID
type IDMixin struct {
	mixin.Schema
}

// Fields of the IDMixin
func (i IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(nanox.ID("")).
			Unique().
			Immutable().
			//			DefaultFunc(func() nanox.ID { return nanox.ID(nanox.MustGetNewID()) }).
			Annotations(
				entoas.Schema(ogen.String()),
				entgql.OrderField("ID"),
			),
	}
}

func (IDMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		IDHook(),
	}
}

func IDHook() ent.Hook {
	type IDSetter interface {
		SetID(nanox.ID)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(IDSetter)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}
			id := nanox.MustGetNewID()

			is.SetID(id)
			return next.Mutate(ctx, m)
		})
	}
}
