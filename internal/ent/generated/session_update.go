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
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/nanox"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SessionUpdate) SetUpdatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetCreatedBy sets the "created_by" field.
func (su *SessionUpdate) SetCreatedBy(s string) *SessionUpdate {
	su.mutation.SetCreatedBy(s)
	return su
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCreatedBy(s *string) *SessionUpdate {
	if s != nil {
		su.SetCreatedBy(*s)
	}
	return su
}

// ClearCreatedBy clears the value of the "created_by" field.
func (su *SessionUpdate) ClearCreatedBy() *SessionUpdate {
	su.mutation.ClearCreatedBy()
	return su
}

// SetUpdatedBy sets the "updated_by" field.
func (su *SessionUpdate) SetUpdatedBy(s string) *SessionUpdate {
	su.mutation.SetUpdatedBy(s)
	return su
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUpdatedBy(s *string) *SessionUpdate {
	if s != nil {
		su.SetUpdatedBy(*s)
	}
	return su
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (su *SessionUpdate) ClearUpdatedBy() *SessionUpdate {
	su.mutation.ClearUpdatedBy()
	return su
}

// SetDisabled sets the "disabled" field.
func (su *SessionUpdate) SetDisabled(b bool) *SessionUpdate {
	su.mutation.SetDisabled(b)
	return su
}

// SetUserAgent sets the "user_agent" field.
func (su *SessionUpdate) SetUserAgent(s string) *SessionUpdate {
	su.mutation.SetUserAgent(s)
	return su
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUserAgent(s *string) *SessionUpdate {
	if s != nil {
		su.SetUserAgent(*s)
	}
	return su
}

// ClearUserAgent clears the value of the "user_agent" field.
func (su *SessionUpdate) ClearUserAgent() *SessionUpdate {
	su.mutation.ClearUserAgent()
	return su
}

// SetIps sets the "ips" field.
func (su *SessionUpdate) SetIps(s string) *SessionUpdate {
	su.mutation.SetIps(s)
	return su
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (su *SessionUpdate) SetUsersID(id nanox.ID) *SessionUpdate {
	su.mutation.SetUsersID(id)
	return su
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (su *SessionUpdate) SetNillableUsersID(id *nanox.ID) *SessionUpdate {
	if id != nil {
		su = su.SetUsersID(*id)
	}
	return su
}

// SetUsers sets the "users" edge to the User entity.
func (su *SessionUpdate) SetUsers(u *User) *SessionUpdate {
	return su.SetUsersID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (su *SessionUpdate) ClearUsers() *SessionUpdate {
	su.mutation.ClearUsers()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	if err := su.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SessionUpdate) defaults() error {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		if session.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized session.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := session.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.CreatedBy(); ok {
		_spec.SetField(session.FieldCreatedBy, field.TypeString, value)
	}
	if su.mutation.CreatedByCleared() {
		_spec.ClearField(session.FieldCreatedBy, field.TypeString)
	}
	if value, ok := su.mutation.UpdatedBy(); ok {
		_spec.SetField(session.FieldUpdatedBy, field.TypeString, value)
	}
	if su.mutation.UpdatedByCleared() {
		_spec.ClearField(session.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := su.mutation.Disabled(); ok {
		_spec.SetField(session.FieldDisabled, field.TypeBool, value)
	}
	if value, ok := su.mutation.UserAgent(); ok {
		_spec.SetField(session.FieldUserAgent, field.TypeString, value)
	}
	if su.mutation.UserAgentCleared() {
		_spec.ClearField(session.FieldUserAgent, field.TypeString)
	}
	if value, ok := su.mutation.Ips(); ok {
		_spec.SetField(session.FieldIps, field.TypeString, value)
	}
	if su.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.UsersTable,
			Columns: []string{session.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.Session
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.UsersTable,
			Columns: []string{session.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.Session
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = su.schemaConfig.Session
	ctx = internal.NewSchemaConfigContext(ctx, su.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SessionUpdateOne) SetUpdatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetCreatedBy sets the "created_by" field.
func (suo *SessionUpdateOne) SetCreatedBy(s string) *SessionUpdateOne {
	suo.mutation.SetCreatedBy(s)
	return suo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCreatedBy(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetCreatedBy(*s)
	}
	return suo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (suo *SessionUpdateOne) ClearCreatedBy() *SessionUpdateOne {
	suo.mutation.ClearCreatedBy()
	return suo
}

// SetUpdatedBy sets the "updated_by" field.
func (suo *SessionUpdateOne) SetUpdatedBy(s string) *SessionUpdateOne {
	suo.mutation.SetUpdatedBy(s)
	return suo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUpdatedBy(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetUpdatedBy(*s)
	}
	return suo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (suo *SessionUpdateOne) ClearUpdatedBy() *SessionUpdateOne {
	suo.mutation.ClearUpdatedBy()
	return suo
}

// SetDisabled sets the "disabled" field.
func (suo *SessionUpdateOne) SetDisabled(b bool) *SessionUpdateOne {
	suo.mutation.SetDisabled(b)
	return suo
}

// SetUserAgent sets the "user_agent" field.
func (suo *SessionUpdateOne) SetUserAgent(s string) *SessionUpdateOne {
	suo.mutation.SetUserAgent(s)
	return suo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserAgent(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetUserAgent(*s)
	}
	return suo
}

// ClearUserAgent clears the value of the "user_agent" field.
func (suo *SessionUpdateOne) ClearUserAgent() *SessionUpdateOne {
	suo.mutation.ClearUserAgent()
	return suo
}

// SetIps sets the "ips" field.
func (suo *SessionUpdateOne) SetIps(s string) *SessionUpdateOne {
	suo.mutation.SetIps(s)
	return suo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (suo *SessionUpdateOne) SetUsersID(id nanox.ID) *SessionUpdateOne {
	suo.mutation.SetUsersID(id)
	return suo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUsersID(id *nanox.ID) *SessionUpdateOne {
	if id != nil {
		suo = suo.SetUsersID(*id)
	}
	return suo
}

// SetUsers sets the "users" edge to the User entity.
func (suo *SessionUpdateOne) SetUsers(u *User) *SessionUpdateOne {
	return suo.SetUsersID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (suo *SessionUpdateOne) ClearUsers() *SessionUpdateOne {
	suo.mutation.ClearUsers()
	return suo
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SessionUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		if session.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized session.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := session.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.CreatedBy(); ok {
		_spec.SetField(session.FieldCreatedBy, field.TypeString, value)
	}
	if suo.mutation.CreatedByCleared() {
		_spec.ClearField(session.FieldCreatedBy, field.TypeString)
	}
	if value, ok := suo.mutation.UpdatedBy(); ok {
		_spec.SetField(session.FieldUpdatedBy, field.TypeString, value)
	}
	if suo.mutation.UpdatedByCleared() {
		_spec.ClearField(session.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := suo.mutation.Disabled(); ok {
		_spec.SetField(session.FieldDisabled, field.TypeBool, value)
	}
	if value, ok := suo.mutation.UserAgent(); ok {
		_spec.SetField(session.FieldUserAgent, field.TypeString, value)
	}
	if suo.mutation.UserAgentCleared() {
		_spec.ClearField(session.FieldUserAgent, field.TypeString)
	}
	if value, ok := suo.mutation.Ips(); ok {
		_spec.SetField(session.FieldIps, field.TypeString, value)
	}
	if suo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.UsersTable,
			Columns: []string{session.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.Session
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   session.UsersTable,
			Columns: []string{session.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.Session
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = suo.schemaConfig.Session
	ctx = internal.NewSchemaConfigContext(ctx, suo.schemaConfig)
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
