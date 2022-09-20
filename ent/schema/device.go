package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("dev_id").
			Unique(),
		field.String("type"),
		field.JSON("conn", map[string]any{}),
		field.Int("read_interval").Optional(),
		field.Int("store_interval").Optional(),
		field.String("name").Optional(),
		field.Bool("delete_flag").Optional(),
		field.String("summary").Optional(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization_position", OrganizationPosition.Type).
			Ref("devices").
			Unique(),
		edge.From("gateway", Gateway.Type).
			Ref("devices").
			Unique(),
	}
}

// Mixin of the Dev
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Indexes of the Device
func (Device) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("dev_id"),
	}
}
