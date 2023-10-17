// Code generated by ent, DO NOT EDIT.

package integration

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldID, id))
}

// Kind applies equality check predicate on the "kind" field. It's identical to KindEQ.
func Kind(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldKind, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldDescription, v))
}

// SecretName applies equality check predicate on the "secret_name" field. It's identical to SecretNameEQ.
func SecretName(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldSecretName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldCreatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldDeletedAt, v))
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldKind, v))
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldKind, v))
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldKind, vs...))
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldKind, vs...))
}

// KindGT applies the GT predicate on the "kind" field.
func KindGT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldKind, v))
}

// KindGTE applies the GTE predicate on the "kind" field.
func KindGTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldKind, v))
}

// KindLT applies the LT predicate on the "kind" field.
func KindLT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldKind, v))
}

// KindLTE applies the LTE predicate on the "kind" field.
func KindLTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldKind, v))
}

// KindContains applies the Contains predicate on the "kind" field.
func KindContains(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContains(FieldKind, v))
}

// KindHasPrefix applies the HasPrefix predicate on the "kind" field.
func KindHasPrefix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasPrefix(FieldKind, v))
}

// KindHasSuffix applies the HasSuffix predicate on the "kind" field.
func KindHasSuffix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasSuffix(FieldKind, v))
}

// KindEqualFold applies the EqualFold predicate on the "kind" field.
func KindEqualFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEqualFold(FieldKind, v))
}

// KindContainsFold applies the ContainsFold predicate on the "kind" field.
func KindContainsFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContainsFold(FieldKind, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Integration {
	return predicate.Integration(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Integration {
	return predicate.Integration(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContainsFold(FieldDescription, v))
}

// SecretNameEQ applies the EQ predicate on the "secret_name" field.
func SecretNameEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldSecretName, v))
}

// SecretNameNEQ applies the NEQ predicate on the "secret_name" field.
func SecretNameNEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldSecretName, v))
}

// SecretNameIn applies the In predicate on the "secret_name" field.
func SecretNameIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldSecretName, vs...))
}

// SecretNameNotIn applies the NotIn predicate on the "secret_name" field.
func SecretNameNotIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldSecretName, vs...))
}

// SecretNameGT applies the GT predicate on the "secret_name" field.
func SecretNameGT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldSecretName, v))
}

// SecretNameGTE applies the GTE predicate on the "secret_name" field.
func SecretNameGTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldSecretName, v))
}

// SecretNameLT applies the LT predicate on the "secret_name" field.
func SecretNameLT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldSecretName, v))
}

// SecretNameLTE applies the LTE predicate on the "secret_name" field.
func SecretNameLTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldSecretName, v))
}

// SecretNameContains applies the Contains predicate on the "secret_name" field.
func SecretNameContains(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContains(FieldSecretName, v))
}

// SecretNameHasPrefix applies the HasPrefix predicate on the "secret_name" field.
func SecretNameHasPrefix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasPrefix(FieldSecretName, v))
}

// SecretNameHasSuffix applies the HasSuffix predicate on the "secret_name" field.
func SecretNameHasSuffix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasSuffix(FieldSecretName, v))
}

// SecretNameEqualFold applies the EqualFold predicate on the "secret_name" field.
func SecretNameEqualFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEqualFold(FieldSecretName, v))
}

// SecretNameContainsFold applies the ContainsFold predicate on the "secret_name" field.
func SecretNameContainsFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContainsFold(FieldSecretName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldCreatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Integration {
	return predicate.Integration(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Integration {
	return predicate.Integration(sql.FieldNotNull(FieldDeletedAt))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.Integration {
	return predicate.Integration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.Integration {
	return predicate.Integration(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Integration) predicate.Integration {
	return predicate.Integration(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Integration) predicate.Integration {
	return predicate.Integration(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Integration) predicate.Integration {
	return predicate.Integration(sql.NotPredicates(p))
}
