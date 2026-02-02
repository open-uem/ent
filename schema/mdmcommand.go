package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MDMCommand holds the schema definition for the MDMCommand entity.
type MDMCommand struct {
	ent.Schema
}

// Fields of the MDMCommand.
func (MDMCommand) Fields() []ent.Field {
	return []ent.Field{
		field.Time("when").Optional(),
		field.Enum("type").Values("DeviceInformation", "UsersList", "InstalledApllicationsList").Optional().Default("DeviceInformation"),
	}
}

// Edges of the MDMCommand.
func (MDMCommand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agents", Agent.Type).Ref("mdmcommands").Unique(),
	}
}
