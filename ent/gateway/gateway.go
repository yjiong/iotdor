// Code generated by ent, DO NOT EDIT.

package gateway

import (
	"time"
)

const (
	// Label holds the string label denoting the gateway type in the database.
	Label = "gateway"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldGwid holds the string denoting the gwid field in the database.
	FieldGwid = "gwid"
	// FieldSvrid holds the string denoting the svrid field in the database.
	FieldSvrid = "svrid"
	// FieldBroker holds the string denoting the broker field in the database.
	FieldBroker = "broker"
	// FieldInstallationLocation holds the string denoting the installationlocation field in the database.
	FieldInstallationLocation = "installation_location"
	// FieldOnline holds the string denoting the online field in the database.
	FieldOnline = "online"
	// FieldDeleteFlag holds the string denoting the deleteflag field in the database.
	FieldDeleteFlag = "delete_flag"
	// FieldUpInterval holds the string denoting the upinterval field in the database.
	FieldUpInterval = "up_interval"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// EdgeDevices holds the string denoting the devices edge name in mutations.
	EdgeDevices = "devices"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the gateway in the database.
	Table = "gateways"
	// DevicesTable is the table that holds the devices relation/edge.
	DevicesTable = "devices"
	// DevicesInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DevicesInverseTable = "devices"
	// DevicesColumn is the table column denoting the devices relation/edge.
	DevicesColumn = "gateway_devices"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "gateways"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_gateways"
)

// Columns holds all SQL columns for gateway fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldGwid,
	FieldSvrid,
	FieldBroker,
	FieldInstallationLocation,
	FieldOnline,
	FieldDeleteFlag,
	FieldUpInterval,
	FieldSummary,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "gateways"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_gateways",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultUpInterval holds the default value on creation for the "upInterval" field.
	DefaultUpInterval int
)
