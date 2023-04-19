// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chatconfig"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chathistory"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/openaitoken"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeChatConfig  = "ChatConfig"
	TypeChatHistory = "ChatHistory"
	TypeOpenAIToken = "OpenAIToken"
)

// ChatConfigMutation represents an operation that mutates the ChatConfig nodes in the graph.
type ChatConfigMutation struct {
	config
	op            Op
	typ           string
	id            *int
	chat_id       *int64
	addchat_id    *int64
	json          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ChatConfig, error)
	predicates    []predicate.ChatConfig
}

var _ ent.Mutation = (*ChatConfigMutation)(nil)

// chatconfigOption allows management of the mutation configuration using functional options.
type chatconfigOption func(*ChatConfigMutation)

// newChatConfigMutation creates new mutation for the ChatConfig entity.
func newChatConfigMutation(c config, op Op, opts ...chatconfigOption) *ChatConfigMutation {
	m := &ChatConfigMutation{
		config:        c,
		op:            op,
		typ:           TypeChatConfig,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withChatConfigID sets the ID field of the mutation.
func withChatConfigID(id int) chatconfigOption {
	return func(m *ChatConfigMutation) {
		var (
			err   error
			once  sync.Once
			value *ChatConfig
		)
		m.oldValue = func(ctx context.Context) (*ChatConfig, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ChatConfig.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withChatConfig sets the old ChatConfig of the mutation.
func withChatConfig(node *ChatConfig) chatconfigOption {
	return func(m *ChatConfigMutation) {
		m.oldValue = func(context.Context) (*ChatConfig, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ChatConfigMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ChatConfigMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ChatConfigMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ChatConfigMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ChatConfig.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetChatID sets the "chat_id" field.
func (m *ChatConfigMutation) SetChatID(i int64) {
	m.chat_id = &i
	m.addchat_id = nil
}

// ChatID returns the value of the "chat_id" field in the mutation.
func (m *ChatConfigMutation) ChatID() (r int64, exists bool) {
	v := m.chat_id
	if v == nil {
		return
	}
	return *v, true
}

// OldChatID returns the old "chat_id" field's value of the ChatConfig entity.
// If the ChatConfig object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatConfigMutation) OldChatID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldChatID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldChatID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldChatID: %w", err)
	}
	return oldValue.ChatID, nil
}

// AddChatID adds i to the "chat_id" field.
func (m *ChatConfigMutation) AddChatID(i int64) {
	if m.addchat_id != nil {
		*m.addchat_id += i
	} else {
		m.addchat_id = &i
	}
}

// AddedChatID returns the value that was added to the "chat_id" field in this mutation.
func (m *ChatConfigMutation) AddedChatID() (r int64, exists bool) {
	v := m.addchat_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetChatID resets all changes to the "chat_id" field.
func (m *ChatConfigMutation) ResetChatID() {
	m.chat_id = nil
	m.addchat_id = nil
}

// SetJSON sets the "json" field.
func (m *ChatConfigMutation) SetJSON(s string) {
	m.json = &s
}

// JSON returns the value of the "json" field in the mutation.
func (m *ChatConfigMutation) JSON() (r string, exists bool) {
	v := m.json
	if v == nil {
		return
	}
	return *v, true
}

// OldJSON returns the old "json" field's value of the ChatConfig entity.
// If the ChatConfig object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatConfigMutation) OldJSON(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldJSON is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldJSON requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldJSON: %w", err)
	}
	return oldValue.JSON, nil
}

// ResetJSON resets all changes to the "json" field.
func (m *ChatConfigMutation) ResetJSON() {
	m.json = nil
}

// Where appends a list predicates to the ChatConfigMutation builder.
func (m *ChatConfigMutation) Where(ps ...predicate.ChatConfig) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ChatConfigMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ChatConfigMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ChatConfig, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ChatConfigMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ChatConfigMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ChatConfig).
func (m *ChatConfigMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ChatConfigMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.chat_id != nil {
		fields = append(fields, chatconfig.FieldChatID)
	}
	if m.json != nil {
		fields = append(fields, chatconfig.FieldJSON)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ChatConfigMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case chatconfig.FieldChatID:
		return m.ChatID()
	case chatconfig.FieldJSON:
		return m.JSON()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ChatConfigMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case chatconfig.FieldChatID:
		return m.OldChatID(ctx)
	case chatconfig.FieldJSON:
		return m.OldJSON(ctx)
	}
	return nil, fmt.Errorf("unknown ChatConfig field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatConfigMutation) SetField(name string, value ent.Value) error {
	switch name {
	case chatconfig.FieldChatID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetChatID(v)
		return nil
	case chatconfig.FieldJSON:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetJSON(v)
		return nil
	}
	return fmt.Errorf("unknown ChatConfig field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ChatConfigMutation) AddedFields() []string {
	var fields []string
	if m.addchat_id != nil {
		fields = append(fields, chatconfig.FieldChatID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ChatConfigMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case chatconfig.FieldChatID:
		return m.AddedChatID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatConfigMutation) AddField(name string, value ent.Value) error {
	switch name {
	case chatconfig.FieldChatID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddChatID(v)
		return nil
	}
	return fmt.Errorf("unknown ChatConfig numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ChatConfigMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ChatConfigMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ChatConfigMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ChatConfig nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ChatConfigMutation) ResetField(name string) error {
	switch name {
	case chatconfig.FieldChatID:
		m.ResetChatID()
		return nil
	case chatconfig.FieldJSON:
		m.ResetJSON()
		return nil
	}
	return fmt.Errorf("unknown ChatConfig field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ChatConfigMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ChatConfigMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ChatConfigMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ChatConfigMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ChatConfigMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ChatConfigMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ChatConfigMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ChatConfig unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ChatConfigMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ChatConfig edge %s", name)
}

// ChatHistoryMutation represents an operation that mutates the ChatHistory nodes in the graph.
type ChatHistoryMutation struct {
	config
	op            Op
	typ           string
	id            *int
	chat_id       *int64
	addchat_id    *int64
	content       *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ChatHistory, error)
	predicates    []predicate.ChatHistory
}

var _ ent.Mutation = (*ChatHistoryMutation)(nil)

// chathistoryOption allows management of the mutation configuration using functional options.
type chathistoryOption func(*ChatHistoryMutation)

// newChatHistoryMutation creates new mutation for the ChatHistory entity.
func newChatHistoryMutation(c config, op Op, opts ...chathistoryOption) *ChatHistoryMutation {
	m := &ChatHistoryMutation{
		config:        c,
		op:            op,
		typ:           TypeChatHistory,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withChatHistoryID sets the ID field of the mutation.
func withChatHistoryID(id int) chathistoryOption {
	return func(m *ChatHistoryMutation) {
		var (
			err   error
			once  sync.Once
			value *ChatHistory
		)
		m.oldValue = func(ctx context.Context) (*ChatHistory, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ChatHistory.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withChatHistory sets the old ChatHistory of the mutation.
func withChatHistory(node *ChatHistory) chathistoryOption {
	return func(m *ChatHistoryMutation) {
		m.oldValue = func(context.Context) (*ChatHistory, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ChatHistoryMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ChatHistoryMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ChatHistoryMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ChatHistoryMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ChatHistory.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetChatID sets the "chat_id" field.
func (m *ChatHistoryMutation) SetChatID(i int64) {
	m.chat_id = &i
	m.addchat_id = nil
}

// ChatID returns the value of the "chat_id" field in the mutation.
func (m *ChatHistoryMutation) ChatID() (r int64, exists bool) {
	v := m.chat_id
	if v == nil {
		return
	}
	return *v, true
}

// OldChatID returns the old "chat_id" field's value of the ChatHistory entity.
// If the ChatHistory object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatHistoryMutation) OldChatID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldChatID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldChatID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldChatID: %w", err)
	}
	return oldValue.ChatID, nil
}

// AddChatID adds i to the "chat_id" field.
func (m *ChatHistoryMutation) AddChatID(i int64) {
	if m.addchat_id != nil {
		*m.addchat_id += i
	} else {
		m.addchat_id = &i
	}
}

// AddedChatID returns the value that was added to the "chat_id" field in this mutation.
func (m *ChatHistoryMutation) AddedChatID() (r int64, exists bool) {
	v := m.addchat_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetChatID resets all changes to the "chat_id" field.
func (m *ChatHistoryMutation) ResetChatID() {
	m.chat_id = nil
	m.addchat_id = nil
}

// SetContent sets the "content" field.
func (m *ChatHistoryMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *ChatHistoryMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the ChatHistory entity.
// If the ChatHistory object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatHistoryMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *ChatHistoryMutation) ResetContent() {
	m.content = nil
}

// Where appends a list predicates to the ChatHistoryMutation builder.
func (m *ChatHistoryMutation) Where(ps ...predicate.ChatHistory) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ChatHistoryMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ChatHistoryMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ChatHistory, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ChatHistoryMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ChatHistoryMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ChatHistory).
func (m *ChatHistoryMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ChatHistoryMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.chat_id != nil {
		fields = append(fields, chathistory.FieldChatID)
	}
	if m.content != nil {
		fields = append(fields, chathistory.FieldContent)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ChatHistoryMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case chathistory.FieldChatID:
		return m.ChatID()
	case chathistory.FieldContent:
		return m.Content()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ChatHistoryMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case chathistory.FieldChatID:
		return m.OldChatID(ctx)
	case chathistory.FieldContent:
		return m.OldContent(ctx)
	}
	return nil, fmt.Errorf("unknown ChatHistory field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatHistoryMutation) SetField(name string, value ent.Value) error {
	switch name {
	case chathistory.FieldChatID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetChatID(v)
		return nil
	case chathistory.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	}
	return fmt.Errorf("unknown ChatHistory field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ChatHistoryMutation) AddedFields() []string {
	var fields []string
	if m.addchat_id != nil {
		fields = append(fields, chathistory.FieldChatID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ChatHistoryMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case chathistory.FieldChatID:
		return m.AddedChatID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatHistoryMutation) AddField(name string, value ent.Value) error {
	switch name {
	case chathistory.FieldChatID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddChatID(v)
		return nil
	}
	return fmt.Errorf("unknown ChatHistory numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ChatHistoryMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ChatHistoryMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ChatHistoryMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ChatHistory nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ChatHistoryMutation) ResetField(name string) error {
	switch name {
	case chathistory.FieldChatID:
		m.ResetChatID()
		return nil
	case chathistory.FieldContent:
		m.ResetContent()
		return nil
	}
	return fmt.Errorf("unknown ChatHistory field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ChatHistoryMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ChatHistoryMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ChatHistoryMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ChatHistoryMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ChatHistoryMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ChatHistoryMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ChatHistoryMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ChatHistory unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ChatHistoryMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ChatHistory edge %s", name)
}

// OpenAITokenMutation represents an operation that mutates the OpenAIToken nodes in the graph.
type OpenAITokenMutation struct {
	config
	op            Op
	typ           string
	id            *int
	token         *string
	is_active     *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*OpenAIToken, error)
	predicates    []predicate.OpenAIToken
}

var _ ent.Mutation = (*OpenAITokenMutation)(nil)

// openaitokenOption allows management of the mutation configuration using functional options.
type openaitokenOption func(*OpenAITokenMutation)

// newOpenAITokenMutation creates new mutation for the OpenAIToken entity.
func newOpenAITokenMutation(c config, op Op, opts ...openaitokenOption) *OpenAITokenMutation {
	m := &OpenAITokenMutation{
		config:        c,
		op:            op,
		typ:           TypeOpenAIToken,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOpenAITokenID sets the ID field of the mutation.
func withOpenAITokenID(id int) openaitokenOption {
	return func(m *OpenAITokenMutation) {
		var (
			err   error
			once  sync.Once
			value *OpenAIToken
		)
		m.oldValue = func(ctx context.Context) (*OpenAIToken, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().OpenAIToken.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOpenAIToken sets the old OpenAIToken of the mutation.
func withOpenAIToken(node *OpenAIToken) openaitokenOption {
	return func(m *OpenAITokenMutation) {
		m.oldValue = func(context.Context) (*OpenAIToken, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OpenAITokenMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OpenAITokenMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OpenAITokenMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *OpenAITokenMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().OpenAIToken.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetToken sets the "token" field.
func (m *OpenAITokenMutation) SetToken(s string) {
	m.token = &s
}

// Token returns the value of the "token" field in the mutation.
func (m *OpenAITokenMutation) Token() (r string, exists bool) {
	v := m.token
	if v == nil {
		return
	}
	return *v, true
}

// OldToken returns the old "token" field's value of the OpenAIToken entity.
// If the OpenAIToken object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OpenAITokenMutation) OldToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldToken: %w", err)
	}
	return oldValue.Token, nil
}

// ResetToken resets all changes to the "token" field.
func (m *OpenAITokenMutation) ResetToken() {
	m.token = nil
}

// SetIsActive sets the "is_active" field.
func (m *OpenAITokenMutation) SetIsActive(b bool) {
	m.is_active = &b
}

// IsActive returns the value of the "is_active" field in the mutation.
func (m *OpenAITokenMutation) IsActive() (r bool, exists bool) {
	v := m.is_active
	if v == nil {
		return
	}
	return *v, true
}

// OldIsActive returns the old "is_active" field's value of the OpenAIToken entity.
// If the OpenAIToken object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OpenAITokenMutation) OldIsActive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsActive is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsActive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsActive: %w", err)
	}
	return oldValue.IsActive, nil
}

// ResetIsActive resets all changes to the "is_active" field.
func (m *OpenAITokenMutation) ResetIsActive() {
	m.is_active = nil
}

// Where appends a list predicates to the OpenAITokenMutation builder.
func (m *OpenAITokenMutation) Where(ps ...predicate.OpenAIToken) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the OpenAITokenMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *OpenAITokenMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.OpenAIToken, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *OpenAITokenMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *OpenAITokenMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (OpenAIToken).
func (m *OpenAITokenMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OpenAITokenMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.token != nil {
		fields = append(fields, openaitoken.FieldToken)
	}
	if m.is_active != nil {
		fields = append(fields, openaitoken.FieldIsActive)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OpenAITokenMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case openaitoken.FieldToken:
		return m.Token()
	case openaitoken.FieldIsActive:
		return m.IsActive()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OpenAITokenMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case openaitoken.FieldToken:
		return m.OldToken(ctx)
	case openaitoken.FieldIsActive:
		return m.OldIsActive(ctx)
	}
	return nil, fmt.Errorf("unknown OpenAIToken field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OpenAITokenMutation) SetField(name string, value ent.Value) error {
	switch name {
	case openaitoken.FieldToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetToken(v)
		return nil
	case openaitoken.FieldIsActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsActive(v)
		return nil
	}
	return fmt.Errorf("unknown OpenAIToken field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OpenAITokenMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OpenAITokenMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OpenAITokenMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown OpenAIToken numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OpenAITokenMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OpenAITokenMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OpenAITokenMutation) ClearField(name string) error {
	return fmt.Errorf("unknown OpenAIToken nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OpenAITokenMutation) ResetField(name string) error {
	switch name {
	case openaitoken.FieldToken:
		m.ResetToken()
		return nil
	case openaitoken.FieldIsActive:
		m.ResetIsActive()
		return nil
	}
	return fmt.Errorf("unknown OpenAIToken field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OpenAITokenMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OpenAITokenMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OpenAITokenMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OpenAITokenMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OpenAITokenMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OpenAITokenMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OpenAITokenMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown OpenAIToken unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OpenAITokenMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown OpenAIToken edge %s", name)
}