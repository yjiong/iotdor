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
	SubOrganizationFromID(id int) (eos []*ent.OrganizationTree, err error)
	CreateOrganizationSubNode(o ent.OrganizationTree) error
	AddOrganizationNode(id int, name, leftOrRight string) error
	UpdateOrganizationTree(o ent.OrganizationTree) error
	DeleteOrganizationTree(id int) error
	RelatePositioinToOranizationTree(id, posid int) error
	RemovePositioinFromOranizationTree(id, posid int) error
}

// OrganiPositionControler ....
type OrganiPositionControler interface {
	// OrganizationPosition ....
	//OrganizationPosition() ([]ent.OrganizationPosition, error)
	CreateOrganizationPosition(o ent.OrganizationPosition) error
	UpdateOrganizationPosition(oName string, o ent.OrganizationPosition) error
	DeleteOrganizationPosition(oName string) error
	AddDeviceToOrganizationPosition(oName, devid string) error
	RemoveDeviceFromOrganizationPosition(devid string) error
	QueryOrganizationPositionDevices(oName string) ([]string, error)
}

// ManageAPI ...
type ManageAPI interface {
	UserControler
	GatewayControler
	DeviceControler
	OrganiTreeControler
	OrganiPositionControler
}
