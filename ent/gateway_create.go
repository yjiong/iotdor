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
	"github.com/yjiong/iotdor/ent/group"
)

// GatewayCreate is the builder for creating a Gateway entity.
type GatewayCreate struct {
	config
	mutation *GatewayMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (gc *GatewayCreate) SetCreateTime(t time.Time) *GatewayCreate {
	gc.mutation.SetCreateTime(t)
	return gc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableCreateTime(t *time.Time) *GatewayCreate {
	if t != nil {
		gc.SetCreateTime(*t)
	}
	return gc
}

// SetUpdateTime sets the "update_time" field.
func (gc *GatewayCreate) SetUpdateTime(t time.Time) *GatewayCreate {
	gc.mutation.SetUpdateTime(t)
	return gc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableUpdateTime(t *time.Time) *GatewayCreate {
	if t != nil {
		gc.SetUpdateTime(*t)
	}
	return gc
}

// SetGwid sets the "gwid" field.
func (gc *GatewayCreate) SetGwid(s string) *GatewayCreate {
	gc.mutation.SetGwid(s)
	return gc
}

// SetSvrid sets the "svrid" field.
func (gc *GatewayCreate) SetSvrid(s string) *GatewayCreate {
	gc.mutation.SetSvrid(s)
	return gc
}

// SetBroker sets the "broker" field.
func (gc *GatewayCreate) SetBroker(s string) *GatewayCreate {
	gc.mutation.SetBroker(s)
	return gc
}

// SetInstallationLocation sets the "installation_location" field.
func (gc *GatewayCreate) SetInstallationLocation(s string) *GatewayCreate {
	gc.mutation.SetInstallationLocation(s)
	return gc
}

// SetNillableInstallationLocation sets the "installation_location" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableInstallationLocation(s *string) *GatewayCreate {
	if s != nil {
		gc.SetInstallationLocation(*s)
	}
	return gc
}

// SetStat sets the "stat" field.
func (gc *GatewayCreate) SetStat(s string) *GatewayCreate {
	gc.mutation.SetStat(s)
	return gc
}

// SetNillableStat sets the "stat" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableStat(s *string) *GatewayCreate {
	if s != nil {
		gc.SetStat(*s)
	}
	return gc
}

// SetDeleteFlag sets the "delete_flag" field.
func (gc *GatewayCreate) SetDeleteFlag(b bool) *GatewayCreate {
	gc.mutation.SetDeleteFlag(b)
	return gc
}

// SetNillableDeleteFlag sets the "delete_flag" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableDeleteFlag(b *bool) *GatewayCreate {
	if b != nil {
		gc.SetDeleteFlag(*b)
	}
	return gc
}

// SetUpInterval sets the "up_interval" field.
func (gc *GatewayCreate) SetUpInterval(i int) *GatewayCreate {
	gc.mutation.SetUpInterval(i)
	return gc
}

// SetNillableUpInterval sets the "up_interval" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableUpInterval(i *int) *GatewayCreate {
	if i != nil {
		gc.SetUpInterval(*i)
	}
	return gc
}

// SetVersion sets the "version" field.
func (gc *GatewayCreate) SetVersion(s string) *GatewayCreate {
	gc.mutation.SetVersion(s)
	return gc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableVersion(s *string) *GatewayCreate {
	if s != nil {
		gc.SetVersion(*s)
	}
	return gc
}

// SetSummary sets the "summary" field.
func (gc *GatewayCreate) SetSummary(s string) *GatewayCreate {
	gc.mutation.SetSummary(s)
	return gc
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableSummary(s *string) *GatewayCreate {
	if s != nil {
		gc.SetSummary(*s)
	}
	return gc
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (gc *GatewayCreate) AddDeviceIDs(ids ...int) *GatewayCreate {
	gc.mutation.AddDeviceIDs(ids...)
	return gc
}

// AddDevices adds the "devices" edges to the Device entity.
func (gc *GatewayCreate) AddDevices(d ...*Device) *GatewayCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gc.AddDeviceIDs(ids...)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gc *GatewayCreate) SetGroupID(id int) *GatewayCreate {
	gc.mutation.SetGroupID(id)
	return gc
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (gc *GatewayCreate) SetNillableGroupID(id *int) *GatewayCreate {
	if id != nil {
		gc = gc.SetGroupID(*id)
	}
	return gc
}

// SetGroup sets the "group" edge to the Group entity.
func (gc *GatewayCreate) SetGroup(g *Group) *GatewayCreate {
	return gc.SetGroupID(g.ID)
}

// Mutation returns the GatewayMutation object of the builder.
func (gc *GatewayCreate) Mutation() *GatewayMutation {
	return gc.mutation
}

// Save creates the Gateway in the database.
func (gc *GatewayCreate) Save(ctx context.Context) (*Gateway, error) {
	var (
		err  error
		node *Gateway
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GatewayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Gateway)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GatewayMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GatewayCreate) SaveX(ctx context.Context) *Gateway {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GatewayCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GatewayCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GatewayCreate) defaults() {
	if _, ok := gc.mutation.CreateTime(); !ok {
		v := gateway.DefaultCreateTime()
		gc.mutation.SetCreateTime(v)
	}
	if _, ok := gc.mutation.UpdateTime(); !ok {
		v := gateway.DefaultUpdateTime()
		gc.mutation.SetUpdateTime(v)
	}
	if _, ok := gc.mutation.UpInterval(); !ok {
		v := gateway.DefaultUpInterval
		gc.mutation.SetUpInterval(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GatewayCreate) check() error {
	if _, ok := gc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Gateway.create_time"`)}
	}
	if _, ok := gc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Gateway.update_time"`)}
	}
	if _, ok := gc.mutation.Gwid(); !ok {
		return &ValidationError{Name: "gwid", err: errors.New(`ent: missing required field "Gateway.gwid"`)}
	}
	if _, ok := gc.mutation.Svrid(); !ok {
		return &ValidationError{Name: "svrid", err: errors.New(`ent: missing required field "Gateway.svrid"`)}
	}
	if _, ok := gc.mutation.Broker(); !ok {
		return &ValidationError{Name: "broker", err: errors.New(`ent: missing required field "Gateway.broker"`)}
	}
	if _, ok := gc.mutation.UpInterval(); !ok {
		return &ValidationError{Name: "up_interval", err: errors.New(`ent: missing required field "Gateway.up_interval"`)}
	}
	return nil
}

func (gc *GatewayCreate) sqlSave(ctx context.Context) (*Gateway, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gc *GatewayCreate) createSpec() (*Gateway, *sqlgraph.CreateSpec) {
	var (
		_node = &Gateway{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gateway.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gateway.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gateway.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := gc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gateway.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := gc.mutation.Gwid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldGwid,
		})
		_node.Gwid = value
	}
	if value, ok := gc.mutation.Svrid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldSvrid,
		})
		_node.Svrid = value
	}
	if value, ok := gc.mutation.Broker(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldBroker,
		})
		_node.Broker = value
	}
	if value, ok := gc.mutation.InstallationLocation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldInstallationLocation,
		})
		_node.InstallationLocation = value
	}
	if value, ok := gc.mutation.Stat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldStat,
		})
		_node.Stat = value
	}
	if value, ok := gc.mutation.DeleteFlag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: gateway.FieldDeleteFlag,
		})
		_node.DeleteFlag = value
	}
	if value, ok := gc.mutation.UpInterval(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gateway.FieldUpInterval,
		})
		_node.UpInterval = value
	}
	if value, ok := gc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldVersion,
		})
		_node.Version = value
	}
	if value, ok := gc.mutation.Summary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldSummary,
		})
		_node.Summary = value
	}
	if nodes := gc.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gateway.DevicesTable,
			Columns: []string{gateway.DevicesColumn},
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
	if nodes := gc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gateway.GroupTable,
			Columns: []string{gateway.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_gateways = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GatewayCreateBulk is the builder for creating many Gateway entities in bulk.
type GatewayCreateBulk struct {
	config
	builders []*GatewayCreate
}

// Save creates the Gateway entities in the database.
func (gcb *GatewayCreateBulk) Save(ctx context.Context) ([]*Gateway, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Gateway, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GatewayMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GatewayCreateBulk) SaveX(ctx context.Context) []*Gateway {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GatewayCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GatewayCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
