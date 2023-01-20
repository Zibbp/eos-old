// Code generated by ent, DO NOT EDIT.

package comment

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldLikeCount holds the string denoting the like_count field in the database.
	FieldLikeCount = "like_count"
	// FieldIsFavorited holds the string denoting the is_favorited field in the database.
	FieldIsFavorited = "is_favorited"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldAuthorID holds the string denoting the author_id field in the database.
	FieldAuthorID = "author_id"
	// FieldAuthorThumbnail holds the string denoting the author_thumbnail field in the database.
	FieldAuthorThumbnail = "author_thumbnail"
	// FieldAuthorIsUploader holds the string denoting the author_is_uploader field in the database.
	FieldAuthorIsUploader = "author_is_uploader"
	// FieldParent holds the string denoting the parent field in the database.
	FieldParent = "parent"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// VideoTable is the table that holds the video relation/edge.
	VideoTable = "comments"
	// VideoInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideoInverseTable = "videos"
	// VideoColumn is the table column denoting the video relation/edge.
	VideoColumn = "video_comments"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldTimestamp,
	FieldLikeCount,
	FieldIsFavorited,
	FieldAuthor,
	FieldAuthorID,
	FieldAuthorThumbnail,
	FieldAuthorIsUploader,
	FieldParent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"video_comments",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}