// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/restService/pkg/ent/enrichedfio"
	"github.com/wtkeqrf0/restService/pkg/ent/predicate"
)

// EnrichedFioDelete is the builder for deleting a EnrichedFio entity.
type EnrichedFioDelete struct {
	config
	hooks    []Hook
	mutation *EnrichedFioMutation
}

// Where appends a list predicates to the EnrichedFioDelete builder.
func (efd *EnrichedFioDelete) Where(ps ...predicate.EnrichedFio) *EnrichedFioDelete {
	efd.mutation.Where(ps...)
	return efd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (efd *EnrichedFioDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, efd.sqlExec, efd.mutation, efd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (efd *EnrichedFioDelete) ExecX(ctx context.Context) int {
	n, err := efd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (efd *EnrichedFioDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(enrichedfio.Table, sqlgraph.NewFieldSpec(enrichedfio.FieldID, field.TypeInt))
	if ps := efd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, efd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	efd.mutation.done = true
	return affected, err
}

// EnrichedFioDeleteOne is the builder for deleting a single EnrichedFio entity.
type EnrichedFioDeleteOne struct {
	efd *EnrichedFioDelete
}

// Where appends a list predicates to the EnrichedFioDelete builder.
func (efdo *EnrichedFioDeleteOne) Where(ps ...predicate.EnrichedFio) *EnrichedFioDeleteOne {
	efdo.efd.mutation.Where(ps...)
	return efdo
}

// Exec executes the deletion query.
func (efdo *EnrichedFioDeleteOne) Exec(ctx context.Context) error {
	n, err := efdo.efd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{enrichedfio.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (efdo *EnrichedFioDeleteOne) ExecX(ctx context.Context) {
	if err := efdo.Exec(ctx); err != nil {
		panic(err)
	}
}
