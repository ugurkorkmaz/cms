// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/gallery"
	"app/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GalleryUpdate is the builder for updating Gallery entities.
type GalleryUpdate struct {
	config
	hooks    []Hook
	mutation *GalleryMutation
}

// Where appends a list predicates to the GalleryUpdate builder.
func (gu *GalleryUpdate) Where(ps ...predicate.Gallery) *GalleryUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GalleryUpdate) SetUpdatedAt(t time.Time) *GalleryUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetName sets the "name" field.
func (gu *GalleryUpdate) SetName(s string) *GalleryUpdate {
	gu.mutation.SetName(s)
	return gu
}

// Mutation returns the GalleryMutation object of the builder.
func (gu *GalleryUpdate) Mutation() *GalleryMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GalleryUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks[int, GalleryMutation](ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GalleryUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GalleryUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GalleryUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GalleryUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := gallery.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GalleryUpdate) check() error {
	if v, ok := gu.mutation.Name(); ok {
		if err := gallery.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Gallery.name": %w`, err)}
		}
	}
	return nil
}

func (gu *GalleryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(gallery.Table, gallery.Columns, sqlgraph.NewFieldSpec(gallery.FieldID, field.TypeUUID))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(gallery.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(gallery.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gallery.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GalleryUpdateOne is the builder for updating a single Gallery entity.
type GalleryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GalleryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GalleryUpdateOne) SetUpdatedAt(t time.Time) *GalleryUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetName sets the "name" field.
func (guo *GalleryUpdateOne) SetName(s string) *GalleryUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// Mutation returns the GalleryMutation object of the builder.
func (guo *GalleryUpdateOne) Mutation() *GalleryMutation {
	return guo.mutation
}

// Where appends a list predicates to the GalleryUpdate builder.
func (guo *GalleryUpdateOne) Where(ps ...predicate.Gallery) *GalleryUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GalleryUpdateOne) Select(field string, fields ...string) *GalleryUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Gallery entity.
func (guo *GalleryUpdateOne) Save(ctx context.Context) (*Gallery, error) {
	guo.defaults()
	return withHooks[*Gallery, GalleryMutation](ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GalleryUpdateOne) SaveX(ctx context.Context) *Gallery {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GalleryUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GalleryUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GalleryUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := gallery.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GalleryUpdateOne) check() error {
	if v, ok := guo.mutation.Name(); ok {
		if err := gallery.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Gallery.name": %w`, err)}
		}
	}
	return nil
}

func (guo *GalleryUpdateOne) sqlSave(ctx context.Context) (_node *Gallery, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(gallery.Table, gallery.Columns, sqlgraph.NewFieldSpec(gallery.FieldID, field.TypeUUID))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Gallery.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gallery.FieldID)
		for _, f := range fields {
			if !gallery.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gallery.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(gallery.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(gallery.FieldName, field.TypeString, value)
	}
	_node = &Gallery{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gallery.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
