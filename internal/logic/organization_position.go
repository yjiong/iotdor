package logic

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/device"
	"github.com/yjiong/iotdor/ent/organizationposition"
	"github.com/yjiong/iotdor/ent/user"
	"github.com/yjiong/iotdor/utils"
)

// OrganizationPosition ...
func (m *Manage) OrganizationPosition() ([]*ent.OrganizationPosition, error) {
	return m.entC.OrganizationPosition.Query().All(m.ctx)
}

// ListOrganizationPositionDevices ....
func (m *Manage) ListOrganizationPositionDevices(posid string) ([]*ent.Device, error) {
	return m.entC.OrganizationPosition.Query().
		Where(organizationposition.PositionID(posid)).
		QueryDevices().All(m.ctx)
}

// ListOrganizationPositionCharges ....
func (m *Manage) ListOrganizationPositionCharges(posid string) ([]*ent.User, error) {
	return m.entC.OrganizationPosition.Query().
		Where(organizationposition.PositionID(posid)).
		QueryPersonCharges().All(m.ctx)
}

// CreateOrganizationPosition .....
func (m *Manage) CreateOrganizationPosition(o ent.OrganizationPosition) error {
	opid := utils.GetSnowflakeID()
	return m.entC.OrganizationPosition.Create().
		SetPositionID(opid).
		SetFloor(o.Floor).
		SetAddress(o.Address).
		SetUnitNo(o.UnitNo).
		SetLongitudeAndLatitude(o.LongitudeAndLatitude).
		SetSummary(o.Summary).Exec(m.ctx)
}

// UpdateOrganizationPosition .....
func (m *Manage) UpdateOrganizationPosition(o ent.OrganizationPosition) error {
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

// AddDeviceToOrganizationPosition .....
func (m *Manage) AddDeviceToOrganizationPosition(posid, devid string) error {
	var rerr error
	o, oerr := m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	if oerr == nil {
		d, derr := m.entC.Device.Query().Where(device.DevID(devid)).Only(m.ctx)
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
func (m *Manage) RemoveDeviceFromOrganizationPosition(posid, devid string) error {
	d, err := m.entC.Device.Query().
		Where(device.DevID(devid),
			device.HasOrganizationPositionWith(organizationposition.PositionID(posid))).
		Only(m.ctx)
	if d != nil && err == nil {
		return m.entC.OrganizationPosition.Update().
			Where(organizationposition.PositionID(posid)).
			RemoveDeviceIDs(d.ID).
			Exec(m.ctx)
	}
	return errors.Wrap(err, fmt.Sprintf("%s not exist", devid))
}

// SetPersonChargeWithOrganizationPosition .....
func (m *Manage) SetPersonChargeWithOrganizationPosition(posid, uname string) error {
	o, oerr := m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	u, uerr := m.entC.User.Query().Where(user.Name(uname)).Only(m.ctx)
	if oerr == nil && uerr == nil {
		//return u.Update().AddPersonCharges(o).Exec(m.ctx)
		return o.Update().AddPersonCharges(u).Exec(m.ctx)
	}
	return fmt.Errorf("%v-%v", oerr, uerr)
}

// RemovePersonChargeWithOrganizationPosition .....
func (m *Manage) RemovePersonChargeWithOrganizationPosition(posid, uname string) error {
	o, oerr := m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	u, uerr := o.QueryPersonCharges().Where(user.Name(uname)).Only(m.ctx)
	if oerr == nil && uerr == nil {
		return o.Update().RemovePersonChargeIDs(u.ID).Exec(m.ctx)
	}
	return fmt.Errorf("%v-%v", oerr, uerr)
}
