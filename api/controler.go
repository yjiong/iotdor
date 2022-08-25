package api

import "github.com/yjiong/iotdor/ent"

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
}

// OrganiTreeControler ....
type OrganiTreeControler interface {
	OrganizationTree() ([]ent.OrganizationTree, error)
	AddOrganizationTree(o ent.OrganizationTree) error
	//UpdateOrganizationTree(oName string, o ent.OrganizationTree) error
	//DeleteOrganizationTree(oName string) error
}

// OrganiPositionControler ....
type OrganiPositionControler interface {
	// OrganizationPosition ....
	//OrganizationPosition() ([]ent.OrganizationPosition, error)
	AddOrganizationPosition(o ent.OrganizationPosition) error
	UpdateOrganizationPosition(oName string, o ent.OrganizationPosition) error
	DeleteOrganizationPosition(oName string) error
	BeRelatedDeviceToOrganizationPosition(oName, devid string) error
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
