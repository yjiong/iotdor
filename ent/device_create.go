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
	"github.com/yjiong/iotdor/ent/gateway"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	mutation *DeviceMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (dc *DeviceCreate) SetCreateTime(t time.Time) *DeviceCreate {
	dc.mutation.SetCreateTime(t)
	return dc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableCreateTime(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetCreateTime(*t)
	}
	return dc
}

// SetUpdateTime sets the "update_time" field.
func (dc *DeviceCreate) SetUpdateTime(t time.Time) *DeviceCreate {
	dc.mutation.SetUpdateTime(t)
	return dc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableUpdateTime(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetUpdateTime(*t)
	}
	return dc
}

// SetDevID sets the "devID" field.
func (dc *DeviceCreate) SetDevID(s string) *DeviceCreate {
	dc.mutation.SetDevID(s)
	return dc
}

// SetDevType sets the "devType" field.
func (dc *DeviceCreate) SetDevType(s string) *DeviceCreate {
	dc.mutation.SetDevType(s)
	return dc
}

// SetDevAddr sets the "devAddr" field.
func (dc *DeviceCreate) SetDevAddr(s string) *DeviceCreate {
	dc.mutation.SetDevAddr(s)
	return dc
}

// SetConn sets the "conn" field.
func (dc *DeviceCreate) SetConn(s string) *DeviceCreate {
	dc.mutation.SetConn(s)
	return dc
}

// SetName sets the "name" field.
func (dc *DeviceCreate) SetName(s string) *DeviceCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableName(s *string) *DeviceCreate {
	if s != nil {
		dc.SetName(*s)
	}
	return dc
}

// SetDeleteFlag sets the "DeleteFlag" field.
func (dc *DeviceCreate) SetDeleteFlag(b bool) *DeviceCreate {
	dc.mutation.SetDeleteFlag(b)
	return dc
}

// SetNillableDeleteFlag sets the "DeleteFlag" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableDeleteFlag(b *bool) *DeviceCreate {
	if b != nil {
		dc.SetDeleteFlag(*b)
	}
	return dc
}

// SetSummary sets the "summary" field.
func (dc *DeviceCreate) SetSummary(s string) *DeviceCreate {
	dc.mutation.SetSummary(s)
	return dc
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableSummary(s *string) *DeviceCreate {
	if s != nil {
		dc.SetSummary(*s)
	}
	return dc
}

// SetGatewayID sets the "gateway" edge to the Gateway entity by ID.
func (dc *DeviceCreate) SetGatewayID(id int) *DeviceCreate {
	dc.mutation.SetGatewayID(id)
	return dc
}

// SetNillableGatewayID sets the "gateway" edge to the Gateway entity by ID if the given value is not nil.
func (dc *DeviceCreate) SetNillableGatewayID(id *int) *DeviceCreate {
	if id != nil {
		dc = dc.SetGatewayID(*id)
	}
	return dc
}

// SetGateway sets the "gateway" edge to the Gateway entity.
func (dc *DeviceCreate) SetGateway(g *Gateway) *DeviceCreate {
	return dc.SetGatewayID(g.ID)
}

// Mutation returns the DeviceMutation object of the builder.
func (dc *DeviceCreate) Mutation() *DeviceMutation {
	return dc.mutation
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	var (
		err  error
		node *Device
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Device)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DeviceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeviceCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeviceCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeviceCreate) defaults() {
	if _, ok := dc.mutation.CreateTime(); !ok {
		v := device.DefaultCreateTime()
		dc.mutation.SetCreateTime(v)
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		v := device.DefaultUpdateTime()
		dc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeviceCreate) check() error {
	if _, ok := dc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Device.create_time"`)}
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Device.update_time"`)}
	}
	if _, ok := dc.mutation.DevID(); !ok {
		return &ValidationError{Name: "devID", err: errors.New(`ent: missing required field "Device.devID"`)}
	}
	if _, ok := dc.mutation.DevType(); !ok {
		return &ValidationError{Name: "devType", err: errors.New(`ent: missing required field "Device.devType"`)}
	}
	if _, ok := dc.mutation.DevAddr(); !ok {
		return &ValidationError{Name: "devAddr", err: errors.New(`ent: missing required field "Device.devAddr"`)}
	}
	if _, ok := dc.mutation.Conn(); !ok {
		return &ValidationError{Name: "conn", err: errors.New(`ent: missing required field "Device.conn"`)}
	}
	return nil
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DeviceCreate) createSpec() (*Device, *sqlgraph.CreateSpec) {
	var (
		_node = &Device{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: device.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := dc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := dc.mutation.DevID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevID,
		})
		_node.DevID = value
	}
	if value, ok := dc.mutation.DevType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevType,
		})
		_node.DevType = value
	}
	if value, ok := dc.mutation.DevAddr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevAddr,
		})
		_node.DevAddr = value
	}
	if value, ok := dc.mutation.Conn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldConn,
		})
		_node.Conn = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dc.mutation.DeleteFlag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: device.FieldDeleteFlag,
		})
		_node.DeleteFlag = value
	}
	if value, ok := dc.mutation.Summary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldSummary,
		})
		_node.Summary = value
	}
	if nodes := dc.mutation.GatewayIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.GatewayTable,
			Columns: []string{device.GatewayColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gateway.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.gateway_devices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeviceCreateBulk is the builder for creating many Device entities in bulk.
type DeviceCreateBulk struct {
	config
	builders []*DeviceCreate
}

// Save creates the Device entities in the database.
func (dcb *DeviceCreateBulk) Save(ctx context.Context) ([]*Device, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Device, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeviceCreateBulk) SaveX(ctx context.Context) []*Device {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeviceCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
