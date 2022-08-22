package logic

import (
	"fmt"
	"strings"
	"time"

	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/organization"
	"github.com/yjiong/iotdor/ent/user"
)

// AllDevices ....
func (m *Manage) AllDevices() (ids []string) {
	keys, _ := m.redisC.Keys(m.ctx, fmt.Sprintf("%s:*:%s", strings.ToUpper(m.iotdName), DeviceValue)).Result()
	for _, key := range keys {
		kids := strings.Split(key, ":")
		if len(kids) > 3 {
			ids = append(ids, kids[2])
		}
	}
	return
}

// DeviceRealTimeValue ....
func (m *Manage) DeviceRealTimeValue(devid string) map[string]string {
	keys, _ := m.redisC.Keys(m.ctx, fmt.Sprintf("*:%s:*", devid)).Result()
	vs, _ := m.redisC.HGetAll(m.ctx, keys[0]).Result()
	return vs
}

// UsersInfo for api
func (m *Manage) UsersInfo() (us []*ent.User, e error) {
	if eg, err := queryGroupByName(m.ctx, m.entC, m.iotdName); err != nil {
		e = err
	} else {
		us, _ = eg.QueryUsers().All(m.ctx)
	}
	return
}

// UserInfo for api
func (m *Manage) UserInfo(name string) (u *ent.User, e error) {
	if us, err := queryUsers(m.ctx, m.entC); err != nil {
		e = err
	} else {
		for _, eu := range us {
			if eu.Name == name {
				u = eu
			}
		}
	}
	return
}

// AddUser ...
func (m *Manage) AddUser(name, passwd string, adminFlag bool, phone ...string) error {
	eg, _ := queryGroupByName(m.ctx, m.entC, m.iotdName)
	_, err := addUser(m.ctx, m.entC, name, passwd, eg, adminFlag, phone...)
	return err
}

// DelUser ...
func (m *Manage) DelUser(name string) error {
	_, err := m.entC.User.Delete().Where(user.Name(name)).Exec(m.ctx)
	return err
}

// UpdateUser ...
func (m *Manage) UpdateUser(name, passwd string, adminFlag bool, phone ...string) error {
	eg, _ := queryGroupByName(m.ctx, m.entC, m.iotdName)
	err := updateUser(m.ctx, m.entC, name, passwd, eg, adminFlag, phone...)
	return err
}

// UpdateUserLoginInfo ...
func (m *Manage) UpdateUserLoginInfo(name, lip string) error {
	return m.entC.User.Update().Where(user.Name(name)).
		SetLastLoginTime(time.Now()).
		SetLastLoginIP(lip).Exec(m.ctx)
}

// UserAdminFlag ....
func (m *Manage) UserAdminFlag(uname string) bool {
	u, _ := queryUserByName(m.ctx, m.entC, uname)
	return userAdminFlag(m.ctx, u, uname)
}

// OrganizationInfo ...
func (m *Manage) OrganizationInfo() ([]*ent.Organization, error) {
	return m.entC.Organization.Query().All(m.ctx)
}

// AddOrganization .....
func (m *Manage) AddOrganization(o ent.Organization) error {
	return m.entC.Organization.Create().
		SetFloor(o.Floor).
		SetAddress(o.Address).
		SetName(o.Name).
		SetUnitNo(o.UnitNo).
		SetLongitudeAndLatituude(o.LongitudeAndLatituude).
		SetSummary(o.Summary).Exec(m.ctx)
}

// UpdateOrganization .....
func (m *Manage) UpdateOrganization(name string, o ent.Organization) error {
	return m.entC.Organization.Update().
		Where(organization.Name(name)).
		SetFloor(o.Floor).
		SetAddress(o.Address).
		SetUnitNo(o.UnitNo).
		SetLongitudeAndLatituude(o.LongitudeAndLatituude).
		SetSummary(o.Summary).Exec(m.ctx)
}

// DeleteOrganization .....
func (m *Manage) DeleteOrganization(name string) error {
	_, err := m.entC.Organization.Delete().Where(organization.Name(name)).Exec(m.ctx)
	return err
}

// QueryOrganizationDevices ....
func (m *Manage) QueryOrganizationDevices(name string) (ds []string, err error) {
	var o *ent.Organization
	o, err = m.entC.Organization.Query().Where(organization.Name(name)).Only(m.ctx)
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

// BeRelatedDeviceToOrganization .....
func (m *Manage) BeRelatedDeviceToOrganization(name, devid string) error {
	var rerr error
	o, oerr := m.entC.Organization.Query().Where(organization.Name(name)).Only(m.ctx)
	if oerr == nil {
		d, derr := queryDeviceByDevID(m.ctx, m.entC, devid)
		if derr == nil {
			rerr = d.Update().SetOrganization(o).Exec(m.ctx)
		} else {
			rerr = derr
		}
	} else {
		rerr = oerr
	}
	return rerr
}

// RemoveDeviceFromOrganization .....
func (m *Manage) RemoveDeviceFromOrganization(devid string) error {
	var rerr error
	d, derr := queryDeviceByDevID(m.ctx, m.entC, devid)
	if derr == nil {
		rerr = d.Update().ClearOrganization().Exec(m.ctx)
	} else {
		rerr = derr
	}
	return rerr
}

// SetPersonChargeWithOrganization .....
func (m *Manage) SetPersonChargeWithOrganization(oname, uname string) error {
	o, oerr := m.entC.Organization.Query().Where(organization.Name(oname)).Only(m.ctx)
	u, uerr := m.entC.User.Query().Where(user.Name(uname)).Only(m.ctx)
	if oerr == nil && uerr == nil {
		return u.Update().AddPersonCharges(o).Exec(m.ctx)
	}
	return fmt.Errorf("%v-%v", oerr, uerr)
}

// RemovePersonChargeWithOrganization .....
func (m *Manage) RemovePersonChargeWithOrganization(oname, uname string) error {
	o, oerr := m.entC.Organization.Query().Where(organization.Name(oname)).Only(m.ctx)
	u, uerr := m.entC.User.Query().Where(user.Name(uname)).Only(m.ctx)
	if oerr == nil && uerr == nil {
		return o.Update().RemovePersonCharges(u).Exec(m.ctx)
	}
	return fmt.Errorf("%v-%v", oerr, uerr)
}
