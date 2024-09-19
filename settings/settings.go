// Code generated by ent, DO NOT EDIT.

package settings

import (
	"time"

	"entgo.io/ent/dialect/sql"
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
	// FieldMessageFrom holds the string denoting the message_from field in the database.
	FieldMessageFrom = "message_from"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// Table holds the table name of the settings in the database.
	Table = "settings"
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
	FieldMessageFrom,
	FieldCreated,
	FieldModified,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSMTPPort holds the default value on creation for the "smtp_port" field.
	DefaultSMTPPort int
	// DefaultSMTPAuth holds the default value on creation for the "smtp_auth" field.
	DefaultSMTPAuth string
	// DefaultSMTPTLS holds the default value on creation for the "smtp_tls" field.
	DefaultSMTPTLS bool
	// DefaultSMTPStarttls holds the default value on creation for the "smtp_starttls" field.
	DefaultSMTPStarttls bool
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// DefaultModified holds the default value on creation for the "modified" field.
	DefaultModified func() time.Time
	// UpdateDefaultModified holds the default value on update for the "modified" field.
	UpdateDefaultModified func() time.Time
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

// ByMessageFrom orders the results by the message_from field.
func ByMessageFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessageFrom, opts...).ToFunc()
}

// ByCreated orders the results by the created field.
func ByCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreated, opts...).ToFunc()
}

// ByModified orders the results by the modified field.
func ByModified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModified, opts...).ToFunc()
}
