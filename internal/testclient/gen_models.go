// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package testclient

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// CreateIntegrationInput is used for create Integration object.
// Input was generated by ent.
type CreateIntegrationInput struct {
	Kind           string     `json:"kind"`
	Description    *string    `json:"description,omitempty"`
	SecretName     string     `json:"secretName"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty"`
	OrganizationID string     `json:"organizationID"`
}

// CreateMembershipInput is used for create Membership object.
// Input was generated by ent.
type CreateMembershipInput struct {
	Current        *bool      `json:"current,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	OrganizationID string     `json:"organizationID"`
	UserID         string     `json:"userID"`
}

// CreateOrganizationInput is used for create Organization object.
// Input was generated by ent.
type CreateOrganizationInput struct {
	Name           *string    `json:"name,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	MembershipIDs  []string   `json:"membershipIDs,omitempty"`
	IntegrationIDs []string   `json:"integrationIDs,omitempty"`
}

// CreateUserInput is used for create User object.
// Input was generated by ent.
type CreateUserInput struct {
	Email         string     `json:"email"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	MembershipIDs []string   `json:"membershipIDs,omitempty"`
}

type Integration struct {
	ID           string       `json:"id"`
	Kind         string       `json:"kind"`
	Description  *string      `json:"description,omitempty"`
	SecretName   string       `json:"secretName"`
	CreatedAt    time.Time    `json:"createdAt"`
	DeletedAt    *time.Time   `json:"deletedAt,omitempty"`
	Organization Organization `json:"organization"`
}

func (Integration) IsNode() {}

// Return response for createIntegration mutation
type IntegrationCreatePayload struct {
	// Created integration
	Integration Integration `json:"integration"`
}

// Return response for deleteIntegration mutation
type IntegrationDeletePayload struct {
	// Deleted integration ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateIntegration mutation
type IntegrationUpdatePayload struct {
	// Updated integration
	Integration Integration `json:"integration"`
}

// IntegrationWhereInput is used for filtering Integration objects.
// Input was generated by ent.
type IntegrationWhereInput struct {
	Not *IntegrationWhereInput   `json:"not,omitempty"`
	And []*IntegrationWhereInput `json:"and,omitempty"`
	Or  []*IntegrationWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// kind field predicates
	Kind             *string  `json:"kind,omitempty"`
	KindNeq          *string  `json:"kindNEQ,omitempty"`
	KindIn           []string `json:"kindIn,omitempty"`
	KindNotIn        []string `json:"kindNotIn,omitempty"`
	KindGt           *string  `json:"kindGT,omitempty"`
	KindGte          *string  `json:"kindGTE,omitempty"`
	KindLt           *string  `json:"kindLT,omitempty"`
	KindLte          *string  `json:"kindLTE,omitempty"`
	KindContains     *string  `json:"kindContains,omitempty"`
	KindHasPrefix    *string  `json:"kindHasPrefix,omitempty"`
	KindHasSuffix    *string  `json:"kindHasSuffix,omitempty"`
	KindEqualFold    *string  `json:"kindEqualFold,omitempty"`
	KindContainsFold *string  `json:"kindContainsFold,omitempty"`
	// description field predicates
	Description             *string  `json:"description,omitempty"`
	DescriptionNeq          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGt           *string  `json:"descriptionGT,omitempty"`
	DescriptionGte          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLt           *string  `json:"descriptionLT,omitempty"`
	DescriptionLte          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionIsNil        *bool    `json:"descriptionIsNil,omitempty"`
	DescriptionNotNil       *bool    `json:"descriptionNotNil,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`
	// secret_name field predicates
	SecretName             *string  `json:"secretName,omitempty"`
	SecretNameNeq          *string  `json:"secretNameNEQ,omitempty"`
	SecretNameIn           []string `json:"secretNameIn,omitempty"`
	SecretNameNotIn        []string `json:"secretNameNotIn,omitempty"`
	SecretNameGt           *string  `json:"secretNameGT,omitempty"`
	SecretNameGte          *string  `json:"secretNameGTE,omitempty"`
	SecretNameLt           *string  `json:"secretNameLT,omitempty"`
	SecretNameLte          *string  `json:"secretNameLTE,omitempty"`
	SecretNameContains     *string  `json:"secretNameContains,omitempty"`
	SecretNameHasPrefix    *string  `json:"secretNameHasPrefix,omitempty"`
	SecretNameHasSuffix    *string  `json:"secretNameHasSuffix,omitempty"`
	SecretNameEqualFold    *string  `json:"secretNameEqualFold,omitempty"`
	SecretNameContainsFold *string  `json:"secretNameContainsFold,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// deleted_at field predicates
	DeletedAt       *time.Time   `json:"deletedAt,omitempty"`
	DeletedAtNeq    *time.Time   `json:"deletedAtNEQ,omitempty"`
	DeletedAtIn     []*time.Time `json:"deletedAtIn,omitempty"`
	DeletedAtNotIn  []*time.Time `json:"deletedAtNotIn,omitempty"`
	DeletedAtGt     *time.Time   `json:"deletedAtGT,omitempty"`
	DeletedAtGte    *time.Time   `json:"deletedAtGTE,omitempty"`
	DeletedAtLt     *time.Time   `json:"deletedAtLT,omitempty"`
	DeletedAtLte    *time.Time   `json:"deletedAtLTE,omitempty"`
	DeletedAtIsNil  *bool        `json:"deletedAtIsNil,omitempty"`
	DeletedAtNotNil *bool        `json:"deletedAtNotNil,omitempty"`
	// organization edge predicates
	HasOrganization     *bool                     `json:"hasOrganization,omitempty"`
	HasOrganizationWith []*OrganizationWhereInput `json:"hasOrganizationWith,omitempty"`
}

type Membership struct {
	ID           string       `json:"id"`
	Current      bool         `json:"current"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	Organization Organization `json:"organization"`
	User         User         `json:"user"`
}

func (Membership) IsNode() {}

// Return response for createMembership mutation
type MembershipCreatePayload struct {
	// Created membership
	Membership Membership `json:"membership"`
}

// Return response for deleteMembership mutation
type MembershipDeletePayload struct {
	// Deleted membership ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateMembership mutation
type MembershipUpdatePayload struct {
	// Updated membership
	Membership Membership `json:"membership"`
}

// MembershipWhereInput is used for filtering Membership objects.
// Input was generated by ent.
type MembershipWhereInput struct {
	Not *MembershipWhereInput   `json:"not,omitempty"`
	And []*MembershipWhereInput `json:"and,omitempty"`
	Or  []*MembershipWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// current field predicates
	Current    *bool `json:"current,omitempty"`
	CurrentNeq *bool `json:"currentNEQ,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// organization edge predicates
	HasOrganization     *bool                     `json:"hasOrganization,omitempty"`
	HasOrganizationWith []*OrganizationWhereInput `json:"hasOrganizationWith,omitempty"`
	// user edge predicates
	HasUser     *bool             `json:"hasUser,omitempty"`
	HasUserWith []*UserWhereInput `json:"hasUserWith,omitempty"`
}

type Organization struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	CreatedAt    time.Time      `json:"createdAt"`
	Memberships  []*Membership  `json:"memberships,omitempty"`
	Integrations []*Integration `json:"integrations,omitempty"`
}

func (Organization) IsNode() {}

// Return response for createOrganization mutation
type OrganizationCreatePayload struct {
	// Created organization
	Organization Organization `json:"organization"`
}

// Return response for deleteOrganization mutation
type OrganizationDeletePayload struct {
	// Deleted organization ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOrganization mutation
type OrganizationUpdatePayload struct {
	// Updated organization
	Organization Organization `json:"organization"`
}

// OrganizationWhereInput is used for filtering Organization objects.
// Input was generated by ent.
type OrganizationWhereInput struct {
	Not *OrganizationWhereInput   `json:"not,omitempty"`
	And []*OrganizationWhereInput `json:"and,omitempty"`
	Or  []*OrganizationWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// memberships edge predicates
	HasMemberships     *bool                   `json:"hasMemberships,omitempty"`
	HasMembershipsWith []*MembershipWhereInput `json:"hasMembershipsWith,omitempty"`
	// integrations edge predicates
	HasIntegrations     *bool                    `json:"hasIntegrations,omitempty"`
	HasIntegrationsWith []*IntegrationWhereInput `json:"hasIntegrationsWith,omitempty"`
}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor,omitempty"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor,omitempty"`
}

// UpdateIntegrationInput is used for update Integration object.
// Input was generated by ent.
type UpdateIntegrationInput struct {
	Description      *string    `json:"description,omitempty"`
	ClearDescription *bool      `json:"clearDescription,omitempty"`
	DeletedAt        *time.Time `json:"deletedAt,omitempty"`
	ClearDeletedAt   *bool      `json:"clearDeletedAt,omitempty"`
	OrganizationID   *string    `json:"organizationID,omitempty"`
}

// UpdateMembershipInput is used for update Membership object.
// Input was generated by ent.
type UpdateMembershipInput struct {
	Current        *bool      `json:"current,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	OrganizationID *string    `json:"organizationID,omitempty"`
	UserID         *string    `json:"userID,omitempty"`
}

// UpdateOrganizationInput is used for update Organization object.
// Input was generated by ent.
type UpdateOrganizationInput struct {
	Name                 *string  `json:"name,omitempty"`
	AddMembershipIDs     []string `json:"addMembershipIDs,omitempty"`
	RemoveMembershipIDs  []string `json:"removeMembershipIDs,omitempty"`
	ClearMemberships     *bool    `json:"clearMemberships,omitempty"`
	AddIntegrationIDs    []string `json:"addIntegrationIDs,omitempty"`
	RemoveIntegrationIDs []string `json:"removeIntegrationIDs,omitempty"`
	ClearIntegrations    *bool    `json:"clearIntegrations,omitempty"`
}

// UpdateUserInput is used for update User object.
// Input was generated by ent.
type UpdateUserInput struct {
	Email               *string  `json:"email,omitempty"`
	AddMembershipIDs    []string `json:"addMembershipIDs,omitempty"`
	RemoveMembershipIDs []string `json:"removeMembershipIDs,omitempty"`
	ClearMemberships    *bool    `json:"clearMemberships,omitempty"`
}

type User struct {
	ID          string        `json:"id"`
	Email       string        `json:"email"`
	CreatedAt   time.Time     `json:"createdAt"`
	Memberships []*Membership `json:"memberships,omitempty"`
}

func (User) IsNode() {}

// Return response for createUser mutation
type UserCreatePayload struct {
	// Created user
	User User `json:"user"`
}

// Return response for deleteUser mutation
type UserDeletePayload struct {
	// Deleted user ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateUser mutation
type UserUpdatePayload struct {
	// Updated user
	User User `json:"user"`
}

// UserWhereInput is used for filtering User objects.
// Input was generated by ent.
type UserWhereInput struct {
	Not *UserWhereInput   `json:"not,omitempty"`
	And []*UserWhereInput `json:"and,omitempty"`
	Or  []*UserWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// email field predicates
	Email             *string  `json:"email,omitempty"`
	EmailNeq          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGt           *string  `json:"emailGT,omitempty"`
	EmailGte          *string  `json:"emailGTE,omitempty"`
	EmailLt           *string  `json:"emailLT,omitempty"`
	EmailLte          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// memberships edge predicates
	HasMemberships     *bool                   `json:"hasMemberships,omitempty"`
	HasMembershipsWith []*MembershipWhereInput `json:"hasMembershipsWith,omitempty"`
}

type Service struct {
	Sdl *string `json:"sdl,omitempty"`
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
