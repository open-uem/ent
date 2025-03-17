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
	"github.com/open-uem/ent/profileissue"
)

// ProfileIssueCreate is the builder for creating a ProfileIssue entity.
type ProfileIssueCreate struct {
	config
	mutation *ProfileIssueMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetError sets the "error" field.
func (pic *ProfileIssueCreate) SetError(s string) *ProfileIssueCreate {
	pic.mutation.SetError(s)
	return pic
}

// SetNillableError sets the "error" field if the given value is not nil.
func (pic *ProfileIssueCreate) SetNillableError(s *string) *ProfileIssueCreate {
	if s != nil {
		pic.SetError(*s)
	}
	return pic
}

// SetWhen sets the "when" field.
func (pic *ProfileIssueCreate) SetWhen(t time.Time) *ProfileIssueCreate {
	pic.mutation.SetWhen(t)
	return pic
}

// SetNillableWhen sets the "when" field if the given value is not nil.
func (pic *ProfileIssueCreate) SetNillableWhen(t *time.Time) *ProfileIssueCreate {
	if t != nil {
		pic.SetWhen(*t)
	}
	return pic
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (pic *ProfileIssueCreate) SetProfileID(id int) *ProfileIssueCreate {
	pic.mutation.SetProfileID(id)
	return pic
}

// SetNillableProfileID sets the "profile" edge to the Profile entity by ID if the given value is not nil.
func (pic *ProfileIssueCreate) SetNillableProfileID(id *int) *ProfileIssueCreate {
	if id != nil {
		pic = pic.SetProfileID(*id)
	}
	return pic
}

// SetProfile sets the "profile" edge to the Profile entity.
func (pic *ProfileIssueCreate) SetProfile(p *Profile) *ProfileIssueCreate {
	return pic.SetProfileID(p.ID)
}

// SetAgentsID sets the "agents" edge to the Agent entity by ID.
func (pic *ProfileIssueCreate) SetAgentsID(id string) *ProfileIssueCreate {
	pic.mutation.SetAgentsID(id)
	return pic
}

// SetNillableAgentsID sets the "agents" edge to the Agent entity by ID if the given value is not nil.
func (pic *ProfileIssueCreate) SetNillableAgentsID(id *string) *ProfileIssueCreate {
	if id != nil {
		pic = pic.SetAgentsID(*id)
	}
	return pic
}

// SetAgents sets the "agents" edge to the Agent entity.
func (pic *ProfileIssueCreate) SetAgents(a *Agent) *ProfileIssueCreate {
	return pic.SetAgentsID(a.ID)
}

// Mutation returns the ProfileIssueMutation object of the builder.
func (pic *ProfileIssueCreate) Mutation() *ProfileIssueMutation {
	return pic.mutation
}

// Save creates the ProfileIssue in the database.
func (pic *ProfileIssueCreate) Save(ctx context.Context) (*ProfileIssue, error) {
	pic.defaults()
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *ProfileIssueCreate) SaveX(ctx context.Context) *ProfileIssue {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *ProfileIssueCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *ProfileIssueCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *ProfileIssueCreate) defaults() {
	if _, ok := pic.mutation.When(); !ok {
		v := profileissue.DefaultWhen
		pic.mutation.SetWhen(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *ProfileIssueCreate) check() error {
	return nil
}

func (pic *ProfileIssueCreate) sqlSave(ctx context.Context) (*ProfileIssue, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *ProfileIssueCreate) createSpec() (*ProfileIssue, *sqlgraph.CreateSpec) {
	var (
		_node = &ProfileIssue{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(profileissue.Table, sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pic.conflict
	if value, ok := pic.mutation.Error(); ok {
		_spec.SetField(profileissue.FieldError, field.TypeString, value)
		_node.Error = value
	}
	if value, ok := pic.mutation.When(); ok {
		_spec.SetField(profileissue.FieldWhen, field.TypeTime, value)
		_node.When = value
	}
	if nodes := pic.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profileissue.ProfileTable,
			Columns: []string{profileissue.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.profile_issues = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.AgentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   profileissue.AgentsTable,
			Columns: []string{profileissue.AgentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.profile_issue_agents = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProfileIssue.Create().
//		SetError(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProfileIssueUpsert) {
//			SetError(v+v).
//		}).
//		Exec(ctx)
func (pic *ProfileIssueCreate) OnConflict(opts ...sql.ConflictOption) *ProfileIssueUpsertOne {
	pic.conflict = opts
	return &ProfileIssueUpsertOne{
		create: pic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProfileIssue.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pic *ProfileIssueCreate) OnConflictColumns(columns ...string) *ProfileIssueUpsertOne {
	pic.conflict = append(pic.conflict, sql.ConflictColumns(columns...))
	return &ProfileIssueUpsertOne{
		create: pic,
	}
}

type (
	// ProfileIssueUpsertOne is the builder for "upsert"-ing
	//  one ProfileIssue node.
	ProfileIssueUpsertOne struct {
		create *ProfileIssueCreate
	}

	// ProfileIssueUpsert is the "OnConflict" setter.
	ProfileIssueUpsert struct {
		*sql.UpdateSet
	}
)

// SetError sets the "error" field.
func (u *ProfileIssueUpsert) SetError(v string) *ProfileIssueUpsert {
	u.Set(profileissue.FieldError, v)
	return u
}

// UpdateError sets the "error" field to the value that was provided on create.
func (u *ProfileIssueUpsert) UpdateError() *ProfileIssueUpsert {
	u.SetExcluded(profileissue.FieldError)
	return u
}

// ClearError clears the value of the "error" field.
func (u *ProfileIssueUpsert) ClearError() *ProfileIssueUpsert {
	u.SetNull(profileissue.FieldError)
	return u
}

// SetWhen sets the "when" field.
func (u *ProfileIssueUpsert) SetWhen(v time.Time) *ProfileIssueUpsert {
	u.Set(profileissue.FieldWhen, v)
	return u
}

// UpdateWhen sets the "when" field to the value that was provided on create.
func (u *ProfileIssueUpsert) UpdateWhen() *ProfileIssueUpsert {
	u.SetExcluded(profileissue.FieldWhen)
	return u
}

// ClearWhen clears the value of the "when" field.
func (u *ProfileIssueUpsert) ClearWhen() *ProfileIssueUpsert {
	u.SetNull(profileissue.FieldWhen)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.ProfileIssue.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProfileIssueUpsertOne) UpdateNewValues() *ProfileIssueUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProfileIssue.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProfileIssueUpsertOne) Ignore() *ProfileIssueUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProfileIssueUpsertOne) DoNothing() *ProfileIssueUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProfileIssueCreate.OnConflict
// documentation for more info.
func (u *ProfileIssueUpsertOne) Update(set func(*ProfileIssueUpsert)) *ProfileIssueUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProfileIssueUpsert{UpdateSet: update})
	}))
	return u
}

// SetError sets the "error" field.
func (u *ProfileIssueUpsertOne) SetError(v string) *ProfileIssueUpsertOne {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.SetError(v)
	})
}

// UpdateError sets the "error" field to the value that was provided on create.
func (u *ProfileIssueUpsertOne) UpdateError() *ProfileIssueUpsertOne {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.UpdateError()
	})
}

// ClearError clears the value of the "error" field.
func (u *ProfileIssueUpsertOne) ClearError() *ProfileIssueUpsertOne {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.ClearError()
	})
}

// SetWhen sets the "when" field.
func (u *ProfileIssueUpsertOne) SetWhen(v time.Time) *ProfileIssueUpsertOne {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.SetWhen(v)
	})
}

// UpdateWhen sets the "when" field to the value that was provided on create.
func (u *ProfileIssueUpsertOne) UpdateWhen() *ProfileIssueUpsertOne {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.UpdateWhen()
	})
}

// ClearWhen clears the value of the "when" field.
func (u *ProfileIssueUpsertOne) ClearWhen() *ProfileIssueUpsertOne {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.ClearWhen()
	})
}

// Exec executes the query.
func (u *ProfileIssueUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProfileIssueCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProfileIssueUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProfileIssueUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProfileIssueUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProfileIssueCreateBulk is the builder for creating many ProfileIssue entities in bulk.
type ProfileIssueCreateBulk struct {
	config
	err      error
	builders []*ProfileIssueCreate
	conflict []sql.ConflictOption
}

// Save creates the ProfileIssue entities in the database.
func (picb *ProfileIssueCreateBulk) Save(ctx context.Context) ([]*ProfileIssue, error) {
	if picb.err != nil {
		return nil, picb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*ProfileIssue, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileIssueMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = picb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *ProfileIssueCreateBulk) SaveX(ctx context.Context) []*ProfileIssue {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *ProfileIssueCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *ProfileIssueCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProfileIssue.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProfileIssueUpsert) {
//			SetError(v+v).
//		}).
//		Exec(ctx)
func (picb *ProfileIssueCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProfileIssueUpsertBulk {
	picb.conflict = opts
	return &ProfileIssueUpsertBulk{
		create: picb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProfileIssue.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (picb *ProfileIssueCreateBulk) OnConflictColumns(columns ...string) *ProfileIssueUpsertBulk {
	picb.conflict = append(picb.conflict, sql.ConflictColumns(columns...))
	return &ProfileIssueUpsertBulk{
		create: picb,
	}
}

// ProfileIssueUpsertBulk is the builder for "upsert"-ing
// a bulk of ProfileIssue nodes.
type ProfileIssueUpsertBulk struct {
	create *ProfileIssueCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ProfileIssue.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProfileIssueUpsertBulk) UpdateNewValues() *ProfileIssueUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProfileIssue.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProfileIssueUpsertBulk) Ignore() *ProfileIssueUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProfileIssueUpsertBulk) DoNothing() *ProfileIssueUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProfileIssueCreateBulk.OnConflict
// documentation for more info.
func (u *ProfileIssueUpsertBulk) Update(set func(*ProfileIssueUpsert)) *ProfileIssueUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProfileIssueUpsert{UpdateSet: update})
	}))
	return u
}

// SetError sets the "error" field.
func (u *ProfileIssueUpsertBulk) SetError(v string) *ProfileIssueUpsertBulk {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.SetError(v)
	})
}

// UpdateError sets the "error" field to the value that was provided on create.
func (u *ProfileIssueUpsertBulk) UpdateError() *ProfileIssueUpsertBulk {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.UpdateError()
	})
}

// ClearError clears the value of the "error" field.
func (u *ProfileIssueUpsertBulk) ClearError() *ProfileIssueUpsertBulk {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.ClearError()
	})
}

// SetWhen sets the "when" field.
func (u *ProfileIssueUpsertBulk) SetWhen(v time.Time) *ProfileIssueUpsertBulk {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.SetWhen(v)
	})
}

// UpdateWhen sets the "when" field to the value that was provided on create.
func (u *ProfileIssueUpsertBulk) UpdateWhen() *ProfileIssueUpsertBulk {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.UpdateWhen()
	})
}

// ClearWhen clears the value of the "when" field.
func (u *ProfileIssueUpsertBulk) ClearWhen() *ProfileIssueUpsertBulk {
	return u.Update(func(s *ProfileIssueUpsert) {
		s.ClearWhen()
	})
}

// Exec executes the query.
func (u *ProfileIssueUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProfileIssueCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProfileIssueCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProfileIssueUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
