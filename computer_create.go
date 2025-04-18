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
	"github.com/open-uem/ent/computer"
)

// ComputerCreate is the builder for creating a Computer entity.
type ComputerCreate struct {
	config
	mutation *ComputerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetManufacturer sets the "manufacturer" field.
func (cc *ComputerCreate) SetManufacturer(s string) *ComputerCreate {
	cc.mutation.SetManufacturer(s)
	return cc
}

// SetNillableManufacturer sets the "manufacturer" field if the given value is not nil.
func (cc *ComputerCreate) SetNillableManufacturer(s *string) *ComputerCreate {
	if s != nil {
		cc.SetManufacturer(*s)
	}
	return cc
}

// SetModel sets the "model" field.
func (cc *ComputerCreate) SetModel(s string) *ComputerCreate {
	cc.mutation.SetModel(s)
	return cc
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (cc *ComputerCreate) SetNillableModel(s *string) *ComputerCreate {
	if s != nil {
		cc.SetModel(*s)
	}
	return cc
}

// SetSerial sets the "serial" field.
func (cc *ComputerCreate) SetSerial(s string) *ComputerCreate {
	cc.mutation.SetSerial(s)
	return cc
}

// SetNillableSerial sets the "serial" field if the given value is not nil.
func (cc *ComputerCreate) SetNillableSerial(s *string) *ComputerCreate {
	if s != nil {
		cc.SetSerial(*s)
	}
	return cc
}

// SetMemory sets the "memory" field.
func (cc *ComputerCreate) SetMemory(u uint64) *ComputerCreate {
	cc.mutation.SetMemory(u)
	return cc
}

// SetNillableMemory sets the "memory" field if the given value is not nil.
func (cc *ComputerCreate) SetNillableMemory(u *uint64) *ComputerCreate {
	if u != nil {
		cc.SetMemory(*u)
	}
	return cc
}

// SetProcessor sets the "processor" field.
func (cc *ComputerCreate) SetProcessor(s string) *ComputerCreate {
	cc.mutation.SetProcessor(s)
	return cc
}

// SetNillableProcessor sets the "processor" field if the given value is not nil.
func (cc *ComputerCreate) SetNillableProcessor(s *string) *ComputerCreate {
	if s != nil {
		cc.SetProcessor(*s)
	}
	return cc
}

// SetProcessorCores sets the "processor_cores" field.
func (cc *ComputerCreate) SetProcessorCores(i int64) *ComputerCreate {
	cc.mutation.SetProcessorCores(i)
	return cc
}

// SetNillableProcessorCores sets the "processor_cores" field if the given value is not nil.
func (cc *ComputerCreate) SetNillableProcessorCores(i *int64) *ComputerCreate {
	if i != nil {
		cc.SetProcessorCores(*i)
	}
	return cc
}

// SetProcessorArch sets the "processor_arch" field.
func (cc *ComputerCreate) SetProcessorArch(s string) *ComputerCreate {
	cc.mutation.SetProcessorArch(s)
	return cc
}

// SetNillableProcessorArch sets the "processor_arch" field if the given value is not nil.
func (cc *ComputerCreate) SetNillableProcessorArch(s *string) *ComputerCreate {
	if s != nil {
		cc.SetProcessorArch(*s)
	}
	return cc
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (cc *ComputerCreate) SetOwnerID(id string) *ComputerCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetOwner sets the "owner" edge to the Agent entity.
func (cc *ComputerCreate) SetOwner(a *Agent) *ComputerCreate {
	return cc.SetOwnerID(a.ID)
}

// Mutation returns the ComputerMutation object of the builder.
func (cc *ComputerCreate) Mutation() *ComputerMutation {
	return cc.mutation
}

// Save creates the Computer in the database.
func (cc *ComputerCreate) Save(ctx context.Context) (*Computer, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ComputerCreate) SaveX(ctx context.Context) *Computer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ComputerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ComputerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ComputerCreate) check() error {
	if len(cc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Computer.owner"`)}
	}
	return nil
}

func (cc *ComputerCreate) sqlSave(ctx context.Context) (*Computer, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ComputerCreate) createSpec() (*Computer, *sqlgraph.CreateSpec) {
	var (
		_node = &Computer{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(computer.Table, sqlgraph.NewFieldSpec(computer.FieldID, field.TypeInt))
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.Manufacturer(); ok {
		_spec.SetField(computer.FieldManufacturer, field.TypeString, value)
		_node.Manufacturer = value
	}
	if value, ok := cc.mutation.Model(); ok {
		_spec.SetField(computer.FieldModel, field.TypeString, value)
		_node.Model = value
	}
	if value, ok := cc.mutation.Serial(); ok {
		_spec.SetField(computer.FieldSerial, field.TypeString, value)
		_node.Serial = value
	}
	if value, ok := cc.mutation.Memory(); ok {
		_spec.SetField(computer.FieldMemory, field.TypeUint64, value)
		_node.Memory = value
	}
	if value, ok := cc.mutation.Processor(); ok {
		_spec.SetField(computer.FieldProcessor, field.TypeString, value)
		_node.Processor = value
	}
	if value, ok := cc.mutation.ProcessorCores(); ok {
		_spec.SetField(computer.FieldProcessorCores, field.TypeInt64, value)
		_node.ProcessorCores = value
	}
	if value, ok := cc.mutation.ProcessorArch(); ok {
		_spec.SetField(computer.FieldProcessorArch, field.TypeString, value)
		_node.ProcessorArch = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   computer.OwnerTable,
			Columns: []string{computer.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.agent_computer = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Computer.Create().
//		SetManufacturer(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ComputerUpsert) {
//			SetManufacturer(v+v).
//		}).
//		Exec(ctx)
func (cc *ComputerCreate) OnConflict(opts ...sql.ConflictOption) *ComputerUpsertOne {
	cc.conflict = opts
	return &ComputerUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Computer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *ComputerCreate) OnConflictColumns(columns ...string) *ComputerUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ComputerUpsertOne{
		create: cc,
	}
}

type (
	// ComputerUpsertOne is the builder for "upsert"-ing
	//  one Computer node.
	ComputerUpsertOne struct {
		create *ComputerCreate
	}

	// ComputerUpsert is the "OnConflict" setter.
	ComputerUpsert struct {
		*sql.UpdateSet
	}
)

// SetManufacturer sets the "manufacturer" field.
func (u *ComputerUpsert) SetManufacturer(v string) *ComputerUpsert {
	u.Set(computer.FieldManufacturer, v)
	return u
}

// UpdateManufacturer sets the "manufacturer" field to the value that was provided on create.
func (u *ComputerUpsert) UpdateManufacturer() *ComputerUpsert {
	u.SetExcluded(computer.FieldManufacturer)
	return u
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (u *ComputerUpsert) ClearManufacturer() *ComputerUpsert {
	u.SetNull(computer.FieldManufacturer)
	return u
}

// SetModel sets the "model" field.
func (u *ComputerUpsert) SetModel(v string) *ComputerUpsert {
	u.Set(computer.FieldModel, v)
	return u
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *ComputerUpsert) UpdateModel() *ComputerUpsert {
	u.SetExcluded(computer.FieldModel)
	return u
}

// ClearModel clears the value of the "model" field.
func (u *ComputerUpsert) ClearModel() *ComputerUpsert {
	u.SetNull(computer.FieldModel)
	return u
}

// SetSerial sets the "serial" field.
func (u *ComputerUpsert) SetSerial(v string) *ComputerUpsert {
	u.Set(computer.FieldSerial, v)
	return u
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *ComputerUpsert) UpdateSerial() *ComputerUpsert {
	u.SetExcluded(computer.FieldSerial)
	return u
}

// ClearSerial clears the value of the "serial" field.
func (u *ComputerUpsert) ClearSerial() *ComputerUpsert {
	u.SetNull(computer.FieldSerial)
	return u
}

// SetMemory sets the "memory" field.
func (u *ComputerUpsert) SetMemory(v uint64) *ComputerUpsert {
	u.Set(computer.FieldMemory, v)
	return u
}

// UpdateMemory sets the "memory" field to the value that was provided on create.
func (u *ComputerUpsert) UpdateMemory() *ComputerUpsert {
	u.SetExcluded(computer.FieldMemory)
	return u
}

// AddMemory adds v to the "memory" field.
func (u *ComputerUpsert) AddMemory(v uint64) *ComputerUpsert {
	u.Add(computer.FieldMemory, v)
	return u
}

// ClearMemory clears the value of the "memory" field.
func (u *ComputerUpsert) ClearMemory() *ComputerUpsert {
	u.SetNull(computer.FieldMemory)
	return u
}

// SetProcessor sets the "processor" field.
func (u *ComputerUpsert) SetProcessor(v string) *ComputerUpsert {
	u.Set(computer.FieldProcessor, v)
	return u
}

// UpdateProcessor sets the "processor" field to the value that was provided on create.
func (u *ComputerUpsert) UpdateProcessor() *ComputerUpsert {
	u.SetExcluded(computer.FieldProcessor)
	return u
}

// ClearProcessor clears the value of the "processor" field.
func (u *ComputerUpsert) ClearProcessor() *ComputerUpsert {
	u.SetNull(computer.FieldProcessor)
	return u
}

// SetProcessorCores sets the "processor_cores" field.
func (u *ComputerUpsert) SetProcessorCores(v int64) *ComputerUpsert {
	u.Set(computer.FieldProcessorCores, v)
	return u
}

// UpdateProcessorCores sets the "processor_cores" field to the value that was provided on create.
func (u *ComputerUpsert) UpdateProcessorCores() *ComputerUpsert {
	u.SetExcluded(computer.FieldProcessorCores)
	return u
}

// AddProcessorCores adds v to the "processor_cores" field.
func (u *ComputerUpsert) AddProcessorCores(v int64) *ComputerUpsert {
	u.Add(computer.FieldProcessorCores, v)
	return u
}

// ClearProcessorCores clears the value of the "processor_cores" field.
func (u *ComputerUpsert) ClearProcessorCores() *ComputerUpsert {
	u.SetNull(computer.FieldProcessorCores)
	return u
}

// SetProcessorArch sets the "processor_arch" field.
func (u *ComputerUpsert) SetProcessorArch(v string) *ComputerUpsert {
	u.Set(computer.FieldProcessorArch, v)
	return u
}

// UpdateProcessorArch sets the "processor_arch" field to the value that was provided on create.
func (u *ComputerUpsert) UpdateProcessorArch() *ComputerUpsert {
	u.SetExcluded(computer.FieldProcessorArch)
	return u
}

// ClearProcessorArch clears the value of the "processor_arch" field.
func (u *ComputerUpsert) ClearProcessorArch() *ComputerUpsert {
	u.SetNull(computer.FieldProcessorArch)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Computer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ComputerUpsertOne) UpdateNewValues() *ComputerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Computer.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ComputerUpsertOne) Ignore() *ComputerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ComputerUpsertOne) DoNothing() *ComputerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ComputerCreate.OnConflict
// documentation for more info.
func (u *ComputerUpsertOne) Update(set func(*ComputerUpsert)) *ComputerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ComputerUpsert{UpdateSet: update})
	}))
	return u
}

// SetManufacturer sets the "manufacturer" field.
func (u *ComputerUpsertOne) SetManufacturer(v string) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.SetManufacturer(v)
	})
}

// UpdateManufacturer sets the "manufacturer" field to the value that was provided on create.
func (u *ComputerUpsertOne) UpdateManufacturer() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateManufacturer()
	})
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (u *ComputerUpsertOne) ClearManufacturer() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearManufacturer()
	})
}

// SetModel sets the "model" field.
func (u *ComputerUpsertOne) SetModel(v string) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.SetModel(v)
	})
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *ComputerUpsertOne) UpdateModel() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateModel()
	})
}

// ClearModel clears the value of the "model" field.
func (u *ComputerUpsertOne) ClearModel() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearModel()
	})
}

// SetSerial sets the "serial" field.
func (u *ComputerUpsertOne) SetSerial(v string) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.SetSerial(v)
	})
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *ComputerUpsertOne) UpdateSerial() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateSerial()
	})
}

// ClearSerial clears the value of the "serial" field.
func (u *ComputerUpsertOne) ClearSerial() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearSerial()
	})
}

// SetMemory sets the "memory" field.
func (u *ComputerUpsertOne) SetMemory(v uint64) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.SetMemory(v)
	})
}

// AddMemory adds v to the "memory" field.
func (u *ComputerUpsertOne) AddMemory(v uint64) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.AddMemory(v)
	})
}

// UpdateMemory sets the "memory" field to the value that was provided on create.
func (u *ComputerUpsertOne) UpdateMemory() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateMemory()
	})
}

// ClearMemory clears the value of the "memory" field.
func (u *ComputerUpsertOne) ClearMemory() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearMemory()
	})
}

// SetProcessor sets the "processor" field.
func (u *ComputerUpsertOne) SetProcessor(v string) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.SetProcessor(v)
	})
}

// UpdateProcessor sets the "processor" field to the value that was provided on create.
func (u *ComputerUpsertOne) UpdateProcessor() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateProcessor()
	})
}

// ClearProcessor clears the value of the "processor" field.
func (u *ComputerUpsertOne) ClearProcessor() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearProcessor()
	})
}

// SetProcessorCores sets the "processor_cores" field.
func (u *ComputerUpsertOne) SetProcessorCores(v int64) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.SetProcessorCores(v)
	})
}

// AddProcessorCores adds v to the "processor_cores" field.
func (u *ComputerUpsertOne) AddProcessorCores(v int64) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.AddProcessorCores(v)
	})
}

// UpdateProcessorCores sets the "processor_cores" field to the value that was provided on create.
func (u *ComputerUpsertOne) UpdateProcessorCores() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateProcessorCores()
	})
}

// ClearProcessorCores clears the value of the "processor_cores" field.
func (u *ComputerUpsertOne) ClearProcessorCores() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearProcessorCores()
	})
}

// SetProcessorArch sets the "processor_arch" field.
func (u *ComputerUpsertOne) SetProcessorArch(v string) *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.SetProcessorArch(v)
	})
}

// UpdateProcessorArch sets the "processor_arch" field to the value that was provided on create.
func (u *ComputerUpsertOne) UpdateProcessorArch() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateProcessorArch()
	})
}

// ClearProcessorArch clears the value of the "processor_arch" field.
func (u *ComputerUpsertOne) ClearProcessorArch() *ComputerUpsertOne {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearProcessorArch()
	})
}

// Exec executes the query.
func (u *ComputerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ComputerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ComputerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ComputerUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ComputerUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ComputerCreateBulk is the builder for creating many Computer entities in bulk.
type ComputerCreateBulk struct {
	config
	err      error
	builders []*ComputerCreate
	conflict []sql.ConflictOption
}

// Save creates the Computer entities in the database.
func (ccb *ComputerCreateBulk) Save(ctx context.Context) ([]*Computer, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Computer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ComputerMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ComputerCreateBulk) SaveX(ctx context.Context) []*Computer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ComputerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ComputerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Computer.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ComputerUpsert) {
//			SetManufacturer(v+v).
//		}).
//		Exec(ctx)
func (ccb *ComputerCreateBulk) OnConflict(opts ...sql.ConflictOption) *ComputerUpsertBulk {
	ccb.conflict = opts
	return &ComputerUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Computer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *ComputerCreateBulk) OnConflictColumns(columns ...string) *ComputerUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ComputerUpsertBulk{
		create: ccb,
	}
}

// ComputerUpsertBulk is the builder for "upsert"-ing
// a bulk of Computer nodes.
type ComputerUpsertBulk struct {
	create *ComputerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Computer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ComputerUpsertBulk) UpdateNewValues() *ComputerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Computer.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ComputerUpsertBulk) Ignore() *ComputerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ComputerUpsertBulk) DoNothing() *ComputerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ComputerCreateBulk.OnConflict
// documentation for more info.
func (u *ComputerUpsertBulk) Update(set func(*ComputerUpsert)) *ComputerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ComputerUpsert{UpdateSet: update})
	}))
	return u
}

// SetManufacturer sets the "manufacturer" field.
func (u *ComputerUpsertBulk) SetManufacturer(v string) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.SetManufacturer(v)
	})
}

// UpdateManufacturer sets the "manufacturer" field to the value that was provided on create.
func (u *ComputerUpsertBulk) UpdateManufacturer() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateManufacturer()
	})
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (u *ComputerUpsertBulk) ClearManufacturer() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearManufacturer()
	})
}

// SetModel sets the "model" field.
func (u *ComputerUpsertBulk) SetModel(v string) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.SetModel(v)
	})
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *ComputerUpsertBulk) UpdateModel() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateModel()
	})
}

// ClearModel clears the value of the "model" field.
func (u *ComputerUpsertBulk) ClearModel() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearModel()
	})
}

// SetSerial sets the "serial" field.
func (u *ComputerUpsertBulk) SetSerial(v string) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.SetSerial(v)
	})
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *ComputerUpsertBulk) UpdateSerial() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateSerial()
	})
}

// ClearSerial clears the value of the "serial" field.
func (u *ComputerUpsertBulk) ClearSerial() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearSerial()
	})
}

// SetMemory sets the "memory" field.
func (u *ComputerUpsertBulk) SetMemory(v uint64) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.SetMemory(v)
	})
}

// AddMemory adds v to the "memory" field.
func (u *ComputerUpsertBulk) AddMemory(v uint64) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.AddMemory(v)
	})
}

// UpdateMemory sets the "memory" field to the value that was provided on create.
func (u *ComputerUpsertBulk) UpdateMemory() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateMemory()
	})
}

// ClearMemory clears the value of the "memory" field.
func (u *ComputerUpsertBulk) ClearMemory() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearMemory()
	})
}

// SetProcessor sets the "processor" field.
func (u *ComputerUpsertBulk) SetProcessor(v string) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.SetProcessor(v)
	})
}

// UpdateProcessor sets the "processor" field to the value that was provided on create.
func (u *ComputerUpsertBulk) UpdateProcessor() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateProcessor()
	})
}

// ClearProcessor clears the value of the "processor" field.
func (u *ComputerUpsertBulk) ClearProcessor() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearProcessor()
	})
}

// SetProcessorCores sets the "processor_cores" field.
func (u *ComputerUpsertBulk) SetProcessorCores(v int64) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.SetProcessorCores(v)
	})
}

// AddProcessorCores adds v to the "processor_cores" field.
func (u *ComputerUpsertBulk) AddProcessorCores(v int64) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.AddProcessorCores(v)
	})
}

// UpdateProcessorCores sets the "processor_cores" field to the value that was provided on create.
func (u *ComputerUpsertBulk) UpdateProcessorCores() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateProcessorCores()
	})
}

// ClearProcessorCores clears the value of the "processor_cores" field.
func (u *ComputerUpsertBulk) ClearProcessorCores() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearProcessorCores()
	})
}

// SetProcessorArch sets the "processor_arch" field.
func (u *ComputerUpsertBulk) SetProcessorArch(v string) *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.SetProcessorArch(v)
	})
}

// UpdateProcessorArch sets the "processor_arch" field to the value that was provided on create.
func (u *ComputerUpsertBulk) UpdateProcessorArch() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.UpdateProcessorArch()
	})
}

// ClearProcessorArch clears the value of the "processor_arch" field.
func (u *ComputerUpsertBulk) ClearProcessorArch() *ComputerUpsertBulk {
	return u.Update(func(s *ComputerUpsert) {
		s.ClearProcessorArch()
	})
}

// Exec executes the query.
func (u *ComputerUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ComputerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ComputerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ComputerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
