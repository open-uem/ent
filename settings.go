// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/settings"
	"github.com/open-uem/ent/tag"
)

// Settings is the model entity for the Settings schema.
type Settings struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// Organization holds the value of the "organization" field.
	Organization string `json:"organization,omitempty"`
	// PostalAddress holds the value of the "postal_address" field.
	PostalAddress string `json:"postal_address,omitempty"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode string `json:"postal_code,omitempty"`
	// Locality holds the value of the "locality" field.
	Locality string `json:"locality,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// SMTPServer holds the value of the "smtp_server" field.
	SMTPServer string `json:"smtp_server,omitempty"`
	// SMTPPort holds the value of the "smtp_port" field.
	SMTPPort int `json:"smtp_port,omitempty"`
	// SMTPUser holds the value of the "smtp_user" field.
	SMTPUser string `json:"smtp_user,omitempty"`
	// SMTPPassword holds the value of the "smtp_password" field.
	SMTPPassword string `json:"smtp_password,omitempty"`
	// SMTPAuth holds the value of the "smtp_auth" field.
	SMTPAuth string `json:"smtp_auth,omitempty"`
	// SMTPTLS holds the value of the "smtp_tls" field.
	SMTPTLS bool `json:"smtp_tls,omitempty"`
	// SMTPStarttls holds the value of the "smtp_starttls" field.
	SMTPStarttls bool `json:"smtp_starttls,omitempty"`
	// NatsServer holds the value of the "nats_server" field.
	NatsServer string `json:"nats_server,omitempty"`
	// NatsPort holds the value of the "nats_port" field.
	NatsPort string `json:"nats_port,omitempty"`
	// MessageFrom holds the value of the "message_from" field.
	MessageFrom string `json:"message_from,omitempty"`
	// MaxUploadSize holds the value of the "max_upload_size" field.
	MaxUploadSize string `json:"max_upload_size,omitempty"`
	// UserCertYearsValid holds the value of the "user_cert_years_valid" field.
	UserCertYearsValid int `json:"user_cert_years_valid,omitempty"`
	// NatsRequestTimeoutSeconds holds the value of the "nats_request_timeout_seconds" field.
	NatsRequestTimeoutSeconds int `json:"nats_request_timeout_seconds,omitempty"`
	// RefreshTimeInMinutes holds the value of the "refresh_time_in_minutes" field.
	RefreshTimeInMinutes int `json:"refresh_time_in_minutes,omitempty"`
	// SessionLifetimeInMinutes holds the value of the "session_lifetime_in_minutes" field.
	SessionLifetimeInMinutes int `json:"session_lifetime_in_minutes,omitempty"`
	// UpdateChannel holds the value of the "update_channel" field.
	UpdateChannel string `json:"update_channel,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
	// AgentReportFrequenceInMinutes holds the value of the "agent_report_frequence_in_minutes" field.
	AgentReportFrequenceInMinutes int `json:"agent_report_frequence_in_minutes,omitempty"`
	// RequestVncPin holds the value of the "request_vnc_pin" field.
	RequestVncPin bool `json:"request_vnc_pin,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SettingsQuery when eager-loading is set.
	Edges        SettingsEdges `json:"edges"`
	settings_tag *int
	selectValues sql.SelectValues
}

// SettingsEdges holds the relations/edges for other nodes in the graph.
type SettingsEdges struct {
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SettingsEdges) TagOrErr() (*Tag, error) {
	if e.Tag != nil {
		return e.Tag, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: tag.Label}
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Settings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case settings.FieldSMTPTLS, settings.FieldSMTPStarttls, settings.FieldRequestVncPin:
			values[i] = new(sql.NullBool)
		case settings.FieldID, settings.FieldSMTPPort, settings.FieldUserCertYearsValid, settings.FieldNatsRequestTimeoutSeconds, settings.FieldRefreshTimeInMinutes, settings.FieldSessionLifetimeInMinutes, settings.FieldAgentReportFrequenceInMinutes:
			values[i] = new(sql.NullInt64)
		case settings.FieldLanguage, settings.FieldOrganization, settings.FieldPostalAddress, settings.FieldPostalCode, settings.FieldLocality, settings.FieldProvince, settings.FieldState, settings.FieldCountry, settings.FieldSMTPServer, settings.FieldSMTPUser, settings.FieldSMTPPassword, settings.FieldSMTPAuth, settings.FieldNatsServer, settings.FieldNatsPort, settings.FieldMessageFrom, settings.FieldMaxUploadSize, settings.FieldUpdateChannel:
			values[i] = new(sql.NullString)
		case settings.FieldCreated, settings.FieldModified:
			values[i] = new(sql.NullTime)
		case settings.ForeignKeys[0]: // settings_tag
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Settings fields.
func (s *Settings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case settings.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case settings.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				s.Language = value.String
			}
		case settings.FieldOrganization:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization", values[i])
			} else if value.Valid {
				s.Organization = value.String
			}
		case settings.FieldPostalAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_address", values[i])
			} else if value.Valid {
				s.PostalAddress = value.String
			}
		case settings.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				s.PostalCode = value.String
			}
		case settings.FieldLocality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locality", values[i])
			} else if value.Valid {
				s.Locality = value.String
			}
		case settings.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				s.Province = value.String
			}
		case settings.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				s.State = value.String
			}
		case settings.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				s.Country = value.String
			}
		case settings.FieldSMTPServer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field smtp_server", values[i])
			} else if value.Valid {
				s.SMTPServer = value.String
			}
		case settings.FieldSMTPPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field smtp_port", values[i])
			} else if value.Valid {
				s.SMTPPort = int(value.Int64)
			}
		case settings.FieldSMTPUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field smtp_user", values[i])
			} else if value.Valid {
				s.SMTPUser = value.String
			}
		case settings.FieldSMTPPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field smtp_password", values[i])
			} else if value.Valid {
				s.SMTPPassword = value.String
			}
		case settings.FieldSMTPAuth:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field smtp_auth", values[i])
			} else if value.Valid {
				s.SMTPAuth = value.String
			}
		case settings.FieldSMTPTLS:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field smtp_tls", values[i])
			} else if value.Valid {
				s.SMTPTLS = value.Bool
			}
		case settings.FieldSMTPStarttls:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field smtp_starttls", values[i])
			} else if value.Valid {
				s.SMTPStarttls = value.Bool
			}
		case settings.FieldNatsServer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nats_server", values[i])
			} else if value.Valid {
				s.NatsServer = value.String
			}
		case settings.FieldNatsPort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nats_port", values[i])
			} else if value.Valid {
				s.NatsPort = value.String
			}
		case settings.FieldMessageFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message_from", values[i])
			} else if value.Valid {
				s.MessageFrom = value.String
			}
		case settings.FieldMaxUploadSize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field max_upload_size", values[i])
			} else if value.Valid {
				s.MaxUploadSize = value.String
			}
		case settings.FieldUserCertYearsValid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_cert_years_valid", values[i])
			} else if value.Valid {
				s.UserCertYearsValid = int(value.Int64)
			}
		case settings.FieldNatsRequestTimeoutSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nats_request_timeout_seconds", values[i])
			} else if value.Valid {
				s.NatsRequestTimeoutSeconds = int(value.Int64)
			}
		case settings.FieldRefreshTimeInMinutes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_time_in_minutes", values[i])
			} else if value.Valid {
				s.RefreshTimeInMinutes = int(value.Int64)
			}
		case settings.FieldSessionLifetimeInMinutes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field session_lifetime_in_minutes", values[i])
			} else if value.Valid {
				s.SessionLifetimeInMinutes = int(value.Int64)
			}
		case settings.FieldUpdateChannel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field update_channel", values[i])
			} else if value.Valid {
				s.UpdateChannel = value.String
			}
		case settings.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				s.Created = value.Time
			}
		case settings.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				s.Modified = value.Time
			}
		case settings.FieldAgentReportFrequenceInMinutes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field agent_report_frequence_in_minutes", values[i])
			} else if value.Valid {
				s.AgentReportFrequenceInMinutes = int(value.Int64)
			}
		case settings.FieldRequestVncPin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field request_vnc_pin", values[i])
			} else if value.Valid {
				s.RequestVncPin = value.Bool
			}
		case settings.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field settings_tag", value)
			} else if value.Valid {
				s.settings_tag = new(int)
				*s.settings_tag = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Settings.
// This includes values selected through modifiers, order, etc.
func (s *Settings) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryTag queries the "tag" edge of the Settings entity.
func (s *Settings) QueryTag() *TagQuery {
	return NewSettingsClient(s.config).QueryTag(s)
}

// Update returns a builder for updating this Settings.
// Note that you need to call Settings.Unwrap() before calling this method if this Settings
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Settings) Update() *SettingsUpdateOne {
	return NewSettingsClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Settings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Settings) Unwrap() *Settings {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Settings is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Settings) String() string {
	var builder strings.Builder
	builder.WriteString("Settings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("language=")
	builder.WriteString(s.Language)
	builder.WriteString(", ")
	builder.WriteString("organization=")
	builder.WriteString(s.Organization)
	builder.WriteString(", ")
	builder.WriteString("postal_address=")
	builder.WriteString(s.PostalAddress)
	builder.WriteString(", ")
	builder.WriteString("postal_code=")
	builder.WriteString(s.PostalCode)
	builder.WriteString(", ")
	builder.WriteString("locality=")
	builder.WriteString(s.Locality)
	builder.WriteString(", ")
	builder.WriteString("province=")
	builder.WriteString(s.Province)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(s.State)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(s.Country)
	builder.WriteString(", ")
	builder.WriteString("smtp_server=")
	builder.WriteString(s.SMTPServer)
	builder.WriteString(", ")
	builder.WriteString("smtp_port=")
	builder.WriteString(fmt.Sprintf("%v", s.SMTPPort))
	builder.WriteString(", ")
	builder.WriteString("smtp_user=")
	builder.WriteString(s.SMTPUser)
	builder.WriteString(", ")
	builder.WriteString("smtp_password=")
	builder.WriteString(s.SMTPPassword)
	builder.WriteString(", ")
	builder.WriteString("smtp_auth=")
	builder.WriteString(s.SMTPAuth)
	builder.WriteString(", ")
	builder.WriteString("smtp_tls=")
	builder.WriteString(fmt.Sprintf("%v", s.SMTPTLS))
	builder.WriteString(", ")
	builder.WriteString("smtp_starttls=")
	builder.WriteString(fmt.Sprintf("%v", s.SMTPStarttls))
	builder.WriteString(", ")
	builder.WriteString("nats_server=")
	builder.WriteString(s.NatsServer)
	builder.WriteString(", ")
	builder.WriteString("nats_port=")
	builder.WriteString(s.NatsPort)
	builder.WriteString(", ")
	builder.WriteString("message_from=")
	builder.WriteString(s.MessageFrom)
	builder.WriteString(", ")
	builder.WriteString("max_upload_size=")
	builder.WriteString(s.MaxUploadSize)
	builder.WriteString(", ")
	builder.WriteString("user_cert_years_valid=")
	builder.WriteString(fmt.Sprintf("%v", s.UserCertYearsValid))
	builder.WriteString(", ")
	builder.WriteString("nats_request_timeout_seconds=")
	builder.WriteString(fmt.Sprintf("%v", s.NatsRequestTimeoutSeconds))
	builder.WriteString(", ")
	builder.WriteString("refresh_time_in_minutes=")
	builder.WriteString(fmt.Sprintf("%v", s.RefreshTimeInMinutes))
	builder.WriteString(", ")
	builder.WriteString("session_lifetime_in_minutes=")
	builder.WriteString(fmt.Sprintf("%v", s.SessionLifetimeInMinutes))
	builder.WriteString(", ")
	builder.WriteString("update_channel=")
	builder.WriteString(s.UpdateChannel)
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(s.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(s.Modified.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("agent_report_frequence_in_minutes=")
	builder.WriteString(fmt.Sprintf("%v", s.AgentReportFrequenceInMinutes))
	builder.WriteString(", ")
	builder.WriteString("request_vnc_pin=")
	builder.WriteString(fmt.Sprintf("%v", s.RequestVncPin))
	builder.WriteByte(')')
	return builder.String()
}

// SettingsSlice is a parsable slice of Settings.
type SettingsSlice []*Settings
