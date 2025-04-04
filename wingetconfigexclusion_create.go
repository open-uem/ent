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
	"github.com/open-uem/ent/wingetconfigexclusion"
)

// WingetConfigExclusionCreate is the builder for creating a WingetConfigExclusion entity.
type WingetConfigExclusionCreate struct {
	config
	mutation *WingetConfigExclusionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPackageID sets the "package_id" field.
func (wcec *WingetConfigExclusionCreate) SetPackageID(s string) *WingetConfigExclusionCreate {
	wcec.mutation.SetPackageID(s)
	return wcec
}

// SetWhen sets the "when" field.
func (wcec *WingetConfigExclusionCreate) SetWhen(t time.Time) *WingetConfigExclusionCreate {
	wcec.mutation.SetWhen(t)
	return wcec
}

// SetNillableWhen sets the "when" field if the given value is not nil.
func (wcec *WingetConfigExclusionCreate) SetNillableWhen(t *time.Time) *WingetConfigExclusionCreate {
	if t != nil {
		wcec.SetWhen(*t)
	}
	return wcec
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (wcec *WingetConfigExclusionCreate) SetOwnerID(id string) *WingetConfigExclusionCreate {
	wcec.mutation.SetOwnerID(id)
	return wcec
}

// SetOwner sets the "owner" edge to the Agent entity.
func (wcec *WingetConfigExclusionCreate) SetOwner(a *Agent) *WingetConfigExclusionCreate {
	return wcec.SetOwnerID(a.ID)
}

// Mutation returns the WingetConfigExclusionMutation object of the builder.
func (wcec *WingetConfigExclusionCreate) Mutation() *WingetConfigExclusionMutation {
	return wcec.mutation
}

// Save creates the WingetConfigExclusion in the database.
func (wcec *WingetConfigExclusionCreate) Save(ctx context.Context) (*WingetConfigExclusion, error) {
	wcec.defaults()
	return withHooks(ctx, wcec.sqlSave, wcec.mutation, wcec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wcec *WingetConfigExclusionCreate) SaveX(ctx context.Context) *WingetConfigExclusion {
	v, err := wcec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcec *WingetConfigExclusionCreate) Exec(ctx context.Context) error {
	_, err := wcec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcec *WingetConfigExclusionCreate) ExecX(ctx context.Context) {
	if err := wcec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcec *WingetConfigExclusionCreate) defaults() {
	if _, ok := wcec.mutation.When(); !ok {
		v := wingetconfigexclusion.DefaultWhen()
		wcec.mutation.SetWhen(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcec *WingetConfigExclusionCreate) check() error {
	if _, ok := wcec.mutation.PackageID(); !ok {
		return &ValidationError{Name: "package_id", err: errors.New(`ent: missing required field "WingetConfigExclusion.package_id"`)}
	}
	if len(wcec.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "WingetConfigExclusion.owner"`)}
	}
	return nil
}

func (wcec *WingetConfigExclusionCreate) sqlSave(ctx context.Context) (*WingetConfigExclusion, error) {
	if err := wcec.check(); err != nil {
		return nil, err
	}
	_node, _spec := wcec.createSpec()
	if err := sqlgraph.CreateNode(ctx, wcec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wcec.mutation.id = &_node.ID
	wcec.mutation.done = true
	return _node, nil
}

func (wcec *WingetConfigExclusionCreate) createSpec() (*WingetConfigExclusion, *sqlgraph.CreateSpec) {
	var (
		_node = &WingetConfigExclusion{config: wcec.config}
		_spec = sqlgraph.NewCreateSpec(wingetconfigexclusion.Table, sqlgraph.NewFieldSpec(wingetconfigexclusion.FieldID, field.TypeInt))
	)
	_spec.OnConflict = wcec.conflict
	if value, ok := wcec.mutation.PackageID(); ok {
		_spec.SetField(wingetconfigexclusion.FieldPackageID, field.TypeString, value)
		_node.PackageID = value
	}
	if value, ok := wcec.mutation.When(); ok {
		_spec.SetField(wingetconfigexclusion.FieldWhen, field.TypeTime, value)
		_node.When = value
	}
	if nodes := wcec.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wingetconfigexclusion.OwnerTable,
			Columns: []string{wingetconfigexclusion.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.agent_wingetcfgexclusions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.WingetConfigExclusion.Create().
//		SetPackageID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WingetConfigExclusionUpsert) {
//			SetPackageID(v+v).
//		}).
//		Exec(ctx)
func (wcec *WingetConfigExclusionCreate) OnConflict(opts ...sql.ConflictOption) *WingetConfigExclusionUpsertOne {
	wcec.conflict = opts
	return &WingetConfigExclusionUpsertOne{
		create: wcec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.WingetConfigExclusion.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wcec *WingetConfigExclusionCreate) OnConflictColumns(columns ...string) *WingetConfigExclusionUpsertOne {
	wcec.conflict = append(wcec.conflict, sql.ConflictColumns(columns...))
	return &WingetConfigExclusionUpsertOne{
		create: wcec,
	}
}

type (
	// WingetConfigExclusionUpsertOne is the builder for "upsert"-ing
	//  one WingetConfigExclusion node.
	WingetConfigExclusionUpsertOne struct {
		create *WingetConfigExclusionCreate
	}

	// WingetConfigExclusionUpsert is the "OnConflict" setter.
	WingetConfigExclusionUpsert struct {
		*sql.UpdateSet
	}
)

// SetPackageID sets the "package_id" field.
func (u *WingetConfigExclusionUpsert) SetPackageID(v string) *WingetConfigExclusionUpsert {
	u.Set(wingetconfigexclusion.FieldPackageID, v)
	return u
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *WingetConfigExclusionUpsert) UpdatePackageID() *WingetConfigExclusionUpsert {
	u.SetExcluded(wingetconfigexclusion.FieldPackageID)
	return u
}

// SetWhen sets the "when" field.
func (u *WingetConfigExclusionUpsert) SetWhen(v time.Time) *WingetConfigExclusionUpsert {
	u.Set(wingetconfigexclusion.FieldWhen, v)
	return u
}

// UpdateWhen sets the "when" field to the value that was provided on create.
func (u *WingetConfigExclusionUpsert) UpdateWhen() *WingetConfigExclusionUpsert {
	u.SetExcluded(wingetconfigexclusion.FieldWhen)
	return u
}

// ClearWhen clears the value of the "when" field.
func (u *WingetConfigExclusionUpsert) ClearWhen() *WingetConfigExclusionUpsert {
	u.SetNull(wingetconfigexclusion.FieldWhen)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.WingetConfigExclusion.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *WingetConfigExclusionUpsertOne) UpdateNewValues() *WingetConfigExclusionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.WingetConfigExclusion.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *WingetConfigExclusionUpsertOne) Ignore() *WingetConfigExclusionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WingetConfigExclusionUpsertOne) DoNothing() *WingetConfigExclusionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WingetConfigExclusionCreate.OnConflict
// documentation for more info.
func (u *WingetConfigExclusionUpsertOne) Update(set func(*WingetConfigExclusionUpsert)) *WingetConfigExclusionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WingetConfigExclusionUpsert{UpdateSet: update})
	}))
	return u
}

// SetPackageID sets the "package_id" field.
func (u *WingetConfigExclusionUpsertOne) SetPackageID(v string) *WingetConfigExclusionUpsertOne {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.SetPackageID(v)
	})
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *WingetConfigExclusionUpsertOne) UpdatePackageID() *WingetConfigExclusionUpsertOne {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.UpdatePackageID()
	})
}

// SetWhen sets the "when" field.
func (u *WingetConfigExclusionUpsertOne) SetWhen(v time.Time) *WingetConfigExclusionUpsertOne {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.SetWhen(v)
	})
}

// UpdateWhen sets the "when" field to the value that was provided on create.
func (u *WingetConfigExclusionUpsertOne) UpdateWhen() *WingetConfigExclusionUpsertOne {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.UpdateWhen()
	})
}

// ClearWhen clears the value of the "when" field.
func (u *WingetConfigExclusionUpsertOne) ClearWhen() *WingetConfigExclusionUpsertOne {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.ClearWhen()
	})
}

// Exec executes the query.
func (u *WingetConfigExclusionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WingetConfigExclusionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WingetConfigExclusionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *WingetConfigExclusionUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *WingetConfigExclusionUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// WingetConfigExclusionCreateBulk is the builder for creating many WingetConfigExclusion entities in bulk.
type WingetConfigExclusionCreateBulk struct {
	config
	err      error
	builders []*WingetConfigExclusionCreate
	conflict []sql.ConflictOption
}

// Save creates the WingetConfigExclusion entities in the database.
func (wcecb *WingetConfigExclusionCreateBulk) Save(ctx context.Context) ([]*WingetConfigExclusion, error) {
	if wcecb.err != nil {
		return nil, wcecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcecb.builders))
	nodes := make([]*WingetConfigExclusion, len(wcecb.builders))
	mutators := make([]Mutator, len(wcecb.builders))
	for i := range wcecb.builders {
		func(i int, root context.Context) {
			builder := wcecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WingetConfigExclusionMutation)
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
					_, err = mutators[i+1].Mutate(root, wcecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = wcecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcecb *WingetConfigExclusionCreateBulk) SaveX(ctx context.Context) []*WingetConfigExclusion {
	v, err := wcecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcecb *WingetConfigExclusionCreateBulk) Exec(ctx context.Context) error {
	_, err := wcecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcecb *WingetConfigExclusionCreateBulk) ExecX(ctx context.Context) {
	if err := wcecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.WingetConfigExclusion.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WingetConfigExclusionUpsert) {
//			SetPackageID(v+v).
//		}).
//		Exec(ctx)
func (wcecb *WingetConfigExclusionCreateBulk) OnConflict(opts ...sql.ConflictOption) *WingetConfigExclusionUpsertBulk {
	wcecb.conflict = opts
	return &WingetConfigExclusionUpsertBulk{
		create: wcecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.WingetConfigExclusion.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wcecb *WingetConfigExclusionCreateBulk) OnConflictColumns(columns ...string) *WingetConfigExclusionUpsertBulk {
	wcecb.conflict = append(wcecb.conflict, sql.ConflictColumns(columns...))
	return &WingetConfigExclusionUpsertBulk{
		create: wcecb,
	}
}

// WingetConfigExclusionUpsertBulk is the builder for "upsert"-ing
// a bulk of WingetConfigExclusion nodes.
type WingetConfigExclusionUpsertBulk struct {
	create *WingetConfigExclusionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.WingetConfigExclusion.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *WingetConfigExclusionUpsertBulk) UpdateNewValues() *WingetConfigExclusionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.WingetConfigExclusion.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *WingetConfigExclusionUpsertBulk) Ignore() *WingetConfigExclusionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WingetConfigExclusionUpsertBulk) DoNothing() *WingetConfigExclusionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WingetConfigExclusionCreateBulk.OnConflict
// documentation for more info.
func (u *WingetConfigExclusionUpsertBulk) Update(set func(*WingetConfigExclusionUpsert)) *WingetConfigExclusionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WingetConfigExclusionUpsert{UpdateSet: update})
	}))
	return u
}

// SetPackageID sets the "package_id" field.
func (u *WingetConfigExclusionUpsertBulk) SetPackageID(v string) *WingetConfigExclusionUpsertBulk {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.SetPackageID(v)
	})
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *WingetConfigExclusionUpsertBulk) UpdatePackageID() *WingetConfigExclusionUpsertBulk {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.UpdatePackageID()
	})
}

// SetWhen sets the "when" field.
func (u *WingetConfigExclusionUpsertBulk) SetWhen(v time.Time) *WingetConfigExclusionUpsertBulk {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.SetWhen(v)
	})
}

// UpdateWhen sets the "when" field to the value that was provided on create.
func (u *WingetConfigExclusionUpsertBulk) UpdateWhen() *WingetConfigExclusionUpsertBulk {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.UpdateWhen()
	})
}

// ClearWhen clears the value of the "when" field.
func (u *WingetConfigExclusionUpsertBulk) ClearWhen() *WingetConfigExclusionUpsertBulk {
	return u.Update(func(s *WingetConfigExclusionUpsert) {
		s.ClearWhen()
	})
}

// Exec executes the query.
func (u *WingetConfigExclusionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the WingetConfigExclusionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WingetConfigExclusionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WingetConfigExclusionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
