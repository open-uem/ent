package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NanoMDMSettings holds the schema definition for the NanoMDMSettings entity.
type NanoMDMSettings struct {
	ent.Schema
}

// Fields of the NanoMDMSettings.
func (NanoMDMSettings) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Optional().Default(""),
		field.String("password").Optional().Default(""),
		field.String("server_url").Optional().Default(""),
		field.String("ca_cer_file").Optional().Default(""),
		field.String("profile_payload_id").Optional().Default(""),
	}
}

// Edges of the NanoMDMSettings.
func (NanoMDMSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Ref("nanomdm"),
	}
}
