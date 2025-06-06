// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/predicate"
	"github.com/open-uem/ent/profile"
	"github.com/open-uem/ent/profileissue"
	"github.com/open-uem/ent/site"
	"github.com/open-uem/ent/tag"
	"github.com/open-uem/ent/task"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks     []Hook
	mutation  *ProfileMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ProfileUpdate builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProfileUpdate) SetName(s string) *ProfileUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableName(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetApplyToAll sets the "apply_to_all" field.
func (pu *ProfileUpdate) SetApplyToAll(b bool) *ProfileUpdate {
	pu.mutation.SetApplyToAll(b)
	return pu
}

// SetNillableApplyToAll sets the "apply_to_all" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableApplyToAll(b *bool) *ProfileUpdate {
	if b != nil {
		pu.SetApplyToAll(*b)
	}
	return pu
}

// SetType sets the "type" field.
func (pu *ProfileUpdate) SetType(pr profile.Type) *ProfileUpdate {
	pu.mutation.SetType(pr)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableType(pr *profile.Type) *ProfileUpdate {
	if pr != nil {
		pu.SetType(*pr)
	}
	return pu
}

// ClearType clears the value of the "type" field.
func (pu *ProfileUpdate) ClearType() *ProfileUpdate {
	pu.mutation.ClearType()
	return pu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (pu *ProfileUpdate) AddTagIDs(ids ...int) *ProfileUpdate {
	pu.mutation.AddTagIDs(ids...)
	return pu
}

// AddTags adds the "tags" edges to the Tag entity.
func (pu *ProfileUpdate) AddTags(t ...*Tag) *ProfileUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTagIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (pu *ProfileUpdate) AddTaskIDs(ids ...int) *ProfileUpdate {
	pu.mutation.AddTaskIDs(ids...)
	return pu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (pu *ProfileUpdate) AddTasks(t ...*Task) *ProfileUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTaskIDs(ids...)
}

// AddIssueIDs adds the "issues" edge to the ProfileIssue entity by IDs.
func (pu *ProfileUpdate) AddIssueIDs(ids ...int) *ProfileUpdate {
	pu.mutation.AddIssueIDs(ids...)
	return pu
}

// AddIssues adds the "issues" edges to the ProfileIssue entity.
func (pu *ProfileUpdate) AddIssues(p ...*ProfileIssue) *ProfileUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddIssueIDs(ids...)
}

// SetSiteID sets the "site" edge to the Site entity by ID.
func (pu *ProfileUpdate) SetSiteID(id int) *ProfileUpdate {
	pu.mutation.SetSiteID(id)
	return pu
}

// SetNillableSiteID sets the "site" edge to the Site entity by ID if the given value is not nil.
func (pu *ProfileUpdate) SetNillableSiteID(id *int) *ProfileUpdate {
	if id != nil {
		pu = pu.SetSiteID(*id)
	}
	return pu
}

// SetSite sets the "site" edge to the Site entity.
func (pu *ProfileUpdate) SetSite(s *Site) *ProfileUpdate {
	return pu.SetSiteID(s.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (pu *ProfileUpdate) ClearTags() *ProfileUpdate {
	pu.mutation.ClearTags()
	return pu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (pu *ProfileUpdate) RemoveTagIDs(ids ...int) *ProfileUpdate {
	pu.mutation.RemoveTagIDs(ids...)
	return pu
}

// RemoveTags removes "tags" edges to Tag entities.
func (pu *ProfileUpdate) RemoveTags(t ...*Tag) *ProfileUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTagIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (pu *ProfileUpdate) ClearTasks() *ProfileUpdate {
	pu.mutation.ClearTasks()
	return pu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (pu *ProfileUpdate) RemoveTaskIDs(ids ...int) *ProfileUpdate {
	pu.mutation.RemoveTaskIDs(ids...)
	return pu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (pu *ProfileUpdate) RemoveTasks(t ...*Task) *ProfileUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTaskIDs(ids...)
}

// ClearIssues clears all "issues" edges to the ProfileIssue entity.
func (pu *ProfileUpdate) ClearIssues() *ProfileUpdate {
	pu.mutation.ClearIssues()
	return pu
}

// RemoveIssueIDs removes the "issues" edge to ProfileIssue entities by IDs.
func (pu *ProfileUpdate) RemoveIssueIDs(ids ...int) *ProfileUpdate {
	pu.mutation.RemoveIssueIDs(ids...)
	return pu
}

// RemoveIssues removes "issues" edges to ProfileIssue entities.
func (pu *ProfileUpdate) RemoveIssues(p ...*ProfileIssue) *ProfileUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveIssueIDs(ids...)
}

// ClearSite clears the "site" edge to the Site entity.
func (pu *ProfileUpdate) ClearSite() *ProfileUpdate {
	pu.mutation.ClearSite()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProfileUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := profile.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Profile.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.GetType(); ok {
		if err := profile.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Profile.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *ProfileUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProfileUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(profile.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.ApplyToAll(); ok {
		_spec.SetField(profile.FieldApplyToAll, field.TypeBool, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(profile.FieldType, field.TypeEnum, value)
	}
	if pu.mutation.TypeCleared() {
		_spec.ClearField(profile.FieldType, field.TypeEnum)
	}
	if pu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.TagsTable,
			Columns: profile.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !pu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.TagsTable,
			Columns: profile.TagsPrimaryKey,
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
	if nodes := pu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.TagsTable,
			Columns: profile.TagsPrimaryKey,
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
	if pu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.TasksTable,
			Columns: []string{profile.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !pu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.TasksTable,
			Columns: []string{profile.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.TasksTable,
			Columns: []string{profile.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.IssuesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.IssuesTable,
			Columns: []string{profile.IssuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedIssuesIDs(); len(nodes) > 0 && !pu.mutation.IssuesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.IssuesTable,
			Columns: []string{profile.IssuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.IssuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.IssuesTable,
			Columns: []string{profile.IssuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SiteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.SiteTable,
			Columns: []string{profile.SiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(site.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.SiteTable,
			Columns: []string{profile.SiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(site.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ProfileMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (puo *ProfileUpdateOne) SetName(s string) *ProfileUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableName(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetApplyToAll sets the "apply_to_all" field.
func (puo *ProfileUpdateOne) SetApplyToAll(b bool) *ProfileUpdateOne {
	puo.mutation.SetApplyToAll(b)
	return puo
}

// SetNillableApplyToAll sets the "apply_to_all" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableApplyToAll(b *bool) *ProfileUpdateOne {
	if b != nil {
		puo.SetApplyToAll(*b)
	}
	return puo
}

// SetType sets the "type" field.
func (puo *ProfileUpdateOne) SetType(pr profile.Type) *ProfileUpdateOne {
	puo.mutation.SetType(pr)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableType(pr *profile.Type) *ProfileUpdateOne {
	if pr != nil {
		puo.SetType(*pr)
	}
	return puo
}

// ClearType clears the value of the "type" field.
func (puo *ProfileUpdateOne) ClearType() *ProfileUpdateOne {
	puo.mutation.ClearType()
	return puo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (puo *ProfileUpdateOne) AddTagIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.AddTagIDs(ids...)
	return puo
}

// AddTags adds the "tags" edges to the Tag entity.
func (puo *ProfileUpdateOne) AddTags(t ...*Tag) *ProfileUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTagIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (puo *ProfileUpdateOne) AddTaskIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.AddTaskIDs(ids...)
	return puo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (puo *ProfileUpdateOne) AddTasks(t ...*Task) *ProfileUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTaskIDs(ids...)
}

// AddIssueIDs adds the "issues" edge to the ProfileIssue entity by IDs.
func (puo *ProfileUpdateOne) AddIssueIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.AddIssueIDs(ids...)
	return puo
}

// AddIssues adds the "issues" edges to the ProfileIssue entity.
func (puo *ProfileUpdateOne) AddIssues(p ...*ProfileIssue) *ProfileUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddIssueIDs(ids...)
}

// SetSiteID sets the "site" edge to the Site entity by ID.
func (puo *ProfileUpdateOne) SetSiteID(id int) *ProfileUpdateOne {
	puo.mutation.SetSiteID(id)
	return puo
}

// SetNillableSiteID sets the "site" edge to the Site entity by ID if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableSiteID(id *int) *ProfileUpdateOne {
	if id != nil {
		puo = puo.SetSiteID(*id)
	}
	return puo
}

// SetSite sets the "site" edge to the Site entity.
func (puo *ProfileUpdateOne) SetSite(s *Site) *ProfileUpdateOne {
	return puo.SetSiteID(s.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (puo *ProfileUpdateOne) ClearTags() *ProfileUpdateOne {
	puo.mutation.ClearTags()
	return puo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (puo *ProfileUpdateOne) RemoveTagIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.RemoveTagIDs(ids...)
	return puo
}

// RemoveTags removes "tags" edges to Tag entities.
func (puo *ProfileUpdateOne) RemoveTags(t ...*Tag) *ProfileUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTagIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (puo *ProfileUpdateOne) ClearTasks() *ProfileUpdateOne {
	puo.mutation.ClearTasks()
	return puo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (puo *ProfileUpdateOne) RemoveTaskIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.RemoveTaskIDs(ids...)
	return puo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (puo *ProfileUpdateOne) RemoveTasks(t ...*Task) *ProfileUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTaskIDs(ids...)
}

// ClearIssues clears all "issues" edges to the ProfileIssue entity.
func (puo *ProfileUpdateOne) ClearIssues() *ProfileUpdateOne {
	puo.mutation.ClearIssues()
	return puo
}

// RemoveIssueIDs removes the "issues" edge to ProfileIssue entities by IDs.
func (puo *ProfileUpdateOne) RemoveIssueIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.RemoveIssueIDs(ids...)
	return puo
}

// RemoveIssues removes "issues" edges to ProfileIssue entities.
func (puo *ProfileUpdateOne) RemoveIssues(p ...*ProfileIssue) *ProfileUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveIssueIDs(ids...)
}

// ClearSite clears the "site" edge to the Site entity.
func (puo *ProfileUpdateOne) ClearSite() *ProfileUpdateOne {
	puo.mutation.ClearSite()
	return puo
}

// Where appends a list predicates to the ProfileUpdate builder.
func (puo *ProfileUpdateOne) Where(ps ...predicate.Profile) *ProfileUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProfileUpdateOne) Select(field string, fields ...string) *ProfileUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Profile entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProfileUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := profile.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Profile.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.GetType(); ok {
		if err := profile.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Profile.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *ProfileUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProfileUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (_node *Profile, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Profile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for _, f := range fields {
			if !profile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(profile.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.ApplyToAll(); ok {
		_spec.SetField(profile.FieldApplyToAll, field.TypeBool, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(profile.FieldType, field.TypeEnum, value)
	}
	if puo.mutation.TypeCleared() {
		_spec.ClearField(profile.FieldType, field.TypeEnum)
	}
	if puo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.TagsTable,
			Columns: profile.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !puo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.TagsTable,
			Columns: profile.TagsPrimaryKey,
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
	if nodes := puo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.TagsTable,
			Columns: profile.TagsPrimaryKey,
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
	if puo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.TasksTable,
			Columns: []string{profile.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !puo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.TasksTable,
			Columns: []string{profile.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.TasksTable,
			Columns: []string{profile.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.IssuesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.IssuesTable,
			Columns: []string{profile.IssuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedIssuesIDs(); len(nodes) > 0 && !puo.mutation.IssuesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.IssuesTable,
			Columns: []string{profile.IssuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.IssuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.IssuesTable,
			Columns: []string{profile.IssuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SiteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.SiteTable,
			Columns: []string{profile.SiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(site.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.SiteTable,
			Columns: []string{profile.SiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(site.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Profile{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
