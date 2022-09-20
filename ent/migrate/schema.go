// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "dev_id", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeString},
		{Name: "conn", Type: field.TypeJSON},
		{Name: "read_interval", Type: field.TypeInt, Nullable: true},
		{Name: "store_interval", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "delete_flag", Type: field.TypeBool, Nullable: true},
		{Name: "summary", Type: field.TypeString, Nullable: true},
		{Name: "gateway_devices", Type: field.TypeInt, Nullable: true},
		{Name: "organization_position_devices", Type: field.TypeInt, Nullable: true},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_gateways_devices",
				Columns:    []*schema.Column{DevicesColumns[11]},
				RefColumns: []*schema.Column{GatewaysColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "devices_organization_positions_devices",
				Columns:    []*schema.Column{DevicesColumns[12]},
				RefColumns: []*schema.Column{OrganizationPositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "device_dev_id",
				Unique:  false,
				Columns: []*schema.Column{DevicesColumns[3]},
			},
		},
	}
	// GatewaysColumns holds the columns for the "gateways" table.
	GatewaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "gwid", Type: field.TypeString, Unique: true},
		{Name: "svrid", Type: field.TypeString},
		{Name: "broker", Type: field.TypeString},
		{Name: "installation_location", Type: field.TypeString, Nullable: true},
		{Name: "online", Type: field.TypeBool, Nullable: true},
		{Name: "delete_flag", Type: field.TypeBool, Nullable: true},
		{Name: "up_interval", Type: field.TypeInt, Default: 60},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "summary", Type: field.TypeString, Nullable: true},
		{Name: "group_gateways", Type: field.TypeInt, Nullable: true},
	}
	// GatewaysTable holds the schema information for the "gateways" table.
	GatewaysTable = &schema.Table{
		Name:       "gateways",
		Columns:    GatewaysColumns,
		PrimaryKey: []*schema.Column{GatewaysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "gateways_groups_gateways",
				Columns:    []*schema.Column{GatewaysColumns[12]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "gateway_gwid",
				Unique:  true,
				Columns: []*schema.Column{GatewaysColumns[3]},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "summary", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// OrganizationPositionsColumns holds the columns for the "organization_positions" table.
	OrganizationPositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "position_id", Type: field.TypeString, Unique: true},
		{Name: "address", Type: field.TypeString},
		{Name: "floor", Type: field.TypeString, Nullable: true},
		{Name: "unit_no", Type: field.TypeString, Nullable: true},
		{Name: "longitude_and_latitude", Type: field.TypeString},
		{Name: "summary", Type: field.TypeString, Nullable: true},
		{Name: "organization_tree_organization_positions", Type: field.TypeInt, Nullable: true},
	}
	// OrganizationPositionsTable holds the schema information for the "organization_positions" table.
	OrganizationPositionsTable = &schema.Table{
		Name:       "organization_positions",
		Columns:    OrganizationPositionsColumns,
		PrimaryKey: []*schema.Column{OrganizationPositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_positions_organization_trees_organization_positions",
				Columns:    []*schema.Column{OrganizationPositionsColumns[9]},
				RefColumns: []*schema.Column{OrganizationTreesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "organizationposition_position_id",
				Unique:  false,
				Columns: []*schema.Column{OrganizationPositionsColumns[3]},
			},
		},
	}
	// OrganizationTreesColumns holds the columns for the "organization_trees" table.
	OrganizationTreesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "parent_id", Type: field.TypeInt},
		{Name: "left", Type: field.TypeInt},
		{Name: "right", Type: field.TypeInt},
		{Name: "level", Type: field.TypeInt},
	}
	// OrganizationTreesTable holds the schema information for the "organization_trees" table.
	OrganizationTreesTable = &schema.Table{
		Name:       "organization_trees",
		Columns:    OrganizationTreesColumns,
		PrimaryKey: []*schema.Column{OrganizationTreesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "passwd", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "last_login_ip", Type: field.TypeString, Nullable: true},
		{Name: "last_login_time", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_group_id",
				Columns:    []*schema.Column{GroupUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_user_id",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupAdminsColumns holds the columns for the "group_admins" table.
	GroupAdminsColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupAdminsTable holds the schema information for the "group_admins" table.
	GroupAdminsTable = &schema.Table{
		Name:       "group_admins",
		Columns:    GroupAdminsColumns,
		PrimaryKey: []*schema.Column{GroupAdminsColumns[0], GroupAdminsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_admins_group_id",
				Columns:    []*schema.Column{GroupAdminsColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_admins_user_id",
				Columns:    []*schema.Column{GroupAdminsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OrganizationPositionPersonChargesColumns holds the columns for the "organization_position_person_charges" table.
	OrganizationPositionPersonChargesColumns = []*schema.Column{
		{Name: "organization_position_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// OrganizationPositionPersonChargesTable holds the schema information for the "organization_position_person_charges" table.
	OrganizationPositionPersonChargesTable = &schema.Table{
		Name:       "organization_position_person_charges",
		Columns:    OrganizationPositionPersonChargesColumns,
		PrimaryKey: []*schema.Column{OrganizationPositionPersonChargesColumns[0], OrganizationPositionPersonChargesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_position_person_charges_organization_position_id",
				Columns:    []*schema.Column{OrganizationPositionPersonChargesColumns[0]},
				RefColumns: []*schema.Column{OrganizationPositionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organization_position_person_charges_user_id",
				Columns:    []*schema.Column{OrganizationPositionPersonChargesColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DevicesTable,
		GatewaysTable,
		GroupsTable,
		OrganizationPositionsTable,
		OrganizationTreesTable,
		UsersTable,
		GroupUsersTable,
		GroupAdminsTable,
		OrganizationPositionPersonChargesTable,
	}
)

func init() {
	DevicesTable.ForeignKeys[0].RefTable = GatewaysTable
	DevicesTable.ForeignKeys[1].RefTable = OrganizationPositionsTable
	GatewaysTable.ForeignKeys[0].RefTable = GroupsTable
	OrganizationPositionsTable.ForeignKeys[0].RefTable = OrganizationTreesTable
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
	GroupAdminsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupAdminsTable.ForeignKeys[1].RefTable = UsersTable
	OrganizationPositionPersonChargesTable.ForeignKeys[0].RefTable = OrganizationPositionsTable
	OrganizationPositionPersonChargesTable.ForeignKeys[1].RefTable = UsersTable
}
