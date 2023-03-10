// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/meta"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MetaCreate is the builder for creating a Meta entity.
type MetaCreate struct {
	config
	mutation *MetaMutation
	hooks    []Hook
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MetaCreate) SetUpdatedAt(t time.Time) *MetaCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MetaCreate) SetNillableUpdatedAt(t *time.Time) *MetaCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *MetaCreate) SetCreatedAt(t time.Time) *MetaCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MetaCreate) SetNillableCreatedAt(t *time.Time) *MetaCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MetaCreate) SetID(u uuid.UUID) *MetaCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MetaCreate) SetNillableID(u *uuid.UUID) *MetaCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// Mutation returns the MetaMutation object of the builder.
func (mc *MetaCreate) Mutation() *MetaMutation {
	return mc.mutation
}

// Save creates the Meta in the database.
func (mc *MetaCreate) Save(ctx context.Context) (*Meta, error) {
	mc.defaults()
	return withHooks[*Meta, MetaMutation](ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetaCreate) SaveX(ctx context.Context) *Meta {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MetaCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MetaCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MetaCreate) defaults() {
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := meta.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := meta.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := meta.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetaCreate) check() error {
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Meta.updated_at"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Meta.created_at"`)}
	}
	return nil
}

func (mc *MetaCreate) sqlSave(ctx context.Context) (*Meta, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MetaCreate) createSpec() (*Meta, *sqlgraph.CreateSpec) {
	var (
		_node = &Meta{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(meta.Table, sqlgraph.NewFieldSpec(meta.FieldID, field.TypeUUID))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(meta.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(meta.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// MetaCreateBulk is the builder for creating many Meta entities in bulk.
type MetaCreateBulk struct {
	config
	builders []*MetaCreate
}

// Save creates the Meta entities in the database.
func (mcb *MetaCreateBulk) Save(ctx context.Context) ([]*Meta, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Meta, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetaCreateBulk) SaveX(ctx context.Context) []*Meta {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MetaCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MetaCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
