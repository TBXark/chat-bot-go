// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/TBXark/chat-bot-go/pkg/dao/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chatconfig"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chathistory"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/openaitoken"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ChatConfig is the client for interacting with the ChatConfig builders.
	ChatConfig *ChatConfigClient
	// ChatHistory is the client for interacting with the ChatHistory builders.
	ChatHistory *ChatHistoryClient
	// OpenAIToken is the client for interacting with the OpenAIToken builders.
	OpenAIToken *OpenAITokenClient
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
	c.ChatConfig = NewChatConfigClient(c.config)
	c.ChatHistory = NewChatHistoryClient(c.config)
	c.OpenAIToken = NewOpenAITokenClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		ctx:         ctx,
		config:      cfg,
		ChatConfig:  NewChatConfigClient(cfg),
		ChatHistory: NewChatHistoryClient(cfg),
		OpenAIToken: NewOpenAITokenClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		ChatConfig:  NewChatConfigClient(cfg),
		ChatHistory: NewChatHistoryClient(cfg),
		OpenAIToken: NewOpenAITokenClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ChatConfig.
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
	c.ChatConfig.Use(hooks...)
	c.ChatHistory.Use(hooks...)
	c.OpenAIToken.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.ChatConfig.Intercept(interceptors...)
	c.ChatHistory.Intercept(interceptors...)
	c.OpenAIToken.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ChatConfigMutation:
		return c.ChatConfig.mutate(ctx, m)
	case *ChatHistoryMutation:
		return c.ChatHistory.mutate(ctx, m)
	case *OpenAITokenMutation:
		return c.OpenAIToken.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ChatConfigClient is a client for the ChatConfig schema.
type ChatConfigClient struct {
	config
}

// NewChatConfigClient returns a client for the ChatConfig from the given config.
func NewChatConfigClient(c config) *ChatConfigClient {
	return &ChatConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chatconfig.Hooks(f(g(h())))`.
func (c *ChatConfigClient) Use(hooks ...Hook) {
	c.hooks.ChatConfig = append(c.hooks.ChatConfig, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chatconfig.Intercept(f(g(h())))`.
func (c *ChatConfigClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChatConfig = append(c.inters.ChatConfig, interceptors...)
}

// Create returns a builder for creating a ChatConfig entity.
func (c *ChatConfigClient) Create() *ChatConfigCreate {
	mutation := newChatConfigMutation(c.config, OpCreate)
	return &ChatConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChatConfig entities.
func (c *ChatConfigClient) CreateBulk(builders ...*ChatConfigCreate) *ChatConfigCreateBulk {
	return &ChatConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChatConfig.
func (c *ChatConfigClient) Update() *ChatConfigUpdate {
	mutation := newChatConfigMutation(c.config, OpUpdate)
	return &ChatConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatConfigClient) UpdateOne(cc *ChatConfig) *ChatConfigUpdateOne {
	mutation := newChatConfigMutation(c.config, OpUpdateOne, withChatConfig(cc))
	return &ChatConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatConfigClient) UpdateOneID(id int) *ChatConfigUpdateOne {
	mutation := newChatConfigMutation(c.config, OpUpdateOne, withChatConfigID(id))
	return &ChatConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChatConfig.
func (c *ChatConfigClient) Delete() *ChatConfigDelete {
	mutation := newChatConfigMutation(c.config, OpDelete)
	return &ChatConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChatConfigClient) DeleteOne(cc *ChatConfig) *ChatConfigDeleteOne {
	return c.DeleteOneID(cc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChatConfigClient) DeleteOneID(id int) *ChatConfigDeleteOne {
	builder := c.Delete().Where(chatconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChatConfigDeleteOne{builder}
}

// Query returns a query builder for ChatConfig.
func (c *ChatConfigClient) Query() *ChatConfigQuery {
	return &ChatConfigQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChatConfig},
		inters: c.Interceptors(),
	}
}

// Get returns a ChatConfig entity by its id.
func (c *ChatConfigClient) Get(ctx context.Context, id int) (*ChatConfig, error) {
	return c.Query().Where(chatconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatConfigClient) GetX(ctx context.Context, id int) *ChatConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChatConfigClient) Hooks() []Hook {
	return c.hooks.ChatConfig
}

// Interceptors returns the client interceptors.
func (c *ChatConfigClient) Interceptors() []Interceptor {
	return c.inters.ChatConfig
}

func (c *ChatConfigClient) mutate(ctx context.Context, m *ChatConfigMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChatConfigCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChatConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChatConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChatConfigDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChatConfig mutation op: %q", m.Op())
	}
}

// ChatHistoryClient is a client for the ChatHistory schema.
type ChatHistoryClient struct {
	config
}

// NewChatHistoryClient returns a client for the ChatHistory from the given config.
func NewChatHistoryClient(c config) *ChatHistoryClient {
	return &ChatHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chathistory.Hooks(f(g(h())))`.
func (c *ChatHistoryClient) Use(hooks ...Hook) {
	c.hooks.ChatHistory = append(c.hooks.ChatHistory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chathistory.Intercept(f(g(h())))`.
func (c *ChatHistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChatHistory = append(c.inters.ChatHistory, interceptors...)
}

// Create returns a builder for creating a ChatHistory entity.
func (c *ChatHistoryClient) Create() *ChatHistoryCreate {
	mutation := newChatHistoryMutation(c.config, OpCreate)
	return &ChatHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChatHistory entities.
func (c *ChatHistoryClient) CreateBulk(builders ...*ChatHistoryCreate) *ChatHistoryCreateBulk {
	return &ChatHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChatHistory.
func (c *ChatHistoryClient) Update() *ChatHistoryUpdate {
	mutation := newChatHistoryMutation(c.config, OpUpdate)
	return &ChatHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatHistoryClient) UpdateOne(ch *ChatHistory) *ChatHistoryUpdateOne {
	mutation := newChatHistoryMutation(c.config, OpUpdateOne, withChatHistory(ch))
	return &ChatHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatHistoryClient) UpdateOneID(id int) *ChatHistoryUpdateOne {
	mutation := newChatHistoryMutation(c.config, OpUpdateOne, withChatHistoryID(id))
	return &ChatHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChatHistory.
func (c *ChatHistoryClient) Delete() *ChatHistoryDelete {
	mutation := newChatHistoryMutation(c.config, OpDelete)
	return &ChatHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChatHistoryClient) DeleteOne(ch *ChatHistory) *ChatHistoryDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChatHistoryClient) DeleteOneID(id int) *ChatHistoryDeleteOne {
	builder := c.Delete().Where(chathistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChatHistoryDeleteOne{builder}
}

// Query returns a query builder for ChatHistory.
func (c *ChatHistoryClient) Query() *ChatHistoryQuery {
	return &ChatHistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChatHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a ChatHistory entity by its id.
func (c *ChatHistoryClient) Get(ctx context.Context, id int) (*ChatHistory, error) {
	return c.Query().Where(chathistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatHistoryClient) GetX(ctx context.Context, id int) *ChatHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChatHistoryClient) Hooks() []Hook {
	return c.hooks.ChatHistory
}

// Interceptors returns the client interceptors.
func (c *ChatHistoryClient) Interceptors() []Interceptor {
	return c.inters.ChatHistory
}

func (c *ChatHistoryClient) mutate(ctx context.Context, m *ChatHistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChatHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChatHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChatHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChatHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChatHistory mutation op: %q", m.Op())
	}
}

// OpenAITokenClient is a client for the OpenAIToken schema.
type OpenAITokenClient struct {
	config
}

// NewOpenAITokenClient returns a client for the OpenAIToken from the given config.
func NewOpenAITokenClient(c config) *OpenAITokenClient {
	return &OpenAITokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `openaitoken.Hooks(f(g(h())))`.
func (c *OpenAITokenClient) Use(hooks ...Hook) {
	c.hooks.OpenAIToken = append(c.hooks.OpenAIToken, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `openaitoken.Intercept(f(g(h())))`.
func (c *OpenAITokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.OpenAIToken = append(c.inters.OpenAIToken, interceptors...)
}

// Create returns a builder for creating a OpenAIToken entity.
func (c *OpenAITokenClient) Create() *OpenAITokenCreate {
	mutation := newOpenAITokenMutation(c.config, OpCreate)
	return &OpenAITokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OpenAIToken entities.
func (c *OpenAITokenClient) CreateBulk(builders ...*OpenAITokenCreate) *OpenAITokenCreateBulk {
	return &OpenAITokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OpenAIToken.
func (c *OpenAITokenClient) Update() *OpenAITokenUpdate {
	mutation := newOpenAITokenMutation(c.config, OpUpdate)
	return &OpenAITokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OpenAITokenClient) UpdateOne(oat *OpenAIToken) *OpenAITokenUpdateOne {
	mutation := newOpenAITokenMutation(c.config, OpUpdateOne, withOpenAIToken(oat))
	return &OpenAITokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OpenAITokenClient) UpdateOneID(id int) *OpenAITokenUpdateOne {
	mutation := newOpenAITokenMutation(c.config, OpUpdateOne, withOpenAITokenID(id))
	return &OpenAITokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OpenAIToken.
func (c *OpenAITokenClient) Delete() *OpenAITokenDelete {
	mutation := newOpenAITokenMutation(c.config, OpDelete)
	return &OpenAITokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OpenAITokenClient) DeleteOne(oat *OpenAIToken) *OpenAITokenDeleteOne {
	return c.DeleteOneID(oat.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OpenAITokenClient) DeleteOneID(id int) *OpenAITokenDeleteOne {
	builder := c.Delete().Where(openaitoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OpenAITokenDeleteOne{builder}
}

// Query returns a query builder for OpenAIToken.
func (c *OpenAITokenClient) Query() *OpenAITokenQuery {
	return &OpenAITokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOpenAIToken},
		inters: c.Interceptors(),
	}
}

// Get returns a OpenAIToken entity by its id.
func (c *OpenAITokenClient) Get(ctx context.Context, id int) (*OpenAIToken, error) {
	return c.Query().Where(openaitoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OpenAITokenClient) GetX(ctx context.Context, id int) *OpenAIToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OpenAITokenClient) Hooks() []Hook {
	return c.hooks.OpenAIToken
}

// Interceptors returns the client interceptors.
func (c *OpenAITokenClient) Interceptors() []Interceptor {
	return c.inters.OpenAIToken
}

func (c *OpenAITokenClient) mutate(ctx context.Context, m *OpenAITokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OpenAITokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OpenAITokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OpenAITokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OpenAITokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown OpenAIToken mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ChatConfig, ChatHistory, OpenAIToken []ent.Hook
	}
	inters struct {
		ChatConfig, ChatHistory, OpenAIToken []ent.Interceptor
	}
)
