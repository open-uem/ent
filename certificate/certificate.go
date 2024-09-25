// Code generated by ent, DO NOT EDIT.

package certificate

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the certificate type in the database.
	Label = "certificate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "serial"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// Table holds the table name of the certificate in the database.
	Table = "certificates"
)

// Columns holds all SQL columns for certificate fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldDescription,
	FieldExpiry,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeConsole Type = "console"
	TypeWorker  Type = "worker"
	TypeAgent   Type = "agent"
	TypeUser    Type = "user"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeConsole, TypeWorker, TypeAgent, TypeUser:
		return nil
	default:
		return fmt.Errorf("certificate: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Certificate queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByExpiry orders the results by the expiry field.
func ByExpiry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiry, opts...).ToFunc()
}
