// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/yjiong/iotdor/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Floor applies equality check predicate on the "floor" field. It's identical to FloorEQ.
func Floor(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFloor), v))
	})
}

// UnitNo applies equality check predicate on the "unitNo" field. It's identical to UnitNoEQ.
func UnitNo(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnitNo), v))
	})
}

// LongitudeAndLatituude applies equality check predicate on the "longitudeAndLatituude" field. It's identical to LongitudeAndLatituudeEQ.
func LongitudeAndLatituude(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitudeAndLatituude), v))
	})
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// FloorEQ applies the EQ predicate on the "floor" field.
func FloorEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFloor), v))
	})
}

// FloorNEQ applies the NEQ predicate on the "floor" field.
func FloorNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFloor), v))
	})
}

// FloorIn applies the In predicate on the "floor" field.
func FloorIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFloor), v...))
	})
}

// FloorNotIn applies the NotIn predicate on the "floor" field.
func FloorNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFloor), v...))
	})
}

// FloorGT applies the GT predicate on the "floor" field.
func FloorGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFloor), v))
	})
}

// FloorGTE applies the GTE predicate on the "floor" field.
func FloorGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFloor), v))
	})
}

// FloorLT applies the LT predicate on the "floor" field.
func FloorLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFloor), v))
	})
}

// FloorLTE applies the LTE predicate on the "floor" field.
func FloorLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFloor), v))
	})
}

// FloorContains applies the Contains predicate on the "floor" field.
func FloorContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFloor), v))
	})
}

// FloorHasPrefix applies the HasPrefix predicate on the "floor" field.
func FloorHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFloor), v))
	})
}

// FloorHasSuffix applies the HasSuffix predicate on the "floor" field.
func FloorHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFloor), v))
	})
}

// FloorIsNil applies the IsNil predicate on the "floor" field.
func FloorIsNil() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFloor)))
	})
}

// FloorNotNil applies the NotNil predicate on the "floor" field.
func FloorNotNil() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFloor)))
	})
}

// FloorEqualFold applies the EqualFold predicate on the "floor" field.
func FloorEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFloor), v))
	})
}

// FloorContainsFold applies the ContainsFold predicate on the "floor" field.
func FloorContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFloor), v))
	})
}

// UnitNoEQ applies the EQ predicate on the "unitNo" field.
func UnitNoEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnitNo), v))
	})
}

// UnitNoNEQ applies the NEQ predicate on the "unitNo" field.
func UnitNoNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnitNo), v))
	})
}

// UnitNoIn applies the In predicate on the "unitNo" field.
func UnitNoIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnitNo), v...))
	})
}

// UnitNoNotIn applies the NotIn predicate on the "unitNo" field.
func UnitNoNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnitNo), v...))
	})
}

// UnitNoGT applies the GT predicate on the "unitNo" field.
func UnitNoGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnitNo), v))
	})
}

// UnitNoGTE applies the GTE predicate on the "unitNo" field.
func UnitNoGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnitNo), v))
	})
}

// UnitNoLT applies the LT predicate on the "unitNo" field.
func UnitNoLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnitNo), v))
	})
}

// UnitNoLTE applies the LTE predicate on the "unitNo" field.
func UnitNoLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnitNo), v))
	})
}

// UnitNoContains applies the Contains predicate on the "unitNo" field.
func UnitNoContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUnitNo), v))
	})
}

// UnitNoHasPrefix applies the HasPrefix predicate on the "unitNo" field.
func UnitNoHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUnitNo), v))
	})
}

// UnitNoHasSuffix applies the HasSuffix predicate on the "unitNo" field.
func UnitNoHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUnitNo), v))
	})
}

// UnitNoIsNil applies the IsNil predicate on the "unitNo" field.
func UnitNoIsNil() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUnitNo)))
	})
}

// UnitNoNotNil applies the NotNil predicate on the "unitNo" field.
func UnitNoNotNil() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUnitNo)))
	})
}

// UnitNoEqualFold applies the EqualFold predicate on the "unitNo" field.
func UnitNoEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUnitNo), v))
	})
}

// UnitNoContainsFold applies the ContainsFold predicate on the "unitNo" field.
func UnitNoContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUnitNo), v))
	})
}

// LongitudeAndLatituudeEQ applies the EQ predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeNEQ applies the NEQ predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeIn applies the In predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLongitudeAndLatituude), v...))
	})
}

// LongitudeAndLatituudeNotIn applies the NotIn predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLongitudeAndLatituude), v...))
	})
}

// LongitudeAndLatituudeGT applies the GT predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeGTE applies the GTE predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeLT applies the LT predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeLTE applies the LTE predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeContains applies the Contains predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeHasPrefix applies the HasPrefix predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeHasSuffix applies the HasSuffix predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeEqualFold applies the EqualFold predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLongitudeAndLatituude), v))
	})
}

// LongitudeAndLatituudeContainsFold applies the ContainsFold predicate on the "longitudeAndLatituude" field.
func LongitudeAndLatituudeContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLongitudeAndLatituude), v))
	})
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSummary), v))
	})
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSummary), v...))
	})
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSummary), v...))
	})
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSummary), v))
	})
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSummary), v))
	})
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSummary), v))
	})
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSummary), v))
	})
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSummary), v))
	})
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSummary), v))
	})
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSummary), v))
	})
}

// SummaryIsNil applies the IsNil predicate on the "summary" field.
func SummaryIsNil() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSummary)))
	})
}

// SummaryNotNil applies the NotNil predicate on the "summary" field.
func SummaryNotNil() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSummary)))
	})
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSummary), v))
	})
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSummary), v))
	})
}

// HasDevices applies the HasEdge predicate on the "devices" edge.
func HasDevices() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DevicesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDevicesWith applies the HasEdge predicate on the "devices" edge with a given conditions (other predicates).
func HasDevicesWith(preds ...predicate.Device) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DevicesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonCharges applies the HasEdge predicate on the "personCharges" edge.
func HasPersonCharges() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonChargesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PersonChargesTable, PersonChargesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonChargesWith applies the HasEdge predicate on the "personCharges" edge with a given conditions (other predicates).
func HasPersonChargesWith(preds ...predicate.User) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonChargesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PersonChargesTable, PersonChargesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
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
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		p(s.Not())
	})
}
