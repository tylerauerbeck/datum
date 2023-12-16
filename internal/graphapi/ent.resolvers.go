package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

// AuthStyle is the resolver for the authStyle field.
func (r *oauthProviderResolver) AuthStyle(ctx context.Context, obj *generated.OauthProvider) (int, error) {
	panic(fmt.Errorf("not implemented: AuthStyle - authStyle"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (generated.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]generated.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// AccessTokens is the resolver for the accessTokens field.
func (r *queryResolver) AccessTokens(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.AccessTokenWhereInput) (*generated.AccessTokenConnection, error) {
	panic(fmt.Errorf("not implemented: AccessTokens - accessTokens"))
}

// Entitlements is the resolver for the entitlements field.
func (r *queryResolver) Entitlements(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.EntitlementWhereInput) (*generated.EntitlementConnection, error) {
	panic(fmt.Errorf("not implemented: Entitlements - entitlements"))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *generated.GroupOrder, where *generated.GroupWhereInput) (*generated.GroupConnection, error) {
	// if auth is disabled, policy decisions will be skipped
	if r.authDisabled {
		ctx = privacy.DecisionContext(ctx, privacy.Allow)
	}

	return r.client.Group.Query().Paginate(ctx, after, first, before, last, generated.WithGroupOrder(orderBy), generated.WithGroupFilter(where.Filter))
}

// GroupSettings is the resolver for the groupSettings field.
func (r *queryResolver) GroupSettings(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.GroupSettingWhereInput) (*generated.GroupSettingConnection, error) {
	panic(fmt.Errorf("not implemented: GroupSettings - groupSettings"))
}

// Integrations is the resolver for the integrations field.
func (r *queryResolver) Integrations(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *generated.IntegrationOrder, where *generated.IntegrationWhereInput) (*generated.IntegrationConnection, error) {
	return r.client.Integration.Query().Paginate(ctx, after, first, before, last, generated.WithIntegrationOrder(orderBy), generated.WithIntegrationFilter(where.Filter))
}

// OauthProviders is the resolver for the oauthProviders field.
func (r *queryResolver) OauthProviders(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.OauthProviderWhereInput) (*generated.OauthProviderConnection, error) {
	panic(fmt.Errorf("not implemented: OauthProviders - oauthProviders"))
}

// OhAuthTooTokens is the resolver for the ohAuthTooTokens field.
func (r *queryResolver) OhAuthTooTokens(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.OhAuthTooTokenWhereInput) (*generated.OhAuthTooTokenConnection, error) {
	panic(fmt.Errorf("not implemented: OhAuthTooTokens - ohAuthTooTokens"))
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *generated.OrganizationOrder, where *generated.OrganizationWhereInput) (*generated.OrganizationConnection, error) {
	// if auth is disabled, policy decisions will be skipped
	if r.authDisabled {
		ctx = privacy.DecisionContext(ctx, privacy.Allow)
	}

	return r.client.Organization.Query().Paginate(ctx, after, first, before, last, generated.WithOrganizationOrder(orderBy), generated.WithOrganizationFilter(where.Filter))
}

// OrganizationSettings is the resolver for the organizationSettings field.
func (r *queryResolver) OrganizationSettings(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.OrganizationSettingWhereInput) (*generated.OrganizationSettingConnection, error) {
	panic(fmt.Errorf("not implemented: OrganizationSettings - organizationSettings"))
}

// PersonalAccessTokens is the resolver for the personalAccessTokens field.
func (r *queryResolver) PersonalAccessTokens(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.PersonalAccessTokenWhereInput) (*generated.PersonalAccessTokenConnection, error) {
	panic(fmt.Errorf("not implemented: PersonalAccessTokens - personalAccessTokens"))
}

// RefreshTokens is the resolver for the refreshTokens field.
func (r *queryResolver) RefreshTokens(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.RefreshTokenWhereInput) (*generated.RefreshTokenConnection, error) {
	panic(fmt.Errorf("not implemented: RefreshTokens - refreshTokens"))
}

// Sessions is the resolver for the sessions field.
func (r *queryResolver) Sessions(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.SessionWhereInput) (*generated.SessionConnection, error) {
	return r.client.Session.Query().Paginate(ctx, after, first, before, last, generated.WithSessionFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *generated.UserOrder, where *generated.UserWhereInput) (*generated.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, generated.WithUserOrder(orderBy), generated.WithUserFilter(where.Filter))
}

// UserSettings is the resolver for the userSettings field.
func (r *queryResolver) UserSettings(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.UserSettingWhereInput) (*generated.UserSettingConnection, error) {
	panic(fmt.Errorf("not implemented: UserSettings - userSettings"))
}

// AuthStyle is the resolver for the authStyle field.
func (r *createOauthProviderInputResolver) AuthStyle(ctx context.Context, obj *generated.CreateOauthProviderInput, data int) error {
	panic(fmt.Errorf("not implemented: AuthStyle - authStyle"))
}

// AuthStyle is the resolver for the authStyle field.
func (r *oauthProviderWhereInputResolver) AuthStyle(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyle - authStyle"))
}

// AuthStyleNeq is the resolver for the authStyleNEQ field.
func (r *oauthProviderWhereInputResolver) AuthStyleNeq(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleNeq - authStyleNEQ"))
}

// AuthStyleIn is the resolver for the authStyleIn field.
func (r *oauthProviderWhereInputResolver) AuthStyleIn(ctx context.Context, obj *generated.OauthProviderWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: AuthStyleIn - authStyleIn"))
}

// AuthStyleNotIn is the resolver for the authStyleNotIn field.
func (r *oauthProviderWhereInputResolver) AuthStyleNotIn(ctx context.Context, obj *generated.OauthProviderWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: AuthStyleNotIn - authStyleNotIn"))
}

// AuthStyleGt is the resolver for the authStyleGT field.
func (r *oauthProviderWhereInputResolver) AuthStyleGt(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleGt - authStyleGT"))
}

// AuthStyleGte is the resolver for the authStyleGTE field.
func (r *oauthProviderWhereInputResolver) AuthStyleGte(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleGte - authStyleGTE"))
}

// AuthStyleLt is the resolver for the authStyleLT field.
func (r *oauthProviderWhereInputResolver) AuthStyleLt(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleLt - authStyleLT"))
}

// AuthStyleLte is the resolver for the authStyleLTE field.
func (r *oauthProviderWhereInputResolver) AuthStyleLte(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleLte - authStyleLTE"))
}

// AuthStyle is the resolver for the authStyle field.
func (r *updateOauthProviderInputResolver) AuthStyle(ctx context.Context, obj *generated.UpdateOauthProviderInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyle - authStyle"))
}

// OauthProvider returns OauthProviderResolver implementation.
func (r *Resolver) OauthProvider() OauthProviderResolver { return &oauthProviderResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// CreateOauthProviderInput returns CreateOauthProviderInputResolver implementation.
func (r *Resolver) CreateOauthProviderInput() CreateOauthProviderInputResolver {
	return &createOauthProviderInputResolver{r}
}

// OauthProviderWhereInput returns OauthProviderWhereInputResolver implementation.
func (r *Resolver) OauthProviderWhereInput() OauthProviderWhereInputResolver {
	return &oauthProviderWhereInputResolver{r}
}

// UpdateOauthProviderInput returns UpdateOauthProviderInputResolver implementation.
func (r *Resolver) UpdateOauthProviderInput() UpdateOauthProviderInputResolver {
	return &updateOauthProviderInputResolver{r}
}

type oauthProviderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type createOauthProviderInputResolver struct{ *Resolver }
type oauthProviderWhereInputResolver struct{ *Resolver }
type updateOauthProviderInputResolver struct{ *Resolver }
