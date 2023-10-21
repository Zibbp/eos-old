package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("title"),
		field.String("description").Optional(),
		field.Time("upload_date"),
		field.String("uploader"),
		field.Int64("duration"),
		field.Int64("view_count"),
		field.Int64("like_count"),
		field.Int64("dislike_count").Optional(),
		field.String("format").Optional(),
		field.Int64("width").Optional(),
		field.Int64("height").Optional(),
		field.String("resolution").Optional(),
		field.Float("fps").Optional(),
		field.String("audio_codec").Optional(),
		field.String("video_codec").Optional(),
		field.Float("abr").Optional(),
		field.Float("vbr").Optional(),
		field.Int64("epoch").Optional(),
		field.Int64("comment_count").Optional(),
		field.String("tags").Optional(),
		field.String("categories").Optional(),
		field.String("video_path"),
		field.String("thumbnail_path"),
		field.String("json_path"),
		field.String("caption_path").Optional(),
		field.String("path"),
		field.String("thumbnails_path").Optional().Comment("path to storyboard of thumbnails"),
		field.Int("thumbnails_width").Optional().Comment("width of storyboard of thumbnails"),
		field.Int("thumbnails_height").Optional().Comment("height of storyboard of thumbnails"),
		field.Float("thumbnails_interval").Optional().Comment("interval of storyboard of thumbnails"),
		field.Int("thumbnails_rows").Optional().Comment("rows of storyboard of thumbnails"),
		field.Bool("eos_generated_thumbnails").Optional().Comment("whether the storyboard of thumbnails is generated by eos"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("channel", Channel.Type).Ref("videos").Unique().Required(),
		edge.To("chapters", Chapter.Type),
		edge.To("comments", Comment.Type),
	}
}
