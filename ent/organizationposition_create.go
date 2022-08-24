// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yjiong/iotdor/ent/device"
	"github.com/yjiong/iotdor/ent/organizationposition"
	"github.com/yjiong/iotdor/ent/organizationtree"
	"github.com/yjiong/iotdor/ent/user"
)

// OrganizationPositionCreate is the builder for creating a OrganizationPosition entity.
type OrganizationPositionCreate struct {
	config
	mutation *OrganizationPositionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (opc *OrganizationPositionCreate) SetCreateTime(t time.Time) *OrganizationPositionCreate {
	opc.mutation.SetCreateTime(t)
	return opc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (opc *OrganizationPositionCreate) SetNillableCreateTime(t *time.Time) *OrganizationPositionCreate {
	if t != nil {
		opc.SetCreateTime(*t)
	}
	return opc
}

// SetUpdateTime sets the "update_time" field.
func (opc *OrganizationPositionCreate) SetUpdateTime(t time.Time) *OrganizationPositionCreate {
	opc.mutation.SetUpdateTime(t)
	return opc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (opc *OrganizationPositionCreate) SetNillableUpdateTime(t *time.Time) *OrganizationPositionCreate {
	if t != nil {
		opc.SetUpdateTime(*t)
	}
	return opc
}

// SetPositionID sets the "position_id" field.
func (opc *OrganizationPositionCreate) SetPositionID(s string) *OrganizationPositionCreate {
	opc.mutation.SetPositionID(s)
	return opc
}

// SetAddress sets the "address" field.
func (opc *OrganizationPositionCreate) SetAddress(s string) *OrganizationPositionCreate {
	opc.mutation.SetAddress(s)
	return opc
}

// SetFloor sets the "floor" field.
func (opc *OrganizationPositionCreate) SetFloor(s string) *OrganizationPositionCreate {
	opc.mutation.SetFloor(s)
	return opc
}

// SetNillableFloor sets the "floor" field if the given value is not nil.
func (opc *OrganizationPositionCreate) SetNillableFloor(s *string) *OrganizationPositionCreate {
	if s != nil {
		opc.SetFloor(*s)
	}
	return opc
}

// SetUnitNo sets the "unit_no" field.
func (opc *OrganizationPositionCreate) SetUnitNo(s string) *OrganizationPositionCreate {
	opc.mutation.SetUnitNo(s)
	return opc
}

// SetNillableUnitNo sets the "unit_no" field if the given value is not nil.
func (opc *OrganizationPositionCreate) SetNillableUnitNo(s *string) *OrganizationPositionCreate {
	if s != nil {
		opc.SetUnitNo(*s)
	}
	return opc
}

// SetLongitudeAndLatitude sets the "longitude_and_latitude" field.
func (opc *OrganizationPositionCreate) SetLongitudeAndLatitude(s string) *OrganizationPositionCreate {
	opc.mutation.SetLongitudeAndLatitude(s)
	return opc
}

// SetSummary sets the "summary" field.
func (opc *OrganizationPositionCreate) SetSummary(s string) *OrganizationPositionCreate {
	opc.mutation.SetSummary(s)
	return opc
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (opc *OrganizationPositionCreate) SetNillableSummary(s *string) *OrganizationPositionCreate {
	if s != nil {
		opc.SetSummary(*s)
	}
	return opc
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (opc *OrganizationPositionCreate) AddDeviceIDs(ids ...int) *OrganizationPositionCreate {
	opc.mutation.AddDeviceIDs(ids...)
	return opc
}

// AddDevices adds the "devices" edges to the Device entity.
func (opc *OrganizationPositionCreate) AddDevices(d ...*Device) *OrganizationPositionCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return opc.AddDeviceIDs(ids...)
}

// AddPersonChargeIDs adds the "person_charges" edge to the User entity by IDs.
func (opc *OrganizationPositionCreate) AddPersonChargeIDs(ids ...int) *OrganizationPositionCreate {
	opc.mutation.AddPersonChargeIDs(ids...)
	return opc
}

// AddPersonCharges adds the "person_charges" edges to the User entity.
func (opc *OrganizationPositionCreate) AddPersonCharges(u ...*User) *OrganizationPositionCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return opc.AddPersonChargeIDs(ids...)
}

// AddOrganizationTreeIDs adds the "organization_tree" edge to the OrganizationTree entity by IDs.
func (opc *OrganizationPositionCreate) AddOrganizationTreeIDs(ids ...int) *OrganizationPositionCreate {
	opc.mutation.AddOrganizationTreeIDs(ids...)
	return opc
}

// AddOrganizationTree adds the "organization_tree" edges to the OrganizationTree entity.
func (opc *OrganizationPositionCreate) AddOrganizationTree(o ...*OrganizationTree) *OrganizationPositionCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return opc.AddOrganizationTreeIDs(ids...)
}

// Mutation returns the OrganizationPositionMutation object of the builder.
func (opc *OrganizationPositionCreate) Mutation() *OrganizationPositionMutation {
	return opc.mutation
}

// Save creates the OrganizationPosition in the database.
func (opc *OrganizationPositionCreate) Save(ctx context.Context) (*OrganizationPosition, error) {
	var (
		err  error
		node *OrganizationPosition
	)
	opc.defaults()
	if len(opc.hooks) == 0 {
		if err = opc.check(); err != nil {
			return nil, err
		}
		node, err = opc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = opc.check(); err != nil {
				return nil, err
			}
			opc.mutation = mutation
			if node, err = opc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(opc.hooks) - 1; i >= 0; i-- {
			if opc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = opc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, opc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrganizationPosition)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrganizationPositionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (opc *OrganizationPositionCreate) SaveX(ctx context.Context) *OrganizationPosition {
	v, err := opc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opc *OrganizationPositionCreate) Exec(ctx context.Context) error {
	_, err := opc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opc *OrganizationPositionCreate) ExecX(ctx context.Context) {
	if err := opc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opc *OrganizationPositionCreate) defaults() {
	if _, ok := opc.mutation.CreateTime(); !ok {
		v := organizationposition.DefaultCreateTime()
		opc.mutation.SetCreateTime(v)
	}
	if _, ok := opc.mutation.UpdateTime(); !ok {
		v := organizationposition.DefaultUpdateTime()
		opc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (opc *OrganizationPositionCreate) check() error {
	if _, ok := opc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "OrganizationPosition.create_time"`)}
	}
	if _, ok := opc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "OrganizationPosition.update_time"`)}
	}
	if _, ok := opc.mutation.PositionID(); !ok {
		return &ValidationError{Name: "position_id", err: errors.New(`ent: missing required field "OrganizationPosition.position_id"`)}
	}
	if _, ok := opc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "OrganizationPosition.address"`)}
	}
	if _, ok := opc.mutation.LongitudeAndLatitude(); !ok {
		return &ValidationError{Name: "longitude_and_latitude", err: errors.New(`ent: missing required field "OrganizationPosition.longitude_and_latitude"`)}
	}
	return nil
}

func (opc *OrganizationPositionCreate) sqlSave(ctx context.Context) (*OrganizationPosition, error) {
	_node, _spec := opc.createSpec()
	if err := sqlgraph.CreateNode(ctx, opc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (opc *OrganizationPositionCreate) createSpec() (*OrganizationPosition, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationPosition{config: opc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: organizationposition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationposition.FieldID,
			},
		}
	)
	if value, ok := opc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationposition.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := opc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationposition.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := opc.mutation.PositionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldPositionID,
		})
		_node.PositionID = value
	}
	if value, ok := opc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := opc.mutation.Floor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldFloor,
		})
		_node.Floor = value
	}
	if value, ok := opc.mutation.UnitNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldUnitNo,
		})
		_node.UnitNo = value
	}
	if value, ok := opc.mutation.LongitudeAndLatitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldLongitudeAndLatitude,
		})
		_node.LongitudeAndLatitude = value
	}
	if value, ok := opc.mutation.Summary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldSummary,
		})
		_node.Summary = value
	}
	if nodes := opc.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationposition.DevicesTable,
			Columns: []string{organizationposition.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := opc.mutation.PersonChargesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organizationposition.PersonChargesTable,
			Columns: organizationposition.PersonChargesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := opc.mutation.OrganizationTreeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organizationposition.OrganizationTreeTable,
			Columns: []string{organizationposition.OrganizationTreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationtree.FieldID,
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

// OrganizationPositionCreateBulk is the builder for creating many OrganizationPosition entities in bulk.
type OrganizationPositionCreateBulk struct {
	config
	builders []*OrganizationPositionCreate
}

// Save creates the OrganizationPosition entities in the database.
func (opcb *OrganizationPositionCreateBulk) Save(ctx context.Context) ([]*OrganizationPosition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(opcb.builders))
	nodes := make([]*OrganizationPosition, len(opcb.builders))
	mutators := make([]Mutator, len(opcb.builders))
	for i := range opcb.builders {
		func(i int, root context.Context) {
			builder := opcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationPositionMutation)
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
					_, err = mutators[i+1].Mutate(root, opcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, opcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, opcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (opcb *OrganizationPositionCreateBulk) SaveX(ctx context.Context) []*OrganizationPosition {
	v, err := opcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opcb *OrganizationPositionCreateBulk) Exec(ctx context.Context) error {
	_, err := opcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opcb *OrganizationPositionCreateBulk) ExecX(ctx context.Context) {
	if err := opcb.Exec(ctx); err != nil {
		panic(err)
	}
}
