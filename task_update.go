// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/predicate"
	"github.com/open-uem/ent/profile"
	"github.com/open-uem/ent/tag"
	"github.com/open-uem/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks     []Hook
	mutation  *TaskMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TaskUpdate) SetName(s string) *TaskUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableName(s *string) *TaskUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetType sets the "type" field.
func (tu *TaskUpdate) SetType(t task.Type) *TaskUpdate {
	tu.mutation.SetType(t)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableType(t *task.Type) *TaskUpdate {
	if t != nil {
		tu.SetType(*t)
	}
	return tu
}

// SetExecute sets the "execute" field.
func (tu *TaskUpdate) SetExecute(s string) *TaskUpdate {
	tu.mutation.SetExecute(s)
	return tu
}

// SetNillableExecute sets the "execute" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableExecute(s *string) *TaskUpdate {
	if s != nil {
		tu.SetExecute(*s)
	}
	return tu
}

// ClearExecute clears the value of the "execute" field.
func (tu *TaskUpdate) ClearExecute() *TaskUpdate {
	tu.mutation.ClearExecute()
	return tu
}

// SetPackageID sets the "package_id" field.
func (tu *TaskUpdate) SetPackageID(s string) *TaskUpdate {
	tu.mutation.SetPackageID(s)
	return tu
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePackageID(s *string) *TaskUpdate {
	if s != nil {
		tu.SetPackageID(*s)
	}
	return tu
}

// ClearPackageID clears the value of the "package_id" field.
func (tu *TaskUpdate) ClearPackageID() *TaskUpdate {
	tu.mutation.ClearPackageID()
	return tu
}

// SetPackageName sets the "package_name" field.
func (tu *TaskUpdate) SetPackageName(s string) *TaskUpdate {
	tu.mutation.SetPackageName(s)
	return tu
}

// SetNillablePackageName sets the "package_name" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePackageName(s *string) *TaskUpdate {
	if s != nil {
		tu.SetPackageName(*s)
	}
	return tu
}

// ClearPackageName clears the value of the "package_name" field.
func (tu *TaskUpdate) ClearPackageName() *TaskUpdate {
	tu.mutation.ClearPackageName()
	return tu
}

// SetWhen sets the "when" field.
func (tu *TaskUpdate) SetWhen(t time.Time) *TaskUpdate {
	tu.mutation.SetWhen(t)
	return tu
}

// SetNillableWhen sets the "when" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableWhen(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetWhen(*t)
	}
	return tu
}

// ClearWhen clears the value of the "when" field.
func (tu *TaskUpdate) ClearWhen() *TaskUpdate {
	tu.mutation.ClearWhen()
	return tu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tu *TaskUpdate) AddTagIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddTagIDs(ids...)
	return tu
}

// AddTags adds the "tags" edges to the Tag entity.
func (tu *TaskUpdate) AddTags(t ...*Tag) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagIDs(ids...)
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (tu *TaskUpdate) SetProfileID(id int) *TaskUpdate {
	tu.mutation.SetProfileID(id)
	return tu
}

// SetNillableProfileID sets the "profile" edge to the Profile entity by ID if the given value is not nil.
func (tu *TaskUpdate) SetNillableProfileID(id *int) *TaskUpdate {
	if id != nil {
		tu = tu.SetProfileID(*id)
	}
	return tu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (tu *TaskUpdate) SetProfile(p *Profile) *TaskUpdate {
	return tu.SetProfileID(p.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (tu *TaskUpdate) ClearTags() *TaskUpdate {
	tu.mutation.ClearTags()
	return tu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (tu *TaskUpdate) RemoveTagIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveTagIDs(ids...)
	return tu
}

// RemoveTags removes "tags" edges to Tag entities.
func (tu *TaskUpdate) RemoveTags(t ...*Tag) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagIDs(ids...)
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (tu *TaskUpdate) ClearProfile() *TaskUpdate {
	tu.mutation.ClearProfile()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := task.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Task.name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.GetType(); ok {
		if err := task.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Task.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TaskUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Execute(); ok {
		_spec.SetField(task.FieldExecute, field.TypeString, value)
	}
	if tu.mutation.ExecuteCleared() {
		_spec.ClearField(task.FieldExecute, field.TypeString)
	}
	if value, ok := tu.mutation.PackageID(); ok {
		_spec.SetField(task.FieldPackageID, field.TypeString, value)
	}
	if tu.mutation.PackageIDCleared() {
		_spec.ClearField(task.FieldPackageID, field.TypeString)
	}
	if value, ok := tu.mutation.PackageName(); ok {
		_spec.SetField(task.FieldPackageName, field.TypeString, value)
	}
	if tu.mutation.PackageNameCleared() {
		_spec.ClearField(task.FieldPackageName, field.TypeString)
	}
	if value, ok := tu.mutation.When(); ok {
		_spec.SetField(task.FieldWhen, field.TypeTime, value)
	}
	if tu.mutation.WhenCleared() {
		_spec.ClearField(task.FieldWhen, field.TypeTime)
	}
	if tu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: []string{task.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: []string{task.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: []string{task.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ProfileTable,
			Columns: []string{task.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ProfileTable,
			Columns: []string{task.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TaskMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (tuo *TaskUpdateOne) SetName(s string) *TaskUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableName(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetType sets the "type" field.
func (tuo *TaskUpdateOne) SetType(t task.Type) *TaskUpdateOne {
	tuo.mutation.SetType(t)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableType(t *task.Type) *TaskUpdateOne {
	if t != nil {
		tuo.SetType(*t)
	}
	return tuo
}

// SetExecute sets the "execute" field.
func (tuo *TaskUpdateOne) SetExecute(s string) *TaskUpdateOne {
	tuo.mutation.SetExecute(s)
	return tuo
}

// SetNillableExecute sets the "execute" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableExecute(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetExecute(*s)
	}
	return tuo
}

// ClearExecute clears the value of the "execute" field.
func (tuo *TaskUpdateOne) ClearExecute() *TaskUpdateOne {
	tuo.mutation.ClearExecute()
	return tuo
}

// SetPackageID sets the "package_id" field.
func (tuo *TaskUpdateOne) SetPackageID(s string) *TaskUpdateOne {
	tuo.mutation.SetPackageID(s)
	return tuo
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePackageID(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetPackageID(*s)
	}
	return tuo
}

// ClearPackageID clears the value of the "package_id" field.
func (tuo *TaskUpdateOne) ClearPackageID() *TaskUpdateOne {
	tuo.mutation.ClearPackageID()
	return tuo
}

// SetPackageName sets the "package_name" field.
func (tuo *TaskUpdateOne) SetPackageName(s string) *TaskUpdateOne {
	tuo.mutation.SetPackageName(s)
	return tuo
}

// SetNillablePackageName sets the "package_name" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePackageName(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetPackageName(*s)
	}
	return tuo
}

// ClearPackageName clears the value of the "package_name" field.
func (tuo *TaskUpdateOne) ClearPackageName() *TaskUpdateOne {
	tuo.mutation.ClearPackageName()
	return tuo
}

// SetWhen sets the "when" field.
func (tuo *TaskUpdateOne) SetWhen(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetWhen(t)
	return tuo
}

// SetNillableWhen sets the "when" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableWhen(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetWhen(*t)
	}
	return tuo
}

// ClearWhen clears the value of the "when" field.
func (tuo *TaskUpdateOne) ClearWhen() *TaskUpdateOne {
	tuo.mutation.ClearWhen()
	return tuo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tuo *TaskUpdateOne) AddTagIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddTagIDs(ids...)
	return tuo
}

// AddTags adds the "tags" edges to the Tag entity.
func (tuo *TaskUpdateOne) AddTags(t ...*Tag) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagIDs(ids...)
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (tuo *TaskUpdateOne) SetProfileID(id int) *TaskUpdateOne {
	tuo.mutation.SetProfileID(id)
	return tuo
}

// SetNillableProfileID sets the "profile" edge to the Profile entity by ID if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableProfileID(id *int) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetProfileID(*id)
	}
	return tuo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (tuo *TaskUpdateOne) SetProfile(p *Profile) *TaskUpdateOne {
	return tuo.SetProfileID(p.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (tuo *TaskUpdateOne) ClearTags() *TaskUpdateOne {
	tuo.mutation.ClearTags()
	return tuo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (tuo *TaskUpdateOne) RemoveTagIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveTagIDs(ids...)
	return tuo
}

// RemoveTags removes "tags" edges to Tag entities.
func (tuo *TaskUpdateOne) RemoveTags(t ...*Tag) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagIDs(ids...)
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (tuo *TaskUpdateOne) ClearProfile() *TaskUpdateOne {
	tuo.mutation.ClearProfile()
	return tuo
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := task.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Task.name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.GetType(); ok {
		if err := task.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Task.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TaskUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Execute(); ok {
		_spec.SetField(task.FieldExecute, field.TypeString, value)
	}
	if tuo.mutation.ExecuteCleared() {
		_spec.ClearField(task.FieldExecute, field.TypeString)
	}
	if value, ok := tuo.mutation.PackageID(); ok {
		_spec.SetField(task.FieldPackageID, field.TypeString, value)
	}
	if tuo.mutation.PackageIDCleared() {
		_spec.ClearField(task.FieldPackageID, field.TypeString)
	}
	if value, ok := tuo.mutation.PackageName(); ok {
		_spec.SetField(task.FieldPackageName, field.TypeString, value)
	}
	if tuo.mutation.PackageNameCleared() {
		_spec.ClearField(task.FieldPackageName, field.TypeString)
	}
	if value, ok := tuo.mutation.When(); ok {
		_spec.SetField(task.FieldWhen, field.TypeTime, value)
	}
	if tuo.mutation.WhenCleared() {
		_spec.ClearField(task.FieldWhen, field.TypeTime)
	}
	if tuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: []string{task.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: []string{task.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: []string{task.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ProfileTable,
			Columns: []string{task.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ProfileTable,
			Columns: []string{task.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
