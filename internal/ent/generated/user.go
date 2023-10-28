// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy uuid.UUID `json:"updated_by,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// The user's displayed 'friendly' name
	DisplayName string `json:"display_name,omitempty"`
	// user account is locked if unconfirmed or explicitly locked
	Locked bool `json:"locked,omitempty"`
	// URL of the user's remote avatar
	AvatarRemoteURL *string `json:"avatar_remote_url,omitempty"`
	// The user's local avatar file
	AvatarLocalFile *string `json:"avatar_local_file,omitempty"`
	// The time the user's (local) avatar was last updated
	AvatarUpdatedAt *time.Time `json:"avatar_updated_at,omitempty"`
	// The time the user was silenced
	SilencedAt *time.Time `json:"silenced_at,omitempty"`
	// The time the user was suspended
	SuspendedAt *time.Time `json:"suspended_at,omitempty"`
	// local Actor password recovery code generated during account creation
	RecoveryCode *string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Memberships holds the value of the memberships edge.
	Memberships []*Membership `json:"memberships,omitempty"`
	// Sessions holds the value of the sessions edge.
	Sessions []*Session `json:"sessions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedMemberships map[string][]*Membership
	namedSessions    map[string][]*Session
}

// MembershipsOrErr returns the Memberships value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MembershipsOrErr() ([]*Membership, error) {
	if e.loadedTypes[0] {
		return e.Memberships, nil
	}
	return nil, &NotLoadedError{edge: "memberships"}
}

// SessionsOrErr returns the Sessions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SessionsOrErr() ([]*Session, error) {
	if e.loadedTypes[1] {
		return e.Sessions, nil
	}
	return nil, &NotLoadedError{edge: "sessions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldLocked:
			values[i] = new(sql.NullBool)
		case user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldDisplayName, user.FieldAvatarRemoteURL, user.FieldAvatarLocalFile, user.FieldRecoveryCode:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldAvatarUpdatedAt, user.FieldSilencedAt, user.FieldSuspendedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID, user.FieldCreatedBy, user.FieldUpdatedBy:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				u.CreatedBy = *value
			}
		case user.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				u.UpdatedBy = *value
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldLocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				u.Locked = value.Bool
			}
		case user.FieldAvatarRemoteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_remote_url", values[i])
			} else if value.Valid {
				u.AvatarRemoteURL = new(string)
				*u.AvatarRemoteURL = value.String
			}
		case user.FieldAvatarLocalFile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_local_file", values[i])
			} else if value.Valid {
				u.AvatarLocalFile = new(string)
				*u.AvatarLocalFile = value.String
			}
		case user.FieldAvatarUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_updated_at", values[i])
			} else if value.Valid {
				u.AvatarUpdatedAt = new(time.Time)
				*u.AvatarUpdatedAt = value.Time
			}
		case user.FieldSilencedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field silenced_at", values[i])
			} else if value.Valid {
				u.SilencedAt = new(time.Time)
				*u.SilencedAt = value.Time
			}
		case user.FieldSuspendedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field suspended_at", values[i])
			} else if value.Valid {
				u.SuspendedAt = new(time.Time)
				*u.SuspendedAt = value.Time
			}
		case user.FieldRecoveryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recovery_code", values[i])
			} else if value.Valid {
				u.RecoveryCode = new(string)
				*u.RecoveryCode = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryMemberships queries the "memberships" edge of the User entity.
func (u *User) QueryMemberships() *MembershipQuery {
	return NewUserClient(u.config).QueryMemberships(u)
}

// QuerySessions queries the "sessions" edge of the User entity.
func (u *User) QuerySessions() *SessionQuery {
	return NewUserClient(u.config).QuerySessions(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("generated: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", u.Locked))
	builder.WriteString(", ")
	if v := u.AvatarRemoteURL; v != nil {
		builder.WriteString("avatar_remote_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.AvatarLocalFile; v != nil {
		builder.WriteString("avatar_local_file=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.AvatarUpdatedAt; v != nil {
		builder.WriteString("avatar_updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.SilencedAt; v != nil {
		builder.WriteString("silenced_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.SuspendedAt; v != nil {
		builder.WriteString("suspended_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("recovery_code=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// NamedMemberships returns the Memberships named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedMemberships(name string) ([]*Membership, error) {
	if u.Edges.namedMemberships == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedMemberships[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedMemberships(name string, edges ...*Membership) {
	if u.Edges.namedMemberships == nil {
		u.Edges.namedMemberships = make(map[string][]*Membership)
	}
	if len(edges) == 0 {
		u.Edges.namedMemberships[name] = []*Membership{}
	} else {
		u.Edges.namedMemberships[name] = append(u.Edges.namedMemberships[name], edges...)
	}
}

// NamedSessions returns the Sessions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedSessions(name string) ([]*Session, error) {
	if u.Edges.namedSessions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedSessions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedSessions(name string, edges ...*Session) {
	if u.Edges.namedSessions == nil {
		u.Edges.namedSessions = make(map[string][]*Session)
	}
	if len(edges) == 0 {
		u.Edges.namedSessions[name] = []*Session{}
	} else {
		u.Edges.namedSessions[name] = append(u.Edges.namedSessions[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
