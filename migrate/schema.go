// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgentsColumns holds the columns for the "agents" table.
	AgentsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeString, Unique: true},
		{Name: "os", Type: field.TypeString},
		{Name: "hostname", Type: field.TypeString},
		{Name: "ip", Type: field.TypeString, Default: ""},
		{Name: "mac", Type: field.TypeString, Default: ""},
		{Name: "first_contact", Type: field.TypeTime, Nullable: true},
		{Name: "last_contact", Type: field.TypeTime, Nullable: true},
		{Name: "vnc", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "notes", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "update_task_status", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "update_task_description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "update_task_result", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "update_task_execution", Type: field.TypeTime, Nullable: true},
		{Name: "update_task_version", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "vnc_proxy_port", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "sftp_port", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"WaitingForAdmission", "Enabled", "Disabled"}, Default: "WaitingForAdmission"},
		{Name: "certificate_ready", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "restart_required", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "release_agents", Type: field.TypeInt, Nullable: true},
	}
	// AgentsTable holds the schema information for the "agents" table.
	AgentsTable = &schema.Table{
		Name:       "agents",
		Columns:    AgentsColumns,
		PrimaryKey: []*schema.Column{AgentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "agents_releases_agents",
				Columns:    []*schema.Column{AgentsColumns[19]},
				RefColumns: []*schema.Column{ReleasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AntiviriColumns holds the columns for the "antiviri" table.
	AntiviriColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "is_updated", Type: field.TypeBool},
		{Name: "agent_antivirus", Type: field.TypeString, Unique: true},
	}
	// AntiviriTable holds the schema information for the "antiviri" table.
	AntiviriTable = &schema.Table{
		Name:       "antiviri",
		Columns:    AntiviriColumns,
		PrimaryKey: []*schema.Column{AntiviriColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "antiviri_agents_antivirus",
				Columns:    []*schema.Column{AntiviriColumns[4]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AppsColumns holds the columns for the "apps" table.
	AppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "publisher", Type: field.TypeString, Nullable: true},
		{Name: "install_date", Type: field.TypeString, Nullable: true},
		{Name: "agent_apps", Type: field.TypeString},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:       "apps",
		Columns:    AppsColumns,
		PrimaryKey: []*schema.Column{AppsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "apps_agents_apps",
				Columns:    []*schema.Column{AppsColumns[5]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CertificatesColumns holds the columns for the "certificates" table.
	CertificatesColumns = []*schema.Column{
		{Name: "serial", Type: field.TypeInt64, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"console", "worker", "agent", "user", "ocsp", "nats"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, Nullable: true},
		{Name: "uid", Type: field.TypeString, Nullable: true},
	}
	// CertificatesTable holds the schema information for the "certificates" table.
	CertificatesTable = &schema.Table{
		Name:       "certificates",
		Columns:    CertificatesColumns,
		PrimaryKey: []*schema.Column{CertificatesColumns[0]},
	}
	// ComputersColumns holds the columns for the "computers" table.
	ComputersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "manufacturer", Type: field.TypeString, Nullable: true},
		{Name: "model", Type: field.TypeString, Nullable: true},
		{Name: "serial", Type: field.TypeString, Nullable: true},
		{Name: "memory", Type: field.TypeUint64, Nullable: true},
		{Name: "processor", Type: field.TypeString, Nullable: true},
		{Name: "processor_cores", Type: field.TypeInt64, Nullable: true},
		{Name: "processor_arch", Type: field.TypeString, Nullable: true},
		{Name: "agent_computer", Type: field.TypeString, Unique: true},
	}
	// ComputersTable holds the schema information for the "computers" table.
	ComputersTable = &schema.Table{
		Name:       "computers",
		Columns:    ComputersColumns,
		PrimaryKey: []*schema.Column{ComputersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "computers_agents_computer",
				Columns:    []*schema.Column{ComputersColumns[8]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DeploymentsColumns holds the columns for the "deployments" table.
	DeploymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "package_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "installed", Type: field.TypeTime, Nullable: true},
		{Name: "updated", Type: field.TypeTime, Nullable: true},
		{Name: "agent_deployments", Type: field.TypeString},
	}
	// DeploymentsTable holds the schema information for the "deployments" table.
	DeploymentsTable = &schema.Table{
		Name:       "deployments",
		Columns:    DeploymentsColumns,
		PrimaryKey: []*schema.Column{DeploymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployments_agents_deployments",
				Columns:    []*schema.Column{DeploymentsColumns[6]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LogicalDisksColumns holds the columns for the "logical_disks" table.
	LogicalDisksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "label", Type: field.TypeString},
		{Name: "filesystem", Type: field.TypeString, Nullable: true},
		{Name: "usage", Type: field.TypeInt8, Default: 0},
		{Name: "size_in_units", Type: field.TypeString, Nullable: true},
		{Name: "remaining_space_in_units", Type: field.TypeString, Nullable: true},
		{Name: "volume_name", Type: field.TypeString, Nullable: true},
		{Name: "bitlocker_status", Type: field.TypeString, Nullable: true},
		{Name: "agent_logicaldisks", Type: field.TypeString},
	}
	// LogicalDisksTable holds the schema information for the "logical_disks" table.
	LogicalDisksTable = &schema.Table{
		Name:       "logical_disks",
		Columns:    LogicalDisksColumns,
		PrimaryKey: []*schema.Column{LogicalDisksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "logical_disks_agents_logicaldisks",
				Columns:    []*schema.Column{LogicalDisksColumns[8]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MetadataColumns holds the columns for the "metadata" table.
	MetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString},
		{Name: "agent_metadata", Type: field.TypeString},
		{Name: "org_metadata_metadata", Type: field.TypeInt},
	}
	// MetadataTable holds the schema information for the "metadata" table.
	MetadataTable = &schema.Table{
		Name:       "metadata",
		Columns:    MetadataColumns,
		PrimaryKey: []*schema.Column{MetadataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "metadata_agents_metadata",
				Columns:    []*schema.Column{MetadataColumns[2]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "metadata_org_metadata_metadata",
				Columns:    []*schema.Column{MetadataColumns[3]},
				RefColumns: []*schema.Column{OrgMetadataColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "metadata_org_metadata_metadata_agent_metadata",
				Unique:  true,
				Columns: []*schema.Column{MetadataColumns[3], MetadataColumns[2]},
			},
		},
	}
	// MonitorsColumns holds the columns for the "monitors" table.
	MonitorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "manufacturer", Type: field.TypeString, Nullable: true},
		{Name: "model", Type: field.TypeString, Nullable: true},
		{Name: "serial", Type: field.TypeString, Nullable: true},
		{Name: "agent_monitors", Type: field.TypeString},
	}
	// MonitorsTable holds the schema information for the "monitors" table.
	MonitorsTable = &schema.Table{
		Name:       "monitors",
		Columns:    MonitorsColumns,
		PrimaryKey: []*schema.Column{MonitorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "monitors_agents_monitors",
				Columns:    []*schema.Column{MonitorsColumns[4]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NetworkAdaptersColumns holds the columns for the "network_adapters" table.
	NetworkAdaptersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "mac_address", Type: field.TypeString},
		{Name: "addresses", Type: field.TypeString},
		{Name: "subnet", Type: field.TypeString, Nullable: true},
		{Name: "default_gateway", Type: field.TypeString, Nullable: true},
		{Name: "dns_servers", Type: field.TypeString, Nullable: true},
		{Name: "dns_domain", Type: field.TypeString, Nullable: true},
		{Name: "dhcp_enabled", Type: field.TypeBool, Nullable: true},
		{Name: "dhcp_lease_obtained", Type: field.TypeTime, Nullable: true},
		{Name: "dhcp_lease_expired", Type: field.TypeTime, Nullable: true},
		{Name: "speed", Type: field.TypeString},
		{Name: "agent_networkadapters", Type: field.TypeString},
	}
	// NetworkAdaptersTable holds the schema information for the "network_adapters" table.
	NetworkAdaptersTable = &schema.Table{
		Name:       "network_adapters",
		Columns:    NetworkAdaptersColumns,
		PrimaryKey: []*schema.Column{NetworkAdaptersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "network_adapters_agents_networkadapters",
				Columns:    []*schema.Column{NetworkAdaptersColumns[12]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OperatingSystemsColumns holds the columns for the "operating_systems" table.
	OperatingSystemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "version", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "edition", Type: field.TypeString, Nullable: true},
		{Name: "install_date", Type: field.TypeTime, Nullable: true},
		{Name: "arch", Type: field.TypeString, Nullable: true},
		{Name: "username", Type: field.TypeString},
		{Name: "last_bootup_time", Type: field.TypeTime, Nullable: true},
		{Name: "agent_operatingsystem", Type: field.TypeString, Unique: true},
	}
	// OperatingSystemsTable holds the schema information for the "operating_systems" table.
	OperatingSystemsTable = &schema.Table{
		Name:       "operating_systems",
		Columns:    OperatingSystemsColumns,
		PrimaryKey: []*schema.Column{OperatingSystemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "operating_systems_agents_operatingsystem",
				Columns:    []*schema.Column{OperatingSystemsColumns[9]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OrgMetadataColumns holds the columns for the "org_metadata" table.
	OrgMetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// OrgMetadataTable holds the schema information for the "org_metadata" table.
	OrgMetadataTable = &schema.Table{
		Name:       "org_metadata",
		Columns:    OrgMetadataColumns,
		PrimaryKey: []*schema.Column{OrgMetadataColumns[0]},
	}
	// PrintersColumns holds the columns for the "printers" table.
	PrintersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "port", Type: field.TypeString, Nullable: true},
		{Name: "is_default", Type: field.TypeBool, Nullable: true},
		{Name: "is_network", Type: field.TypeBool, Nullable: true},
		{Name: "agent_printers", Type: field.TypeString},
	}
	// PrintersTable holds the schema information for the "printers" table.
	PrintersTable = &schema.Table{
		Name:       "printers",
		Columns:    PrintersColumns,
		PrimaryKey: []*schema.Column{PrintersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "printers_agents_printers",
				Columns:    []*schema.Column{PrintersColumns[5]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ReleasesColumns holds the columns for the "releases" table.
	ReleasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "channel", Type: field.TypeString, Nullable: true},
		{Name: "summary", Type: field.TypeString, Nullable: true},
		{Name: "release_notes", Type: field.TypeString, Nullable: true},
		{Name: "file_url", Type: field.TypeString, Nullable: true},
		{Name: "checksum", Type: field.TypeString, Nullable: true},
		{Name: "is_critical", Type: field.TypeBool, Nullable: true},
		{Name: "release_date", Type: field.TypeTime, Nullable: true},
		{Name: "os", Type: field.TypeString, Nullable: true},
		{Name: "arch", Type: field.TypeString, Nullable: true},
	}
	// ReleasesTable holds the schema information for the "releases" table.
	ReleasesTable = &schema.Table{
		Name:       "releases",
		Columns:    ReleasesColumns,
		PrimaryKey: []*schema.Column{ReleasesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "release_version_channel_os_arch",
				Unique:  true,
				Columns: []*schema.Column{ReleasesColumns[1], ReleasesColumns[2], ReleasesColumns[9], ReleasesColumns[10]},
			},
		},
	}
	// RevocationsColumns holds the columns for the "revocations" table.
	RevocationsColumns = []*schema.Column{
		{Name: "serial", Type: field.TypeInt64, Increment: true},
		{Name: "reason", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "info", Type: field.TypeString, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, Nullable: true},
		{Name: "revoked", Type: field.TypeTime},
	}
	// RevocationsTable holds the schema information for the "revocations" table.
	RevocationsTable = &schema.Table{
		Name:       "revocations",
		Columns:    RevocationsColumns,
		PrimaryKey: []*schema.Column{RevocationsColumns[0]},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "token", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "data", Type: field.TypeBytes},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "user_sessions", Type: field.TypeString, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sessions_expiry_idx",
				Unique:  false,
				Columns: []*schema.Column{SessionsColumns[2]},
			},
		},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "language", Type: field.TypeString, Nullable: true},
		{Name: "organization", Type: field.TypeString, Nullable: true},
		{Name: "postal_address", Type: field.TypeString, Nullable: true},
		{Name: "postal_code", Type: field.TypeString, Nullable: true},
		{Name: "locality", Type: field.TypeString, Nullable: true},
		{Name: "province", Type: field.TypeString, Nullable: true},
		{Name: "state", Type: field.TypeString, Nullable: true},
		{Name: "country", Type: field.TypeString, Nullable: true, Default: "ES"},
		{Name: "smtp_server", Type: field.TypeString, Nullable: true},
		{Name: "smtp_port", Type: field.TypeInt, Nullable: true, Default: 587},
		{Name: "smtp_user", Type: field.TypeString, Nullable: true},
		{Name: "smtp_password", Type: field.TypeString, Nullable: true},
		{Name: "smtp_auth", Type: field.TypeString, Nullable: true, Default: "LOGIN"},
		{Name: "smtp_tls", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "smtp_starttls", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "nats_server", Type: field.TypeString, Nullable: true},
		{Name: "nats_port", Type: field.TypeString, Nullable: true},
		{Name: "message_from", Type: field.TypeString, Nullable: true},
		{Name: "max_upload_size", Type: field.TypeString, Nullable: true, Default: "512M"},
		{Name: "user_cert_years_valid", Type: field.TypeInt, Nullable: true, Default: 1},
		{Name: "nats_request_timeout_seconds", Type: field.TypeInt, Nullable: true, Default: 20},
		{Name: "refresh_time_in_minutes", Type: field.TypeInt, Nullable: true, Default: 5},
		{Name: "session_lifetime_in_minutes", Type: field.TypeInt, Nullable: true, Default: 1440},
		{Name: "update_channel", Type: field.TypeString, Nullable: true, Default: "stable"},
		{Name: "created", Type: field.TypeTime, Nullable: true},
		{Name: "modified", Type: field.TypeTime, Nullable: true},
		{Name: "agent_report_frequence_in_minutes", Type: field.TypeInt, Nullable: true, Default: 60},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
	}
	// SharesColumns holds the columns for the "shares" table.
	SharesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "path", Type: field.TypeString, Nullable: true},
		{Name: "agent_shares", Type: field.TypeString},
	}
	// SharesTable holds the schema information for the "shares" table.
	SharesTable = &schema.Table{
		Name:       "shares",
		Columns:    SharesColumns,
		PrimaryKey: []*schema.Column{SharesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "shares_agents_shares",
				Columns:    []*schema.Column{SharesColumns[4]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SystemUpdatesColumns holds the columns for the "system_updates" table.
	SystemUpdatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString},
		{Name: "last_install", Type: field.TypeTime},
		{Name: "last_search", Type: field.TypeTime},
		{Name: "pending_updates", Type: field.TypeBool},
		{Name: "agent_systemupdate", Type: field.TypeString, Unique: true},
	}
	// SystemUpdatesTable holds the schema information for the "system_updates" table.
	SystemUpdatesTable = &schema.Table{
		Name:       "system_updates",
		Columns:    SystemUpdatesColumns,
		PrimaryKey: []*schema.Column{SystemUpdatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "system_updates_agents_systemupdate",
				Columns:    []*schema.Column{SystemUpdatesColumns[5]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tag", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "color", Type: field.TypeString},
		{Name: "tag_children", Type: field.TypeInt, Nullable: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tags_tags_children",
				Columns:    []*schema.Column{TagsColumns[4]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UpdatesColumns holds the columns for the "updates" table.
	UpdatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "support_url", Type: field.TypeString, Nullable: true},
		{Name: "agent_updates", Type: field.TypeString},
	}
	// UpdatesTable holds the schema information for the "updates" table.
	UpdatesTable = &schema.Table{
		Name:       "updates",
		Columns:    UpdatesColumns,
		PrimaryKey: []*schema.Column{UpdatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "updates_agents_updates",
				Columns:    []*schema.Column{UpdatesColumns[4]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "uid", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "country", Type: field.TypeString, Nullable: true},
		{Name: "email_verified", Type: field.TypeBool, Default: false},
		{Name: "register", Type: field.TypeString, Default: "users.pending_email_confirmation"},
		{Name: "cert_clear_password", Type: field.TypeString, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, Nullable: true},
		{Name: "created", Type: field.TypeTime, Nullable: true},
		{Name: "modified", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "users_email_idx",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[2]},
			},
		},
	}
	// AgentTagsColumns holds the columns for the "agent_tags" table.
	AgentTagsColumns = []*schema.Column{
		{Name: "agent_id", Type: field.TypeString},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// AgentTagsTable holds the schema information for the "agent_tags" table.
	AgentTagsTable = &schema.Table{
		Name:       "agent_tags",
		Columns:    AgentTagsColumns,
		PrimaryKey: []*schema.Column{AgentTagsColumns[0], AgentTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "agent_tags_agent_id",
				Columns:    []*schema.Column{AgentTagsColumns[0]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "agent_tags_tag_id",
				Columns:    []*schema.Column{AgentTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgentsTable,
		AntiviriTable,
		AppsTable,
		CertificatesTable,
		ComputersTable,
		DeploymentsTable,
		LogicalDisksTable,
		MetadataTable,
		MonitorsTable,
		NetworkAdaptersTable,
		OperatingSystemsTable,
		OrgMetadataTable,
		PrintersTable,
		ReleasesTable,
		RevocationsTable,
		SessionsTable,
		SettingsTable,
		SharesTable,
		SystemUpdatesTable,
		TagsTable,
		UpdatesTable,
		UsersTable,
		AgentTagsTable,
	}
)

func init() {
	AgentsTable.ForeignKeys[0].RefTable = ReleasesTable
	AntiviriTable.ForeignKeys[0].RefTable = AgentsTable
	AppsTable.ForeignKeys[0].RefTable = AgentsTable
	ComputersTable.ForeignKeys[0].RefTable = AgentsTable
	DeploymentsTable.ForeignKeys[0].RefTable = AgentsTable
	LogicalDisksTable.ForeignKeys[0].RefTable = AgentsTable
	MetadataTable.ForeignKeys[0].RefTable = AgentsTable
	MetadataTable.ForeignKeys[1].RefTable = OrgMetadataTable
	MonitorsTable.ForeignKeys[0].RefTable = AgentsTable
	NetworkAdaptersTable.ForeignKeys[0].RefTable = AgentsTable
	OperatingSystemsTable.ForeignKeys[0].RefTable = AgentsTable
	PrintersTable.ForeignKeys[0].RefTable = AgentsTable
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
	SharesTable.ForeignKeys[0].RefTable = AgentsTable
	SystemUpdatesTable.ForeignKeys[0].RefTable = AgentsTable
	TagsTable.ForeignKeys[0].RefTable = TagsTable
	UpdatesTable.ForeignKeys[0].RefTable = AgentsTable
	AgentTagsTable.ForeignKeys[0].RefTable = AgentsTable
	AgentTagsTable.ForeignKeys[1].RefTable = TagsTable
}
