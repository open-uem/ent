// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/site"
	"github.com/open-uem/ent/tenant"
)

// Site is the model entity for the Site schema.
type Site struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsDefault holds the value of the "is_default" field.
	IsDefault bool `json:"is_default,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SiteQuery when eager-loading is set.
	Edges        SiteEdges `json:"edges"`
	tenant_sites *int
	selectValues sql.SelectValues
}

// SiteEdges holds the relations/edges for other nodes in the graph.
type SiteEdges struct {
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// Agents holds the value of the agents edge.
	Agents []*Agent `json:"agents,omitempty"`
	// Profiles holds the value of the profiles edge.
	Profiles []*Profile `json:"profiles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SiteEdges) TenantOrErr() (*Tenant, error) {
	if e.Tenant != nil {
		return e.Tenant, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: tenant.Label}
	}
	return nil, &NotLoadedError{edge: "tenant"}
}

// AgentsOrErr returns the Agents value or an error if the edge
// was not loaded in eager-loading.
func (e SiteEdges) AgentsOrErr() ([]*Agent, error) {
	if e.loadedTypes[1] {
		return e.Agents, nil
	}
	return nil, &NotLoadedError{edge: "agents"}
}

// ProfilesOrErr returns the Profiles value or an error if the edge
// was not loaded in eager-loading.
func (e SiteEdges) ProfilesOrErr() ([]*Profile, error) {
	if e.loadedTypes[2] {
		return e.Profiles, nil
	}
	return nil, &NotLoadedError{edge: "profiles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Site) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case site.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case site.FieldID:
			values[i] = new(sql.NullInt64)
		case site.FieldDescription, site.FieldDomain:
			values[i] = new(sql.NullString)
		case site.FieldCreated, site.FieldModified:
			values[i] = new(sql.NullTime)
		case site.ForeignKeys[0]: // tenant_sites
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Site fields.
func (s *Site) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case site.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case site.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case site.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				s.IsDefault = value.Bool
			}
		case site.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				s.Domain = value.String
			}
		case site.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				s.Created = value.Time
			}
		case site.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				s.Modified = value.Time
			}
		case site.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field tenant_sites", value)
			} else if value.Valid {
				s.tenant_sites = new(int)
				*s.tenant_sites = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Site.
// This includes values selected through modifiers, order, etc.
func (s *Site) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryTenant queries the "tenant" edge of the Site entity.
func (s *Site) QueryTenant() *TenantQuery {
	return NewSiteClient(s.config).QueryTenant(s)
}

// QueryAgents queries the "agents" edge of the Site entity.
func (s *Site) QueryAgents() *AgentQuery {
	return NewSiteClient(s.config).QueryAgents(s)
}

// QueryProfiles queries the "profiles" edge of the Site entity.
func (s *Site) QueryProfiles() *ProfileQuery {
	return NewSiteClient(s.config).QueryProfiles(s)
}

// Update returns a builder for updating this Site.
// Note that you need to call Site.Unwrap() before calling this method if this Site
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Site) Update() *SiteUpdateOne {
	return NewSiteClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Site entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Site) Unwrap() *Site {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Site is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Site) String() string {
	var builder strings.Builder
	builder.WriteString("Site(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", s.IsDefault))
	builder.WriteString(", ")
	builder.WriteString("domain=")
	builder.WriteString(s.Domain)
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(s.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(s.Modified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sites is a parsable slice of Site.
type Sites []*Site
