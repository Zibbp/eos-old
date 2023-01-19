// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/zibbp/eos/ent/comment"
	"github.com/zibbp/eos/ent/video"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// LikeCount holds the value of the "like_count" field.
	LikeCount int64 `json:"like_count,omitempty"`
	// IsFavorited holds the value of the "is_favorited" field.
	IsFavorited bool `json:"is_favorited,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// AuthorID holds the value of the "author_id" field.
	AuthorID string `json:"author_id,omitempty"`
	// AuthorThumbnail holds the value of the "author_thumbnail" field.
	AuthorThumbnail string `json:"author_thumbnail,omitempty"`
	// AuthorIsUploader holds the value of the "author_is_uploader" field.
	AuthorIsUploader bool `json:"author_is_uploader,omitempty"`
	// Parent holds the value of the "parent" field.
	Parent string `json:"parent,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges          CommentEdges `json:"edges"`
	video_comments *string
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Video holds the value of the video edge.
	Video *Video `json:"video,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) VideoOrErr() (*Video, error) {
	if e.loadedTypes[0] {
		if e.Video == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: video.Label}
		}
		return e.Video, nil
	}
	return nil, &NotLoadedError{edge: "video"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldIsFavorited, comment.FieldAuthorIsUploader:
			values[i] = new(sql.NullBool)
		case comment.FieldLikeCount:
			values[i] = new(sql.NullInt64)
		case comment.FieldID, comment.FieldText, comment.FieldAuthor, comment.FieldAuthorID, comment.FieldAuthorThumbnail, comment.FieldParent:
			values[i] = new(sql.NullString)
		case comment.FieldTimestamp:
			values[i] = new(sql.NullTime)
		case comment.ForeignKeys[0]: // video_comments
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case comment.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				c.Text = value.String
			}
		case comment.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				c.Timestamp = value.Time
			}
		case comment.FieldLikeCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field like_count", values[i])
			} else if value.Valid {
				c.LikeCount = value.Int64
			}
		case comment.FieldIsFavorited:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_favorited", values[i])
			} else if value.Valid {
				c.IsFavorited = value.Bool
			}
		case comment.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				c.Author = value.String
			}
		case comment.FieldAuthorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value.Valid {
				c.AuthorID = value.String
			}
		case comment.FieldAuthorThumbnail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_thumbnail", values[i])
			} else if value.Valid {
				c.AuthorThumbnail = value.String
			}
		case comment.FieldAuthorIsUploader:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field author_is_uploader", values[i])
			} else if value.Valid {
				c.AuthorIsUploader = value.Bool
			}
		case comment.FieldParent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent", values[i])
			} else if value.Valid {
				c.Parent = value.String
			}
		case comment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_comments", values[i])
			} else if value.Valid {
				c.video_comments = new(string)
				*c.video_comments = value.String
			}
		}
	}
	return nil
}

// QueryVideo queries the "video" edge of the Comment entity.
func (c *Comment) QueryVideo() *VideoQuery {
	return (&CommentClient{config: c.config}).QueryVideo(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("text=")
	builder.WriteString(c.Text)
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(c.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("like_count=")
	builder.WriteString(fmt.Sprintf("%v", c.LikeCount))
	builder.WriteString(", ")
	builder.WriteString("is_favorited=")
	builder.WriteString(fmt.Sprintf("%v", c.IsFavorited))
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(c.Author)
	builder.WriteString(", ")
	builder.WriteString("author_id=")
	builder.WriteString(c.AuthorID)
	builder.WriteString(", ")
	builder.WriteString("author_thumbnail=")
	builder.WriteString(c.AuthorThumbnail)
	builder.WriteString(", ")
	builder.WriteString("author_is_uploader=")
	builder.WriteString(fmt.Sprintf("%v", c.AuthorIsUploader))
	builder.WriteString(", ")
	builder.WriteString("parent=")
	builder.WriteString(c.Parent)
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
