// Code generated by ent, DO NOT EDIT.

package channel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/zibbp/eos/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldDescription, v))
}

// ImagePath applies equality check predicate on the "image_path" field. It's identical to ImagePathEQ.
func ImagePath(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldImagePath, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Channel {
	return predicate.Channel(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Channel {
	return predicate.Channel(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldDescription, v))
}

// ImagePathEQ applies the EQ predicate on the "image_path" field.
func ImagePathEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldImagePath, v))
}

// ImagePathNEQ applies the NEQ predicate on the "image_path" field.
func ImagePathNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldImagePath, v))
}

// ImagePathIn applies the In predicate on the "image_path" field.
func ImagePathIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldImagePath, vs...))
}

// ImagePathNotIn applies the NotIn predicate on the "image_path" field.
func ImagePathNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldImagePath, vs...))
}

// ImagePathGT applies the GT predicate on the "image_path" field.
func ImagePathGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldImagePath, v))
}

// ImagePathGTE applies the GTE predicate on the "image_path" field.
func ImagePathGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldImagePath, v))
}

// ImagePathLT applies the LT predicate on the "image_path" field.
func ImagePathLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldImagePath, v))
}

// ImagePathLTE applies the LTE predicate on the "image_path" field.
func ImagePathLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldImagePath, v))
}

// ImagePathContains applies the Contains predicate on the "image_path" field.
func ImagePathContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldImagePath, v))
}

// ImagePathHasPrefix applies the HasPrefix predicate on the "image_path" field.
func ImagePathHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldImagePath, v))
}

// ImagePathHasSuffix applies the HasSuffix predicate on the "image_path" field.
func ImagePathHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldImagePath, v))
}

// ImagePathIsNil applies the IsNil predicate on the "image_path" field.
func ImagePathIsNil() predicate.Channel {
	return predicate.Channel(sql.FieldIsNull(FieldImagePath))
}

// ImagePathNotNil applies the NotNil predicate on the "image_path" field.
func ImagePathNotNil() predicate.Channel {
	return predicate.Channel(sql.FieldNotNull(FieldImagePath))
}

// ImagePathEqualFold applies the EqualFold predicate on the "image_path" field.
func ImagePathEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldImagePath, v))
}

// ImagePathContainsFold applies the ContainsFold predicate on the "image_path" field.
func ImagePathContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldImagePath, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasVideos applies the HasEdge predicate on the "videos" edge.
func HasVideos() predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VideosTable, VideosColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideosWith applies the HasEdge predicate on the "videos" edge with a given conditions (other predicates).
func HasVideosWith(preds ...predicate.Video) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VideosInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VideosTable, VideosColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
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
func Not(p predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		p(s.Not())
	})
}