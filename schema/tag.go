package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("tag").NotEmpty(),
		field.String("description").Optional(),
		field.String("color"),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Ref("tags"),
		edge.To("children", Tag.Type).From("parent").Unique(),
		edge.From("profile", Profile.Type).Ref("tags"),
		edge.From("tenant", Tenant.Type).Unique().Ref("tags"),
	}
}

// Indexes of the Tag.
func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tag").Edges("tenant").Unique(),
	}
}
