// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/openaitoken"
)

// OpenAITokenCreate is the builder for creating a OpenAIToken entity.
type OpenAITokenCreate struct {
	config
	mutation *OpenAITokenMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (oatc *OpenAITokenCreate) SetToken(s string) *OpenAITokenCreate {
	oatc.mutation.SetToken(s)
	return oatc
}

// SetIsActive sets the "is_active" field.
func (oatc *OpenAITokenCreate) SetIsActive(b bool) *OpenAITokenCreate {
	oatc.mutation.SetIsActive(b)
	return oatc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (oatc *OpenAITokenCreate) SetNillableIsActive(b *bool) *OpenAITokenCreate {
	if b != nil {
		oatc.SetIsActive(*b)
	}
	return oatc
}

// Mutation returns the OpenAITokenMutation object of the builder.
func (oatc *OpenAITokenCreate) Mutation() *OpenAITokenMutation {
	return oatc.mutation
}

// Save creates the OpenAIToken in the database.
func (oatc *OpenAITokenCreate) Save(ctx context.Context) (*OpenAIToken, error) {
	oatc.defaults()
	return withHooks[*OpenAIToken, OpenAITokenMutation](ctx, oatc.sqlSave, oatc.mutation, oatc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oatc *OpenAITokenCreate) SaveX(ctx context.Context) *OpenAIToken {
	v, err := oatc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oatc *OpenAITokenCreate) Exec(ctx context.Context) error {
	_, err := oatc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oatc *OpenAITokenCreate) ExecX(ctx context.Context) {
	if err := oatc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oatc *OpenAITokenCreate) defaults() {
	if _, ok := oatc.mutation.IsActive(); !ok {
		v := openaitoken.DefaultIsActive
		oatc.mutation.SetIsActive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oatc *OpenAITokenCreate) check() error {
	if _, ok := oatc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "OpenAIToken.token"`)}
	}
	if v, ok := oatc.mutation.Token(); ok {
		if err := openaitoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "OpenAIToken.token": %w`, err)}
		}
	}
	if _, ok := oatc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "OpenAIToken.is_active"`)}
	}
	return nil
}

func (oatc *OpenAITokenCreate) sqlSave(ctx context.Context) (*OpenAIToken, error) {
	if err := oatc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oatc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oatc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oatc.mutation.id = &_node.ID
	oatc.mutation.done = true
	return _node, nil
}

func (oatc *OpenAITokenCreate) createSpec() (*OpenAIToken, *sqlgraph.CreateSpec) {
	var (
		_node = &OpenAIToken{config: oatc.config}
		_spec = sqlgraph.NewCreateSpec(openaitoken.Table, sqlgraph.NewFieldSpec(openaitoken.FieldID, field.TypeInt))
	)
	if value, ok := oatc.mutation.Token(); ok {
		_spec.SetField(openaitoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := oatc.mutation.IsActive(); ok {
		_spec.SetField(openaitoken.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	return _node, _spec
}

// OpenAITokenCreateBulk is the builder for creating many OpenAIToken entities in bulk.
type OpenAITokenCreateBulk struct {
	config
	builders []*OpenAITokenCreate
}

// Save creates the OpenAIToken entities in the database.
func (oatcb *OpenAITokenCreateBulk) Save(ctx context.Context) ([]*OpenAIToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oatcb.builders))
	nodes := make([]*OpenAIToken, len(oatcb.builders))
	mutators := make([]Mutator, len(oatcb.builders))
	for i := range oatcb.builders {
		func(i int, root context.Context) {
			builder := oatcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OpenAITokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oatcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oatcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oatcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oatcb *OpenAITokenCreateBulk) SaveX(ctx context.Context) []*OpenAIToken {
	v, err := oatcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oatcb *OpenAITokenCreateBulk) Exec(ctx context.Context) error {
	_, err := oatcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oatcb *OpenAITokenCreateBulk) ExecX(ctx context.Context) {
	if err := oatcb.Exec(ctx); err != nil {
		panic(err)
	}
}
