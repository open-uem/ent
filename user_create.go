// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem_ent/sessions"
	"github.com/doncicuto/openuem_ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetCsr sets the "csr" field.
func (uc *UserCreate) SetCsr(s string) *UserCreate {
	uc.mutation.SetCsr(s)
	return uc
}

// SetNillableCsr sets the "csr" field if the given value is not nil.
func (uc *UserCreate) SetNillableCsr(s *string) *UserCreate {
	if s != nil {
		uc.SetCsr(*s)
	}
	return uc
}

// SetCertSerial sets the "certSerial" field.
func (uc *UserCreate) SetCertSerial(s string) *UserCreate {
	uc.mutation.SetCertSerial(s)
	return uc
}

// SetNillableCertSerial sets the "certSerial" field if the given value is not nil.
func (uc *UserCreate) SetNillableCertSerial(s *string) *UserCreate {
	if s != nil {
		uc.SetCertSerial(*s)
	}
	return uc
}

// SetRegister sets the "register" field.
func (uc *UserCreate) SetRegister(s string) *UserCreate {
	uc.mutation.SetRegister(s)
	return uc
}

// SetNillableRegister sets the "register" field if the given value is not nil.
func (uc *UserCreate) SetNillableRegister(s *string) *UserCreate {
	if s != nil {
		uc.SetRegister(*s)
	}
	return uc
}

// SetExpiry sets the "expiry" field.
func (uc *UserCreate) SetExpiry(t time.Time) *UserCreate {
	uc.mutation.SetExpiry(t)
	return uc
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (uc *UserCreate) SetNillableExpiry(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetExpiry(*t)
	}
	return uc
}

// SetCreated sets the "created" field.
func (uc *UserCreate) SetCreated(t time.Time) *UserCreate {
	uc.mutation.SetCreated(t)
	return uc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreated(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreated(*t)
	}
	return uc
}

// SetModified sets the "modified" field.
func (uc *UserCreate) SetModified(t time.Time) *UserCreate {
	uc.mutation.SetModified(t)
	return uc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (uc *UserCreate) SetNillableModified(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetModified(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddSessionIDs adds the "sessions" edge to the Sessions entity by IDs.
func (uc *UserCreate) AddSessionIDs(ids ...string) *UserCreate {
	uc.mutation.AddSessionIDs(ids...)
	return uc
}

// AddSessions adds the "sessions" edges to the Sessions entity.
func (uc *UserCreate) AddSessions(s ...*Sessions) *UserCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddSessionIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Register(); !ok {
		v := user.DefaultRegister
		uc.mutation.SetRegister(v)
	}
	if _, ok := uc.mutation.Created(); !ok {
		v := user.DefaultCreated()
		uc.mutation.SetCreated(v)
	}
	if _, ok := uc.mutation.Modified(); !ok {
		v := user.DefaultModified()
		uc.mutation.SetModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`openuem_ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.Register(); !ok {
		return &ValidationError{Name: "register", err: errors.New(`openuem_ent: missing required field "User.register"`)}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := user.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`openuem_ent: validator failed for field "User.id": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := uc.mutation.Csr(); ok {
		_spec.SetField(user.FieldCsr, field.TypeString, value)
		_node.Csr = value
	}
	if value, ok := uc.mutation.CertSerial(); ok {
		_spec.SetField(user.FieldCertSerial, field.TypeString, value)
		_node.CertSerial = value
	}
	if value, ok := uc.mutation.Register(); ok {
		_spec.SetField(user.FieldRegister, field.TypeString, value)
		_node.Register = value
	}
	if value, ok := uc.mutation.Expiry(); ok {
		_spec.SetField(user.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	if value, ok := uc.mutation.Created(); ok {
		_spec.SetField(user.FieldCreated, field.TypeTime, value)
		_node.Created = value
	}
	if value, ok := uc.mutation.Modified(); ok {
		_spec.SetField(user.FieldModified, field.TypeTime, value)
		_node.Modified = value
	}
	if nodes := uc.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeString),
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
//	client.User.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *UserUpsert) SetName(v string) *UserUpsert {
	u.Set(user.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsert) UpdateName() *UserUpsert {
	u.SetExcluded(user.FieldName)
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsert) ClearEmail() *UserUpsert {
	u.SetNull(user.FieldEmail)
	return u
}

// SetPhone sets the "phone" field.
func (u *UserUpsert) SetPhone(v string) *UserUpsert {
	u.Set(user.FieldPhone, v)
	return u
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *UserUpsert) UpdatePhone() *UserUpsert {
	u.SetExcluded(user.FieldPhone)
	return u
}

// ClearPhone clears the value of the "phone" field.
func (u *UserUpsert) ClearPhone() *UserUpsert {
	u.SetNull(user.FieldPhone)
	return u
}

// SetCsr sets the "csr" field.
func (u *UserUpsert) SetCsr(v string) *UserUpsert {
	u.Set(user.FieldCsr, v)
	return u
}

// UpdateCsr sets the "csr" field to the value that was provided on create.
func (u *UserUpsert) UpdateCsr() *UserUpsert {
	u.SetExcluded(user.FieldCsr)
	return u
}

// ClearCsr clears the value of the "csr" field.
func (u *UserUpsert) ClearCsr() *UserUpsert {
	u.SetNull(user.FieldCsr)
	return u
}

// SetCertSerial sets the "certSerial" field.
func (u *UserUpsert) SetCertSerial(v string) *UserUpsert {
	u.Set(user.FieldCertSerial, v)
	return u
}

// UpdateCertSerial sets the "certSerial" field to the value that was provided on create.
func (u *UserUpsert) UpdateCertSerial() *UserUpsert {
	u.SetExcluded(user.FieldCertSerial)
	return u
}

// ClearCertSerial clears the value of the "certSerial" field.
func (u *UserUpsert) ClearCertSerial() *UserUpsert {
	u.SetNull(user.FieldCertSerial)
	return u
}

// SetRegister sets the "register" field.
func (u *UserUpsert) SetRegister(v string) *UserUpsert {
	u.Set(user.FieldRegister, v)
	return u
}

// UpdateRegister sets the "register" field to the value that was provided on create.
func (u *UserUpsert) UpdateRegister() *UserUpsert {
	u.SetExcluded(user.FieldRegister)
	return u
}

// SetExpiry sets the "expiry" field.
func (u *UserUpsert) SetExpiry(v time.Time) *UserUpsert {
	u.Set(user.FieldExpiry, v)
	return u
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *UserUpsert) UpdateExpiry() *UserUpsert {
	u.SetExcluded(user.FieldExpiry)
	return u
}

// ClearExpiry clears the value of the "expiry" field.
func (u *UserUpsert) ClearExpiry() *UserUpsert {
	u.SetNull(user.FieldExpiry)
	return u
}

// SetCreated sets the "created" field.
func (u *UserUpsert) SetCreated(v time.Time) *UserUpsert {
	u.Set(user.FieldCreated, v)
	return u
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *UserUpsert) UpdateCreated() *UserUpsert {
	u.SetExcluded(user.FieldCreated)
	return u
}

// ClearCreated clears the value of the "created" field.
func (u *UserUpsert) ClearCreated() *UserUpsert {
	u.SetNull(user.FieldCreated)
	return u
}

// SetModified sets the "modified" field.
func (u *UserUpsert) SetModified(v time.Time) *UserUpsert {
	u.Set(user.FieldModified, v)
	return u
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *UserUpsert) UpdateModified() *UserUpsert {
	u.SetExcluded(user.FieldModified)
	return u
}

// ClearModified clears the value of the "modified" field.
func (u *UserUpsert) ClearModified() *UserUpsert {
	u.SetNull(user.FieldModified)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *UserUpsertOne) SetName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsertOne) ClearEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearEmail()
	})
}

// SetPhone sets the "phone" field.
func (u *UserUpsertOne) SetPhone(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePhone() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePhone()
	})
}

// ClearPhone clears the value of the "phone" field.
func (u *UserUpsertOne) ClearPhone() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearPhone()
	})
}

// SetCsr sets the "csr" field.
func (u *UserUpsertOne) SetCsr(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCsr(v)
	})
}

// UpdateCsr sets the "csr" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCsr() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCsr()
	})
}

// ClearCsr clears the value of the "csr" field.
func (u *UserUpsertOne) ClearCsr() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearCsr()
	})
}

// SetCertSerial sets the "certSerial" field.
func (u *UserUpsertOne) SetCertSerial(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCertSerial(v)
	})
}

// UpdateCertSerial sets the "certSerial" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCertSerial() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCertSerial()
	})
}

// ClearCertSerial clears the value of the "certSerial" field.
func (u *UserUpsertOne) ClearCertSerial() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearCertSerial()
	})
}

// SetRegister sets the "register" field.
func (u *UserUpsertOne) SetRegister(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetRegister(v)
	})
}

// UpdateRegister sets the "register" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateRegister() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateRegister()
	})
}

// SetExpiry sets the "expiry" field.
func (u *UserUpsertOne) SetExpiry(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetExpiry(v)
	})
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateExpiry() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateExpiry()
	})
}

// ClearExpiry clears the value of the "expiry" field.
func (u *UserUpsertOne) ClearExpiry() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearExpiry()
	})
}

// SetCreated sets the "created" field.
func (u *UserUpsertOne) SetCreated(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCreated(v)
	})
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCreated() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreated()
	})
}

// ClearCreated clears the value of the "created" field.
func (u *UserUpsertOne) ClearCreated() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearCreated()
	})
}

// SetModified sets the "modified" field.
func (u *UserUpsertOne) SetModified(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetModified(v)
	})
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateModified() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateModified()
	})
}

// ClearModified clears the value of the "modified" field.
func (u *UserUpsertOne) ClearModified() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearModified()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("openuem_ent: UserUpsertOne.ID is not supported by MySQL driver. Use UserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *UserUpsertBulk) SetName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsertBulk) ClearEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearEmail()
	})
}

// SetPhone sets the "phone" field.
func (u *UserUpsertBulk) SetPhone(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePhone() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePhone()
	})
}

// ClearPhone clears the value of the "phone" field.
func (u *UserUpsertBulk) ClearPhone() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearPhone()
	})
}

// SetCsr sets the "csr" field.
func (u *UserUpsertBulk) SetCsr(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCsr(v)
	})
}

// UpdateCsr sets the "csr" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCsr() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCsr()
	})
}

// ClearCsr clears the value of the "csr" field.
func (u *UserUpsertBulk) ClearCsr() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearCsr()
	})
}

// SetCertSerial sets the "certSerial" field.
func (u *UserUpsertBulk) SetCertSerial(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCertSerial(v)
	})
}

// UpdateCertSerial sets the "certSerial" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCertSerial() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCertSerial()
	})
}

// ClearCertSerial clears the value of the "certSerial" field.
func (u *UserUpsertBulk) ClearCertSerial() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearCertSerial()
	})
}

// SetRegister sets the "register" field.
func (u *UserUpsertBulk) SetRegister(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetRegister(v)
	})
}

// UpdateRegister sets the "register" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateRegister() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateRegister()
	})
}

// SetExpiry sets the "expiry" field.
func (u *UserUpsertBulk) SetExpiry(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetExpiry(v)
	})
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateExpiry() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateExpiry()
	})
}

// ClearExpiry clears the value of the "expiry" field.
func (u *UserUpsertBulk) ClearExpiry() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearExpiry()
	})
}

// SetCreated sets the "created" field.
func (u *UserUpsertBulk) SetCreated(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCreated(v)
	})
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCreated() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreated()
	})
}

// ClearCreated clears the value of the "created" field.
func (u *UserUpsertBulk) ClearCreated() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearCreated()
	})
}

// SetModified sets the "modified" field.
func (u *UserUpsertBulk) SetModified(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetModified(v)
	})
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateModified() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateModified()
	})
}

// ClearModified clears the value of the "modified" field.
func (u *UserUpsertBulk) ClearModified() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearModified()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("openuem_ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
