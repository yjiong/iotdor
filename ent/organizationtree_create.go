// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yjiong/iotdor/ent/organizationposition"
	"github.com/yjiong/iotdor/ent/organizationtree"
)

// OrganizationTreeCreate is the builder for creating a OrganizationTree entity.
type OrganizationTreeCreate struct {
	config
	mutation *OrganizationTreeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (otc *OrganizationTreeCreate) SetCreateTime(t time.Time) *OrganizationTreeCreate {
	otc.mutation.SetCreateTime(t)
	return otc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (otc *OrganizationTreeCreate) SetNillableCreateTime(t *time.Time) *OrganizationTreeCreate {
	if t != nil {
		otc.SetCreateTime(*t)
	}
	return otc
}

// SetUpdateTime sets the "update_time" field.
func (otc *OrganizationTreeCreate) SetUpdateTime(t time.Time) *OrganizationTreeCreate {
	otc.mutation.SetUpdateTime(t)
	return otc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (otc *OrganizationTreeCreate) SetNillableUpdateTime(t *time.Time) *OrganizationTreeCreate {
	if t != nil {
		otc.SetUpdateTime(*t)
	}
	return otc
}

// SetName sets the "name" field.
func (otc *OrganizationTreeCreate) SetName(s string) *OrganizationTreeCreate {
	otc.mutation.SetName(s)
	return otc
}

// SetParentID sets the "parent_id" field.
func (otc *OrganizationTreeCreate) SetParentID(i int) *OrganizationTreeCreate {
	otc.mutation.SetParentID(i)
	return otc
}

// SetLeft sets the "left" field.
func (otc *OrganizationTreeCreate) SetLeft(i int) *OrganizationTreeCreate {
	otc.mutation.SetLeft(i)
	return otc
}

// SetRight sets the "right" field.
func (otc *OrganizationTreeCreate) SetRight(i int) *OrganizationTreeCreate {
	otc.mutation.SetRight(i)
	return otc
}

// SetLevel sets the "level" field.
func (otc *OrganizationTreeCreate) SetLevel(i int) *OrganizationTreeCreate {
	otc.mutation.SetLevel(i)
	return otc
}

// AddOrganizationPositionIDs adds the "organization_positions" edge to the OrganizationPosition entity by IDs.
func (otc *OrganizationTreeCreate) AddOrganizationPositionIDs(ids ...int) *OrganizationTreeCreate {
	otc.mutation.AddOrganizationPositionIDs(ids...)
	return otc
}

// AddOrganizationPositions adds the "organization_positions" edges to the OrganizationPosition entity.
func (otc *OrganizationTreeCreate) AddOrganizationPositions(o ...*OrganizationPosition) *OrganizationTreeCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otc.AddOrganizationPositionIDs(ids...)
}

// Mutation returns the OrganizationTreeMutation object of the builder.
func (otc *OrganizationTreeCreate) Mutation() *OrganizationTreeMutation {
	return otc.mutation
}

// Save creates the OrganizationTree in the database.
func (otc *OrganizationTreeCreate) Save(ctx context.Context) (*OrganizationTree, error) {
	var (
		err  error
		node *OrganizationTree
	)
	otc.defaults()
	if len(otc.hooks) == 0 {
		if err = otc.check(); err != nil {
			return nil, err
		}
		node, err = otc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationTreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = otc.check(); err != nil {
				return nil, err
			}
			otc.mutation = mutation
			if node, err = otc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(otc.hooks) - 1; i >= 0; i-- {
			if otc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = otc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, otc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrganizationTree)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrganizationTreeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (otc *OrganizationTreeCreate) SaveX(ctx context.Context) *OrganizationTree {
	v, err := otc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (otc *OrganizationTreeCreate) Exec(ctx context.Context) error {
	_, err := otc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otc *OrganizationTreeCreate) ExecX(ctx context.Context) {
	if err := otc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (otc *OrganizationTreeCreate) defaults() {
	if _, ok := otc.mutation.CreateTime(); !ok {
		v := organizationtree.DefaultCreateTime()
		otc.mutation.SetCreateTime(v)
	}
	if _, ok := otc.mutation.UpdateTime(); !ok {
		v := organizationtree.DefaultUpdateTime()
		otc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otc *OrganizationTreeCreate) check() error {
	if _, ok := otc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "OrganizationTree.create_time"`)}
	}
	if _, ok := otc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "OrganizationTree.update_time"`)}
	}
	if _, ok := otc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OrganizationTree.name"`)}
	}
	if _, ok := otc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`ent: missing required field "OrganizationTree.parent_id"`)}
	}
	if _, ok := otc.mutation.Left(); !ok {
		return &ValidationError{Name: "left", err: errors.New(`ent: missing required field "OrganizationTree.left"`)}
	}
	if _, ok := otc.mutation.Right(); !ok {
		return &ValidationError{Name: "right", err: errors.New(`ent: missing required field "OrganizationTree.right"`)}
	}
	if _, ok := otc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "OrganizationTree.level"`)}
	}
	return nil
}

func (otc *OrganizationTreeCreate) sqlSave(ctx context.Context) (*OrganizationTree, error) {
	_node, _spec := otc.createSpec()
	if err := sqlgraph.CreateNode(ctx, otc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (otc *OrganizationTreeCreate) createSpec() (*OrganizationTree, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationTree{config: otc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: organizationtree.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationtree.FieldID,
			},
		}
	)
	if value, ok := otc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationtree.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := otc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationtree.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := otc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationtree.FieldName,
		})
		_node.Name = value
	}
	if value, ok := otc.mutation.ParentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldParentID,
		})
		_node.ParentID = value
	}
	if value, ok := otc.mutation.Left(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLeft,
		})
		_node.Left = value
	}
	if value, ok := otc.mutation.Right(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldRight,
		})
		_node.Right = value
	}
	if value, ok := otc.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLevel,
		})
		_node.Level = value
	}
	if nodes := otc.mutation.OrganizationPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationtree.OrganizationPositionsTable,
			Columns: []string{organizationtree.OrganizationPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationposition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationTreeCreateBulk is the builder for creating many OrganizationTree entities in bulk.
type OrganizationTreeCreateBulk struct {
	config
	builders []*OrganizationTreeCreate
}

// Save creates the OrganizationTree entities in the database.
func (otcb *OrganizationTreeCreateBulk) Save(ctx context.Context) ([]*OrganizationTree, error) {
	specs := make([]*sqlgraph.CreateSpec, len(otcb.builders))
	nodes := make([]*OrganizationTree, len(otcb.builders))
	mutators := make([]Mutator, len(otcb.builders))
	for i := range otcb.builders {
		func(i int, root context.Context) {
			builder := otcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationTreeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, otcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, otcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, otcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (otcb *OrganizationTreeCreateBulk) SaveX(ctx context.Context) []*OrganizationTree {
	v, err := otcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (otcb *OrganizationTreeCreateBulk) Exec(ctx context.Context) error {
	_, err := otcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otcb *OrganizationTreeCreateBulk) ExecX(ctx context.Context) {
	if err := otcb.Exec(ctx); err != nil {
		panic(err)
	}
}
