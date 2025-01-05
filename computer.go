// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/computer"
)

// Computer is the model entity for the Computer schema.
type Computer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Manufacturer holds the value of the "manufacturer" field.
	Manufacturer string `json:"manufacturer,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Serial holds the value of the "serial" field.
	Serial string `json:"serial,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory uint64 `json:"memory,omitempty"`
	// Processor holds the value of the "processor" field.
	Processor string `json:"processor,omitempty"`
	// ProcessorCores holds the value of the "processor_cores" field.
	ProcessorCores int64 `json:"processor_cores,omitempty"`
	// ProcessorArch holds the value of the "processor_arch" field.
	ProcessorArch string `json:"processor_arch,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ComputerQuery when eager-loading is set.
	Edges          ComputerEdges `json:"edges"`
	agent_computer *string
	selectValues   sql.SelectValues
}

// ComputerEdges holds the relations/edges for other nodes in the graph.
type ComputerEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ComputerEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Computer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case computer.FieldID, computer.FieldMemory, computer.FieldProcessorCores:
			values[i] = new(sql.NullInt64)
		case computer.FieldManufacturer, computer.FieldModel, computer.FieldSerial, computer.FieldProcessor, computer.FieldProcessorArch:
			values[i] = new(sql.NullString)
		case computer.ForeignKeys[0]: // agent_computer
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Computer fields.
func (c *Computer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case computer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case computer.FieldManufacturer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer", values[i])
			} else if value.Valid {
				c.Manufacturer = value.String
			}
		case computer.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				c.Model = value.String
			}
		case computer.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				c.Serial = value.String
			}
		case computer.FieldMemory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				c.Memory = uint64(value.Int64)
			}
		case computer.FieldProcessor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field processor", values[i])
			} else if value.Valid {
				c.Processor = value.String
			}
		case computer.FieldProcessorCores:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field processor_cores", values[i])
			} else if value.Valid {
				c.ProcessorCores = value.Int64
			}
		case computer.FieldProcessorArch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field processor_arch", values[i])
			} else if value.Valid {
				c.ProcessorArch = value.String
			}
		case computer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_computer", values[i])
			} else if value.Valid {
				c.agent_computer = new(string)
				*c.agent_computer = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Computer.
// This includes values selected through modifiers, order, etc.
func (c *Computer) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Computer entity.
func (c *Computer) QueryOwner() *AgentQuery {
	return NewComputerClient(c.config).QueryOwner(c)
}

// Update returns a builder for updating this Computer.
// Note that you need to call Computer.Unwrap() before calling this method if this Computer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Computer) Update() *ComputerUpdateOne {
	return NewComputerClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Computer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Computer) Unwrap() *Computer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Computer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Computer) String() string {
	var builder strings.Builder
	builder.WriteString("Computer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("manufacturer=")
	builder.WriteString(c.Manufacturer)
	builder.WriteString(", ")
	builder.WriteString("model=")
	builder.WriteString(c.Model)
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(c.Serial)
	builder.WriteString(", ")
	builder.WriteString("memory=")
	builder.WriteString(fmt.Sprintf("%v", c.Memory))
	builder.WriteString(", ")
	builder.WriteString("processor=")
	builder.WriteString(c.Processor)
	builder.WriteString(", ")
	builder.WriteString("processor_cores=")
	builder.WriteString(fmt.Sprintf("%v", c.ProcessorCores))
	builder.WriteString(", ")
	builder.WriteString("processor_arch=")
	builder.WriteString(c.ProcessorArch)
	builder.WriteByte(')')
	return builder.String()
}

// Computers is a parsable slice of Computer.
type Computers []*Computer
