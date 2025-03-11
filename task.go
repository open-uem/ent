// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/profile"
	"github.com/open-uem/ent/task"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type task.Type `json:"type,omitempty"`
	// PackageID holds the value of the "package_id" field.
	PackageID string `json:"package_id,omitempty"`
	// PackageName holds the value of the "package_name" field.
	PackageName string `json:"package_name,omitempty"`
	// RegistryKey holds the value of the "registry_key" field.
	RegistryKey string `json:"registry_key,omitempty"`
	// RegistryKeyValueName holds the value of the "registry_key_value_name" field.
	RegistryKeyValueName string `json:"registry_key_value_name,omitempty"`
	// RegistryKeyValueType holds the value of the "registry_key_value_type" field.
	RegistryKeyValueType task.RegistryKeyValueType `json:"registry_key_value_type,omitempty"`
	// RegistryKeyValueData holds the value of the "registry_key_value_data" field.
	RegistryKeyValueData string `json:"registry_key_value_data,omitempty"`
	// RegistryHex holds the value of the "registry_hex" field.
	RegistryHex bool `json:"registry_hex,omitempty"`
	// RegistryForce holds the value of the "registry_force" field.
	RegistryForce bool `json:"registry_force,omitempty"`
	// LocalUserUsername holds the value of the "local_user_username" field.
	LocalUserUsername string `json:"local_user_username,omitempty"`
	// LocalUserDescription holds the value of the "local_user_description" field.
	LocalUserDescription string `json:"local_user_description,omitempty"`
	// LocalUserDisable holds the value of the "local_user_disable" field.
	LocalUserDisable bool `json:"local_user_disable,omitempty"`
	// LocalUserFullname holds the value of the "local_user_fullname" field.
	LocalUserFullname string `json:"local_user_fullname,omitempty"`
	// LocalUserPassword holds the value of the "local_user_password" field.
	LocalUserPassword string `json:"local_user_password,omitempty"`
	// LocalUserPasswordChangeNotAllowed holds the value of the "local_user_password_change_not_allowed" field.
	LocalUserPasswordChangeNotAllowed bool `json:"local_user_password_change_not_allowed,omitempty"`
	// LocalUserPasswordChangeRequired holds the value of the "local_user_password_change_required" field.
	LocalUserPasswordChangeRequired bool `json:"local_user_password_change_required,omitempty"`
	// LocalUserPasswordNeverExpires holds the value of the "local_user_password_never_expires" field.
	LocalUserPasswordNeverExpires bool `json:"local_user_password_never_expires,omitempty"`
	// When holds the value of the "when" field.
	When time.Time `json:"when,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges         TaskEdges `json:"edges"`
	profile_tasks *int
	selectValues  sql.SelectValues
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Profile holds the value of the profile edge.
	Profile *Profile `json:"profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) ProfileOrErr() (*Profile, error) {
	if e.Profile != nil {
		return e.Profile, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: profile.Label}
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldRegistryHex, task.FieldRegistryForce, task.FieldLocalUserDisable, task.FieldLocalUserPasswordChangeNotAllowed, task.FieldLocalUserPasswordChangeRequired, task.FieldLocalUserPasswordNeverExpires:
			values[i] = new(sql.NullBool)
		case task.FieldID:
			values[i] = new(sql.NullInt64)
		case task.FieldName, task.FieldType, task.FieldPackageID, task.FieldPackageName, task.FieldRegistryKey, task.FieldRegistryKeyValueName, task.FieldRegistryKeyValueType, task.FieldRegistryKeyValueData, task.FieldLocalUserUsername, task.FieldLocalUserDescription, task.FieldLocalUserFullname, task.FieldLocalUserPassword:
			values[i] = new(sql.NullString)
		case task.FieldWhen:
			values[i] = new(sql.NullTime)
		case task.ForeignKeys[0]: // profile_tasks
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case task.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case task.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = task.Type(value.String)
			}
		case task.FieldPackageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field package_id", values[i])
			} else if value.Valid {
				t.PackageID = value.String
			}
		case task.FieldPackageName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field package_name", values[i])
			} else if value.Valid {
				t.PackageName = value.String
			}
		case task.FieldRegistryKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field registry_key", values[i])
			} else if value.Valid {
				t.RegistryKey = value.String
			}
		case task.FieldRegistryKeyValueName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field registry_key_value_name", values[i])
			} else if value.Valid {
				t.RegistryKeyValueName = value.String
			}
		case task.FieldRegistryKeyValueType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field registry_key_value_type", values[i])
			} else if value.Valid {
				t.RegistryKeyValueType = task.RegistryKeyValueType(value.String)
			}
		case task.FieldRegistryKeyValueData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field registry_key_value_data", values[i])
			} else if value.Valid {
				t.RegistryKeyValueData = value.String
			}
		case task.FieldRegistryHex:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field registry_hex", values[i])
			} else if value.Valid {
				t.RegistryHex = value.Bool
			}
		case task.FieldRegistryForce:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field registry_force", values[i])
			} else if value.Valid {
				t.RegistryForce = value.Bool
			}
		case task.FieldLocalUserUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field local_user_username", values[i])
			} else if value.Valid {
				t.LocalUserUsername = value.String
			}
		case task.FieldLocalUserDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field local_user_description", values[i])
			} else if value.Valid {
				t.LocalUserDescription = value.String
			}
		case task.FieldLocalUserDisable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field local_user_disable", values[i])
			} else if value.Valid {
				t.LocalUserDisable = value.Bool
			}
		case task.FieldLocalUserFullname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field local_user_fullname", values[i])
			} else if value.Valid {
				t.LocalUserFullname = value.String
			}
		case task.FieldLocalUserPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field local_user_password", values[i])
			} else if value.Valid {
				t.LocalUserPassword = value.String
			}
		case task.FieldLocalUserPasswordChangeNotAllowed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field local_user_password_change_not_allowed", values[i])
			} else if value.Valid {
				t.LocalUserPasswordChangeNotAllowed = value.Bool
			}
		case task.FieldLocalUserPasswordChangeRequired:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field local_user_password_change_required", values[i])
			} else if value.Valid {
				t.LocalUserPasswordChangeRequired = value.Bool
			}
		case task.FieldLocalUserPasswordNeverExpires:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field local_user_password_never_expires", values[i])
			} else if value.Valid {
				t.LocalUserPasswordNeverExpires = value.Bool
			}
		case task.FieldWhen:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field when", values[i])
			} else if value.Valid {
				t.When = value.Time
			}
		case task.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field profile_tasks", value)
			} else if value.Valid {
				t.profile_tasks = new(int)
				*t.profile_tasks = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Task.
// This includes values selected through modifiers, order, etc.
func (t *Task) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryTags queries the "tags" edge of the Task entity.
func (t *Task) QueryTags() *TagQuery {
	return NewTaskClient(t.config).QueryTags(t)
}

// QueryProfile queries the "profile" edge of the Task entity.
func (t *Task) QueryProfile() *ProfileQuery {
	return NewTaskClient(t.config).QueryProfile(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return NewTaskClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", ")
	builder.WriteString("package_id=")
	builder.WriteString(t.PackageID)
	builder.WriteString(", ")
	builder.WriteString("package_name=")
	builder.WriteString(t.PackageName)
	builder.WriteString(", ")
	builder.WriteString("registry_key=")
	builder.WriteString(t.RegistryKey)
	builder.WriteString(", ")
	builder.WriteString("registry_key_value_name=")
	builder.WriteString(t.RegistryKeyValueName)
	builder.WriteString(", ")
	builder.WriteString("registry_key_value_type=")
	builder.WriteString(fmt.Sprintf("%v", t.RegistryKeyValueType))
	builder.WriteString(", ")
	builder.WriteString("registry_key_value_data=")
	builder.WriteString(t.RegistryKeyValueData)
	builder.WriteString(", ")
	builder.WriteString("registry_hex=")
	builder.WriteString(fmt.Sprintf("%v", t.RegistryHex))
	builder.WriteString(", ")
	builder.WriteString("registry_force=")
	builder.WriteString(fmt.Sprintf("%v", t.RegistryForce))
	builder.WriteString(", ")
	builder.WriteString("local_user_username=")
	builder.WriteString(t.LocalUserUsername)
	builder.WriteString(", ")
	builder.WriteString("local_user_description=")
	builder.WriteString(t.LocalUserDescription)
	builder.WriteString(", ")
	builder.WriteString("local_user_disable=")
	builder.WriteString(fmt.Sprintf("%v", t.LocalUserDisable))
	builder.WriteString(", ")
	builder.WriteString("local_user_fullname=")
	builder.WriteString(t.LocalUserFullname)
	builder.WriteString(", ")
	builder.WriteString("local_user_password=")
	builder.WriteString(t.LocalUserPassword)
	builder.WriteString(", ")
	builder.WriteString("local_user_password_change_not_allowed=")
	builder.WriteString(fmt.Sprintf("%v", t.LocalUserPasswordChangeNotAllowed))
	builder.WriteString(", ")
	builder.WriteString("local_user_password_change_required=")
	builder.WriteString(fmt.Sprintf("%v", t.LocalUserPasswordChangeRequired))
	builder.WriteString(", ")
	builder.WriteString("local_user_password_never_expires=")
	builder.WriteString(fmt.Sprintf("%v", t.LocalUserPasswordNeverExpires))
	builder.WriteString(", ")
	builder.WriteString("when=")
	builder.WriteString(t.When.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task
