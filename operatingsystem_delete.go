// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/operatingsystem"
	"github.com/open-uem/ent/predicate"
)

// OperatingSystemDelete is the builder for deleting a OperatingSystem entity.
type OperatingSystemDelete struct {
	config
	hooks    []Hook
	mutation *OperatingSystemMutation
}

// Where appends a list predicates to the OperatingSystemDelete builder.
func (osd *OperatingSystemDelete) Where(ps ...predicate.OperatingSystem) *OperatingSystemDelete {
	osd.mutation.Where(ps...)
	return osd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (osd *OperatingSystemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, osd.sqlExec, osd.mutation, osd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (osd *OperatingSystemDelete) ExecX(ctx context.Context) int {
	n, err := osd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (osd *OperatingSystemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(operatingsystem.Table, sqlgraph.NewFieldSpec(operatingsystem.FieldID, field.TypeInt))
	if ps := osd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, osd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	osd.mutation.done = true
	return affected, err
}

// OperatingSystemDeleteOne is the builder for deleting a single OperatingSystem entity.
type OperatingSystemDeleteOne struct {
	osd *OperatingSystemDelete
}

// Where appends a list predicates to the OperatingSystemDelete builder.
func (osdo *OperatingSystemDeleteOne) Where(ps ...predicate.OperatingSystem) *OperatingSystemDeleteOne {
	osdo.osd.mutation.Where(ps...)
	return osdo
}

// Exec executes the deletion query.
func (osdo *OperatingSystemDeleteOne) Exec(ctx context.Context) error {
	n, err := osdo.osd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{operatingsystem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (osdo *OperatingSystemDeleteOne) ExecX(ctx context.Context) {
	if err := osdo.Exec(ctx); err != nil {
		panic(err)
	}
}
