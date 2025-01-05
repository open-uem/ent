// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/antivirus"
	"github.com/open-uem/ent/predicate"
)

// AntivirusUpdate is the builder for updating Antivirus entities.
type AntivirusUpdate struct {
	config
	hooks     []Hook
	mutation  *AntivirusMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AntivirusUpdate builder.
func (au *AntivirusUpdate) Where(ps ...predicate.Antivirus) *AntivirusUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AntivirusUpdate) SetName(s string) *AntivirusUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AntivirusUpdate) SetNillableName(s *string) *AntivirusUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetIsActive sets the "is_active" field.
func (au *AntivirusUpdate) SetIsActive(b bool) *AntivirusUpdate {
	au.mutation.SetIsActive(b)
	return au
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (au *AntivirusUpdate) SetNillableIsActive(b *bool) *AntivirusUpdate {
	if b != nil {
		au.SetIsActive(*b)
	}
	return au
}

// SetIsUpdated sets the "is_updated" field.
func (au *AntivirusUpdate) SetIsUpdated(b bool) *AntivirusUpdate {
	au.mutation.SetIsUpdated(b)
	return au
}

// SetNillableIsUpdated sets the "is_updated" field if the given value is not nil.
func (au *AntivirusUpdate) SetNillableIsUpdated(b *bool) *AntivirusUpdate {
	if b != nil {
		au.SetIsUpdated(*b)
	}
	return au
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (au *AntivirusUpdate) SetOwnerID(id string) *AntivirusUpdate {
	au.mutation.SetOwnerID(id)
	return au
}

// SetOwner sets the "owner" edge to the Agent entity.
func (au *AntivirusUpdate) SetOwner(a *Agent) *AntivirusUpdate {
	return au.SetOwnerID(a.ID)
}

// Mutation returns the AntivirusMutation object of the builder.
func (au *AntivirusUpdate) Mutation() *AntivirusMutation {
	return au.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (au *AntivirusUpdate) ClearOwner() *AntivirusUpdate {
	au.mutation.ClearOwner()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AntivirusUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AntivirusUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AntivirusUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AntivirusUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AntivirusUpdate) check() error {
	if au.mutation.OwnerCleared() && len(au.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Antivirus.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AntivirusUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AntivirusUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AntivirusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(antivirus.Table, antivirus.Columns, sqlgraph.NewFieldSpec(antivirus.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(antivirus.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.IsActive(); ok {
		_spec.SetField(antivirus.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := au.mutation.IsUpdated(); ok {
		_spec.SetField(antivirus.FieldIsUpdated, field.TypeBool, value)
	}
	if au.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   antivirus.OwnerTable,
			Columns: []string{antivirus.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   antivirus.OwnerTable,
			Columns: []string{antivirus.OwnerColumn},
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
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{antivirus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AntivirusUpdateOne is the builder for updating a single Antivirus entity.
type AntivirusUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AntivirusMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (auo *AntivirusUpdateOne) SetName(s string) *AntivirusUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AntivirusUpdateOne) SetNillableName(s *string) *AntivirusUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetIsActive sets the "is_active" field.
func (auo *AntivirusUpdateOne) SetIsActive(b bool) *AntivirusUpdateOne {
	auo.mutation.SetIsActive(b)
	return auo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (auo *AntivirusUpdateOne) SetNillableIsActive(b *bool) *AntivirusUpdateOne {
	if b != nil {
		auo.SetIsActive(*b)
	}
	return auo
}

// SetIsUpdated sets the "is_updated" field.
func (auo *AntivirusUpdateOne) SetIsUpdated(b bool) *AntivirusUpdateOne {
	auo.mutation.SetIsUpdated(b)
	return auo
}

// SetNillableIsUpdated sets the "is_updated" field if the given value is not nil.
func (auo *AntivirusUpdateOne) SetNillableIsUpdated(b *bool) *AntivirusUpdateOne {
	if b != nil {
		auo.SetIsUpdated(*b)
	}
	return auo
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (auo *AntivirusUpdateOne) SetOwnerID(id string) *AntivirusUpdateOne {
	auo.mutation.SetOwnerID(id)
	return auo
}

// SetOwner sets the "owner" edge to the Agent entity.
func (auo *AntivirusUpdateOne) SetOwner(a *Agent) *AntivirusUpdateOne {
	return auo.SetOwnerID(a.ID)
}

// Mutation returns the AntivirusMutation object of the builder.
func (auo *AntivirusUpdateOne) Mutation() *AntivirusMutation {
	return auo.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (auo *AntivirusUpdateOne) ClearOwner() *AntivirusUpdateOne {
	auo.mutation.ClearOwner()
	return auo
}

// Where appends a list predicates to the AntivirusUpdate builder.
func (auo *AntivirusUpdateOne) Where(ps ...predicate.Antivirus) *AntivirusUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AntivirusUpdateOne) Select(field string, fields ...string) *AntivirusUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Antivirus entity.
func (auo *AntivirusUpdateOne) Save(ctx context.Context) (*Antivirus, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AntivirusUpdateOne) SaveX(ctx context.Context) *Antivirus {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AntivirusUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AntivirusUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AntivirusUpdateOne) check() error {
	if auo.mutation.OwnerCleared() && len(auo.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Antivirus.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AntivirusUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AntivirusUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AntivirusUpdateOne) sqlSave(ctx context.Context) (_node *Antivirus, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(antivirus.Table, antivirus.Columns, sqlgraph.NewFieldSpec(antivirus.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Antivirus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, antivirus.FieldID)
		for _, f := range fields {
			if !antivirus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != antivirus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(antivirus.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.IsActive(); ok {
		_spec.SetField(antivirus.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := auo.mutation.IsUpdated(); ok {
		_spec.SetField(antivirus.FieldIsUpdated, field.TypeBool, value)
	}
	if auo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   antivirus.OwnerTable,
			Columns: []string{antivirus.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   antivirus.OwnerTable,
			Columns: []string{antivirus.OwnerColumn},
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
	_spec.AddModifiers(auo.modifiers...)
	_node = &Antivirus{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{antivirus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
