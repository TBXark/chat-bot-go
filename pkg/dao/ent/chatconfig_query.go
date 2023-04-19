// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chatconfig"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/predicate"
)

// ChatConfigQuery is the builder for querying ChatConfig entities.
type ChatConfigQuery struct {
	config
	ctx        *QueryContext
	order      []chatconfig.OrderOption
	inters     []Interceptor
	predicates []predicate.ChatConfig
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChatConfigQuery builder.
func (ccq *ChatConfigQuery) Where(ps ...predicate.ChatConfig) *ChatConfigQuery {
	ccq.predicates = append(ccq.predicates, ps...)
	return ccq
}

// Limit the number of records to be returned by this query.
func (ccq *ChatConfigQuery) Limit(limit int) *ChatConfigQuery {
	ccq.ctx.Limit = &limit
	return ccq
}

// Offset to start from.
func (ccq *ChatConfigQuery) Offset(offset int) *ChatConfigQuery {
	ccq.ctx.Offset = &offset
	return ccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccq *ChatConfigQuery) Unique(unique bool) *ChatConfigQuery {
	ccq.ctx.Unique = &unique
	return ccq
}

// Order specifies how the records should be ordered.
func (ccq *ChatConfigQuery) Order(o ...chatconfig.OrderOption) *ChatConfigQuery {
	ccq.order = append(ccq.order, o...)
	return ccq
}

// First returns the first ChatConfig entity from the query.
// Returns a *NotFoundError when no ChatConfig was found.
func (ccq *ChatConfigQuery) First(ctx context.Context) (*ChatConfig, error) {
	nodes, err := ccq.Limit(1).All(setContextOp(ctx, ccq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chatconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccq *ChatConfigQuery) FirstX(ctx context.Context) *ChatConfig {
	node, err := ccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChatConfig ID from the query.
// Returns a *NotFoundError when no ChatConfig ID was found.
func (ccq *ChatConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ccq.Limit(1).IDs(setContextOp(ctx, ccq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chatconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccq *ChatConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := ccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChatConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChatConfig entity is found.
// Returns a *NotFoundError when no ChatConfig entities are found.
func (ccq *ChatConfigQuery) Only(ctx context.Context) (*ChatConfig, error) {
	nodes, err := ccq.Limit(2).All(setContextOp(ctx, ccq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chatconfig.Label}
	default:
		return nil, &NotSingularError{chatconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccq *ChatConfigQuery) OnlyX(ctx context.Context) *ChatConfig {
	node, err := ccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChatConfig ID in the query.
// Returns a *NotSingularError when more than one ChatConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (ccq *ChatConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ccq.Limit(2).IDs(setContextOp(ctx, ccq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chatconfig.Label}
	default:
		err = &NotSingularError{chatconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccq *ChatConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := ccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChatConfigs.
func (ccq *ChatConfigQuery) All(ctx context.Context) ([]*ChatConfig, error) {
	ctx = setContextOp(ctx, ccq.ctx, "All")
	if err := ccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ChatConfig, *ChatConfigQuery]()
	return withInterceptors[[]*ChatConfig](ctx, ccq, qr, ccq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ccq *ChatConfigQuery) AllX(ctx context.Context) []*ChatConfig {
	nodes, err := ccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChatConfig IDs.
func (ccq *ChatConfigQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ccq.ctx.Unique == nil && ccq.path != nil {
		ccq.Unique(true)
	}
	ctx = setContextOp(ctx, ccq.ctx, "IDs")
	if err = ccq.Select(chatconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccq *ChatConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := ccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccq *ChatConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ccq.ctx, "Count")
	if err := ccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ccq, querierCount[*ChatConfigQuery](), ccq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ccq *ChatConfigQuery) CountX(ctx context.Context) int {
	count, err := ccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccq *ChatConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ccq.ctx, "Exist")
	switch _, err := ccq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ccq *ChatConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := ccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChatConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccq *ChatConfigQuery) Clone() *ChatConfigQuery {
	if ccq == nil {
		return nil
	}
	return &ChatConfigQuery{
		config:     ccq.config,
		ctx:        ccq.ctx.Clone(),
		order:      append([]chatconfig.OrderOption{}, ccq.order...),
		inters:     append([]Interceptor{}, ccq.inters...),
		predicates: append([]predicate.ChatConfig{}, ccq.predicates...),
		// clone intermediate query.
		sql:  ccq.sql.Clone(),
		path: ccq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ChatID int64 `json:"chat_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ChatConfig.Query().
//		GroupBy(chatconfig.FieldChatID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ccq *ChatConfigQuery) GroupBy(field string, fields ...string) *ChatConfigGroupBy {
	ccq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChatConfigGroupBy{build: ccq}
	grbuild.flds = &ccq.ctx.Fields
	grbuild.label = chatconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ChatID int64 `json:"chat_id,omitempty"`
//	}
//
//	client.ChatConfig.Query().
//		Select(chatconfig.FieldChatID).
//		Scan(ctx, &v)
func (ccq *ChatConfigQuery) Select(fields ...string) *ChatConfigSelect {
	ccq.ctx.Fields = append(ccq.ctx.Fields, fields...)
	sbuild := &ChatConfigSelect{ChatConfigQuery: ccq}
	sbuild.label = chatconfig.Label
	sbuild.flds, sbuild.scan = &ccq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChatConfigSelect configured with the given aggregations.
func (ccq *ChatConfigQuery) Aggregate(fns ...AggregateFunc) *ChatConfigSelect {
	return ccq.Select().Aggregate(fns...)
}

func (ccq *ChatConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ccq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ccq); err != nil {
				return err
			}
		}
	}
	for _, f := range ccq.ctx.Fields {
		if !chatconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccq.path != nil {
		prev, err := ccq.path(ctx)
		if err != nil {
			return err
		}
		ccq.sql = prev
	}
	return nil
}

func (ccq *ChatConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChatConfig, error) {
	var (
		nodes = []*ChatConfig{}
		_spec = ccq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ChatConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ChatConfig{config: ccq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ccq *ChatConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccq.querySpec()
	_spec.Node.Columns = ccq.ctx.Fields
	if len(ccq.ctx.Fields) > 0 {
		_spec.Unique = ccq.ctx.Unique != nil && *ccq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ccq.driver, _spec)
}

func (ccq *ChatConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chatconfig.Table, chatconfig.Columns, sqlgraph.NewFieldSpec(chatconfig.FieldID, field.TypeInt))
	_spec.From = ccq.sql
	if unique := ccq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ccq.path != nil {
		_spec.Unique = true
	}
	if fields := ccq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatconfig.FieldID)
		for i := range fields {
			if fields[i] != chatconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccq *ChatConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccq.driver.Dialect())
	t1 := builder.Table(chatconfig.Table)
	columns := ccq.ctx.Fields
	if len(columns) == 0 {
		columns = chatconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccq.sql != nil {
		selector = ccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccq.ctx.Unique != nil && *ccq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ccq.predicates {
		p(selector)
	}
	for _, p := range ccq.order {
		p(selector)
	}
	if offset := ccq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChatConfigGroupBy is the group-by builder for ChatConfig entities.
type ChatConfigGroupBy struct {
	selector
	build *ChatConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccgb *ChatConfigGroupBy) Aggregate(fns ...AggregateFunc) *ChatConfigGroupBy {
	ccgb.fns = append(ccgb.fns, fns...)
	return ccgb
}

// Scan applies the selector query and scans the result into the given value.
func (ccgb *ChatConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccgb.build.ctx, "GroupBy")
	if err := ccgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatConfigQuery, *ChatConfigGroupBy](ctx, ccgb.build, ccgb, ccgb.build.inters, v)
}

func (ccgb *ChatConfigGroupBy) sqlScan(ctx context.Context, root *ChatConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ccgb.fns))
	for _, fn := range ccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ccgb.flds)+len(ccgb.fns))
		for _, f := range *ccgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ccgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChatConfigSelect is the builder for selecting fields of ChatConfig entities.
type ChatConfigSelect struct {
	*ChatConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ccs *ChatConfigSelect) Aggregate(fns ...AggregateFunc) *ChatConfigSelect {
	ccs.fns = append(ccs.fns, fns...)
	return ccs
}

// Scan applies the selector query and scans the result into the given value.
func (ccs *ChatConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccs.ctx, "Select")
	if err := ccs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatConfigQuery, *ChatConfigSelect](ctx, ccs.ChatConfigQuery, ccs, ccs.inters, v)
}

func (ccs *ChatConfigSelect) sqlScan(ctx context.Context, root *ChatConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ccs.fns))
	for _, fn := range ccs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ccs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
