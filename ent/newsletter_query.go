// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/newsletter"
	"app/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// NewsletterQuery is the builder for querying Newsletter entities.
type NewsletterQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Newsletter
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Newsletter) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NewsletterQuery builder.
func (nq *NewsletterQuery) Where(ps ...predicate.Newsletter) *NewsletterQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NewsletterQuery) Limit(limit int) *NewsletterQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NewsletterQuery) Offset(offset int) *NewsletterQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NewsletterQuery) Unique(unique bool) *NewsletterQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NewsletterQuery) Order(o ...OrderFunc) *NewsletterQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// First returns the first Newsletter entity from the query.
// Returns a *NotFoundError when no Newsletter was found.
func (nq *NewsletterQuery) First(ctx context.Context) (*Newsletter, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{newsletter.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NewsletterQuery) FirstX(ctx context.Context) *Newsletter {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Newsletter ID from the query.
// Returns a *NotFoundError when no Newsletter ID was found.
func (nq *NewsletterQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{newsletter.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NewsletterQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Newsletter entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Newsletter entity is found.
// Returns a *NotFoundError when no Newsletter entities are found.
func (nq *NewsletterQuery) Only(ctx context.Context) (*Newsletter, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{newsletter.Label}
	default:
		return nil, &NotSingularError{newsletter.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NewsletterQuery) OnlyX(ctx context.Context) *Newsletter {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Newsletter ID in the query.
// Returns a *NotSingularError when more than one Newsletter ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NewsletterQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{newsletter.Label}
	default:
		err = &NotSingularError{newsletter.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NewsletterQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Newsletters.
func (nq *NewsletterQuery) All(ctx context.Context) ([]*Newsletter, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Newsletter, *NewsletterQuery]()
	return withInterceptors[[]*Newsletter](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NewsletterQuery) AllX(ctx context.Context) []*Newsletter {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Newsletter IDs.
func (nq *NewsletterQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(newsletter.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NewsletterQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NewsletterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NewsletterQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NewsletterQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NewsletterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NewsletterQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NewsletterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NewsletterQuery) Clone() *NewsletterQuery {
	if nq == nil {
		return nil
	}
	return &NewsletterQuery{
		config:     nq.config,
		ctx:        nq.ctx.Clone(),
		order:      append([]OrderFunc{}, nq.order...),
		inters:     append([]Interceptor{}, nq.inters...),
		predicates: append([]predicate.Newsletter{}, nq.predicates...),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UpdatedAt time.Time `json:"updated_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Newsletter.Query().
//		GroupBy(newsletter.FieldUpdatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NewsletterQuery) GroupBy(field string, fields ...string) *NewsletterGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NewsletterGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = newsletter.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UpdatedAt time.Time `json:"updated_at,omitempty"`
//	}
//
//	client.Newsletter.Query().
//		Select(newsletter.FieldUpdatedAt).
//		Scan(ctx, &v)
func (nq *NewsletterQuery) Select(fields ...string) *NewsletterSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NewsletterSelect{NewsletterQuery: nq}
	sbuild.label = newsletter.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NewsletterSelect configured with the given aggregations.
func (nq *NewsletterQuery) Aggregate(fns ...AggregateFunc) *NewsletterSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NewsletterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !newsletter.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NewsletterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Newsletter, error) {
	var (
		nodes = []*Newsletter{}
		_spec = nq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Newsletter).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Newsletter{config: nq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range nq.loadTotal {
		if err := nq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NewsletterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NewsletterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(newsletter.Table, newsletter.Columns, sqlgraph.NewFieldSpec(newsletter.FieldID, field.TypeUUID))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, newsletter.FieldID)
		for i := range fields {
			if fields[i] != newsletter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NewsletterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(newsletter.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = newsletter.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NewsletterGroupBy is the group-by builder for Newsletter entities.
type NewsletterGroupBy struct {
	selector
	build *NewsletterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NewsletterGroupBy) Aggregate(fns ...AggregateFunc) *NewsletterGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NewsletterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NewsletterQuery, *NewsletterGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NewsletterGroupBy) sqlScan(ctx context.Context, root *NewsletterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NewsletterSelect is the builder for selecting fields of Newsletter entities.
type NewsletterSelect struct {
	*NewsletterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NewsletterSelect) Aggregate(fns ...AggregateFunc) *NewsletterSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NewsletterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NewsletterQuery, *NewsletterSelect](ctx, ns.NewsletterQuery, ns, ns.inters, v)
}

func (ns *NewsletterSelect) sqlScan(ctx context.Context, root *NewsletterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
