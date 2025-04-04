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
	"github.com/open-uem/ent/profile"
	"github.com/open-uem/ent/profileissue"
)

// ProfileIssueQuery is the builder for querying ProfileIssue entities.
type ProfileIssueQuery struct {
	config
	ctx         *QueryContext
	order       []profileissue.OrderOption
	inters      []Interceptor
	predicates  []predicate.ProfileIssue
	withProfile *ProfileQuery
	withAgents  *AgentQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProfileIssueQuery builder.
func (piq *ProfileIssueQuery) Where(ps ...predicate.ProfileIssue) *ProfileIssueQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit the number of records to be returned by this query.
func (piq *ProfileIssueQuery) Limit(limit int) *ProfileIssueQuery {
	piq.ctx.Limit = &limit
	return piq
}

// Offset to start from.
func (piq *ProfileIssueQuery) Offset(offset int) *ProfileIssueQuery {
	piq.ctx.Offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *ProfileIssueQuery) Unique(unique bool) *ProfileIssueQuery {
	piq.ctx.Unique = &unique
	return piq
}

// Order specifies how the records should be ordered.
func (piq *ProfileIssueQuery) Order(o ...profileissue.OrderOption) *ProfileIssueQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// QueryProfile chains the current query on the "profile" edge.
func (piq *ProfileIssueQuery) QueryProfile() *ProfileQuery {
	query := (&ProfileClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profileissue.Table, profileissue.FieldID, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, profileissue.ProfileTable, profileissue.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAgents chains the current query on the "agents" edge.
func (piq *ProfileIssueQuery) QueryAgents() *AgentQuery {
	query := (&AgentClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profileissue.Table, profileissue.FieldID, selector),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, profileissue.AgentsTable, profileissue.AgentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProfileIssue entity from the query.
// Returns a *NotFoundError when no ProfileIssue was found.
func (piq *ProfileIssueQuery) First(ctx context.Context) (*ProfileIssue, error) {
	nodes, err := piq.Limit(1).All(setContextOp(ctx, piq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{profileissue.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *ProfileIssueQuery) FirstX(ctx context.Context) *ProfileIssue {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProfileIssue ID from the query.
// Returns a *NotFoundError when no ProfileIssue ID was found.
func (piq *ProfileIssueQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(1).IDs(setContextOp(ctx, piq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{profileissue.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *ProfileIssueQuery) FirstIDX(ctx context.Context) int {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProfileIssue entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProfileIssue entity is found.
// Returns a *NotFoundError when no ProfileIssue entities are found.
func (piq *ProfileIssueQuery) Only(ctx context.Context) (*ProfileIssue, error) {
	nodes, err := piq.Limit(2).All(setContextOp(ctx, piq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{profileissue.Label}
	default:
		return nil, &NotSingularError{profileissue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *ProfileIssueQuery) OnlyX(ctx context.Context) *ProfileIssue {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProfileIssue ID in the query.
// Returns a *NotSingularError when more than one ProfileIssue ID is found.
// Returns a *NotFoundError when no entities are found.
func (piq *ProfileIssueQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(2).IDs(setContextOp(ctx, piq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{profileissue.Label}
	default:
		err = &NotSingularError{profileissue.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *ProfileIssueQuery) OnlyIDX(ctx context.Context) int {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProfileIssues.
func (piq *ProfileIssueQuery) All(ctx context.Context) ([]*ProfileIssue, error) {
	ctx = setContextOp(ctx, piq.ctx, ent.OpQueryAll)
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProfileIssue, *ProfileIssueQuery]()
	return withInterceptors[[]*ProfileIssue](ctx, piq, qr, piq.inters)
}

// AllX is like All, but panics if an error occurs.
func (piq *ProfileIssueQuery) AllX(ctx context.Context) []*ProfileIssue {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProfileIssue IDs.
func (piq *ProfileIssueQuery) IDs(ctx context.Context) (ids []int, err error) {
	if piq.ctx.Unique == nil && piq.path != nil {
		piq.Unique(true)
	}
	ctx = setContextOp(ctx, piq.ctx, ent.OpQueryIDs)
	if err = piq.Select(profileissue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *ProfileIssueQuery) IDsX(ctx context.Context) []int {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *ProfileIssueQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, piq.ctx, ent.OpQueryCount)
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, piq, querierCount[*ProfileIssueQuery](), piq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (piq *ProfileIssueQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *ProfileIssueQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, piq.ctx, ent.OpQueryExist)
	switch _, err := piq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *ProfileIssueQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProfileIssueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *ProfileIssueQuery) Clone() *ProfileIssueQuery {
	if piq == nil {
		return nil
	}
	return &ProfileIssueQuery{
		config:      piq.config,
		ctx:         piq.ctx.Clone(),
		order:       append([]profileissue.OrderOption{}, piq.order...),
		inters:      append([]Interceptor{}, piq.inters...),
		predicates:  append([]predicate.ProfileIssue{}, piq.predicates...),
		withProfile: piq.withProfile.Clone(),
		withAgents:  piq.withAgents.Clone(),
		// clone intermediate query.
		sql:       piq.sql.Clone(),
		path:      piq.path,
		modifiers: append([]func(*sql.Selector){}, piq.modifiers...),
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProfileIssueQuery) WithProfile(opts ...func(*ProfileQuery)) *ProfileIssueQuery {
	query := (&ProfileClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withProfile = query
	return piq
}

// WithAgents tells the query-builder to eager-load the nodes that are connected to
// the "agents" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProfileIssueQuery) WithAgents(opts ...func(*AgentQuery)) *ProfileIssueQuery {
	query := (&AgentClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withAgents = query
	return piq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Error string `json:"error,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProfileIssue.Query().
//		GroupBy(profileissue.FieldError).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (piq *ProfileIssueQuery) GroupBy(field string, fields ...string) *ProfileIssueGroupBy {
	piq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProfileIssueGroupBy{build: piq}
	grbuild.flds = &piq.ctx.Fields
	grbuild.label = profileissue.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Error string `json:"error,omitempty"`
//	}
//
//	client.ProfileIssue.Query().
//		Select(profileissue.FieldError).
//		Scan(ctx, &v)
func (piq *ProfileIssueQuery) Select(fields ...string) *ProfileIssueSelect {
	piq.ctx.Fields = append(piq.ctx.Fields, fields...)
	sbuild := &ProfileIssueSelect{ProfileIssueQuery: piq}
	sbuild.label = profileissue.Label
	sbuild.flds, sbuild.scan = &piq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProfileIssueSelect configured with the given aggregations.
func (piq *ProfileIssueQuery) Aggregate(fns ...AggregateFunc) *ProfileIssueSelect {
	return piq.Select().Aggregate(fns...)
}

func (piq *ProfileIssueQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range piq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, piq); err != nil {
				return err
			}
		}
	}
	for _, f := range piq.ctx.Fields {
		if !profileissue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *ProfileIssueQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProfileIssue, error) {
	var (
		nodes       = []*ProfileIssue{}
		withFKs     = piq.withFKs
		_spec       = piq.querySpec()
		loadedTypes = [2]bool{
			piq.withProfile != nil,
			piq.withAgents != nil,
		}
	)
	if piq.withProfile != nil || piq.withAgents != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, profileissue.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProfileIssue).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProfileIssue{config: piq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(piq.modifiers) > 0 {
		_spec.Modifiers = piq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := piq.withProfile; query != nil {
		if err := piq.loadProfile(ctx, query, nodes, nil,
			func(n *ProfileIssue, e *Profile) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	if query := piq.withAgents; query != nil {
		if err := piq.loadAgents(ctx, query, nodes, nil,
			func(n *ProfileIssue, e *Agent) { n.Edges.Agents = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (piq *ProfileIssueQuery) loadProfile(ctx context.Context, query *ProfileQuery, nodes []*ProfileIssue, init func(*ProfileIssue), assign func(*ProfileIssue, *Profile)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProfileIssue)
	for i := range nodes {
		if nodes[i].profile_issues == nil {
			continue
		}
		fk := *nodes[i].profile_issues
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(profile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "profile_issues" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (piq *ProfileIssueQuery) loadAgents(ctx context.Context, query *AgentQuery, nodes []*ProfileIssue, init func(*ProfileIssue), assign func(*ProfileIssue, *Agent)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ProfileIssue)
	for i := range nodes {
		if nodes[i].profile_issue_agents == nil {
			continue
		}
		fk := *nodes[i].profile_issue_agents
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
			return fmt.Errorf(`unexpected foreign-key "profile_issue_agents" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (piq *ProfileIssueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	if len(piq.modifiers) > 0 {
		_spec.Modifiers = piq.modifiers
	}
	_spec.Node.Columns = piq.ctx.Fields
	if len(piq.ctx.Fields) > 0 {
		_spec.Unique = piq.ctx.Unique != nil && *piq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *ProfileIssueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(profileissue.Table, profileissue.Columns, sqlgraph.NewFieldSpec(profileissue.FieldID, field.TypeInt))
	_spec.From = piq.sql
	if unique := piq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if piq.path != nil {
		_spec.Unique = true
	}
	if fields := piq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profileissue.FieldID)
		for i := range fields {
			if fields[i] != profileissue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *ProfileIssueQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(profileissue.Table)
	columns := piq.ctx.Fields
	if len(columns) == 0 {
		columns = profileissue.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.ctx.Unique != nil && *piq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range piq.modifiers {
		m(selector)
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (piq *ProfileIssueQuery) Modify(modifiers ...func(s *sql.Selector)) *ProfileIssueSelect {
	piq.modifiers = append(piq.modifiers, modifiers...)
	return piq.Select()
}

// ProfileIssueGroupBy is the group-by builder for ProfileIssue entities.
type ProfileIssueGroupBy struct {
	selector
	build *ProfileIssueQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *ProfileIssueGroupBy) Aggregate(fns ...AggregateFunc) *ProfileIssueGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the selector query and scans the result into the given value.
func (pigb *ProfileIssueGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pigb.build.ctx, ent.OpQueryGroupBy)
	if err := pigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfileIssueQuery, *ProfileIssueGroupBy](ctx, pigb.build, pigb, pigb.build.inters, v)
}

func (pigb *ProfileIssueGroupBy) sqlScan(ctx context.Context, root *ProfileIssueQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pigb.flds)+len(pigb.fns))
		for _, f := range *pigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProfileIssueSelect is the builder for selecting fields of ProfileIssue entities.
type ProfileIssueSelect struct {
	*ProfileIssueQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pis *ProfileIssueSelect) Aggregate(fns ...AggregateFunc) *ProfileIssueSelect {
	pis.fns = append(pis.fns, fns...)
	return pis
}

// Scan applies the selector query and scans the result into the given value.
func (pis *ProfileIssueSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pis.ctx, ent.OpQuerySelect)
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfileIssueQuery, *ProfileIssueSelect](ctx, pis.ProfileIssueQuery, pis, pis.inters, v)
}

func (pis *ProfileIssueSelect) sqlScan(ctx context.Context, root *ProfileIssueQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pis.fns))
	for _, fn := range pis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pis *ProfileIssueSelect) Modify(modifiers ...func(s *sql.Selector)) *ProfileIssueSelect {
	pis.modifiers = append(pis.modifiers, modifiers...)
	return pis
}
