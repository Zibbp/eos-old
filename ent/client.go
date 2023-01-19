// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/zibbp/eos/ent/migrate"

	"github.com/zibbp/eos/ent/channel"
	"github.com/zibbp/eos/ent/chapter"
	"github.com/zibbp/eos/ent/comment"
	"github.com/zibbp/eos/ent/video"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Channel is the client for interacting with the Channel builders.
	Channel *ChannelClient
	// Chapter is the client for interacting with the Chapter builders.
	Chapter *ChapterClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// Video is the client for interacting with the Video builders.
	Video *VideoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Channel = NewChannelClient(c.config)
	c.Chapter = NewChapterClient(c.config)
	c.Comment = NewCommentClient(c.config)
	c.Video = NewVideoClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Channel: NewChannelClient(cfg),
		Chapter: NewChapterClient(cfg),
		Comment: NewCommentClient(cfg),
		Video:   NewVideoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Channel: NewChannelClient(cfg),
		Chapter: NewChapterClient(cfg),
		Comment: NewCommentClient(cfg),
		Video:   NewVideoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Channel.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Channel.Use(hooks...)
	c.Chapter.Use(hooks...)
	c.Comment.Use(hooks...)
	c.Video.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Channel.Intercept(interceptors...)
	c.Chapter.Intercept(interceptors...)
	c.Comment.Intercept(interceptors...)
	c.Video.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ChannelMutation:
		return c.Channel.mutate(ctx, m)
	case *ChapterMutation:
		return c.Chapter.mutate(ctx, m)
	case *CommentMutation:
		return c.Comment.mutate(ctx, m)
	case *VideoMutation:
		return c.Video.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ChannelClient is a client for the Channel schema.
type ChannelClient struct {
	config
}

// NewChannelClient returns a client for the Channel from the given config.
func NewChannelClient(c config) *ChannelClient {
	return &ChannelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `channel.Hooks(f(g(h())))`.
func (c *ChannelClient) Use(hooks ...Hook) {
	c.hooks.Channel = append(c.hooks.Channel, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `channel.Intercept(f(g(h())))`.
func (c *ChannelClient) Intercept(interceptors ...Interceptor) {
	c.inters.Channel = append(c.inters.Channel, interceptors...)
}

// Create returns a builder for creating a Channel entity.
func (c *ChannelClient) Create() *ChannelCreate {
	mutation := newChannelMutation(c.config, OpCreate)
	return &ChannelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Channel entities.
func (c *ChannelClient) CreateBulk(builders ...*ChannelCreate) *ChannelCreateBulk {
	return &ChannelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Channel.
func (c *ChannelClient) Update() *ChannelUpdate {
	mutation := newChannelMutation(c.config, OpUpdate)
	return &ChannelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChannelClient) UpdateOne(ch *Channel) *ChannelUpdateOne {
	mutation := newChannelMutation(c.config, OpUpdateOne, withChannel(ch))
	return &ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChannelClient) UpdateOneID(id string) *ChannelUpdateOne {
	mutation := newChannelMutation(c.config, OpUpdateOne, withChannelID(id))
	return &ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Channel.
func (c *ChannelClient) Delete() *ChannelDelete {
	mutation := newChannelMutation(c.config, OpDelete)
	return &ChannelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChannelClient) DeleteOne(ch *Channel) *ChannelDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChannelClient) DeleteOneID(id string) *ChannelDeleteOne {
	builder := c.Delete().Where(channel.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChannelDeleteOne{builder}
}

// Query returns a query builder for Channel.
func (c *ChannelClient) Query() *ChannelQuery {
	return &ChannelQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a Channel entity by its id.
func (c *ChannelClient) Get(ctx context.Context, id string) (*Channel, error) {
	return c.Query().Where(channel.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChannelClient) GetX(ctx context.Context, id string) *Channel {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVideos queries the videos edge of a Channel.
func (c *ChannelClient) QueryVideos(ch *Channel) *VideoQuery {
	query := (&VideoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(channel.Table, channel.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, channel.VideosTable, channel.VideosColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChannelClient) Hooks() []Hook {
	return c.hooks.Channel
}

// Interceptors returns the client interceptors.
func (c *ChannelClient) Interceptors() []Interceptor {
	return c.inters.Channel
}

func (c *ChannelClient) mutate(ctx context.Context, m *ChannelMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChannelCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChannelUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChannelDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Channel mutation op: %q", m.Op())
	}
}

// ChapterClient is a client for the Chapter schema.
type ChapterClient struct {
	config
}

// NewChapterClient returns a client for the Chapter from the given config.
func NewChapterClient(c config) *ChapterClient {
	return &ChapterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chapter.Hooks(f(g(h())))`.
func (c *ChapterClient) Use(hooks ...Hook) {
	c.hooks.Chapter = append(c.hooks.Chapter, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chapter.Intercept(f(g(h())))`.
func (c *ChapterClient) Intercept(interceptors ...Interceptor) {
	c.inters.Chapter = append(c.inters.Chapter, interceptors...)
}

// Create returns a builder for creating a Chapter entity.
func (c *ChapterClient) Create() *ChapterCreate {
	mutation := newChapterMutation(c.config, OpCreate)
	return &ChapterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Chapter entities.
func (c *ChapterClient) CreateBulk(builders ...*ChapterCreate) *ChapterCreateBulk {
	return &ChapterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Chapter.
func (c *ChapterClient) Update() *ChapterUpdate {
	mutation := newChapterMutation(c.config, OpUpdate)
	return &ChapterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChapterClient) UpdateOne(ch *Chapter) *ChapterUpdateOne {
	mutation := newChapterMutation(c.config, OpUpdateOne, withChapter(ch))
	return &ChapterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChapterClient) UpdateOneID(id string) *ChapterUpdateOne {
	mutation := newChapterMutation(c.config, OpUpdateOne, withChapterID(id))
	return &ChapterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Chapter.
func (c *ChapterClient) Delete() *ChapterDelete {
	mutation := newChapterMutation(c.config, OpDelete)
	return &ChapterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChapterClient) DeleteOne(ch *Chapter) *ChapterDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChapterClient) DeleteOneID(id string) *ChapterDeleteOne {
	builder := c.Delete().Where(chapter.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChapterDeleteOne{builder}
}

// Query returns a query builder for Chapter.
func (c *ChapterClient) Query() *ChapterQuery {
	return &ChapterQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a Chapter entity by its id.
func (c *ChapterClient) Get(ctx context.Context, id string) (*Chapter, error) {
	return c.Query().Where(chapter.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChapterClient) GetX(ctx context.Context, id string) *Chapter {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVideo queries the video edge of a Chapter.
func (c *ChapterClient) QueryVideo(ch *Chapter) *VideoQuery {
	query := (&VideoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chapter.Table, chapter.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chapter.VideoTable, chapter.VideoColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChapterClient) Hooks() []Hook {
	return c.hooks.Chapter
}

// Interceptors returns the client interceptors.
func (c *ChapterClient) Interceptors() []Interceptor {
	return c.inters.Chapter
}

func (c *ChapterClient) mutate(ctx context.Context, m *ChapterMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChapterCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChapterUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChapterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChapterDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Chapter mutation op: %q", m.Op())
	}
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `comment.Intercept(f(g(h())))`.
func (c *CommentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Comment = append(c.inters.Comment, interceptors...)
}

// Create returns a builder for creating a Comment entity.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id string) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CommentClient) DeleteOneID(id string) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id string) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id string) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVideo queries the video edge of a Comment.
func (c *CommentClient) QueryVideo(co *Comment) *VideoQuery {
	query := (&VideoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.VideoTable, comment.VideoColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// Interceptors returns the client interceptors.
func (c *CommentClient) Interceptors() []Interceptor {
	return c.inters.Comment
}

func (c *CommentClient) mutate(ctx context.Context, m *CommentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CommentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CommentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Comment mutation op: %q", m.Op())
	}
}

// VideoClient is a client for the Video schema.
type VideoClient struct {
	config
}

// NewVideoClient returns a client for the Video from the given config.
func NewVideoClient(c config) *VideoClient {
	return &VideoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `video.Hooks(f(g(h())))`.
func (c *VideoClient) Use(hooks ...Hook) {
	c.hooks.Video = append(c.hooks.Video, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `video.Intercept(f(g(h())))`.
func (c *VideoClient) Intercept(interceptors ...Interceptor) {
	c.inters.Video = append(c.inters.Video, interceptors...)
}

// Create returns a builder for creating a Video entity.
func (c *VideoClient) Create() *VideoCreate {
	mutation := newVideoMutation(c.config, OpCreate)
	return &VideoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Video entities.
func (c *VideoClient) CreateBulk(builders ...*VideoCreate) *VideoCreateBulk {
	return &VideoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Video.
func (c *VideoClient) Update() *VideoUpdate {
	mutation := newVideoMutation(c.config, OpUpdate)
	return &VideoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VideoClient) UpdateOne(v *Video) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideo(v))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VideoClient) UpdateOneID(id string) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideoID(id))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Video.
func (c *VideoClient) Delete() *VideoDelete {
	mutation := newVideoMutation(c.config, OpDelete)
	return &VideoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VideoClient) DeleteOne(v *Video) *VideoDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VideoClient) DeleteOneID(id string) *VideoDeleteOne {
	builder := c.Delete().Where(video.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VideoDeleteOne{builder}
}

// Query returns a query builder for Video.
func (c *VideoClient) Query() *VideoQuery {
	return &VideoQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a Video entity by its id.
func (c *VideoClient) Get(ctx context.Context, id string) (*Video, error) {
	return c.Query().Where(video.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VideoClient) GetX(ctx context.Context, id string) *Video {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChannel queries the channel edge of a Video.
func (c *VideoClient) QueryChannel(v *Video) *ChannelQuery {
	query := (&ChannelClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, video.ChannelTable, video.ChannelColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChapters queries the chapters edge of a Video.
func (c *VideoClient) QueryChapters(v *Video) *ChapterQuery {
	query := (&ChapterClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(chapter.Table, chapter.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, video.ChaptersTable, video.ChaptersColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a Video.
func (c *VideoClient) QueryComments(v *Video) *CommentQuery {
	query := (&CommentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, video.CommentsTable, video.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VideoClient) Hooks() []Hook {
	return c.hooks.Video
}

// Interceptors returns the client interceptors.
func (c *VideoClient) Interceptors() []Interceptor {
	return c.inters.Video
}

func (c *VideoClient) mutate(ctx context.Context, m *VideoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VideoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VideoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VideoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Video mutation op: %q", m.Op())
	}
}
