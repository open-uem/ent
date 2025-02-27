// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/systemupdate"
)

// SystemUpdate is the model entity for the SystemUpdate schema.
type SystemUpdate struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SystemUpdateStatus holds the value of the "system_update_status" field.
	SystemUpdateStatus string `json:"system_update_status,omitempty"`
	// LastInstall holds the value of the "last_install" field.
	LastInstall time.Time `json:"last_install,omitempty"`
	// LastSearch holds the value of the "last_search" field.
	LastSearch time.Time `json:"last_search,omitempty"`
	// PendingUpdates holds the value of the "pending_updates" field.
	PendingUpdates bool `json:"pending_updates,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SystemUpdateQuery when eager-loading is set.
	Edges              SystemUpdateEdges `json:"edges"`
	agent_systemupdate *string
	selectValues       sql.SelectValues
}

// SystemUpdateEdges holds the relations/edges for other nodes in the graph.
type SystemUpdateEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SystemUpdateEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SystemUpdate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case systemupdate.FieldPendingUpdates:
			values[i] = new(sql.NullBool)
		case systemupdate.FieldID:
			values[i] = new(sql.NullInt64)
		case systemupdate.FieldSystemUpdateStatus:
			values[i] = new(sql.NullString)
		case systemupdate.FieldLastInstall, systemupdate.FieldLastSearch:
			values[i] = new(sql.NullTime)
		case systemupdate.ForeignKeys[0]: // agent_systemupdate
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SystemUpdate fields.
func (su *SystemUpdate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case systemupdate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			su.ID = int(value.Int64)
		case systemupdate.FieldSystemUpdateStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field system_update_status", values[i])
			} else if value.Valid {
				su.SystemUpdateStatus = value.String
			}
		case systemupdate.FieldLastInstall:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_install", values[i])
			} else if value.Valid {
				su.LastInstall = value.Time
			}
		case systemupdate.FieldLastSearch:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_search", values[i])
			} else if value.Valid {
				su.LastSearch = value.Time
			}
		case systemupdate.FieldPendingUpdates:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pending_updates", values[i])
			} else if value.Valid {
				su.PendingUpdates = value.Bool
			}
		case systemupdate.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_systemupdate", values[i])
			} else if value.Valid {
				su.agent_systemupdate = new(string)
				*su.agent_systemupdate = value.String
			}
		default:
			su.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SystemUpdate.
// This includes values selected through modifiers, order, etc.
func (su *SystemUpdate) Value(name string) (ent.Value, error) {
	return su.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the SystemUpdate entity.
func (su *SystemUpdate) QueryOwner() *AgentQuery {
	return NewSystemUpdateClient(su.config).QueryOwner(su)
}

// Update returns a builder for updating this SystemUpdate.
// Note that you need to call SystemUpdate.Unwrap() before calling this method if this SystemUpdate
// was returned from a transaction, and the transaction was committed or rolled back.
func (su *SystemUpdate) Update() *SystemUpdateUpdateOne {
	return NewSystemUpdateClient(su.config).UpdateOne(su)
}

// Unwrap unwraps the SystemUpdate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (su *SystemUpdate) Unwrap() *SystemUpdate {
	_tx, ok := su.config.driver.(*txDriver)
	if !ok {
		panic("ent: SystemUpdate is not a transactional entity")
	}
	su.config.driver = _tx.drv
	return su
}

// String implements the fmt.Stringer.
func (su *SystemUpdate) String() string {
	var builder strings.Builder
	builder.WriteString("SystemUpdate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", su.ID))
	builder.WriteString("system_update_status=")
	builder.WriteString(su.SystemUpdateStatus)
	builder.WriteString(", ")
	builder.WriteString("last_install=")
	builder.WriteString(su.LastInstall.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_search=")
	builder.WriteString(su.LastSearch.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("pending_updates=")
	builder.WriteString(fmt.Sprintf("%v", su.PendingUpdates))
	builder.WriteByte(')')
	return builder.String()
}

// SystemUpdates is a parsable slice of SystemUpdate.
type SystemUpdates []*SystemUpdate
