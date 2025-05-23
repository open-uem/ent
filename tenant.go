// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/settings"
	"github.com/open-uem/ent/tenant"
)

// Tenant is the model entity for the Tenant schema.
type Tenant struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsDefault holds the value of the "is_default" field.
	IsDefault bool `json:"is_default,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TenantQuery when eager-loading is set.
	Edges        TenantEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TenantEdges holds the relations/edges for other nodes in the graph.
type TenantEdges struct {
	// Sites holds the value of the sites edge.
	Sites []*Site `json:"sites,omitempty"`
	// Settings holds the value of the settings edge.
	Settings *Settings `json:"settings,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Metadata holds the value of the metadata edge.
	Metadata []*OrgMetadata `json:"metadata,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// SitesOrErr returns the Sites value or an error if the edge
// was not loaded in eager-loading.
func (e TenantEdges) SitesOrErr() ([]*Site, error) {
	if e.loadedTypes[0] {
		return e.Sites, nil
	}
	return nil, &NotLoadedError{edge: "sites"}
}

// SettingsOrErr returns the Settings value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TenantEdges) SettingsOrErr() (*Settings, error) {
	if e.Settings != nil {
		return e.Settings, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: settings.Label}
	}
	return nil, &NotLoadedError{edge: "settings"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e TenantEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading.
func (e TenantEdges) MetadataOrErr() ([]*OrgMetadata, error) {
	if e.loadedTypes[3] {
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tenant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tenant.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case tenant.FieldID:
			values[i] = new(sql.NullInt64)
		case tenant.FieldDescription:
			values[i] = new(sql.NullString)
		case tenant.FieldCreated, tenant.FieldModified:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tenant fields.
func (t *Tenant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tenant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tenant.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case tenant.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				t.IsDefault = value.Bool
			}
		case tenant.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				t.Created = value.Time
			}
		case tenant.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				t.Modified = value.Time
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tenant.
// This includes values selected through modifiers, order, etc.
func (t *Tenant) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QuerySites queries the "sites" edge of the Tenant entity.
func (t *Tenant) QuerySites() *SiteQuery {
	return NewTenantClient(t.config).QuerySites(t)
}

// QuerySettings queries the "settings" edge of the Tenant entity.
func (t *Tenant) QuerySettings() *SettingsQuery {
	return NewTenantClient(t.config).QuerySettings(t)
}

// QueryTags queries the "tags" edge of the Tenant entity.
func (t *Tenant) QueryTags() *TagQuery {
	return NewTenantClient(t.config).QueryTags(t)
}

// QueryMetadata queries the "metadata" edge of the Tenant entity.
func (t *Tenant) QueryMetadata() *OrgMetadataQuery {
	return NewTenantClient(t.config).QueryMetadata(t)
}

// Update returns a builder for updating this Tenant.
// Note that you need to call Tenant.Unwrap() before calling this method if this Tenant
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tenant) Update() *TenantUpdateOne {
	return NewTenantClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tenant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tenant) Unwrap() *Tenant {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tenant is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tenant) String() string {
	var builder strings.Builder
	builder.WriteString("Tenant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", t.IsDefault))
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(t.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(t.Modified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Tenants is a parsable slice of Tenant.
type Tenants []*Tenant
