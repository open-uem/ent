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
	"github.com/open-uem/ent/memoryslot"
)

// MemorySlotCreate is the builder for creating a MemorySlot entity.
type MemorySlotCreate struct {
	config
	mutation *MemorySlotMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSlot sets the "slot" field.
func (msc *MemorySlotCreate) SetSlot(s string) *MemorySlotCreate {
	msc.mutation.SetSlot(s)
	return msc
}

// SetNillableSlot sets the "slot" field if the given value is not nil.
func (msc *MemorySlotCreate) SetNillableSlot(s *string) *MemorySlotCreate {
	if s != nil {
		msc.SetSlot(*s)
	}
	return msc
}

// SetSize sets the "size" field.
func (msc *MemorySlotCreate) SetSize(s string) *MemorySlotCreate {
	msc.mutation.SetSize(s)
	return msc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (msc *MemorySlotCreate) SetNillableSize(s *string) *MemorySlotCreate {
	if s != nil {
		msc.SetSize(*s)
	}
	return msc
}

// SetType sets the "type" field.
func (msc *MemorySlotCreate) SetType(s string) *MemorySlotCreate {
	msc.mutation.SetType(s)
	return msc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (msc *MemorySlotCreate) SetNillableType(s *string) *MemorySlotCreate {
	if s != nil {
		msc.SetType(*s)
	}
	return msc
}

// SetSerialNumber sets the "serial_number" field.
func (msc *MemorySlotCreate) SetSerialNumber(s string) *MemorySlotCreate {
	msc.mutation.SetSerialNumber(s)
	return msc
}

// SetNillableSerialNumber sets the "serial_number" field if the given value is not nil.
func (msc *MemorySlotCreate) SetNillableSerialNumber(s *string) *MemorySlotCreate {
	if s != nil {
		msc.SetSerialNumber(*s)
	}
	return msc
}

// SetPartNumber sets the "part_number" field.
func (msc *MemorySlotCreate) SetPartNumber(s string) *MemorySlotCreate {
	msc.mutation.SetPartNumber(s)
	return msc
}

// SetNillablePartNumber sets the "part_number" field if the given value is not nil.
func (msc *MemorySlotCreate) SetNillablePartNumber(s *string) *MemorySlotCreate {
	if s != nil {
		msc.SetPartNumber(*s)
	}
	return msc
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (msc *MemorySlotCreate) SetOwnerID(id string) *MemorySlotCreate {
	msc.mutation.SetOwnerID(id)
	return msc
}

// SetOwner sets the "owner" edge to the Agent entity.
func (msc *MemorySlotCreate) SetOwner(a *Agent) *MemorySlotCreate {
	return msc.SetOwnerID(a.ID)
}

// Mutation returns the MemorySlotMutation object of the builder.
func (msc *MemorySlotCreate) Mutation() *MemorySlotMutation {
	return msc.mutation
}

// Save creates the MemorySlot in the database.
func (msc *MemorySlotCreate) Save(ctx context.Context) (*MemorySlot, error) {
	return withHooks(ctx, msc.sqlSave, msc.mutation, msc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (msc *MemorySlotCreate) SaveX(ctx context.Context) *MemorySlot {
	v, err := msc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (msc *MemorySlotCreate) Exec(ctx context.Context) error {
	_, err := msc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msc *MemorySlotCreate) ExecX(ctx context.Context) {
	if err := msc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msc *MemorySlotCreate) check() error {
	if len(msc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "MemorySlot.owner"`)}
	}
	return nil
}

func (msc *MemorySlotCreate) sqlSave(ctx context.Context) (*MemorySlot, error) {
	if err := msc.check(); err != nil {
		return nil, err
	}
	_node, _spec := msc.createSpec()
	if err := sqlgraph.CreateNode(ctx, msc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	msc.mutation.id = &_node.ID
	msc.mutation.done = true
	return _node, nil
}

func (msc *MemorySlotCreate) createSpec() (*MemorySlot, *sqlgraph.CreateSpec) {
	var (
		_node = &MemorySlot{config: msc.config}
		_spec = sqlgraph.NewCreateSpec(memoryslot.Table, sqlgraph.NewFieldSpec(memoryslot.FieldID, field.TypeInt))
	)
	_spec.OnConflict = msc.conflict
	if value, ok := msc.mutation.Slot(); ok {
		_spec.SetField(memoryslot.FieldSlot, field.TypeString, value)
		_node.Slot = value
	}
	if value, ok := msc.mutation.Size(); ok {
		_spec.SetField(memoryslot.FieldSize, field.TypeString, value)
		_node.Size = value
	}
	if value, ok := msc.mutation.GetType(); ok {
		_spec.SetField(memoryslot.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := msc.mutation.SerialNumber(); ok {
		_spec.SetField(memoryslot.FieldSerialNumber, field.TypeString, value)
		_node.SerialNumber = value
	}
	if value, ok := msc.mutation.PartNumber(); ok {
		_spec.SetField(memoryslot.FieldPartNumber, field.TypeString, value)
		_node.PartNumber = value
	}
	if nodes := msc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memoryslot.OwnerTable,
			Columns: []string{memoryslot.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.agent_memoryslots = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MemorySlot.Create().
//		SetSlot(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemorySlotUpsert) {
//			SetSlot(v+v).
//		}).
//		Exec(ctx)
func (msc *MemorySlotCreate) OnConflict(opts ...sql.ConflictOption) *MemorySlotUpsertOne {
	msc.conflict = opts
	return &MemorySlotUpsertOne{
		create: msc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MemorySlot.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (msc *MemorySlotCreate) OnConflictColumns(columns ...string) *MemorySlotUpsertOne {
	msc.conflict = append(msc.conflict, sql.ConflictColumns(columns...))
	return &MemorySlotUpsertOne{
		create: msc,
	}
}

type (
	// MemorySlotUpsertOne is the builder for "upsert"-ing
	//  one MemorySlot node.
	MemorySlotUpsertOne struct {
		create *MemorySlotCreate
	}

	// MemorySlotUpsert is the "OnConflict" setter.
	MemorySlotUpsert struct {
		*sql.UpdateSet
	}
)

// SetSlot sets the "slot" field.
func (u *MemorySlotUpsert) SetSlot(v string) *MemorySlotUpsert {
	u.Set(memoryslot.FieldSlot, v)
	return u
}

// UpdateSlot sets the "slot" field to the value that was provided on create.
func (u *MemorySlotUpsert) UpdateSlot() *MemorySlotUpsert {
	u.SetExcluded(memoryslot.FieldSlot)
	return u
}

// ClearSlot clears the value of the "slot" field.
func (u *MemorySlotUpsert) ClearSlot() *MemorySlotUpsert {
	u.SetNull(memoryslot.FieldSlot)
	return u
}

// SetSize sets the "size" field.
func (u *MemorySlotUpsert) SetSize(v string) *MemorySlotUpsert {
	u.Set(memoryslot.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *MemorySlotUpsert) UpdateSize() *MemorySlotUpsert {
	u.SetExcluded(memoryslot.FieldSize)
	return u
}

// ClearSize clears the value of the "size" field.
func (u *MemorySlotUpsert) ClearSize() *MemorySlotUpsert {
	u.SetNull(memoryslot.FieldSize)
	return u
}

// SetType sets the "type" field.
func (u *MemorySlotUpsert) SetType(v string) *MemorySlotUpsert {
	u.Set(memoryslot.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *MemorySlotUpsert) UpdateType() *MemorySlotUpsert {
	u.SetExcluded(memoryslot.FieldType)
	return u
}

// ClearType clears the value of the "type" field.
func (u *MemorySlotUpsert) ClearType() *MemorySlotUpsert {
	u.SetNull(memoryslot.FieldType)
	return u
}

// SetSerialNumber sets the "serial_number" field.
func (u *MemorySlotUpsert) SetSerialNumber(v string) *MemorySlotUpsert {
	u.Set(memoryslot.FieldSerialNumber, v)
	return u
}

// UpdateSerialNumber sets the "serial_number" field to the value that was provided on create.
func (u *MemorySlotUpsert) UpdateSerialNumber() *MemorySlotUpsert {
	u.SetExcluded(memoryslot.FieldSerialNumber)
	return u
}

// ClearSerialNumber clears the value of the "serial_number" field.
func (u *MemorySlotUpsert) ClearSerialNumber() *MemorySlotUpsert {
	u.SetNull(memoryslot.FieldSerialNumber)
	return u
}

// SetPartNumber sets the "part_number" field.
func (u *MemorySlotUpsert) SetPartNumber(v string) *MemorySlotUpsert {
	u.Set(memoryslot.FieldPartNumber, v)
	return u
}

// UpdatePartNumber sets the "part_number" field to the value that was provided on create.
func (u *MemorySlotUpsert) UpdatePartNumber() *MemorySlotUpsert {
	u.SetExcluded(memoryslot.FieldPartNumber)
	return u
}

// ClearPartNumber clears the value of the "part_number" field.
func (u *MemorySlotUpsert) ClearPartNumber() *MemorySlotUpsert {
	u.SetNull(memoryslot.FieldPartNumber)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.MemorySlot.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MemorySlotUpsertOne) UpdateNewValues() *MemorySlotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MemorySlot.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MemorySlotUpsertOne) Ignore() *MemorySlotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemorySlotUpsertOne) DoNothing() *MemorySlotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemorySlotCreate.OnConflict
// documentation for more info.
func (u *MemorySlotUpsertOne) Update(set func(*MemorySlotUpsert)) *MemorySlotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemorySlotUpsert{UpdateSet: update})
	}))
	return u
}

// SetSlot sets the "slot" field.
func (u *MemorySlotUpsertOne) SetSlot(v string) *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetSlot(v)
	})
}

// UpdateSlot sets the "slot" field to the value that was provided on create.
func (u *MemorySlotUpsertOne) UpdateSlot() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdateSlot()
	})
}

// ClearSlot clears the value of the "slot" field.
func (u *MemorySlotUpsertOne) ClearSlot() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearSlot()
	})
}

// SetSize sets the "size" field.
func (u *MemorySlotUpsertOne) SetSize(v string) *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *MemorySlotUpsertOne) UpdateSize() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdateSize()
	})
}

// ClearSize clears the value of the "size" field.
func (u *MemorySlotUpsertOne) ClearSize() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearSize()
	})
}

// SetType sets the "type" field.
func (u *MemorySlotUpsertOne) SetType(v string) *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *MemorySlotUpsertOne) UpdateType() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *MemorySlotUpsertOne) ClearType() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearType()
	})
}

// SetSerialNumber sets the "serial_number" field.
func (u *MemorySlotUpsertOne) SetSerialNumber(v string) *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetSerialNumber(v)
	})
}

// UpdateSerialNumber sets the "serial_number" field to the value that was provided on create.
func (u *MemorySlotUpsertOne) UpdateSerialNumber() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdateSerialNumber()
	})
}

// ClearSerialNumber clears the value of the "serial_number" field.
func (u *MemorySlotUpsertOne) ClearSerialNumber() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearSerialNumber()
	})
}

// SetPartNumber sets the "part_number" field.
func (u *MemorySlotUpsertOne) SetPartNumber(v string) *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetPartNumber(v)
	})
}

// UpdatePartNumber sets the "part_number" field to the value that was provided on create.
func (u *MemorySlotUpsertOne) UpdatePartNumber() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdatePartNumber()
	})
}

// ClearPartNumber clears the value of the "part_number" field.
func (u *MemorySlotUpsertOne) ClearPartNumber() *MemorySlotUpsertOne {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearPartNumber()
	})
}

// Exec executes the query.
func (u *MemorySlotUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemorySlotCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemorySlotUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MemorySlotUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MemorySlotUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MemorySlotCreateBulk is the builder for creating many MemorySlot entities in bulk.
type MemorySlotCreateBulk struct {
	config
	err      error
	builders []*MemorySlotCreate
	conflict []sql.ConflictOption
}

// Save creates the MemorySlot entities in the database.
func (mscb *MemorySlotCreateBulk) Save(ctx context.Context) ([]*MemorySlot, error) {
	if mscb.err != nil {
		return nil, mscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mscb.builders))
	nodes := make([]*MemorySlot, len(mscb.builders))
	mutators := make([]Mutator, len(mscb.builders))
	for i := range mscb.builders {
		func(i int, root context.Context) {
			builder := mscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemorySlotMutation)
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
					_, err = mutators[i+1].Mutate(root, mscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mscb *MemorySlotCreateBulk) SaveX(ctx context.Context) []*MemorySlot {
	v, err := mscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mscb *MemorySlotCreateBulk) Exec(ctx context.Context) error {
	_, err := mscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mscb *MemorySlotCreateBulk) ExecX(ctx context.Context) {
	if err := mscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MemorySlot.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemorySlotUpsert) {
//			SetSlot(v+v).
//		}).
//		Exec(ctx)
func (mscb *MemorySlotCreateBulk) OnConflict(opts ...sql.ConflictOption) *MemorySlotUpsertBulk {
	mscb.conflict = opts
	return &MemorySlotUpsertBulk{
		create: mscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MemorySlot.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mscb *MemorySlotCreateBulk) OnConflictColumns(columns ...string) *MemorySlotUpsertBulk {
	mscb.conflict = append(mscb.conflict, sql.ConflictColumns(columns...))
	return &MemorySlotUpsertBulk{
		create: mscb,
	}
}

// MemorySlotUpsertBulk is the builder for "upsert"-ing
// a bulk of MemorySlot nodes.
type MemorySlotUpsertBulk struct {
	create *MemorySlotCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MemorySlot.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MemorySlotUpsertBulk) UpdateNewValues() *MemorySlotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MemorySlot.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MemorySlotUpsertBulk) Ignore() *MemorySlotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemorySlotUpsertBulk) DoNothing() *MemorySlotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemorySlotCreateBulk.OnConflict
// documentation for more info.
func (u *MemorySlotUpsertBulk) Update(set func(*MemorySlotUpsert)) *MemorySlotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemorySlotUpsert{UpdateSet: update})
	}))
	return u
}

// SetSlot sets the "slot" field.
func (u *MemorySlotUpsertBulk) SetSlot(v string) *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetSlot(v)
	})
}

// UpdateSlot sets the "slot" field to the value that was provided on create.
func (u *MemorySlotUpsertBulk) UpdateSlot() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdateSlot()
	})
}

// ClearSlot clears the value of the "slot" field.
func (u *MemorySlotUpsertBulk) ClearSlot() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearSlot()
	})
}

// SetSize sets the "size" field.
func (u *MemorySlotUpsertBulk) SetSize(v string) *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *MemorySlotUpsertBulk) UpdateSize() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdateSize()
	})
}

// ClearSize clears the value of the "size" field.
func (u *MemorySlotUpsertBulk) ClearSize() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearSize()
	})
}

// SetType sets the "type" field.
func (u *MemorySlotUpsertBulk) SetType(v string) *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *MemorySlotUpsertBulk) UpdateType() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *MemorySlotUpsertBulk) ClearType() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearType()
	})
}

// SetSerialNumber sets the "serial_number" field.
func (u *MemorySlotUpsertBulk) SetSerialNumber(v string) *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetSerialNumber(v)
	})
}

// UpdateSerialNumber sets the "serial_number" field to the value that was provided on create.
func (u *MemorySlotUpsertBulk) UpdateSerialNumber() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdateSerialNumber()
	})
}

// ClearSerialNumber clears the value of the "serial_number" field.
func (u *MemorySlotUpsertBulk) ClearSerialNumber() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearSerialNumber()
	})
}

// SetPartNumber sets the "part_number" field.
func (u *MemorySlotUpsertBulk) SetPartNumber(v string) *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.SetPartNumber(v)
	})
}

// UpdatePartNumber sets the "part_number" field to the value that was provided on create.
func (u *MemorySlotUpsertBulk) UpdatePartNumber() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.UpdatePartNumber()
	})
}

// ClearPartNumber clears the value of the "part_number" field.
func (u *MemorySlotUpsertBulk) ClearPartNumber() *MemorySlotUpsertBulk {
	return u.Update(func(s *MemorySlotUpsert) {
		s.ClearPartNumber()
	})
}

// Exec executes the query.
func (u *MemorySlotUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MemorySlotCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemorySlotCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemorySlotUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
