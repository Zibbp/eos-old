// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/eos/ent/channel"
	"github.com/zibbp/eos/ent/chapter"
	"github.com/zibbp/eos/ent/comment"
	"github.com/zibbp/eos/ent/video"
)

// VideoCreate is the builder for creating a Video entity.
type VideoCreate struct {
	config
	mutation *VideoMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (vc *VideoCreate) SetTitle(s string) *VideoCreate {
	vc.mutation.SetTitle(s)
	return vc
}

// SetDescription sets the "description" field.
func (vc *VideoCreate) SetDescription(s string) *VideoCreate {
	vc.mutation.SetDescription(s)
	return vc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vc *VideoCreate) SetNillableDescription(s *string) *VideoCreate {
	if s != nil {
		vc.SetDescription(*s)
	}
	return vc
}

// SetUploadDate sets the "upload_date" field.
func (vc *VideoCreate) SetUploadDate(t time.Time) *VideoCreate {
	vc.mutation.SetUploadDate(t)
	return vc
}

// SetUploader sets the "uploader" field.
func (vc *VideoCreate) SetUploader(s string) *VideoCreate {
	vc.mutation.SetUploader(s)
	return vc
}

// SetDuration sets the "duration" field.
func (vc *VideoCreate) SetDuration(i int64) *VideoCreate {
	vc.mutation.SetDuration(i)
	return vc
}

// SetViewCount sets the "view_count" field.
func (vc *VideoCreate) SetViewCount(i int64) *VideoCreate {
	vc.mutation.SetViewCount(i)
	return vc
}

// SetLikeCount sets the "like_count" field.
func (vc *VideoCreate) SetLikeCount(i int64) *VideoCreate {
	vc.mutation.SetLikeCount(i)
	return vc
}

// SetDislikeCount sets the "dislike_count" field.
func (vc *VideoCreate) SetDislikeCount(i int64) *VideoCreate {
	vc.mutation.SetDislikeCount(i)
	return vc
}

// SetNillableDislikeCount sets the "dislike_count" field if the given value is not nil.
func (vc *VideoCreate) SetNillableDislikeCount(i *int64) *VideoCreate {
	if i != nil {
		vc.SetDislikeCount(*i)
	}
	return vc
}

// SetFormat sets the "format" field.
func (vc *VideoCreate) SetFormat(s string) *VideoCreate {
	vc.mutation.SetFormat(s)
	return vc
}

// SetNillableFormat sets the "format" field if the given value is not nil.
func (vc *VideoCreate) SetNillableFormat(s *string) *VideoCreate {
	if s != nil {
		vc.SetFormat(*s)
	}
	return vc
}

// SetWidth sets the "width" field.
func (vc *VideoCreate) SetWidth(i int64) *VideoCreate {
	vc.mutation.SetWidth(i)
	return vc
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (vc *VideoCreate) SetNillableWidth(i *int64) *VideoCreate {
	if i != nil {
		vc.SetWidth(*i)
	}
	return vc
}

// SetHeight sets the "height" field.
func (vc *VideoCreate) SetHeight(i int64) *VideoCreate {
	vc.mutation.SetHeight(i)
	return vc
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (vc *VideoCreate) SetNillableHeight(i *int64) *VideoCreate {
	if i != nil {
		vc.SetHeight(*i)
	}
	return vc
}

// SetResolution sets the "resolution" field.
func (vc *VideoCreate) SetResolution(s string) *VideoCreate {
	vc.mutation.SetResolution(s)
	return vc
}

// SetNillableResolution sets the "resolution" field if the given value is not nil.
func (vc *VideoCreate) SetNillableResolution(s *string) *VideoCreate {
	if s != nil {
		vc.SetResolution(*s)
	}
	return vc
}

// SetFps sets the "fps" field.
func (vc *VideoCreate) SetFps(f float64) *VideoCreate {
	vc.mutation.SetFps(f)
	return vc
}

// SetNillableFps sets the "fps" field if the given value is not nil.
func (vc *VideoCreate) SetNillableFps(f *float64) *VideoCreate {
	if f != nil {
		vc.SetFps(*f)
	}
	return vc
}

// SetAudioCodec sets the "audio_codec" field.
func (vc *VideoCreate) SetAudioCodec(s string) *VideoCreate {
	vc.mutation.SetAudioCodec(s)
	return vc
}

// SetNillableAudioCodec sets the "audio_codec" field if the given value is not nil.
func (vc *VideoCreate) SetNillableAudioCodec(s *string) *VideoCreate {
	if s != nil {
		vc.SetAudioCodec(*s)
	}
	return vc
}

// SetVideoCodec sets the "video_codec" field.
func (vc *VideoCreate) SetVideoCodec(s string) *VideoCreate {
	vc.mutation.SetVideoCodec(s)
	return vc
}

// SetNillableVideoCodec sets the "video_codec" field if the given value is not nil.
func (vc *VideoCreate) SetNillableVideoCodec(s *string) *VideoCreate {
	if s != nil {
		vc.SetVideoCodec(*s)
	}
	return vc
}

// SetAbr sets the "abr" field.
func (vc *VideoCreate) SetAbr(f float64) *VideoCreate {
	vc.mutation.SetAbr(f)
	return vc
}

// SetNillableAbr sets the "abr" field if the given value is not nil.
func (vc *VideoCreate) SetNillableAbr(f *float64) *VideoCreate {
	if f != nil {
		vc.SetAbr(*f)
	}
	return vc
}

// SetVbr sets the "vbr" field.
func (vc *VideoCreate) SetVbr(f float64) *VideoCreate {
	vc.mutation.SetVbr(f)
	return vc
}

// SetNillableVbr sets the "vbr" field if the given value is not nil.
func (vc *VideoCreate) SetNillableVbr(f *float64) *VideoCreate {
	if f != nil {
		vc.SetVbr(*f)
	}
	return vc
}

// SetEpoch sets the "epoch" field.
func (vc *VideoCreate) SetEpoch(i int64) *VideoCreate {
	vc.mutation.SetEpoch(i)
	return vc
}

// SetNillableEpoch sets the "epoch" field if the given value is not nil.
func (vc *VideoCreate) SetNillableEpoch(i *int64) *VideoCreate {
	if i != nil {
		vc.SetEpoch(*i)
	}
	return vc
}

// SetCommentCount sets the "comment_count" field.
func (vc *VideoCreate) SetCommentCount(i int64) *VideoCreate {
	vc.mutation.SetCommentCount(i)
	return vc
}

// SetNillableCommentCount sets the "comment_count" field if the given value is not nil.
func (vc *VideoCreate) SetNillableCommentCount(i *int64) *VideoCreate {
	if i != nil {
		vc.SetCommentCount(*i)
	}
	return vc
}

// SetTags sets the "tags" field.
func (vc *VideoCreate) SetTags(s string) *VideoCreate {
	vc.mutation.SetTags(s)
	return vc
}

// SetNillableTags sets the "tags" field if the given value is not nil.
func (vc *VideoCreate) SetNillableTags(s *string) *VideoCreate {
	if s != nil {
		vc.SetTags(*s)
	}
	return vc
}

// SetCategories sets the "categories" field.
func (vc *VideoCreate) SetCategories(s string) *VideoCreate {
	vc.mutation.SetCategories(s)
	return vc
}

// SetNillableCategories sets the "categories" field if the given value is not nil.
func (vc *VideoCreate) SetNillableCategories(s *string) *VideoCreate {
	if s != nil {
		vc.SetCategories(*s)
	}
	return vc
}

// SetVideoPath sets the "video_path" field.
func (vc *VideoCreate) SetVideoPath(s string) *VideoCreate {
	vc.mutation.SetVideoPath(s)
	return vc
}

// SetThumbnailPath sets the "thumbnail_path" field.
func (vc *VideoCreate) SetThumbnailPath(s string) *VideoCreate {
	vc.mutation.SetThumbnailPath(s)
	return vc
}

// SetJSONPath sets the "json_path" field.
func (vc *VideoCreate) SetJSONPath(s string) *VideoCreate {
	vc.mutation.SetJSONPath(s)
	return vc
}

// SetCaptionPath sets the "caption_path" field.
func (vc *VideoCreate) SetCaptionPath(s string) *VideoCreate {
	vc.mutation.SetCaptionPath(s)
	return vc
}

// SetNillableCaptionPath sets the "caption_path" field if the given value is not nil.
func (vc *VideoCreate) SetNillableCaptionPath(s *string) *VideoCreate {
	if s != nil {
		vc.SetCaptionPath(*s)
	}
	return vc
}

// SetPath sets the "path" field.
func (vc *VideoCreate) SetPath(s string) *VideoCreate {
	vc.mutation.SetPath(s)
	return vc
}

// SetThumbnailsPath sets the "thumbnails_path" field.
func (vc *VideoCreate) SetThumbnailsPath(s string) *VideoCreate {
	vc.mutation.SetThumbnailsPath(s)
	return vc
}

// SetNillableThumbnailsPath sets the "thumbnails_path" field if the given value is not nil.
func (vc *VideoCreate) SetNillableThumbnailsPath(s *string) *VideoCreate {
	if s != nil {
		vc.SetThumbnailsPath(*s)
	}
	return vc
}

// SetThumbnailsWidth sets the "thumbnails_width" field.
func (vc *VideoCreate) SetThumbnailsWidth(i int) *VideoCreate {
	vc.mutation.SetThumbnailsWidth(i)
	return vc
}

// SetNillableThumbnailsWidth sets the "thumbnails_width" field if the given value is not nil.
func (vc *VideoCreate) SetNillableThumbnailsWidth(i *int) *VideoCreate {
	if i != nil {
		vc.SetThumbnailsWidth(*i)
	}
	return vc
}

// SetThumbnailsHeight sets the "thumbnails_height" field.
func (vc *VideoCreate) SetThumbnailsHeight(i int) *VideoCreate {
	vc.mutation.SetThumbnailsHeight(i)
	return vc
}

// SetNillableThumbnailsHeight sets the "thumbnails_height" field if the given value is not nil.
func (vc *VideoCreate) SetNillableThumbnailsHeight(i *int) *VideoCreate {
	if i != nil {
		vc.SetThumbnailsHeight(*i)
	}
	return vc
}

// SetThumbnailsInterval sets the "thumbnails_interval" field.
func (vc *VideoCreate) SetThumbnailsInterval(f float64) *VideoCreate {
	vc.mutation.SetThumbnailsInterval(f)
	return vc
}

// SetNillableThumbnailsInterval sets the "thumbnails_interval" field if the given value is not nil.
func (vc *VideoCreate) SetNillableThumbnailsInterval(f *float64) *VideoCreate {
	if f != nil {
		vc.SetThumbnailsInterval(*f)
	}
	return vc
}

// SetEosGeneratedThumbnails sets the "eos_generated_thumbnails" field.
func (vc *VideoCreate) SetEosGeneratedThumbnails(b bool) *VideoCreate {
	vc.mutation.SetEosGeneratedThumbnails(b)
	return vc
}

// SetNillableEosGeneratedThumbnails sets the "eos_generated_thumbnails" field if the given value is not nil.
func (vc *VideoCreate) SetNillableEosGeneratedThumbnails(b *bool) *VideoCreate {
	if b != nil {
		vc.SetEosGeneratedThumbnails(*b)
	}
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VideoCreate) SetCreatedAt(t time.Time) *VideoCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableCreatedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VideoCreate) SetUpdatedAt(t time.Time) *VideoCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableUpdatedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VideoCreate) SetID(s string) *VideoCreate {
	vc.mutation.SetID(s)
	return vc
}

// SetChannelID sets the "channel" edge to the Channel entity by ID.
func (vc *VideoCreate) SetChannelID(id string) *VideoCreate {
	vc.mutation.SetChannelID(id)
	return vc
}

// SetChannel sets the "channel" edge to the Channel entity.
func (vc *VideoCreate) SetChannel(c *Channel) *VideoCreate {
	return vc.SetChannelID(c.ID)
}

// AddChapterIDs adds the "chapters" edge to the Chapter entity by IDs.
func (vc *VideoCreate) AddChapterIDs(ids ...string) *VideoCreate {
	vc.mutation.AddChapterIDs(ids...)
	return vc
}

// AddChapters adds the "chapters" edges to the Chapter entity.
func (vc *VideoCreate) AddChapters(c ...*Chapter) *VideoCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vc.AddChapterIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (vc *VideoCreate) AddCommentIDs(ids ...string) *VideoCreate {
	vc.mutation.AddCommentIDs(ids...)
	return vc
}

// AddComments adds the "comments" edges to the Comment entity.
func (vc *VideoCreate) AddComments(c ...*Comment) *VideoCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vc.AddCommentIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vc *VideoCreate) Mutation() *VideoMutation {
	return vc.mutation
}

// Save creates the Video in the database.
func (vc *VideoCreate) Save(ctx context.Context) (*Video, error) {
	vc.defaults()
	return withHooks[*Video, VideoMutation](ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideoCreate) SaveX(ctx context.Context) *Video {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VideoCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VideoCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VideoCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := video.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := video.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VideoCreate) check() error {
	if _, ok := vc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Video.title"`)}
	}
	if _, ok := vc.mutation.UploadDate(); !ok {
		return &ValidationError{Name: "upload_date", err: errors.New(`ent: missing required field "Video.upload_date"`)}
	}
	if _, ok := vc.mutation.Uploader(); !ok {
		return &ValidationError{Name: "uploader", err: errors.New(`ent: missing required field "Video.uploader"`)}
	}
	if _, ok := vc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Video.duration"`)}
	}
	if _, ok := vc.mutation.ViewCount(); !ok {
		return &ValidationError{Name: "view_count", err: errors.New(`ent: missing required field "Video.view_count"`)}
	}
	if _, ok := vc.mutation.LikeCount(); !ok {
		return &ValidationError{Name: "like_count", err: errors.New(`ent: missing required field "Video.like_count"`)}
	}
	if _, ok := vc.mutation.VideoPath(); !ok {
		return &ValidationError{Name: "video_path", err: errors.New(`ent: missing required field "Video.video_path"`)}
	}
	if _, ok := vc.mutation.ThumbnailPath(); !ok {
		return &ValidationError{Name: "thumbnail_path", err: errors.New(`ent: missing required field "Video.thumbnail_path"`)}
	}
	if _, ok := vc.mutation.JSONPath(); !ok {
		return &ValidationError{Name: "json_path", err: errors.New(`ent: missing required field "Video.json_path"`)}
	}
	if _, ok := vc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Video.path"`)}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Video.created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Video.updated_at"`)}
	}
	if _, ok := vc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel", err: errors.New(`ent: missing required edge "Video.channel"`)}
	}
	return nil
}

func (vc *VideoCreate) sqlSave(ctx context.Context) (*Video, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Video.ID type: %T", _spec.ID.Value)
		}
	}
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VideoCreate) createSpec() (*Video, *sqlgraph.CreateSpec) {
	var (
		_node = &Video{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: video.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: video.FieldID,
			},
		}
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.Title(); ok {
		_spec.SetField(video.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := vc.mutation.Description(); ok {
		_spec.SetField(video.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := vc.mutation.UploadDate(); ok {
		_spec.SetField(video.FieldUploadDate, field.TypeTime, value)
		_node.UploadDate = value
	}
	if value, ok := vc.mutation.Uploader(); ok {
		_spec.SetField(video.FieldUploader, field.TypeString, value)
		_node.Uploader = value
	}
	if value, ok := vc.mutation.Duration(); ok {
		_spec.SetField(video.FieldDuration, field.TypeInt64, value)
		_node.Duration = value
	}
	if value, ok := vc.mutation.ViewCount(); ok {
		_spec.SetField(video.FieldViewCount, field.TypeInt64, value)
		_node.ViewCount = value
	}
	if value, ok := vc.mutation.LikeCount(); ok {
		_spec.SetField(video.FieldLikeCount, field.TypeInt64, value)
		_node.LikeCount = value
	}
	if value, ok := vc.mutation.DislikeCount(); ok {
		_spec.SetField(video.FieldDislikeCount, field.TypeInt64, value)
		_node.DislikeCount = value
	}
	if value, ok := vc.mutation.Format(); ok {
		_spec.SetField(video.FieldFormat, field.TypeString, value)
		_node.Format = value
	}
	if value, ok := vc.mutation.Width(); ok {
		_spec.SetField(video.FieldWidth, field.TypeInt64, value)
		_node.Width = value
	}
	if value, ok := vc.mutation.Height(); ok {
		_spec.SetField(video.FieldHeight, field.TypeInt64, value)
		_node.Height = value
	}
	if value, ok := vc.mutation.Resolution(); ok {
		_spec.SetField(video.FieldResolution, field.TypeString, value)
		_node.Resolution = value
	}
	if value, ok := vc.mutation.Fps(); ok {
		_spec.SetField(video.FieldFps, field.TypeFloat64, value)
		_node.Fps = value
	}
	if value, ok := vc.mutation.AudioCodec(); ok {
		_spec.SetField(video.FieldAudioCodec, field.TypeString, value)
		_node.AudioCodec = value
	}
	if value, ok := vc.mutation.VideoCodec(); ok {
		_spec.SetField(video.FieldVideoCodec, field.TypeString, value)
		_node.VideoCodec = value
	}
	if value, ok := vc.mutation.Abr(); ok {
		_spec.SetField(video.FieldAbr, field.TypeFloat64, value)
		_node.Abr = value
	}
	if value, ok := vc.mutation.Vbr(); ok {
		_spec.SetField(video.FieldVbr, field.TypeFloat64, value)
		_node.Vbr = value
	}
	if value, ok := vc.mutation.Epoch(); ok {
		_spec.SetField(video.FieldEpoch, field.TypeInt64, value)
		_node.Epoch = value
	}
	if value, ok := vc.mutation.CommentCount(); ok {
		_spec.SetField(video.FieldCommentCount, field.TypeInt64, value)
		_node.CommentCount = value
	}
	if value, ok := vc.mutation.Tags(); ok {
		_spec.SetField(video.FieldTags, field.TypeString, value)
		_node.Tags = value
	}
	if value, ok := vc.mutation.Categories(); ok {
		_spec.SetField(video.FieldCategories, field.TypeString, value)
		_node.Categories = value
	}
	if value, ok := vc.mutation.VideoPath(); ok {
		_spec.SetField(video.FieldVideoPath, field.TypeString, value)
		_node.VideoPath = value
	}
	if value, ok := vc.mutation.ThumbnailPath(); ok {
		_spec.SetField(video.FieldThumbnailPath, field.TypeString, value)
		_node.ThumbnailPath = value
	}
	if value, ok := vc.mutation.JSONPath(); ok {
		_spec.SetField(video.FieldJSONPath, field.TypeString, value)
		_node.JSONPath = value
	}
	if value, ok := vc.mutation.CaptionPath(); ok {
		_spec.SetField(video.FieldCaptionPath, field.TypeString, value)
		_node.CaptionPath = value
	}
	if value, ok := vc.mutation.Path(); ok {
		_spec.SetField(video.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := vc.mutation.ThumbnailsPath(); ok {
		_spec.SetField(video.FieldThumbnailsPath, field.TypeString, value)
		_node.ThumbnailsPath = value
	}
	if value, ok := vc.mutation.ThumbnailsWidth(); ok {
		_spec.SetField(video.FieldThumbnailsWidth, field.TypeInt, value)
		_node.ThumbnailsWidth = value
	}
	if value, ok := vc.mutation.ThumbnailsHeight(); ok {
		_spec.SetField(video.FieldThumbnailsHeight, field.TypeInt, value)
		_node.ThumbnailsHeight = value
	}
	if value, ok := vc.mutation.ThumbnailsInterval(); ok {
		_spec.SetField(video.FieldThumbnailsInterval, field.TypeFloat64, value)
		_node.ThumbnailsInterval = value
	}
	if value, ok := vc.mutation.EosGeneratedThumbnails(); ok {
		_spec.SetField(video.FieldEosGeneratedThumbnails, field.TypeBool, value)
		_node.EosGeneratedThumbnails = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(video.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(video.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := vc.mutation.ChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.ChannelTable,
			Columns: []string{video.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: channel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.channel_videos = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ChaptersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.ChaptersTable,
			Columns: []string{video.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: chapter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.CommentsTable,
			Columns: []string{video.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VideoCreateBulk is the builder for creating many Video entities in bulk.
type VideoCreateBulk struct {
	config
	builders []*VideoCreate
}

// Save creates the Video entities in the database.
func (vcb *VideoCreateBulk) Save(ctx context.Context) ([]*Video, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Video, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VideoCreateBulk) SaveX(ctx context.Context) []*Video {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VideoCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VideoCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
