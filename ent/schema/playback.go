package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Playback holds the schema definition for the Playback entity.
type Playback struct {
	ent.Schema
}

// Fields of the Playback.
func (Playback) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("video_id").NotEmpty(),
		field.Int("timestamp").Default(0),
		field.Enum("status").Values("in_progress", "finished").Optional().Default("in_progress"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Playback.
func (Playback) Edges() []ent.Edge {
	return nil
}
