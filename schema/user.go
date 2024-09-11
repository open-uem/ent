package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().StorageKey("uid"),
		field.String("name"),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.Time("created").Optional().Default(time.Now()),
		field.Time("modified").Optional(),
	}
}