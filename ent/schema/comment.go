package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("text"),
		field.Time("timestamp"),
		field.Int64("like_count"),
		field.Bool("is_favorited"),
		field.String("author"),
		field.String("author_id"),
		field.String("author_thumbnail"),
		field.Bool("author_is_uploader"),
		field.String("parent"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("comments").Unique().Required(),
	}
}
