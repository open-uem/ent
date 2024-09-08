// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem-ent/ent/antivirus"
	"github.com/doncicuto/openuem-ent/ent/predicate"
)

// AntivirusDelete is the builder for deleting a Antivirus entity.
type AntivirusDelete struct {
	config
	hooks    []Hook
	mutation *AntivirusMutation
}

// Where appends a list predicates to the AntivirusDelete builder.
func (ad *AntivirusDelete) Where(ps ...predicate.Antivirus) *AntivirusDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AntivirusDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AntivirusDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AntivirusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(antivirus.Table, sqlgraph.NewFieldSpec(antivirus.FieldID, field.TypeInt))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AntivirusDeleteOne is the builder for deleting a single Antivirus entity.
type AntivirusDeleteOne struct {
	ad *AntivirusDelete
}

// Where appends a list predicates to the AntivirusDelete builder.
func (ado *AntivirusDeleteOne) Where(ps ...predicate.Antivirus) *AntivirusDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AntivirusDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{antivirus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AntivirusDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
