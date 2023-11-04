// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/nanox"
)

// Integration is the model entity for the Integration schema.
type Integration struct {
	config `json:"-"`
	// ID of the ent.
	ID nanox.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// SecretName holds the value of the "secret_name" field.
	SecretName string `json:"secret_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IntegrationQuery when eager-loading is set.
	Edges                     IntegrationEdges `json:"edges"`
	organization_integrations *nanox.ID
	selectValues              sql.SelectValues
}

// IntegrationEdges holds the relations/edges for other nodes in the graph.
type IntegrationEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IntegrationEdges) OwnerOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Integration) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case integration.FieldID:
			values[i] = new(nanox.ID)
		case integration.FieldCreatedBy, integration.FieldUpdatedBy, integration.FieldName, integration.FieldKind, integration.FieldDescription, integration.FieldSecretName:
			values[i] = new(sql.NullString)
		case integration.FieldCreatedAt, integration.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case integration.ForeignKeys[0]: // organization_integrations
			values[i] = &sql.NullScanner{S: new(nanox.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Integration fields.
func (i *Integration) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case integration.FieldID:
			if value, ok := values[j].(*nanox.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case integration.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case integration.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case integration.FieldCreatedBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[j])
			} else if value.Valid {
				i.CreatedBy = value.String
			}
		case integration.FieldUpdatedBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[j])
			} else if value.Valid {
				i.UpdatedBy = value.String
			}
		case integration.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case integration.FieldKind:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[j])
			} else if value.Valid {
				i.Kind = value.String
			}
		case integration.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		case integration.FieldSecretName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_name", values[j])
			} else if value.Valid {
				i.SecretName = value.String
			}
		case integration.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_integrations", values[j])
			} else if value.Valid {
				i.organization_integrations = new(nanox.ID)
				*i.organization_integrations = *value.S.(*nanox.ID)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Integration.
// This includes values selected through modifiers, order, etc.
func (i *Integration) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Integration entity.
func (i *Integration) QueryOwner() *OrganizationQuery {
	return NewIntegrationClient(i.config).QueryOwner(i)
}

// Update returns a builder for updating this Integration.
// Note that you need to call Integration.Unwrap() before calling this method if this Integration
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Integration) Update() *IntegrationUpdateOne {
	return NewIntegrationClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Integration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Integration) Unwrap() *Integration {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("generated: Integration is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Integration) String() string {
	var builder strings.Builder
	builder.WriteString("Integration(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(i.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(i.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(i.Kind)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteString(", ")
	builder.WriteString("secret_name=")
	builder.WriteString(i.SecretName)
	builder.WriteByte(')')
	return builder.String()
}

// Integrations is a parsable slice of Integration.
type Integrations []*Integration
