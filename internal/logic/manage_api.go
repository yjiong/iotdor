package logic

import (
	"github.com/yjiong/iotdor/ent"
)

// MApi ...
type MApi interface {
	UserInfo(name string) (*ent.User, error)
	AddUser(name, passwd, group string, adminFlag bool) error
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
func (m *Manage) AddUser(name, passwd, group string, adminFlag bool) error {
	eg, _ := queryGroupByName(m.ctx, m.entC, group)
	_, err := addUser(m.ctx, m.entC, name, passwd, eg, adminFlag)
	return err
}
