// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/revocation"
)

// Revocation is the model entity for the Revocation schema.
type Revocation struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason int `json:"reason,omitempty"`
	// Info holds the value of the "info" field.
	Info string `json:"info,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
	// Revoked holds the value of the "revoked" field.
	Revoked      time.Time `json:"revoked,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Revocation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case revocation.FieldID, revocation.FieldReason:
			values[i] = new(sql.NullInt64)
		case revocation.FieldInfo:
			values[i] = new(sql.NullString)
		case revocation.FieldExpiry, revocation.FieldRevoked:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Revocation fields.
func (r *Revocation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case revocation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int64(value.Int64)
		case revocation.FieldReason:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				r.Reason = int(value.Int64)
			}
		case revocation.FieldInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value.Valid {
				r.Info = value.String
			}
		case revocation.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				r.Expiry = value.Time
			}
		case revocation.FieldRevoked:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field revoked", values[i])
			} else if value.Valid {
				r.Revoked = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Revocation.
// This includes values selected through modifiers, order, etc.
func (r *Revocation) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Revocation.
// Note that you need to call Revocation.Unwrap() before calling this method if this Revocation
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Revocation) Update() *RevocationUpdateOne {
	return NewRevocationClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Revocation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Revocation) Unwrap() *Revocation {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Revocation is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Revocation) String() string {
	var builder strings.Builder
	builder.WriteString("Revocation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("reason=")
	builder.WriteString(fmt.Sprintf("%v", r.Reason))
	builder.WriteString(", ")
	builder.WriteString("info=")
	builder.WriteString(r.Info)
	builder.WriteString(", ")
	builder.WriteString("expiry=")
	builder.WriteString(r.Expiry.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("revoked=")
	builder.WriteString(r.Revoked.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Revocations is a parsable slice of Revocation.
type Revocations []*Revocation
