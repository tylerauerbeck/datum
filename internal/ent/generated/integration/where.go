// Code generated by ent, DO NOT EDIT.

package integration

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/nanox"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id nanox.ID) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldUpdatedBy, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldName, v))
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

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Integration {
	return predicate.Integration(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Integration {
	return predicate.Integration(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Integration {
	return predicate.Integration(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Integration {
	return predicate.Integration(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Integration {
	return predicate.Integration(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Integration {
	return predicate.Integration(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Integration {
	return predicate.Integration(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Integration {
	return predicate.Integration(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Integration {
	return predicate.Integration(sql.FieldContainsFold(FieldName, v))
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

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Integration {
	return predicate.Integration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Integration
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Organization) predicate.Integration {
	return predicate.Integration(func(s *sql.Selector) {
		step := newOwnerStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Integration
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
