// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/predicate"
	"github.com/doncicuto/openuem_ent/release"
)

// ReleaseUpdate is the builder for updating Release entities.
type ReleaseUpdate struct {
	config
	hooks     []Hook
	mutation  *ReleaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ReleaseUpdate builder.
func (ru *ReleaseUpdate) Where(ps ...predicate.Release) *ReleaseUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetVersion sets the "version" field.
func (ru *ReleaseUpdate) SetVersion(s string) *ReleaseUpdate {
	ru.mutation.SetVersion(s)
	return ru
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableVersion(s *string) *ReleaseUpdate {
	if s != nil {
		ru.SetVersion(*s)
	}
	return ru
}

// ClearVersion clears the value of the "version" field.
func (ru *ReleaseUpdate) ClearVersion() *ReleaseUpdate {
	ru.mutation.ClearVersion()
	return ru
}

// SetChannel sets the "channel" field.
func (ru *ReleaseUpdate) SetChannel(s string) *ReleaseUpdate {
	ru.mutation.SetChannel(s)
	return ru
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableChannel(s *string) *ReleaseUpdate {
	if s != nil {
		ru.SetChannel(*s)
	}
	return ru
}

// ClearChannel clears the value of the "channel" field.
func (ru *ReleaseUpdate) ClearChannel() *ReleaseUpdate {
	ru.mutation.ClearChannel()
	return ru
}

// SetSummary sets the "summary" field.
func (ru *ReleaseUpdate) SetSummary(s string) *ReleaseUpdate {
	ru.mutation.SetSummary(s)
	return ru
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableSummary(s *string) *ReleaseUpdate {
	if s != nil {
		ru.SetSummary(*s)
	}
	return ru
}

// ClearSummary clears the value of the "summary" field.
func (ru *ReleaseUpdate) ClearSummary() *ReleaseUpdate {
	ru.mutation.ClearSummary()
	return ru
}

// SetReleaseNotes sets the "release_notes" field.
func (ru *ReleaseUpdate) SetReleaseNotes(s string) *ReleaseUpdate {
	ru.mutation.SetReleaseNotes(s)
	return ru
}

// SetNillableReleaseNotes sets the "release_notes" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableReleaseNotes(s *string) *ReleaseUpdate {
	if s != nil {
		ru.SetReleaseNotes(*s)
	}
	return ru
}

// ClearReleaseNotes clears the value of the "release_notes" field.
func (ru *ReleaseUpdate) ClearReleaseNotes() *ReleaseUpdate {
	ru.mutation.ClearReleaseNotes()
	return ru
}

// SetFileURL sets the "file_url" field.
func (ru *ReleaseUpdate) SetFileURL(s string) *ReleaseUpdate {
	ru.mutation.SetFileURL(s)
	return ru
}

// SetNillableFileURL sets the "file_url" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableFileURL(s *string) *ReleaseUpdate {
	if s != nil {
		ru.SetFileURL(*s)
	}
	return ru
}

// ClearFileURL clears the value of the "file_url" field.
func (ru *ReleaseUpdate) ClearFileURL() *ReleaseUpdate {
	ru.mutation.ClearFileURL()
	return ru
}

// SetChecksum sets the "checksum" field.
func (ru *ReleaseUpdate) SetChecksum(s string) *ReleaseUpdate {
	ru.mutation.SetChecksum(s)
	return ru
}

// SetNillableChecksum sets the "checksum" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableChecksum(s *string) *ReleaseUpdate {
	if s != nil {
		ru.SetChecksum(*s)
	}
	return ru
}

// ClearChecksum clears the value of the "checksum" field.
func (ru *ReleaseUpdate) ClearChecksum() *ReleaseUpdate {
	ru.mutation.ClearChecksum()
	return ru
}

// SetIsCritical sets the "is_critical" field.
func (ru *ReleaseUpdate) SetIsCritical(b bool) *ReleaseUpdate {
	ru.mutation.SetIsCritical(b)
	return ru
}

// SetNillableIsCritical sets the "is_critical" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableIsCritical(b *bool) *ReleaseUpdate {
	if b != nil {
		ru.SetIsCritical(*b)
	}
	return ru
}

// ClearIsCritical clears the value of the "is_critical" field.
func (ru *ReleaseUpdate) ClearIsCritical() *ReleaseUpdate {
	ru.mutation.ClearIsCritical()
	return ru
}

// SetReleaseDate sets the "release_date" field.
func (ru *ReleaseUpdate) SetReleaseDate(t time.Time) *ReleaseUpdate {
	ru.mutation.SetReleaseDate(t)
	return ru
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableReleaseDate(t *time.Time) *ReleaseUpdate {
	if t != nil {
		ru.SetReleaseDate(*t)
	}
	return ru
}

// ClearReleaseDate clears the value of the "release_date" field.
func (ru *ReleaseUpdate) ClearReleaseDate() *ReleaseUpdate {
	ru.mutation.ClearReleaseDate()
	return ru
}

// SetOs sets the "os" field.
func (ru *ReleaseUpdate) SetOs(s string) *ReleaseUpdate {
	ru.mutation.SetOs(s)
	return ru
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableOs(s *string) *ReleaseUpdate {
	if s != nil {
		ru.SetOs(*s)
	}
	return ru
}

// ClearOs clears the value of the "os" field.
func (ru *ReleaseUpdate) ClearOs() *ReleaseUpdate {
	ru.mutation.ClearOs()
	return ru
}

// SetArch sets the "arch" field.
func (ru *ReleaseUpdate) SetArch(s string) *ReleaseUpdate {
	ru.mutation.SetArch(s)
	return ru
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (ru *ReleaseUpdate) SetNillableArch(s *string) *ReleaseUpdate {
	if s != nil {
		ru.SetArch(*s)
	}
	return ru
}

// ClearArch clears the value of the "arch" field.
func (ru *ReleaseUpdate) ClearArch() *ReleaseUpdate {
	ru.mutation.ClearArch()
	return ru
}

// AddAgentIDs adds the "agents" edge to the Agent entity by IDs.
func (ru *ReleaseUpdate) AddAgentIDs(ids ...string) *ReleaseUpdate {
	ru.mutation.AddAgentIDs(ids...)
	return ru
}

// AddAgents adds the "agents" edges to the Agent entity.
func (ru *ReleaseUpdate) AddAgents(a ...*Agent) *ReleaseUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddAgentIDs(ids...)
}

// Mutation returns the ReleaseMutation object of the builder.
func (ru *ReleaseUpdate) Mutation() *ReleaseMutation {
	return ru.mutation
}

// ClearAgents clears all "agents" edges to the Agent entity.
func (ru *ReleaseUpdate) ClearAgents() *ReleaseUpdate {
	ru.mutation.ClearAgents()
	return ru
}

// RemoveAgentIDs removes the "agents" edge to Agent entities by IDs.
func (ru *ReleaseUpdate) RemoveAgentIDs(ids ...string) *ReleaseUpdate {
	ru.mutation.RemoveAgentIDs(ids...)
	return ru
}

// RemoveAgents removes "agents" edges to Agent entities.
func (ru *ReleaseUpdate) RemoveAgents(a ...*Agent) *ReleaseUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveAgentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReleaseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReleaseUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReleaseUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReleaseUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *ReleaseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReleaseUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *ReleaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(release.Table, release.Columns, sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Version(); ok {
		_spec.SetField(release.FieldVersion, field.TypeString, value)
	}
	if ru.mutation.VersionCleared() {
		_spec.ClearField(release.FieldVersion, field.TypeString)
	}
	if value, ok := ru.mutation.Channel(); ok {
		_spec.SetField(release.FieldChannel, field.TypeString, value)
	}
	if ru.mutation.ChannelCleared() {
		_spec.ClearField(release.FieldChannel, field.TypeString)
	}
	if value, ok := ru.mutation.Summary(); ok {
		_spec.SetField(release.FieldSummary, field.TypeString, value)
	}
	if ru.mutation.SummaryCleared() {
		_spec.ClearField(release.FieldSummary, field.TypeString)
	}
	if value, ok := ru.mutation.ReleaseNotes(); ok {
		_spec.SetField(release.FieldReleaseNotes, field.TypeString, value)
	}
	if ru.mutation.ReleaseNotesCleared() {
		_spec.ClearField(release.FieldReleaseNotes, field.TypeString)
	}
	if value, ok := ru.mutation.FileURL(); ok {
		_spec.SetField(release.FieldFileURL, field.TypeString, value)
	}
	if ru.mutation.FileURLCleared() {
		_spec.ClearField(release.FieldFileURL, field.TypeString)
	}
	if value, ok := ru.mutation.Checksum(); ok {
		_spec.SetField(release.FieldChecksum, field.TypeString, value)
	}
	if ru.mutation.ChecksumCleared() {
		_spec.ClearField(release.FieldChecksum, field.TypeString)
	}
	if value, ok := ru.mutation.IsCritical(); ok {
		_spec.SetField(release.FieldIsCritical, field.TypeBool, value)
	}
	if ru.mutation.IsCriticalCleared() {
		_spec.ClearField(release.FieldIsCritical, field.TypeBool)
	}
	if value, ok := ru.mutation.ReleaseDate(); ok {
		_spec.SetField(release.FieldReleaseDate, field.TypeTime, value)
	}
	if ru.mutation.ReleaseDateCleared() {
		_spec.ClearField(release.FieldReleaseDate, field.TypeTime)
	}
	if value, ok := ru.mutation.Os(); ok {
		_spec.SetField(release.FieldOs, field.TypeString, value)
	}
	if ru.mutation.OsCleared() {
		_spec.ClearField(release.FieldOs, field.TypeString)
	}
	if value, ok := ru.mutation.Arch(); ok {
		_spec.SetField(release.FieldArch, field.TypeString, value)
	}
	if ru.mutation.ArchCleared() {
		_spec.ClearField(release.FieldArch, field.TypeString)
	}
	if ru.mutation.AgentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   release.AgentsTable,
			Columns: []string{release.AgentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedAgentsIDs(); len(nodes) > 0 && !ru.mutation.AgentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   release.AgentsTable,
			Columns: []string{release.AgentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AgentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   release.AgentsTable,
			Columns: []string{release.AgentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{release.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReleaseUpdateOne is the builder for updating a single Release entity.
type ReleaseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ReleaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetVersion sets the "version" field.
func (ruo *ReleaseUpdateOne) SetVersion(s string) *ReleaseUpdateOne {
	ruo.mutation.SetVersion(s)
	return ruo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableVersion(s *string) *ReleaseUpdateOne {
	if s != nil {
		ruo.SetVersion(*s)
	}
	return ruo
}

// ClearVersion clears the value of the "version" field.
func (ruo *ReleaseUpdateOne) ClearVersion() *ReleaseUpdateOne {
	ruo.mutation.ClearVersion()
	return ruo
}

// SetChannel sets the "channel" field.
func (ruo *ReleaseUpdateOne) SetChannel(s string) *ReleaseUpdateOne {
	ruo.mutation.SetChannel(s)
	return ruo
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableChannel(s *string) *ReleaseUpdateOne {
	if s != nil {
		ruo.SetChannel(*s)
	}
	return ruo
}

// ClearChannel clears the value of the "channel" field.
func (ruo *ReleaseUpdateOne) ClearChannel() *ReleaseUpdateOne {
	ruo.mutation.ClearChannel()
	return ruo
}

// SetSummary sets the "summary" field.
func (ruo *ReleaseUpdateOne) SetSummary(s string) *ReleaseUpdateOne {
	ruo.mutation.SetSummary(s)
	return ruo
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableSummary(s *string) *ReleaseUpdateOne {
	if s != nil {
		ruo.SetSummary(*s)
	}
	return ruo
}

// ClearSummary clears the value of the "summary" field.
func (ruo *ReleaseUpdateOne) ClearSummary() *ReleaseUpdateOne {
	ruo.mutation.ClearSummary()
	return ruo
}

// SetReleaseNotes sets the "release_notes" field.
func (ruo *ReleaseUpdateOne) SetReleaseNotes(s string) *ReleaseUpdateOne {
	ruo.mutation.SetReleaseNotes(s)
	return ruo
}

// SetNillableReleaseNotes sets the "release_notes" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableReleaseNotes(s *string) *ReleaseUpdateOne {
	if s != nil {
		ruo.SetReleaseNotes(*s)
	}
	return ruo
}

// ClearReleaseNotes clears the value of the "release_notes" field.
func (ruo *ReleaseUpdateOne) ClearReleaseNotes() *ReleaseUpdateOne {
	ruo.mutation.ClearReleaseNotes()
	return ruo
}

// SetFileURL sets the "file_url" field.
func (ruo *ReleaseUpdateOne) SetFileURL(s string) *ReleaseUpdateOne {
	ruo.mutation.SetFileURL(s)
	return ruo
}

// SetNillableFileURL sets the "file_url" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableFileURL(s *string) *ReleaseUpdateOne {
	if s != nil {
		ruo.SetFileURL(*s)
	}
	return ruo
}

// ClearFileURL clears the value of the "file_url" field.
func (ruo *ReleaseUpdateOne) ClearFileURL() *ReleaseUpdateOne {
	ruo.mutation.ClearFileURL()
	return ruo
}

// SetChecksum sets the "checksum" field.
func (ruo *ReleaseUpdateOne) SetChecksum(s string) *ReleaseUpdateOne {
	ruo.mutation.SetChecksum(s)
	return ruo
}

// SetNillableChecksum sets the "checksum" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableChecksum(s *string) *ReleaseUpdateOne {
	if s != nil {
		ruo.SetChecksum(*s)
	}
	return ruo
}

// ClearChecksum clears the value of the "checksum" field.
func (ruo *ReleaseUpdateOne) ClearChecksum() *ReleaseUpdateOne {
	ruo.mutation.ClearChecksum()
	return ruo
}

// SetIsCritical sets the "is_critical" field.
func (ruo *ReleaseUpdateOne) SetIsCritical(b bool) *ReleaseUpdateOne {
	ruo.mutation.SetIsCritical(b)
	return ruo
}

// SetNillableIsCritical sets the "is_critical" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableIsCritical(b *bool) *ReleaseUpdateOne {
	if b != nil {
		ruo.SetIsCritical(*b)
	}
	return ruo
}

// ClearIsCritical clears the value of the "is_critical" field.
func (ruo *ReleaseUpdateOne) ClearIsCritical() *ReleaseUpdateOne {
	ruo.mutation.ClearIsCritical()
	return ruo
}

// SetReleaseDate sets the "release_date" field.
func (ruo *ReleaseUpdateOne) SetReleaseDate(t time.Time) *ReleaseUpdateOne {
	ruo.mutation.SetReleaseDate(t)
	return ruo
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableReleaseDate(t *time.Time) *ReleaseUpdateOne {
	if t != nil {
		ruo.SetReleaseDate(*t)
	}
	return ruo
}

// ClearReleaseDate clears the value of the "release_date" field.
func (ruo *ReleaseUpdateOne) ClearReleaseDate() *ReleaseUpdateOne {
	ruo.mutation.ClearReleaseDate()
	return ruo
}

// SetOs sets the "os" field.
func (ruo *ReleaseUpdateOne) SetOs(s string) *ReleaseUpdateOne {
	ruo.mutation.SetOs(s)
	return ruo
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableOs(s *string) *ReleaseUpdateOne {
	if s != nil {
		ruo.SetOs(*s)
	}
	return ruo
}

// ClearOs clears the value of the "os" field.
func (ruo *ReleaseUpdateOne) ClearOs() *ReleaseUpdateOne {
	ruo.mutation.ClearOs()
	return ruo
}

// SetArch sets the "arch" field.
func (ruo *ReleaseUpdateOne) SetArch(s string) *ReleaseUpdateOne {
	ruo.mutation.SetArch(s)
	return ruo
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (ruo *ReleaseUpdateOne) SetNillableArch(s *string) *ReleaseUpdateOne {
	if s != nil {
		ruo.SetArch(*s)
	}
	return ruo
}

// ClearArch clears the value of the "arch" field.
func (ruo *ReleaseUpdateOne) ClearArch() *ReleaseUpdateOne {
	ruo.mutation.ClearArch()
	return ruo
}

// AddAgentIDs adds the "agents" edge to the Agent entity by IDs.
func (ruo *ReleaseUpdateOne) AddAgentIDs(ids ...string) *ReleaseUpdateOne {
	ruo.mutation.AddAgentIDs(ids...)
	return ruo
}

// AddAgents adds the "agents" edges to the Agent entity.
func (ruo *ReleaseUpdateOne) AddAgents(a ...*Agent) *ReleaseUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddAgentIDs(ids...)
}

// Mutation returns the ReleaseMutation object of the builder.
func (ruo *ReleaseUpdateOne) Mutation() *ReleaseMutation {
	return ruo.mutation
}

// ClearAgents clears all "agents" edges to the Agent entity.
func (ruo *ReleaseUpdateOne) ClearAgents() *ReleaseUpdateOne {
	ruo.mutation.ClearAgents()
	return ruo
}

// RemoveAgentIDs removes the "agents" edge to Agent entities by IDs.
func (ruo *ReleaseUpdateOne) RemoveAgentIDs(ids ...string) *ReleaseUpdateOne {
	ruo.mutation.RemoveAgentIDs(ids...)
	return ruo
}

// RemoveAgents removes "agents" edges to Agent entities.
func (ruo *ReleaseUpdateOne) RemoveAgents(a ...*Agent) *ReleaseUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveAgentIDs(ids...)
}

// Where appends a list predicates to the ReleaseUpdate builder.
func (ruo *ReleaseUpdateOne) Where(ps ...predicate.Release) *ReleaseUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReleaseUpdateOne) Select(field string, fields ...string) *ReleaseUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Release entity.
func (ruo *ReleaseUpdateOne) Save(ctx context.Context) (*Release, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReleaseUpdateOne) SaveX(ctx context.Context) *Release {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReleaseUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReleaseUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *ReleaseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReleaseUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *ReleaseUpdateOne) sqlSave(ctx context.Context) (_node *Release, err error) {
	_spec := sqlgraph.NewUpdateSpec(release.Table, release.Columns, sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`openuem_ent: missing "Release.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, release.FieldID)
		for _, f := range fields {
			if !release.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
			}
			if f != release.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Version(); ok {
		_spec.SetField(release.FieldVersion, field.TypeString, value)
	}
	if ruo.mutation.VersionCleared() {
		_spec.ClearField(release.FieldVersion, field.TypeString)
	}
	if value, ok := ruo.mutation.Channel(); ok {
		_spec.SetField(release.FieldChannel, field.TypeString, value)
	}
	if ruo.mutation.ChannelCleared() {
		_spec.ClearField(release.FieldChannel, field.TypeString)
	}
	if value, ok := ruo.mutation.Summary(); ok {
		_spec.SetField(release.FieldSummary, field.TypeString, value)
	}
	if ruo.mutation.SummaryCleared() {
		_spec.ClearField(release.FieldSummary, field.TypeString)
	}
	if value, ok := ruo.mutation.ReleaseNotes(); ok {
		_spec.SetField(release.FieldReleaseNotes, field.TypeString, value)
	}
	if ruo.mutation.ReleaseNotesCleared() {
		_spec.ClearField(release.FieldReleaseNotes, field.TypeString)
	}
	if value, ok := ruo.mutation.FileURL(); ok {
		_spec.SetField(release.FieldFileURL, field.TypeString, value)
	}
	if ruo.mutation.FileURLCleared() {
		_spec.ClearField(release.FieldFileURL, field.TypeString)
	}
	if value, ok := ruo.mutation.Checksum(); ok {
		_spec.SetField(release.FieldChecksum, field.TypeString, value)
	}
	if ruo.mutation.ChecksumCleared() {
		_spec.ClearField(release.FieldChecksum, field.TypeString)
	}
	if value, ok := ruo.mutation.IsCritical(); ok {
		_spec.SetField(release.FieldIsCritical, field.TypeBool, value)
	}
	if ruo.mutation.IsCriticalCleared() {
		_spec.ClearField(release.FieldIsCritical, field.TypeBool)
	}
	if value, ok := ruo.mutation.ReleaseDate(); ok {
		_spec.SetField(release.FieldReleaseDate, field.TypeTime, value)
	}
	if ruo.mutation.ReleaseDateCleared() {
		_spec.ClearField(release.FieldReleaseDate, field.TypeTime)
	}
	if value, ok := ruo.mutation.Os(); ok {
		_spec.SetField(release.FieldOs, field.TypeString, value)
	}
	if ruo.mutation.OsCleared() {
		_spec.ClearField(release.FieldOs, field.TypeString)
	}
	if value, ok := ruo.mutation.Arch(); ok {
		_spec.SetField(release.FieldArch, field.TypeString, value)
	}
	if ruo.mutation.ArchCleared() {
		_spec.ClearField(release.FieldArch, field.TypeString)
	}
	if ruo.mutation.AgentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   release.AgentsTable,
			Columns: []string{release.AgentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedAgentsIDs(); len(nodes) > 0 && !ruo.mutation.AgentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   release.AgentsTable,
			Columns: []string{release.AgentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AgentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   release.AgentsTable,
			Columns: []string{release.AgentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Release{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{release.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
