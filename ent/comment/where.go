// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/zibbp/eos/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldText, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldTimestamp, v))
}

// LikeCount applies equality check predicate on the "like_count" field. It's identical to LikeCountEQ.
func LikeCount(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldLikeCount, v))
}

// IsFavorited applies equality check predicate on the "is_favorited" field. It's identical to IsFavoritedEQ.
func IsFavorited(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldIsFavorited, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldAuthor, v))
}

// AuthorID applies equality check predicate on the "author_id" field. It's identical to AuthorIDEQ.
func AuthorID(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldAuthorID, v))
}

// AuthorThumbnail applies equality check predicate on the "author_thumbnail" field. It's identical to AuthorThumbnailEQ.
func AuthorThumbnail(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldAuthorThumbnail, v))
}

// AuthorIsUploader applies equality check predicate on the "author_is_uploader" field. It's identical to AuthorIsUploaderEQ.
func AuthorIsUploader(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldAuthorIsUploader, v))
}

// Parent applies equality check predicate on the "parent" field. It's identical to ParentEQ.
func Parent(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldParent, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldText, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldTimestamp, v))
}

// LikeCountEQ applies the EQ predicate on the "like_count" field.
func LikeCountEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldLikeCount, v))
}

// LikeCountNEQ applies the NEQ predicate on the "like_count" field.
func LikeCountNEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldLikeCount, v))
}

// LikeCountIn applies the In predicate on the "like_count" field.
func LikeCountIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldLikeCount, vs...))
}

// LikeCountNotIn applies the NotIn predicate on the "like_count" field.
func LikeCountNotIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldLikeCount, vs...))
}

// LikeCountGT applies the GT predicate on the "like_count" field.
func LikeCountGT(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldLikeCount, v))
}

// LikeCountGTE applies the GTE predicate on the "like_count" field.
func LikeCountGTE(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldLikeCount, v))
}

// LikeCountLT applies the LT predicate on the "like_count" field.
func LikeCountLT(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldLikeCount, v))
}

// LikeCountLTE applies the LTE predicate on the "like_count" field.
func LikeCountLTE(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldLikeCount, v))
}

// IsFavoritedEQ applies the EQ predicate on the "is_favorited" field.
func IsFavoritedEQ(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldIsFavorited, v))
}

// IsFavoritedNEQ applies the NEQ predicate on the "is_favorited" field.
func IsFavoritedNEQ(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldIsFavorited, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldAuthor, v))
}

// AuthorIDEQ applies the EQ predicate on the "author_id" field.
func AuthorIDEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldAuthorID, v))
}

// AuthorIDNEQ applies the NEQ predicate on the "author_id" field.
func AuthorIDNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldAuthorID, v))
}

// AuthorIDIn applies the In predicate on the "author_id" field.
func AuthorIDIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldAuthorID, vs...))
}

// AuthorIDNotIn applies the NotIn predicate on the "author_id" field.
func AuthorIDNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldAuthorID, vs...))
}

// AuthorIDGT applies the GT predicate on the "author_id" field.
func AuthorIDGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldAuthorID, v))
}

// AuthorIDGTE applies the GTE predicate on the "author_id" field.
func AuthorIDGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldAuthorID, v))
}

// AuthorIDLT applies the LT predicate on the "author_id" field.
func AuthorIDLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldAuthorID, v))
}

// AuthorIDLTE applies the LTE predicate on the "author_id" field.
func AuthorIDLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldAuthorID, v))
}

// AuthorIDContains applies the Contains predicate on the "author_id" field.
func AuthorIDContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldAuthorID, v))
}

// AuthorIDHasPrefix applies the HasPrefix predicate on the "author_id" field.
func AuthorIDHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldAuthorID, v))
}

// AuthorIDHasSuffix applies the HasSuffix predicate on the "author_id" field.
func AuthorIDHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldAuthorID, v))
}

// AuthorIDEqualFold applies the EqualFold predicate on the "author_id" field.
func AuthorIDEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldAuthorID, v))
}

// AuthorIDContainsFold applies the ContainsFold predicate on the "author_id" field.
func AuthorIDContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldAuthorID, v))
}

// AuthorThumbnailEQ applies the EQ predicate on the "author_thumbnail" field.
func AuthorThumbnailEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldAuthorThumbnail, v))
}

// AuthorThumbnailNEQ applies the NEQ predicate on the "author_thumbnail" field.
func AuthorThumbnailNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldAuthorThumbnail, v))
}

// AuthorThumbnailIn applies the In predicate on the "author_thumbnail" field.
func AuthorThumbnailIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldAuthorThumbnail, vs...))
}

// AuthorThumbnailNotIn applies the NotIn predicate on the "author_thumbnail" field.
func AuthorThumbnailNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldAuthorThumbnail, vs...))
}

// AuthorThumbnailGT applies the GT predicate on the "author_thumbnail" field.
func AuthorThumbnailGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldAuthorThumbnail, v))
}

// AuthorThumbnailGTE applies the GTE predicate on the "author_thumbnail" field.
func AuthorThumbnailGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldAuthorThumbnail, v))
}

// AuthorThumbnailLT applies the LT predicate on the "author_thumbnail" field.
func AuthorThumbnailLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldAuthorThumbnail, v))
}

// AuthorThumbnailLTE applies the LTE predicate on the "author_thumbnail" field.
func AuthorThumbnailLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldAuthorThumbnail, v))
}

// AuthorThumbnailContains applies the Contains predicate on the "author_thumbnail" field.
func AuthorThumbnailContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldAuthorThumbnail, v))
}

// AuthorThumbnailHasPrefix applies the HasPrefix predicate on the "author_thumbnail" field.
func AuthorThumbnailHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldAuthorThumbnail, v))
}

// AuthorThumbnailHasSuffix applies the HasSuffix predicate on the "author_thumbnail" field.
func AuthorThumbnailHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldAuthorThumbnail, v))
}

// AuthorThumbnailEqualFold applies the EqualFold predicate on the "author_thumbnail" field.
func AuthorThumbnailEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldAuthorThumbnail, v))
}

// AuthorThumbnailContainsFold applies the ContainsFold predicate on the "author_thumbnail" field.
func AuthorThumbnailContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldAuthorThumbnail, v))
}

// AuthorIsUploaderEQ applies the EQ predicate on the "author_is_uploader" field.
func AuthorIsUploaderEQ(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldAuthorIsUploader, v))
}

// AuthorIsUploaderNEQ applies the NEQ predicate on the "author_is_uploader" field.
func AuthorIsUploaderNEQ(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldAuthorIsUploader, v))
}

// ParentEQ applies the EQ predicate on the "parent" field.
func ParentEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldParent, v))
}

// ParentNEQ applies the NEQ predicate on the "parent" field.
func ParentNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldParent, v))
}

// ParentIn applies the In predicate on the "parent" field.
func ParentIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldParent, vs...))
}

// ParentNotIn applies the NotIn predicate on the "parent" field.
func ParentNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldParent, vs...))
}

// ParentGT applies the GT predicate on the "parent" field.
func ParentGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldParent, v))
}

// ParentGTE applies the GTE predicate on the "parent" field.
func ParentGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldParent, v))
}

// ParentLT applies the LT predicate on the "parent" field.
func ParentLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldParent, v))
}

// ParentLTE applies the LTE predicate on the "parent" field.
func ParentLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldParent, v))
}

// ParentContains applies the Contains predicate on the "parent" field.
func ParentContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldParent, v))
}

// ParentHasPrefix applies the HasPrefix predicate on the "parent" field.
func ParentHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldParent, v))
}

// ParentHasSuffix applies the HasSuffix predicate on the "parent" field.
func ParentHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldParent, v))
}

// ParentEqualFold applies the EqualFold predicate on the "parent" field.
func ParentEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldParent, v))
}

// ParentContainsFold applies the ContainsFold predicate on the "parent" field.
func ParentContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldParent, v))
}

// HasVideo applies the HasEdge predicate on the "video" edge.
func HasVideo() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoWith applies the HasEdge predicate on the "video" edge with a given conditions (other predicates).
func HasVideoWith(preds ...predicate.Video) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VideoInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		p(s.Not())
	})
}
