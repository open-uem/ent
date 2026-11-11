package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PendingUpdate holds the schema definition for the PendingUpdate entity.
type PendingUpdate struct {
	ent.Schema
}

// Fields of the PendingUpdate.
func (PendingUpdate) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("kb_article_ids").Optional(),
		field.Bool("is_downloaded").Optional().Default(false),
	}
}

// Edges of the PendingUpdate.
func (PendingUpdate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("pending_updates").Required(),
	}
}
