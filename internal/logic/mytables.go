package logic

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	//AmmtersColumns .....
	AmmtersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "devid", Type: field.TypeString, Unique: false},
		{Name: "status", Type: field.TypeString},
		{Name: "devname", Type: field.TypeString},
		{Name: "timestamp", Type: field.TypeString, Unique: true},
		{Name: "ua", Type: field.TypeFloat32, Nullable: true},
		{Name: "ub", Type: field.TypeFloat32, Nullable: true},
		{Name: "uc", Type: field.TypeFloat32, Nullable: true},
		{Name: "uab", Type: field.TypeFloat32, Nullable: true},
		{Name: "ubc", Type: field.TypeFloat32, Nullable: true},
		{Name: "uca", Type: field.TypeFloat32, Nullable: true},
		{Name: "pha_i", Type: field.TypeFloat32, Nullable: true},
		{Name: "phb_i", Type: field.TypeFloat32, Nullable: true},
		{Name: "phc_i", Type: field.TypeFloat32, Nullable: true},
		{Name: "pha_p", Type: field.TypeFloat32, Nullable: true},
		{Name: "phb_p", Type: field.TypeFloat32, Nullable: true},
		{Name: "phc_p", Type: field.TypeFloat32, Nullable: true},
		{Name: "p", Type: field.TypeFloat32, Nullable: true},
		{Name: "pha_q", Type: field.TypeFloat32, Nullable: true},
		{Name: "phb_q", Type: field.TypeFloat32, Nullable: true},
		{Name: "phc_q", Type: field.TypeFloat32, Nullable: true},
		{Name: "q", Type: field.TypeFloat32, Nullable: true},
		{Name: "pha_s", Type: field.TypeFloat32, Nullable: true},
		{Name: "phb_s", Type: field.TypeFloat32, Nullable: true},
		{Name: "phc_s", Type: field.TypeFloat32, Nullable: true},
		{Name: "s", Type: field.TypeFloat32, Nullable: true},
		{Name: "pha_cos", Type: field.TypeFloat32, Nullable: true},
		{Name: "phb_cos", Type: field.TypeFloat32, Nullable: true},
		{Name: "phc_cos", Type: field.TypeFloat32, Nullable: true},
		{Name: "cos", Type: field.TypeFloat32, Nullable: true},
		{Name: "pha_u_ang", Type: field.TypeFloat32, Nullable: true},
		{Name: "phb_u_ang", Type: field.TypeFloat32, Nullable: true},
		{Name: "phc_u_ang", Type: field.TypeFloat32, Nullable: true},
		{Name: "pha_i_ang", Type: field.TypeFloat32, Nullable: true},
		{Name: "phb_i_ang", Type: field.TypeFloat32, Nullable: true},
		{Name: "phc_i_ang", Type: field.TypeFloat32, Nullable: true},
		{Name: "f", Type: field.TypeFloat32, Nullable: true},
		{Name: "forward_kwh", Type: field.TypeFloat32, Nullable: true},
		{Name: "reverse_kwh", Type: field.TypeFloat32, Nullable: true},
		{Name: "forward_kvarh", Type: field.TypeFloat32, Nullable: true},
		{Name: "reverse_kvarh", Type: field.TypeFloat32, Nullable: true},
		{Name: "kwh", Type: field.TypeFloat32, Nullable: true},
	}
	// AmmtersTable holds the schema information for the "gateways" table.
	AmmtersTable = &schema.Table{
		Name:       "ammeters",
		Columns:    AmmtersColumns,
		PrimaryKey: []*schema.Column{AmmtersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "ammter_devid",
				Unique:  true,
				Columns: []*schema.Column{AmmtersColumns[1]},
			},
			{
				Name:    "ammter_devname",
				Columns: []*schema.Column{AmmtersColumns[3]},
			},
			{
				Name:    "ammter_timestamp",
				Columns: []*schema.Column{AmmtersColumns[4]},
			},
		},
	}
	myTables = []*schema.Table{AmmtersTable}
)
