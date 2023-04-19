// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chathistory"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/predicate"
)

// ChatHistoryUpdate is the builder for updating ChatHistory entities.
type ChatHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *ChatHistoryMutation
}

// Where appends a list predicates to the ChatHistoryUpdate builder.
func (chu *ChatHistoryUpdate) Where(ps ...predicate.ChatHistory) *ChatHistoryUpdate {
	chu.mutation.Where(ps...)
	return chu
}

// SetChatID sets the "chat_id" field.
func (chu *ChatHistoryUpdate) SetChatID(i int64) *ChatHistoryUpdate {
	chu.mutation.ResetChatID()
	chu.mutation.SetChatID(i)
	return chu
}

// AddChatID adds i to the "chat_id" field.
func (chu *ChatHistoryUpdate) AddChatID(i int64) *ChatHistoryUpdate {
	chu.mutation.AddChatID(i)
	return chu
}

// SetContent sets the "content" field.
func (chu *ChatHistoryUpdate) SetContent(s string) *ChatHistoryUpdate {
	chu.mutation.SetContent(s)
	return chu
}

// Mutation returns the ChatHistoryMutation object of the builder.
func (chu *ChatHistoryUpdate) Mutation() *ChatHistoryMutation {
	return chu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (chu *ChatHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ChatHistoryMutation](ctx, chu.sqlSave, chu.mutation, chu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (chu *ChatHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := chu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (chu *ChatHistoryUpdate) Exec(ctx context.Context) error {
	_, err := chu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chu *ChatHistoryUpdate) ExecX(ctx context.Context) {
	if err := chu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (chu *ChatHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(chathistory.Table, chathistory.Columns, sqlgraph.NewFieldSpec(chathistory.FieldID, field.TypeInt))
	if ps := chu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := chu.mutation.ChatID(); ok {
		_spec.SetField(chathistory.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := chu.mutation.AddedChatID(); ok {
		_spec.AddField(chathistory.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := chu.mutation.Content(); ok {
		_spec.SetField(chathistory.FieldContent, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, chu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chathistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	chu.mutation.done = true
	return n, nil
}

// ChatHistoryUpdateOne is the builder for updating a single ChatHistory entity.
type ChatHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatHistoryMutation
}

// SetChatID sets the "chat_id" field.
func (chuo *ChatHistoryUpdateOne) SetChatID(i int64) *ChatHistoryUpdateOne {
	chuo.mutation.ResetChatID()
	chuo.mutation.SetChatID(i)
	return chuo
}

// AddChatID adds i to the "chat_id" field.
func (chuo *ChatHistoryUpdateOne) AddChatID(i int64) *ChatHistoryUpdateOne {
	chuo.mutation.AddChatID(i)
	return chuo
}

// SetContent sets the "content" field.
func (chuo *ChatHistoryUpdateOne) SetContent(s string) *ChatHistoryUpdateOne {
	chuo.mutation.SetContent(s)
	return chuo
}

// Mutation returns the ChatHistoryMutation object of the builder.
func (chuo *ChatHistoryUpdateOne) Mutation() *ChatHistoryMutation {
	return chuo.mutation
}

// Where appends a list predicates to the ChatHistoryUpdate builder.
func (chuo *ChatHistoryUpdateOne) Where(ps ...predicate.ChatHistory) *ChatHistoryUpdateOne {
	chuo.mutation.Where(ps...)
	return chuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (chuo *ChatHistoryUpdateOne) Select(field string, fields ...string) *ChatHistoryUpdateOne {
	chuo.fields = append([]string{field}, fields...)
	return chuo
}

// Save executes the query and returns the updated ChatHistory entity.
func (chuo *ChatHistoryUpdateOne) Save(ctx context.Context) (*ChatHistory, error) {
	return withHooks[*ChatHistory, ChatHistoryMutation](ctx, chuo.sqlSave, chuo.mutation, chuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (chuo *ChatHistoryUpdateOne) SaveX(ctx context.Context) *ChatHistory {
	node, err := chuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (chuo *ChatHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := chuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chuo *ChatHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := chuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (chuo *ChatHistoryUpdateOne) sqlSave(ctx context.Context) (_node *ChatHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(chathistory.Table, chathistory.Columns, sqlgraph.NewFieldSpec(chathistory.FieldID, field.TypeInt))
	id, ok := chuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChatHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := chuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chathistory.FieldID)
		for _, f := range fields {
			if !chathistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chathistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := chuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := chuo.mutation.ChatID(); ok {
		_spec.SetField(chathistory.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := chuo.mutation.AddedChatID(); ok {
		_spec.AddField(chathistory.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := chuo.mutation.Content(); ok {
		_spec.SetField(chathistory.FieldContent, field.TypeString, value)
	}
	_node = &ChatHistory{config: chuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, chuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chathistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	chuo.mutation.done = true
	return _node, nil
}