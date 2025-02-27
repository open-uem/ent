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
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/predicate"
	"github.com/open-uem/ent/systemupdate"
)

// SystemUpdateUpdate is the builder for updating SystemUpdate entities.
type SystemUpdateUpdate struct {
	config
	hooks     []Hook
	mutation  *SystemUpdateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SystemUpdateUpdate builder.
func (suu *SystemUpdateUpdate) Where(ps ...predicate.SystemUpdate) *SystemUpdateUpdate {
	suu.mutation.Where(ps...)
	return suu
}

// SetSystemUpdateStatus sets the "system_update_status" field.
func (suu *SystemUpdateUpdate) SetSystemUpdateStatus(s string) *SystemUpdateUpdate {
	suu.mutation.SetSystemUpdateStatus(s)
	return suu
}

// SetNillableSystemUpdateStatus sets the "system_update_status" field if the given value is not nil.
func (suu *SystemUpdateUpdate) SetNillableSystemUpdateStatus(s *string) *SystemUpdateUpdate {
	if s != nil {
		suu.SetSystemUpdateStatus(*s)
	}
	return suu
}

// SetLastInstall sets the "last_install" field.
func (suu *SystemUpdateUpdate) SetLastInstall(t time.Time) *SystemUpdateUpdate {
	suu.mutation.SetLastInstall(t)
	return suu
}

// SetNillableLastInstall sets the "last_install" field if the given value is not nil.
func (suu *SystemUpdateUpdate) SetNillableLastInstall(t *time.Time) *SystemUpdateUpdate {
	if t != nil {
		suu.SetLastInstall(*t)
	}
	return suu
}

// SetLastSearch sets the "last_search" field.
func (suu *SystemUpdateUpdate) SetLastSearch(t time.Time) *SystemUpdateUpdate {
	suu.mutation.SetLastSearch(t)
	return suu
}

// SetNillableLastSearch sets the "last_search" field if the given value is not nil.
func (suu *SystemUpdateUpdate) SetNillableLastSearch(t *time.Time) *SystemUpdateUpdate {
	if t != nil {
		suu.SetLastSearch(*t)
	}
	return suu
}

// SetPendingUpdates sets the "pending_updates" field.
func (suu *SystemUpdateUpdate) SetPendingUpdates(b bool) *SystemUpdateUpdate {
	suu.mutation.SetPendingUpdates(b)
	return suu
}

// SetNillablePendingUpdates sets the "pending_updates" field if the given value is not nil.
func (suu *SystemUpdateUpdate) SetNillablePendingUpdates(b *bool) *SystemUpdateUpdate {
	if b != nil {
		suu.SetPendingUpdates(*b)
	}
	return suu
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (suu *SystemUpdateUpdate) SetOwnerID(id string) *SystemUpdateUpdate {
	suu.mutation.SetOwnerID(id)
	return suu
}

// SetOwner sets the "owner" edge to the Agent entity.
func (suu *SystemUpdateUpdate) SetOwner(a *Agent) *SystemUpdateUpdate {
	return suu.SetOwnerID(a.ID)
}

// Mutation returns the SystemUpdateMutation object of the builder.
func (suu *SystemUpdateUpdate) Mutation() *SystemUpdateMutation {
	return suu.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (suu *SystemUpdateUpdate) ClearOwner() *SystemUpdateUpdate {
	suu.mutation.ClearOwner()
	return suu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (suu *SystemUpdateUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, suu.sqlSave, suu.mutation, suu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suu *SystemUpdateUpdate) SaveX(ctx context.Context) int {
	affected, err := suu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (suu *SystemUpdateUpdate) Exec(ctx context.Context) error {
	_, err := suu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suu *SystemUpdateUpdate) ExecX(ctx context.Context) {
	if err := suu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suu *SystemUpdateUpdate) check() error {
	if suu.mutation.OwnerCleared() && len(suu.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "SystemUpdate.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suu *SystemUpdateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SystemUpdateUpdate {
	suu.modifiers = append(suu.modifiers, modifiers...)
	return suu
}

func (suu *SystemUpdateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := suu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(systemupdate.Table, systemupdate.Columns, sqlgraph.NewFieldSpec(systemupdate.FieldID, field.TypeInt))
	if ps := suu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suu.mutation.SystemUpdateStatus(); ok {
		_spec.SetField(systemupdate.FieldSystemUpdateStatus, field.TypeString, value)
	}
	if value, ok := suu.mutation.LastInstall(); ok {
		_spec.SetField(systemupdate.FieldLastInstall, field.TypeTime, value)
	}
	if value, ok := suu.mutation.LastSearch(); ok {
		_spec.SetField(systemupdate.FieldLastSearch, field.TypeTime, value)
	}
	if value, ok := suu.mutation.PendingUpdates(); ok {
		_spec.SetField(systemupdate.FieldPendingUpdates, field.TypeBool, value)
	}
	if suu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   systemupdate.OwnerTable,
			Columns: []string{systemupdate.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   systemupdate.OwnerTable,
			Columns: []string{systemupdate.OwnerColumn},
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
	_spec.AddModifiers(suu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, suu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemupdate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	suu.mutation.done = true
	return n, nil
}

// SystemUpdateUpdateOne is the builder for updating a single SystemUpdate entity.
type SystemUpdateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SystemUpdateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetSystemUpdateStatus sets the "system_update_status" field.
func (suuo *SystemUpdateUpdateOne) SetSystemUpdateStatus(s string) *SystemUpdateUpdateOne {
	suuo.mutation.SetSystemUpdateStatus(s)
	return suuo
}

// SetNillableSystemUpdateStatus sets the "system_update_status" field if the given value is not nil.
func (suuo *SystemUpdateUpdateOne) SetNillableSystemUpdateStatus(s *string) *SystemUpdateUpdateOne {
	if s != nil {
		suuo.SetSystemUpdateStatus(*s)
	}
	return suuo
}

// SetLastInstall sets the "last_install" field.
func (suuo *SystemUpdateUpdateOne) SetLastInstall(t time.Time) *SystemUpdateUpdateOne {
	suuo.mutation.SetLastInstall(t)
	return suuo
}

// SetNillableLastInstall sets the "last_install" field if the given value is not nil.
func (suuo *SystemUpdateUpdateOne) SetNillableLastInstall(t *time.Time) *SystemUpdateUpdateOne {
	if t != nil {
		suuo.SetLastInstall(*t)
	}
	return suuo
}

// SetLastSearch sets the "last_search" field.
func (suuo *SystemUpdateUpdateOne) SetLastSearch(t time.Time) *SystemUpdateUpdateOne {
	suuo.mutation.SetLastSearch(t)
	return suuo
}

// SetNillableLastSearch sets the "last_search" field if the given value is not nil.
func (suuo *SystemUpdateUpdateOne) SetNillableLastSearch(t *time.Time) *SystemUpdateUpdateOne {
	if t != nil {
		suuo.SetLastSearch(*t)
	}
	return suuo
}

// SetPendingUpdates sets the "pending_updates" field.
func (suuo *SystemUpdateUpdateOne) SetPendingUpdates(b bool) *SystemUpdateUpdateOne {
	suuo.mutation.SetPendingUpdates(b)
	return suuo
}

// SetNillablePendingUpdates sets the "pending_updates" field if the given value is not nil.
func (suuo *SystemUpdateUpdateOne) SetNillablePendingUpdates(b *bool) *SystemUpdateUpdateOne {
	if b != nil {
		suuo.SetPendingUpdates(*b)
	}
	return suuo
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (suuo *SystemUpdateUpdateOne) SetOwnerID(id string) *SystemUpdateUpdateOne {
	suuo.mutation.SetOwnerID(id)
	return suuo
}

// SetOwner sets the "owner" edge to the Agent entity.
func (suuo *SystemUpdateUpdateOne) SetOwner(a *Agent) *SystemUpdateUpdateOne {
	return suuo.SetOwnerID(a.ID)
}

// Mutation returns the SystemUpdateMutation object of the builder.
func (suuo *SystemUpdateUpdateOne) Mutation() *SystemUpdateMutation {
	return suuo.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (suuo *SystemUpdateUpdateOne) ClearOwner() *SystemUpdateUpdateOne {
	suuo.mutation.ClearOwner()
	return suuo
}

// Where appends a list predicates to the SystemUpdateUpdate builder.
func (suuo *SystemUpdateUpdateOne) Where(ps ...predicate.SystemUpdate) *SystemUpdateUpdateOne {
	suuo.mutation.Where(ps...)
	return suuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suuo *SystemUpdateUpdateOne) Select(field string, fields ...string) *SystemUpdateUpdateOne {
	suuo.fields = append([]string{field}, fields...)
	return suuo
}

// Save executes the query and returns the updated SystemUpdate entity.
func (suuo *SystemUpdateUpdateOne) Save(ctx context.Context) (*SystemUpdate, error) {
	return withHooks(ctx, suuo.sqlSave, suuo.mutation, suuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suuo *SystemUpdateUpdateOne) SaveX(ctx context.Context) *SystemUpdate {
	node, err := suuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suuo *SystemUpdateUpdateOne) Exec(ctx context.Context) error {
	_, err := suuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suuo *SystemUpdateUpdateOne) ExecX(ctx context.Context) {
	if err := suuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suuo *SystemUpdateUpdateOne) check() error {
	if suuo.mutation.OwnerCleared() && len(suuo.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "SystemUpdate.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suuo *SystemUpdateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SystemUpdateUpdateOne {
	suuo.modifiers = append(suuo.modifiers, modifiers...)
	return suuo
}

func (suuo *SystemUpdateUpdateOne) sqlSave(ctx context.Context) (_node *SystemUpdate, err error) {
	if err := suuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(systemupdate.Table, systemupdate.Columns, sqlgraph.NewFieldSpec(systemupdate.FieldID, field.TypeInt))
	id, ok := suuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SystemUpdate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemupdate.FieldID)
		for _, f := range fields {
			if !systemupdate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != systemupdate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suuo.mutation.SystemUpdateStatus(); ok {
		_spec.SetField(systemupdate.FieldSystemUpdateStatus, field.TypeString, value)
	}
	if value, ok := suuo.mutation.LastInstall(); ok {
		_spec.SetField(systemupdate.FieldLastInstall, field.TypeTime, value)
	}
	if value, ok := suuo.mutation.LastSearch(); ok {
		_spec.SetField(systemupdate.FieldLastSearch, field.TypeTime, value)
	}
	if value, ok := suuo.mutation.PendingUpdates(); ok {
		_spec.SetField(systemupdate.FieldPendingUpdates, field.TypeBool, value)
	}
	if suuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   systemupdate.OwnerTable,
			Columns: []string{systemupdate.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   systemupdate.OwnerTable,
			Columns: []string{systemupdate.OwnerColumn},
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
	_spec.AddModifiers(suuo.modifiers...)
	_node = &SystemUpdate{config: suuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemupdate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suuo.mutation.done = true
	return _node, nil
}
