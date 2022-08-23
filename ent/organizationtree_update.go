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
	"github.com/yjiong/iotdor/ent/organizationposition"
	"github.com/yjiong/iotdor/ent/organizationtree"
	"github.com/yjiong/iotdor/ent/predicate"
)

// OrganizationTreeUpdate is the builder for updating OrganizationTree entities.
type OrganizationTreeUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationTreeMutation
}

// Where appends a list predicates to the OrganizationTreeUpdate builder.
func (otu *OrganizationTreeUpdate) Where(ps ...predicate.OrganizationTree) *OrganizationTreeUpdate {
	otu.mutation.Where(ps...)
	return otu
}

// SetUpdateTime sets the "update_time" field.
func (otu *OrganizationTreeUpdate) SetUpdateTime(t time.Time) *OrganizationTreeUpdate {
	otu.mutation.SetUpdateTime(t)
	return otu
}

// SetName sets the "name" field.
func (otu *OrganizationTreeUpdate) SetName(s string) *OrganizationTreeUpdate {
	otu.mutation.SetName(s)
	return otu
}

// SetParentID sets the "parent_id" field.
func (otu *OrganizationTreeUpdate) SetParentID(i int) *OrganizationTreeUpdate {
	otu.mutation.ResetParentID()
	otu.mutation.SetParentID(i)
	return otu
}

// AddParentID adds i to the "parent_id" field.
func (otu *OrganizationTreeUpdate) AddParentID(i int) *OrganizationTreeUpdate {
	otu.mutation.AddParentID(i)
	return otu
}

// SetLeft sets the "left" field.
func (otu *OrganizationTreeUpdate) SetLeft(i int) *OrganizationTreeUpdate {
	otu.mutation.ResetLeft()
	otu.mutation.SetLeft(i)
	return otu
}

// AddLeft adds i to the "left" field.
func (otu *OrganizationTreeUpdate) AddLeft(i int) *OrganizationTreeUpdate {
	otu.mutation.AddLeft(i)
	return otu
}

// SetRight sets the "right" field.
func (otu *OrganizationTreeUpdate) SetRight(i int) *OrganizationTreeUpdate {
	otu.mutation.ResetRight()
	otu.mutation.SetRight(i)
	return otu
}

// AddRight adds i to the "right" field.
func (otu *OrganizationTreeUpdate) AddRight(i int) *OrganizationTreeUpdate {
	otu.mutation.AddRight(i)
	return otu
}

// SetLevel sets the "level" field.
func (otu *OrganizationTreeUpdate) SetLevel(i int) *OrganizationTreeUpdate {
	otu.mutation.ResetLevel()
	otu.mutation.SetLevel(i)
	return otu
}

// AddLevel adds i to the "level" field.
func (otu *OrganizationTreeUpdate) AddLevel(i int) *OrganizationTreeUpdate {
	otu.mutation.AddLevel(i)
	return otu
}

// SetOrganizationPositionsID sets the "organization_positions" edge to the OrganizationPosition entity by ID.
func (otu *OrganizationTreeUpdate) SetOrganizationPositionsID(id int) *OrganizationTreeUpdate {
	otu.mutation.SetOrganizationPositionsID(id)
	return otu
}

// SetOrganizationPositions sets the "organization_positions" edge to the OrganizationPosition entity.
func (otu *OrganizationTreeUpdate) SetOrganizationPositions(o *OrganizationPosition) *OrganizationTreeUpdate {
	return otu.SetOrganizationPositionsID(o.ID)
}

// Mutation returns the OrganizationTreeMutation object of the builder.
func (otu *OrganizationTreeUpdate) Mutation() *OrganizationTreeMutation {
	return otu.mutation
}

// ClearOrganizationPositions clears the "organization_positions" edge to the OrganizationPosition entity.
func (otu *OrganizationTreeUpdate) ClearOrganizationPositions() *OrganizationTreeUpdate {
	otu.mutation.ClearOrganizationPositions()
	return otu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (otu *OrganizationTreeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	otu.defaults()
	if len(otu.hooks) == 0 {
		if err = otu.check(); err != nil {
			return 0, err
		}
		affected, err = otu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationTreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = otu.check(); err != nil {
				return 0, err
			}
			otu.mutation = mutation
			affected, err = otu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(otu.hooks) - 1; i >= 0; i-- {
			if otu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = otu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, otu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (otu *OrganizationTreeUpdate) SaveX(ctx context.Context) int {
	affected, err := otu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (otu *OrganizationTreeUpdate) Exec(ctx context.Context) error {
	_, err := otu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otu *OrganizationTreeUpdate) ExecX(ctx context.Context) {
	if err := otu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (otu *OrganizationTreeUpdate) defaults() {
	if _, ok := otu.mutation.UpdateTime(); !ok {
		v := organizationtree.UpdateDefaultUpdateTime()
		otu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otu *OrganizationTreeUpdate) check() error {
	if _, ok := otu.mutation.OrganizationPositionsID(); otu.mutation.OrganizationPositionsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrganizationTree.organization_positions"`)
	}
	return nil
}

func (otu *OrganizationTreeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   organizationtree.Table,
			Columns: organizationtree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationtree.FieldID,
			},
		},
	}
	if ps := otu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationtree.FieldUpdateTime,
		})
	}
	if value, ok := otu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationtree.FieldName,
		})
	}
	if value, ok := otu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldParentID,
		})
	}
	if value, ok := otu.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldParentID,
		})
	}
	if value, ok := otu.mutation.Left(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLeft,
		})
	}
	if value, ok := otu.mutation.AddedLeft(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLeft,
		})
	}
	if value, ok := otu.mutation.Right(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldRight,
		})
	}
	if value, ok := otu.mutation.AddedRight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldRight,
		})
	}
	if value, ok := otu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLevel,
		})
	}
	if value, ok := otu.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLevel,
		})
	}
	if otu.mutation.OrganizationPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otu.mutation.OrganizationPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, otu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationtree.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrganizationTreeUpdateOne is the builder for updating a single OrganizationTree entity.
type OrganizationTreeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationTreeMutation
}

// SetUpdateTime sets the "update_time" field.
func (otuo *OrganizationTreeUpdateOne) SetUpdateTime(t time.Time) *OrganizationTreeUpdateOne {
	otuo.mutation.SetUpdateTime(t)
	return otuo
}

// SetName sets the "name" field.
func (otuo *OrganizationTreeUpdateOne) SetName(s string) *OrganizationTreeUpdateOne {
	otuo.mutation.SetName(s)
	return otuo
}

// SetParentID sets the "parent_id" field.
func (otuo *OrganizationTreeUpdateOne) SetParentID(i int) *OrganizationTreeUpdateOne {
	otuo.mutation.ResetParentID()
	otuo.mutation.SetParentID(i)
	return otuo
}

// AddParentID adds i to the "parent_id" field.
func (otuo *OrganizationTreeUpdateOne) AddParentID(i int) *OrganizationTreeUpdateOne {
	otuo.mutation.AddParentID(i)
	return otuo
}

// SetLeft sets the "left" field.
func (otuo *OrganizationTreeUpdateOne) SetLeft(i int) *OrganizationTreeUpdateOne {
	otuo.mutation.ResetLeft()
	otuo.mutation.SetLeft(i)
	return otuo
}

// AddLeft adds i to the "left" field.
func (otuo *OrganizationTreeUpdateOne) AddLeft(i int) *OrganizationTreeUpdateOne {
	otuo.mutation.AddLeft(i)
	return otuo
}

// SetRight sets the "right" field.
func (otuo *OrganizationTreeUpdateOne) SetRight(i int) *OrganizationTreeUpdateOne {
	otuo.mutation.ResetRight()
	otuo.mutation.SetRight(i)
	return otuo
}

// AddRight adds i to the "right" field.
func (otuo *OrganizationTreeUpdateOne) AddRight(i int) *OrganizationTreeUpdateOne {
	otuo.mutation.AddRight(i)
	return otuo
}

// SetLevel sets the "level" field.
func (otuo *OrganizationTreeUpdateOne) SetLevel(i int) *OrganizationTreeUpdateOne {
	otuo.mutation.ResetLevel()
	otuo.mutation.SetLevel(i)
	return otuo
}

// AddLevel adds i to the "level" field.
func (otuo *OrganizationTreeUpdateOne) AddLevel(i int) *OrganizationTreeUpdateOne {
	otuo.mutation.AddLevel(i)
	return otuo
}

// SetOrganizationPositionsID sets the "organization_positions" edge to the OrganizationPosition entity by ID.
func (otuo *OrganizationTreeUpdateOne) SetOrganizationPositionsID(id int) *OrganizationTreeUpdateOne {
	otuo.mutation.SetOrganizationPositionsID(id)
	return otuo
}

// SetOrganizationPositions sets the "organization_positions" edge to the OrganizationPosition entity.
func (otuo *OrganizationTreeUpdateOne) SetOrganizationPositions(o *OrganizationPosition) *OrganizationTreeUpdateOne {
	return otuo.SetOrganizationPositionsID(o.ID)
}

// Mutation returns the OrganizationTreeMutation object of the builder.
func (otuo *OrganizationTreeUpdateOne) Mutation() *OrganizationTreeMutation {
	return otuo.mutation
}

// ClearOrganizationPositions clears the "organization_positions" edge to the OrganizationPosition entity.
func (otuo *OrganizationTreeUpdateOne) ClearOrganizationPositions() *OrganizationTreeUpdateOne {
	otuo.mutation.ClearOrganizationPositions()
	return otuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (otuo *OrganizationTreeUpdateOne) Select(field string, fields ...string) *OrganizationTreeUpdateOne {
	otuo.fields = append([]string{field}, fields...)
	return otuo
}

// Save executes the query and returns the updated OrganizationTree entity.
func (otuo *OrganizationTreeUpdateOne) Save(ctx context.Context) (*OrganizationTree, error) {
	var (
		err  error
		node *OrganizationTree
	)
	otuo.defaults()
	if len(otuo.hooks) == 0 {
		if err = otuo.check(); err != nil {
			return nil, err
		}
		node, err = otuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationTreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = otuo.check(); err != nil {
				return nil, err
			}
			otuo.mutation = mutation
			node, err = otuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(otuo.hooks) - 1; i >= 0; i-- {
			if otuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = otuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, otuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (otuo *OrganizationTreeUpdateOne) SaveX(ctx context.Context) *OrganizationTree {
	node, err := otuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (otuo *OrganizationTreeUpdateOne) Exec(ctx context.Context) error {
	_, err := otuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otuo *OrganizationTreeUpdateOne) ExecX(ctx context.Context) {
	if err := otuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (otuo *OrganizationTreeUpdateOne) defaults() {
	if _, ok := otuo.mutation.UpdateTime(); !ok {
		v := organizationtree.UpdateDefaultUpdateTime()
		otuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otuo *OrganizationTreeUpdateOne) check() error {
	if _, ok := otuo.mutation.OrganizationPositionsID(); otuo.mutation.OrganizationPositionsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrganizationTree.organization_positions"`)
	}
	return nil
}

func (otuo *OrganizationTreeUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationTree, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   organizationtree.Table,
			Columns: organizationtree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationtree.FieldID,
			},
		},
	}
	id, ok := otuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrganizationTree.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := otuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationtree.FieldID)
		for _, f := range fields {
			if !organizationtree.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organizationtree.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := otuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationtree.FieldUpdateTime,
		})
	}
	if value, ok := otuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationtree.FieldName,
		})
	}
	if value, ok := otuo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldParentID,
		})
	}
	if value, ok := otuo.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldParentID,
		})
	}
	if value, ok := otuo.mutation.Left(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLeft,
		})
	}
	if value, ok := otuo.mutation.AddedLeft(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLeft,
		})
	}
	if value, ok := otuo.mutation.Right(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldRight,
		})
	}
	if value, ok := otuo.mutation.AddedRight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldRight,
		})
	}
	if value, ok := otuo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLevel,
		})
	}
	if value, ok := otuo.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: organizationtree.FieldLevel,
		})
	}
	if otuo.mutation.OrganizationPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otuo.mutation.OrganizationPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrganizationTree{config: otuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, otuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationtree.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
