// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/deployment"
)

// Deployment is the model entity for the Deployment schema.
type Deployment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PackageID holds the value of the "package_id" field.
	PackageID string `json:"package_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Installed holds the value of the "installed" field.
	Installed time.Time `json:"installed,omitempty"`
	// Updated holds the value of the "updated" field.
	Updated time.Time `json:"updated,omitempty"`
	// ByProfile holds the value of the "by_profile" field.
	ByProfile bool `json:"by_profile,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeploymentQuery when eager-loading is set.
	Edges             DeploymentEdges `json:"edges"`
	agent_deployments *string
	selectValues      sql.SelectValues
}

// DeploymentEdges holds the relations/edges for other nodes in the graph.
type DeploymentEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deployment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deployment.FieldByProfile:
			values[i] = new(sql.NullBool)
		case deployment.FieldID:
			values[i] = new(sql.NullInt64)
		case deployment.FieldPackageID, deployment.FieldName, deployment.FieldVersion:
			values[i] = new(sql.NullString)
		case deployment.FieldInstalled, deployment.FieldUpdated:
			values[i] = new(sql.NullTime)
		case deployment.ForeignKeys[0]: // agent_deployments
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deployment fields.
func (d *Deployment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deployment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case deployment.FieldPackageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field package_id", values[i])
			} else if value.Valid {
				d.PackageID = value.String
			}
		case deployment.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case deployment.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				d.Version = value.String
			}
		case deployment.FieldInstalled:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field installed", values[i])
			} else if value.Valid {
				d.Installed = value.Time
			}
		case deployment.FieldUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated", values[i])
			} else if value.Valid {
				d.Updated = value.Time
			}
		case deployment.FieldByProfile:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field by_profile", values[i])
			} else if value.Valid {
				d.ByProfile = value.Bool
			}
		case deployment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_deployments", values[i])
			} else if value.Valid {
				d.agent_deployments = new(string)
				*d.agent_deployments = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Deployment.
// This includes values selected through modifiers, order, etc.
func (d *Deployment) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Deployment entity.
func (d *Deployment) QueryOwner() *AgentQuery {
	return NewDeploymentClient(d.config).QueryOwner(d)
}

// Update returns a builder for updating this Deployment.
// Note that you need to call Deployment.Unwrap() before calling this method if this Deployment
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deployment) Update() *DeploymentUpdateOne {
	return NewDeploymentClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Deployment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Deployment) Unwrap() *Deployment {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deployment is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deployment) String() string {
	var builder strings.Builder
	builder.WriteString("Deployment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("package_id=")
	builder.WriteString(d.PackageID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(d.Version)
	builder.WriteString(", ")
	builder.WriteString("installed=")
	builder.WriteString(d.Installed.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated=")
	builder.WriteString(d.Updated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("by_profile=")
	builder.WriteString(fmt.Sprintf("%v", d.ByProfile))
	builder.WriteByte(')')
	return builder.String()
}

// Deployments is a parsable slice of Deployment.
type Deployments []*Deployment
