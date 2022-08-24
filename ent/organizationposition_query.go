// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yjiong/iotdor/ent/device"
	"github.com/yjiong/iotdor/ent/organizationposition"
	"github.com/yjiong/iotdor/ent/organizationtree"
	"github.com/yjiong/iotdor/ent/predicate"
	"github.com/yjiong/iotdor/ent/user"
)

// OrganizationPositionQuery is the builder for querying OrganizationPosition entities.
type OrganizationPositionQuery struct {
	config
	limit                *int
	offset               *int
	unique               *bool
	order                []OrderFunc
	fields               []string
	predicates           []predicate.OrganizationPosition
	withDevices          *DeviceQuery
	withPersonCharges    *UserQuery
	withOrganizationTree *OrganizationTreeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationPositionQuery builder.
func (opq *OrganizationPositionQuery) Where(ps ...predicate.OrganizationPosition) *OrganizationPositionQuery {
	opq.predicates = append(opq.predicates, ps...)
	return opq
}

// Limit adds a limit step to the query.
func (opq *OrganizationPositionQuery) Limit(limit int) *OrganizationPositionQuery {
	opq.limit = &limit
	return opq
}

// Offset adds an offset step to the query.
func (opq *OrganizationPositionQuery) Offset(offset int) *OrganizationPositionQuery {
	opq.offset = &offset
	return opq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (opq *OrganizationPositionQuery) Unique(unique bool) *OrganizationPositionQuery {
	opq.unique = &unique
	return opq
}

// Order adds an order step to the query.
func (opq *OrganizationPositionQuery) Order(o ...OrderFunc) *OrganizationPositionQuery {
	opq.order = append(opq.order, o...)
	return opq
}

// QueryDevices chains the current query on the "devices" edge.
func (opq *OrganizationPositionQuery) QueryDevices() *DeviceQuery {
	query := &DeviceQuery{config: opq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationposition.Table, organizationposition.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organizationposition.DevicesTable, organizationposition.DevicesColumn),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPersonCharges chains the current query on the "person_charges" edge.
func (opq *OrganizationPositionQuery) QueryPersonCharges() *UserQuery {
	query := &UserQuery{config: opq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationposition.Table, organizationposition.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, organizationposition.PersonChargesTable, organizationposition.PersonChargesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganizationTree chains the current query on the "organization_tree" edge.
func (opq *OrganizationPositionQuery) QueryOrganizationTree() *OrganizationTreeQuery {
	query := &OrganizationTreeQuery{config: opq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationposition.Table, organizationposition.FieldID, selector),
			sqlgraph.To(organizationtree.Table, organizationtree.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, organizationposition.OrganizationTreeTable, organizationposition.OrganizationTreeColumn),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrganizationPosition entity from the query.
// Returns a *NotFoundError when no OrganizationPosition was found.
func (opq *OrganizationPositionQuery) First(ctx context.Context) (*OrganizationPosition, error) {
	nodes, err := opq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationposition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (opq *OrganizationPositionQuery) FirstX(ctx context.Context) *OrganizationPosition {
	node, err := opq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationPosition ID from the query.
// Returns a *NotFoundError when no OrganizationPosition ID was found.
func (opq *OrganizationPositionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = opq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationposition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (opq *OrganizationPositionQuery) FirstIDX(ctx context.Context) int {
	id, err := opq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationPosition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationPosition entity is found.
// Returns a *NotFoundError when no OrganizationPosition entities are found.
func (opq *OrganizationPositionQuery) Only(ctx context.Context) (*OrganizationPosition, error) {
	nodes, err := opq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationposition.Label}
	default:
		return nil, &NotSingularError{organizationposition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (opq *OrganizationPositionQuery) OnlyX(ctx context.Context) *OrganizationPosition {
	node, err := opq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationPosition ID in the query.
// Returns a *NotSingularError when more than one OrganizationPosition ID is found.
// Returns a *NotFoundError when no entities are found.
func (opq *OrganizationPositionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = opq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationposition.Label}
	default:
		err = &NotSingularError{organizationposition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (opq *OrganizationPositionQuery) OnlyIDX(ctx context.Context) int {
	id, err := opq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationPositions.
func (opq *OrganizationPositionQuery) All(ctx context.Context) ([]*OrganizationPosition, error) {
	if err := opq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return opq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (opq *OrganizationPositionQuery) AllX(ctx context.Context) []*OrganizationPosition {
	nodes, err := opq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationPosition IDs.
func (opq *OrganizationPositionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := opq.Select(organizationposition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (opq *OrganizationPositionQuery) IDsX(ctx context.Context) []int {
	ids, err := opq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (opq *OrganizationPositionQuery) Count(ctx context.Context) (int, error) {
	if err := opq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return opq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (opq *OrganizationPositionQuery) CountX(ctx context.Context) int {
	count, err := opq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (opq *OrganizationPositionQuery) Exist(ctx context.Context) (bool, error) {
	if err := opq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return opq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (opq *OrganizationPositionQuery) ExistX(ctx context.Context) bool {
	exist, err := opq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationPositionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (opq *OrganizationPositionQuery) Clone() *OrganizationPositionQuery {
	if opq == nil {
		return nil
	}
	return &OrganizationPositionQuery{
		config:               opq.config,
		limit:                opq.limit,
		offset:               opq.offset,
		order:                append([]OrderFunc{}, opq.order...),
		predicates:           append([]predicate.OrganizationPosition{}, opq.predicates...),
		withDevices:          opq.withDevices.Clone(),
		withPersonCharges:    opq.withPersonCharges.Clone(),
		withOrganizationTree: opq.withOrganizationTree.Clone(),
		// clone intermediate query.
		sql:    opq.sql.Clone(),
		path:   opq.path,
		unique: opq.unique,
	}
}

// WithDevices tells the query-builder to eager-load the nodes that are connected to
// the "devices" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrganizationPositionQuery) WithDevices(opts ...func(*DeviceQuery)) *OrganizationPositionQuery {
	query := &DeviceQuery{config: opq.config}
	for _, opt := range opts {
		opt(query)
	}
	opq.withDevices = query
	return opq
}

// WithPersonCharges tells the query-builder to eager-load the nodes that are connected to
// the "person_charges" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrganizationPositionQuery) WithPersonCharges(opts ...func(*UserQuery)) *OrganizationPositionQuery {
	query := &UserQuery{config: opq.config}
	for _, opt := range opts {
		opt(query)
	}
	opq.withPersonCharges = query
	return opq
}

// WithOrganizationTree tells the query-builder to eager-load the nodes that are connected to
// the "organization_tree" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrganizationPositionQuery) WithOrganizationTree(opts ...func(*OrganizationTreeQuery)) *OrganizationPositionQuery {
	query := &OrganizationTreeQuery{config: opq.config}
	for _, opt := range opts {
		opt(query)
	}
	opq.withOrganizationTree = query
	return opq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrganizationPosition.Query().
//		GroupBy(organizationposition.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (opq *OrganizationPositionQuery) GroupBy(field string, fields ...string) *OrganizationPositionGroupBy {
	grbuild := &OrganizationPositionGroupBy{config: opq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return opq.sqlQuery(ctx), nil
	}
	grbuild.label = organizationposition.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.OrganizationPosition.Query().
//		Select(organizationposition.FieldCreateTime).
//		Scan(ctx, &v)
//
func (opq *OrganizationPositionQuery) Select(fields ...string) *OrganizationPositionSelect {
	opq.fields = append(opq.fields, fields...)
	selbuild := &OrganizationPositionSelect{OrganizationPositionQuery: opq}
	selbuild.label = organizationposition.Label
	selbuild.flds, selbuild.scan = &opq.fields, selbuild.Scan
	return selbuild
}

func (opq *OrganizationPositionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range opq.fields {
		if !organizationposition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if opq.path != nil {
		prev, err := opq.path(ctx)
		if err != nil {
			return err
		}
		opq.sql = prev
	}
	return nil
}

func (opq *OrganizationPositionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationPosition, error) {
	var (
		nodes       = []*OrganizationPosition{}
		_spec       = opq.querySpec()
		loadedTypes = [3]bool{
			opq.withDevices != nil,
			opq.withPersonCharges != nil,
			opq.withOrganizationTree != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OrganizationPosition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OrganizationPosition{config: opq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, opq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := opq.withDevices; query != nil {
		if err := opq.loadDevices(ctx, query, nodes,
			func(n *OrganizationPosition) { n.Edges.Devices = []*Device{} },
			func(n *OrganizationPosition, e *Device) { n.Edges.Devices = append(n.Edges.Devices, e) }); err != nil {
			return nil, err
		}
	}
	if query := opq.withPersonCharges; query != nil {
		if err := opq.loadPersonCharges(ctx, query, nodes,
			func(n *OrganizationPosition) { n.Edges.PersonCharges = []*User{} },
			func(n *OrganizationPosition, e *User) { n.Edges.PersonCharges = append(n.Edges.PersonCharges, e) }); err != nil {
			return nil, err
		}
	}
	if query := opq.withOrganizationTree; query != nil {
		if err := opq.loadOrganizationTree(ctx, query, nodes,
			func(n *OrganizationPosition) { n.Edges.OrganizationTree = []*OrganizationTree{} },
			func(n *OrganizationPosition, e *OrganizationTree) {
				n.Edges.OrganizationTree = append(n.Edges.OrganizationTree, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (opq *OrganizationPositionQuery) loadDevices(ctx context.Context, query *DeviceQuery, nodes []*OrganizationPosition, init func(*OrganizationPosition), assign func(*OrganizationPosition, *Device)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*OrganizationPosition)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Device(func(s *sql.Selector) {
		s.Where(sql.InValues(organizationposition.DevicesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.organization_position_devices
		if fk == nil {
			return fmt.Errorf(`foreign-key "organization_position_devices" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_position_devices" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (opq *OrganizationPositionQuery) loadPersonCharges(ctx context.Context, query *UserQuery, nodes []*OrganizationPosition, init func(*OrganizationPosition), assign func(*OrganizationPosition, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*OrganizationPosition)
	nids := make(map[int]map[*OrganizationPosition]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(organizationposition.PersonChargesTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(organizationposition.PersonChargesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(organizationposition.PersonChargesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(organizationposition.PersonChargesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*OrganizationPosition]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "person_charges" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (opq *OrganizationPositionQuery) loadOrganizationTree(ctx context.Context, query *OrganizationTreeQuery, nodes []*OrganizationPosition, init func(*OrganizationPosition), assign func(*OrganizationPosition, *OrganizationTree)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*OrganizationPosition)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.InValues(organizationposition.OrganizationTreeColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.organization_tree_organization_positions
		if fk == nil {
			return fmt.Errorf(`foreign-key "organization_tree_organization_positions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_tree_organization_positions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (opq *OrganizationPositionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := opq.querySpec()
	_spec.Node.Columns = opq.fields
	if len(opq.fields) > 0 {
		_spec.Unique = opq.unique != nil && *opq.unique
	}
	return sqlgraph.CountNodes(ctx, opq.driver, _spec)
}

func (opq *OrganizationPositionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := opq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (opq *OrganizationPositionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   organizationposition.Table,
			Columns: organizationposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationposition.FieldID,
			},
		},
		From:   opq.sql,
		Unique: true,
	}
	if unique := opq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := opq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationposition.FieldID)
		for i := range fields {
			if fields[i] != organizationposition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := opq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := opq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := opq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := opq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (opq *OrganizationPositionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(opq.driver.Dialect())
	t1 := builder.Table(organizationposition.Table)
	columns := opq.fields
	if len(columns) == 0 {
		columns = organizationposition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if opq.sql != nil {
		selector = opq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if opq.unique != nil && *opq.unique {
		selector.Distinct()
	}
	for _, p := range opq.predicates {
		p(selector)
	}
	for _, p := range opq.order {
		p(selector)
	}
	if offset := opq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := opq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrganizationPositionGroupBy is the group-by builder for OrganizationPosition entities.
type OrganizationPositionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (opgb *OrganizationPositionGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationPositionGroupBy {
	opgb.fns = append(opgb.fns, fns...)
	return opgb
}

// Scan applies the group-by query and scans the result into the given value.
func (opgb *OrganizationPositionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := opgb.path(ctx)
	if err != nil {
		return err
	}
	opgb.sql = query
	return opgb.sqlScan(ctx, v)
}

func (opgb *OrganizationPositionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range opgb.fields {
		if !organizationposition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := opgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := opgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (opgb *OrganizationPositionGroupBy) sqlQuery() *sql.Selector {
	selector := opgb.sql.Select()
	aggregation := make([]string, 0, len(opgb.fns))
	for _, fn := range opgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(opgb.fields)+len(opgb.fns))
		for _, f := range opgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(opgb.fields...)...)
}

// OrganizationPositionSelect is the builder for selecting fields of OrganizationPosition entities.
type OrganizationPositionSelect struct {
	*OrganizationPositionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ops *OrganizationPositionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ops.prepareQuery(ctx); err != nil {
		return err
	}
	ops.sql = ops.OrganizationPositionQuery.sqlQuery(ctx)
	return ops.sqlScan(ctx, v)
}

func (ops *OrganizationPositionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ops.sql.Query()
	if err := ops.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}