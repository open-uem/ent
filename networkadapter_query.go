// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/networkadapter"
	"github.com/open-uem/ent/predicate"
)

// NetworkAdapterQuery is the builder for querying NetworkAdapter entities.
type NetworkAdapterQuery struct {
	config
	ctx        *QueryContext
	order      []networkadapter.OrderOption
	inters     []Interceptor
	predicates []predicate.NetworkAdapter
	withOwner  *AgentQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NetworkAdapterQuery builder.
func (naq *NetworkAdapterQuery) Where(ps ...predicate.NetworkAdapter) *NetworkAdapterQuery {
	naq.predicates = append(naq.predicates, ps...)
	return naq
}

// Limit the number of records to be returned by this query.
func (naq *NetworkAdapterQuery) Limit(limit int) *NetworkAdapterQuery {
	naq.ctx.Limit = &limit
	return naq
}

// Offset to start from.
func (naq *NetworkAdapterQuery) Offset(offset int) *NetworkAdapterQuery {
	naq.ctx.Offset = &offset
	return naq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (naq *NetworkAdapterQuery) Unique(unique bool) *NetworkAdapterQuery {
	naq.ctx.Unique = &unique
	return naq
}

// Order specifies how the records should be ordered.
func (naq *NetworkAdapterQuery) Order(o ...networkadapter.OrderOption) *NetworkAdapterQuery {
	naq.order = append(naq.order, o...)
	return naq
}

// QueryOwner chains the current query on the "owner" edge.
func (naq *NetworkAdapterQuery) QueryOwner() *AgentQuery {
	query := (&AgentClient{config: naq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := naq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := naq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(networkadapter.Table, networkadapter.FieldID, selector),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, networkadapter.OwnerTable, networkadapter.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(naq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NetworkAdapter entity from the query.
// Returns a *NotFoundError when no NetworkAdapter was found.
func (naq *NetworkAdapterQuery) First(ctx context.Context) (*NetworkAdapter, error) {
	nodes, err := naq.Limit(1).All(setContextOp(ctx, naq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{networkadapter.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (naq *NetworkAdapterQuery) FirstX(ctx context.Context) *NetworkAdapter {
	node, err := naq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NetworkAdapter ID from the query.
// Returns a *NotFoundError when no NetworkAdapter ID was found.
func (naq *NetworkAdapterQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = naq.Limit(1).IDs(setContextOp(ctx, naq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{networkadapter.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (naq *NetworkAdapterQuery) FirstIDX(ctx context.Context) int {
	id, err := naq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NetworkAdapter entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NetworkAdapter entity is found.
// Returns a *NotFoundError when no NetworkAdapter entities are found.
func (naq *NetworkAdapterQuery) Only(ctx context.Context) (*NetworkAdapter, error) {
	nodes, err := naq.Limit(2).All(setContextOp(ctx, naq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{networkadapter.Label}
	default:
		return nil, &NotSingularError{networkadapter.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (naq *NetworkAdapterQuery) OnlyX(ctx context.Context) *NetworkAdapter {
	node, err := naq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NetworkAdapter ID in the query.
// Returns a *NotSingularError when more than one NetworkAdapter ID is found.
// Returns a *NotFoundError when no entities are found.
func (naq *NetworkAdapterQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = naq.Limit(2).IDs(setContextOp(ctx, naq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{networkadapter.Label}
	default:
		err = &NotSingularError{networkadapter.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (naq *NetworkAdapterQuery) OnlyIDX(ctx context.Context) int {
	id, err := naq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NetworkAdapters.
func (naq *NetworkAdapterQuery) All(ctx context.Context) ([]*NetworkAdapter, error) {
	ctx = setContextOp(ctx, naq.ctx, ent.OpQueryAll)
	if err := naq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NetworkAdapter, *NetworkAdapterQuery]()
	return withInterceptors[[]*NetworkAdapter](ctx, naq, qr, naq.inters)
}

// AllX is like All, but panics if an error occurs.
func (naq *NetworkAdapterQuery) AllX(ctx context.Context) []*NetworkAdapter {
	nodes, err := naq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NetworkAdapter IDs.
func (naq *NetworkAdapterQuery) IDs(ctx context.Context) (ids []int, err error) {
	if naq.ctx.Unique == nil && naq.path != nil {
		naq.Unique(true)
	}
	ctx = setContextOp(ctx, naq.ctx, ent.OpQueryIDs)
	if err = naq.Select(networkadapter.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (naq *NetworkAdapterQuery) IDsX(ctx context.Context) []int {
	ids, err := naq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (naq *NetworkAdapterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, naq.ctx, ent.OpQueryCount)
	if err := naq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, naq, querierCount[*NetworkAdapterQuery](), naq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (naq *NetworkAdapterQuery) CountX(ctx context.Context) int {
	count, err := naq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (naq *NetworkAdapterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, naq.ctx, ent.OpQueryExist)
	switch _, err := naq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (naq *NetworkAdapterQuery) ExistX(ctx context.Context) bool {
	exist, err := naq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NetworkAdapterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (naq *NetworkAdapterQuery) Clone() *NetworkAdapterQuery {
	if naq == nil {
		return nil
	}
	return &NetworkAdapterQuery{
		config:     naq.config,
		ctx:        naq.ctx.Clone(),
		order:      append([]networkadapter.OrderOption{}, naq.order...),
		inters:     append([]Interceptor{}, naq.inters...),
		predicates: append([]predicate.NetworkAdapter{}, naq.predicates...),
		withOwner:  naq.withOwner.Clone(),
		// clone intermediate query.
		sql:       naq.sql.Clone(),
		path:      naq.path,
		modifiers: append([]func(*sql.Selector){}, naq.modifiers...),
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (naq *NetworkAdapterQuery) WithOwner(opts ...func(*AgentQuery)) *NetworkAdapterQuery {
	query := (&AgentClient{config: naq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	naq.withOwner = query
	return naq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NetworkAdapter.Query().
//		GroupBy(networkadapter.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (naq *NetworkAdapterQuery) GroupBy(field string, fields ...string) *NetworkAdapterGroupBy {
	naq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NetworkAdapterGroupBy{build: naq}
	grbuild.flds = &naq.ctx.Fields
	grbuild.label = networkadapter.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.NetworkAdapter.Query().
//		Select(networkadapter.FieldName).
//		Scan(ctx, &v)
func (naq *NetworkAdapterQuery) Select(fields ...string) *NetworkAdapterSelect {
	naq.ctx.Fields = append(naq.ctx.Fields, fields...)
	sbuild := &NetworkAdapterSelect{NetworkAdapterQuery: naq}
	sbuild.label = networkadapter.Label
	sbuild.flds, sbuild.scan = &naq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NetworkAdapterSelect configured with the given aggregations.
func (naq *NetworkAdapterQuery) Aggregate(fns ...AggregateFunc) *NetworkAdapterSelect {
	return naq.Select().Aggregate(fns...)
}

func (naq *NetworkAdapterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range naq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, naq); err != nil {
				return err
			}
		}
	}
	for _, f := range naq.ctx.Fields {
		if !networkadapter.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if naq.path != nil {
		prev, err := naq.path(ctx)
		if err != nil {
			return err
		}
		naq.sql = prev
	}
	return nil
}

func (naq *NetworkAdapterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NetworkAdapter, error) {
	var (
		nodes       = []*NetworkAdapter{}
		withFKs     = naq.withFKs
		_spec       = naq.querySpec()
		loadedTypes = [1]bool{
			naq.withOwner != nil,
		}
	)
	if naq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, networkadapter.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NetworkAdapter).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NetworkAdapter{config: naq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(naq.modifiers) > 0 {
		_spec.Modifiers = naq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, naq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := naq.withOwner; query != nil {
		if err := naq.loadOwner(ctx, query, nodes, nil,
			func(n *NetworkAdapter, e *Agent) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (naq *NetworkAdapterQuery) loadOwner(ctx context.Context, query *AgentQuery, nodes []*NetworkAdapter, init func(*NetworkAdapter), assign func(*NetworkAdapter, *Agent)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*NetworkAdapter)
	for i := range nodes {
		if nodes[i].agent_networkadapters == nil {
			continue
		}
		fk := *nodes[i].agent_networkadapters
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(agent.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "agent_networkadapters" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (naq *NetworkAdapterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := naq.querySpec()
	if len(naq.modifiers) > 0 {
		_spec.Modifiers = naq.modifiers
	}
	_spec.Node.Columns = naq.ctx.Fields
	if len(naq.ctx.Fields) > 0 {
		_spec.Unique = naq.ctx.Unique != nil && *naq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, naq.driver, _spec)
}

func (naq *NetworkAdapterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(networkadapter.Table, networkadapter.Columns, sqlgraph.NewFieldSpec(networkadapter.FieldID, field.TypeInt))
	_spec.From = naq.sql
	if unique := naq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if naq.path != nil {
		_spec.Unique = true
	}
	if fields := naq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, networkadapter.FieldID)
		for i := range fields {
			if fields[i] != networkadapter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := naq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := naq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := naq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := naq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (naq *NetworkAdapterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(naq.driver.Dialect())
	t1 := builder.Table(networkadapter.Table)
	columns := naq.ctx.Fields
	if len(columns) == 0 {
		columns = networkadapter.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if naq.sql != nil {
		selector = naq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if naq.ctx.Unique != nil && *naq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range naq.modifiers {
		m(selector)
	}
	for _, p := range naq.predicates {
		p(selector)
	}
	for _, p := range naq.order {
		p(selector)
	}
	if offset := naq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := naq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (naq *NetworkAdapterQuery) Modify(modifiers ...func(s *sql.Selector)) *NetworkAdapterSelect {
	naq.modifiers = append(naq.modifiers, modifiers...)
	return naq.Select()
}

// NetworkAdapterGroupBy is the group-by builder for NetworkAdapter entities.
type NetworkAdapterGroupBy struct {
	selector
	build *NetworkAdapterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nagb *NetworkAdapterGroupBy) Aggregate(fns ...AggregateFunc) *NetworkAdapterGroupBy {
	nagb.fns = append(nagb.fns, fns...)
	return nagb
}

// Scan applies the selector query and scans the result into the given value.
func (nagb *NetworkAdapterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nagb.build.ctx, ent.OpQueryGroupBy)
	if err := nagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NetworkAdapterQuery, *NetworkAdapterGroupBy](ctx, nagb.build, nagb, nagb.build.inters, v)
}

func (nagb *NetworkAdapterGroupBy) sqlScan(ctx context.Context, root *NetworkAdapterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nagb.fns))
	for _, fn := range nagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nagb.flds)+len(nagb.fns))
		for _, f := range *nagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NetworkAdapterSelect is the builder for selecting fields of NetworkAdapter entities.
type NetworkAdapterSelect struct {
	*NetworkAdapterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nas *NetworkAdapterSelect) Aggregate(fns ...AggregateFunc) *NetworkAdapterSelect {
	nas.fns = append(nas.fns, fns...)
	return nas
}

// Scan applies the selector query and scans the result into the given value.
func (nas *NetworkAdapterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nas.ctx, ent.OpQuerySelect)
	if err := nas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NetworkAdapterQuery, *NetworkAdapterSelect](ctx, nas.NetworkAdapterQuery, nas, nas.inters, v)
}

func (nas *NetworkAdapterSelect) sqlScan(ctx context.Context, root *NetworkAdapterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nas.fns))
	for _, fn := range nas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (nas *NetworkAdapterSelect) Modify(modifiers ...func(s *sql.Selector)) *NetworkAdapterSelect {
	nas.modifiers = append(nas.modifiers, modifiers...)
	return nas
}
