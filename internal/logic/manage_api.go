package logic

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/user"
	"github.com/yjiong/iotdor/utils"
)

// AllGateways ....
func (m *Manage) AllGateways() []*ent.Gateway {
	gws, _ := m.entC.Gateway.Query().All(m.ctx)
	return gws
}

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

// DeviceHistoryValue ....
func (m *Manage) DeviceHistoryValue(devid string, qs utils.QuerySection) any {
	var pages int64
	cstr := fmt.Sprintf("select count(*) from %s where devid='%s' and timestamp between '%s' and '%s';",
		DEVICETABLE,
		devid,
		qs.Since.Local().Format(TFORMAT),
		qs.Until.Local().Format(TFORMAT))
	_, counts, _ := RawQuery(m.DB, cstr)
	if len(counts) > 0 {
		if c, err := strconv.ParseInt(counts[0], 10, 64); err == nil {
			pages = utils.CalcPages(c, qs.PageSize)
		}
	}
	qstr := fmt.Sprintf("select * from %s where devid='%s' and timestamp between '%s' and '%s' limit %d offset %d;",
		DEVICETABLE,
		devid,
		qs.Since.Local().Format(TFORMAT),
		qs.Until.Local().Format(TFORMAT),
		qs.PageSize,
		qs.PageSize*qs.PagesIndex)
	columns, values, _ := RawQuery(m.DB, qstr)
	return map[string]any{
		"Columns": columns,
		"pages":   pages,
		"datas":   values,
	}
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
