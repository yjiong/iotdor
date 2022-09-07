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
	"github.com/yjiong/iotdor/ent/organizationposition"
	"github.com/yjiong/iotdor/ent/organizationtree"
	"github.com/yjiong/iotdor/ent/predicate"
	"github.com/yjiong/iotdor/ent/user"
)

// OrganizationPositionUpdate is the builder for updating OrganizationPosition entities.
type OrganizationPositionUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationPositionMutation
}

// Where appends a list predicates to the OrganizationPositionUpdate builder.
func (opu *OrganizationPositionUpdate) Where(ps ...predicate.OrganizationPosition) *OrganizationPositionUpdate {
	opu.mutation.Where(ps...)
	return opu
}

// SetUpdateTime sets the "update_time" field.
func (opu *OrganizationPositionUpdate) SetUpdateTime(t time.Time) *OrganizationPositionUpdate {
	opu.mutation.SetUpdateTime(t)
	return opu
}

// SetPositionID sets the "position_id" field.
func (opu *OrganizationPositionUpdate) SetPositionID(s string) *OrganizationPositionUpdate {
	opu.mutation.SetPositionID(s)
	return opu
}

// SetAddress sets the "address" field.
func (opu *OrganizationPositionUpdate) SetAddress(s string) *OrganizationPositionUpdate {
	opu.mutation.SetAddress(s)
	return opu
}

// SetFloor sets the "floor" field.
func (opu *OrganizationPositionUpdate) SetFloor(s string) *OrganizationPositionUpdate {
	opu.mutation.SetFloor(s)
	return opu
}

// SetNillableFloor sets the "floor" field if the given value is not nil.
func (opu *OrganizationPositionUpdate) SetNillableFloor(s *string) *OrganizationPositionUpdate {
	if s != nil {
		opu.SetFloor(*s)
	}
	return opu
}

// ClearFloor clears the value of the "floor" field.
func (opu *OrganizationPositionUpdate) ClearFloor() *OrganizationPositionUpdate {
	opu.mutation.ClearFloor()
	return opu
}

// SetUnitNo sets the "unit_no" field.
func (opu *OrganizationPositionUpdate) SetUnitNo(s string) *OrganizationPositionUpdate {
	opu.mutation.SetUnitNo(s)
	return opu
}

// SetNillableUnitNo sets the "unit_no" field if the given value is not nil.
func (opu *OrganizationPositionUpdate) SetNillableUnitNo(s *string) *OrganizationPositionUpdate {
	if s != nil {
		opu.SetUnitNo(*s)
	}
	return opu
}

// ClearUnitNo clears the value of the "unit_no" field.
func (opu *OrganizationPositionUpdate) ClearUnitNo() *OrganizationPositionUpdate {
	opu.mutation.ClearUnitNo()
	return opu
}

// SetLongitudeAndLatitude sets the "longitude_and_latitude" field.
func (opu *OrganizationPositionUpdate) SetLongitudeAndLatitude(s string) *OrganizationPositionUpdate {
	opu.mutation.SetLongitudeAndLatitude(s)
	return opu
}

// SetSummary sets the "summary" field.
func (opu *OrganizationPositionUpdate) SetSummary(s string) *OrganizationPositionUpdate {
	opu.mutation.SetSummary(s)
	return opu
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (opu *OrganizationPositionUpdate) SetNillableSummary(s *string) *OrganizationPositionUpdate {
	if s != nil {
		opu.SetSummary(*s)
	}
	return opu
}

// ClearSummary clears the value of the "summary" field.
func (opu *OrganizationPositionUpdate) ClearSummary() *OrganizationPositionUpdate {
	opu.mutation.ClearSummary()
	return opu
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (opu *OrganizationPositionUpdate) AddDeviceIDs(ids ...int) *OrganizationPositionUpdate {
	opu.mutation.AddDeviceIDs(ids...)
	return opu
}

// AddDevices adds the "devices" edges to the Device entity.
func (opu *OrganizationPositionUpdate) AddDevices(d ...*Device) *OrganizationPositionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return opu.AddDeviceIDs(ids...)
}

// AddPersonChargeIDs adds the "person_charges" edge to the User entity by IDs.
func (opu *OrganizationPositionUpdate) AddPersonChargeIDs(ids ...int) *OrganizationPositionUpdate {
	opu.mutation.AddPersonChargeIDs(ids...)
	return opu
}

// AddPersonCharges adds the "person_charges" edges to the User entity.
func (opu *OrganizationPositionUpdate) AddPersonCharges(u ...*User) *OrganizationPositionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return opu.AddPersonChargeIDs(ids...)
}

// SetOrganizationTreeID sets the "organization_tree" edge to the OrganizationTree entity by ID.
func (opu *OrganizationPositionUpdate) SetOrganizationTreeID(id int) *OrganizationPositionUpdate {
	opu.mutation.SetOrganizationTreeID(id)
	return opu
}

// SetNillableOrganizationTreeID sets the "organization_tree" edge to the OrganizationTree entity by ID if the given value is not nil.
func (opu *OrganizationPositionUpdate) SetNillableOrganizationTreeID(id *int) *OrganizationPositionUpdate {
	if id != nil {
		opu = opu.SetOrganizationTreeID(*id)
	}
	return opu
}

// SetOrganizationTree sets the "organization_tree" edge to the OrganizationTree entity.
func (opu *OrganizationPositionUpdate) SetOrganizationTree(o *OrganizationTree) *OrganizationPositionUpdate {
	return opu.SetOrganizationTreeID(o.ID)
}

// Mutation returns the OrganizationPositionMutation object of the builder.
func (opu *OrganizationPositionUpdate) Mutation() *OrganizationPositionMutation {
	return opu.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (opu *OrganizationPositionUpdate) ClearDevices() *OrganizationPositionUpdate {
	opu.mutation.ClearDevices()
	return opu
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (opu *OrganizationPositionUpdate) RemoveDeviceIDs(ids ...int) *OrganizationPositionUpdate {
	opu.mutation.RemoveDeviceIDs(ids...)
	return opu
}

// RemoveDevices removes "devices" edges to Device entities.
func (opu *OrganizationPositionUpdate) RemoveDevices(d ...*Device) *OrganizationPositionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return opu.RemoveDeviceIDs(ids...)
}

// ClearPersonCharges clears all "person_charges" edges to the User entity.
func (opu *OrganizationPositionUpdate) ClearPersonCharges() *OrganizationPositionUpdate {
	opu.mutation.ClearPersonCharges()
	return opu
}

// RemovePersonChargeIDs removes the "person_charges" edge to User entities by IDs.
func (opu *OrganizationPositionUpdate) RemovePersonChargeIDs(ids ...int) *OrganizationPositionUpdate {
	opu.mutation.RemovePersonChargeIDs(ids...)
	return opu
}

// RemovePersonCharges removes "person_charges" edges to User entities.
func (opu *OrganizationPositionUpdate) RemovePersonCharges(u ...*User) *OrganizationPositionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return opu.RemovePersonChargeIDs(ids...)
}

// ClearOrganizationTree clears the "organization_tree" edge to the OrganizationTree entity.
func (opu *OrganizationPositionUpdate) ClearOrganizationTree() *OrganizationPositionUpdate {
	opu.mutation.ClearOrganizationTree()
	return opu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opu *OrganizationPositionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	opu.defaults()
	if len(opu.hooks) == 0 {
		affected, err = opu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			opu.mutation = mutation
			affected, err = opu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(opu.hooks) - 1; i >= 0; i-- {
			if opu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = opu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, opu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (opu *OrganizationPositionUpdate) SaveX(ctx context.Context) int {
	affected, err := opu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opu *OrganizationPositionUpdate) Exec(ctx context.Context) error {
	_, err := opu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opu *OrganizationPositionUpdate) ExecX(ctx context.Context) {
	if err := opu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opu *OrganizationPositionUpdate) defaults() {
	if _, ok := opu.mutation.UpdateTime(); !ok {
		v := organizationposition.UpdateDefaultUpdateTime()
		opu.mutation.SetUpdateTime(v)
	}
}

func (opu *OrganizationPositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   organizationposition.Table,
			Columns: organizationposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationposition.FieldID,
			},
		},
	}
	if ps := opu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationposition.FieldUpdateTime,
		})
	}
	if value, ok := opu.mutation.PositionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldPositionID,
		})
	}
	if value, ok := opu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldAddress,
		})
	}
	if value, ok := opu.mutation.Floor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldFloor,
		})
	}
	if opu.mutation.FloorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: organizationposition.FieldFloor,
		})
	}
	if value, ok := opu.mutation.UnitNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldUnitNo,
		})
	}
	if opu.mutation.UnitNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: organizationposition.FieldUnitNo,
		})
	}
	if value, ok := opu.mutation.LongitudeAndLatitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldLongitudeAndLatitude,
		})
	}
	if value, ok := opu.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldSummary,
		})
	}
	if opu.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: organizationposition.FieldSummary,
		})
	}
	if opu.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !opu.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.DevicesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if opu.mutation.PersonChargesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.RemovedPersonChargesIDs(); len(nodes) > 0 && !opu.mutation.PersonChargesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.PersonChargesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if opu.mutation.OrganizationTreeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.OrganizationTreeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, opu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationposition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrganizationPositionUpdateOne is the builder for updating a single OrganizationPosition entity.
type OrganizationPositionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationPositionMutation
}

// SetUpdateTime sets the "update_time" field.
func (opuo *OrganizationPositionUpdateOne) SetUpdateTime(t time.Time) *OrganizationPositionUpdateOne {
	opuo.mutation.SetUpdateTime(t)
	return opuo
}

// SetPositionID sets the "position_id" field.
func (opuo *OrganizationPositionUpdateOne) SetPositionID(s string) *OrganizationPositionUpdateOne {
	opuo.mutation.SetPositionID(s)
	return opuo
}

// SetAddress sets the "address" field.
func (opuo *OrganizationPositionUpdateOne) SetAddress(s string) *OrganizationPositionUpdateOne {
	opuo.mutation.SetAddress(s)
	return opuo
}

// SetFloor sets the "floor" field.
func (opuo *OrganizationPositionUpdateOne) SetFloor(s string) *OrganizationPositionUpdateOne {
	opuo.mutation.SetFloor(s)
	return opuo
}

// SetNillableFloor sets the "floor" field if the given value is not nil.
func (opuo *OrganizationPositionUpdateOne) SetNillableFloor(s *string) *OrganizationPositionUpdateOne {
	if s != nil {
		opuo.SetFloor(*s)
	}
	return opuo
}

// ClearFloor clears the value of the "floor" field.
func (opuo *OrganizationPositionUpdateOne) ClearFloor() *OrganizationPositionUpdateOne {
	opuo.mutation.ClearFloor()
	return opuo
}

// SetUnitNo sets the "unit_no" field.
func (opuo *OrganizationPositionUpdateOne) SetUnitNo(s string) *OrganizationPositionUpdateOne {
	opuo.mutation.SetUnitNo(s)
	return opuo
}

// SetNillableUnitNo sets the "unit_no" field if the given value is not nil.
func (opuo *OrganizationPositionUpdateOne) SetNillableUnitNo(s *string) *OrganizationPositionUpdateOne {
	if s != nil {
		opuo.SetUnitNo(*s)
	}
	return opuo
}

// ClearUnitNo clears the value of the "unit_no" field.
func (opuo *OrganizationPositionUpdateOne) ClearUnitNo() *OrganizationPositionUpdateOne {
	opuo.mutation.ClearUnitNo()
	return opuo
}

// SetLongitudeAndLatitude sets the "longitude_and_latitude" field.
func (opuo *OrganizationPositionUpdateOne) SetLongitudeAndLatitude(s string) *OrganizationPositionUpdateOne {
	opuo.mutation.SetLongitudeAndLatitude(s)
	return opuo
}

// SetSummary sets the "summary" field.
func (opuo *OrganizationPositionUpdateOne) SetSummary(s string) *OrganizationPositionUpdateOne {
	opuo.mutation.SetSummary(s)
	return opuo
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (opuo *OrganizationPositionUpdateOne) SetNillableSummary(s *string) *OrganizationPositionUpdateOne {
	if s != nil {
		opuo.SetSummary(*s)
	}
	return opuo
}

// ClearSummary clears the value of the "summary" field.
func (opuo *OrganizationPositionUpdateOne) ClearSummary() *OrganizationPositionUpdateOne {
	opuo.mutation.ClearSummary()
	return opuo
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (opuo *OrganizationPositionUpdateOne) AddDeviceIDs(ids ...int) *OrganizationPositionUpdateOne {
	opuo.mutation.AddDeviceIDs(ids...)
	return opuo
}

// AddDevices adds the "devices" edges to the Device entity.
func (opuo *OrganizationPositionUpdateOne) AddDevices(d ...*Device) *OrganizationPositionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return opuo.AddDeviceIDs(ids...)
}

// AddPersonChargeIDs adds the "person_charges" edge to the User entity by IDs.
func (opuo *OrganizationPositionUpdateOne) AddPersonChargeIDs(ids ...int) *OrganizationPositionUpdateOne {
	opuo.mutation.AddPersonChargeIDs(ids...)
	return opuo
}

// AddPersonCharges adds the "person_charges" edges to the User entity.
func (opuo *OrganizationPositionUpdateOne) AddPersonCharges(u ...*User) *OrganizationPositionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return opuo.AddPersonChargeIDs(ids...)
}

// SetOrganizationTreeID sets the "organization_tree" edge to the OrganizationTree entity by ID.
func (opuo *OrganizationPositionUpdateOne) SetOrganizationTreeID(id int) *OrganizationPositionUpdateOne {
	opuo.mutation.SetOrganizationTreeID(id)
	return opuo
}

// SetNillableOrganizationTreeID sets the "organization_tree" edge to the OrganizationTree entity by ID if the given value is not nil.
func (opuo *OrganizationPositionUpdateOne) SetNillableOrganizationTreeID(id *int) *OrganizationPositionUpdateOne {
	if id != nil {
		opuo = opuo.SetOrganizationTreeID(*id)
	}
	return opuo
}

// SetOrganizationTree sets the "organization_tree" edge to the OrganizationTree entity.
func (opuo *OrganizationPositionUpdateOne) SetOrganizationTree(o *OrganizationTree) *OrganizationPositionUpdateOne {
	return opuo.SetOrganizationTreeID(o.ID)
}

// Mutation returns the OrganizationPositionMutation object of the builder.
func (opuo *OrganizationPositionUpdateOne) Mutation() *OrganizationPositionMutation {
	return opuo.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (opuo *OrganizationPositionUpdateOne) ClearDevices() *OrganizationPositionUpdateOne {
	opuo.mutation.ClearDevices()
	return opuo
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (opuo *OrganizationPositionUpdateOne) RemoveDeviceIDs(ids ...int) *OrganizationPositionUpdateOne {
	opuo.mutation.RemoveDeviceIDs(ids...)
	return opuo
}

// RemoveDevices removes "devices" edges to Device entities.
func (opuo *OrganizationPositionUpdateOne) RemoveDevices(d ...*Device) *OrganizationPositionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return opuo.RemoveDeviceIDs(ids...)
}

// ClearPersonCharges clears all "person_charges" edges to the User entity.
func (opuo *OrganizationPositionUpdateOne) ClearPersonCharges() *OrganizationPositionUpdateOne {
	opuo.mutation.ClearPersonCharges()
	return opuo
}

// RemovePersonChargeIDs removes the "person_charges" edge to User entities by IDs.
func (opuo *OrganizationPositionUpdateOne) RemovePersonChargeIDs(ids ...int) *OrganizationPositionUpdateOne {
	opuo.mutation.RemovePersonChargeIDs(ids...)
	return opuo
}

// RemovePersonCharges removes "person_charges" edges to User entities.
func (opuo *OrganizationPositionUpdateOne) RemovePersonCharges(u ...*User) *OrganizationPositionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return opuo.RemovePersonChargeIDs(ids...)
}

// ClearOrganizationTree clears the "organization_tree" edge to the OrganizationTree entity.
func (opuo *OrganizationPositionUpdateOne) ClearOrganizationTree() *OrganizationPositionUpdateOne {
	opuo.mutation.ClearOrganizationTree()
	return opuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opuo *OrganizationPositionUpdateOne) Select(field string, fields ...string) *OrganizationPositionUpdateOne {
	opuo.fields = append([]string{field}, fields...)
	return opuo
}

// Save executes the query and returns the updated OrganizationPosition entity.
func (opuo *OrganizationPositionUpdateOne) Save(ctx context.Context) (*OrganizationPosition, error) {
	var (
		err  error
		node *OrganizationPosition
	)
	opuo.defaults()
	if len(opuo.hooks) == 0 {
		node, err = opuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			opuo.mutation = mutation
			node, err = opuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(opuo.hooks) - 1; i >= 0; i-- {
			if opuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = opuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, opuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (opuo *OrganizationPositionUpdateOne) SaveX(ctx context.Context) *OrganizationPosition {
	node, err := opuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opuo *OrganizationPositionUpdateOne) Exec(ctx context.Context) error {
	_, err := opuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opuo *OrganizationPositionUpdateOne) ExecX(ctx context.Context) {
	if err := opuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opuo *OrganizationPositionUpdateOne) defaults() {
	if _, ok := opuo.mutation.UpdateTime(); !ok {
		v := organizationposition.UpdateDefaultUpdateTime()
		opuo.mutation.SetUpdateTime(v)
	}
}

func (opuo *OrganizationPositionUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationPosition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   organizationposition.Table,
			Columns: organizationposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationposition.FieldID,
			},
		},
	}
	id, ok := opuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrganizationPosition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationposition.FieldID)
		for _, f := range fields {
			if !organizationposition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organizationposition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationposition.FieldUpdateTime,
		})
	}
	if value, ok := opuo.mutation.PositionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldPositionID,
		})
	}
	if value, ok := opuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldAddress,
		})
	}
	if value, ok := opuo.mutation.Floor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldFloor,
		})
	}
	if opuo.mutation.FloorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: organizationposition.FieldFloor,
		})
	}
	if value, ok := opuo.mutation.UnitNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldUnitNo,
		})
	}
	if opuo.mutation.UnitNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: organizationposition.FieldUnitNo,
		})
	}
	if value, ok := opuo.mutation.LongitudeAndLatitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldLongitudeAndLatitude,
		})
	}
	if value, ok := opuo.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationposition.FieldSummary,
		})
	}
	if opuo.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: organizationposition.FieldSummary,
		})
	}
	if opuo.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !opuo.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.DevicesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if opuo.mutation.PersonChargesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.RemovedPersonChargesIDs(); len(nodes) > 0 && !opuo.mutation.PersonChargesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.PersonChargesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if opuo.mutation.OrganizationTreeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.OrganizationTreeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrganizationPosition{config: opuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationposition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
