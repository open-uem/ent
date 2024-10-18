// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/metadata"
	"github.com/doncicuto/openuem_ent/orgmetadata"
	"github.com/doncicuto/openuem_ent/predicate"
)

// MetadataUpdate is the builder for updating Metadata entities.
type MetadataUpdate struct {
	config
	hooks     []Hook
	mutation  *MetadataMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MetadataUpdate builder.
func (mu *MetadataUpdate) Where(ps ...predicate.Metadata) *MetadataUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetValue sets the "value" field.
func (mu *MetadataUpdate) SetValue(s string) *MetadataUpdate {
	mu.mutation.SetValue(s)
	return mu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableValue(s *string) *MetadataUpdate {
	if s != nil {
		mu.SetValue(*s)
	}
	return mu
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (mu *MetadataUpdate) SetOwnerID(id string) *MetadataUpdate {
	mu.mutation.SetOwnerID(id)
	return mu
}

// SetOwner sets the "owner" edge to the Agent entity.
func (mu *MetadataUpdate) SetOwner(a *Agent) *MetadataUpdate {
	return mu.SetOwnerID(a.ID)
}

// SetOrgID sets the "org" edge to the OrgMetadata entity by ID.
func (mu *MetadataUpdate) SetOrgID(id int) *MetadataUpdate {
	mu.mutation.SetOrgID(id)
	return mu
}

// SetOrg sets the "org" edge to the OrgMetadata entity.
func (mu *MetadataUpdate) SetOrg(o *OrgMetadata) *MetadataUpdate {
	return mu.SetOrgID(o.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (mu *MetadataUpdate) Mutation() *MetadataMutation {
	return mu.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (mu *MetadataUpdate) ClearOwner() *MetadataUpdate {
	mu.mutation.ClearOwner()
	return mu
}

// ClearOrg clears the "org" edge to the OrgMetadata entity.
func (mu *MetadataUpdate) ClearOrg() *MetadataUpdate {
	mu.mutation.ClearOrg()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetadataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetadataUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetadataUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MetadataUpdate) check() error {
	if mu.mutation.OwnerCleared() && len(mu.mutation.OwnerIDs()) > 0 {
		return errors.New(`openuem_ent: clearing a required unique edge "Metadata.owner"`)
	}
	if mu.mutation.OrgCleared() && len(mu.mutation.OrgIDs()) > 0 {
		return errors.New(`openuem_ent: clearing a required unique edge "Metadata.org"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MetadataUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MetadataUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Value(); ok {
		_spec.SetField(metadata.FieldValue, field.TypeString, value)
	}
	if mu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.OwnerTable,
			Columns: []string{metadata.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.OwnerTable,
			Columns: []string{metadata.OwnerColumn},
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
	if mu.mutation.OrgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.OrgTable,
			Columns: []string{metadata.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgmetadata.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.OrgTable,
			Columns: []string{metadata.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgmetadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MetadataUpdateOne is the builder for updating a single Metadata entity.
type MetadataUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MetadataMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetValue sets the "value" field.
func (muo *MetadataUpdateOne) SetValue(s string) *MetadataUpdateOne {
	muo.mutation.SetValue(s)
	return muo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableValue(s *string) *MetadataUpdateOne {
	if s != nil {
		muo.SetValue(*s)
	}
	return muo
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (muo *MetadataUpdateOne) SetOwnerID(id string) *MetadataUpdateOne {
	muo.mutation.SetOwnerID(id)
	return muo
}

// SetOwner sets the "owner" edge to the Agent entity.
func (muo *MetadataUpdateOne) SetOwner(a *Agent) *MetadataUpdateOne {
	return muo.SetOwnerID(a.ID)
}

// SetOrgID sets the "org" edge to the OrgMetadata entity by ID.
func (muo *MetadataUpdateOne) SetOrgID(id int) *MetadataUpdateOne {
	muo.mutation.SetOrgID(id)
	return muo
}

// SetOrg sets the "org" edge to the OrgMetadata entity.
func (muo *MetadataUpdateOne) SetOrg(o *OrgMetadata) *MetadataUpdateOne {
	return muo.SetOrgID(o.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (muo *MetadataUpdateOne) Mutation() *MetadataMutation {
	return muo.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (muo *MetadataUpdateOne) ClearOwner() *MetadataUpdateOne {
	muo.mutation.ClearOwner()
	return muo
}

// ClearOrg clears the "org" edge to the OrgMetadata entity.
func (muo *MetadataUpdateOne) ClearOrg() *MetadataUpdateOne {
	muo.mutation.ClearOrg()
	return muo
}

// Where appends a list predicates to the MetadataUpdate builder.
func (muo *MetadataUpdateOne) Where(ps ...predicate.Metadata) *MetadataUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetadataUpdateOne) Select(field string, fields ...string) *MetadataUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Metadata entity.
func (muo *MetadataUpdateOne) Save(ctx context.Context) (*Metadata, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetadataUpdateOne) SaveX(ctx context.Context) *Metadata {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetadataUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MetadataUpdateOne) check() error {
	if muo.mutation.OwnerCleared() && len(muo.mutation.OwnerIDs()) > 0 {
		return errors.New(`openuem_ent: clearing a required unique edge "Metadata.owner"`)
	}
	if muo.mutation.OrgCleared() && len(muo.mutation.OrgIDs()) > 0 {
		return errors.New(`openuem_ent: clearing a required unique edge "Metadata.org"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MetadataUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MetadataUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MetadataUpdateOne) sqlSave(ctx context.Context) (_node *Metadata, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`openuem_ent: missing "Metadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadata.FieldID)
		for _, f := range fields {
			if !metadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
			}
			if f != metadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Value(); ok {
		_spec.SetField(metadata.FieldValue, field.TypeString, value)
	}
	if muo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.OwnerTable,
			Columns: []string{metadata.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.OwnerTable,
			Columns: []string{metadata.OwnerColumn},
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
	if muo.mutation.OrgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.OrgTable,
			Columns: []string{metadata.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgmetadata.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.OrgTable,
			Columns: []string{metadata.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgmetadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(muo.modifiers...)
	_node = &Metadata{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}