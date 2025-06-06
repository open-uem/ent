// Code generated by ent, DO NOT EDIT.

package tag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// AgentFieldID holds the string denoting the ID field of the Agent.
	AgentFieldID = "oid"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// OwnerTable is the table that holds the owner relation/edge. The primary key declared below.
	OwnerTable = "agent_tags"
	// OwnerInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	OwnerInverseTable = "agents"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "tags"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "tag_children"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "tags"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "tag_children"
	// ProfileTable is the table that holds the profile relation/edge. The primary key declared below.
	ProfileTable = "profile_tags"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "tags"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "tenant_tags"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldTag,
	FieldDescription,
	FieldColor,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tags"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"tag_children",
	"task_tags",
	"tenant_tags",
}

var (
	// OwnerPrimaryKey and OwnerColumn2 are the table columns denoting the
	// primary key for the owner relation (M2M).
	OwnerPrimaryKey = []string{"agent_id", "tag_id"}
	// ProfilePrimaryKey and ProfileColumn2 are the table columns denoting the
	// primary key for the profile relation (M2M).
	ProfilePrimaryKey = []string{"profile_id", "tag_id"}
)

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
	// TagValidator is a validator for the "tag" field. It is called by the builders before save.
	TagValidator func(string) error
)

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTag orders the results by the tag field.
func ByTag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTag, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByOwnerCount orders the results by owner count.
func ByOwnerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOwnerStep(), opts...)
	}
}

// ByOwner orders the results by owner terms.
func ByOwner(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProfileCount orders the results by profile count.
func ByProfileCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProfileStep(), opts...)
	}
}

// ByProfile orders the results by profile terms.
func ByProfile(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTenantField orders the results by tenant field.
func ByTenantField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTenantStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, AgentFieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, OwnerTable, OwnerPrimaryKey...),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProfileTable, ProfilePrimaryKey...),
	)
}
func newTenantStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TenantInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TenantTable, TenantColumn),
	)
}
