package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MDMCommand holds the schema definition for the Agent entity.
type MDMCommand struct {
	ent.Schema
}

// Fields of the Agent.
func (MDMCommand) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().StorageKey("uuid"),
		field.Time("when").Optional(),
		field.Enum("agent_status").Values("DeviceInformation", "UsersList", "InstalledApllicationsList").Optional().Default("DeviceInformation"),
	}
}

// Edges of the MDMCommand.
func (MDMCommand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agents", Agent.Type).Ref("mdmcommands").Unique(),
	}
}
