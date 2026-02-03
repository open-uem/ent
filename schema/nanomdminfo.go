package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NanoMDMInfo holds the schema definition for the Agent entity.
type NanoMDMInfo struct {
	ent.Schema
}

// Fields of the NanoMDMInfo.
func (NanoMDMInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Float("available_device_capacity").Optional().Default(0),
		field.Bool("awaiting_configuration").Optional().Default(false),
		field.Float("battery_level").Optional().Default(-1),
		field.String("bluetooth_mac").Optional().Default(""),
		field.String("build_version").Optional().Default(""),
		field.String("current_console_managed_user").Optional().Default(""),
		field.Float("device_capacity").Optional().Default(-1),
		field.String("device_name").Optional().Default(""),
		field.String("eacs_preflight").Optional().Default(""),
		field.String("ethernet_mac").Optional().Default(""),
		field.Bool("has_battery").Optional().Default(false),
		field.String("hostname").Optional().Default(""),
		field.Bool("is_activation_lock_enabled").Optional().Default(false),
		field.Bool("is_activation_lock_supported").Optional().Default(false),
		field.Bool("is_apple_silicon").Optional().Default(false),
		field.Bool("is_supervised").Optional().Default(false),
		field.String("localhostname").Optional().Default(""),
		field.String("model").Optional().Default(""),
		field.String("model_name").Optional().Default(""),
		field.Bool("auto_check_enabled").Optional().Default(false),
		field.Bool("automatic_app_installation_enabled").Optional().Default(false),
		field.Bool("automatic_os_installation_enabled").Optional().Default(false),
		field.Bool("automatic_security_updates_enabled").Optional().Default(false),
		field.Bool("background_download_enabled").Optional().Default(false),
		field.String("catalog_url").Optional().Default(""),
		field.Bool("is_default_catalog").Optional().Default(false),
		field.Time("previous_scan_date").Optional(),
		field.Int("previous_scan_result").Optional().Default(0),
		field.String("os_version").Optional().Default(""),
		field.Bool("pin_required_for_device_lock").Optional().Default(false),
		field.Bool("pin_required_for_erase_device").Optional().Default(false),
		field.String("product_name").Optional().Default(""),
		field.String("provisioning_udid").Optional().Default(""),
		field.String("serial_number").Optional().Default(""),
		field.String("software_update_device_id").Optional().Default(""),
		field.String("supplemental_build_version").Optional().Default(""),
		field.Bool("supports_lom_device").Optional().Default(false),
		field.Bool("supports_ios_app_installs").Optional().Default(false),
		field.Bool("system_integrity_protection_enabled").Optional().Default(false),
		field.String("udid").Optional().Default(""),
	}
}

// Edges of the NanoMDMInfo.
func (NanoMDMInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agent", Agent.Type).Ref("nanomdminfo").Unique(),
	}
}
