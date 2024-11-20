// Code generated by ent, DO NOT EDIT.

package release

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the release type in the database.
	Label = "release"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldChannel holds the string denoting the channel field in the database.
	FieldChannel = "channel"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// FieldReleaseNotes holds the string denoting the release_notes field in the database.
	FieldReleaseNotes = "release_notes"
	// FieldFileURL holds the string denoting the file_url field in the database.
	FieldFileURL = "file_url"
	// FieldChecksum holds the string denoting the checksum field in the database.
	FieldChecksum = "checksum"
	// FieldIsCritical holds the string denoting the is_critical field in the database.
	FieldIsCritical = "is_critical"
	// FieldReleaseDate holds the string denoting the release_date field in the database.
	FieldReleaseDate = "release_date"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldArch holds the string denoting the arch field in the database.
	FieldArch = "arch"
	// EdgeAgents holds the string denoting the agents edge name in mutations.
	EdgeAgents = "agents"
	// AgentFieldID holds the string denoting the ID field of the Agent.
	AgentFieldID = "oid"
	// Table holds the table name of the release in the database.
	Table = "releases"
	// AgentsTable is the table that holds the agents relation/edge.
	AgentsTable = "agents"
	// AgentsInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	AgentsInverseTable = "agents"
	// AgentsColumn is the table column denoting the agents relation/edge.
	AgentsColumn = "release_agents"
)

// Columns holds all SQL columns for release fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldChannel,
	FieldSummary,
	FieldReleaseNotes,
	FieldFileURL,
	FieldChecksum,
	FieldIsCritical,
	FieldReleaseDate,
	FieldOs,
	FieldArch,
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

// OrderOption defines the ordering options for the Release queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByChannel orders the results by the channel field.
func ByChannel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannel, opts...).ToFunc()
}

// BySummary orders the results by the summary field.
func BySummary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSummary, opts...).ToFunc()
}

// ByReleaseNotes orders the results by the release_notes field.
func ByReleaseNotes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleaseNotes, opts...).ToFunc()
}

// ByFileURL orders the results by the file_url field.
func ByFileURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileURL, opts...).ToFunc()
}

// ByChecksum orders the results by the checksum field.
func ByChecksum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChecksum, opts...).ToFunc()
}

// ByIsCritical orders the results by the is_critical field.
func ByIsCritical(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsCritical, opts...).ToFunc()
}

// ByReleaseDate orders the results by the release_date field.
func ByReleaseDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleaseDate, opts...).ToFunc()
}

// ByOs orders the results by the os field.
func ByOs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOs, opts...).ToFunc()
}

// ByArch orders the results by the arch field.
func ByArch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArch, opts...).ToFunc()
}

// ByAgentsCount orders the results by agents count.
func ByAgentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAgentsStep(), opts...)
	}
}

// ByAgents orders the results by agents terms.
func ByAgents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAgentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAgentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AgentsInverseTable, AgentFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AgentsTable, AgentsColumn),
	)
}
