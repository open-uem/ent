// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/app"
)

// App is the model entity for the App schema.
type App struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Publisher holds the value of the "publisher" field.
	Publisher string `json:"publisher,omitempty"`
	// InstallDate holds the value of the "install_date" field.
	InstallDate string `json:"install_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppQuery when eager-loading is set.
	Edges        AppEdges `json:"edges"`
	agent_apps   *string
	selectValues sql.SelectValues
}

// AppEdges holds the relations/edges for other nodes in the graph.
type AppEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*App) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case app.FieldID:
			values[i] = new(sql.NullInt64)
		case app.FieldName, app.FieldVersion, app.FieldPublisher, app.FieldInstallDate:
			values[i] = new(sql.NullString)
		case app.ForeignKeys[0]: // agent_apps
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the App fields.
func (a *App) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case app.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case app.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case app.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				a.Version = value.String
			}
		case app.FieldPublisher:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publisher", values[i])
			} else if value.Valid {
				a.Publisher = value.String
			}
		case app.FieldInstallDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field install_date", values[i])
			} else if value.Valid {
				a.InstallDate = value.String
			}
		case app.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_apps", values[i])
			} else if value.Valid {
				a.agent_apps = new(string)
				*a.agent_apps = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the App.
// This includes values selected through modifiers, order, etc.
func (a *App) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the App entity.
func (a *App) QueryOwner() *AgentQuery {
	return NewAppClient(a.config).QueryOwner(a)
}

// Update returns a builder for updating this App.
// Note that you need to call App.Unwrap() before calling this method if this App
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *App) Update() *AppUpdateOne {
	return NewAppClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the App entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *App) Unwrap() *App {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("openuem_ent: App is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *App) String() string {
	var builder strings.Builder
	builder.WriteString("App(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(a.Version)
	builder.WriteString(", ")
	builder.WriteString("publisher=")
	builder.WriteString(a.Publisher)
	builder.WriteString(", ")
	builder.WriteString("install_date=")
	builder.WriteString(a.InstallDate)
	builder.WriteByte(')')
	return builder.String()
}

// Apps is a parsable slice of App.
type Apps []*App
