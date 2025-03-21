// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/antivirus"
)

// Antivirus is the model entity for the Antivirus schema.
type Antivirus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// IsUpdated holds the value of the "is_updated" field.
	IsUpdated bool `json:"is_updated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AntivirusQuery when eager-loading is set.
	Edges           AntivirusEdges `json:"edges"`
	agent_antivirus *string
	selectValues    sql.SelectValues
}

// AntivirusEdges holds the relations/edges for other nodes in the graph.
type AntivirusEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AntivirusEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Antivirus) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case antivirus.FieldIsActive, antivirus.FieldIsUpdated:
			values[i] = new(sql.NullBool)
		case antivirus.FieldID:
			values[i] = new(sql.NullInt64)
		case antivirus.FieldName:
			values[i] = new(sql.NullString)
		case antivirus.ForeignKeys[0]: // agent_antivirus
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Antivirus fields.
func (a *Antivirus) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case antivirus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case antivirus.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case antivirus.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				a.IsActive = value.Bool
			}
		case antivirus.FieldIsUpdated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_updated", values[i])
			} else if value.Valid {
				a.IsUpdated = value.Bool
			}
		case antivirus.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_antivirus", values[i])
			} else if value.Valid {
				a.agent_antivirus = new(string)
				*a.agent_antivirus = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Antivirus.
// This includes values selected through modifiers, order, etc.
func (a *Antivirus) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Antivirus entity.
func (a *Antivirus) QueryOwner() *AgentQuery {
	return NewAntivirusClient(a.config).QueryOwner(a)
}

// Update returns a builder for updating this Antivirus.
// Note that you need to call Antivirus.Unwrap() before calling this method if this Antivirus
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Antivirus) Update() *AntivirusUpdateOne {
	return NewAntivirusClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Antivirus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Antivirus) Unwrap() *Antivirus {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Antivirus is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Antivirus) String() string {
	var builder strings.Builder
	builder.WriteString("Antivirus(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", a.IsActive))
	builder.WriteString(", ")
	builder.WriteString("is_updated=")
	builder.WriteString(fmt.Sprintf("%v", a.IsUpdated))
	builder.WriteByte(')')
	return builder.String()
}

// Antiviri is a parsable slice of Antivirus.
type Antiviri []*Antivirus
