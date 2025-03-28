// Code generated by ent, DO NOT EDIT.

package share

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the share type in the database.
	Label = "share"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// AgentFieldID holds the string denoting the ID field of the Agent.
	AgentFieldID = "oid"
	// Table holds the table name of the share in the database.
	Table = "shares"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "shares"
	// OwnerInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	OwnerInverseTable = "agents"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "agent_shares"
)

// Columns holds all SQL columns for share fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldPath,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "shares"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"agent_shares",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Share queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, AgentFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
