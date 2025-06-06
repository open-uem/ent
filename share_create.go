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
	"github.com/open-uem/ent/share"
)

// ShareCreate is the builder for creating a Share entity.
type ShareCreate struct {
	config
	mutation *ShareMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (sc *ShareCreate) SetName(s string) *ShareCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *ShareCreate) SetDescription(s string) *ShareCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetPath sets the "path" field.
func (sc *ShareCreate) SetPath(s string) *ShareCreate {
	sc.mutation.SetPath(s)
	return sc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (sc *ShareCreate) SetNillablePath(s *string) *ShareCreate {
	if s != nil {
		sc.SetPath(*s)
	}
	return sc
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (sc *ShareCreate) SetOwnerID(id string) *ShareCreate {
	sc.mutation.SetOwnerID(id)
	return sc
}

// SetOwner sets the "owner" edge to the Agent entity.
func (sc *ShareCreate) SetOwner(a *Agent) *ShareCreate {
	return sc.SetOwnerID(a.ID)
}

// Mutation returns the ShareMutation object of the builder.
func (sc *ShareCreate) Mutation() *ShareMutation {
	return sc.mutation
}

// Save creates the Share in the database.
func (sc *ShareCreate) Save(ctx context.Context) (*Share, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShareCreate) SaveX(ctx context.Context) *Share {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ShareCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ShareCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShareCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Share.name"`)}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Share.description"`)}
	}
	if len(sc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Share.owner"`)}
	}
	return nil
}

func (sc *ShareCreate) sqlSave(ctx context.Context) (*Share, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ShareCreate) createSpec() (*Share, *sqlgraph.CreateSpec) {
	var (
		_node = &Share{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(share.Table, sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(share.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(share.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.Path(); ok {
		_spec.SetField(share.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   share.OwnerTable,
			Columns: []string{share.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.agent_shares = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Share.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShareUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sc *ShareCreate) OnConflict(opts ...sql.ConflictOption) *ShareUpsertOne {
	sc.conflict = opts
	return &ShareUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Share.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *ShareCreate) OnConflictColumns(columns ...string) *ShareUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &ShareUpsertOne{
		create: sc,
	}
}

type (
	// ShareUpsertOne is the builder for "upsert"-ing
	//  one Share node.
	ShareUpsertOne struct {
		create *ShareCreate
	}

	// ShareUpsert is the "OnConflict" setter.
	ShareUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *ShareUpsert) SetName(v string) *ShareUpsert {
	u.Set(share.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ShareUpsert) UpdateName() *ShareUpsert {
	u.SetExcluded(share.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *ShareUpsert) SetDescription(v string) *ShareUpsert {
	u.Set(share.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ShareUpsert) UpdateDescription() *ShareUpsert {
	u.SetExcluded(share.FieldDescription)
	return u
}

// SetPath sets the "path" field.
func (u *ShareUpsert) SetPath(v string) *ShareUpsert {
	u.Set(share.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ShareUpsert) UpdatePath() *ShareUpsert {
	u.SetExcluded(share.FieldPath)
	return u
}

// ClearPath clears the value of the "path" field.
func (u *ShareUpsert) ClearPath() *ShareUpsert {
	u.SetNull(share.FieldPath)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Share.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ShareUpsertOne) UpdateNewValues() *ShareUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Share.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ShareUpsertOne) Ignore() *ShareUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShareUpsertOne) DoNothing() *ShareUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShareCreate.OnConflict
// documentation for more info.
func (u *ShareUpsertOne) Update(set func(*ShareUpsert)) *ShareUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShareUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ShareUpsertOne) SetName(v string) *ShareUpsertOne {
	return u.Update(func(s *ShareUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ShareUpsertOne) UpdateName() *ShareUpsertOne {
	return u.Update(func(s *ShareUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *ShareUpsertOne) SetDescription(v string) *ShareUpsertOne {
	return u.Update(func(s *ShareUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ShareUpsertOne) UpdateDescription() *ShareUpsertOne {
	return u.Update(func(s *ShareUpsert) {
		s.UpdateDescription()
	})
}

// SetPath sets the "path" field.
func (u *ShareUpsertOne) SetPath(v string) *ShareUpsertOne {
	return u.Update(func(s *ShareUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ShareUpsertOne) UpdatePath() *ShareUpsertOne {
	return u.Update(func(s *ShareUpsert) {
		s.UpdatePath()
	})
}

// ClearPath clears the value of the "path" field.
func (u *ShareUpsertOne) ClearPath() *ShareUpsertOne {
	return u.Update(func(s *ShareUpsert) {
		s.ClearPath()
	})
}

// Exec executes the query.
func (u *ShareUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ShareCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShareUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ShareUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ShareUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ShareCreateBulk is the builder for creating many Share entities in bulk.
type ShareCreateBulk struct {
	config
	err      error
	builders []*ShareCreate
	conflict []sql.ConflictOption
}

// Save creates the Share entities in the database.
func (scb *ShareCreateBulk) Save(ctx context.Context) ([]*Share, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Share, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShareMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShareCreateBulk) SaveX(ctx context.Context) []*Share {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ShareCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ShareCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Share.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShareUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (scb *ShareCreateBulk) OnConflict(opts ...sql.ConflictOption) *ShareUpsertBulk {
	scb.conflict = opts
	return &ShareUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Share.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *ShareCreateBulk) OnConflictColumns(columns ...string) *ShareUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &ShareUpsertBulk{
		create: scb,
	}
}

// ShareUpsertBulk is the builder for "upsert"-ing
// a bulk of Share nodes.
type ShareUpsertBulk struct {
	create *ShareCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Share.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ShareUpsertBulk) UpdateNewValues() *ShareUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Share.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ShareUpsertBulk) Ignore() *ShareUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShareUpsertBulk) DoNothing() *ShareUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShareCreateBulk.OnConflict
// documentation for more info.
func (u *ShareUpsertBulk) Update(set func(*ShareUpsert)) *ShareUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShareUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ShareUpsertBulk) SetName(v string) *ShareUpsertBulk {
	return u.Update(func(s *ShareUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ShareUpsertBulk) UpdateName() *ShareUpsertBulk {
	return u.Update(func(s *ShareUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *ShareUpsertBulk) SetDescription(v string) *ShareUpsertBulk {
	return u.Update(func(s *ShareUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ShareUpsertBulk) UpdateDescription() *ShareUpsertBulk {
	return u.Update(func(s *ShareUpsert) {
		s.UpdateDescription()
	})
}

// SetPath sets the "path" field.
func (u *ShareUpsertBulk) SetPath(v string) *ShareUpsertBulk {
	return u.Update(func(s *ShareUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ShareUpsertBulk) UpdatePath() *ShareUpsertBulk {
	return u.Update(func(s *ShareUpsert) {
		s.UpdatePath()
	})
}

// ClearPath clears the value of the "path" field.
func (u *ShareUpsertBulk) ClearPath() *ShareUpsertBulk {
	return u.Update(func(s *ShareUpsert) {
		s.ClearPath()
	})
}

// Exec executes the query.
func (u *ShareUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ShareCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ShareCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShareUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
