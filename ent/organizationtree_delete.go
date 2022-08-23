// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yjiong/iotdor/ent/organizationtree"
	"github.com/yjiong/iotdor/ent/predicate"
)

// OrganizationTreeDelete is the builder for deleting a OrganizationTree entity.
type OrganizationTreeDelete struct {
	config
	hooks    []Hook
	mutation *OrganizationTreeMutation
}

// Where appends a list predicates to the OrganizationTreeDelete builder.
func (otd *OrganizationTreeDelete) Where(ps ...predicate.OrganizationTree) *OrganizationTreeDelete {
	otd.mutation.Where(ps...)
	return otd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (otd *OrganizationTreeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(otd.hooks) == 0 {
		affected, err = otd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationTreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			otd.mutation = mutation
			affected, err = otd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(otd.hooks) - 1; i >= 0; i-- {
			if otd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = otd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, otd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (otd *OrganizationTreeDelete) ExecX(ctx context.Context) int {
	n, err := otd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (otd *OrganizationTreeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: organizationtree.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationtree.FieldID,
			},
		},
	}
	if ps := otd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, otd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OrganizationTreeDeleteOne is the builder for deleting a single OrganizationTree entity.
type OrganizationTreeDeleteOne struct {
	otd *OrganizationTreeDelete
}

// Exec executes the deletion query.
func (otdo *OrganizationTreeDeleteOne) Exec(ctx context.Context) error {
	n, err := otdo.otd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{organizationtree.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (otdo *OrganizationTreeDeleteOne) ExecX(ctx context.Context) {
	otdo.otd.ExecX(ctx)
}
