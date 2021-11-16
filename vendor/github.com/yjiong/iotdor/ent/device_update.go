// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"iotdor/ent/device"
	"iotdor/ent/gateway"
	"iotdor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetDevID sets the "devID" field.
func (du *DeviceUpdate) SetDevID(s string) *DeviceUpdate {
	du.mutation.SetDevID(s)
	return du
}

// SetDevType sets the "devType" field.
func (du *DeviceUpdate) SetDevType(s string) *DeviceUpdate {
	du.mutation.SetDevType(s)
	return du
}

// SetDevAddr sets the "devAddr" field.
func (du *DeviceUpdate) SetDevAddr(s string) *DeviceUpdate {
	du.mutation.SetDevAddr(s)
	return du
}

// SetConn sets the "conn" field.
func (du *DeviceUpdate) SetConn(s string) *DeviceUpdate {
	du.mutation.SetConn(s)
	return du
}

// SetName sets the "name" field.
func (du *DeviceUpdate) SetName(s string) *DeviceUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableName(s *string) *DeviceUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// ClearName clears the value of the "name" field.
func (du *DeviceUpdate) ClearName() *DeviceUpdate {
	du.mutation.ClearName()
	return du
}

// SetIdDelete sets the "idDelete" field.
func (du *DeviceUpdate) SetIdDelete(b bool) *DeviceUpdate {
	du.mutation.SetIdDelete(b)
	return du
}

// SetGatewayID sets the "gateway" edge to the Gateway entity by ID.
func (du *DeviceUpdate) SetGatewayID(id int) *DeviceUpdate {
	du.mutation.SetGatewayID(id)
	return du
}

// SetNillableGatewayID sets the "gateway" edge to the Gateway entity by ID if the given value is not nil.
func (du *DeviceUpdate) SetNillableGatewayID(id *int) *DeviceUpdate {
	if id != nil {
		du = du.SetGatewayID(*id)
	}
	return du
}

// SetGateway sets the "gateway" edge to the Gateway entity.
func (du *DeviceUpdate) SetGateway(g *Gateway) *DeviceUpdate {
	return du.SetGatewayID(g.ID)
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// ClearGateway clears the "gateway" edge to the Gateway entity.
func (du *DeviceUpdate) ClearGateway() *DeviceUpdate {
	du.mutation.ClearGateway()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DeviceUpdate) defaults() {
	if _, ok := du.mutation.UpdateTime(); !ok {
		v := device.UpdateDefaultUpdateTime()
		du.mutation.SetUpdateTime(v)
	}
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldUpdateTime,
		})
	}
	if value, ok := du.mutation.DevID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevID,
		})
	}
	if value, ok := du.mutation.DevType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevType,
		})
	}
	if value, ok := du.mutation.DevAddr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevAddr,
		})
	}
	if value, ok := du.mutation.Conn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldConn,
		})
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldName,
		})
	}
	if du.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: device.FieldName,
		})
	}
	if value, ok := du.mutation.IdDelete(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: device.FieldIdDelete,
		})
	}
	if du.mutation.GatewayCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.GatewayIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceMutation
}

// SetDevID sets the "devID" field.
func (duo *DeviceUpdateOne) SetDevID(s string) *DeviceUpdateOne {
	duo.mutation.SetDevID(s)
	return duo
}

// SetDevType sets the "devType" field.
func (duo *DeviceUpdateOne) SetDevType(s string) *DeviceUpdateOne {
	duo.mutation.SetDevType(s)
	return duo
}

// SetDevAddr sets the "devAddr" field.
func (duo *DeviceUpdateOne) SetDevAddr(s string) *DeviceUpdateOne {
	duo.mutation.SetDevAddr(s)
	return duo
}

// SetConn sets the "conn" field.
func (duo *DeviceUpdateOne) SetConn(s string) *DeviceUpdateOne {
	duo.mutation.SetConn(s)
	return duo
}

// SetName sets the "name" field.
func (duo *DeviceUpdateOne) SetName(s string) *DeviceUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableName(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// ClearName clears the value of the "name" field.
func (duo *DeviceUpdateOne) ClearName() *DeviceUpdateOne {
	duo.mutation.ClearName()
	return duo
}

// SetIdDelete sets the "idDelete" field.
func (duo *DeviceUpdateOne) SetIdDelete(b bool) *DeviceUpdateOne {
	duo.mutation.SetIdDelete(b)
	return duo
}

// SetGatewayID sets the "gateway" edge to the Gateway entity by ID.
func (duo *DeviceUpdateOne) SetGatewayID(id int) *DeviceUpdateOne {
	duo.mutation.SetGatewayID(id)
	return duo
}

// SetNillableGatewayID sets the "gateway" edge to the Gateway entity by ID if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableGatewayID(id *int) *DeviceUpdateOne {
	if id != nil {
		duo = duo.SetGatewayID(*id)
	}
	return duo
}

// SetGateway sets the "gateway" edge to the Gateway entity.
func (duo *DeviceUpdateOne) SetGateway(g *Gateway) *DeviceUpdateOne {
	return duo.SetGatewayID(g.ID)
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// ClearGateway clears the "gateway" edge to the Gateway entity.
func (duo *DeviceUpdateOne) ClearGateway() *DeviceUpdateOne {
	duo.mutation.ClearGateway()
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeviceUpdateOne) Select(field string, fields ...string) *DeviceUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Device entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	var (
		err  error
		node *Device
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DeviceUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdateTime(); !ok {
		v := device.UpdateDefaultUpdateTime()
		duo.mutation.SetUpdateTime(v)
	}
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (_node *Device, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Device.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for _, f := range fields {
			if !device.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldUpdateTime,
		})
	}
	if value, ok := duo.mutation.DevID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevID,
		})
	}
	if value, ok := duo.mutation.DevType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevType,
		})
	}
	if value, ok := duo.mutation.DevAddr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDevAddr,
		})
	}
	if value, ok := duo.mutation.Conn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldConn,
		})
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldName,
		})
	}
	if duo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: device.FieldName,
		})
	}
	if value, ok := duo.mutation.IdDelete(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: device.FieldIdDelete,
		})
	}
	if duo.mutation.GatewayCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.GatewayIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Device{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
