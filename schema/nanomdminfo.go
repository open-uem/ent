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
		field.Float("available_device_capacity").Default(0),
		field.Bool("awaiting_configuration").Default(false),
		field.Float("battery_level").Default(-1),
		field.String("bluetooth_mac").Default(""),
		field.String("build_version").Default(""),
		field.String("current_console_managed_user").Default(""),
		field.Float("device_capacity").Default(-1),
		field.String("device_name").Default(""),
		field.String("eacs_preflight").Default(""),
		field.String("ethernet_mac").Default(""),
		field.Bool("has_battery").Default(false),
		field.String("hostname").Default(""),
		field.Bool("is_activation_lock_enabled").Default(false),
		field.Bool("is_activation_lock_supported").Default(false),
		field.Bool("is_apple_silicon").Default(false),
		field.Bool("is_supervised").Default(false),
		field.String("localhostname").Default(""),
		field.String("model").Default(""),
		field.String("model_name").Default(""),
		field.Bool("auto_check_enabled").Default(false),
		field.Bool("automatic_app_installation_enabled").Default(false),
		field.Bool("automatic_os_installation_enabled").Default(false),
		field.Bool("automatic_security_updates_enabled").Default(false),
		field.Bool("background_download_enabled").Default(false),
		field.String("catalog_url").Default(""),
		field.Bool("is_default_catalog").Default(false),
		field.Time("previous_scan_date"),
		field.Int("previous_scan_result").Default(0),
		field.String("os_version").Default(""),
		field.Bool("pin_required_for_device_lock").Default(false),
		field.Bool("pin_required_for_erase_device").Default(false),
		field.String("product_name").Default(""),
		field.String("provisioning_udid").Default(""),
		field.String("serial_number").Default(""),
		field.String("software_update_device_id").Default(""),
		field.String("supplemental_build_version").Default(""),
		field.Bool("supports_lom_device").Default(false),
		field.Bool("supports_ios_app_installs").Default(false),
		field.Bool("system_integrity_protection_enabled").Default(false),
		field.String("udid").Default(""),
	}
}

// Edges of the NanoMDMInfo.
func (NanoMDMInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agent", Agent.Type).Ref("nanomdminfo").Unique(),
	}
}
