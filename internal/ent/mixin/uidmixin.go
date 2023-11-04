package mixin

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/ogen-go/ogen"

	"github.com/datumforge/datum/internal/nanox"
)

var _ ent.Mixin = (*IDMixin)(nil)

// IDMixin holds the schema definition for the ID
type IDMixin struct {
	mixin.Schema
}

// Fields of the Doc.
func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(nanox.ID("")).
			Unique().
			Immutable().
			DefaultFunc(func() nanox.ID { return nanox.ID(nanox.MustGetNewID()) }).
			Annotations(entoas.Schema(ogen.String())),
	}
}
