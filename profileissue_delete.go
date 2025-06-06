// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/predicate"
	"github.com/open-uem/ent/profileissue"
)

// ProfileIssueDelete is the builder for deleting a ProfileIssue entity.
type ProfileIssueDelete struct {
	config
	hooks    []Hook
	mutation *ProfileIssueMutation
}

// Where appends a list predicates to the ProfileIssueDelete builder.
func (pid *ProfileIssueDelete) Where(ps ...predicate.ProfileIssue) *ProfileIssueDelete {
	pid.mutation.Where(ps...)
	return pid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pid *ProfileIssueDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pid.sqlExec, pid.mutation, pid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pid *ProfileIssueDelete) ExecX(ctx context.Context) int {
	n, err := pid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pid *ProfileIssueDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(profileissue.Table, sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt))
	if ps := pid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pid.mutation.done = true
	return affected, err
}

// ProfileIssueDeleteOne is the builder for deleting a single ProfileIssue entity.
type ProfileIssueDeleteOne struct {
	pid *ProfileIssueDelete
}

// Where appends a list predicates to the ProfileIssueDelete builder.
func (pido *ProfileIssueDeleteOne) Where(ps ...predicate.ProfileIssue) *ProfileIssueDeleteOne {
	pido.pid.mutation.Where(ps...)
	return pido
}

// Exec executes the deletion query.
func (pido *ProfileIssueDeleteOne) Exec(ctx context.Context) error {
	n, err := pido.pid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{profileissue.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pido *ProfileIssueDeleteOne) ExecX(ctx context.Context) {
	if err := pido.Exec(ctx); err != nil {
		panic(err)
	}
}
