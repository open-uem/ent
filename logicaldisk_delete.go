// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/logicaldisk"
	"github.com/open-uem/ent/predicate"
)

// LogicalDiskDelete is the builder for deleting a LogicalDisk entity.
type LogicalDiskDelete struct {
	config
	hooks    []Hook
	mutation *LogicalDiskMutation
}

// Where appends a list predicates to the LogicalDiskDelete builder.
func (ldd *LogicalDiskDelete) Where(ps ...predicate.LogicalDisk) *LogicalDiskDelete {
	ldd.mutation.Where(ps...)
	return ldd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ldd *LogicalDiskDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ldd.sqlExec, ldd.mutation, ldd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ldd *LogicalDiskDelete) ExecX(ctx context.Context) int {
	n, err := ldd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ldd *LogicalDiskDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(logicaldisk.Table, sqlgraph.NewFieldSpec(logicaldisk.FieldID, field.TypeInt))
	if ps := ldd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ldd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ldd.mutation.done = true
	return affected, err
}

// LogicalDiskDeleteOne is the builder for deleting a single LogicalDisk entity.
type LogicalDiskDeleteOne struct {
	ldd *LogicalDiskDelete
}

// Where appends a list predicates to the LogicalDiskDelete builder.
func (lddo *LogicalDiskDeleteOne) Where(ps ...predicate.LogicalDisk) *LogicalDiskDeleteOne {
	lddo.ldd.mutation.Where(ps...)
	return lddo
}

// Exec executes the deletion query.
func (lddo *LogicalDiskDeleteOne) Exec(ctx context.Context) error {
	n, err := lddo.ldd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{logicaldisk.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lddo *LogicalDiskDeleteOne) ExecX(ctx context.Context) {
	if err := lddo.Exec(ctx); err != nil {
		panic(err)
	}
}
