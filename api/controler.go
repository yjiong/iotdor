package api

import (
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/utils"
)

// UserControler ....
type UserControler interface {
	UsersInfo() ([]*ent.User, error)
	UserInfo(uName string) (*ent.User, error)
	AddUser(uName, passwd string, adminFlag bool, phone ...string) error
	UpdateUser(uName, passwd string, adminFlag bool, phone ...string) error
	UpdateUserLoginInfo(uName, lip string) error
	UserAdminFlag(uName string) bool
	DelUser(uName string) error
}

// GatewayControler .....
type GatewayControler interface {
	AllGateways() []*ent.Gateway
}

// DeviceControler ....
type DeviceControler interface {
	AllDevices() []string
	DeviceRealTimeValue(devid string) map[string]string
	DeviceHistoryValue(devid string, qs utils.QuerySection) any
}

// OrganiTreeControler ....
type OrganiTreeControler interface {
	OrganizationTree() ([]*ent.OrganizationTree, error)
	ListOrganizationTreePositions(*ent.OrganizationTree) ([]*ent.OrganizationPosition, error)
	SubOrganizationFromID(id int) (eos []*ent.OrganizationTree, err error)
	CreateOrganizationSubNode(o ent.OrganizationTree) error
	AddOrganizationNode(id int, name, leftOrRight string) error
	UpdateOrganizationTree(o ent.OrganizationTree) error
	DeleteOrganizationTree(id int) error
	RelatePositioinToOranizationTree(id, posid string) error
	RemovePositioinFromOranizationTree(id, posid string) error
}

// OrganiPositionControler ....
type OrganiPositionControler interface {
	OrganizationPosition() ([]*ent.OrganizationPosition, error)
	ListOrganizationPositionDevices(*ent.OrganizationPosition) ([]*ent.Device, error)
	CreateOrganizationPosition(o ent.OrganizationPosition) error
	UpdateOrganizationPosition(o ent.OrganizationPosition) error
	DeleteOrganizationPosition(posid string) error
	AddDeviceToOrganizationPosition(posid, devid string) error
	SetPersonChargeWithOrganizationPosition(posid, uname string) error
	RemovePersonChargeWithOrganizationPosition(posid, uname string) error
	RemoveDeviceFromOrganizationPosition(posid, devid string) error
	QueryOrganizationPositionDevices(posid string) ([]string, error)
}

// ManageAPI ...
type ManageAPI interface {
	UserControler
	GatewayControler
	DeviceControler
	OrganiTreeControler
	OrganiPositionControler
}
