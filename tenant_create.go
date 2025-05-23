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
	"github.com/open-uem/ent/orgmetadata"
	"github.com/open-uem/ent/settings"
	"github.com/open-uem/ent/site"
	"github.com/open-uem/ent/tag"
	"github.com/open-uem/ent/tenant"
)

// TenantCreate is the builder for creating a Tenant entity.
type TenantCreate struct {
	config
	mutation *TenantMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDescription sets the "description" field.
func (tc *TenantCreate) SetDescription(s string) *TenantCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TenantCreate) SetNillableDescription(s *string) *TenantCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetIsDefault sets the "is_default" field.
func (tc *TenantCreate) SetIsDefault(b bool) *TenantCreate {
	tc.mutation.SetIsDefault(b)
	return tc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (tc *TenantCreate) SetNillableIsDefault(b *bool) *TenantCreate {
	if b != nil {
		tc.SetIsDefault(*b)
	}
	return tc
}

// SetCreated sets the "created" field.
func (tc *TenantCreate) SetCreated(t time.Time) *TenantCreate {
	tc.mutation.SetCreated(t)
	return tc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (tc *TenantCreate) SetNillableCreated(t *time.Time) *TenantCreate {
	if t != nil {
		tc.SetCreated(*t)
	}
	return tc
}

// SetModified sets the "modified" field.
func (tc *TenantCreate) SetModified(t time.Time) *TenantCreate {
	tc.mutation.SetModified(t)
	return tc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (tc *TenantCreate) SetNillableModified(t *time.Time) *TenantCreate {
	if t != nil {
		tc.SetModified(*t)
	}
	return tc
}

// AddSiteIDs adds the "sites" edge to the Site entity by IDs.
func (tc *TenantCreate) AddSiteIDs(ids ...int) *TenantCreate {
	tc.mutation.AddSiteIDs(ids...)
	return tc
}

// AddSites adds the "sites" edges to the Site entity.
func (tc *TenantCreate) AddSites(s ...*Site) *TenantCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddSiteIDs(ids...)
}

// SetSettingsID sets the "settings" edge to the Settings entity by ID.
func (tc *TenantCreate) SetSettingsID(id int) *TenantCreate {
	tc.mutation.SetSettingsID(id)
	return tc
}

// SetNillableSettingsID sets the "settings" edge to the Settings entity by ID if the given value is not nil.
func (tc *TenantCreate) SetNillableSettingsID(id *int) *TenantCreate {
	if id != nil {
		tc = tc.SetSettingsID(*id)
	}
	return tc
}

// SetSettings sets the "settings" edge to the Settings entity.
func (tc *TenantCreate) SetSettings(s *Settings) *TenantCreate {
	return tc.SetSettingsID(s.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tc *TenantCreate) AddTagIDs(ids ...int) *TenantCreate {
	tc.mutation.AddTagIDs(ids...)
	return tc
}

// AddTags adds the "tags" edges to the Tag entity.
func (tc *TenantCreate) AddTags(t ...*Tag) *TenantCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTagIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the OrgMetadata entity by IDs.
func (tc *TenantCreate) AddMetadatumIDs(ids ...int) *TenantCreate {
	tc.mutation.AddMetadatumIDs(ids...)
	return tc
}

// AddMetadata adds the "metadata" edges to the OrgMetadata entity.
func (tc *TenantCreate) AddMetadata(o ...*OrgMetadata) *TenantCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return tc.AddMetadatumIDs(ids...)
}

// Mutation returns the TenantMutation object of the builder.
func (tc *TenantCreate) Mutation() *TenantMutation {
	return tc.mutation
}

// Save creates the Tenant in the database.
func (tc *TenantCreate) Save(ctx context.Context) (*Tenant, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TenantCreate) SaveX(ctx context.Context) *Tenant {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TenantCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TenantCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TenantCreate) defaults() {
	if _, ok := tc.mutation.Created(); !ok {
		v := tenant.DefaultCreated()
		tc.mutation.SetCreated(v)
	}
	if _, ok := tc.mutation.Modified(); !ok {
		v := tenant.DefaultModified()
		tc.mutation.SetModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TenantCreate) check() error {
	return nil
}

func (tc *TenantCreate) sqlSave(ctx context.Context) (*Tenant, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TenantCreate) createSpec() (*Tenant, *sqlgraph.CreateSpec) {
	var (
		_node = &Tenant{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tenant.Table, sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(tenant.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.IsDefault(); ok {
		_spec.SetField(tenant.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if value, ok := tc.mutation.Created(); ok {
		_spec.SetField(tenant.FieldCreated, field.TypeTime, value)
		_node.Created = value
	}
	if value, ok := tc.mutation.Modified(); ok {
		_spec.SetField(tenant.FieldModified, field.TypeTime, value)
		_node.Modified = value
	}
	if nodes := tc.mutation.SitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tenant.SitesTable,
			Columns: []string{tenant.SitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(site.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tenant.SettingsTable,
			Columns: []string{tenant.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tenant.TagsTable,
			Columns: []string{tenant.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tenant.MetadataTable,
			Columns: []string{tenant.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgmetadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tenant.Create().
//		SetDescription(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TenantUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (tc *TenantCreate) OnConflict(opts ...sql.ConflictOption) *TenantUpsertOne {
	tc.conflict = opts
	return &TenantUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tenant.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TenantCreate) OnConflictColumns(columns ...string) *TenantUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TenantUpsertOne{
		create: tc,
	}
}

type (
	// TenantUpsertOne is the builder for "upsert"-ing
	//  one Tenant node.
	TenantUpsertOne struct {
		create *TenantCreate
	}

	// TenantUpsert is the "OnConflict" setter.
	TenantUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *TenantUpsert) SetDescription(v string) *TenantUpsert {
	u.Set(tenant.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TenantUpsert) UpdateDescription() *TenantUpsert {
	u.SetExcluded(tenant.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *TenantUpsert) ClearDescription() *TenantUpsert {
	u.SetNull(tenant.FieldDescription)
	return u
}

// SetIsDefault sets the "is_default" field.
func (u *TenantUpsert) SetIsDefault(v bool) *TenantUpsert {
	u.Set(tenant.FieldIsDefault, v)
	return u
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *TenantUpsert) UpdateIsDefault() *TenantUpsert {
	u.SetExcluded(tenant.FieldIsDefault)
	return u
}

// ClearIsDefault clears the value of the "is_default" field.
func (u *TenantUpsert) ClearIsDefault() *TenantUpsert {
	u.SetNull(tenant.FieldIsDefault)
	return u
}

// SetCreated sets the "created" field.
func (u *TenantUpsert) SetCreated(v time.Time) *TenantUpsert {
	u.Set(tenant.FieldCreated, v)
	return u
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *TenantUpsert) UpdateCreated() *TenantUpsert {
	u.SetExcluded(tenant.FieldCreated)
	return u
}

// ClearCreated clears the value of the "created" field.
func (u *TenantUpsert) ClearCreated() *TenantUpsert {
	u.SetNull(tenant.FieldCreated)
	return u
}

// SetModified sets the "modified" field.
func (u *TenantUpsert) SetModified(v time.Time) *TenantUpsert {
	u.Set(tenant.FieldModified, v)
	return u
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *TenantUpsert) UpdateModified() *TenantUpsert {
	u.SetExcluded(tenant.FieldModified)
	return u
}

// ClearModified clears the value of the "modified" field.
func (u *TenantUpsert) ClearModified() *TenantUpsert {
	u.SetNull(tenant.FieldModified)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Tenant.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TenantUpsertOne) UpdateNewValues() *TenantUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tenant.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TenantUpsertOne) Ignore() *TenantUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TenantUpsertOne) DoNothing() *TenantUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TenantCreate.OnConflict
// documentation for more info.
func (u *TenantUpsertOne) Update(set func(*TenantUpsert)) *TenantUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TenantUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *TenantUpsertOne) SetDescription(v string) *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TenantUpsertOne) UpdateDescription() *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *TenantUpsertOne) ClearDescription() *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.ClearDescription()
	})
}

// SetIsDefault sets the "is_default" field.
func (u *TenantUpsertOne) SetIsDefault(v bool) *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.SetIsDefault(v)
	})
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *TenantUpsertOne) UpdateIsDefault() *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.UpdateIsDefault()
	})
}

// ClearIsDefault clears the value of the "is_default" field.
func (u *TenantUpsertOne) ClearIsDefault() *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.ClearIsDefault()
	})
}

// SetCreated sets the "created" field.
func (u *TenantUpsertOne) SetCreated(v time.Time) *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.SetCreated(v)
	})
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *TenantUpsertOne) UpdateCreated() *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.UpdateCreated()
	})
}

// ClearCreated clears the value of the "created" field.
func (u *TenantUpsertOne) ClearCreated() *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.ClearCreated()
	})
}

// SetModified sets the "modified" field.
func (u *TenantUpsertOne) SetModified(v time.Time) *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.SetModified(v)
	})
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *TenantUpsertOne) UpdateModified() *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.UpdateModified()
	})
}

// ClearModified clears the value of the "modified" field.
func (u *TenantUpsertOne) ClearModified() *TenantUpsertOne {
	return u.Update(func(s *TenantUpsert) {
		s.ClearModified()
	})
}

// Exec executes the query.
func (u *TenantUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TenantCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TenantUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TenantUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TenantUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TenantCreateBulk is the builder for creating many Tenant entities in bulk.
type TenantCreateBulk struct {
	config
	err      error
	builders []*TenantCreate
	conflict []sql.ConflictOption
}

// Save creates the Tenant entities in the database.
func (tcb *TenantCreateBulk) Save(ctx context.Context) ([]*Tenant, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tenant, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TenantMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TenantCreateBulk) SaveX(ctx context.Context) []*Tenant {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TenantCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TenantCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tenant.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TenantUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (tcb *TenantCreateBulk) OnConflict(opts ...sql.ConflictOption) *TenantUpsertBulk {
	tcb.conflict = opts
	return &TenantUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tenant.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TenantCreateBulk) OnConflictColumns(columns ...string) *TenantUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TenantUpsertBulk{
		create: tcb,
	}
}

// TenantUpsertBulk is the builder for "upsert"-ing
// a bulk of Tenant nodes.
type TenantUpsertBulk struct {
	create *TenantCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Tenant.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TenantUpsertBulk) UpdateNewValues() *TenantUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tenant.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TenantUpsertBulk) Ignore() *TenantUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TenantUpsertBulk) DoNothing() *TenantUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TenantCreateBulk.OnConflict
// documentation for more info.
func (u *TenantUpsertBulk) Update(set func(*TenantUpsert)) *TenantUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TenantUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *TenantUpsertBulk) SetDescription(v string) *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TenantUpsertBulk) UpdateDescription() *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *TenantUpsertBulk) ClearDescription() *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.ClearDescription()
	})
}

// SetIsDefault sets the "is_default" field.
func (u *TenantUpsertBulk) SetIsDefault(v bool) *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.SetIsDefault(v)
	})
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *TenantUpsertBulk) UpdateIsDefault() *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.UpdateIsDefault()
	})
}

// ClearIsDefault clears the value of the "is_default" field.
func (u *TenantUpsertBulk) ClearIsDefault() *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.ClearIsDefault()
	})
}

// SetCreated sets the "created" field.
func (u *TenantUpsertBulk) SetCreated(v time.Time) *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.SetCreated(v)
	})
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *TenantUpsertBulk) UpdateCreated() *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.UpdateCreated()
	})
}

// ClearCreated clears the value of the "created" field.
func (u *TenantUpsertBulk) ClearCreated() *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.ClearCreated()
	})
}

// SetModified sets the "modified" field.
func (u *TenantUpsertBulk) SetModified(v time.Time) *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.SetModified(v)
	})
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *TenantUpsertBulk) UpdateModified() *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.UpdateModified()
	})
}

// ClearModified clears the value of the "modified" field.
func (u *TenantUpsertBulk) ClearModified() *TenantUpsertBulk {
	return u.Update(func(s *TenantUpsert) {
		s.ClearModified()
	})
}

// Exec executes the query.
func (u *TenantUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TenantCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TenantCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TenantUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
