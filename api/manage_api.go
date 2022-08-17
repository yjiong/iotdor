package api

import (
	"github.com/yjiong/iotdor/ent"
)

// ManageAPI ...
type ManageAPI interface {
	UserInfo(name string) (*ent.User, error)
	AddUser(name, passwd, group string, adminFlag bool, phone ...string) error
	UpdateUser(name, passwd, group string, adminFlag bool, phone ...string) error
}
