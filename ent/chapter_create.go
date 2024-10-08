// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/eos/ent/chapter"
	"github.com/zibbp/eos/ent/video"
)

// ChapterCreate is the builder for creating a Chapter entity.
type ChapterCreate struct {
	config
	mutation *ChapterMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (cc *ChapterCreate) SetTitle(s string) *ChapterCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetStartTime sets the "start_time" field.
func (cc *ChapterCreate) SetStartTime(f float64) *ChapterCreate {
	cc.mutation.SetStartTime(f)
	return cc
}

// SetEndTime sets the "end_time" field.
func (cc *ChapterCreate) SetEndTime(f float64) *ChapterCreate {
	cc.mutation.SetEndTime(f)
	return cc
}

// SetID sets the "id" field.
func (cc *ChapterCreate) SetID(s string) *ChapterCreate {
	cc.mutation.SetID(s)
	return cc
}

// SetVideoID sets the "video" edge to the Video entity by ID.
func (cc *ChapterCreate) SetVideoID(id string) *ChapterCreate {
	cc.mutation.SetVideoID(id)
	return cc
}

// SetVideo sets the "video" edge to the Video entity.
func (cc *ChapterCreate) SetVideo(v *Video) *ChapterCreate {
	return cc.SetVideoID(v.ID)
}

// Mutation returns the ChapterMutation object of the builder.
func (cc *ChapterCreate) Mutation() *ChapterMutation {
	return cc.mutation
}

// Save creates the Chapter in the database.
func (cc *ChapterCreate) Save(ctx context.Context) (*Chapter, error) {
	return withHooks[*Chapter, ChapterMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChapterCreate) SaveX(ctx context.Context) *Chapter {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChapterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChapterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChapterCreate) check() error {
	if _, ok := cc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Chapter.title"`)}
	}
	if _, ok := cc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "Chapter.start_time"`)}
	}
	if _, ok := cc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "Chapter.end_time"`)}
	}
	if _, ok := cc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video", err: errors.New(`ent: missing required edge "Chapter.video"`)}
	}
	return nil
}

func (cc *ChapterCreate) sqlSave(ctx context.Context) (*Chapter, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Chapter.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChapterCreate) createSpec() (*Chapter, *sqlgraph.CreateSpec) {
	var (
		_node = &Chapter{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: chapter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chapter.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Title(); ok {
		_spec.SetField(chapter.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := cc.mutation.StartTime(); ok {
		_spec.SetField(chapter.FieldStartTime, field.TypeFloat64, value)
		_node.StartTime = value
	}
	if value, ok := cc.mutation.EndTime(); ok {
		_spec.SetField(chapter.FieldEndTime, field.TypeFloat64, value)
		_node.EndTime = value
	}
	if nodes := cc.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chapter.VideoTable,
			Columns: []string{chapter.VideoColumn},
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
		_node.video_chapters = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChapterCreateBulk is the builder for creating many Chapter entities in bulk.
type ChapterCreateBulk struct {
	config
	builders []*ChapterCreate
}

// Save creates the Chapter entities in the database.
func (ccb *ChapterCreateBulk) Save(ctx context.Context) ([]*Chapter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chapter, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChapterMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChapterCreateBulk) SaveX(ctx context.Context) []*Chapter {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChapterCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChapterCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
