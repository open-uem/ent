package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NanoMDMUser holds the schema definition for the Agent entity.
type NanoMDMUser struct {
	ent.Schema
}

// Fields of the NanoMDMUser.
func (NanoMDMUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("data_quota").Optional(),
		field.Int("data_used").Optional(),
		field.Bool("has_data_to_sync").Optional(),
		field.Bool("has_secure_token").Optional(),
		field.Bool("is_logged_in").Optional(),
		field.String("username").Optional(),
		field.String("fullname").Optional(),
		field.Bool("mobile_account").Optional(),
		field.Int("uid").Optional(),
		field.String("user_guid").Optional(),
	}
}

// Edges of the NanoMDMUser.
func (NanoMDMUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agent", Agent.Type).Ref("nanomdmusers").Unique(),
	}
}
