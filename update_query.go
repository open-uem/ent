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
	"github.com/open-uem/ent/predicate"
	"github.com/open-uem/ent/update"
)

// UpdateQuery is the builder for querying Update entities.
type UpdateQuery struct {
	config
	ctx        *QueryContext
	order      []update.OrderOption
	inters     []Interceptor
	predicates []predicate.Update
	withOwner  *AgentQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UpdateQuery builder.
func (uq *UpdateQuery) Where(ps ...predicate.Update) *UpdateQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UpdateQuery) Limit(limit int) *UpdateQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UpdateQuery) Offset(offset int) *UpdateQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UpdateQuery) Unique(unique bool) *UpdateQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UpdateQuery) Order(o ...update.OrderOption) *UpdateQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryOwner chains the current query on the "owner" edge.
func (uq *UpdateQuery) QueryOwner() *AgentQuery {
	query := (&AgentClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(update.Table, update.FieldID, selector),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, update.OwnerTable, update.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Update entity from the query.
// Returns a *NotFoundError when no Update was found.
func (uq *UpdateQuery) First(ctx context.Context) (*Update, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{update.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UpdateQuery) FirstX(ctx context.Context) *Update {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Update ID from the query.
// Returns a *NotFoundError when no Update ID was found.
func (uq *UpdateQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{update.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UpdateQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Update entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Update entity is found.
// Returns a *NotFoundError when no Update entities are found.
func (uq *UpdateQuery) Only(ctx context.Context) (*Update, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{update.Label}
	default:
		return nil, &NotSingularError{update.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UpdateQuery) OnlyX(ctx context.Context) *Update {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Update ID in the query.
// Returns a *NotSingularError when more than one Update ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UpdateQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{update.Label}
	default:
		err = &NotSingularError{update.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UpdateQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Updates.
func (uq *UpdateQuery) All(ctx context.Context) ([]*Update, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryAll)
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Update, *UpdateQuery]()
	return withInterceptors[[]*Update](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UpdateQuery) AllX(ctx context.Context) []*Update {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Update IDs.
func (uq *UpdateQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryIDs)
	if err = uq.Select(update.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UpdateQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UpdateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryCount)
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UpdateQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UpdateQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UpdateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryExist)
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UpdateQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UpdateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UpdateQuery) Clone() *UpdateQuery {
	if uq == nil {
		return nil
	}
	return &UpdateQuery{
		config:     uq.config,
		ctx:        uq.ctx.Clone(),
		order:      append([]update.OrderOption{}, uq.order...),
		inters:     append([]Interceptor{}, uq.inters...),
		predicates: append([]predicate.Update{}, uq.predicates...),
		withOwner:  uq.withOwner.Clone(),
		// clone intermediate query.
		sql:       uq.sql.Clone(),
		path:      uq.path,
		modifiers: append([]func(*sql.Selector){}, uq.modifiers...),
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UpdateQuery) WithOwner(opts ...func(*AgentQuery)) *UpdateQuery {
	query := (&AgentClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withOwner = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Update.Query().
//		GroupBy(update.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UpdateQuery) GroupBy(field string, fields ...string) *UpdateGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UpdateGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = update.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Update.Query().
//		Select(update.FieldTitle).
//		Scan(ctx, &v)
func (uq *UpdateQuery) Select(fields ...string) *UpdateSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UpdateSelect{UpdateQuery: uq}
	sbuild.label = update.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UpdateSelect configured with the given aggregations.
func (uq *UpdateQuery) Aggregate(fns ...AggregateFunc) *UpdateSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UpdateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !update.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UpdateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Update, error) {
	var (
		nodes       = []*Update{}
		withFKs     = uq.withFKs
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withOwner != nil,
		}
	)
	if uq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, update.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Update).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Update{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withOwner; query != nil {
		if err := uq.loadOwner(ctx, query, nodes, nil,
			func(n *Update, e *Agent) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UpdateQuery) loadOwner(ctx context.Context, query *AgentQuery, nodes []*Update, init func(*Update), assign func(*Update, *Agent)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Update)
	for i := range nodes {
		if nodes[i].agent_updates == nil {
			continue
		}
		fk := *nodes[i].agent_updates
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
			return fmt.Errorf(`unexpected foreign-key "agent_updates" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uq *UpdateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UpdateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(update.Table, update.Columns, sqlgraph.NewFieldSpec(update.FieldID, field.TypeInt))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, update.FieldID)
		for i := range fields {
			if fields[i] != update.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UpdateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(update.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = update.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range uq.modifiers {
		m(selector)
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (uq *UpdateQuery) Modify(modifiers ...func(s *sql.Selector)) *UpdateSelect {
	uq.modifiers = append(uq.modifiers, modifiers...)
	return uq.Select()
}

// UpdateGroupBy is the group-by builder for Update entities.
type UpdateGroupBy struct {
	selector
	build *UpdateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UpdateGroupBy) Aggregate(fns ...AggregateFunc) *UpdateGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UpdateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, ent.OpQueryGroupBy)
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpdateQuery, *UpdateGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UpdateGroupBy) sqlScan(ctx context.Context, root *UpdateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UpdateSelect is the builder for selecting fields of Update entities.
type UpdateSelect struct {
	*UpdateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UpdateSelect) Aggregate(fns ...AggregateFunc) *UpdateSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UpdateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, ent.OpQuerySelect)
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpdateQuery, *UpdateSelect](ctx, us.UpdateQuery, us, us.inters, v)
}

func (us *UpdateSelect) sqlScan(ctx context.Context, root *UpdateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (us *UpdateSelect) Modify(modifiers ...func(s *sql.Selector)) *UpdateSelect {
	us.modifiers = append(us.modifiers, modifiers...)
	return us
}
