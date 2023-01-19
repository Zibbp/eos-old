// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/eos/ent/comment"
	"github.com/zibbp/eos/ent/video"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (cc *CommentCreate) SetText(s string) *CommentCreate {
	cc.mutation.SetText(s)
	return cc
}

// SetTimestamp sets the "timestamp" field.
func (cc *CommentCreate) SetTimestamp(t time.Time) *CommentCreate {
	cc.mutation.SetTimestamp(t)
	return cc
}

// SetLikeCount sets the "like_count" field.
func (cc *CommentCreate) SetLikeCount(i int64) *CommentCreate {
	cc.mutation.SetLikeCount(i)
	return cc
}

// SetIsFavorited sets the "is_favorited" field.
func (cc *CommentCreate) SetIsFavorited(b bool) *CommentCreate {
	cc.mutation.SetIsFavorited(b)
	return cc
}

// SetAuthor sets the "author" field.
func (cc *CommentCreate) SetAuthor(s string) *CommentCreate {
	cc.mutation.SetAuthor(s)
	return cc
}

// SetAuthorID sets the "author_id" field.
func (cc *CommentCreate) SetAuthorID(s string) *CommentCreate {
	cc.mutation.SetAuthorID(s)
	return cc
}

// SetAuthorThumbnail sets the "author_thumbnail" field.
func (cc *CommentCreate) SetAuthorThumbnail(s string) *CommentCreate {
	cc.mutation.SetAuthorThumbnail(s)
	return cc
}

// SetAuthorIsUploader sets the "author_is_uploader" field.
func (cc *CommentCreate) SetAuthorIsUploader(b bool) *CommentCreate {
	cc.mutation.SetAuthorIsUploader(b)
	return cc
}

// SetParent sets the "parent" field.
func (cc *CommentCreate) SetParent(s string) *CommentCreate {
	cc.mutation.SetParent(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CommentCreate) SetID(s string) *CommentCreate {
	cc.mutation.SetID(s)
	return cc
}

// SetVideoID sets the "video" edge to the Video entity by ID.
func (cc *CommentCreate) SetVideoID(id string) *CommentCreate {
	cc.mutation.SetVideoID(id)
	return cc
}

// SetVideo sets the "video" edge to the Video entity.
func (cc *CommentCreate) SetVideo(v *Video) *CommentCreate {
	return cc.SetVideoID(v.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	return withHooks[*Comment, CommentMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Comment.text"`)}
	}
	if _, ok := cc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Comment.timestamp"`)}
	}
	if _, ok := cc.mutation.LikeCount(); !ok {
		return &ValidationError{Name: "like_count", err: errors.New(`ent: missing required field "Comment.like_count"`)}
	}
	if _, ok := cc.mutation.IsFavorited(); !ok {
		return &ValidationError{Name: "is_favorited", err: errors.New(`ent: missing required field "Comment.is_favorited"`)}
	}
	if _, ok := cc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "Comment.author"`)}
	}
	if _, ok := cc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author_id", err: errors.New(`ent: missing required field "Comment.author_id"`)}
	}
	if _, ok := cc.mutation.AuthorThumbnail(); !ok {
		return &ValidationError{Name: "author_thumbnail", err: errors.New(`ent: missing required field "Comment.author_thumbnail"`)}
	}
	if _, ok := cc.mutation.AuthorIsUploader(); !ok {
		return &ValidationError{Name: "author_is_uploader", err: errors.New(`ent: missing required field "Comment.author_is_uploader"`)}
	}
	if _, ok := cc.mutation.Parent(); !ok {
		return &ValidationError{Name: "parent", err: errors.New(`ent: missing required field "Comment.parent"`)}
	}
	if _, ok := cc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video", err: errors.New(`ent: missing required edge "Comment.video"`)}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
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
			return nil, fmt.Errorf("unexpected Comment.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: comment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: comment.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Text(); ok {
		_spec.SetField(comment.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := cc.mutation.Timestamp(); ok {
		_spec.SetField(comment.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := cc.mutation.LikeCount(); ok {
		_spec.SetField(comment.FieldLikeCount, field.TypeInt64, value)
		_node.LikeCount = value
	}
	if value, ok := cc.mutation.IsFavorited(); ok {
		_spec.SetField(comment.FieldIsFavorited, field.TypeBool, value)
		_node.IsFavorited = value
	}
	if value, ok := cc.mutation.Author(); ok {
		_spec.SetField(comment.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := cc.mutation.AuthorID(); ok {
		_spec.SetField(comment.FieldAuthorID, field.TypeString, value)
		_node.AuthorID = value
	}
	if value, ok := cc.mutation.AuthorThumbnail(); ok {
		_spec.SetField(comment.FieldAuthorThumbnail, field.TypeString, value)
		_node.AuthorThumbnail = value
	}
	if value, ok := cc.mutation.AuthorIsUploader(); ok {
		_spec.SetField(comment.FieldAuthorIsUploader, field.TypeBool, value)
		_node.AuthorIsUploader = value
	}
	if value, ok := cc.mutation.Parent(); ok {
		_spec.SetField(comment.FieldParent, field.TypeString, value)
		_node.Parent = value
	}
	if nodes := cc.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.VideoTable,
			Columns: []string{comment.VideoColumn},
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
		_node.video_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	builders []*CommentCreate
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
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
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
