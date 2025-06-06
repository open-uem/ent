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
	"github.com/open-uem/ent/certificate"
	"github.com/open-uem/ent/predicate"
)

// CertificateUpdate is the builder for updating Certificate entities.
type CertificateUpdate struct {
	config
	hooks     []Hook
	mutation  *CertificateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CertificateUpdate builder.
func (cu *CertificateUpdate) Where(ps ...predicate.Certificate) *CertificateUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetType sets the "type" field.
func (cu *CertificateUpdate) SetType(c certificate.Type) *CertificateUpdate {
	cu.mutation.SetType(c)
	return cu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cu *CertificateUpdate) SetNillableType(c *certificate.Type) *CertificateUpdate {
	if c != nil {
		cu.SetType(*c)
	}
	return cu
}

// SetDescription sets the "description" field.
func (cu *CertificateUpdate) SetDescription(s string) *CertificateUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CertificateUpdate) SetNillableDescription(s *string) *CertificateUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CertificateUpdate) ClearDescription() *CertificateUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetExpiry sets the "expiry" field.
func (cu *CertificateUpdate) SetExpiry(t time.Time) *CertificateUpdate {
	cu.mutation.SetExpiry(t)
	return cu
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (cu *CertificateUpdate) SetNillableExpiry(t *time.Time) *CertificateUpdate {
	if t != nil {
		cu.SetExpiry(*t)
	}
	return cu
}

// ClearExpiry clears the value of the "expiry" field.
func (cu *CertificateUpdate) ClearExpiry() *CertificateUpdate {
	cu.mutation.ClearExpiry()
	return cu
}

// SetUID sets the "uid" field.
func (cu *CertificateUpdate) SetUID(s string) *CertificateUpdate {
	cu.mutation.SetUID(s)
	return cu
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (cu *CertificateUpdate) SetNillableUID(s *string) *CertificateUpdate {
	if s != nil {
		cu.SetUID(*s)
	}
	return cu
}

// ClearUID clears the value of the "uid" field.
func (cu *CertificateUpdate) ClearUID() *CertificateUpdate {
	cu.mutation.ClearUID()
	return cu
}

// Mutation returns the CertificateMutation object of the builder.
func (cu *CertificateUpdate) Mutation() *CertificateMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CertificateUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CertificateUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CertificateUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CertificateUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CertificateUpdate) check() error {
	if v, ok := cu.mutation.GetType(); ok {
		if err := certificate.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Certificate.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CertificateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CertificateUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CertificateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(certificate.Table, certificate.Columns, sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(certificate.FieldType, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(certificate.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(certificate.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.Expiry(); ok {
		_spec.SetField(certificate.FieldExpiry, field.TypeTime, value)
	}
	if cu.mutation.ExpiryCleared() {
		_spec.ClearField(certificate.FieldExpiry, field.TypeTime)
	}
	if value, ok := cu.mutation.UID(); ok {
		_spec.SetField(certificate.FieldUID, field.TypeString, value)
	}
	if cu.mutation.UIDCleared() {
		_spec.ClearField(certificate.FieldUID, field.TypeString)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certificate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CertificateUpdateOne is the builder for updating a single Certificate entity.
type CertificateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CertificateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetType sets the "type" field.
func (cuo *CertificateUpdateOne) SetType(c certificate.Type) *CertificateUpdateOne {
	cuo.mutation.SetType(c)
	return cuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cuo *CertificateUpdateOne) SetNillableType(c *certificate.Type) *CertificateUpdateOne {
	if c != nil {
		cuo.SetType(*c)
	}
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CertificateUpdateOne) SetDescription(s string) *CertificateUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CertificateUpdateOne) SetNillableDescription(s *string) *CertificateUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CertificateUpdateOne) ClearDescription() *CertificateUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetExpiry sets the "expiry" field.
func (cuo *CertificateUpdateOne) SetExpiry(t time.Time) *CertificateUpdateOne {
	cuo.mutation.SetExpiry(t)
	return cuo
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (cuo *CertificateUpdateOne) SetNillableExpiry(t *time.Time) *CertificateUpdateOne {
	if t != nil {
		cuo.SetExpiry(*t)
	}
	return cuo
}

// ClearExpiry clears the value of the "expiry" field.
func (cuo *CertificateUpdateOne) ClearExpiry() *CertificateUpdateOne {
	cuo.mutation.ClearExpiry()
	return cuo
}

// SetUID sets the "uid" field.
func (cuo *CertificateUpdateOne) SetUID(s string) *CertificateUpdateOne {
	cuo.mutation.SetUID(s)
	return cuo
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (cuo *CertificateUpdateOne) SetNillableUID(s *string) *CertificateUpdateOne {
	if s != nil {
		cuo.SetUID(*s)
	}
	return cuo
}

// ClearUID clears the value of the "uid" field.
func (cuo *CertificateUpdateOne) ClearUID() *CertificateUpdateOne {
	cuo.mutation.ClearUID()
	return cuo
}

// Mutation returns the CertificateMutation object of the builder.
func (cuo *CertificateUpdateOne) Mutation() *CertificateMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CertificateUpdate builder.
func (cuo *CertificateUpdateOne) Where(ps ...predicate.Certificate) *CertificateUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CertificateUpdateOne) Select(field string, fields ...string) *CertificateUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Certificate entity.
func (cuo *CertificateUpdateOne) Save(ctx context.Context) (*Certificate, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CertificateUpdateOne) SaveX(ctx context.Context) *Certificate {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CertificateUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CertificateUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CertificateUpdateOne) check() error {
	if v, ok := cuo.mutation.GetType(); ok {
		if err := certificate.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Certificate.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CertificateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CertificateUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CertificateUpdateOne) sqlSave(ctx context.Context) (_node *Certificate, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(certificate.Table, certificate.Columns, sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Certificate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certificate.FieldID)
		for _, f := range fields {
			if !certificate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != certificate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(certificate.FieldType, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(certificate.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(certificate.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.Expiry(); ok {
		_spec.SetField(certificate.FieldExpiry, field.TypeTime, value)
	}
	if cuo.mutation.ExpiryCleared() {
		_spec.ClearField(certificate.FieldExpiry, field.TypeTime)
	}
	if value, ok := cuo.mutation.UID(); ok {
		_spec.SetField(certificate.FieldUID, field.TypeString, value)
	}
	if cuo.mutation.UIDCleared() {
		_spec.ClearField(certificate.FieldUID, field.TypeString)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Certificate{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certificate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
