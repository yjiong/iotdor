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
	"github.com/yjiong/iotdor/ent/device"
	"github.com/yjiong/iotdor/ent/gateway"
	"github.com/yjiong/iotdor/ent/group"
	"github.com/yjiong/iotdor/ent/predicate"
)

// GatewayUpdate is the builder for updating Gateway entities.
type GatewayUpdate struct {
	config
	hooks    []Hook
	mutation *GatewayMutation
}

// Where appends a list predicates to the GatewayUpdate builder.
func (gu *GatewayUpdate) Where(ps ...predicate.Gateway) *GatewayUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdateTime sets the "update_time" field.
func (gu *GatewayUpdate) SetUpdateTime(t time.Time) *GatewayUpdate {
	gu.mutation.SetUpdateTime(t)
	return gu
}

// SetGwid sets the "gwid" field.
func (gu *GatewayUpdate) SetGwid(s string) *GatewayUpdate {
	gu.mutation.SetGwid(s)
	return gu
}

// SetSvrid sets the "svrid" field.
func (gu *GatewayUpdate) SetSvrid(s string) *GatewayUpdate {
	gu.mutation.SetSvrid(s)
	return gu
}

// SetBroker sets the "broker" field.
func (gu *GatewayUpdate) SetBroker(s string) *GatewayUpdate {
	gu.mutation.SetBroker(s)
	return gu
}

// SetInstallationLocation sets the "installation_location" field.
func (gu *GatewayUpdate) SetInstallationLocation(s string) *GatewayUpdate {
	gu.mutation.SetInstallationLocation(s)
	return gu
}

// SetNillableInstallationLocation sets the "installation_location" field if the given value is not nil.
func (gu *GatewayUpdate) SetNillableInstallationLocation(s *string) *GatewayUpdate {
	if s != nil {
		gu.SetInstallationLocation(*s)
	}
	return gu
}

// ClearInstallationLocation clears the value of the "installation_location" field.
func (gu *GatewayUpdate) ClearInstallationLocation() *GatewayUpdate {
	gu.mutation.ClearInstallationLocation()
	return gu
}

// SetStat sets the "stat" field.
func (gu *GatewayUpdate) SetStat(s string) *GatewayUpdate {
	gu.mutation.SetStat(s)
	return gu
}

// SetNillableStat sets the "stat" field if the given value is not nil.
func (gu *GatewayUpdate) SetNillableStat(s *string) *GatewayUpdate {
	if s != nil {
		gu.SetStat(*s)
	}
	return gu
}

// ClearStat clears the value of the "stat" field.
func (gu *GatewayUpdate) ClearStat() *GatewayUpdate {
	gu.mutation.ClearStat()
	return gu
}

// SetDeleteFlag sets the "delete_flag" field.
func (gu *GatewayUpdate) SetDeleteFlag(b bool) *GatewayUpdate {
	gu.mutation.SetDeleteFlag(b)
	return gu
}

// SetNillableDeleteFlag sets the "delete_flag" field if the given value is not nil.
func (gu *GatewayUpdate) SetNillableDeleteFlag(b *bool) *GatewayUpdate {
	if b != nil {
		gu.SetDeleteFlag(*b)
	}
	return gu
}

// ClearDeleteFlag clears the value of the "delete_flag" field.
func (gu *GatewayUpdate) ClearDeleteFlag() *GatewayUpdate {
	gu.mutation.ClearDeleteFlag()
	return gu
}

// SetUpInterval sets the "up_interval" field.
func (gu *GatewayUpdate) SetUpInterval(i int) *GatewayUpdate {
	gu.mutation.ResetUpInterval()
	gu.mutation.SetUpInterval(i)
	return gu
}

// SetNillableUpInterval sets the "up_interval" field if the given value is not nil.
func (gu *GatewayUpdate) SetNillableUpInterval(i *int) *GatewayUpdate {
	if i != nil {
		gu.SetUpInterval(*i)
	}
	return gu
}

// AddUpInterval adds i to the "up_interval" field.
func (gu *GatewayUpdate) AddUpInterval(i int) *GatewayUpdate {
	gu.mutation.AddUpInterval(i)
	return gu
}

// SetVersion sets the "version" field.
func (gu *GatewayUpdate) SetVersion(s string) *GatewayUpdate {
	gu.mutation.SetVersion(s)
	return gu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (gu *GatewayUpdate) SetNillableVersion(s *string) *GatewayUpdate {
	if s != nil {
		gu.SetVersion(*s)
	}
	return gu
}

// ClearVersion clears the value of the "version" field.
func (gu *GatewayUpdate) ClearVersion() *GatewayUpdate {
	gu.mutation.ClearVersion()
	return gu
}

// SetSummary sets the "summary" field.
func (gu *GatewayUpdate) SetSummary(s string) *GatewayUpdate {
	gu.mutation.SetSummary(s)
	return gu
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (gu *GatewayUpdate) SetNillableSummary(s *string) *GatewayUpdate {
	if s != nil {
		gu.SetSummary(*s)
	}
	return gu
}

// ClearSummary clears the value of the "summary" field.
func (gu *GatewayUpdate) ClearSummary() *GatewayUpdate {
	gu.mutation.ClearSummary()
	return gu
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (gu *GatewayUpdate) AddDeviceIDs(ids ...int) *GatewayUpdate {
	gu.mutation.AddDeviceIDs(ids...)
	return gu
}

// AddDevices adds the "devices" edges to the Device entity.
func (gu *GatewayUpdate) AddDevices(d ...*Device) *GatewayUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gu.AddDeviceIDs(ids...)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gu *GatewayUpdate) SetGroupID(id int) *GatewayUpdate {
	gu.mutation.SetGroupID(id)
	return gu
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (gu *GatewayUpdate) SetNillableGroupID(id *int) *GatewayUpdate {
	if id != nil {
		gu = gu.SetGroupID(*id)
	}
	return gu
}

// SetGroup sets the "group" edge to the Group entity.
func (gu *GatewayUpdate) SetGroup(g *Group) *GatewayUpdate {
	return gu.SetGroupID(g.ID)
}

// Mutation returns the GatewayMutation object of the builder.
func (gu *GatewayUpdate) Mutation() *GatewayMutation {
	return gu.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (gu *GatewayUpdate) ClearDevices() *GatewayUpdate {
	gu.mutation.ClearDevices()
	return gu
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (gu *GatewayUpdate) RemoveDeviceIDs(ids ...int) *GatewayUpdate {
	gu.mutation.RemoveDeviceIDs(ids...)
	return gu
}

// RemoveDevices removes "devices" edges to Device entities.
func (gu *GatewayUpdate) RemoveDevices(d ...*Device) *GatewayUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gu.RemoveDeviceIDs(ids...)
}

// ClearGroup clears the "group" edge to the Group entity.
func (gu *GatewayUpdate) ClearGroup() *GatewayUpdate {
	gu.mutation.ClearGroup()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GatewayUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gu.defaults()
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GatewayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GatewayUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GatewayUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GatewayUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GatewayUpdate) defaults() {
	if _, ok := gu.mutation.UpdateTime(); !ok {
		v := gateway.UpdateDefaultUpdateTime()
		gu.mutation.SetUpdateTime(v)
	}
}

func (gu *GatewayUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gateway.Table,
			Columns: gateway.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gateway.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gateway.FieldUpdateTime,
		})
	}
	if value, ok := gu.mutation.Gwid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldGwid,
		})
	}
	if value, ok := gu.mutation.Svrid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldSvrid,
		})
	}
	if value, ok := gu.mutation.Broker(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldBroker,
		})
	}
	if value, ok := gu.mutation.InstallationLocation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldInstallationLocation,
		})
	}
	if gu.mutation.InstallationLocationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gateway.FieldInstallationLocation,
		})
	}
	if value, ok := gu.mutation.Stat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldStat,
		})
	}
	if gu.mutation.StatCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gateway.FieldStat,
		})
	}
	if value, ok := gu.mutation.DeleteFlag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: gateway.FieldDeleteFlag,
		})
	}
	if gu.mutation.DeleteFlagCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: gateway.FieldDeleteFlag,
		})
	}
	if value, ok := gu.mutation.UpInterval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gateway.FieldUpInterval,
		})
	}
	if value, ok := gu.mutation.AddedUpInterval(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gateway.FieldUpInterval,
		})
	}
	if value, ok := gu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldVersion,
		})
	}
	if gu.mutation.VersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gateway.FieldVersion,
		})
	}
	if value, ok := gu.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldSummary,
		})
	}
	if gu.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gateway.FieldSummary,
		})
	}
	if gu.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !gu.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.DevicesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.GroupCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.GroupIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gateway.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GatewayUpdateOne is the builder for updating a single Gateway entity.
type GatewayUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GatewayMutation
}

// SetUpdateTime sets the "update_time" field.
func (guo *GatewayUpdateOne) SetUpdateTime(t time.Time) *GatewayUpdateOne {
	guo.mutation.SetUpdateTime(t)
	return guo
}

// SetGwid sets the "gwid" field.
func (guo *GatewayUpdateOne) SetGwid(s string) *GatewayUpdateOne {
	guo.mutation.SetGwid(s)
	return guo
}

// SetSvrid sets the "svrid" field.
func (guo *GatewayUpdateOne) SetSvrid(s string) *GatewayUpdateOne {
	guo.mutation.SetSvrid(s)
	return guo
}

// SetBroker sets the "broker" field.
func (guo *GatewayUpdateOne) SetBroker(s string) *GatewayUpdateOne {
	guo.mutation.SetBroker(s)
	return guo
}

// SetInstallationLocation sets the "installation_location" field.
func (guo *GatewayUpdateOne) SetInstallationLocation(s string) *GatewayUpdateOne {
	guo.mutation.SetInstallationLocation(s)
	return guo
}

// SetNillableInstallationLocation sets the "installation_location" field if the given value is not nil.
func (guo *GatewayUpdateOne) SetNillableInstallationLocation(s *string) *GatewayUpdateOne {
	if s != nil {
		guo.SetInstallationLocation(*s)
	}
	return guo
}

// ClearInstallationLocation clears the value of the "installation_location" field.
func (guo *GatewayUpdateOne) ClearInstallationLocation() *GatewayUpdateOne {
	guo.mutation.ClearInstallationLocation()
	return guo
}

// SetStat sets the "stat" field.
func (guo *GatewayUpdateOne) SetStat(s string) *GatewayUpdateOne {
	guo.mutation.SetStat(s)
	return guo
}

// SetNillableStat sets the "stat" field if the given value is not nil.
func (guo *GatewayUpdateOne) SetNillableStat(s *string) *GatewayUpdateOne {
	if s != nil {
		guo.SetStat(*s)
	}
	return guo
}

// ClearStat clears the value of the "stat" field.
func (guo *GatewayUpdateOne) ClearStat() *GatewayUpdateOne {
	guo.mutation.ClearStat()
	return guo
}

// SetDeleteFlag sets the "delete_flag" field.
func (guo *GatewayUpdateOne) SetDeleteFlag(b bool) *GatewayUpdateOne {
	guo.mutation.SetDeleteFlag(b)
	return guo
}

// SetNillableDeleteFlag sets the "delete_flag" field if the given value is not nil.
func (guo *GatewayUpdateOne) SetNillableDeleteFlag(b *bool) *GatewayUpdateOne {
	if b != nil {
		guo.SetDeleteFlag(*b)
	}
	return guo
}

// ClearDeleteFlag clears the value of the "delete_flag" field.
func (guo *GatewayUpdateOne) ClearDeleteFlag() *GatewayUpdateOne {
	guo.mutation.ClearDeleteFlag()
	return guo
}

// SetUpInterval sets the "up_interval" field.
func (guo *GatewayUpdateOne) SetUpInterval(i int) *GatewayUpdateOne {
	guo.mutation.ResetUpInterval()
	guo.mutation.SetUpInterval(i)
	return guo
}

// SetNillableUpInterval sets the "up_interval" field if the given value is not nil.
func (guo *GatewayUpdateOne) SetNillableUpInterval(i *int) *GatewayUpdateOne {
	if i != nil {
		guo.SetUpInterval(*i)
	}
	return guo
}

// AddUpInterval adds i to the "up_interval" field.
func (guo *GatewayUpdateOne) AddUpInterval(i int) *GatewayUpdateOne {
	guo.mutation.AddUpInterval(i)
	return guo
}

// SetVersion sets the "version" field.
func (guo *GatewayUpdateOne) SetVersion(s string) *GatewayUpdateOne {
	guo.mutation.SetVersion(s)
	return guo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (guo *GatewayUpdateOne) SetNillableVersion(s *string) *GatewayUpdateOne {
	if s != nil {
		guo.SetVersion(*s)
	}
	return guo
}

// ClearVersion clears the value of the "version" field.
func (guo *GatewayUpdateOne) ClearVersion() *GatewayUpdateOne {
	guo.mutation.ClearVersion()
	return guo
}

// SetSummary sets the "summary" field.
func (guo *GatewayUpdateOne) SetSummary(s string) *GatewayUpdateOne {
	guo.mutation.SetSummary(s)
	return guo
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (guo *GatewayUpdateOne) SetNillableSummary(s *string) *GatewayUpdateOne {
	if s != nil {
		guo.SetSummary(*s)
	}
	return guo
}

// ClearSummary clears the value of the "summary" field.
func (guo *GatewayUpdateOne) ClearSummary() *GatewayUpdateOne {
	guo.mutation.ClearSummary()
	return guo
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (guo *GatewayUpdateOne) AddDeviceIDs(ids ...int) *GatewayUpdateOne {
	guo.mutation.AddDeviceIDs(ids...)
	return guo
}

// AddDevices adds the "devices" edges to the Device entity.
func (guo *GatewayUpdateOne) AddDevices(d ...*Device) *GatewayUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return guo.AddDeviceIDs(ids...)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (guo *GatewayUpdateOne) SetGroupID(id int) *GatewayUpdateOne {
	guo.mutation.SetGroupID(id)
	return guo
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (guo *GatewayUpdateOne) SetNillableGroupID(id *int) *GatewayUpdateOne {
	if id != nil {
		guo = guo.SetGroupID(*id)
	}
	return guo
}

// SetGroup sets the "group" edge to the Group entity.
func (guo *GatewayUpdateOne) SetGroup(g *Group) *GatewayUpdateOne {
	return guo.SetGroupID(g.ID)
}

// Mutation returns the GatewayMutation object of the builder.
func (guo *GatewayUpdateOne) Mutation() *GatewayMutation {
	return guo.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (guo *GatewayUpdateOne) ClearDevices() *GatewayUpdateOne {
	guo.mutation.ClearDevices()
	return guo
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (guo *GatewayUpdateOne) RemoveDeviceIDs(ids ...int) *GatewayUpdateOne {
	guo.mutation.RemoveDeviceIDs(ids...)
	return guo
}

// RemoveDevices removes "devices" edges to Device entities.
func (guo *GatewayUpdateOne) RemoveDevices(d ...*Device) *GatewayUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return guo.RemoveDeviceIDs(ids...)
}

// ClearGroup clears the "group" edge to the Group entity.
func (guo *GatewayUpdateOne) ClearGroup() *GatewayUpdateOne {
	guo.mutation.ClearGroup()
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GatewayUpdateOne) Select(field string, fields ...string) *GatewayUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Gateway entity.
func (guo *GatewayUpdateOne) Save(ctx context.Context) (*Gateway, error) {
	var (
		err  error
		node *Gateway
	)
	guo.defaults()
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GatewayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (guo *GatewayUpdateOne) SaveX(ctx context.Context) *Gateway {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GatewayUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GatewayUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GatewayUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdateTime(); !ok {
		v := gateway.UpdateDefaultUpdateTime()
		guo.mutation.SetUpdateTime(v)
	}
}

func (guo *GatewayUpdateOne) sqlSave(ctx context.Context) (_node *Gateway, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gateway.Table,
			Columns: gateway.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gateway.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Gateway.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gateway.FieldID)
		for _, f := range fields {
			if !gateway.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gateway.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gateway.FieldUpdateTime,
		})
	}
	if value, ok := guo.mutation.Gwid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldGwid,
		})
	}
	if value, ok := guo.mutation.Svrid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldSvrid,
		})
	}
	if value, ok := guo.mutation.Broker(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldBroker,
		})
	}
	if value, ok := guo.mutation.InstallationLocation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldInstallationLocation,
		})
	}
	if guo.mutation.InstallationLocationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gateway.FieldInstallationLocation,
		})
	}
	if value, ok := guo.mutation.Stat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldStat,
		})
	}
	if guo.mutation.StatCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gateway.FieldStat,
		})
	}
	if value, ok := guo.mutation.DeleteFlag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: gateway.FieldDeleteFlag,
		})
	}
	if guo.mutation.DeleteFlagCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: gateway.FieldDeleteFlag,
		})
	}
	if value, ok := guo.mutation.UpInterval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gateway.FieldUpInterval,
		})
	}
	if value, ok := guo.mutation.AddedUpInterval(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gateway.FieldUpInterval,
		})
	}
	if value, ok := guo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldVersion,
		})
	}
	if guo.mutation.VersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gateway.FieldVersion,
		})
	}
	if value, ok := guo.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gateway.FieldSummary,
		})
	}
	if guo.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gateway.FieldSummary,
		})
	}
	if guo.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !guo.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.DevicesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.GroupCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.GroupIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Gateway{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gateway.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
