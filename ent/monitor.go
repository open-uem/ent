// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/doncicuto/openuem-ent/ent/agent"
	"github.com/doncicuto/openuem-ent/ent/monitor"
)

// Monitor is the model entity for the Monitor schema.
type Monitor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Manufacturer holds the value of the "manufacturer" field.
	Manufacturer string `json:"manufacturer,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Serial holds the value of the "serial" field.
	Serial string `json:"serial,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MonitorQuery when eager-loading is set.
	Edges          MonitorEdges `json:"edges"`
	agent_monitors *string
	selectValues   sql.SelectValues
}

// MonitorEdges holds the relations/edges for other nodes in the graph.
type MonitorEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MonitorEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Monitor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case monitor.FieldID:
			values[i] = new(sql.NullInt64)
		case monitor.FieldManufacturer, monitor.FieldModel, monitor.FieldSerial:
			values[i] = new(sql.NullString)
		case monitor.ForeignKeys[0]: // agent_monitors
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Monitor fields.
func (m *Monitor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case monitor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case monitor.FieldManufacturer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer", values[i])
			} else if value.Valid {
				m.Manufacturer = value.String
			}
		case monitor.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				m.Model = value.String
			}
		case monitor.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				m.Serial = value.String
			}
		case monitor.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_monitors", values[i])
			} else if value.Valid {
				m.agent_monitors = new(string)
				*m.agent_monitors = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Monitor.
// This includes values selected through modifiers, order, etc.
func (m *Monitor) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Monitor entity.
func (m *Monitor) QueryOwner() *AgentQuery {
	return NewMonitorClient(m.config).QueryOwner(m)
}

// Update returns a builder for updating this Monitor.
// Note that you need to call Monitor.Unwrap() before calling this method if this Monitor
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Monitor) Update() *MonitorUpdateOne {
	return NewMonitorClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Monitor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Monitor) Unwrap() *Monitor {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Monitor is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Monitor) String() string {
	var builder strings.Builder
	builder.WriteString("Monitor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("manufacturer=")
	builder.WriteString(m.Manufacturer)
	builder.WriteString(", ")
	builder.WriteString("model=")
	builder.WriteString(m.Model)
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(m.Serial)
	builder.WriteByte(')')
	return builder.String()
}

// Monitors is a parsable slice of Monitor.
type Monitors []*Monitor
