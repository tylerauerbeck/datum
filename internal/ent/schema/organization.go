package schema

import (
	"strings"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
)

const (
	orgNameMaxLen = 160
)

// Organization holds the schema definition for the Organization entity - organizations are the top level tenancy construct in the system
type Organization struct {
	ent.Schema
}

// Fields of the Organization
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().
			MaxLen(orgNameMaxLen).
			NotEmpty().
			Annotations(
				entgql.OrderField("name"),
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("display_name").
			Comment("The organization's displayed 'friendly' name").
			MaxLen(nameMaxLen).
			NotEmpty().
			Default("unknown").
			Annotations(
				entgql.OrderField("display_name"),
			).
			Validate(
				func(s string) error {
					if strings.Contains(s, " ") {
						return ErrContainsSpaces
					}
					return nil
				},
			),
		field.String("description").
			Comment("An optional description of the Organization").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("parent_organization_id").Optional().Immutable().
			Comment("The ID of the parent organization for the organization.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipType),
				entoas.Schema(ogen.String()),
			),
	}
}

// Edges of the Organization
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Organization.Type).
			Annotations(
				entgql.RelayConnection(),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			).
			From("parent").
			Field("parent_organization_id").
			Immutable().
			Unique(),
		// an org can have and belong to many users
		edge.From("users", User.Type).Ref("organizations"),
		edge.To("groups", Group.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("integrations", Integration.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("setting", OrganizationSetting.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("entitlements", Entitlement.Type),
		edge.To("oauthprovider", OauthProvider.Type),
	}
}

// Annotations of the Organization
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Organization
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Policy defines the privacy policy of the Organization.
func (Organization) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoSubject(),      // Requires a user to be authenticated with a valid JWT
			rule.HasOrgMutationAccess(), // Requires edit for Update, and delete for Delete mutations
			privacy.AlwaysAllowRule(),   // Allow all other users (e.g. a user with a JWT should be able to create a new org)
		},
		Query: privacy.QueryPolicy{
			rule.DenyIfNoSubject(),   // Requires a user to be authenticated with a valid JWT
			rule.HasOrgReadAccess(),  // Requires a user to have can_view access of the org
			privacy.AlwaysDenyRule(), // Deny all other users
		},
	}
}

// Hooks of the Organization
func (Organization) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookOrganization(),
	}
}
