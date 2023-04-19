// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/openaitoken"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/predicate"
)

// OpenAITokenUpdate is the builder for updating OpenAIToken entities.
type OpenAITokenUpdate struct {
	config
	hooks    []Hook
	mutation *OpenAITokenMutation
}

// Where appends a list predicates to the OpenAITokenUpdate builder.
func (oatu *OpenAITokenUpdate) Where(ps ...predicate.OpenAIToken) *OpenAITokenUpdate {
	oatu.mutation.Where(ps...)
	return oatu
}

// SetToken sets the "token" field.
func (oatu *OpenAITokenUpdate) SetToken(s string) *OpenAITokenUpdate {
	oatu.mutation.SetToken(s)
	return oatu
}

// SetIsActive sets the "is_active" field.
func (oatu *OpenAITokenUpdate) SetIsActive(b bool) *OpenAITokenUpdate {
	oatu.mutation.SetIsActive(b)
	return oatu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (oatu *OpenAITokenUpdate) SetNillableIsActive(b *bool) *OpenAITokenUpdate {
	if b != nil {
		oatu.SetIsActive(*b)
	}
	return oatu
}

// Mutation returns the OpenAITokenMutation object of the builder.
func (oatu *OpenAITokenUpdate) Mutation() *OpenAITokenMutation {
	return oatu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oatu *OpenAITokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OpenAITokenMutation](ctx, oatu.sqlSave, oatu.mutation, oatu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oatu *OpenAITokenUpdate) SaveX(ctx context.Context) int {
	affected, err := oatu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oatu *OpenAITokenUpdate) Exec(ctx context.Context) error {
	_, err := oatu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oatu *OpenAITokenUpdate) ExecX(ctx context.Context) {
	if err := oatu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oatu *OpenAITokenUpdate) check() error {
	if v, ok := oatu.mutation.Token(); ok {
		if err := openaitoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "OpenAIToken.token": %w`, err)}
		}
	}
	return nil
}

func (oatu *OpenAITokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oatu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(openaitoken.Table, openaitoken.Columns, sqlgraph.NewFieldSpec(openaitoken.FieldID, field.TypeInt))
	if ps := oatu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oatu.mutation.Token(); ok {
		_spec.SetField(openaitoken.FieldToken, field.TypeString, value)
	}
	if value, ok := oatu.mutation.IsActive(); ok {
		_spec.SetField(openaitoken.FieldIsActive, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oatu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{openaitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oatu.mutation.done = true
	return n, nil
}

// OpenAITokenUpdateOne is the builder for updating a single OpenAIToken entity.
type OpenAITokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OpenAITokenMutation
}

// SetToken sets the "token" field.
func (oatuo *OpenAITokenUpdateOne) SetToken(s string) *OpenAITokenUpdateOne {
	oatuo.mutation.SetToken(s)
	return oatuo
}

// SetIsActive sets the "is_active" field.
func (oatuo *OpenAITokenUpdateOne) SetIsActive(b bool) *OpenAITokenUpdateOne {
	oatuo.mutation.SetIsActive(b)
	return oatuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (oatuo *OpenAITokenUpdateOne) SetNillableIsActive(b *bool) *OpenAITokenUpdateOne {
	if b != nil {
		oatuo.SetIsActive(*b)
	}
	return oatuo
}

// Mutation returns the OpenAITokenMutation object of the builder.
func (oatuo *OpenAITokenUpdateOne) Mutation() *OpenAITokenMutation {
	return oatuo.mutation
}

// Where appends a list predicates to the OpenAITokenUpdate builder.
func (oatuo *OpenAITokenUpdateOne) Where(ps ...predicate.OpenAIToken) *OpenAITokenUpdateOne {
	oatuo.mutation.Where(ps...)
	return oatuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oatuo *OpenAITokenUpdateOne) Select(field string, fields ...string) *OpenAITokenUpdateOne {
	oatuo.fields = append([]string{field}, fields...)
	return oatuo
}

// Save executes the query and returns the updated OpenAIToken entity.
func (oatuo *OpenAITokenUpdateOne) Save(ctx context.Context) (*OpenAIToken, error) {
	return withHooks[*OpenAIToken, OpenAITokenMutation](ctx, oatuo.sqlSave, oatuo.mutation, oatuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oatuo *OpenAITokenUpdateOne) SaveX(ctx context.Context) *OpenAIToken {
	node, err := oatuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oatuo *OpenAITokenUpdateOne) Exec(ctx context.Context) error {
	_, err := oatuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oatuo *OpenAITokenUpdateOne) ExecX(ctx context.Context) {
	if err := oatuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oatuo *OpenAITokenUpdateOne) check() error {
	if v, ok := oatuo.mutation.Token(); ok {
		if err := openaitoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "OpenAIToken.token": %w`, err)}
		}
	}
	return nil
}

func (oatuo *OpenAITokenUpdateOne) sqlSave(ctx context.Context) (_node *OpenAIToken, err error) {
	if err := oatuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(openaitoken.Table, openaitoken.Columns, sqlgraph.NewFieldSpec(openaitoken.FieldID, field.TypeInt))
	id, ok := oatuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OpenAIToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oatuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, openaitoken.FieldID)
		for _, f := range fields {
			if !openaitoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != openaitoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oatuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oatuo.mutation.Token(); ok {
		_spec.SetField(openaitoken.FieldToken, field.TypeString, value)
	}
	if value, ok := oatuo.mutation.IsActive(); ok {
		_spec.SetField(openaitoken.FieldIsActive, field.TypeBool, value)
	}
	_node = &OpenAIToken{config: oatuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oatuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{openaitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oatuo.mutation.done = true
	return _node, nil
}
