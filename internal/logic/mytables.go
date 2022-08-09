package logic

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	//AmmtersColumns .....
	AmmtersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "devid", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeString, Unique: true},
		{Name: "timestamp", Type: field.TypeString},
		{Name: "总有功功率", Type: field.TypeFloat32},
	}
	// AmmtersTable holds the schema information for the "gateways" table.
	AmmtersTable = &schema.Table{
		Name:       "ammters",
		Columns:    AmmtersColumns,
		PrimaryKey: []*schema.Column{AmmtersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "ammter_devid",
				Unique:  true,
				Columns: []*schema.Column{AmmtersColumns[3]},
			},
		},
	}
)
