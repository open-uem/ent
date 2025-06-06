// Code generated by ent, DO NOT EDIT.

package settings

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the settings type in the database.
	Label = "settings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldOrganization holds the string denoting the organization field in the database.
	FieldOrganization = "organization"
	// FieldPostalAddress holds the string denoting the postal_address field in the database.
	FieldPostalAddress = "postal_address"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldLocality holds the string denoting the locality field in the database.
	FieldLocality = "locality"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldSMTPServer holds the string denoting the smtp_server field in the database.
	FieldSMTPServer = "smtp_server"
	// FieldSMTPPort holds the string denoting the smtp_port field in the database.
	FieldSMTPPort = "smtp_port"
	// FieldSMTPUser holds the string denoting the smtp_user field in the database.
	FieldSMTPUser = "smtp_user"
	// FieldSMTPPassword holds the string denoting the smtp_password field in the database.
	FieldSMTPPassword = "smtp_password"
	// FieldSMTPAuth holds the string denoting the smtp_auth field in the database.
	FieldSMTPAuth = "smtp_auth"
	// FieldSMTPTLS holds the string denoting the smtp_tls field in the database.
	FieldSMTPTLS = "smtp_tls"
	// FieldSMTPStarttls holds the string denoting the smtp_starttls field in the database.
	FieldSMTPStarttls = "smtp_starttls"
	// FieldNatsServer holds the string denoting the nats_server field in the database.
	FieldNatsServer = "nats_server"
	// FieldNatsPort holds the string denoting the nats_port field in the database.
	FieldNatsPort = "nats_port"
	// FieldMessageFrom holds the string denoting the message_from field in the database.
	FieldMessageFrom = "message_from"
	// FieldMaxUploadSize holds the string denoting the max_upload_size field in the database.
	FieldMaxUploadSize = "max_upload_size"
	// FieldUserCertYearsValid holds the string denoting the user_cert_years_valid field in the database.
	FieldUserCertYearsValid = "user_cert_years_valid"
	// FieldNatsRequestTimeoutSeconds holds the string denoting the nats_request_timeout_seconds field in the database.
	FieldNatsRequestTimeoutSeconds = "nats_request_timeout_seconds"
	// FieldRefreshTimeInMinutes holds the string denoting the refresh_time_in_minutes field in the database.
	FieldRefreshTimeInMinutes = "refresh_time_in_minutes"
	// FieldSessionLifetimeInMinutes holds the string denoting the session_lifetime_in_minutes field in the database.
	FieldSessionLifetimeInMinutes = "session_lifetime_in_minutes"
	// FieldUpdateChannel holds the string denoting the update_channel field in the database.
	FieldUpdateChannel = "update_channel"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// FieldAgentReportFrequenceInMinutes holds the string denoting the agent_report_frequence_in_minutes field in the database.
	FieldAgentReportFrequenceInMinutes = "agent_report_frequence_in_minutes"
	// FieldRequestVncPin holds the string denoting the request_vnc_pin field in the database.
	FieldRequestVncPin = "request_vnc_pin"
	// FieldProfilesApplicationFrequenceInMinutes holds the string denoting the profiles_application_frequence_in_minutes field in the database.
	FieldProfilesApplicationFrequenceInMinutes = "profiles_application_frequence_in_minutes"
	// FieldUseWinget holds the string denoting the use_winget field in the database.
	FieldUseWinget = "use_winget"
	// FieldUseFlatpak holds the string denoting the use_flatpak field in the database.
	FieldUseFlatpak = "use_flatpak"
	// FieldUseBrew holds the string denoting the use_brew field in the database.
	FieldUseBrew = "use_brew"
	// FieldDisableSftp holds the string denoting the disable_sftp field in the database.
	FieldDisableSftp = "disable_sftp"
	// FieldDisableRemoteAssistance holds the string denoting the disable_remote_assistance field in the database.
	FieldDisableRemoteAssistance = "disable_remote_assistance"
	// FieldDetectRemoteAgents holds the string denoting the detect_remote_agents field in the database.
	FieldDetectRemoteAgents = "detect_remote_agents"
	// FieldAutoAdmitAgents holds the string denoting the auto_admit_agents field in the database.
	FieldAutoAdmitAgents = "auto_admit_agents"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// Table holds the table name of the settings in the database.
	Table = "settings"
	// TagTable is the table that holds the tag relation/edge.
	TagTable = "settings"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "settings_tag"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "settings"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "tenant_settings"
)

// Columns holds all SQL columns for settings fields.
var Columns = []string{
	FieldID,
	FieldLanguage,
	FieldOrganization,
	FieldPostalAddress,
	FieldPostalCode,
	FieldLocality,
	FieldProvince,
	FieldState,
	FieldCountry,
	FieldSMTPServer,
	FieldSMTPPort,
	FieldSMTPUser,
	FieldSMTPPassword,
	FieldSMTPAuth,
	FieldSMTPTLS,
	FieldSMTPStarttls,
	FieldNatsServer,
	FieldNatsPort,
	FieldMessageFrom,
	FieldMaxUploadSize,
	FieldUserCertYearsValid,
	FieldNatsRequestTimeoutSeconds,
	FieldRefreshTimeInMinutes,
	FieldSessionLifetimeInMinutes,
	FieldUpdateChannel,
	FieldCreated,
	FieldModified,
	FieldAgentReportFrequenceInMinutes,
	FieldRequestVncPin,
	FieldProfilesApplicationFrequenceInMinutes,
	FieldUseWinget,
	FieldUseFlatpak,
	FieldUseBrew,
	FieldDisableSftp,
	FieldDisableRemoteAssistance,
	FieldDetectRemoteAgents,
	FieldAutoAdmitAgents,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"settings_tag",
	"tenant_settings",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCountry holds the default value on creation for the "country" field.
	DefaultCountry string
	// DefaultSMTPPort holds the default value on creation for the "smtp_port" field.
	DefaultSMTPPort int
	// DefaultSMTPAuth holds the default value on creation for the "smtp_auth" field.
	DefaultSMTPAuth string
	// DefaultSMTPTLS holds the default value on creation for the "smtp_tls" field.
	DefaultSMTPTLS bool
	// DefaultSMTPStarttls holds the default value on creation for the "smtp_starttls" field.
	DefaultSMTPStarttls bool
	// DefaultMaxUploadSize holds the default value on creation for the "max_upload_size" field.
	DefaultMaxUploadSize string
	// DefaultUserCertYearsValid holds the default value on creation for the "user_cert_years_valid" field.
	DefaultUserCertYearsValid int
	// DefaultNatsRequestTimeoutSeconds holds the default value on creation for the "nats_request_timeout_seconds" field.
	DefaultNatsRequestTimeoutSeconds int
	// DefaultRefreshTimeInMinutes holds the default value on creation for the "refresh_time_in_minutes" field.
	DefaultRefreshTimeInMinutes int
	// DefaultSessionLifetimeInMinutes holds the default value on creation for the "session_lifetime_in_minutes" field.
	DefaultSessionLifetimeInMinutes int
	// DefaultUpdateChannel holds the default value on creation for the "update_channel" field.
	DefaultUpdateChannel string
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// DefaultModified holds the default value on creation for the "modified" field.
	DefaultModified func() time.Time
	// UpdateDefaultModified holds the default value on update for the "modified" field.
	UpdateDefaultModified func() time.Time
	// DefaultAgentReportFrequenceInMinutes holds the default value on creation for the "agent_report_frequence_in_minutes" field.
	DefaultAgentReportFrequenceInMinutes int
	// DefaultRequestVncPin holds the default value on creation for the "request_vnc_pin" field.
	DefaultRequestVncPin bool
	// DefaultProfilesApplicationFrequenceInMinutes holds the default value on creation for the "profiles_application_frequence_in_minutes" field.
	DefaultProfilesApplicationFrequenceInMinutes int
	// DefaultUseWinget holds the default value on creation for the "use_winget" field.
	DefaultUseWinget bool
	// DefaultUseFlatpak holds the default value on creation for the "use_flatpak" field.
	DefaultUseFlatpak bool
	// DefaultUseBrew holds the default value on creation for the "use_brew" field.
	DefaultUseBrew bool
	// DefaultDisableSftp holds the default value on creation for the "disable_sftp" field.
	DefaultDisableSftp bool
	// DefaultDisableRemoteAssistance holds the default value on creation for the "disable_remote_assistance" field.
	DefaultDisableRemoteAssistance bool
	// DefaultDetectRemoteAgents holds the default value on creation for the "detect_remote_agents" field.
	DefaultDetectRemoteAgents bool
	// DefaultAutoAdmitAgents holds the default value on creation for the "auto_admit_agents" field.
	DefaultAutoAdmitAgents bool
)

// OrderOption defines the ordering options for the Settings queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByOrganization orders the results by the organization field.
func ByOrganization(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganization, opts...).ToFunc()
}

// ByPostalAddress orders the results by the postal_address field.
func ByPostalAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalAddress, opts...).ToFunc()
}

// ByPostalCode orders the results by the postal_code field.
func ByPostalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCode, opts...).ToFunc()
}

// ByLocality orders the results by the locality field.
func ByLocality(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocality, opts...).ToFunc()
}

// ByProvince orders the results by the province field.
func ByProvince(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProvince, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// BySMTPServer orders the results by the smtp_server field.
func BySMTPServer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSMTPServer, opts...).ToFunc()
}

// BySMTPPort orders the results by the smtp_port field.
func BySMTPPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSMTPPort, opts...).ToFunc()
}

// BySMTPUser orders the results by the smtp_user field.
func BySMTPUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSMTPUser, opts...).ToFunc()
}

// BySMTPPassword orders the results by the smtp_password field.
func BySMTPPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSMTPPassword, opts...).ToFunc()
}

// BySMTPAuth orders the results by the smtp_auth field.
func BySMTPAuth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSMTPAuth, opts...).ToFunc()
}

// BySMTPTLS orders the results by the smtp_tls field.
func BySMTPTLS(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSMTPTLS, opts...).ToFunc()
}

// BySMTPStarttls orders the results by the smtp_starttls field.
func BySMTPStarttls(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSMTPStarttls, opts...).ToFunc()
}

// ByNatsServer orders the results by the nats_server field.
func ByNatsServer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNatsServer, opts...).ToFunc()
}

// ByNatsPort orders the results by the nats_port field.
func ByNatsPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNatsPort, opts...).ToFunc()
}

// ByMessageFrom orders the results by the message_from field.
func ByMessageFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessageFrom, opts...).ToFunc()
}

// ByMaxUploadSize orders the results by the max_upload_size field.
func ByMaxUploadSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxUploadSize, opts...).ToFunc()
}

// ByUserCertYearsValid orders the results by the user_cert_years_valid field.
func ByUserCertYearsValid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserCertYearsValid, opts...).ToFunc()
}

// ByNatsRequestTimeoutSeconds orders the results by the nats_request_timeout_seconds field.
func ByNatsRequestTimeoutSeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNatsRequestTimeoutSeconds, opts...).ToFunc()
}

// ByRefreshTimeInMinutes orders the results by the refresh_time_in_minutes field.
func ByRefreshTimeInMinutes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshTimeInMinutes, opts...).ToFunc()
}

// BySessionLifetimeInMinutes orders the results by the session_lifetime_in_minutes field.
func BySessionLifetimeInMinutes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSessionLifetimeInMinutes, opts...).ToFunc()
}

// ByUpdateChannel orders the results by the update_channel field.
func ByUpdateChannel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateChannel, opts...).ToFunc()
}

// ByCreated orders the results by the created field.
func ByCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreated, opts...).ToFunc()
}

// ByModified orders the results by the modified field.
func ByModified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModified, opts...).ToFunc()
}

// ByAgentReportFrequenceInMinutes orders the results by the agent_report_frequence_in_minutes field.
func ByAgentReportFrequenceInMinutes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAgentReportFrequenceInMinutes, opts...).ToFunc()
}

// ByRequestVncPin orders the results by the request_vnc_pin field.
func ByRequestVncPin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestVncPin, opts...).ToFunc()
}

// ByProfilesApplicationFrequenceInMinutes orders the results by the profiles_application_frequence_in_minutes field.
func ByProfilesApplicationFrequenceInMinutes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfilesApplicationFrequenceInMinutes, opts...).ToFunc()
}

// ByUseWinget orders the results by the use_winget field.
func ByUseWinget(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUseWinget, opts...).ToFunc()
}

// ByUseFlatpak orders the results by the use_flatpak field.
func ByUseFlatpak(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUseFlatpak, opts...).ToFunc()
}

// ByUseBrew orders the results by the use_brew field.
func ByUseBrew(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUseBrew, opts...).ToFunc()
}

// ByDisableSftp orders the results by the disable_sftp field.
func ByDisableSftp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisableSftp, opts...).ToFunc()
}

// ByDisableRemoteAssistance orders the results by the disable_remote_assistance field.
func ByDisableRemoteAssistance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisableRemoteAssistance, opts...).ToFunc()
}

// ByDetectRemoteAgents orders the results by the detect_remote_agents field.
func ByDetectRemoteAgents(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDetectRemoteAgents, opts...).ToFunc()
}

// ByAutoAdmitAgents orders the results by the auto_admit_agents field.
func ByAutoAdmitAgents(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAutoAdmitAgents, opts...).ToFunc()
}

// ByTagField orders the results by tag field.
func ByTagField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagStep(), sql.OrderByField(field, opts...))
	}
}

// ByTenantField orders the results by tenant field.
func ByTenantField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTenantStep(), sql.OrderByField(field, opts...))
	}
}
func newTagStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
	)
}
func newTenantStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TenantInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, TenantTable, TenantColumn),
	)
}
