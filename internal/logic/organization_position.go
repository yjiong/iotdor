package logic

import (
	"fmt"

	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/organizationposition"
	"github.com/yjiong/iotdor/ent/user"
)

// OrganizationPosition ...
func (m *Manage) OrganizationPosition() ([]*ent.OrganizationPosition, error) {
	return m.entC.OrganizationPosition.Query().All(m.ctx)
}

// AddOrganizationPosition .....
func (m *Manage) AddOrganizationPosition(o ent.OrganizationPosition) error {
	return m.entC.OrganizationPosition.Create().
		SetPositionID(o.PositionID).
		SetFloor(o.Floor).
		SetAddress(o.Address).
		SetUnitNo(o.UnitNo).
		SetLongitudeAndLatitude(o.LongitudeAndLatitude).
		SetSummary(o.Summary).Exec(m.ctx)
}

// UpdateOrganizationPosition .....
func (m *Manage) UpdateOrganizationPosition(posid string, o ent.OrganizationPosition) error {
	return m.entC.OrganizationPosition.Update().
		Where(organizationposition.PositionID(o.PositionID)).
		SetFloor(o.Floor).
		SetAddress(o.Address).
		SetUnitNo(o.UnitNo).
		SetLongitudeAndLatitude(o.LongitudeAndLatitude).
		SetSummary(o.Summary).Exec(m.ctx)
}

// DeleteOrganizationPosition .....
func (m *Manage) DeleteOrganizationPosition(posid string) error {
	_, err := m.entC.OrganizationPosition.Delete().Where(organizationposition.PositionID(posid)).Exec(m.ctx)
	return err
}

// QueryOrganizationPositionDevices ....
func (m *Manage) QueryOrganizationPositionDevices(posid string) (ds []string, err error) {
	var o *ent.OrganizationPosition
	o, err = m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	if err == nil {
		eds, _ := o.QueryDevices().All(m.ctx)
		if len(eds) > 0 {
			for _, ed := range eds {
				ds = append(ds, ed.DevID)
			}
		}
	}
	return
}

// BeRelatedDeviceToOrganizationPosition .....
func (m *Manage) BeRelatedDeviceToOrganizationPosition(posid, devid string) error {
	var rerr error
	o, oerr := m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	if oerr == nil {
		d, derr := queryDeviceByDevID(m.ctx, m.entC, devid)
		if derr == nil {
			rerr = d.Update().SetOrganizationPosition(o).Exec(m.ctx)
		} else {
			rerr = derr
		}
	} else {
		rerr = oerr
	}
	return rerr
}

// RemoveDeviceFromOrganizationPosition .....
func (m *Manage) RemoveDeviceFromOrganizationPosition(devid string) error {
	var rerr error
	d, derr := queryDeviceByDevID(m.ctx, m.entC, devid)
	if derr == nil {
		rerr = d.Update().ClearOrganizationPosition().Exec(m.ctx)
	} else {
		rerr = derr
	}
	return rerr
}

// SetPersonChargeWithOrganizationPosition .....
func (m *Manage) SetPersonChargeWithOrganizationPosition(posid, uname string) error {
	o, oerr := m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	u, uerr := m.entC.User.Query().Where(user.Name(uname)).Only(m.ctx)
	if oerr == nil && uerr == nil {
		return u.Update().AddPersonCharges(o).Exec(m.ctx)
	}
	return fmt.Errorf("%v-%v", oerr, uerr)
}

// RemovePersonChargeWithOrganizationPosition .....
func (m *Manage) RemovePersonChargeWithOrganizationPosition(posid, uname string) error {
	o, oerr := m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	u, uerr := m.entC.User.Query().Where(user.Name(uname)).Only(m.ctx)
	if oerr == nil && uerr == nil {
		return o.Update().RemovePersonCharges(u).Exec(m.ctx)
	}
	return fmt.Errorf("%v-%v", oerr, uerr)
}
