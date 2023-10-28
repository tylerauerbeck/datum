// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/membership"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/google/uuid"
)

// OrganizationUpdate is the builder for updating Organization entities.
type OrganizationUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationMutation
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ou *OrganizationUpdate) Where(ps ...predicate.Organization) *OrganizationUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrganizationUpdate) SetUpdatedAt(t time.Time) *OrganizationUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetCreatedBy sets the "created_by" field.
func (ou *OrganizationUpdate) SetCreatedBy(u uuid.UUID) *OrganizationUpdate {
	ou.mutation.SetCreatedBy(u)
	return ou
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableCreatedBy(u *uuid.UUID) *OrganizationUpdate {
	if u != nil {
		ou.SetCreatedBy(*u)
	}
	return ou
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ou *OrganizationUpdate) ClearCreatedBy() *OrganizationUpdate {
	ou.mutation.ClearCreatedBy()
	return ou
}

// SetUpdatedBy sets the "updated_by" field.
func (ou *OrganizationUpdate) SetUpdatedBy(u uuid.UUID) *OrganizationUpdate {
	ou.mutation.SetUpdatedBy(u)
	return ou
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableUpdatedBy(u *uuid.UUID) *OrganizationUpdate {
	if u != nil {
		ou.SetUpdatedBy(*u)
	}
	return ou
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ou *OrganizationUpdate) ClearUpdatedBy() *OrganizationUpdate {
	ou.mutation.ClearUpdatedBy()
	return ou
}

// SetName sets the "name" field.
func (ou *OrganizationUpdate) SetName(s string) *OrganizationUpdate {
	ou.mutation.SetName(s)
	return ou
}

// AddMembershipIDs adds the "memberships" edge to the Membership entity by IDs.
func (ou *OrganizationUpdate) AddMembershipIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.AddMembershipIDs(ids...)
	return ou
}

// AddMemberships adds the "memberships" edges to the Membership entity.
func (ou *OrganizationUpdate) AddMemberships(m ...*Membership) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ou.AddMembershipIDs(ids...)
}

// AddIntegrationIDs adds the "integrations" edge to the Integration entity by IDs.
func (ou *OrganizationUpdate) AddIntegrationIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.AddIntegrationIDs(ids...)
	return ou
}

// AddIntegrations adds the "integrations" edges to the Integration entity.
func (ou *OrganizationUpdate) AddIntegrations(i ...*Integration) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ou.AddIntegrationIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ou *OrganizationUpdate) Mutation() *OrganizationMutation {
	return ou.mutation
}

// ClearMemberships clears all "memberships" edges to the Membership entity.
func (ou *OrganizationUpdate) ClearMemberships() *OrganizationUpdate {
	ou.mutation.ClearMemberships()
	return ou
}

// RemoveMembershipIDs removes the "memberships" edge to Membership entities by IDs.
func (ou *OrganizationUpdate) RemoveMembershipIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.RemoveMembershipIDs(ids...)
	return ou
}

// RemoveMemberships removes "memberships" edges to Membership entities.
func (ou *OrganizationUpdate) RemoveMemberships(m ...*Membership) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ou.RemoveMembershipIDs(ids...)
}

// ClearIntegrations clears all "integrations" edges to the Integration entity.
func (ou *OrganizationUpdate) ClearIntegrations() *OrganizationUpdate {
	ou.mutation.ClearIntegrations()
	return ou
}

// RemoveIntegrationIDs removes the "integrations" edge to Integration entities by IDs.
func (ou *OrganizationUpdate) RemoveIntegrationIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.RemoveIntegrationIDs(ids...)
	return ou
}

// RemoveIntegrations removes "integrations" edges to Integration entities.
func (ou *OrganizationUpdate) RemoveIntegrations(i ...*Integration) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ou.RemoveIntegrationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrganizationUpdate) Save(ctx context.Context) (int, error) {
	if err := ou.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrganizationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrganizationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrganizationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrganizationUpdate) defaults() error {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		if organization.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized organization.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := organization.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrganizationUpdate) check() error {
	if v, ok := ou.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Organization.name": %w`, err)}
		}
	}
	return nil
}

func (ou *OrganizationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.CreatedBy(); ok {
		_spec.SetField(organization.FieldCreatedBy, field.TypeUUID, value)
	}
	if ou.mutation.CreatedByCleared() {
		_spec.ClearField(organization.FieldCreatedBy, field.TypeUUID)
	}
	if value, ok := ou.mutation.UpdatedBy(); ok {
		_spec.SetField(organization.FieldUpdatedBy, field.TypeUUID, value)
	}
	if ou.mutation.UpdatedByCleared() {
		_spec.ClearField(organization.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if ou.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.MembershipsTable,
			Columns: []string{organization.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedMembershipsIDs(); len(nodes) > 0 && !ou.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.MembershipsTable,
			Columns: []string{organization.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.MembershipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.MembershipsTable,
			Columns: []string{organization.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.IntegrationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.IntegrationsTable,
			Columns: []string{organization.IntegrationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedIntegrationsIDs(); len(nodes) > 0 && !ou.mutation.IntegrationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.IntegrationsTable,
			Columns: []string{organization.IntegrationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.IntegrationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.IntegrationsTable,
			Columns: []string{organization.IntegrationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrganizationUpdateOne is the builder for updating a single Organization entity.
type OrganizationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrganizationUpdateOne) SetUpdatedAt(t time.Time) *OrganizationUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetCreatedBy sets the "created_by" field.
func (ouo *OrganizationUpdateOne) SetCreatedBy(u uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.SetCreatedBy(u)
	return ouo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *OrganizationUpdateOne {
	if u != nil {
		ouo.SetCreatedBy(*u)
	}
	return ouo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ouo *OrganizationUpdateOne) ClearCreatedBy() *OrganizationUpdateOne {
	ouo.mutation.ClearCreatedBy()
	return ouo
}

// SetUpdatedBy sets the "updated_by" field.
func (ouo *OrganizationUpdateOne) SetUpdatedBy(u uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.SetUpdatedBy(u)
	return ouo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *OrganizationUpdateOne {
	if u != nil {
		ouo.SetUpdatedBy(*u)
	}
	return ouo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ouo *OrganizationUpdateOne) ClearUpdatedBy() *OrganizationUpdateOne {
	ouo.mutation.ClearUpdatedBy()
	return ouo
}

// SetName sets the "name" field.
func (ouo *OrganizationUpdateOne) SetName(s string) *OrganizationUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// AddMembershipIDs adds the "memberships" edge to the Membership entity by IDs.
func (ouo *OrganizationUpdateOne) AddMembershipIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.AddMembershipIDs(ids...)
	return ouo
}

// AddMemberships adds the "memberships" edges to the Membership entity.
func (ouo *OrganizationUpdateOne) AddMemberships(m ...*Membership) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ouo.AddMembershipIDs(ids...)
}

// AddIntegrationIDs adds the "integrations" edge to the Integration entity by IDs.
func (ouo *OrganizationUpdateOne) AddIntegrationIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.AddIntegrationIDs(ids...)
	return ouo
}

// AddIntegrations adds the "integrations" edges to the Integration entity.
func (ouo *OrganizationUpdateOne) AddIntegrations(i ...*Integration) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ouo.AddIntegrationIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ouo *OrganizationUpdateOne) Mutation() *OrganizationMutation {
	return ouo.mutation
}

// ClearMemberships clears all "memberships" edges to the Membership entity.
func (ouo *OrganizationUpdateOne) ClearMemberships() *OrganizationUpdateOne {
	ouo.mutation.ClearMemberships()
	return ouo
}

// RemoveMembershipIDs removes the "memberships" edge to Membership entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveMembershipIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.RemoveMembershipIDs(ids...)
	return ouo
}

// RemoveMemberships removes "memberships" edges to Membership entities.
func (ouo *OrganizationUpdateOne) RemoveMemberships(m ...*Membership) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ouo.RemoveMembershipIDs(ids...)
}

// ClearIntegrations clears all "integrations" edges to the Integration entity.
func (ouo *OrganizationUpdateOne) ClearIntegrations() *OrganizationUpdateOne {
	ouo.mutation.ClearIntegrations()
	return ouo
}

// RemoveIntegrationIDs removes the "integrations" edge to Integration entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveIntegrationIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.RemoveIntegrationIDs(ids...)
	return ouo
}

// RemoveIntegrations removes "integrations" edges to Integration entities.
func (ouo *OrganizationUpdateOne) RemoveIntegrations(i ...*Integration) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ouo.RemoveIntegrationIDs(ids...)
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ouo *OrganizationUpdateOne) Where(ps ...predicate.Organization) *OrganizationUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrganizationUpdateOne) Select(field string, fields ...string) *OrganizationUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Organization entity.
func (ouo *OrganizationUpdateOne) Save(ctx context.Context) (*Organization, error) {
	if err := ouo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) SaveX(ctx context.Context) *Organization {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrganizationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrganizationUpdateOne) defaults() error {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		if organization.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized organization.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := organization.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrganizationUpdateOne) check() error {
	if v, ok := ouo.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Organization.name": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrganizationUpdateOne) sqlSave(ctx context.Context) (_node *Organization, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Organization.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for _, f := range fields {
			if !organization.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.CreatedBy(); ok {
		_spec.SetField(organization.FieldCreatedBy, field.TypeUUID, value)
	}
	if ouo.mutation.CreatedByCleared() {
		_spec.ClearField(organization.FieldCreatedBy, field.TypeUUID)
	}
	if value, ok := ouo.mutation.UpdatedBy(); ok {
		_spec.SetField(organization.FieldUpdatedBy, field.TypeUUID, value)
	}
	if ouo.mutation.UpdatedByCleared() {
		_spec.ClearField(organization.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if ouo.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.MembershipsTable,
			Columns: []string{organization.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedMembershipsIDs(); len(nodes) > 0 && !ouo.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.MembershipsTable,
			Columns: []string{organization.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.MembershipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.MembershipsTable,
			Columns: []string{organization.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.IntegrationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.IntegrationsTable,
			Columns: []string{organization.IntegrationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedIntegrationsIDs(); len(nodes) > 0 && !ouo.mutation.IntegrationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.IntegrationsTable,
			Columns: []string{organization.IntegrationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.IntegrationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.IntegrationsTable,
			Columns: []string{organization.IntegrationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Organization{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
