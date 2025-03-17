// Code generated by ent, DO NOT EDIT.

package profileissue

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the profileissue type in the database.
	Label = "profile_issue"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// FieldWhen holds the string denoting the when field in the database.
	FieldWhen = "when"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeAgents holds the string denoting the agents edge name in mutations.
	EdgeAgents = "agents"
	// AgentFieldID holds the string denoting the ID field of the Agent.
	AgentFieldID = "oid"
	// Table holds the table name of the profileissue in the database.
	Table = "profile_issues"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "profile_issues"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "profile_issues"
	// AgentsTable is the table that holds the agents relation/edge.
	AgentsTable = "profile_issues"
	// AgentsInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	AgentsInverseTable = "agents"
	// AgentsColumn is the table column denoting the agents relation/edge.
	AgentsColumn = "profile_issue_agents"
)

// Columns holds all SQL columns for profileissue fields.
var Columns = []string{
	FieldID,
	FieldError,
	FieldWhen,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "profile_issues"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profile_issues",
	"profile_issue_agents",
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

var (
	// DefaultWhen holds the default value on creation for the "when" field.
	DefaultWhen time.Time
	// UpdateDefaultWhen holds the default value on update for the "when" field.
	UpdateDefaultWhen func() time.Time
)

// OrderOption defines the ordering options for the ProfileIssue queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByError orders the results by the error field.
func ByError(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldError, opts...).ToFunc()
}

// ByWhen orders the results by the when field.
func ByWhen(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhen, opts...).ToFunc()
}

// ByProfileField orders the results by profile field.
func ByProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByAgentsField orders the results by agents field.
func ByAgentsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAgentsStep(), sql.OrderByField(field, opts...))
	}
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
	)
}
func newAgentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AgentsInverseTable, AgentFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AgentsTable, AgentsColumn),
	)
}
