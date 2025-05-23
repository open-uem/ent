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
	"github.com/open-uem/ent/profile"
	"github.com/open-uem/ent/site"
	"github.com/open-uem/ent/tenant"
)

// SiteCreate is the builder for creating a Site entity.
type SiteCreate struct {
	config
	mutation *SiteMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDescription sets the "description" field.
func (sc *SiteCreate) SetDescription(s string) *SiteCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *SiteCreate) SetNillableDescription(s *string) *SiteCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetIsDefault sets the "is_default" field.
func (sc *SiteCreate) SetIsDefault(b bool) *SiteCreate {
	sc.mutation.SetIsDefault(b)
	return sc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (sc *SiteCreate) SetNillableIsDefault(b *bool) *SiteCreate {
	if b != nil {
		sc.SetIsDefault(*b)
	}
	return sc
}

// SetDomain sets the "domain" field.
func (sc *SiteCreate) SetDomain(s string) *SiteCreate {
	sc.mutation.SetDomain(s)
	return sc
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (sc *SiteCreate) SetNillableDomain(s *string) *SiteCreate {
	if s != nil {
		sc.SetDomain(*s)
	}
	return sc
}

// SetCreated sets the "created" field.
func (sc *SiteCreate) SetCreated(t time.Time) *SiteCreate {
	sc.mutation.SetCreated(t)
	return sc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (sc *SiteCreate) SetNillableCreated(t *time.Time) *SiteCreate {
	if t != nil {
		sc.SetCreated(*t)
	}
	return sc
}

// SetModified sets the "modified" field.
func (sc *SiteCreate) SetModified(t time.Time) *SiteCreate {
	sc.mutation.SetModified(t)
	return sc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (sc *SiteCreate) SetNillableModified(t *time.Time) *SiteCreate {
	if t != nil {
		sc.SetModified(*t)
	}
	return sc
}

// SetTenantID sets the "tenant" edge to the Tenant entity by ID.
func (sc *SiteCreate) SetTenantID(id int) *SiteCreate {
	sc.mutation.SetTenantID(id)
	return sc
}

// SetNillableTenantID sets the "tenant" edge to the Tenant entity by ID if the given value is not nil.
func (sc *SiteCreate) SetNillableTenantID(id *int) *SiteCreate {
	if id != nil {
		sc = sc.SetTenantID(*id)
	}
	return sc
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (sc *SiteCreate) SetTenant(t *Tenant) *SiteCreate {
	return sc.SetTenantID(t.ID)
}

// AddAgentIDs adds the "agents" edge to the Agent entity by IDs.
func (sc *SiteCreate) AddAgentIDs(ids ...string) *SiteCreate {
	sc.mutation.AddAgentIDs(ids...)
	return sc
}

// AddAgents adds the "agents" edges to the Agent entity.
func (sc *SiteCreate) AddAgents(a ...*Agent) *SiteCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAgentIDs(ids...)
}

// AddProfileIDs adds the "profiles" edge to the Profile entity by IDs.
func (sc *SiteCreate) AddProfileIDs(ids ...int) *SiteCreate {
	sc.mutation.AddProfileIDs(ids...)
	return sc
}

// AddProfiles adds the "profiles" edges to the Profile entity.
func (sc *SiteCreate) AddProfiles(p ...*Profile) *SiteCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddProfileIDs(ids...)
}

// Mutation returns the SiteMutation object of the builder.
func (sc *SiteCreate) Mutation() *SiteMutation {
	return sc.mutation
}

// Save creates the Site in the database.
func (sc *SiteCreate) Save(ctx context.Context) (*Site, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SiteCreate) SaveX(ctx context.Context) *Site {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SiteCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SiteCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SiteCreate) defaults() {
	if _, ok := sc.mutation.Created(); !ok {
		v := site.DefaultCreated()
		sc.mutation.SetCreated(v)
	}
	if _, ok := sc.mutation.Modified(); !ok {
		v := site.DefaultModified()
		sc.mutation.SetModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SiteCreate) check() error {
	return nil
}

func (sc *SiteCreate) sqlSave(ctx context.Context) (*Site, error) {
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

func (sc *SiteCreate) createSpec() (*Site, *sqlgraph.CreateSpec) {
	var (
		_node = &Site{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(site.Table, sqlgraph.NewFieldSpec(site.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(site.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.IsDefault(); ok {
		_spec.SetField(site.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if value, ok := sc.mutation.Domain(); ok {
		_spec.SetField(site.FieldDomain, field.TypeString, value)
		_node.Domain = value
	}
	if value, ok := sc.mutation.Created(); ok {
		_spec.SetField(site.FieldCreated, field.TypeTime, value)
		_node.Created = value
	}
	if value, ok := sc.mutation.Modified(); ok {
		_spec.SetField(site.FieldModified, field.TypeTime, value)
		_node.Modified = value
	}
	if nodes := sc.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   site.TenantTable,
			Columns: []string{site.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.tenant_sites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.AgentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   site.AgentsTable,
			Columns: site.AgentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ProfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   site.ProfilesTable,
			Columns: []string{site.ProfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
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
//	client.Site.Create().
//		SetDescription(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SiteUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (sc *SiteCreate) OnConflict(opts ...sql.ConflictOption) *SiteUpsertOne {
	sc.conflict = opts
	return &SiteUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Site.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SiteCreate) OnConflictColumns(columns ...string) *SiteUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SiteUpsertOne{
		create: sc,
	}
}

type (
	// SiteUpsertOne is the builder for "upsert"-ing
	//  one Site node.
	SiteUpsertOne struct {
		create *SiteCreate
	}

	// SiteUpsert is the "OnConflict" setter.
	SiteUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *SiteUpsert) SetDescription(v string) *SiteUpsert {
	u.Set(site.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SiteUpsert) UpdateDescription() *SiteUpsert {
	u.SetExcluded(site.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *SiteUpsert) ClearDescription() *SiteUpsert {
	u.SetNull(site.FieldDescription)
	return u
}

// SetIsDefault sets the "is_default" field.
func (u *SiteUpsert) SetIsDefault(v bool) *SiteUpsert {
	u.Set(site.FieldIsDefault, v)
	return u
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *SiteUpsert) UpdateIsDefault() *SiteUpsert {
	u.SetExcluded(site.FieldIsDefault)
	return u
}

// ClearIsDefault clears the value of the "is_default" field.
func (u *SiteUpsert) ClearIsDefault() *SiteUpsert {
	u.SetNull(site.FieldIsDefault)
	return u
}

// SetDomain sets the "domain" field.
func (u *SiteUpsert) SetDomain(v string) *SiteUpsert {
	u.Set(site.FieldDomain, v)
	return u
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *SiteUpsert) UpdateDomain() *SiteUpsert {
	u.SetExcluded(site.FieldDomain)
	return u
}

// ClearDomain clears the value of the "domain" field.
func (u *SiteUpsert) ClearDomain() *SiteUpsert {
	u.SetNull(site.FieldDomain)
	return u
}

// SetCreated sets the "created" field.
func (u *SiteUpsert) SetCreated(v time.Time) *SiteUpsert {
	u.Set(site.FieldCreated, v)
	return u
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *SiteUpsert) UpdateCreated() *SiteUpsert {
	u.SetExcluded(site.FieldCreated)
	return u
}

// ClearCreated clears the value of the "created" field.
func (u *SiteUpsert) ClearCreated() *SiteUpsert {
	u.SetNull(site.FieldCreated)
	return u
}

// SetModified sets the "modified" field.
func (u *SiteUpsert) SetModified(v time.Time) *SiteUpsert {
	u.Set(site.FieldModified, v)
	return u
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *SiteUpsert) UpdateModified() *SiteUpsert {
	u.SetExcluded(site.FieldModified)
	return u
}

// ClearModified clears the value of the "modified" field.
func (u *SiteUpsert) ClearModified() *SiteUpsert {
	u.SetNull(site.FieldModified)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Site.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SiteUpsertOne) UpdateNewValues() *SiteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Site.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SiteUpsertOne) Ignore() *SiteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SiteUpsertOne) DoNothing() *SiteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SiteCreate.OnConflict
// documentation for more info.
func (u *SiteUpsertOne) Update(set func(*SiteUpsert)) *SiteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SiteUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *SiteUpsertOne) SetDescription(v string) *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SiteUpsertOne) UpdateDescription() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *SiteUpsertOne) ClearDescription() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.ClearDescription()
	})
}

// SetIsDefault sets the "is_default" field.
func (u *SiteUpsertOne) SetIsDefault(v bool) *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.SetIsDefault(v)
	})
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *SiteUpsertOne) UpdateIsDefault() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateIsDefault()
	})
}

// ClearIsDefault clears the value of the "is_default" field.
func (u *SiteUpsertOne) ClearIsDefault() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.ClearIsDefault()
	})
}

// SetDomain sets the "domain" field.
func (u *SiteUpsertOne) SetDomain(v string) *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.SetDomain(v)
	})
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *SiteUpsertOne) UpdateDomain() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateDomain()
	})
}

// ClearDomain clears the value of the "domain" field.
func (u *SiteUpsertOne) ClearDomain() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.ClearDomain()
	})
}

// SetCreated sets the "created" field.
func (u *SiteUpsertOne) SetCreated(v time.Time) *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.SetCreated(v)
	})
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *SiteUpsertOne) UpdateCreated() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateCreated()
	})
}

// ClearCreated clears the value of the "created" field.
func (u *SiteUpsertOne) ClearCreated() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.ClearCreated()
	})
}

// SetModified sets the "modified" field.
func (u *SiteUpsertOne) SetModified(v time.Time) *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.SetModified(v)
	})
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *SiteUpsertOne) UpdateModified() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateModified()
	})
}

// ClearModified clears the value of the "modified" field.
func (u *SiteUpsertOne) ClearModified() *SiteUpsertOne {
	return u.Update(func(s *SiteUpsert) {
		s.ClearModified()
	})
}

// Exec executes the query.
func (u *SiteUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SiteCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SiteUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SiteUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SiteUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SiteCreateBulk is the builder for creating many Site entities in bulk.
type SiteCreateBulk struct {
	config
	err      error
	builders []*SiteCreate
	conflict []sql.ConflictOption
}

// Save creates the Site entities in the database.
func (scb *SiteCreateBulk) Save(ctx context.Context) ([]*Site, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Site, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SiteMutation)
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
func (scb *SiteCreateBulk) SaveX(ctx context.Context) []*Site {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SiteCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SiteCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Site.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SiteUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (scb *SiteCreateBulk) OnConflict(opts ...sql.ConflictOption) *SiteUpsertBulk {
	scb.conflict = opts
	return &SiteUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Site.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SiteCreateBulk) OnConflictColumns(columns ...string) *SiteUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SiteUpsertBulk{
		create: scb,
	}
}

// SiteUpsertBulk is the builder for "upsert"-ing
// a bulk of Site nodes.
type SiteUpsertBulk struct {
	create *SiteCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Site.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SiteUpsertBulk) UpdateNewValues() *SiteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Site.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SiteUpsertBulk) Ignore() *SiteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SiteUpsertBulk) DoNothing() *SiteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SiteCreateBulk.OnConflict
// documentation for more info.
func (u *SiteUpsertBulk) Update(set func(*SiteUpsert)) *SiteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SiteUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *SiteUpsertBulk) SetDescription(v string) *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SiteUpsertBulk) UpdateDescription() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *SiteUpsertBulk) ClearDescription() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.ClearDescription()
	})
}

// SetIsDefault sets the "is_default" field.
func (u *SiteUpsertBulk) SetIsDefault(v bool) *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.SetIsDefault(v)
	})
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *SiteUpsertBulk) UpdateIsDefault() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateIsDefault()
	})
}

// ClearIsDefault clears the value of the "is_default" field.
func (u *SiteUpsertBulk) ClearIsDefault() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.ClearIsDefault()
	})
}

// SetDomain sets the "domain" field.
func (u *SiteUpsertBulk) SetDomain(v string) *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.SetDomain(v)
	})
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *SiteUpsertBulk) UpdateDomain() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateDomain()
	})
}

// ClearDomain clears the value of the "domain" field.
func (u *SiteUpsertBulk) ClearDomain() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.ClearDomain()
	})
}

// SetCreated sets the "created" field.
func (u *SiteUpsertBulk) SetCreated(v time.Time) *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.SetCreated(v)
	})
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *SiteUpsertBulk) UpdateCreated() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateCreated()
	})
}

// ClearCreated clears the value of the "created" field.
func (u *SiteUpsertBulk) ClearCreated() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.ClearCreated()
	})
}

// SetModified sets the "modified" field.
func (u *SiteUpsertBulk) SetModified(v time.Time) *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.SetModified(v)
	})
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *SiteUpsertBulk) UpdateModified() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.UpdateModified()
	})
}

// ClearModified clears the value of the "modified" field.
func (u *SiteUpsertBulk) ClearModified() *SiteUpsertBulk {
	return u.Update(func(s *SiteUpsert) {
		s.ClearModified()
	})
}

// Exec executes the query.
func (u *SiteUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SiteCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SiteCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SiteUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
