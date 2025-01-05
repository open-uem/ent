// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/certificate"
)

// Certificate is the model entity for the Certificate schema.
type Certificate struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type certificate.Type `json:"type,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
	// UID holds the value of the "uid" field.
	UID          string `json:"uid,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Certificate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case certificate.FieldID:
			values[i] = new(sql.NullInt64)
		case certificate.FieldType, certificate.FieldDescription, certificate.FieldUID:
			values[i] = new(sql.NullString)
		case certificate.FieldExpiry:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Certificate fields.
func (c *Certificate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case certificate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case certificate.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = certificate.Type(value.String)
			}
		case certificate.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case certificate.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				c.Expiry = value.Time
			}
		case certificate.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				c.UID = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Certificate.
// This includes values selected through modifiers, order, etc.
func (c *Certificate) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Certificate.
// Note that you need to call Certificate.Unwrap() before calling this method if this Certificate
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Certificate) Update() *CertificateUpdateOne {
	return NewCertificateClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Certificate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Certificate) Unwrap() *Certificate {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Certificate is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Certificate) String() string {
	var builder strings.Builder
	builder.WriteString("Certificate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("expiry=")
	builder.WriteString(c.Expiry.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("uid=")
	builder.WriteString(c.UID)
	builder.WriteByte(')')
	return builder.String()
}

// Certificates is a parsable slice of Certificate.
type Certificates []*Certificate
