package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MDMCommand holds the schema definition for the MDMCommand entity.
type MDMCommand struct {
	ent.Schema
}

// Fields of the MDMCommand.
func (MDMCommand) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().StorageKey("uuid"),
		field.Time("when").Optional(),
		field.Enum("type").Values("DeviceInformation", "UsersList", "InstalledApplicationsList", "RemoveProfile").Optional().Default("DeviceInformation"),
		field.String("agentID"),
	}
}
