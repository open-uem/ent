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
		{Name: "version", Type: field.TypeString},
		{Name: "ip", Type: field.TypeString, Default: ""},
		{Name: "first_contact", Type: field.TypeTime},
		{Name: "last_contact", Type: field.TypeTime},
		{Name: "enabled", Type: field.TypeBool, Default: true},
	}
	// AgentsTable holds the schema information for the "agents" table.
	AgentsTable = &schema.Table{
		Name:       "agents",
		Columns:    AgentsColumns,
		PrimaryKey: []*schema.Column{AgentsColumns[0]},
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
				Columns:    []*schema.Column{OperatingSystemsColumns[8]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
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
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "data", Type: field.TypeBytes},
		{Name: "expiry", Type: field.TypeTime},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sessions_token",
				Unique:  false,
				Columns: []*schema.Column{SessionsColumns[1]},
			},
			{
				Name:    "sessions_expiry_idx",
				Unique:  false,
				Columns: []*schema.Column{SessionsColumns[3]},
			},
		},
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgentsTable,
		AntiviriTable,
		AppsTable,
		ComputersTable,
		LogicalDisksTable,
		MonitorsTable,
		NetworkAdaptersTable,
		OperatingSystemsTable,
		PrintersTable,
		SessionsTable,
		SharesTable,
		SystemUpdatesTable,
	}
)

func init() {
	AntiviriTable.ForeignKeys[0].RefTable = AgentsTable
	AppsTable.ForeignKeys[0].RefTable = AgentsTable
	ComputersTable.ForeignKeys[0].RefTable = AgentsTable
	LogicalDisksTable.ForeignKeys[0].RefTable = AgentsTable
	MonitorsTable.ForeignKeys[0].RefTable = AgentsTable
	NetworkAdaptersTable.ForeignKeys[0].RefTable = AgentsTable
	OperatingSystemsTable.ForeignKeys[0].RefTable = AgentsTable
	PrintersTable.ForeignKeys[0].RefTable = AgentsTable
	SharesTable.ForeignKeys[0].RefTable = AgentsTable
	SystemUpdatesTable.ForeignKeys[0].RefTable = AgentsTable
}
