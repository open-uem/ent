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
	"github.com/doncicuto/openuem_ent/sessions"
)

// SessionsCreate is the builder for creating a Sessions entity.
type SessionsCreate struct {
	config
	mutation *SessionsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetToken sets the "token" field.
func (sc *SessionsCreate) SetToken(s string) *SessionsCreate {
	sc.mutation.SetToken(s)
	return sc
}

// SetData sets the "data" field.
func (sc *SessionsCreate) SetData(b []byte) *SessionsCreate {
	sc.mutation.SetData(b)
	return sc
}

// SetExpiry sets the "expiry" field.
func (sc *SessionsCreate) SetExpiry(t time.Time) *SessionsCreate {
	sc.mutation.SetExpiry(t)
	return sc
}

// Mutation returns the SessionsMutation object of the builder.
func (sc *SessionsCreate) Mutation() *SessionsMutation {
	return sc.mutation
}

// Save creates the Sessions in the database.
func (sc *SessionsCreate) Save(ctx context.Context) (*Sessions, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionsCreate) SaveX(ctx context.Context) *Sessions {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionsCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionsCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionsCreate) check() error {
	if _, ok := sc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`openuem_ent: missing required field "Sessions.token"`)}
	}
	if v, ok := sc.mutation.Token(); ok {
		if err := sessions.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`openuem_ent: validator failed for field "Sessions.token": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`openuem_ent: missing required field "Sessions.data"`)}
	}
	if v, ok := sc.mutation.Data(); ok {
		if err := sessions.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`openuem_ent: validator failed for field "Sessions.data": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`openuem_ent: missing required field "Sessions.expiry"`)}
	}
	return nil
}

func (sc *SessionsCreate) sqlSave(ctx context.Context) (*Sessions, error) {
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

func (sc *SessionsCreate) createSpec() (*Sessions, *sqlgraph.CreateSpec) {
	var (
		_node = &Sessions{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(sessions.Table, sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.Token(); ok {
		_spec.SetField(sessions.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := sc.mutation.Data(); ok {
		_spec.SetField(sessions.FieldData, field.TypeBytes, value)
		_node.Data = value
	}
	if value, ok := sc.mutation.Expiry(); ok {
		_spec.SetField(sessions.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Sessions.Create().
//		SetToken(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SessionsUpsert) {
//			SetToken(v+v).
//		}).
//		Exec(ctx)
func (sc *SessionsCreate) OnConflict(opts ...sql.ConflictOption) *SessionsUpsertOne {
	sc.conflict = opts
	return &SessionsUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Sessions.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SessionsCreate) OnConflictColumns(columns ...string) *SessionsUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SessionsUpsertOne{
		create: sc,
	}
}

type (
	// SessionsUpsertOne is the builder for "upsert"-ing
	//  one Sessions node.
	SessionsUpsertOne struct {
		create *SessionsCreate
	}

	// SessionsUpsert is the "OnConflict" setter.
	SessionsUpsert struct {
		*sql.UpdateSet
	}
)

// SetToken sets the "token" field.
func (u *SessionsUpsert) SetToken(v string) *SessionsUpsert {
	u.Set(sessions.FieldToken, v)
	return u
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *SessionsUpsert) UpdateToken() *SessionsUpsert {
	u.SetExcluded(sessions.FieldToken)
	return u
}

// SetData sets the "data" field.
func (u *SessionsUpsert) SetData(v []byte) *SessionsUpsert {
	u.Set(sessions.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *SessionsUpsert) UpdateData() *SessionsUpsert {
	u.SetExcluded(sessions.FieldData)
	return u
}

// SetExpiry sets the "expiry" field.
func (u *SessionsUpsert) SetExpiry(v time.Time) *SessionsUpsert {
	u.Set(sessions.FieldExpiry, v)
	return u
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *SessionsUpsert) UpdateExpiry() *SessionsUpsert {
	u.SetExcluded(sessions.FieldExpiry)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Sessions.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SessionsUpsertOne) UpdateNewValues() *SessionsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Sessions.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SessionsUpsertOne) Ignore() *SessionsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SessionsUpsertOne) DoNothing() *SessionsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SessionsCreate.OnConflict
// documentation for more info.
func (u *SessionsUpsertOne) Update(set func(*SessionsUpsert)) *SessionsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SessionsUpsert{UpdateSet: update})
	}))
	return u
}

// SetToken sets the "token" field.
func (u *SessionsUpsertOne) SetToken(v string) *SessionsUpsertOne {
	return u.Update(func(s *SessionsUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *SessionsUpsertOne) UpdateToken() *SessionsUpsertOne {
	return u.Update(func(s *SessionsUpsert) {
		s.UpdateToken()
	})
}

// SetData sets the "data" field.
func (u *SessionsUpsertOne) SetData(v []byte) *SessionsUpsertOne {
	return u.Update(func(s *SessionsUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *SessionsUpsertOne) UpdateData() *SessionsUpsertOne {
	return u.Update(func(s *SessionsUpsert) {
		s.UpdateData()
	})
}

// SetExpiry sets the "expiry" field.
func (u *SessionsUpsertOne) SetExpiry(v time.Time) *SessionsUpsertOne {
	return u.Update(func(s *SessionsUpsert) {
		s.SetExpiry(v)
	})
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *SessionsUpsertOne) UpdateExpiry() *SessionsUpsertOne {
	return u.Update(func(s *SessionsUpsert) {
		s.UpdateExpiry()
	})
}

// Exec executes the query.
func (u *SessionsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for SessionsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SessionsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SessionsUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SessionsUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SessionsCreateBulk is the builder for creating many Sessions entities in bulk.
type SessionsCreateBulk struct {
	config
	err      error
	builders []*SessionsCreate
	conflict []sql.ConflictOption
}

// Save creates the Sessions entities in the database.
func (scb *SessionsCreateBulk) Save(ctx context.Context) ([]*Sessions, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sessions, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionsMutation)
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
func (scb *SessionsCreateBulk) SaveX(ctx context.Context) []*Sessions {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionsCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionsCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Sessions.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SessionsUpsert) {
//			SetToken(v+v).
//		}).
//		Exec(ctx)
func (scb *SessionsCreateBulk) OnConflict(opts ...sql.ConflictOption) *SessionsUpsertBulk {
	scb.conflict = opts
	return &SessionsUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Sessions.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SessionsCreateBulk) OnConflictColumns(columns ...string) *SessionsUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SessionsUpsertBulk{
		create: scb,
	}
}

// SessionsUpsertBulk is the builder for "upsert"-ing
// a bulk of Sessions nodes.
type SessionsUpsertBulk struct {
	create *SessionsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Sessions.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SessionsUpsertBulk) UpdateNewValues() *SessionsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Sessions.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SessionsUpsertBulk) Ignore() *SessionsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SessionsUpsertBulk) DoNothing() *SessionsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SessionsCreateBulk.OnConflict
// documentation for more info.
func (u *SessionsUpsertBulk) Update(set func(*SessionsUpsert)) *SessionsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SessionsUpsert{UpdateSet: update})
	}))
	return u
}

// SetToken sets the "token" field.
func (u *SessionsUpsertBulk) SetToken(v string) *SessionsUpsertBulk {
	return u.Update(func(s *SessionsUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *SessionsUpsertBulk) UpdateToken() *SessionsUpsertBulk {
	return u.Update(func(s *SessionsUpsert) {
		s.UpdateToken()
	})
}

// SetData sets the "data" field.
func (u *SessionsUpsertBulk) SetData(v []byte) *SessionsUpsertBulk {
	return u.Update(func(s *SessionsUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *SessionsUpsertBulk) UpdateData() *SessionsUpsertBulk {
	return u.Update(func(s *SessionsUpsert) {
		s.UpdateData()
	})
}

// SetExpiry sets the "expiry" field.
func (u *SessionsUpsertBulk) SetExpiry(v time.Time) *SessionsUpsertBulk {
	return u.Update(func(s *SessionsUpsert) {
		s.SetExpiry(v)
	})
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *SessionsUpsertBulk) UpdateExpiry() *SessionsUpsertBulk {
	return u.Update(func(s *SessionsUpsert) {
		s.UpdateExpiry()
	})
}

// Exec executes the query.
func (u *SessionsUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("openuem_ent: OnConflict was set for builder %d. Set it on the SessionsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for SessionsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SessionsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
