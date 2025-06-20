package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Enum("type").Values("winget_install", "winget_update", "winget_delete", "add_registry_key", "update_registry_key_default_value", "add_registry_key_value", "remove_registry_key", "remove_registry_key_value", "add_local_user", "remove_local_user", "add_local_group", "remove_local_group", "add_users_to_local_group", "remove_users_from_local_group", "msi_install", "msi_uninstall"),
		field.String("package_id").Optional().Default(""),
		field.String("package_name").Optional().Default(""),
		field.String("registry_key").Optional().Default(""),
		field.String("registry_key_value_name").Optional().Default(""),
		field.Enum("registry_key_value_type").Values("String", "Binary", "DWord", "QWord", "MultiString", "ExpandString").Optional(),
		field.String("registry_key_value_data").Optional().Default(""),
		field.Bool("registry_hex").Optional().Default(false),
		field.Bool("registry_force").Optional().Default(false),
		field.String("local_user_username").Optional().Default(""),
		field.String("local_user_description").Optional().Default(""),
		field.Bool("local_user_disable").Optional().Default(false),
		field.String("local_user_fullname").Optional().Default(""),
		field.String("local_user_password").Optional().Default(""),
		field.Bool("local_user_password_change_not_allowed").Optional().Default(false),
		field.Bool("local_user_password_change_required").Optional().Default(false),
		field.Bool("local_user_password_never_expires").Optional().Default(false),
		field.String("local_group_name").Optional().Default(""),
		field.String("local_group_description").Optional().Default(""),
		field.String("local_group_members").Optional().Default(""),
		field.String("local_group_members_to_include").Optional().Default(""),
		field.String("local_group_members_to_exclude").Optional().Default(""),
		field.String("msi_productid").Optional().Default(""),
		field.String("msi_path").Optional().Default(""),
		field.String("msi_arguments").Optional().Default(""),
		field.String("msi_file_hash").Optional().Default(""),
		field.Enum("msi_file_hash_alg").Values("MD5", "RIPEMD160", "SHA1", "SHA256", "SHA384", "SHA512").Optional(),
		field.String("msi_log_path").Optional().Default(""),
		field.Time("when").Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.From("profile", Profile.Type).Unique().Ref("tasks"),
	}
}
