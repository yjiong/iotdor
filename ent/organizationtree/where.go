// Code generated by ent, DO NOT EDIT.

package organizationtree

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/yjiong/iotdor/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentID), v))
	})
}

// Left applies equality check predicate on the "left" field. It's identical to LeftEQ.
func Left(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeft), v))
	})
}

// Right applies equality check predicate on the "right" field. It's identical to RightEQ.
func Right(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRight), v))
	})
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentID), v))
	})
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParentID), v))
	})
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldParentID), v...))
	})
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldParentID), v...))
	})
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParentID), v))
	})
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParentID), v))
	})
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParentID), v))
	})
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParentID), v))
	})
}

// LeftEQ applies the EQ predicate on the "left" field.
func LeftEQ(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeft), v))
	})
}

// LeftNEQ applies the NEQ predicate on the "left" field.
func LeftNEQ(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLeft), v))
	})
}

// LeftIn applies the In predicate on the "left" field.
func LeftIn(vs ...int) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLeft), v...))
	})
}

// LeftNotIn applies the NotIn predicate on the "left" field.
func LeftNotIn(vs ...int) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLeft), v...))
	})
}

// LeftGT applies the GT predicate on the "left" field.
func LeftGT(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLeft), v))
	})
}

// LeftGTE applies the GTE predicate on the "left" field.
func LeftGTE(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLeft), v))
	})
}

// LeftLT applies the LT predicate on the "left" field.
func LeftLT(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLeft), v))
	})
}

// LeftLTE applies the LTE predicate on the "left" field.
func LeftLTE(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLeft), v))
	})
}

// RightEQ applies the EQ predicate on the "right" field.
func RightEQ(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRight), v))
	})
}

// RightNEQ applies the NEQ predicate on the "right" field.
func RightNEQ(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRight), v))
	})
}

// RightIn applies the In predicate on the "right" field.
func RightIn(vs ...int) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRight), v...))
	})
}

// RightNotIn applies the NotIn predicate on the "right" field.
func RightNotIn(vs ...int) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRight), v...))
	})
}

// RightGT applies the GT predicate on the "right" field.
func RightGT(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRight), v))
	})
}

// RightGTE applies the GTE predicate on the "right" field.
func RightGTE(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRight), v))
	})
}

// RightLT applies the LT predicate on the "right" field.
func RightLT(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRight), v))
	})
}

// RightLTE applies the LTE predicate on the "right" field.
func RightLTE(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRight), v))
	})
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLevel), v))
	})
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLevel), v...))
	})
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.OrganizationTree {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLevel), v...))
	})
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLevel), v))
	})
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLevel), v))
	})
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLevel), v))
	})
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLevel), v))
	})
}

// HasOrganizationPositions applies the HasEdge predicate on the "organization_positions" edge.
func HasOrganizationPositions() predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationPositionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationPositionsTable, OrganizationPositionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationPositionsWith applies the HasEdge predicate on the "organization_positions" edge with a given conditions (other predicates).
func HasOrganizationPositionsWith(preds ...predicate.OrganizationPosition) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationPositionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationPositionsTable, OrganizationPositionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrganizationTree) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrganizationTree) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
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
func Not(p predicate.OrganizationTree) predicate.OrganizationTree {
	return predicate.OrganizationTree(func(s *sql.Selector) {
		p(s.Not())
	})
}