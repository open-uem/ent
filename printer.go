// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/printer"
)

// Printer is the model entity for the Printer schema.
type Printer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Port holds the value of the "port" field.
	Port string `json:"port,omitempty"`
	// IsDefault holds the value of the "is_default" field.
	IsDefault bool `json:"is_default,omitempty"`
	// IsNetwork holds the value of the "is_network" field.
	IsNetwork bool `json:"is_network,omitempty"`
	// IsShared holds the value of the "is_shared" field.
	IsShared bool `json:"is_shared,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PrinterQuery when eager-loading is set.
	Edges          PrinterEdges `json:"edges"`
	agent_printers *string
	selectValues   sql.SelectValues
}

// PrinterEdges holds the relations/edges for other nodes in the graph.
type PrinterEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrinterEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Printer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case printer.FieldIsDefault, printer.FieldIsNetwork, printer.FieldIsShared:
			values[i] = new(sql.NullBool)
		case printer.FieldID:
			values[i] = new(sql.NullInt64)
		case printer.FieldName, printer.FieldPort:
			values[i] = new(sql.NullString)
		case printer.ForeignKeys[0]: // agent_printers
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Printer fields.
func (pr *Printer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case printer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case printer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case printer.FieldPort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				pr.Port = value.String
			}
		case printer.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				pr.IsDefault = value.Bool
			}
		case printer.FieldIsNetwork:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_network", values[i])
			} else if value.Valid {
				pr.IsNetwork = value.Bool
			}
		case printer.FieldIsShared:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_shared", values[i])
			} else if value.Valid {
				pr.IsShared = value.Bool
			}
		case printer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_printers", values[i])
			} else if value.Valid {
				pr.agent_printers = new(string)
				*pr.agent_printers = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Printer.
// This includes values selected through modifiers, order, etc.
func (pr *Printer) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Printer entity.
func (pr *Printer) QueryOwner() *AgentQuery {
	return NewPrinterClient(pr.config).QueryOwner(pr)
}

// Update returns a builder for updating this Printer.
// Note that you need to call Printer.Unwrap() before calling this method if this Printer
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Printer) Update() *PrinterUpdateOne {
	return NewPrinterClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Printer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Printer) Unwrap() *Printer {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Printer is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Printer) String() string {
	var builder strings.Builder
	builder.WriteString("Printer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("port=")
	builder.WriteString(pr.Port)
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsDefault))
	builder.WriteString(", ")
	builder.WriteString("is_network=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsNetwork))
	builder.WriteString(", ")
	builder.WriteString("is_shared=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsShared))
	builder.WriteByte(')')
	return builder.String()
}

// Printers is a parsable slice of Printer.
type Printers []*Printer
