// Code generated by entc, DO NOT EDIT.

package gateway

import (
	"dcmbroker/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Gwid applies equality check predicate on the "gwid" field. It's identical to GwidEQ.
func Gwid(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGwid), v))
	})
}

// Broker applies equality check predicate on the "broker" field. It's identical to BrokerEQ.
func Broker(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBroker), v))
	})
}

// InstallationLocation applies equality check predicate on the "installationLocation" field. It's identical to InstallationLocationEQ.
func InstallationLocation(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInstallationLocation), v))
	})
}

// Online applies equality check predicate on the "online" field. It's identical to OnlineEQ.
func Online(v bool) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnline), v))
	})
}

// IdDelete applies equality check predicate on the "idDelete" field. It's identical to IdDeleteEQ.
func IdDelete(v bool) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdDelete), v))
	})
}

// UpInterval applies equality check predicate on the "upInterval" field. It's identical to UpIntervalEQ.
func UpInterval(v int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpInterval), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// GwidEQ applies the EQ predicate on the "gwid" field.
func GwidEQ(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGwid), v))
	})
}

// GwidNEQ applies the NEQ predicate on the "gwid" field.
func GwidNEQ(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGwid), v))
	})
}

// GwidIn applies the In predicate on the "gwid" field.
func GwidIn(vs ...string) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGwid), v...))
	})
}

// GwidNotIn applies the NotIn predicate on the "gwid" field.
func GwidNotIn(vs ...string) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGwid), v...))
	})
}

// GwidGT applies the GT predicate on the "gwid" field.
func GwidGT(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGwid), v))
	})
}

// GwidGTE applies the GTE predicate on the "gwid" field.
func GwidGTE(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGwid), v))
	})
}

// GwidLT applies the LT predicate on the "gwid" field.
func GwidLT(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGwid), v))
	})
}

// GwidLTE applies the LTE predicate on the "gwid" field.
func GwidLTE(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGwid), v))
	})
}

// GwidContains applies the Contains predicate on the "gwid" field.
func GwidContains(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGwid), v))
	})
}

// GwidHasPrefix applies the HasPrefix predicate on the "gwid" field.
func GwidHasPrefix(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGwid), v))
	})
}

// GwidHasSuffix applies the HasSuffix predicate on the "gwid" field.
func GwidHasSuffix(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGwid), v))
	})
}

// GwidEqualFold applies the EqualFold predicate on the "gwid" field.
func GwidEqualFold(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGwid), v))
	})
}

// GwidContainsFold applies the ContainsFold predicate on the "gwid" field.
func GwidContainsFold(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGwid), v))
	})
}

// BrokerEQ applies the EQ predicate on the "broker" field.
func BrokerEQ(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBroker), v))
	})
}

// BrokerNEQ applies the NEQ predicate on the "broker" field.
func BrokerNEQ(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBroker), v))
	})
}

// BrokerIn applies the In predicate on the "broker" field.
func BrokerIn(vs ...string) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBroker), v...))
	})
}

// BrokerNotIn applies the NotIn predicate on the "broker" field.
func BrokerNotIn(vs ...string) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBroker), v...))
	})
}

// BrokerGT applies the GT predicate on the "broker" field.
func BrokerGT(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBroker), v))
	})
}

// BrokerGTE applies the GTE predicate on the "broker" field.
func BrokerGTE(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBroker), v))
	})
}

// BrokerLT applies the LT predicate on the "broker" field.
func BrokerLT(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBroker), v))
	})
}

// BrokerLTE applies the LTE predicate on the "broker" field.
func BrokerLTE(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBroker), v))
	})
}

// BrokerContains applies the Contains predicate on the "broker" field.
func BrokerContains(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBroker), v))
	})
}

// BrokerHasPrefix applies the HasPrefix predicate on the "broker" field.
func BrokerHasPrefix(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBroker), v))
	})
}

// BrokerHasSuffix applies the HasSuffix predicate on the "broker" field.
func BrokerHasSuffix(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBroker), v))
	})
}

// BrokerEqualFold applies the EqualFold predicate on the "broker" field.
func BrokerEqualFold(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBroker), v))
	})
}

// BrokerContainsFold applies the ContainsFold predicate on the "broker" field.
func BrokerContainsFold(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBroker), v))
	})
}

// InstallationLocationEQ applies the EQ predicate on the "installationLocation" field.
func InstallationLocationEQ(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationNEQ applies the NEQ predicate on the "installationLocation" field.
func InstallationLocationNEQ(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationIn applies the In predicate on the "installationLocation" field.
func InstallationLocationIn(vs ...string) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInstallationLocation), v...))
	})
}

// InstallationLocationNotIn applies the NotIn predicate on the "installationLocation" field.
func InstallationLocationNotIn(vs ...string) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInstallationLocation), v...))
	})
}

// InstallationLocationGT applies the GT predicate on the "installationLocation" field.
func InstallationLocationGT(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationGTE applies the GTE predicate on the "installationLocation" field.
func InstallationLocationGTE(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationLT applies the LT predicate on the "installationLocation" field.
func InstallationLocationLT(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationLTE applies the LTE predicate on the "installationLocation" field.
func InstallationLocationLTE(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationContains applies the Contains predicate on the "installationLocation" field.
func InstallationLocationContains(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationHasPrefix applies the HasPrefix predicate on the "installationLocation" field.
func InstallationLocationHasPrefix(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationHasSuffix applies the HasSuffix predicate on the "installationLocation" field.
func InstallationLocationHasSuffix(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationIsNil applies the IsNil predicate on the "installationLocation" field.
func InstallationLocationIsNil() predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInstallationLocation)))
	})
}

// InstallationLocationNotNil applies the NotNil predicate on the "installationLocation" field.
func InstallationLocationNotNil() predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInstallationLocation)))
	})
}

// InstallationLocationEqualFold applies the EqualFold predicate on the "installationLocation" field.
func InstallationLocationEqualFold(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInstallationLocation), v))
	})
}

// InstallationLocationContainsFold applies the ContainsFold predicate on the "installationLocation" field.
func InstallationLocationContainsFold(v string) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInstallationLocation), v))
	})
}

// OnlineEQ applies the EQ predicate on the "online" field.
func OnlineEQ(v bool) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnline), v))
	})
}

// OnlineNEQ applies the NEQ predicate on the "online" field.
func OnlineNEQ(v bool) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOnline), v))
	})
}

// IdDeleteEQ applies the EQ predicate on the "idDelete" field.
func IdDeleteEQ(v bool) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdDelete), v))
	})
}

// IdDeleteNEQ applies the NEQ predicate on the "idDelete" field.
func IdDeleteNEQ(v bool) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdDelete), v))
	})
}

// UpIntervalEQ applies the EQ predicate on the "upInterval" field.
func UpIntervalEQ(v int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpInterval), v))
	})
}

// UpIntervalNEQ applies the NEQ predicate on the "upInterval" field.
func UpIntervalNEQ(v int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpInterval), v))
	})
}

// UpIntervalIn applies the In predicate on the "upInterval" field.
func UpIntervalIn(vs ...int) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpInterval), v...))
	})
}

// UpIntervalNotIn applies the NotIn predicate on the "upInterval" field.
func UpIntervalNotIn(vs ...int) predicate.Gateway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gateway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpInterval), v...))
	})
}

// UpIntervalGT applies the GT predicate on the "upInterval" field.
func UpIntervalGT(v int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpInterval), v))
	})
}

// UpIntervalGTE applies the GTE predicate on the "upInterval" field.
func UpIntervalGTE(v int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpInterval), v))
	})
}

// UpIntervalLT applies the LT predicate on the "upInterval" field.
func UpIntervalLT(v int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpInterval), v))
	})
}

// UpIntervalLTE applies the LTE predicate on the "upInterval" field.
func UpIntervalLTE(v int) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpInterval), v))
	})
}

// HasDevices applies the HasEdge predicate on the "devices" edge.
func HasDevices() predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DevicesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDevicesWith applies the HasEdge predicate on the "devices" edge with a given conditions (other predicates).
func HasDevicesWith(preds ...predicate.Device) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
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

// HasBelong applies the HasEdge predicate on the "belong" edge.
func HasBelong() predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BelongTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BelongTable, BelongColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBelongWith applies the HasEdge predicate on the "belong" edge with a given conditions (other predicates).
func HasBelongWith(preds ...predicate.User) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BelongInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BelongTable, BelongColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Gateway) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Gateway) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
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
func Not(p predicate.Gateway) predicate.Gateway {
	return predicate.Gateway(func(s *sql.Selector) {
		p(s.Not())
	})
}
