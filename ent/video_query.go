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
	"github.com/zibbp/eos/ent/channel"
	"github.com/zibbp/eos/ent/chapter"
	"github.com/zibbp/eos/ent/comment"
	"github.com/zibbp/eos/ent/predicate"
	"github.com/zibbp/eos/ent/video"
)

// VideoQuery is the builder for querying Video entities.
type VideoQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	inters       []Interceptor
	predicates   []predicate.Video
	withChannel  *ChannelQuery
	withChapters *ChapterQuery
	withComments *CommentQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideoQuery builder.
func (vq *VideoQuery) Where(ps ...predicate.Video) *VideoQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VideoQuery) Limit(limit int) *VideoQuery {
	vq.limit = &limit
	return vq
}

// Offset to start from.
func (vq *VideoQuery) Offset(offset int) *VideoQuery {
	vq.offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VideoQuery) Unique(unique bool) *VideoQuery {
	vq.unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VideoQuery) Order(o ...OrderFunc) *VideoQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryChannel chains the current query on the "channel" edge.
func (vq *VideoQuery) QueryChannel() *ChannelQuery {
	query := (&ChannelClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, selector),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, video.ChannelTable, video.ChannelColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChapters chains the current query on the "chapters" edge.
func (vq *VideoQuery) QueryChapters() *ChapterQuery {
	query := (&ChapterClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, selector),
			sqlgraph.To(chapter.Table, chapter.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, video.ChaptersTable, video.ChaptersColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryComments chains the current query on the "comments" edge.
func (vq *VideoQuery) QueryComments() *CommentQuery {
	query := (&CommentClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, video.CommentsTable, video.CommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Video entity from the query.
// Returns a *NotFoundError when no Video was found.
func (vq *VideoQuery) First(ctx context.Context) (*Video, error) {
	nodes, err := vq.Limit(1).All(newQueryContext(ctx, TypeVideo, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{video.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VideoQuery) FirstX(ctx context.Context) *Video {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Video ID from the query.
// Returns a *NotFoundError when no Video ID was found.
func (vq *VideoQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vq.Limit(1).IDs(newQueryContext(ctx, TypeVideo, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{video.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VideoQuery) FirstIDX(ctx context.Context) string {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Video entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Video entity is found.
// Returns a *NotFoundError when no Video entities are found.
func (vq *VideoQuery) Only(ctx context.Context) (*Video, error) {
	nodes, err := vq.Limit(2).All(newQueryContext(ctx, TypeVideo, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{video.Label}
	default:
		return nil, &NotSingularError{video.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VideoQuery) OnlyX(ctx context.Context) *Video {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Video ID in the query.
// Returns a *NotSingularError when more than one Video ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VideoQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vq.Limit(2).IDs(newQueryContext(ctx, TypeVideo, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{video.Label}
	default:
		err = &NotSingularError{video.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VideoQuery) OnlyIDX(ctx context.Context) string {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Videos.
func (vq *VideoQuery) All(ctx context.Context) ([]*Video, error) {
	ctx = newQueryContext(ctx, TypeVideo, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Video, *VideoQuery]()
	return withInterceptors[[]*Video](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VideoQuery) AllX(ctx context.Context) []*Video {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Video IDs.
func (vq *VideoQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	ctx = newQueryContext(ctx, TypeVideo, "IDs")
	if err := vq.Select(video.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VideoQuery) IDsX(ctx context.Context) []string {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VideoQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeVideo, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VideoQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VideoQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VideoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeVideo, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VideoQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VideoQuery) Clone() *VideoQuery {
	if vq == nil {
		return nil
	}
	return &VideoQuery{
		config:       vq.config,
		limit:        vq.limit,
		offset:       vq.offset,
		order:        append([]OrderFunc{}, vq.order...),
		inters:       append([]Interceptor{}, vq.inters...),
		predicates:   append([]predicate.Video{}, vq.predicates...),
		withChannel:  vq.withChannel.Clone(),
		withChapters: vq.withChapters.Clone(),
		withComments: vq.withComments.Clone(),
		// clone intermediate query.
		sql:    vq.sql.Clone(),
		path:   vq.path,
		unique: vq.unique,
	}
}

// WithChannel tells the query-builder to eager-load the nodes that are connected to
// the "channel" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideoQuery) WithChannel(opts ...func(*ChannelQuery)) *VideoQuery {
	query := (&ChannelClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withChannel = query
	return vq
}

// WithChapters tells the query-builder to eager-load the nodes that are connected to
// the "chapters" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideoQuery) WithChapters(opts ...func(*ChapterQuery)) *VideoQuery {
	query := (&ChapterClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withChapters = query
	return vq
}

// WithComments tells the query-builder to eager-load the nodes that are connected to
// the "comments" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideoQuery) WithComments(opts ...func(*CommentQuery)) *VideoQuery {
	query := (&CommentClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withComments = query
	return vq
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
//	client.Video.Query().
//		GroupBy(video.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VideoQuery) GroupBy(field string, fields ...string) *VideoGroupBy {
	vq.fields = append([]string{field}, fields...)
	grbuild := &VideoGroupBy{build: vq}
	grbuild.flds = &vq.fields
	grbuild.label = video.Label
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
//	client.Video.Query().
//		Select(video.FieldTitle).
//		Scan(ctx, &v)
func (vq *VideoQuery) Select(fields ...string) *VideoSelect {
	vq.fields = append(vq.fields, fields...)
	sbuild := &VideoSelect{VideoQuery: vq}
	sbuild.label = video.Label
	sbuild.flds, sbuild.scan = &vq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideoSelect configured with the given aggregations.
func (vq *VideoQuery) Aggregate(fns ...AggregateFunc) *VideoSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VideoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.fields {
		if !video.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VideoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Video, error) {
	var (
		nodes       = []*Video{}
		withFKs     = vq.withFKs
		_spec       = vq.querySpec()
		loadedTypes = [3]bool{
			vq.withChannel != nil,
			vq.withChapters != nil,
			vq.withComments != nil,
		}
	)
	if vq.withChannel != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, video.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Video).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Video{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withChannel; query != nil {
		if err := vq.loadChannel(ctx, query, nodes, nil,
			func(n *Video, e *Channel) { n.Edges.Channel = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withChapters; query != nil {
		if err := vq.loadChapters(ctx, query, nodes,
			func(n *Video) { n.Edges.Chapters = []*Chapter{} },
			func(n *Video, e *Chapter) { n.Edges.Chapters = append(n.Edges.Chapters, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withComments; query != nil {
		if err := vq.loadComments(ctx, query, nodes,
			func(n *Video) { n.Edges.Comments = []*Comment{} },
			func(n *Video, e *Comment) { n.Edges.Comments = append(n.Edges.Comments, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VideoQuery) loadChannel(ctx context.Context, query *ChannelQuery, nodes []*Video, init func(*Video), assign func(*Video, *Channel)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Video)
	for i := range nodes {
		if nodes[i].channel_videos == nil {
			continue
		}
		fk := *nodes[i].channel_videos
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(channel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "channel_videos" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VideoQuery) loadChapters(ctx context.Context, query *ChapterQuery, nodes []*Video, init func(*Video), assign func(*Video, *Chapter)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Video)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.InValues(video.ChaptersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.video_chapters
		if fk == nil {
			return fmt.Errorf(`foreign-key "video_chapters" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "video_chapters" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VideoQuery) loadComments(ctx context.Context, query *CommentQuery, nodes []*Video, init func(*Video), assign func(*Video, *Comment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Video)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.InValues(video.CommentsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.video_comments
		if fk == nil {
			return fmt.Errorf(`foreign-key "video_comments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "video_comments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vq *VideoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.fields
	if len(vq.fields) > 0 {
		_spec.Unique = vq.unique != nil && *vq.unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VideoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   video.Table,
			Columns: video.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: video.FieldID,
			},
		},
		From:   vq.sql,
		Unique: false,
	}
	if unique := vq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, video.FieldID)
		for i := range fields {
			if fields[i] != video.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VideoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(video.Table)
	columns := vq.fields
	if len(columns) == 0 {
		columns = video.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.unique != nil && *vq.unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VideoGroupBy is the group-by builder for Video entities.
type VideoGroupBy struct {
	selector
	build *VideoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VideoGroupBy) Aggregate(fns ...AggregateFunc) *VideoGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VideoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeVideo, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoQuery, *VideoGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VideoGroupBy) sqlScan(ctx context.Context, root *VideoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VideoSelect is the builder for selecting fields of Video entities.
type VideoSelect struct {
	*VideoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VideoSelect) Aggregate(fns ...AggregateFunc) *VideoSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VideoSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeVideo, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoQuery, *VideoSelect](ctx, vs.VideoQuery, vs, vs.inters, v)
}

func (vs *VideoSelect) sqlScan(ctx context.Context, root *VideoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}