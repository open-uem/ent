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
)

// AntivirusCreate is the builder for creating a Antivirus entity.
type AntivirusCreate struct {
	config
	mutation *AntivirusMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ac *AntivirusCreate) SetName(s string) *AntivirusCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetIsActive sets the "is_active" field.
func (ac *AntivirusCreate) SetIsActive(b bool) *AntivirusCreate {
	ac.mutation.SetIsActive(b)
	return ac
}

// SetIsUpdated sets the "is_updated" field.
func (ac *AntivirusCreate) SetIsUpdated(b bool) *AntivirusCreate {
	ac.mutation.SetIsUpdated(b)
	return ac
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (ac *AntivirusCreate) SetOwnerID(id string) *AntivirusCreate {
	ac.mutation.SetOwnerID(id)
	return ac
}

// SetOwner sets the "owner" edge to the Agent entity.
func (ac *AntivirusCreate) SetOwner(a *Agent) *AntivirusCreate {
	return ac.SetOwnerID(a.ID)
}

// Mutation returns the AntivirusMutation object of the builder.
func (ac *AntivirusCreate) Mutation() *AntivirusMutation {
	return ac.mutation
}

// Save creates the Antivirus in the database.
func (ac *AntivirusCreate) Save(ctx context.Context) (*Antivirus, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AntivirusCreate) SaveX(ctx context.Context) *Antivirus {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AntivirusCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AntivirusCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AntivirusCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Antivirus.name"`)}
	}
	if _, ok := ac.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "Antivirus.is_active"`)}
	}
	if _, ok := ac.mutation.IsUpdated(); !ok {
		return &ValidationError{Name: "is_updated", err: errors.New(`ent: missing required field "Antivirus.is_updated"`)}
	}
	if len(ac.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Antivirus.owner"`)}
	}
	return nil
}

func (ac *AntivirusCreate) sqlSave(ctx context.Context) (*Antivirus, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AntivirusCreate) createSpec() (*Antivirus, *sqlgraph.CreateSpec) {
	var (
		_node = &Antivirus{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(antivirus.Table, sqlgraph.NewFieldSpec(antivirus.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(antivirus.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.IsActive(); ok {
		_spec.SetField(antivirus.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := ac.mutation.IsUpdated(); ok {
		_spec.SetField(antivirus.FieldIsUpdated, field.TypeBool, value)
		_node.IsUpdated = value
	}
	if nodes := ac.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_node.agent_antivirus = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Antivirus.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AntivirusUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ac *AntivirusCreate) OnConflict(opts ...sql.ConflictOption) *AntivirusUpsertOne {
	ac.conflict = opts
	return &AntivirusUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Antivirus.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AntivirusCreate) OnConflictColumns(columns ...string) *AntivirusUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AntivirusUpsertOne{
		create: ac,
	}
}

type (
	// AntivirusUpsertOne is the builder for "upsert"-ing
	//  one Antivirus node.
	AntivirusUpsertOne struct {
		create *AntivirusCreate
	}

	// AntivirusUpsert is the "OnConflict" setter.
	AntivirusUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *AntivirusUpsert) SetName(v string) *AntivirusUpsert {
	u.Set(antivirus.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AntivirusUpsert) UpdateName() *AntivirusUpsert {
	u.SetExcluded(antivirus.FieldName)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *AntivirusUpsert) SetIsActive(v bool) *AntivirusUpsert {
	u.Set(antivirus.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *AntivirusUpsert) UpdateIsActive() *AntivirusUpsert {
	u.SetExcluded(antivirus.FieldIsActive)
	return u
}

// SetIsUpdated sets the "is_updated" field.
func (u *AntivirusUpsert) SetIsUpdated(v bool) *AntivirusUpsert {
	u.Set(antivirus.FieldIsUpdated, v)
	return u
}

// UpdateIsUpdated sets the "is_updated" field to the value that was provided on create.
func (u *AntivirusUpsert) UpdateIsUpdated() *AntivirusUpsert {
	u.SetExcluded(antivirus.FieldIsUpdated)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Antivirus.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AntivirusUpsertOne) UpdateNewValues() *AntivirusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Antivirus.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AntivirusUpsertOne) Ignore() *AntivirusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AntivirusUpsertOne) DoNothing() *AntivirusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AntivirusCreate.OnConflict
// documentation for more info.
func (u *AntivirusUpsertOne) Update(set func(*AntivirusUpsert)) *AntivirusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AntivirusUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AntivirusUpsertOne) SetName(v string) *AntivirusUpsertOne {
	return u.Update(func(s *AntivirusUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AntivirusUpsertOne) UpdateName() *AntivirusUpsertOne {
	return u.Update(func(s *AntivirusUpsert) {
		s.UpdateName()
	})
}

// SetIsActive sets the "is_active" field.
func (u *AntivirusUpsertOne) SetIsActive(v bool) *AntivirusUpsertOne {
	return u.Update(func(s *AntivirusUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *AntivirusUpsertOne) UpdateIsActive() *AntivirusUpsertOne {
	return u.Update(func(s *AntivirusUpsert) {
		s.UpdateIsActive()
	})
}

// SetIsUpdated sets the "is_updated" field.
func (u *AntivirusUpsertOne) SetIsUpdated(v bool) *AntivirusUpsertOne {
	return u.Update(func(s *AntivirusUpsert) {
		s.SetIsUpdated(v)
	})
}

// UpdateIsUpdated sets the "is_updated" field to the value that was provided on create.
func (u *AntivirusUpsertOne) UpdateIsUpdated() *AntivirusUpsertOne {
	return u.Update(func(s *AntivirusUpsert) {
		s.UpdateIsUpdated()
	})
}

// Exec executes the query.
func (u *AntivirusUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AntivirusCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AntivirusUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AntivirusUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AntivirusUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AntivirusCreateBulk is the builder for creating many Antivirus entities in bulk.
type AntivirusCreateBulk struct {
	config
	err      error
	builders []*AntivirusCreate
	conflict []sql.ConflictOption
}

// Save creates the Antivirus entities in the database.
func (acb *AntivirusCreateBulk) Save(ctx context.Context) ([]*Antivirus, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Antivirus, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AntivirusMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AntivirusCreateBulk) SaveX(ctx context.Context) []*Antivirus {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AntivirusCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AntivirusCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Antivirus.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AntivirusUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (acb *AntivirusCreateBulk) OnConflict(opts ...sql.ConflictOption) *AntivirusUpsertBulk {
	acb.conflict = opts
	return &AntivirusUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Antivirus.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AntivirusCreateBulk) OnConflictColumns(columns ...string) *AntivirusUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AntivirusUpsertBulk{
		create: acb,
	}
}

// AntivirusUpsertBulk is the builder for "upsert"-ing
// a bulk of Antivirus nodes.
type AntivirusUpsertBulk struct {
	create *AntivirusCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Antivirus.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AntivirusUpsertBulk) UpdateNewValues() *AntivirusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Antivirus.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AntivirusUpsertBulk) Ignore() *AntivirusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AntivirusUpsertBulk) DoNothing() *AntivirusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AntivirusCreateBulk.OnConflict
// documentation for more info.
func (u *AntivirusUpsertBulk) Update(set func(*AntivirusUpsert)) *AntivirusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AntivirusUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AntivirusUpsertBulk) SetName(v string) *AntivirusUpsertBulk {
	return u.Update(func(s *AntivirusUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AntivirusUpsertBulk) UpdateName() *AntivirusUpsertBulk {
	return u.Update(func(s *AntivirusUpsert) {
		s.UpdateName()
	})
}

// SetIsActive sets the "is_active" field.
func (u *AntivirusUpsertBulk) SetIsActive(v bool) *AntivirusUpsertBulk {
	return u.Update(func(s *AntivirusUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *AntivirusUpsertBulk) UpdateIsActive() *AntivirusUpsertBulk {
	return u.Update(func(s *AntivirusUpsert) {
		s.UpdateIsActive()
	})
}

// SetIsUpdated sets the "is_updated" field.
func (u *AntivirusUpsertBulk) SetIsUpdated(v bool) *AntivirusUpsertBulk {
	return u.Update(func(s *AntivirusUpsert) {
		s.SetIsUpdated(v)
	})
}

// UpdateIsUpdated sets the "is_updated" field to the value that was provided on create.
func (u *AntivirusUpsertBulk) UpdateIsUpdated() *AntivirusUpsertBulk {
	return u.Update(func(s *AntivirusUpsert) {
		s.UpdateIsUpdated()
	})
}

// Exec executes the query.
func (u *AntivirusUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AntivirusCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AntivirusCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AntivirusUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
