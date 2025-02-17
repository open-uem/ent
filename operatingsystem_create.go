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
	"github.com/open-uem/ent/operatingsystem"
)

// OperatingSystemCreate is the builder for creating a OperatingSystem entity.
type OperatingSystemCreate struct {
	config
	mutation *OperatingSystemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetType sets the "type" field.
func (osc *OperatingSystemCreate) SetType(s string) *OperatingSystemCreate {
	osc.mutation.SetType(s)
	return osc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (osc *OperatingSystemCreate) SetNillableType(s *string) *OperatingSystemCreate {
	if s != nil {
		osc.SetType(*s)
	}
	return osc
}

// SetVersion sets the "version" field.
func (osc *OperatingSystemCreate) SetVersion(s string) *OperatingSystemCreate {
	osc.mutation.SetVersion(s)
	return osc
}

// SetDescription sets the "description" field.
func (osc *OperatingSystemCreate) SetDescription(s string) *OperatingSystemCreate {
	osc.mutation.SetDescription(s)
	return osc
}

// SetEdition sets the "edition" field.
func (osc *OperatingSystemCreate) SetEdition(s string) *OperatingSystemCreate {
	osc.mutation.SetEdition(s)
	return osc
}

// SetNillableEdition sets the "edition" field if the given value is not nil.
func (osc *OperatingSystemCreate) SetNillableEdition(s *string) *OperatingSystemCreate {
	if s != nil {
		osc.SetEdition(*s)
	}
	return osc
}

// SetInstallDate sets the "install_date" field.
func (osc *OperatingSystemCreate) SetInstallDate(t time.Time) *OperatingSystemCreate {
	osc.mutation.SetInstallDate(t)
	return osc
}

// SetNillableInstallDate sets the "install_date" field if the given value is not nil.
func (osc *OperatingSystemCreate) SetNillableInstallDate(t *time.Time) *OperatingSystemCreate {
	if t != nil {
		osc.SetInstallDate(*t)
	}
	return osc
}

// SetArch sets the "arch" field.
func (osc *OperatingSystemCreate) SetArch(s string) *OperatingSystemCreate {
	osc.mutation.SetArch(s)
	return osc
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (osc *OperatingSystemCreate) SetNillableArch(s *string) *OperatingSystemCreate {
	if s != nil {
		osc.SetArch(*s)
	}
	return osc
}

// SetUsername sets the "username" field.
func (osc *OperatingSystemCreate) SetUsername(s string) *OperatingSystemCreate {
	osc.mutation.SetUsername(s)
	return osc
}

// SetLastBootupTime sets the "last_bootup_time" field.
func (osc *OperatingSystemCreate) SetLastBootupTime(t time.Time) *OperatingSystemCreate {
	osc.mutation.SetLastBootupTime(t)
	return osc
}

// SetNillableLastBootupTime sets the "last_bootup_time" field if the given value is not nil.
func (osc *OperatingSystemCreate) SetNillableLastBootupTime(t *time.Time) *OperatingSystemCreate {
	if t != nil {
		osc.SetLastBootupTime(*t)
	}
	return osc
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (osc *OperatingSystemCreate) SetOwnerID(id string) *OperatingSystemCreate {
	osc.mutation.SetOwnerID(id)
	return osc
}

// SetOwner sets the "owner" edge to the Agent entity.
func (osc *OperatingSystemCreate) SetOwner(a *Agent) *OperatingSystemCreate {
	return osc.SetOwnerID(a.ID)
}

// Mutation returns the OperatingSystemMutation object of the builder.
func (osc *OperatingSystemCreate) Mutation() *OperatingSystemMutation {
	return osc.mutation
}

// Save creates the OperatingSystem in the database.
func (osc *OperatingSystemCreate) Save(ctx context.Context) (*OperatingSystem, error) {
	return withHooks(ctx, osc.sqlSave, osc.mutation, osc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (osc *OperatingSystemCreate) SaveX(ctx context.Context) *OperatingSystem {
	v, err := osc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (osc *OperatingSystemCreate) Exec(ctx context.Context) error {
	_, err := osc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osc *OperatingSystemCreate) ExecX(ctx context.Context) {
	if err := osc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osc *OperatingSystemCreate) check() error {
	if _, ok := osc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "OperatingSystem.version"`)}
	}
	if _, ok := osc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "OperatingSystem.description"`)}
	}
	if _, ok := osc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "OperatingSystem.username"`)}
	}
	if len(osc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "OperatingSystem.owner"`)}
	}
	return nil
}

func (osc *OperatingSystemCreate) sqlSave(ctx context.Context) (*OperatingSystem, error) {
	if err := osc.check(); err != nil {
		return nil, err
	}
	_node, _spec := osc.createSpec()
	if err := sqlgraph.CreateNode(ctx, osc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	osc.mutation.id = &_node.ID
	osc.mutation.done = true
	return _node, nil
}

func (osc *OperatingSystemCreate) createSpec() (*OperatingSystem, *sqlgraph.CreateSpec) {
	var (
		_node = &OperatingSystem{config: osc.config}
		_spec = sqlgraph.NewCreateSpec(operatingsystem.Table, sqlgraph.NewFieldSpec(operatingsystem.FieldID, field.TypeInt))
	)
	_spec.OnConflict = osc.conflict
	if value, ok := osc.mutation.GetType(); ok {
		_spec.SetField(operatingsystem.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := osc.mutation.Version(); ok {
		_spec.SetField(operatingsystem.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := osc.mutation.Description(); ok {
		_spec.SetField(operatingsystem.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := osc.mutation.Edition(); ok {
		_spec.SetField(operatingsystem.FieldEdition, field.TypeString, value)
		_node.Edition = value
	}
	if value, ok := osc.mutation.InstallDate(); ok {
		_spec.SetField(operatingsystem.FieldInstallDate, field.TypeTime, value)
		_node.InstallDate = value
	}
	if value, ok := osc.mutation.Arch(); ok {
		_spec.SetField(operatingsystem.FieldArch, field.TypeString, value)
		_node.Arch = value
	}
	if value, ok := osc.mutation.Username(); ok {
		_spec.SetField(operatingsystem.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := osc.mutation.LastBootupTime(); ok {
		_spec.SetField(operatingsystem.FieldLastBootupTime, field.TypeTime, value)
		_node.LastBootupTime = value
	}
	if nodes := osc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   operatingsystem.OwnerTable,
			Columns: []string{operatingsystem.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.agent_operatingsystem = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OperatingSystem.Create().
//		SetType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OperatingSystemUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
func (osc *OperatingSystemCreate) OnConflict(opts ...sql.ConflictOption) *OperatingSystemUpsertOne {
	osc.conflict = opts
	return &OperatingSystemUpsertOne{
		create: osc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OperatingSystem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (osc *OperatingSystemCreate) OnConflictColumns(columns ...string) *OperatingSystemUpsertOne {
	osc.conflict = append(osc.conflict, sql.ConflictColumns(columns...))
	return &OperatingSystemUpsertOne{
		create: osc,
	}
}

type (
	// OperatingSystemUpsertOne is the builder for "upsert"-ing
	//  one OperatingSystem node.
	OperatingSystemUpsertOne struct {
		create *OperatingSystemCreate
	}

	// OperatingSystemUpsert is the "OnConflict" setter.
	OperatingSystemUpsert struct {
		*sql.UpdateSet
	}
)

// SetType sets the "type" field.
func (u *OperatingSystemUpsert) SetType(v string) *OperatingSystemUpsert {
	u.Set(operatingsystem.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OperatingSystemUpsert) UpdateType() *OperatingSystemUpsert {
	u.SetExcluded(operatingsystem.FieldType)
	return u
}

// ClearType clears the value of the "type" field.
func (u *OperatingSystemUpsert) ClearType() *OperatingSystemUpsert {
	u.SetNull(operatingsystem.FieldType)
	return u
}

// SetVersion sets the "version" field.
func (u *OperatingSystemUpsert) SetVersion(v string) *OperatingSystemUpsert {
	u.Set(operatingsystem.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *OperatingSystemUpsert) UpdateVersion() *OperatingSystemUpsert {
	u.SetExcluded(operatingsystem.FieldVersion)
	return u
}

// SetDescription sets the "description" field.
func (u *OperatingSystemUpsert) SetDescription(v string) *OperatingSystemUpsert {
	u.Set(operatingsystem.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *OperatingSystemUpsert) UpdateDescription() *OperatingSystemUpsert {
	u.SetExcluded(operatingsystem.FieldDescription)
	return u
}

// SetEdition sets the "edition" field.
func (u *OperatingSystemUpsert) SetEdition(v string) *OperatingSystemUpsert {
	u.Set(operatingsystem.FieldEdition, v)
	return u
}

// UpdateEdition sets the "edition" field to the value that was provided on create.
func (u *OperatingSystemUpsert) UpdateEdition() *OperatingSystemUpsert {
	u.SetExcluded(operatingsystem.FieldEdition)
	return u
}

// ClearEdition clears the value of the "edition" field.
func (u *OperatingSystemUpsert) ClearEdition() *OperatingSystemUpsert {
	u.SetNull(operatingsystem.FieldEdition)
	return u
}

// SetInstallDate sets the "install_date" field.
func (u *OperatingSystemUpsert) SetInstallDate(v time.Time) *OperatingSystemUpsert {
	u.Set(operatingsystem.FieldInstallDate, v)
	return u
}

// UpdateInstallDate sets the "install_date" field to the value that was provided on create.
func (u *OperatingSystemUpsert) UpdateInstallDate() *OperatingSystemUpsert {
	u.SetExcluded(operatingsystem.FieldInstallDate)
	return u
}

// ClearInstallDate clears the value of the "install_date" field.
func (u *OperatingSystemUpsert) ClearInstallDate() *OperatingSystemUpsert {
	u.SetNull(operatingsystem.FieldInstallDate)
	return u
}

// SetArch sets the "arch" field.
func (u *OperatingSystemUpsert) SetArch(v string) *OperatingSystemUpsert {
	u.Set(operatingsystem.FieldArch, v)
	return u
}

// UpdateArch sets the "arch" field to the value that was provided on create.
func (u *OperatingSystemUpsert) UpdateArch() *OperatingSystemUpsert {
	u.SetExcluded(operatingsystem.FieldArch)
	return u
}

// ClearArch clears the value of the "arch" field.
func (u *OperatingSystemUpsert) ClearArch() *OperatingSystemUpsert {
	u.SetNull(operatingsystem.FieldArch)
	return u
}

// SetUsername sets the "username" field.
func (u *OperatingSystemUpsert) SetUsername(v string) *OperatingSystemUpsert {
	u.Set(operatingsystem.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *OperatingSystemUpsert) UpdateUsername() *OperatingSystemUpsert {
	u.SetExcluded(operatingsystem.FieldUsername)
	return u
}

// SetLastBootupTime sets the "last_bootup_time" field.
func (u *OperatingSystemUpsert) SetLastBootupTime(v time.Time) *OperatingSystemUpsert {
	u.Set(operatingsystem.FieldLastBootupTime, v)
	return u
}

// UpdateLastBootupTime sets the "last_bootup_time" field to the value that was provided on create.
func (u *OperatingSystemUpsert) UpdateLastBootupTime() *OperatingSystemUpsert {
	u.SetExcluded(operatingsystem.FieldLastBootupTime)
	return u
}

// ClearLastBootupTime clears the value of the "last_bootup_time" field.
func (u *OperatingSystemUpsert) ClearLastBootupTime() *OperatingSystemUpsert {
	u.SetNull(operatingsystem.FieldLastBootupTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.OperatingSystem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OperatingSystemUpsertOne) UpdateNewValues() *OperatingSystemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OperatingSystem.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OperatingSystemUpsertOne) Ignore() *OperatingSystemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OperatingSystemUpsertOne) DoNothing() *OperatingSystemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OperatingSystemCreate.OnConflict
// documentation for more info.
func (u *OperatingSystemUpsertOne) Update(set func(*OperatingSystemUpsert)) *OperatingSystemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OperatingSystemUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *OperatingSystemUpsertOne) SetType(v string) *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OperatingSystemUpsertOne) UpdateType() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *OperatingSystemUpsertOne) ClearType() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearType()
	})
}

// SetVersion sets the "version" field.
func (u *OperatingSystemUpsertOne) SetVersion(v string) *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *OperatingSystemUpsertOne) UpdateVersion() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateVersion()
	})
}

// SetDescription sets the "description" field.
func (u *OperatingSystemUpsertOne) SetDescription(v string) *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *OperatingSystemUpsertOne) UpdateDescription() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateDescription()
	})
}

// SetEdition sets the "edition" field.
func (u *OperatingSystemUpsertOne) SetEdition(v string) *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetEdition(v)
	})
}

// UpdateEdition sets the "edition" field to the value that was provided on create.
func (u *OperatingSystemUpsertOne) UpdateEdition() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateEdition()
	})
}

// ClearEdition clears the value of the "edition" field.
func (u *OperatingSystemUpsertOne) ClearEdition() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearEdition()
	})
}

// SetInstallDate sets the "install_date" field.
func (u *OperatingSystemUpsertOne) SetInstallDate(v time.Time) *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetInstallDate(v)
	})
}

// UpdateInstallDate sets the "install_date" field to the value that was provided on create.
func (u *OperatingSystemUpsertOne) UpdateInstallDate() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateInstallDate()
	})
}

// ClearInstallDate clears the value of the "install_date" field.
func (u *OperatingSystemUpsertOne) ClearInstallDate() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearInstallDate()
	})
}

// SetArch sets the "arch" field.
func (u *OperatingSystemUpsertOne) SetArch(v string) *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetArch(v)
	})
}

// UpdateArch sets the "arch" field to the value that was provided on create.
func (u *OperatingSystemUpsertOne) UpdateArch() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateArch()
	})
}

// ClearArch clears the value of the "arch" field.
func (u *OperatingSystemUpsertOne) ClearArch() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearArch()
	})
}

// SetUsername sets the "username" field.
func (u *OperatingSystemUpsertOne) SetUsername(v string) *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *OperatingSystemUpsertOne) UpdateUsername() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateUsername()
	})
}

// SetLastBootupTime sets the "last_bootup_time" field.
func (u *OperatingSystemUpsertOne) SetLastBootupTime(v time.Time) *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetLastBootupTime(v)
	})
}

// UpdateLastBootupTime sets the "last_bootup_time" field to the value that was provided on create.
func (u *OperatingSystemUpsertOne) UpdateLastBootupTime() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateLastBootupTime()
	})
}

// ClearLastBootupTime clears the value of the "last_bootup_time" field.
func (u *OperatingSystemUpsertOne) ClearLastBootupTime() *OperatingSystemUpsertOne {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearLastBootupTime()
	})
}

// Exec executes the query.
func (u *OperatingSystemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OperatingSystemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OperatingSystemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OperatingSystemUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OperatingSystemUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OperatingSystemCreateBulk is the builder for creating many OperatingSystem entities in bulk.
type OperatingSystemCreateBulk struct {
	config
	err      error
	builders []*OperatingSystemCreate
	conflict []sql.ConflictOption
}

// Save creates the OperatingSystem entities in the database.
func (oscb *OperatingSystemCreateBulk) Save(ctx context.Context) ([]*OperatingSystem, error) {
	if oscb.err != nil {
		return nil, oscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oscb.builders))
	nodes := make([]*OperatingSystem, len(oscb.builders))
	mutators := make([]Mutator, len(oscb.builders))
	for i := range oscb.builders {
		func(i int, root context.Context) {
			builder := oscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OperatingSystemMutation)
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
					_, err = mutators[i+1].Mutate(root, oscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = oscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oscb *OperatingSystemCreateBulk) SaveX(ctx context.Context) []*OperatingSystem {
	v, err := oscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oscb *OperatingSystemCreateBulk) Exec(ctx context.Context) error {
	_, err := oscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscb *OperatingSystemCreateBulk) ExecX(ctx context.Context) {
	if err := oscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OperatingSystem.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OperatingSystemUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
func (oscb *OperatingSystemCreateBulk) OnConflict(opts ...sql.ConflictOption) *OperatingSystemUpsertBulk {
	oscb.conflict = opts
	return &OperatingSystemUpsertBulk{
		create: oscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OperatingSystem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oscb *OperatingSystemCreateBulk) OnConflictColumns(columns ...string) *OperatingSystemUpsertBulk {
	oscb.conflict = append(oscb.conflict, sql.ConflictColumns(columns...))
	return &OperatingSystemUpsertBulk{
		create: oscb,
	}
}

// OperatingSystemUpsertBulk is the builder for "upsert"-ing
// a bulk of OperatingSystem nodes.
type OperatingSystemUpsertBulk struct {
	create *OperatingSystemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OperatingSystem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OperatingSystemUpsertBulk) UpdateNewValues() *OperatingSystemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OperatingSystem.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OperatingSystemUpsertBulk) Ignore() *OperatingSystemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OperatingSystemUpsertBulk) DoNothing() *OperatingSystemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OperatingSystemCreateBulk.OnConflict
// documentation for more info.
func (u *OperatingSystemUpsertBulk) Update(set func(*OperatingSystemUpsert)) *OperatingSystemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OperatingSystemUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *OperatingSystemUpsertBulk) SetType(v string) *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OperatingSystemUpsertBulk) UpdateType() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *OperatingSystemUpsertBulk) ClearType() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearType()
	})
}

// SetVersion sets the "version" field.
func (u *OperatingSystemUpsertBulk) SetVersion(v string) *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *OperatingSystemUpsertBulk) UpdateVersion() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateVersion()
	})
}

// SetDescription sets the "description" field.
func (u *OperatingSystemUpsertBulk) SetDescription(v string) *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *OperatingSystemUpsertBulk) UpdateDescription() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateDescription()
	})
}

// SetEdition sets the "edition" field.
func (u *OperatingSystemUpsertBulk) SetEdition(v string) *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetEdition(v)
	})
}

// UpdateEdition sets the "edition" field to the value that was provided on create.
func (u *OperatingSystemUpsertBulk) UpdateEdition() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateEdition()
	})
}

// ClearEdition clears the value of the "edition" field.
func (u *OperatingSystemUpsertBulk) ClearEdition() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearEdition()
	})
}

// SetInstallDate sets the "install_date" field.
func (u *OperatingSystemUpsertBulk) SetInstallDate(v time.Time) *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetInstallDate(v)
	})
}

// UpdateInstallDate sets the "install_date" field to the value that was provided on create.
func (u *OperatingSystemUpsertBulk) UpdateInstallDate() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateInstallDate()
	})
}

// ClearInstallDate clears the value of the "install_date" field.
func (u *OperatingSystemUpsertBulk) ClearInstallDate() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearInstallDate()
	})
}

// SetArch sets the "arch" field.
func (u *OperatingSystemUpsertBulk) SetArch(v string) *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetArch(v)
	})
}

// UpdateArch sets the "arch" field to the value that was provided on create.
func (u *OperatingSystemUpsertBulk) UpdateArch() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateArch()
	})
}

// ClearArch clears the value of the "arch" field.
func (u *OperatingSystemUpsertBulk) ClearArch() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearArch()
	})
}

// SetUsername sets the "username" field.
func (u *OperatingSystemUpsertBulk) SetUsername(v string) *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *OperatingSystemUpsertBulk) UpdateUsername() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateUsername()
	})
}

// SetLastBootupTime sets the "last_bootup_time" field.
func (u *OperatingSystemUpsertBulk) SetLastBootupTime(v time.Time) *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.SetLastBootupTime(v)
	})
}

// UpdateLastBootupTime sets the "last_bootup_time" field to the value that was provided on create.
func (u *OperatingSystemUpsertBulk) UpdateLastBootupTime() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.UpdateLastBootupTime()
	})
}

// ClearLastBootupTime clears the value of the "last_bootup_time" field.
func (u *OperatingSystemUpsertBulk) ClearLastBootupTime() *OperatingSystemUpsertBulk {
	return u.Update(func(s *OperatingSystemUpsert) {
		s.ClearLastBootupTime()
	})
}

// Exec executes the query.
func (u *OperatingSystemUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OperatingSystemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OperatingSystemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OperatingSystemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
