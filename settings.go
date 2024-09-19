// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/doncicuto/openuem_ent/settings"
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
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified     time.Time `json:"modified,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Settings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case settings.FieldSMTPTLS:
			values[i] = new(sql.NullBool)
		case settings.FieldID, settings.FieldSMTPPort:
			values[i] = new(sql.NullInt64)
		case settings.FieldLanguage, settings.FieldOrganization, settings.FieldPostalAddress, settings.FieldPostalCode, settings.FieldLocality, settings.FieldProvince, settings.FieldState, settings.FieldCountry, settings.FieldSMTPServer, settings.FieldSMTPUser, settings.FieldSMTPPassword, settings.FieldSMTPAuth:
			values[i] = new(sql.NullString)
		case settings.FieldCreated, settings.FieldModified:
			values[i] = new(sql.NullTime)
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
		panic("openuem_ent: Settings is not a transactional entity")
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
	builder.WriteString("created=")
	builder.WriteString(s.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(s.Modified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SettingsSlice is a parsable slice of Settings.
type SettingsSlice []*Settings
