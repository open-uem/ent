// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/predicate"
	"github.com/open-uem/ent/wingetconfigexclusion"
)

// WingetConfigExclusionDelete is the builder for deleting a WingetConfigExclusion entity.
type WingetConfigExclusionDelete struct {
	config
	hooks    []Hook
	mutation *WingetConfigExclusionMutation
}

// Where appends a list predicates to the WingetConfigExclusionDelete builder.
func (wced *WingetConfigExclusionDelete) Where(ps ...predicate.WingetConfigExclusion) *WingetConfigExclusionDelete {
	wced.mutation.Where(ps...)
	return wced
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wced *WingetConfigExclusionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wced.sqlExec, wced.mutation, wced.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wced *WingetConfigExclusionDelete) ExecX(ctx context.Context) int {
	n, err := wced.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wced *WingetConfigExclusionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wingetconfigexclusion.Table, sqlgraph.NewFieldSpec(wingetconfigexclusion.FieldID, field.TypeInt))
	if ps := wced.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wced.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wced.mutation.done = true
	return affected, err
}

// WingetConfigExclusionDeleteOne is the builder for deleting a single WingetConfigExclusion entity.
type WingetConfigExclusionDeleteOne struct {
	wced *WingetConfigExclusionDelete
}

// Where appends a list predicates to the WingetConfigExclusionDelete builder.
func (wcedo *WingetConfigExclusionDeleteOne) Where(ps ...predicate.WingetConfigExclusion) *WingetConfigExclusionDeleteOne {
	wcedo.wced.mutation.Where(ps...)
	return wcedo
}

// Exec executes the deletion query.
func (wcedo *WingetConfigExclusionDeleteOne) Exec(ctx context.Context) error {
	n, err := wcedo.wced.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wingetconfigexclusion.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wcedo *WingetConfigExclusionDeleteOne) ExecX(ctx context.Context) {
	if err := wcedo.Exec(ctx); err != nil {
		panic(err)
	}
}
