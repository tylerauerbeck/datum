package mixin

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// AuditMixin provides auditing for all records where enabled. The created_at, created_by, updated_at, and updated_by records are automatically populated when this mixin is enabled.
type AuditMixin struct {
	mixin.Schema
}

// Fields of the AuditMixin
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.UUID("created_by", GoType(idx.ID("")).Optional().Immutable().
		field.UUID("updated_by", GoType(idx.ID("")).Optional().Immutable().,
	}
}

// Hooks of the AuditMixin
func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}

// AuditHook sets and returns the created_at, updated_at, etc., fields
func AuditHook(next ent.Mutator) ent.Mutator {
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (v time.Time, exists bool) // exists if present before this hook
		SetUpdatedAt(time.Time)
		UpdatedAt() (v time.Time, exists bool)
		SetCreatedBy(idx.ID)
		CreatedBy() (id idx.ID, exists bool)
		SetUpdatedBy(idx.ID)
		UpdatedBy() (id idx.ID, exists bool)
	}

	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, newUnexpectedAuditError(m)
		}

		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())
		}

		return next.Mutate(ctx, m)
	})
}
