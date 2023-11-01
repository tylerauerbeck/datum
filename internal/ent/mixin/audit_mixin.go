package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/datumforge/datum/internal/echox"
	"github.com/google/uuid"
	"github.com/labstack/echo"
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
		field.UUID("created_by", uuid.UUID{}).
			Optional(),
		field.UUID("updated_by", uuid.UUID{}).
			Optional(),
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
		SetCreatedBy(uuid.UUID)
		CreatedBy() (id uuid.UUID, exists bool)
		SetUpdatedBy(uuid.UUID)
		UpdatedBy() (id uuid.UUID, exists bool)
	}

	return ent.MutateFunc(func(c echo.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, newUnexpectedAuditError(m)
		}

		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())
			if _, exists := ml.CreatedBy(); !exists {
				uid, err := getUserID(c)
				if err != nil {
					return nil, err
				}
				ml.SetCreatedBy(uid)
			}

		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())
			if _, exists := ml.UpdatedBy(); !exists {
				uid, err := getUserID(c)
				if err != nil {
					return nil, err
				}
				ml.SetUpdatedBy(uid)
			}
		}
		return next.Mutate(c, m)
	})
}

// Actor retrieves the ActorKey from echo Context.
func getUserID(c echo.Context) (string, err error) {
	if actor, ok := c.Get(echox.ActorKey).(string); ok {
		return actor
	}

	if err != nil {
		return nil, err
	}

	return
}
