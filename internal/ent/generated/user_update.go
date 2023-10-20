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
	"github.com/datumforge/datum/internal/ent/generated/membership"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetCreatedBy sets the "created_by" field.
func (uu *UserUpdate) SetCreatedBy(i int) *UserUpdate {
	uu.mutation.ResetCreatedBy()
	uu.mutation.SetCreatedBy(i)
	return uu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedBy(i *int) *UserUpdate {
	if i != nil {
		uu.SetCreatedBy(*i)
	}
	return uu
}

// AddCreatedBy adds i to the "created_by" field.
func (uu *UserUpdate) AddCreatedBy(i int) *UserUpdate {
	uu.mutation.AddCreatedBy(i)
	return uu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (uu *UserUpdate) ClearCreatedBy() *UserUpdate {
	uu.mutation.ClearCreatedBy()
	return uu
}

// SetUpdatedBy sets the "updated_by" field.
func (uu *UserUpdate) SetUpdatedBy(i int) *UserUpdate {
	uu.mutation.ResetUpdatedBy()
	uu.mutation.SetUpdatedBy(i)
	return uu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedBy(i *int) *UserUpdate {
	if i != nil {
		uu.SetUpdatedBy(*i)
	}
	return uu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (uu *UserUpdate) AddUpdatedBy(i int) *UserUpdate {
	uu.mutation.AddUpdatedBy(i)
	return uu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uu *UserUpdate) ClearUpdatedBy() *UserUpdate {
	uu.mutation.ClearUpdatedBy()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetDisplayName sets the "display_name" field.
func (uu *UserUpdate) SetDisplayName(s string) *UserUpdate {
	uu.mutation.SetDisplayName(s)
	return uu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDisplayName(s *string) *UserUpdate {
	if s != nil {
		uu.SetDisplayName(*s)
	}
	return uu
}

// SetLocked sets the "locked" field.
func (uu *UserUpdate) SetLocked(b bool) *UserUpdate {
	uu.mutation.SetLocked(b)
	return uu
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLocked(b *bool) *UserUpdate {
	if b != nil {
		uu.SetLocked(*b)
	}
	return uu
}

// SetAvatarRemoteURL sets the "avatar_remote_url" field.
func (uu *UserUpdate) SetAvatarRemoteURL(s string) *UserUpdate {
	uu.mutation.SetAvatarRemoteURL(s)
	return uu
}

// SetNillableAvatarRemoteURL sets the "avatar_remote_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatarRemoteURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatarRemoteURL(*s)
	}
	return uu
}

// ClearAvatarRemoteURL clears the value of the "avatar_remote_url" field.
func (uu *UserUpdate) ClearAvatarRemoteURL() *UserUpdate {
	uu.mutation.ClearAvatarRemoteURL()
	return uu
}

// SetAvatarLocalFile sets the "avatar_local_file" field.
func (uu *UserUpdate) SetAvatarLocalFile(s string) *UserUpdate {
	uu.mutation.SetAvatarLocalFile(s)
	return uu
}

// SetNillableAvatarLocalFile sets the "avatar_local_file" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatarLocalFile(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatarLocalFile(*s)
	}
	return uu
}

// ClearAvatarLocalFile clears the value of the "avatar_local_file" field.
func (uu *UserUpdate) ClearAvatarLocalFile() *UserUpdate {
	uu.mutation.ClearAvatarLocalFile()
	return uu
}

// SetAvatarUpdatedAt sets the "avatar_updated_at" field.
func (uu *UserUpdate) SetAvatarUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetAvatarUpdatedAt(t)
	return uu
}

// SetNillableAvatarUpdatedAt sets the "avatar_updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatarUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetAvatarUpdatedAt(*t)
	}
	return uu
}

// ClearAvatarUpdatedAt clears the value of the "avatar_updated_at" field.
func (uu *UserUpdate) ClearAvatarUpdatedAt() *UserUpdate {
	uu.mutation.ClearAvatarUpdatedAt()
	return uu
}

// SetSilencedAt sets the "silenced_at" field.
func (uu *UserUpdate) SetSilencedAt(t time.Time) *UserUpdate {
	uu.mutation.SetSilencedAt(t)
	return uu
}

// SetNillableSilencedAt sets the "silenced_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSilencedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetSilencedAt(*t)
	}
	return uu
}

// ClearSilencedAt clears the value of the "silenced_at" field.
func (uu *UserUpdate) ClearSilencedAt() *UserUpdate {
	uu.mutation.ClearSilencedAt()
	return uu
}

// SetSuspendedAt sets the "suspended_at" field.
func (uu *UserUpdate) SetSuspendedAt(t time.Time) *UserUpdate {
	uu.mutation.SetSuspendedAt(t)
	return uu
}

// SetNillableSuspendedAt sets the "suspended_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSuspendedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetSuspendedAt(*t)
	}
	return uu
}

// ClearSuspendedAt clears the value of the "suspended_at" field.
func (uu *UserUpdate) ClearSuspendedAt() *UserUpdate {
	uu.mutation.ClearSuspendedAt()
	return uu
}

// SetRecoveryCode sets the "recovery_code" field.
func (uu *UserUpdate) SetRecoveryCode(s string) *UserUpdate {
	uu.mutation.SetRecoveryCode(s)
	return uu
}

// SetNillableRecoveryCode sets the "recovery_code" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRecoveryCode(s *string) *UserUpdate {
	if s != nil {
		uu.SetRecoveryCode(*s)
	}
	return uu
}

// ClearRecoveryCode clears the value of the "recovery_code" field.
func (uu *UserUpdate) ClearRecoveryCode() *UserUpdate {
	uu.mutation.ClearRecoveryCode()
	return uu
}

// AddMembershipIDs adds the "memberships" edge to the Membership entity by IDs.
func (uu *UserUpdate) AddMembershipIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddMembershipIDs(ids...)
	return uu
}

// AddMemberships adds the "memberships" edges to the Membership entity.
func (uu *UserUpdate) AddMemberships(m ...*Membership) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.AddMembershipIDs(ids...)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (uu *UserUpdate) AddSessionIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddSessionIDs(ids...)
	return uu
}

// AddSessions adds the "sessions" edges to the Session entity.
func (uu *UserUpdate) AddSessions(s ...*Session) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSessionIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearMemberships clears all "memberships" edges to the Membership entity.
func (uu *UserUpdate) ClearMemberships() *UserUpdate {
	uu.mutation.ClearMemberships()
	return uu
}

// RemoveMembershipIDs removes the "memberships" edge to Membership entities by IDs.
func (uu *UserUpdate) RemoveMembershipIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveMembershipIDs(ids...)
	return uu
}

// RemoveMemberships removes "memberships" edges to Membership entities.
func (uu *UserUpdate) RemoveMemberships(m ...*Membership) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.RemoveMembershipIDs(ids...)
}

// ClearSessions clears all "sessions" edges to the Session entity.
func (uu *UserUpdate) ClearSessions() *UserUpdate {
	uu.mutation.ClearSessions()
	return uu
}

// RemoveSessionIDs removes the "sessions" edge to Session entities by IDs.
func (uu *UserUpdate) RemoveSessionIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveSessionIDs(ids...)
	return uu
}

// RemoveSessions removes "sessions" edges to Session entities.
func (uu *UserUpdate) RemoveSessions(s ...*Session) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized user.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`generated: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`generated: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.DisplayName(); ok {
		if err := user.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`generated: validator failed for field "User.display_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.AvatarRemoteURL(); ok {
		if err := user.AvatarRemoteURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_remote_url", err: fmt.Errorf(`generated: validator failed for field "User.avatar_remote_url": %w`, err)}
		}
	}
	if v, ok := uu.mutation.AvatarLocalFile(); ok {
		if err := user.AvatarLocalFileValidator(v); err != nil {
			return &ValidationError{Name: "avatar_local_file", err: fmt.Errorf(`generated: validator failed for field "User.avatar_local_file": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(user.FieldCreatedBy, field.TypeInt, value)
	}
	if uu.mutation.CreatedByCleared() {
		_spec.ClearField(user.FieldCreatedBy, field.TypeInt)
	}
	if value, ok := uu.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(user.FieldUpdatedBy, field.TypeInt, value)
	}
	if uu.mutation.UpdatedByCleared() {
		_spec.ClearField(user.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Locked(); ok {
		_spec.SetField(user.FieldLocked, field.TypeBool, value)
	}
	if value, ok := uu.mutation.AvatarRemoteURL(); ok {
		_spec.SetField(user.FieldAvatarRemoteURL, field.TypeString, value)
	}
	if uu.mutation.AvatarRemoteURLCleared() {
		_spec.ClearField(user.FieldAvatarRemoteURL, field.TypeString)
	}
	if value, ok := uu.mutation.AvatarLocalFile(); ok {
		_spec.SetField(user.FieldAvatarLocalFile, field.TypeString, value)
	}
	if uu.mutation.AvatarLocalFileCleared() {
		_spec.ClearField(user.FieldAvatarLocalFile, field.TypeString)
	}
	if value, ok := uu.mutation.AvatarUpdatedAt(); ok {
		_spec.SetField(user.FieldAvatarUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.AvatarUpdatedAtCleared() {
		_spec.ClearField(user.FieldAvatarUpdatedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.SilencedAt(); ok {
		_spec.SetField(user.FieldSilencedAt, field.TypeTime, value)
	}
	if uu.mutation.SilencedAtCleared() {
		_spec.ClearField(user.FieldSilencedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.SuspendedAt(); ok {
		_spec.SetField(user.FieldSuspendedAt, field.TypeTime, value)
	}
	if uu.mutation.SuspendedAtCleared() {
		_spec.ClearField(user.FieldSuspendedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.RecoveryCode(); ok {
		_spec.SetField(user.FieldRecoveryCode, field.TypeString, value)
	}
	if uu.mutation.RecoveryCodeCleared() {
		_spec.ClearField(user.FieldRecoveryCode, field.TypeString)
	}
	if uu.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MembershipsTable,
			Columns: []string{user.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedMembershipsIDs(); len(nodes) > 0 && !uu.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MembershipsTable,
			Columns: []string{user.MembershipsColumn},
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
	if nodes := uu.mutation.MembershipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MembershipsTable,
			Columns: []string{user.MembershipsColumn},
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
	if uu.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !uu.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetCreatedBy sets the "created_by" field.
func (uuo *UserUpdateOne) SetCreatedBy(i int) *UserUpdateOne {
	uuo.mutation.ResetCreatedBy()
	uuo.mutation.SetCreatedBy(i)
	return uuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedBy(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetCreatedBy(*i)
	}
	return uuo
}

// AddCreatedBy adds i to the "created_by" field.
func (uuo *UserUpdateOne) AddCreatedBy(i int) *UserUpdateOne {
	uuo.mutation.AddCreatedBy(i)
	return uuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (uuo *UserUpdateOne) ClearCreatedBy() *UserUpdateOne {
	uuo.mutation.ClearCreatedBy()
	return uuo
}

// SetUpdatedBy sets the "updated_by" field.
func (uuo *UserUpdateOne) SetUpdatedBy(i int) *UserUpdateOne {
	uuo.mutation.ResetUpdatedBy()
	uuo.mutation.SetUpdatedBy(i)
	return uuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedBy(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetUpdatedBy(*i)
	}
	return uuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (uuo *UserUpdateOne) AddUpdatedBy(i int) *UserUpdateOne {
	uuo.mutation.AddUpdatedBy(i)
	return uuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uuo *UserUpdateOne) ClearUpdatedBy() *UserUpdateOne {
	uuo.mutation.ClearUpdatedBy()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetDisplayName sets the "display_name" field.
func (uuo *UserUpdateOne) SetDisplayName(s string) *UserUpdateOne {
	uuo.mutation.SetDisplayName(s)
	return uuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDisplayName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDisplayName(*s)
	}
	return uuo
}

// SetLocked sets the "locked" field.
func (uuo *UserUpdateOne) SetLocked(b bool) *UserUpdateOne {
	uuo.mutation.SetLocked(b)
	return uuo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLocked(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetLocked(*b)
	}
	return uuo
}

// SetAvatarRemoteURL sets the "avatar_remote_url" field.
func (uuo *UserUpdateOne) SetAvatarRemoteURL(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarRemoteURL(s)
	return uuo
}

// SetNillableAvatarRemoteURL sets the "avatar_remote_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatarRemoteURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatarRemoteURL(*s)
	}
	return uuo
}

// ClearAvatarRemoteURL clears the value of the "avatar_remote_url" field.
func (uuo *UserUpdateOne) ClearAvatarRemoteURL() *UserUpdateOne {
	uuo.mutation.ClearAvatarRemoteURL()
	return uuo
}

// SetAvatarLocalFile sets the "avatar_local_file" field.
func (uuo *UserUpdateOne) SetAvatarLocalFile(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarLocalFile(s)
	return uuo
}

// SetNillableAvatarLocalFile sets the "avatar_local_file" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatarLocalFile(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatarLocalFile(*s)
	}
	return uuo
}

// ClearAvatarLocalFile clears the value of the "avatar_local_file" field.
func (uuo *UserUpdateOne) ClearAvatarLocalFile() *UserUpdateOne {
	uuo.mutation.ClearAvatarLocalFile()
	return uuo
}

// SetAvatarUpdatedAt sets the "avatar_updated_at" field.
func (uuo *UserUpdateOne) SetAvatarUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetAvatarUpdatedAt(t)
	return uuo
}

// SetNillableAvatarUpdatedAt sets the "avatar_updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatarUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetAvatarUpdatedAt(*t)
	}
	return uuo
}

// ClearAvatarUpdatedAt clears the value of the "avatar_updated_at" field.
func (uuo *UserUpdateOne) ClearAvatarUpdatedAt() *UserUpdateOne {
	uuo.mutation.ClearAvatarUpdatedAt()
	return uuo
}

// SetSilencedAt sets the "silenced_at" field.
func (uuo *UserUpdateOne) SetSilencedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetSilencedAt(t)
	return uuo
}

// SetNillableSilencedAt sets the "silenced_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSilencedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetSilencedAt(*t)
	}
	return uuo
}

// ClearSilencedAt clears the value of the "silenced_at" field.
func (uuo *UserUpdateOne) ClearSilencedAt() *UserUpdateOne {
	uuo.mutation.ClearSilencedAt()
	return uuo
}

// SetSuspendedAt sets the "suspended_at" field.
func (uuo *UserUpdateOne) SetSuspendedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetSuspendedAt(t)
	return uuo
}

// SetNillableSuspendedAt sets the "suspended_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSuspendedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetSuspendedAt(*t)
	}
	return uuo
}

// ClearSuspendedAt clears the value of the "suspended_at" field.
func (uuo *UserUpdateOne) ClearSuspendedAt() *UserUpdateOne {
	uuo.mutation.ClearSuspendedAt()
	return uuo
}

// SetRecoveryCode sets the "recovery_code" field.
func (uuo *UserUpdateOne) SetRecoveryCode(s string) *UserUpdateOne {
	uuo.mutation.SetRecoveryCode(s)
	return uuo
}

// SetNillableRecoveryCode sets the "recovery_code" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRecoveryCode(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRecoveryCode(*s)
	}
	return uuo
}

// ClearRecoveryCode clears the value of the "recovery_code" field.
func (uuo *UserUpdateOne) ClearRecoveryCode() *UserUpdateOne {
	uuo.mutation.ClearRecoveryCode()
	return uuo
}

// AddMembershipIDs adds the "memberships" edge to the Membership entity by IDs.
func (uuo *UserUpdateOne) AddMembershipIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddMembershipIDs(ids...)
	return uuo
}

// AddMemberships adds the "memberships" edges to the Membership entity.
func (uuo *UserUpdateOne) AddMemberships(m ...*Membership) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.AddMembershipIDs(ids...)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (uuo *UserUpdateOne) AddSessionIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddSessionIDs(ids...)
	return uuo
}

// AddSessions adds the "sessions" edges to the Session entity.
func (uuo *UserUpdateOne) AddSessions(s ...*Session) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSessionIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearMemberships clears all "memberships" edges to the Membership entity.
func (uuo *UserUpdateOne) ClearMemberships() *UserUpdateOne {
	uuo.mutation.ClearMemberships()
	return uuo
}

// RemoveMembershipIDs removes the "memberships" edge to Membership entities by IDs.
func (uuo *UserUpdateOne) RemoveMembershipIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveMembershipIDs(ids...)
	return uuo
}

// RemoveMemberships removes "memberships" edges to Membership entities.
func (uuo *UserUpdateOne) RemoveMemberships(m ...*Membership) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.RemoveMembershipIDs(ids...)
}

// ClearSessions clears all "sessions" edges to the Session entity.
func (uuo *UserUpdateOne) ClearSessions() *UserUpdateOne {
	uuo.mutation.ClearSessions()
	return uuo
}

// RemoveSessionIDs removes the "sessions" edge to Session entities by IDs.
func (uuo *UserUpdateOne) RemoveSessionIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveSessionIDs(ids...)
	return uuo
}

// RemoveSessions removes "sessions" edges to Session entities.
func (uuo *UserUpdateOne) RemoveSessions(s ...*Session) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSessionIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized user.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`generated: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`generated: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.DisplayName(); ok {
		if err := user.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`generated: validator failed for field "User.display_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.AvatarRemoteURL(); ok {
		if err := user.AvatarRemoteURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_remote_url", err: fmt.Errorf(`generated: validator failed for field "User.avatar_remote_url": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.AvatarLocalFile(); ok {
		if err := user.AvatarLocalFileValidator(v); err != nil {
			return &ValidationError{Name: "avatar_local_file", err: fmt.Errorf(`generated: validator failed for field "User.avatar_local_file": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(user.FieldCreatedBy, field.TypeInt, value)
	}
	if uuo.mutation.CreatedByCleared() {
		_spec.ClearField(user.FieldCreatedBy, field.TypeInt)
	}
	if value, ok := uuo.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(user.FieldUpdatedBy, field.TypeInt, value)
	}
	if uuo.mutation.UpdatedByCleared() {
		_spec.ClearField(user.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Locked(); ok {
		_spec.SetField(user.FieldLocked, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.AvatarRemoteURL(); ok {
		_spec.SetField(user.FieldAvatarRemoteURL, field.TypeString, value)
	}
	if uuo.mutation.AvatarRemoteURLCleared() {
		_spec.ClearField(user.FieldAvatarRemoteURL, field.TypeString)
	}
	if value, ok := uuo.mutation.AvatarLocalFile(); ok {
		_spec.SetField(user.FieldAvatarLocalFile, field.TypeString, value)
	}
	if uuo.mutation.AvatarLocalFileCleared() {
		_spec.ClearField(user.FieldAvatarLocalFile, field.TypeString)
	}
	if value, ok := uuo.mutation.AvatarUpdatedAt(); ok {
		_spec.SetField(user.FieldAvatarUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.AvatarUpdatedAtCleared() {
		_spec.ClearField(user.FieldAvatarUpdatedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.SilencedAt(); ok {
		_spec.SetField(user.FieldSilencedAt, field.TypeTime, value)
	}
	if uuo.mutation.SilencedAtCleared() {
		_spec.ClearField(user.FieldSilencedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.SuspendedAt(); ok {
		_spec.SetField(user.FieldSuspendedAt, field.TypeTime, value)
	}
	if uuo.mutation.SuspendedAtCleared() {
		_spec.ClearField(user.FieldSuspendedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.RecoveryCode(); ok {
		_spec.SetField(user.FieldRecoveryCode, field.TypeString, value)
	}
	if uuo.mutation.RecoveryCodeCleared() {
		_spec.ClearField(user.FieldRecoveryCode, field.TypeString)
	}
	if uuo.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MembershipsTable,
			Columns: []string{user.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedMembershipsIDs(); len(nodes) > 0 && !uuo.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MembershipsTable,
			Columns: []string{user.MembershipsColumn},
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
	if nodes := uuo.mutation.MembershipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MembershipsTable,
			Columns: []string{user.MembershipsColumn},
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
	if uuo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !uuo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
