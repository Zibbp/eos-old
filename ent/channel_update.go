// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/eos/ent/channel"
	"github.com/zibbp/eos/ent/predicate"
	"github.com/zibbp/eos/ent/video"
)

// ChannelUpdate is the builder for updating Channel entities.
type ChannelUpdate struct {
	config
	hooks    []Hook
	mutation *ChannelMutation
}

// Where appends a list predicates to the ChannelUpdate builder.
func (cu *ChannelUpdate) Where(ps ...predicate.Channel) *ChannelUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ChannelUpdate) SetName(s string) *ChannelUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *ChannelUpdate) SetDescription(s string) *ChannelUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableDescription(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *ChannelUpdate) ClearDescription() *ChannelUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetImagePath sets the "image_path" field.
func (cu *ChannelUpdate) SetImagePath(s string) *ChannelUpdate {
	cu.mutation.SetImagePath(s)
	return cu
}

// SetNillableImagePath sets the "image_path" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableImagePath(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetImagePath(*s)
	}
	return cu
}

// ClearImagePath clears the value of the "image_path" field.
func (cu *ChannelUpdate) ClearImagePath() *ChannelUpdate {
	cu.mutation.ClearImagePath()
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ChannelUpdate) SetUpdatedAt(t time.Time) *ChannelUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// AddVideoIDs adds the "videos" edge to the Video entity by IDs.
func (cu *ChannelUpdate) AddVideoIDs(ids ...string) *ChannelUpdate {
	cu.mutation.AddVideoIDs(ids...)
	return cu
}

// AddVideos adds the "videos" edges to the Video entity.
func (cu *ChannelUpdate) AddVideos(v ...*Video) *ChannelUpdate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.AddVideoIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cu *ChannelUpdate) Mutation() *ChannelMutation {
	return cu.mutation
}

// ClearVideos clears all "videos" edges to the Video entity.
func (cu *ChannelUpdate) ClearVideos() *ChannelUpdate {
	cu.mutation.ClearVideos()
	return cu
}

// RemoveVideoIDs removes the "videos" edge to Video entities by IDs.
func (cu *ChannelUpdate) RemoveVideoIDs(ids ...string) *ChannelUpdate {
	cu.mutation.RemoveVideoIDs(ids...)
	return cu
}

// RemoveVideos removes "videos" edges to Video entities.
func (cu *ChannelUpdate) RemoveVideos(v ...*Video) *ChannelUpdate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.RemoveVideoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChannelUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks[int, ChannelMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChannelUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChannelUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ChannelUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := channel.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: channel.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(channel.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(channel.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(channel.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.ImagePath(); ok {
		_spec.SetField(channel.FieldImagePath, field.TypeString, value)
	}
	if cu.mutation.ImagePathCleared() {
		_spec.ClearField(channel.FieldImagePath, field.TypeString)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(channel.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.VideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VideosTable,
			Columns: []string{channel.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: video.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedVideosIDs(); len(nodes) > 0 && !cu.mutation.VideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VideosTable,
			Columns: []string{channel.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.VideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VideosTable,
			Columns: []string{channel.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChannelUpdateOne is the builder for updating a single Channel entity.
type ChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChannelMutation
}

// SetName sets the "name" field.
func (cuo *ChannelUpdateOne) SetName(s string) *ChannelUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ChannelUpdateOne) SetDescription(s string) *ChannelUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableDescription(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *ChannelUpdateOne) ClearDescription() *ChannelUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetImagePath sets the "image_path" field.
func (cuo *ChannelUpdateOne) SetImagePath(s string) *ChannelUpdateOne {
	cuo.mutation.SetImagePath(s)
	return cuo
}

// SetNillableImagePath sets the "image_path" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableImagePath(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetImagePath(*s)
	}
	return cuo
}

// ClearImagePath clears the value of the "image_path" field.
func (cuo *ChannelUpdateOne) ClearImagePath() *ChannelUpdateOne {
	cuo.mutation.ClearImagePath()
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ChannelUpdateOne) SetUpdatedAt(t time.Time) *ChannelUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// AddVideoIDs adds the "videos" edge to the Video entity by IDs.
func (cuo *ChannelUpdateOne) AddVideoIDs(ids ...string) *ChannelUpdateOne {
	cuo.mutation.AddVideoIDs(ids...)
	return cuo
}

// AddVideos adds the "videos" edges to the Video entity.
func (cuo *ChannelUpdateOne) AddVideos(v ...*Video) *ChannelUpdateOne {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.AddVideoIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cuo *ChannelUpdateOne) Mutation() *ChannelMutation {
	return cuo.mutation
}

// ClearVideos clears all "videos" edges to the Video entity.
func (cuo *ChannelUpdateOne) ClearVideos() *ChannelUpdateOne {
	cuo.mutation.ClearVideos()
	return cuo
}

// RemoveVideoIDs removes the "videos" edge to Video entities by IDs.
func (cuo *ChannelUpdateOne) RemoveVideoIDs(ids ...string) *ChannelUpdateOne {
	cuo.mutation.RemoveVideoIDs(ids...)
	return cuo
}

// RemoveVideos removes "videos" edges to Video entities.
func (cuo *ChannelUpdateOne) RemoveVideos(v ...*Video) *ChannelUpdateOne {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.RemoveVideoIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChannelUpdateOne) Select(field string, fields ...string) *ChannelUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Channel entity.
func (cuo *ChannelUpdateOne) Save(ctx context.Context) (*Channel, error) {
	cuo.defaults()
	return withHooks[*Channel, ChannelMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChannelUpdateOne) SaveX(ctx context.Context) *Channel {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChannelUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ChannelUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := channel.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ChannelUpdateOne) sqlSave(ctx context.Context) (_node *Channel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: channel.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Channel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channel.FieldID)
		for _, f := range fields {
			if !channel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != channel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(channel.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(channel.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(channel.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.ImagePath(); ok {
		_spec.SetField(channel.FieldImagePath, field.TypeString, value)
	}
	if cuo.mutation.ImagePathCleared() {
		_spec.ClearField(channel.FieldImagePath, field.TypeString)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(channel.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.VideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VideosTable,
			Columns: []string{channel.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: video.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedVideosIDs(); len(nodes) > 0 && !cuo.mutation.VideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VideosTable,
			Columns: []string{channel.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.VideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VideosTable,
			Columns: []string{channel.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Channel{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}