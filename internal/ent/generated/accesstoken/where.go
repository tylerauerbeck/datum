// Code generated by ent, DO NOT EDIT.

package accesstoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUpdatedBy, v))
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldAccessToken, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldExpiresAt, v))
}

// IssuedAt applies equality check predicate on the "issued_at" field. It's identical to IssuedAtEQ.
func IssuedAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldIssuedAt, v))
}

// LastUsedAt applies equality check predicate on the "last_used_at" field. It's identical to LastUsedAtEQ.
func LastUsedAt(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldLastUsedAt, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldOrganizationID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUserID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldAccessToken, v))
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldAccessToken, v))
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldAccessToken, vs...))
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldAccessToken, vs...))
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldAccessToken, v))
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldAccessToken, v))
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldAccessToken, v))
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldAccessToken, v))
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldAccessToken, v))
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldAccessToken, v))
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldAccessToken, v))
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldAccessToken, v))
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldAccessToken, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldExpiresAt, v))
}

// IssuedAtEQ applies the EQ predicate on the "issued_at" field.
func IssuedAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldIssuedAt, v))
}

// IssuedAtNEQ applies the NEQ predicate on the "issued_at" field.
func IssuedAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldIssuedAt, v))
}

// IssuedAtIn applies the In predicate on the "issued_at" field.
func IssuedAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldIssuedAt, vs...))
}

// IssuedAtNotIn applies the NotIn predicate on the "issued_at" field.
func IssuedAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldIssuedAt, vs...))
}

// IssuedAtGT applies the GT predicate on the "issued_at" field.
func IssuedAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldIssuedAt, v))
}

// IssuedAtGTE applies the GTE predicate on the "issued_at" field.
func IssuedAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldIssuedAt, v))
}

// IssuedAtLT applies the LT predicate on the "issued_at" field.
func IssuedAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldIssuedAt, v))
}

// IssuedAtLTE applies the LTE predicate on the "issued_at" field.
func IssuedAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldIssuedAt, v))
}

// LastUsedAtEQ applies the EQ predicate on the "last_used_at" field.
func LastUsedAtEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldLastUsedAt, v))
}

// LastUsedAtNEQ applies the NEQ predicate on the "last_used_at" field.
func LastUsedAtNEQ(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldLastUsedAt, v))
}

// LastUsedAtIn applies the In predicate on the "last_used_at" field.
func LastUsedAtIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldLastUsedAt, vs...))
}

// LastUsedAtNotIn applies the NotIn predicate on the "last_used_at" field.
func LastUsedAtNotIn(vs ...time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldLastUsedAt, vs...))
}

// LastUsedAtGT applies the GT predicate on the "last_used_at" field.
func LastUsedAtGT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldLastUsedAt, v))
}

// LastUsedAtGTE applies the GTE predicate on the "last_used_at" field.
func LastUsedAtGTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldLastUsedAt, v))
}

// LastUsedAtLT applies the LT predicate on the "last_used_at" field.
func LastUsedAtLT(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldLastUsedAt, v))
}

// LastUsedAtLTE applies the LTE predicate on the "last_used_at" field.
func LastUsedAtLTE(v time.Time) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldLastUsedAt, v))
}

// LastUsedAtIsNil applies the IsNil predicate on the "last_used_at" field.
func LastUsedAtIsNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIsNull(FieldLastUsedAt))
}

// LastUsedAtNotNil applies the NotNil predicate on the "last_used_at" field.
func LastUsedAtNotNil() predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotNull(FieldLastUsedAt))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// OrganizationIDGT applies the GT predicate on the "organization_id" field.
func OrganizationIDGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldOrganizationID, v))
}

// OrganizationIDGTE applies the GTE predicate on the "organization_id" field.
func OrganizationIDGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldOrganizationID, v))
}

// OrganizationIDLT applies the LT predicate on the "organization_id" field.
func OrganizationIDLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldOrganizationID, v))
}

// OrganizationIDLTE applies the LTE predicate on the "organization_id" field.
func OrganizationIDLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldOrganizationID, v))
}

// OrganizationIDContains applies the Contains predicate on the "organization_id" field.
func OrganizationIDContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldOrganizationID, v))
}

// OrganizationIDHasPrefix applies the HasPrefix predicate on the "organization_id" field.
func OrganizationIDHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldOrganizationID, v))
}

// OrganizationIDHasSuffix applies the HasSuffix predicate on the "organization_id" field.
func OrganizationIDHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldOrganizationID, v))
}

// OrganizationIDEqualFold applies the EqualFold predicate on the "organization_id" field.
func OrganizationIDEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldOrganizationID, v))
}

// OrganizationIDContainsFold applies the ContainsFold predicate on the "organization_id" field.
func OrganizationIDContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldOrganizationID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldUserID, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.AccessToken
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		step := newOwnerStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.AccessToken
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.NotPredicates(p))
}